# Arishem

- Arishem是由字节跳动客服平台架构组自研的一款**轻量、高性能**
  的DSL规则引擎。目的是将变更频繁的业务决策从应用程序中剥离出来，使用可视化界面灵活地编写业务决策，提升业务需求的响应速度。
- Arishem采用完全兼容的JSON语法格式来定义规则语法，通过**组装、嵌套**
  的方式可以灵活地表达业务规则，使用Arishem，可以很容易地使用页面将规则进行可视化，使得不具备编程基础的人员也能快速上手。
- Arishem内部从AST解析生成和规则执行都进行了一系列的优化，使得单个复杂规则的执行可在**µs级别**的时间内完成。
- Arishem支持可**自定义**的规则执行顺序和**并发执行粒度**，支持**运行时下游数据的并发和预测获取**。
- Arishem内部集成了丰富的操作符和内置函数。

## 完全兼容JSON语法

Arishem由条件表达式和目的表达式构成，其最大的特点就是其语法**完全兼容JSON语法**，也能完全兼容IDL（如thrift、protobuf等）,在[examples](./examples/idls)部分定义了thrift和protobuf进行规则定义的样例。

arishem在规则应用到**可视化场景(如规则配置页面)具有非常显著的优势**，在字节跳动的客服平台，有关规则判断的几十个场景都接入了arishem，并实现了产运的可视化规则配置页面，使得规则配置这种原本需要研发去维护的场景，变成产运同学也能进行配置维护，大大解放了研发资源。

## 非常快！

在benchmark下执行单个无网络请求的复杂规则的耗时仅在微秒级别！当然没有最好性能的框架，只有最适合的应用场景。

测试环境：

```go
goos: darwin
goarch: amd64
pkg: */arishem/arishem
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
PASS
```

测试条件样例，共495个AST节点，包含基本的逻辑操作，包括字符串正则、数组遍历、数组交集等：

```json
// 实时数据
{"username":"Andrew","usernames":["Jack","Mike","Andrew"],"news":"Jack hanged out with Mike last weekend.","number1":100,"numbers":["10",99.9,0]}
// 条件表达式
{"OpLogic":"||","ConditionGroups":[{"OpLogic":"&&","Conditions":[{"Operator":"STRING_START_WITH","Lhs":{"VarExpr":"username"},"Rhs":{"Const":{"StrConst":"Banana"}}},{"Operator":"STRING_END_WITH","Lhs":{"VarExpr":"usernames#1"},"Rhs":{"Const":{"StrConst":"A"}}},{"Operator":"STRING_END_WITH","Lhs":{"VarExpr":"usernames#0"},"Rhs":{"Const":{"StrConst":"hahaha"}}},{"Operator":"CONTAIN_REGULAR","Lhs":{"VarExpr":"usernames#0"},"Rhs":{"Const":{"StrConst":"^M.*"}}}]},{"OpLogic":"and","ConditionGroups":[{"OpLogic":"&&","ConditionGroups":[{"OpLogic":"||","Conditions":[{"Operator":"<=","Lhs":{"Const":{"NumConst":100}},"Rhs":{"Const":{"NumConst":10}}},{"Operator":"<=","Lhs":{"Const":{"NumConst":100}},"Rhs":{"Const":{"NumConst":10}}},{"Operator":"<=","Lhs":{"Const":{"NumConst":100}},"Rhs":{"Const":{"NumConst":10}}},{"Operator":"<=","Lhs":{"Const":{"NumConst":100}},"Rhs":{"Const":{"NumConst":10}}},{"Operator":"<","Lhs":{"Const":{"NumConst":100}},"Rhs":{"Const":{"NumConst":10}}}]}]},{"OpLogic":"&&","Conditions":[{"Operator":"LIST_IN","Lhs":{"VarExpr":"number1"},"Rhs":{"ConstList":[{"NumConst":1},{"NumConst":98},{"NumConst":101},{"NumConst":1.32e-3},{"NumConst":12.234},{"NumConst":-1}]}},{"Operator":"!LIST_IN","Lhs":{"VarExpr":"numbers#0"},"Rhs":{"ConstList":[{"NumConst":1},{"NumConst":99},{"NumConst":999},{"NumConst":1.32e-3},{"NumConst":10},{"NumConst":-1}]}},{"Operator":"LIST_CONTAINS","Lhs":{"VarExpr":"numbers"},"Rhs":{"MathExpr":{"OpMath":"+","ParamList":[{"Const":{"NumConst":6}},{"Const":{"StrConst":"5"}}]}}},{"Operator":"LIST_RETAIN","Lhs":{"VarExpr":"numbers"},"Rhs":{"MathExpr":{"OpMath":"+","ParamList":[{"Const":{"NumConst":6}},{"Const":{"StrConst":"5"}}]}}},{"Operator":"LIST_RETAIN","Lhs":{"VarExpr":"numbers"},"Rhs":{"ConstList":[{"StrConst":"-1"},{"BoolConst":true},{"NumConst":-3.1415926}]}},{"Operator":"!LIST_RETAIN","Lhs":{"VarExpr":"numbers"},"Rhs":{"ConstList":[{"StrConst":"-1"},{"BoolConst":true},{"NumConst":-3.1415926}]}}]}]},{"OpLogic":"&&","Conditions":[{"Operator":"==","Lhs":{"Const":{"NumConst":1}},"Rhs":{"Const":{"NumConst":1}}}]}]}
// 目的表达式
{"ActionName":"Greeting2","ParamMap":{"UserAge":{"VarExpr":"user.age"},"TempUsername":{"VarExpr":"user.name"}}}
```

运行结果：平均不到30µs

```go
goos: darwin
goarch: amd64
pkg: */arishem/arishem
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz

BenchmarkSingleComplexRule-12     205094             28959 ns/op            2227 B/op         69 allocs/op
BenchmarkSingleComplexRule-12     208012             28637 ns/op            2226 B/op         69 allocs/op
BenchmarkSingleComplexRule-12     207696             28614 ns/op            2226 B/op         69 allocs/op
PASS
```

## 快速开始

注意：目前可使用的版本是 v1.0.0-rc1和v1.0.0-rc2，其他版本的存在包可见性问题，建议使用v1.0.0-rc2版本
```shell
go get github.com/bytedance/arishem@{version}
```

在使用前必须先调用Initialize方法，否则执行将导致arishem执行异常，一般情况下使用default配置即可，该操作应当在你的init方法中进行。

```go
func init() {
    arishem.Initialize(arishem.DefaultConfiguration())
}
```

- 快速判断一个条件表达式是否符合匹配的预期
[arishem语法参考](./doc/zh/INDEX_zh.md)
```go
func main() {
    condition := `
    {
        "OpLogic": "&&",
        "Conditions": [
            {
                "Operator": "==",
                "Lhs": {
                    "Const": {
                        "NumConst": 1
                    }
                },
                "Rhs": {
                    "Const": {
                        "NumConst": 1
                    }
                }
            }
        ]
    }
    `
    pass, err := arishem.JudgeCondition(condition)
    if err != nil {
        // handle error here
        // ...
        println(err.Error())
    }
    if pass {
        // your business code here
        // ...
        println("condition passed!")
    }
}
```
输出结果:
```go
condition passed!
````
- 在规则中通过关键字VarExpr获取实时数据进行判断
```go
func main() {
    condition := `
    {
        "OpLogic": "&&",
        "Conditions": [
            {
                "Operator": ">",
                "Lhs": {
                    "VarExpr": "user.age"
                },
                "Rhs": {
                    "VarExpr": "user_ages#1"
                }
            }
        ]
    }
    `
    pass, err := arishem.JudgeConditionWithFactMeta(condition, `
    {
        "user": {
            "name": "KJ",
            "age": 24
        },
        "user_ages": [
            15,
            20,
            32
        ]
    }
    `)
    if err != nil {
        // handle error here
        return
    }
    if pass {
        println("KJ's age is greater than 20!")
    }
}
```
- 创建一个规则并进行判断，输出规则目的
```go
func main() {
    condition := `
{
    "OpLogic": "&&",
    "Conditions": [
        {
            "Operator": ">=",
            "Lhs": {
                "VarExpr": "user.age"
            },
            "Rhs": {
                "VarExpr": "user_ages#1"
            }
        }
    ]
}
`
    // create an expression aim
    aim := `
{
    "Const": {
        "StrConst": "rule passed!"
    }
}
`
    rule, err := arishem.NewNoPriorityRule("rule1", condition, aim)
    if err != nil {
        // handle error here
        return
    }
    dc, err := arishem.DataContext(`
{
    "user": {
        "name": "KJ",
        "age": 24
    },
    "user_ages": [
        20,
        18,
        32
    ]
}
`)
    if err != nil {
        // handle error here
        return
    }
    rr := arishem.ExecuteSingleRule(rule, dc)
    if rr.Passed() {
        fmt.Printf("%s pass, output=>%s", rr.Identifier(), rr.Aim().AsExpr())
    }
}
```
输出结果
```go
rule passed!
```

或者通过内置的builder函数来构建条件表达式和目的表达式
```go
func main() {
    condGroup := arishem.NewConditionsCondGroup(arishem.OpLogicAnd)
    cond1 := arishem.NewCondition(operator.Equal)
    cond1.Lhs = arishem.NewConstExpr(arishem.NewNumConst(1.0))
    cond1.Rhs = arishem.NewConstExpr(arishem.NewNumConst(1.0))
    condGroup.AddConditions(cond1)
    expr, _ := condGroup.Build()
    println(expr)
}
```
输出结果
```json
{"OpLogic":"&&","Conditions":[{"Operator":"==","Lhs":{"Const":{"NumConst":1}},"Rhs":{"Const":{"NumConst":1}}}]}
```

更多使用方式请参考[详细文档](doc/zh/INDEX_zh.md)

## 可自定义的规则优先级定义和执行顺序

arishem另一个强大的功能就是支持灵活多变的规则执行顺序，arishem支持优先级规则执行、非优先级规则执行以及优先级和非优先级混合执行。甚至你可以通过自实现arishem规则接口，实现动态的优先级计算。相关使用方式请参考[详细文档](doc/zh/INDEX_zh.md)

## 丰富的操作符支持和内置函数

arishem支持多达20多种操作符，包括常用的值判断、数组判断，字符串判断等，并且arishem另一个特性就是强大的类型自动转换。
arishem在进行判断时，如果左指和右值的类型不一致，将尝试进行类型统一，在进行类型转换时，以右值的类型为标准进行转换。
更多使用方式请参考[详细文档](doc/zh/INDEX_zh.md)

arishem内部集成了一些常见的功能函数，包括日期、数组、map和字符串函数，并且支持注册自定义的函数到arishem中使用！使用[函数表达式](doc/zh/INDEX_zh.md)来使用他们。

## 支持网络数据的并发访问和预取

arishem将实时获取的网络数据定义为一个feature(特征)。arishem在执行规则的时候通过分批的方式进行规则运算，那么在feature获取的时候，也是通过批次去获取的。在使用具有网络数据的场景中，使用[FeatureExpr](doc/zh/INDEX_zh.md)关键字来使用，并实现feature的获取方法。
```json
{
  ...
  "FeatureExpr": {
    // user is the feature name, and username is the field path
    "FeaturePath": "user.username"
  }
}
```
```go
type MyFeatureFetcher struct{}

...

func (m *MyFeatureFetcher) FetchFeature(feat typedef.FeatureParam, dc typedef.DataCtx) (typedef.MetaType, error) {
    println("ready to fetch feature=>%s", feat.FeatureName())
    // you code here
    return nil, nil
}

func init() {
    arishem.Initialize(
        arishem.DefaultConfiguration(),
        arishem.WithFeatureFetcherFactory(func() typedef.FeatureFetcher {
            return &MyFeatureFetcher{}
        }),
    )
}
```
每个feature在获取时都是在一个异步协程中执行的，所以你可以没有顾虑地进行网络IO的访问，更多使用方式请参考[详细文档](doc/zh/INDEX_zh.md)。

## 灵活的自定义配置和监听回调

arishem支持多个配置的自定义，包括规则运算时的缓存实现、批大小的计算方式以及无优先级的最大并大数量等。

在规则运算的场景下，我们小组面临最多的问题便是排查规则为什么通过/没通过？规则运算过程是否有错误？feature获取的时候具体过程是怎么样的，究竟是哪个数据没有获取到？
所以感知规则运算的具体过程是非常有必要的，这将为后续排查规则命中详情提供基础能力的支持。

arishem支持规则匹配过的条件回调和错误回调，并且FeatureFetcher也必须实现observable方法，以便让arishem内部将feature fetch的过程通知给每一个已注册观察者。

```go
// MyObserver implements VisitObserver and FeatureFetchObserver
type MyObserver {}

func (m *MyObserver) OnFeatureFetchStart(feat typedef.FeatureParam) {}

func (m *MyObserver) OnFeatureFetchEnd(featureHash string, featureValue typedef.MetaType, err error) {}

func (m *MyObserver) OnJudgeNodeVisitEnd(info typedef.JudgeNode, vt typedef.VisitTarget) {}

func (m *MyObserver) OnVisitError(node, errMsg string, vt typedef.VisitTarget) {}

func main() {
    arishem.ExecuteSingleRule(rule, dataCtx, WithVisitObserver(myObserver), WithFetchObserver(myObserver))
}
```
有关自定义配置和监听回调的更多使用方式请参考[详细文档](doc/zh/INDEX_zh.md)。

## 开源许可

Arishem 基于[Apache License 2.0](LICENSE) 许可证。

## 联系&&问题反馈

Currently, we only support bug reports, summit your [ISSUES](https://github.com/bytedance/arishem/issues) here.
