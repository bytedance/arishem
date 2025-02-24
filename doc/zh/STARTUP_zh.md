# 快速开始

## 启动arishem

依赖arishem sdk后，必须调用arishem的Initialize方法后才能正常使用
```go
func init() {
    arishem.Initialize(arishem.DefaultConfiguration())
}
```

## 自定义配置
arishem支持多个配置的自定义，包括规则运算时的缓存实现、批大小的计算方式以及无优先级的最大并发数量等。在arishem初始化的时候传入你的自定义配置：

```go
func init() {
    arishem.Initialize(
        arishem.DefaultConfiguration(),
        arishem.WithDefTreeCache(),
        arishem.WithDefMaxParallels(),
        arishem.WithDefGranularity(),
        arishem.WithDefRuleComputePool(),
        arishem.WithDefFeatureFetchPool(),
        arishem.WithDefFeatFetcherFactory(),
        arishem.WithDefFeatVisitCache(),
        arishem.WithCustomListParamFuncs(
            arishem.ListParamFnPair{Name: "SwitchCase", Fn: SwitchCase},
            arishem.ListParamFnPair{Name: "TernaryOpt", Fn: TernaryOpt},
        ),
        arishem.WithEnableSubCondition("OnlyPriceMatch", `{"OpLogic":"&&","Conditions":[{"Operator":">","Lhs":{"VarExpr":"price"},"Rhs":{"Const":{"NumConst":10}}}]}`),
    )
}
```

#### WithDefTreeCache

TreeCache是arishem解析规则时使用的缓存，加上这个缓存后，arishem在解析到同一个条件/目的表达式时，将不再调用Antlr4的解析方法，直接从缓存中获取，默认256条。

如果要实现自定义的解析树缓存池，实现TreeCache接口即可
```go
type TreeCache interface {
    Put(expr string, tc *RuleTree)
    Get(expr string) (tc *RuleTree, ok bool)
}
```

#### WithDefMaxParallels

MaxParallels是arishem在运算非优先级规则时的最大并发度，arishem默认配置中，对非优先级的规则运算采用全部并发执行的配置，但不会超过MaxParallels(2048)。

#### WithDefGranularity

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

#### WithDefRuleComputePool/WithDefFeatureFetchPool

RuleComputePool/FeatureFetchPool分别是arishem在规则运算时和Feature获取时使用的协程池。

要使用自定义的协程池，请实现ConcurrentPool
```go
// Task execute mission with the parameter
type Task func(param interface{})

type ConcurrentPool interface {
    Submit(t Task, p interface{})
    SubmitWait(t Task, p interface{})
}
```

#### WithFeatFetcherFactory

FeatFetcherFactory是用户在使用了featureExpr后必须自实现的FeatureFetcher工厂方法。

要使用自定义的协程池，请实现ConcurrentPool
```go
// Task execute mission with the parameter
type Task func(param interface{})

type ConcurrentPool interface {
    Submit(t Task, p interface{})
    SubmitWait(t Task, p interface{})
}
```

#### WithFeatVisitCache
规则依赖的Feature会在在表达式解析阶段进行收集，同时该功能也会有缓存，默认大小是512个

要使用自定义的解析缓存，请实现SharedVisitCache接口
```go
// SharedVisitCache is the interface that can pass data among multi visitors.
type SharedVisitCache interface {
    // Get will get cached node visit result by its alternative number, Set and Get must concurrent safe.
    Get(key interface{}) (val interface{}, ok bool)
    // Set will set a visit result by its alternative number.
    Set(key interface{}, val interface{})
}
```

#### WithCustomNoParamFuncs <font color="red"><sup>new</sup></font>
在arishem启动时注册自定义的无参数Function，关于Function，[详见Expr部分](EXPR_zh.md)和[操作符部分](OPERATOR_FUNC_zh.md)
 
#### WithCustomMapParamFuncs <font color="red"><sup>new</sup></font>
在arishem启动时注册自定义的Map类型参数Function，关于Function，[详见Expr部分](EXPR_zh.md)和[操作符部分](OPERATOR_FUNC_zh.md)
 
#### WithCustomListParamFuncs <font color="red"><sup>new</sup></font>
在arishem启动时注册自定义的List类型参数Function，关于Function，[详见Expr部分](EXPR_zh.md)和[操作符部分](OPERATOR_FUNC_zh.md)

#### WithEnableSubCondition <font color="red"><sup>new</sup></font>
启用arishem对子条件判断的支持，并采用arishem默认的子查询，在使用此方法时，可以传入 cond1_Name、cond1_expr、cond2_name、cond2_expr...来注册一些初始化子条件

如果要使用自定义的子条件管理，请实现SubConditionManage
```go
type SubConditionManage interface {
	// WhenConditionParsed will be called when a condition expression successfully parsed
	WhenConditionParsed(condName, expr string, tree antlr.ParseTree)
	// GetConditionTree defines arishem how to get the condition parse tree by the condition name/key
	GetConditionTree(condName string) (antlr.ParseTree, error)
	// RuleIdentityMapAsCondName if this function return true,
	// then every rule identity will be the condition name stored into cache when parse rule or condition
	RuleIdentityMapAsCondName() bool
}
```
1. 向arishem添加一个或多个子条件
```go
err := arishem.AddSubCondition(${name1}, ${expression1}, ${name2}, ${expression2} ...)
```
2. SubConditionManage方法说明
- GetConditionTree -> 定义arishem如何通过条件名称来获取该条件的解析树
- RuleIdentityMapAsCondName -> 告诉arishem在解析规则时，是否将规则的identity作为其条件的名称进行存储, arishem默认为true
- WhenConditionParsed -> 该方法会在调用AddSubCondition方法成功解析一段条件表达式后回调，或者在启用了RuleIdentityMapAsCondName的情况下，成功解析规则后进行回调，在这个方法里，业务可以自定实现条件名称和解析树的查找关系
