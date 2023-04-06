# Arishem Builder

在进行arishem的条件或目的表达式的编写时，可以使用arishem内置的builder方法。

## 创建条件表达式

使用NewConditionsCondGroup创建条件列表或使用NewCondGroupsCondGroup创建条件组的条件表达式：

```go
func main() {
    condGroup := NewConditionsCondGroup(OpLogicAnd)
    cond1 := NewCondition(operator.Equal)
    cond1.Lhs = NewConstExpr(NewNumConst(1.0))
    cond1.Rhs = NewConstExpr(NewNumConst(1.0))
    condGroup.AddConditions(cond1)
    expr, err := condGroup.Build()
    assert.Nil(t, err)
    assert.NotEmpty(t, expr)
    assert.Equal(t, `{"OpLogic":"&&","Conditions":[{"Operator":"==","Lhs":{"Const":{"NumConst":1}},"Rhs":{"Const":{"NumConst":1}}}]}`, expr)

    cond2 := NewCondition(operator.ListIn, NOT)
    var varPath VarExpr = "user.user_list#2.name"
    cond2.Lhs = NewVarExpr(&varPath)
    cond2.Rhs = NewConstListExpr([]*Const{
        NewStrConst("Jack"),
        NewStrConst("Jane"),
        NewStrConst("John"),
        NewStrConst("Ezreal"),
    })
    condGroup.AddConditions(cond2)
    expr, err = condGroup.Build()
    assert.Nil(t, err)
    assert.NotEmpty(t, expr)
}
```

## 创建目的表达式
```go
func main() {
    aimStr, err := NewParamListAction("MyAction", []*Expr{NewConstExpr(NewNumConst(1.23))}).Build()
    assert.Nil(t, err)
    t.Log(aimStr)
    aim, err := parser.ParseArishemAim(aimStr)
    assert.Nil(t, err)
    assert.NotNil(t, aim)
    aimStr, err = NewParamMapAction("MyAction", map[string]*Expr{"rate": NewConstExpr(NewNumConst(1.23))}).Build()
    assert.Nil(t, err)
    t.Log(aimStr)
    aim, err = parser.ParseArishemAim(aimStr)
    assert.Nil(t, err)
    assert.NotNil(t, aim)
}
```