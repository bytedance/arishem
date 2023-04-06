# Arishem Rule运算与自定义优先级

## 轻量接口

如果你仅仅是想使用arishem的规则判断能力，arishem提供了轻量级接口JudgeCondition来提供支持：

- JudgeCondition

```go
func main() {
	condition := `{
    "OpLogic": "&&",
    "Conditions": [
        {
            "Operator": "==",
            "Lhs": {
                "FuncExpr": {
    				"FuncName": "GetCurrentYear"
				}
            },
            "Rhs": {
                "Const": {
                    "NumConst": 2023
                }
            }
        }
    ]
}`
	pass, err := arishem.JudgeCondition(condition)
	assert.Nil(t, err)
	assert.True(t, pass)
}
```

- JudgeConditionWithFactMeta

```go
func main() {
	condition := `{"OpLogic":"&&","Conditions":[{"Operator":"==","Lhs":{"Const":{"NumConst":1}},"Rhs":{"Const":{"NumConst":1}}},{"Operator":"!LIST_IN","Lhs":{"VarExpr":"user.user_list#2.name"},"Rhs":{"ConstList":[{"StrConst":"Jack"},{"StrConst":"Jane"},{"StrConst":"John"},{"StrConst":"Ezreal"}]}}]}`
	pass, err := JudgeConditionWithFactMeta(expr, `{"user":{"user_list":[{"name":"Aatrox"},{"name":"Ahri"},{"name":"Ezreal"},{"name":"MalPhite"}]}}`)
	assert.Nil(t, err)
	assert.False(t, pass)
}
```

## 规则运算

arishem内置两种类型的规则，具有优先级和不具有优先级。其中具有优先级的规则在运算时会进行排序，然后再依次做运算，优先级高的规则通过后，便不再继续往下执行。而不具有优先级的规则都会得到执行。

> 注：规则运算如果有错误，默认认为该条规则判断未通过，继续执行下一条规则，如果想要获取错误，参考[自定义监听回调部分](./CALLBACK_zh.md)

- 优先级规则

通过NewPriorityRule创建一条具有优先级的规则，传入ruleName，优先级、条件表达式和目的表达式。

```go
rule, err := NewPriorityRule("p-rule3", 3, `{"OpLogic":"&&","Conditions":[{"Operator":"==","Lhs":{"Const":{"NumConst":99.999999}},"Rhs":{"Const":{"NumConst":100}}}]}`, `{"Const":{"StrConst":"p-rule3 passed!"}}`)
```

- 非优先级规则

通过NewNoPriorityRule创建一条不具有优先级的规则，传入ruleName，条件表达式和目的表达式。

```go
rule, err := NewNoPriorityRule("p-rule3", `{"OpLogic":"&&","Conditions":[{"Operator":"==","Lhs":{"Const":{"NumConst":99.999999}},"Rhs":{"Const":{"NumConst":100}}}]}`, `{"Const":{"StrConst":"p-rule3 passed!"}}`)
```

- 运算规则

通过调用ExecuteSingleRule来执行单条规则

```go
ExecuteSingleRule(rule, dc)
```

通过调用ExecuteRules来执行多条规则，规则可以时带有优先级的或者是非优先级的，也可以时二者混合的：

```go
ExecuteRules(rules, dc)
```

注：当具有优先级和不具有优先级的规则混合执行时，具有优先级的默认执行到命中(或最后一条)时不再继续向下执行，而不具有优先级的规则将全部得到执行。

## 自定义优先级

实现RuleTarget接口，在规则运算时实现实时带权重的优先级计算

```go
type WeightsPriorityRule struct {
    *arishem.NoPriorityRule
    weight   float64
    priority float64
}

func (w *WeightsPriorityRule) Compare(other typedef.Comparable) int {
    o := other.(*WeightsPriorityRule)
    myWeighted := w.weight * w.priority
    otherWeighted := o.weight * o.priority
    if myWeighted < otherWeighted {
        return -1
    } else if myWeighted == otherWeighted {
        return 0
    } else {
        return 1
    }
}
```