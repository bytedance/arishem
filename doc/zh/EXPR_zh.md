# Arishem Expr类型

arishem的语法完全兼容JSON，你可以在JSON编辑器中编辑arishem规则。arishem定义了**条件表达式**和**目的表达式**
，这两种表达式都是由一个或多个Expr通过一些关键字构成的，在构建条件表达式和目的表达式之前，我们先了解Expr的类型。

arishem的最小构成为Expr类型，在arishem中**必须由一对大括号包裹**，表示为 '{' + ExprKey + Expr类型 + '}'
，arishem支持的Expr类型和示例如下：

| 类型            | ExprKey     | 说明                                |
|---------------|-------------|-----------------------------------|
| ListExpr      | ListExpr    | 由多个Expr类型构成的数组                    |
| MapExpr       | MapExpr     | 由Key(string)+Value(Expr)构成的表达式map |
| ConstExpr     | Const       | 常量类型                              |
| ConstListExpr | ConstList   | 由多个常量构成的数组                        |
| MathExpr      | MathExpr    | 数学表达式类型                           |
| FuncExpr      | FuncExpr    | 使用arishem内置函数                     |
| VarExpr       | VarExpr     | 从实时数据中获取数据的路径表达式                  |
| FeatureExpr   | FeatureExpr | 从外部网络中获取数据的路径表达式                  |
| SubCondExpr <font color="red"><sup>new</sup></font>  | SubCondExpr | 右值是作为子条件进行运算的方式 |
| null          | 无           | 空类型的Expr                          |

按照拆解的原则，ListExpr和MapExpr都是Expr的特定组合，故先从最小的Expr类型进行说明。

## 0. Arishem合法Name类型

在arsihem支持的表达式如MapExpr中，会为某个值(名称)取名作为Key，arishem定义了这种key的合法格式，他的正则表达式如下：

```sh
[a-zA-Z][a-zA-Z0-9_#]*
```

即以[a-zA-Z]开头，[a-zA-Z0-9_#]其中任意字符取0个或多个，在[VarExpr](#3. VarExpr)和[FeatureExpr](#4. FeatureExpr)
表达式中，'#'后将紧接数字作为数组类型的下标索引。

## 1. ConstExpr

ConstExpr表示常量类型，由关键字"Const"表示，arishem支持3种常量类型：Bool类型、Number类型以及String类型。示例如下：

- Bool类型：由关键字"BoolConst"+bool值构成，bool值为true/false 二选一，大小写敏感，**'True/False'将会解析报错。**

```json
{
  "Const": {
    "BoolConst": true
  }
}
```

arishem解析后

```go
true
```

- Number类型：由关键字"NumConst"+number值构成，支持整型、浮点型、有无符号、科学计数法等表示方式。

```json
{
  "Const": {
    "NumConst": 1.234e1234
  }
}
```

arishem解析后

```go
1.234e1234
```

- String类型：由关键字"StrConst"+string值构成，string值由一对双引号包裹，StrConst支持所有unicode定义的字符串。

```json
{
  "Const": {
    "StrConst": "Strings:\n\b\t\r\f~!@#$%^&*()-=_+[]\\{}|;':\",.?<>/œ∑´®†¥¨ˆøπ“‘«åß∂ƒ©˙∆˚¬…æΩ≈ç√∫˜µ≤≥÷"
  }
}
```

arishem解析后

```go
"Strings:\n\b\t\r\f~!@#$%^&*()-=_+[]\\{}|;':\",.?<>/œ∑´®†¥¨ˆøπ“‘«åß∂ƒ©˙∆˚¬…æΩ≈ç√∫˜µ≤≥÷"
```

## 2. ConstListExpr

ConstListExpr表示**至少**由一个或多个Const组成，关键字为"ConstList"，由中括号包裹元素，示例如下：

```json
{
  "ConstList": [
    {
      "StrConst": "name"
    },
    {
      "NumConst": 1.234
    },
    {
      "BooConst": false
    }
  ]
}
```

arishem解析后

```go
["name", 1.234, false]
```

## 3. VarExpr <font color="red"><sup>new</sup></font>

VarExpr表示一个路径，该路径表示某个值在**实时数据**
中的取值路径。arishem的取值路径和JSON相似，但又有所不同。VarExpr的格式为a.b.c。每一个用'.'分割的元素都是实时数据中的某个value的key，
它的类型为[ Arishem合法Name类型](#0. Arishem合法Name类型)。其中'#'后紧跟数字，表示某个数组类型的取值下标，下标从0开始计数。

以一个实时数据为例：

```json
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
```

- 获取user.name的VarExpr示例：

```json
{
  "VarExpr": "user.name"
}
```

arishem解析后

```go
"KJ"
```

- 获取user_ages的第一个元素的值：

```json
{
  "VarExpr": "user_ages#0"
}
```

arishem解析后

```go
20
```

- 支持list类型的元素收集，语法为list元素后紧跟‘##’,只支持单个字段的收集

 如有这么一段数据
```json
{
    "item_info": {
        "item_list": [
            {
                "name": "name@1",
                "price": 100
            },
            {
                "name": "name@2",
                "price": 102.13
            },
            {
                "name": "name@3",
                "price": 200
            },
            {
                "name": "name@4",
                "price": 100
            },
            {
                "name": "name@5",
                "price": 101
            },
            {
                "name": "name@6",
                "price": 303.1234
            }
        ]
    }
}
```
收集里面的所有price字段
```json
{
    "VarExpr": "item_info.item_list##price"
}
```
那么arishem解析后为
```json
[100,102.13,200,100,101,303.1234]
```

## 4. FeatureExpr

arishem将依赖外部网络的数据定义为feature，使用FeatureExpr表示该数据从外部网络中获取，FeatureExpr由FeaturePath和BuiltinParam(
可选)
构成，在使用FeatureExpr时，必须实现FeatureFetcher接口，详见[实现网络数据的访问](./FEATURE_zh.md#实现featurefetcher接口)
。FeaturePath和VarExpr表达式的格式是一致的，但在解析时，第一个用'.'
分割的key将作为FeatureName并在FeatureFetcher中进行回调，剩余的path将会作为feature字段的取值路径。

```json
{
  "FeatureExpr": {
    // user is the feature name, age is the field path of user feature's data.
    "FeaturePath": "user.age"
  }
}
```

FeatureExpr支持内置参数(BuiltinParam)，内置参数表示的是Feature配置在规则中的一些固有参数，格式是Map格式的Expr，是在解析阶段可以计算出值的表达式，一般是常量或者规则右值的数据，如：

```json
{
  "FeatureExpr": {
    "FeaturePath": "user.age",
    "BuiltinParam": {
      "QueryIDRange": {
        "ConstList": [
          {
            "NumConst": 5678
          },
          {
            "NumConst": 1234
          }
        ]
      }
    }
  }
}
```

FeatureFetcher在被回调时，可以通过BuiltinParam()
方法获取配置在规则中的内置参数，详见[实现网络数据的访问](./FEATURE_zh.md#使用featureexpr)。

## 5. MathExpr

MathExpr由OpMath和计算的左值(Lhs)和右值(Rhs)构成，OpMath支持+、-、*、/、%，Lhs和Rhs都是一个Expr，例如：

```json
{
  "MathExpr": {
    "OpMath": "+",
    "Lhs": {
      "Const": {
        "NumConst": 10
      }
    },
    "Rhs": {
      "Const": {
        "NumConst": 8
      }
    }
  }
}
```

arishem在执行的时候，将会执行10+8的操作。

MathExpr支持表达式列表作为参数输入，其关键字为"ParamList"，例如：

```json
{
  "MathExpr": {
    "OpMath": "+",
    "ParamList": [
      {
        "Const": {
          "NumConst": 6
        }
      },
      {
        "Const": {
          "NumConst": 5
        }
      },
      {
        "Const": {
          "StrConst": 10
        }
      }
    ]
  }
}
```
arishem在执行时，将会执行6+5+10的操作。

## 6. FuncExpr

使用FuncExpr将指定使用arishem内置的函数进行判断，FuncExpr由FuncName+入参构成，入参支持无参、数组类型ParamList入参或者map类型的ParamMap入参(
3选1)。例如：

- 无参

```json
{
  "FuncExpr": {
    "FuncName": "GetCurrentYear" // 获取当前年份
  }
}
```

arishem将会得到当前年份，如2023。

- ParamList入参

```json
{
  "FuncExpr": {
    "FuncName": "ListAdd",
    "ParamList": [
      {
        "ConstList": [
          {
            "NumConst": 1
          },
          {
            "NumConst": 1
          }
        ]
      },
      {
        "Const": {
          "NumConst": 1
        }
      }
    ]
  }
}
```

ListAdd内置函数将会对所有的参数进行列表转换，并把元素进行相加，在本例中得到入参：[1,1,1]。

- ParamMap入参

```json
{
  "FuncExpr": {
    "FuncName": "ListLength",
    "ParamMap": {
      "list": {
        "ConstList": [
          {
            "NumConst": 1
          },
          {
            "NumConst": 1
          }
        ]
      },
      "filter_null": {
        "Const": {
          "BoolConst": true
        }
      }
    }
  }
}
```

ListLength内置函数将会计算入参中的list对象，filter_null表示是否过滤null值，在本例中得到list长度为2。

支持的函数详见[支持的操作符和内置函数](INDEX_zh.md#支持的操作符和内置函数)。

## 7. ListExpr

ListExpr表示由一个或多个Expr类型组成的表达式，其关键字为ListExpr，例如：

```json
{
  "ListExpr": [
    {
      "FeatureExpr": {
        "FeaturePath": "user.username"
      }
    },
    {
      "VarExpr": "record#1.record_name"
    },
    {
      "Const": {
        "NumConst": 1
      }
    }
  ]
}
```

## 8. MapExpr

MapExpr和ListExpr类型相似，表示有一个或多个由Expr类型组成的表达式，其关键字为MapExpr，每个Expr类型需要有一个key，key的合法形式是a-zA-Z开头，a-zA-Z0-9_
#中任意字符串组成，例如：

```json
{
  "MapExpr": {
    "UserName": {
      "VarExpr": "user.username"
    },
    "UserAge": {
      "MathExpr": {
        "OpMath": "+",
        "Lhs": {
          "Const": {
            "NumConst": 10
          }
        },
        "Rhs": {
          "Const": {
            "NumConst": 8
          }
        }
      }
    }
  }
}
```

## 9. SubCondExpr <font color="red"><sup>new</sup></font>
使用该表达式必须启用了arishem对子条件的判断支持，详见[快速开始的配置部分](STARTUP_zh.md)

该表达式表示一个子条件，只能作用于条件的右值表达式，并用CondName指定要使用子条件，举个具体的例子，要判断入参中的商品列表里的每一个商品都满足 名称包含xxx 并且 价格 大于100:

入参：
```json
{
    "item_info": {
        "item_list": [
            {
                "name": "name@1",
                "price": 100
            },
            {
                "name": "name@2",
                "price": 102.13
            },
            {
                "name": "name@3",
                "price": 200
            },
            {
                "name": "name@4",
                "price": 100
            },
            {
                "name": "name@5",
                "price": 101
            },
            {
                "name": "name@6",
                "price": 303.1234
            }
        ]
    }
}
```
规则的条件表达式:
```json
{
  "OpLogic": "&&",
  "Conditions": [
    {
      "Operator": "FOREACH SUB_COND and",
      "Lhs": {
        "VarExpr": "item_info.item_list"
      },
      "Rhs": {
        "SubCondExpr": {
          "CondName": "NameAndPriceMatch"
        }
      }
    }
  ]
}
```
在运算上述条件前，请确保你已经向arishem中注册了NameAndPriceMatch这样一个子条件
```json
{
    "OpLogic": "&&",
    "Conditions": [
        {
            "Operator": ">",
            "Lhs": {
                "VarExpr": "price"
            },
            "Rhs": {
                "Const": {
                    "NumConst": 10
                }
            }
        },
        {
            "Operator": "STRING_CONTAINS",
            "Lhs": {
                "VarExpr": "name"
            },
            "Rhs": {
                "Const": {
                    "StrConst": "name@"
                }
            }
        }
    ]
}
```
代码描述
```go
err := AddSubCondition("NameAndPriceMatch", ${sub condition expression})
assert.Nil(t, err)
pass, err := JudgeConditionWithFactMeta(context.Background(), ${this rule conition}, ${FactData})
assert.Nil(t, err)
assert.True(t, pass)
```
