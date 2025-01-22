# Arishem自定义配置和回调

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

    OnFeatureFetchStart(feat FeatureParam)
    OnFeatureFetchEnd(featureHash string, featureValue MetaType, err error)
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


- 新的judgeNode类型：ForeachJudgeNode
如果你使用了arishem FOREACH 操作符的能力，你可能需要在判断回调中，得到foreach操作过程的具体详情，如: foreach的操作符、逻辑表达式, 每个item的值是什么以及它是否符合当前foreach的判断. 
你可以在OnJudgeNodeVisitEnd回调中，判断JudgeNode的具体类型，来尝试获取Foreach的记录
```go
func (t *testArisVisitObserver) OnJudgeNodeVisitEnd(ctx context.Context, info typedef.JudgeNode, vt typedef.VisitTarget) {
    foreachNode, ok := info.(typedef.ForeachJudgeNode)
    if ok {
        itemsStr, _ := json.Marshal(foreachNode.ForeachItems())
        logger.Infof("[ForeachJudgeNode] left: %v, leftExpr: %v, right: %v, rightExpr: %v, operator: %v, target: %v, err: %v, foreachOperator: %v, foreachLogic: %v, foreachItems: %v",
            foreachNode.Left(), info.LeftExpr(), foreachNode.Right(), info.RightExpr(), foreachNode.Operator(), vt.Identifier(), foreachNode.Error(),
            foreachNode.ForeachOperator(), foreachNode.ForeachLogic(), string(itemsStr),
        )
    }
}
```