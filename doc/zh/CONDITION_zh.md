# Arishem Condition表达式

## 1. Condition类型
在了解了Expr类型后，就可以编写Condition的类型了，Condtion实际上是由一个操作符(Operator)+左值(Lhs)+右值(Rhs)
构成，支持的操作符号详见[支持的操作符和内置函数](./EXPR_zh.md)。示例如下：

```json
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
```

## 2. ConditionGroup类型
ConditionGroup类型就是arishem的条件类型，ConditionGroup类型由一个逻辑符(OpLogic)+条件列表(Conditions)或条件组列表(
ConditionGroups)构成(二选一)，支持的OpLogic有与、或、非：

| 逻辑操作符 | 关键字                 | 含义                                           |
| ---------- | ---------------------- | ---------------------------------------------- |
| 与         | AND(不区分大小写)、&&  | 所有条件之间的通过为与关系                     |
| 或         | OR(不区分大小写)、\|\| |                                                |
| 非         | NOT(不区分大小写)、!   | 取条件列表的第一个进行运算，对是否通过取非操作 |

- Conditions类型的ConditionGroup示例

Conditions类型的ConditionGroup由OpLogic和Conditions组成。

```json
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
```

- ConditionGroups的ConditionGroup示例

ConditionGroups类型的ConditionGroup由OpLogic和ConditionGroups组成，ConditionGroups里的每一个元素都是一个ConditionGroup，通过这种方式，arishem的条件可以组合并嵌套出很复杂的规则，但ConditionGroups必须至少有一个Conditions类型的ConditionGroup

```json
{
  "OpLogic": "&&",
  "ConditionGroups": [
    {
      "OpLogic": "&&",
      "ConditionGroups": [
        {
          "OpLogic": "&&",
          "ConditionGroups": [
            {
              "OpLogic": "||",
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
          ]
        }
      ]
    },
    {
      "OpLogic": "||",
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
  ]
}
```

