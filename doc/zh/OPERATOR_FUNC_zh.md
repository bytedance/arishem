# Arishem支持的操作符和内置函数

## 操作符

### 类型转换说明

Arishem另一个强大的特性就是根据操作符自动进行类型兼容，在基本操作符判断的时候，**如果左值的类型和右值的类型不匹配，那么arishem将会以右值类型为标准对左值进行转换。**

以下是arishem支持的操作符
- 基本操作符

| 运算符  | 中文名   | 说明                     | 举例                                                   |
| ------- | -------- | ------------------------ | ------------------------------------------------------ |
| ==      | 等于     | 支持数字和字符串和布尔值 | 1==1 或者 "a"=="a"，true/false和数字比较时将转换成0、1 |
| !=      | 不等于   | 支持数字和字符串和布尔值 | 1!=2 或者 "a"!="b"                                     |
| >       | 大于     | 支持数字和字符串和布尔值 |                                                        |
| >=      | 大于等于 | 支持数字和字符串和布尔值 |                                                        |
| <       | 小于     | 支持数字和字符串和布尔值 |                                                        |
| <=      | 小于等于 | 支持数字和字符串和布尔值  |                                                        |
| IS_NULL | 是否为空 | 支持任意类型             |                                                        |

- 其他操作符

arishem在进行列表相关的操作时，会默认尝试对左值和右值进行列表类型转换，转换失败判断将会返回false。该列表下所有的操作符在前面加上'NOT '或者'!'表示取非的逻辑，如'LIST_IN'取非操作就表示为'NOT LIST_IN'或'!LIST_IN'。

| 运算符                                                 | 中文名                           | 说明                                                                                                             | 举例                                                                                                   |
|-----------------------------------------------------|-------------------------------|----------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------|
| LIST_IN                                             | list包含                        | 支持数字和字符串和布尔值；真实含义是：在...中                                                                                       | A在[A,B,C]中                                                                                           |
| LIST_RETAIN                                         | list有交集                       | 支持数字和字符串和布尔值；两个list校验是否有交集，在进行比较时，布尔值将被转换成1(true)和0(false)。                                                    | [A,B,C]跟[B,C,D]有交集                                                                                   |
| LIST_CONTAINS                                       | list包含                        | 一个list包含某一项                                                                                                    | [A,B,C]包含C                                                                                           |
| STRING_START_WITH                                   | 字符串以X开头                       |                                                                                                                | "app"以"a"开头                                                                                          |
| STRING_END_WITH                                     | 字符串以X结尾                       |                                                                                                                | "app"不以"b"开头                                                                                         |
| STRING_CONTAINS                                     | 字符串包含                         | 效率很差，不建议用                                                                                                      | "access"包含"e"                                                                                        |
| CONTAIN_REGULAR                                     | 字符串正则匹配                       | 不建议大量使用，性能不佳                                                                                                   | 正则"鲜花\|鲜花饼\|满地鲜花"与字符串"鲜花"匹配                                                                          |
| SUB_LIST_IN                                         | 子序列(左值)包含于(右值)                | 支持数字和字符串；真实含义是：左值 是 右值的一个子列表(序列)                                                                               | [A,B][A,C]是[A,B,C]的一个子列表[A]是[A]的一个子列表[A,M]不是[A,B,C]的一个子列表                                            |
| SUB_LIST_CONTAINS                                   | (左值)包含子序列(右值)                 | 支持数字和字符串；真实含义是：左值 包含了 右值这个子列表(序列)                                                                              | [A,B,C]包含了[A,B]这个子列表[A,B]不包含[A,B,C]这个子列表[A,B]不包含[M,N]这个子列表                                           |
| BETWEEN_ALL_CLOSE                                   | 左值在右值区间，左右闭                   | 右值为数组类型，元素个数等于2，默认第一个元素为左区间元素，第二个为右区间元素，当右值数组元素大于2时，取第一个和第二个进行判断                                               | 1 在 [1,5] 之间，1 不在 [2,5] 之间                                                                           |
| BETWEEN_ALL_OPEN                                    | 左值在右值区间，左右开                   |                                                                                                                | 1 在 (0, 3) 之间1 不在 (1, 5) 之间                                                                          |
| BETWEEN_LEFT_OPEN_RIGHT_CLOSE                       | 左值在右值区间，左开右闭                  |                                                                                                                | 1 在 (0, 3] 之间1 不在 (1, 3] 之间3 在 (1, 3] 之间                                                             |
| BETWEEN_LEFT_CLOSE_RIGHT_OPEN                       | 左值在右值区间，左闭右开                  |                                                                                                                | 1 在 [1, 3) 之间2 在 [1, 3) 之间3 不在 [1, 3) 之间                                                             |
| SUB_COND    <font color="red"><sup>new</sup></font> | 左值满足子条件                       | 适用于对左值做复杂判断的场景，左值必须是object类型，右值必须是SubCondExpr来指定其运算子条件, 子条件运行时，入参为左值，与当前规则运算的入参数**互相独立**                       | 左值: {"name": "Mike", age: 20} 满足子条件 `MikeIsAdult` name == "Mike" && age >= 18, 详见[表达式部分](EXPR_zh.md) |
| FOREACH ${Operator} ${OpLogic} <font color="red"><sup>new</sup></font>                     | 左值的${每一项\|任一项} 满足 ${Operator} | 适用于左值是数组类型时，对左值的每一个元素做单独判断的场景，只支持单层，不支持多层FOREACH嵌套, ${Operator}是arishem支持的任意一个操作符，${OpLogic}是arishem支持的逻辑与或操作符 | 左值: [1,2,3,4] 满足 `FOREACH > &&` 右值: 0 的含义是：左值里的每一项大于0,详见[表达式部分](EXPR_zh.md)                          |
## 内置函数

arishem内置函数分为3类，不需要参数的func、参数为列表的func以及参数为map的func。

- 不需要参数的func

| funcName           | 返回值类型 | 函数说明                     | 返回值示例                                                   |
| ------------------ | ---------- | ---------------------------- | ------------------------------------------------------------ |
| GetUnixSecond      | int        | 获取当前Unix时间戳，单位为秒 | 1680528585                                                   |
| RandomUUID         | string     | 获取一个随机的UUID           | d9564ecc-020a-1022-4894-2711ee18ded3                         |
| GetCurrentYear     | int64      | 获取当前的年份               | 2023                                                         |
| GetCurrentYearDay  | int64      | 获取当天为当前年的第几天     | 当前时间为1月4日，返回4                                      |
| GetCurrentMonth    | int64      | 获取当前月份                 | 1                                                            |
| GetCurrentMonthDay | int64      | 获取当天为该月的第几天       | 当前时间为1月6日，返回6                                      |
| GetCurrentWeekDay  | string     | 获取当天为星期几             | 可用枚举为：SundayMondayTuesdayWednesdayThursdayFridaySaturday |
| GetCurrentMinute   | int64      | 获取当前分钟                 | 当前时间为10:24，返回24                                      |
| GetCurrentSecond   | int64      | 获取当前的秒数               | 当前时间为10:25:23，返回23                                   |
| GetCurrentLocation | string     | 获取当前所在地区             | LocalAsia/ShanghaiUTC                                        |

- 参数为列表的func

| funcName         | 参数说明                                             | 返回值类型             | 函数说明                                                     | 返回值示例                                                   |
| ---------------- | ---------------------------------------------------- | ---------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| MinFlatWithInt   | 参数为数组，每个元素类型为单个元素或者元素数组       | int64                  | 所有的参数打平求最小值，每个元素会先被转换成int64            | [1, [3,4,5], 6, [0,1,2]] -> 0[1, [0], [2,4]] -> 0            |
| MinFlatWithFloat | 参数为数组，每个元素类型为单个元素或者元素数组       | float64                | 所有的参数打平求最小值，每个元素会先被转换成float64          | [1, [0.2, 2.3, 4], 0.1, [1,2]] -> 0.1                        |
| MaxFlatWithInt   | 参数为数组，每个元素类型为单个元素或者元素数组       | int64                  | 所有的参数打平求最大值，每个元素会先被转换成int64            | [1, [3,4,5], 6, [0,1,2]] -> 6[1, [0], [2,4]] -> 4            |
| MaxFlatWithFloat | 参数为数组，每个元素类型为单个元素或者元素数组       | float64                | 所有的参数打平求最大值，每个元素会先被转换成float64          | [1, [0.2, 2.3, 4], 0.1, [1,2]] -> 4.0                        |
| ListFlatJoin     | 参数为数组，每个元素类型为单个元素或者元素数组       | string                 | 所有的参数打平连接，参数数组第一个为连接符号                 | ["-",1.23,"zhangsan",3,true] -> "1.23-zhangsan-3-true"[",",1.23,"zhangsan",3,true] -> "1.23,zhangsan,3,true"注：数组第一个元素为连接符 |
| ListAdd          | 参数为数组，每个元素类型为单个元素或者元素数组       | []interface{}          | 参数数组第一个参数默认为数组类型，然后与后面的参数相加       | [[1, 2, 3],[1.0, 2.0]] -> [1, 2, 3, 1.0, 2.0]["zhangsan", [1, 2, 3], [1.0, 2.0, 3.1]] -> ["zhangsan", 1, 2, 3, 1, 2, 3.1] |
| ListRemove       | 参数为数组，每个元素类型为单个元素或者元素数组       | []interface{}          | 参数数组第一个参数默认为数组类型，然后与后面的参数相减**注：Number类型会统一转换成uint64/int64/float64类型，然后做差，不同类型数据无法做差** | [[1, 2, 3],[1.0, 2.0]] -> [3.0][[1, 2, 3],[1.0, 2.0,3.1]] -> [3.0][["zhangsan"], [1,2,3]] -> ["zhangsan"] |
| ListUnion        | 参数为数组，每个元素类型为单个元素或者元素数组       | []interface{}          | 数组求并集**注：Number类型会统一转换成uint64/int64/float64类型，然后做同值判断，指针和结构体类型会通过序列化成string类型做同值判断，Number类型和string类型/指针/结构体之间无法比较。** | [[1, 2, 3],[1.0, 2.0]] -> [1.0,2.0,3.0][[1, 2, 3],[1.0, 2.0, 3.1]] -> [1.0,2.0,3.0,3.1][[1, 2, 3],[1.0, 2.0, 3.1],["zhang","san"]] -> [1.0,2.0,3.0,3.1,"zhang","san"] |
| ListIntersect    | 参数为数组，每个元素类型为单个元素或者元素数组       | []interface{}          | 数组求交集**注：Number类型会统一转换成uint64/int64/float64类型，然后做同值判断，指针和结构体类型会通过序列化成string类型做同值判断，Number类型和string类型/指针/结构体之间无法比较。** | [[1, 2, 3],[1.0, 2.0]] -> [1.0, 2.0][[1, 2, 3],[1.0, 2.0, 3.1]] -> [1.0, 2.0][[1, 2, 3],[1.0, 2.0, 3.1],["1", "zhang", "san"]] -> [] （返回一个空数组，因为三个数组之间没有交集，Number与string类型之间无法比较）["1", "zhang", "san"], ["1", "zhang", "gou"]] -> ["1", "zhang"][]*nameAge{{Name: "zhangsan", Age: 18}, {Name: "lisi", Age: 20}} 和 []*nameAge{{Name: "wangwu", Age: 18}, {Name: "lisi", Age: 20}}求交集，会返回一个string数组，每个元素为结构体序列化后的值，如当前例子返回["{"Name":"lisi","Age":20}"] |
| MapUnion         | 参数为数组类型，每个元素为map[string]interface{}类型 | map[string]interface{} | map求并集，**key值的同值判断与上述数组求并/交集一致**        | [{"name": "1"},{"name": "3"},{"name": "10", "age": 18}] -> {"name":"10","age":18} |
| MapIntersect     |                                                      | map[string]interface{} | map求交集，**key值的同值判断与上述数组求并/交集一致**        | [{"name": "1"},{"name": "3"},{"name": "10", "age": 18}] -> {"name":"1"} |

- 参数为Map的func

| funcName                  | 参数说明                                                     | 返回值类型 | 函数说明                                                     | 返回值示例                                                   | 备注                                                         |
| ------------------------- | ------------------------------------------------------------ | ---------- | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| ListLength            | 传入两个参数，分别为list和filter_null(可选)，list为数组类型，filter_null为bool类型，默认不传为false，表示是否过滤数组中的null元素 | int        | 计算targetList的长度，当filterNull为true时，过滤数组中的空元素 | list:"zhangsan" -> 0[1,2,3,4] -> 4["1","2",null,"4"],filter_null为true时 -> 3filterNull不传或者为false时->4 |                                                              |
| DateToUnix                | 四个参数：<br />1.date：格式为yyy-MM-dd HH:mm:ss的日期字符串 <br />2. precision：表示日期字符串的精度，和date格式有关，当date格式精准到秒时，填sec，以此类推，可选的有：day、min、sec、milli<br />3. zone(可选)：时间区域，不传默认为local，剩下可传的为utc、shanghai，或者自行传入符合golang语言的区域格式，如：Asia/Shanghai、Europe/Berlin<br />4. unit(可选)：unix时间戳返回的单位，默认为秒，其余可选的为：sec、milli、micro、nano | int64      | 在参数map中，传入日期格式的字符串**date**，**date的格式与precision保持一致**，函数会将该date转换成日期，并返回该日期的unix时间戳，可以通过传入zone改变date转换成日期时的时区，不传默认为local，返回的时间戳单位可通过传入unit参数来改变。 | date："2023-01-02" precision: "day" -> 1672588800date："2023-01-02" precision: "day" zone: "utc" -> 1672617600 |                                                              |
| UnixToDate                | 两个参数：**date：**unix时间戳，单位为秒**format(可选，不传默认精确到秒):** 格式化后的日志格式，由于golang语言限制，**注意模版为：2006-01-02 15:04:05** | string     | 将参数date所表示的unix时间戳转换成日期，再转换成日期字符串   | date:   "1672655888", format: "2006-01-02", -> 2023-01-02date:   "1672655888", format: "2006-01-02  15:04:05", -> 2023-01-02 18:38:08 |                                                              |
| ListJoin                  | 两个参数：**list**：数组元素,sep分隔符           | string     | 将list中的元素通过sep拼接成一个字符串                        | "list": [1.23, 2.45], "sep":  ",", -> "1.23,2.45""list": [1.23, 2.45], "sep":  "-", -> "1.23-2.45" |                                                              |
| RandomUUIDWithReplacer    | 一个参数：**replacer(可选)：**将替换原有UUID自带的分隔符"-"  | string     | 将UUID自带的"-"分隔符用replacer进行替换，返回新连接符的UUID  | replacer："" 8e1d6b823f6f38bb3eb7c7b3556286c3replacer: "#" c804d25d#b6d0#8e31#1e30#e72188745f42 |                                                              |
| MarshalString | 两个参数：<br />1.target：要encoding的对象<br />2. escape_html：是否转移html特殊字符 | string | target可以为任意类型 |  | |

## 注册自定义函数 <font color="red"><sup>new</sup></font>
在arishem初始化时，通过三种option注册三种类型入参的自定义函数。
```go
func init() {
    arishem.Initialize(
        arishem.DefaultConfiguration(),
        arishem.WithFeatureFetcherFactory(func() typedef.FeatureFetcher {
            return new(MyFeatureFetcher)
        }),
        arishem.WithCustomNoParamFuncs(arishem.NoParamFnPair{Name: "NPGreeting", Fn: func(ctx context.Context) (interface{}, error) {
            return "hello arishem", nil
        }}),
        arishem.WithCustomMapParamFuncs(MapParamFnPair{Name: "MPGreeting", Fn: func(ctx context.Context, params map[string]interface{}) (interface{}, error) {
            return "Hello " + tools.ConvToUnifiedStringType(param["name"]), nil
        }}),
        arishem.WithCustomListParamFuncs(ListParamFnPair{Name: "LPGreeting", Fn: func(ctx context.Context, params []interface{}) (interface{}, error) {
            return "Hello " + tools.ConvToUnifiedStringType(param[0]), nil
        }}),
    )
}
```
- 无参数的自定义函数 <font color="red"><sup>new</sup></font>
```go
WithCustomNoParamFuncs(arishem.NoParamFnPair{Name: "NPGreeting", Fn: func(ctx context.Context) (interface{}, error) {
    return "hello arishem", nil
}}),
```
- Map参数类型的自定义函数 <font color="red"><sup>new</sup></font>
```go
WithCustomMapParamFuncs(MapParamFnPair{Name: "MPGreeting", Fn: func(ctx context.Context, params map[string]interface{}) (interface{}, error) {
    return "Hello " + tools.ConvToUnifiedStringType(param["name"]), nil
}}),
```
- List参数类型的自定义函数 <font color="red"><sup>new</sup></font>
```go
WithCustomListParamFuncs(ListParamFnPair{Name: "LPGreeting", Fn: func(ctx context.Context, params []interface{}) (interface{}, error) {
    return "Hello " + tools.ConvToUnifiedStringType(param[0]), nil
}}),
```
