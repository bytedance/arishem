# Arishem Aim表达式

arishem本身并不执行动作，它只返回某条通过的规则的解析结果。arishem内置了3种目的表达式，通过这3种表达式可以满足大多数业务场景需求。

## 1. Expr类型目的

一个Expr就可以作为一个目的表达式了，例如：

```json
{
  "Const": {
    "StrConst": "condition of this rule is passed, i'm the aim of it."
  }
}
```

arishem解析后将返回字符串: "condition of this rule is passed, i'm the aim of it."
通过访问Aim的AsExpr()方法得到arishem的解析结果
```go
func main() {
    rr := arishem.ExecuteSingleRule(rule, dc)
    if rr.Passed() {
        fmt.Printf("%s pass, output=>%s", rr.Identifier(), rr.Aim().AsExpr())
    }
}
```

## 2. Action类型目的

arishem内置了Action类型的目的表达式，其由关键字ActionName和ParamMap(或ParamList)构成，ActionName为[ Arishem合法Name类型](./EXPR_zh.md#0-arishem合法name类型)
，而ParamMap的实质就是MapExpr，ParamList的实质就是ListExpr。例如：
```json
{
  "ActionName": "Greeting",
  "ParamMap": {
    "Target": {
      "VarExpr": "user.name"
    }
  }
}
```
或者
```json
{
  "ActionName": "Greeting",
  "ParamList": [
    {
      "VarExpr": "user.name"
    },
    {
      "Const": {
        "StrConst": "Mike"
      }
    }
  ]
}
```
通过访问Aim的AsAction()方法得到aim作为Action的结果
```go
expr := aim.AsAction()
// get parameter
param := expr.Parameter()
// get param map
paramMap := expr.ParamAsMap()
// get param as list
paramList := expr.ParamAsList()

```

## 3. ActionAim列表
actionAim列表的aim表达式由一对中括号包裹，包含0个或多个Action类型的目的，如：
```json
[
  {
    "ActionName": "Greeting",
    "ParamList": [
      {
        "VarExpr": "user.name"
      },
      {
        "Const": {
          "StrConst": "Mike"
        }
      }
    ]
  },
  {
    "ActionName": "Greeting2",
    "ParamList": [
      {
        "FeatureExpr": {
          "FeaturePath": "user_profile.username",
          "BuiltinParam": {
            "user_id": {
              "Const": {
                "NumConst": 1234567
              }
            }
          }
        }
      }
    ]
  }
]
```