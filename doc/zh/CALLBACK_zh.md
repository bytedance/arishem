# Arishem自定义配置和回调

## 自定义配置

arishem支持多个配置的自定义，包括规则运算时的缓存实现、批大小的计算方式以及无优先级的最大并大数量等。在arishem初始化的时候传入你的自定义配置：

```go
func DefaultConfiguration() *core.Configuration {
	c := new(core.Configuration)
	core.ApplyOptions(c,
		WithDefTreeCache(),
		WithDefMaxParallels(),
		WithDefGranularity(),
		WithDefRuleComputePool(),
		WithDefFeatureFetchPool(),
		WithDefFeatFetcherFactory(),
		WithDefFeatVisitCache(),
	)
	return c
}
```

- WithDefTreeCache

TreeCache是arishem解析规则时使用的缓存，加上这个缓存后，arishem在解析到同一个条件/目的表达式时，将不再调用Antlr4的解析方法，直接从缓存中获取，默认512条。

- WithDefMaxParallels

MaxParallels是arishem在运算非优先级规则时的最大并发度，arishem默认配置中，对非优先级的规则运算采用全部并发执行的配置，但不会超过MaxParallels。

- WithDefGranularity

Granularity是arishem在运算多条规则时的分批算法，入参是当前的规则组的数量和运算的场景(优先级/非优先级)。默认配置如下：

```go

func granularity(ruleGroupSize int, model core.ExecuteModel) int {
	if model == core.ExecuteModelNoPriority {
		if ruleGroupSize <= arishemConfiguration.MaxParallel {
			return ruleGroupSize
		}
		return arishemConfiguration.MaxParallel
	}
  // dCpu is the double of cpu core num.
	if ruleGroupSize < dCpu {
		return ruleGroupSize
	}
	if ruleGroupSize < dCpu*2 {
		return ruleGroupSize / 2
	}
	if dCpu <= 1 { // when x = 1，2*√x intersects with 2*x
		return 1
	}
	return (dCpu*2 - factor) * 2
}
```

- WithDefRuleComputePool/WithDefFeatureFetchPool

RuleComputePool/FeatureFetchPool分别是arishem在规则运算时和Feature获取时使用的协程池。

- FeatFetcherFactory

FeatFetcherFactory是用户在使用了featureExpr后必须自实现的FeatureFetcher工厂方法。

- FeatVisitCache

Feature在规则表达式解析阶段也会有缓存，默认大小是512个

## 回调监听

arishem一共定义了两种回调类型，分别是规则运算回调和FeatureFetch过程的回调，他们的接口定义如下：

```go
// VisitObserver is the observer that can listen to rule visit events.
type VisitObserver interface {
	Hashable

	// OnJudgeNodeVisitEnd passes the condition node
	OnJudgeNodeVisitEnd(info JudgeNode, vt VisitTarget)
	OnVisitError(node, errMsg string, vt VisitTarget)
}

// FeatureFetchObserver can listen to feature fetch events.
type FeatureFetchObserver interface {
	Hashable

	OnFeaturesFetchStart(feat FeatureParam)
	OnFeaturesFetchEnd(featureHash string, featureValue MetaType, err error)
}
```

注意：回调函数中不要有耗时操作，否则会影响规则运算的性能和耗时。

- VisitObserver 规则运算回调

在规则运算时，通过WithVisitObserver方法传入你的observer

```go
ExecuteRules(rules, dc, WithVisitObserver(observer))
```

- FeatureFetchObserver

在使用前，务必在你的FeatureFether中实现AddObserver和GetFetchObservers方法，注意，AddFetchObserver和GetFetchObservers可能在并发环境下被调用，注意并发安全。

```go
func (m *MyFeatureFetcher) AddFetchObserver(v ...typedef.FeatureFetchObserver) {
	m.lock.Lock()
	defer m.lock.Unlock()
	for _, observer := range v {
		m.observers[observer.HashCode()] = observer
	}
}

func (m *MyFeatureFetcher) GetFetchObservers() []typedef.FeatureFetchObserver {
	m.lock.RLock()
	defer m.lock.RUnlock()
	var obs []typedef.FeatureFetchObserver
	for _, observer := range m.observers {
		obs = append(obs, observer)
	}
	return obs
}
```

在规则运算时，通过WithFeatureFetchObserver方法传入你的observer

```go
ExecuteRules(rules, dc,WithFeatureFetchObserver(observer))
```

