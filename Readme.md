---
module: github.com/tinybear1976/localsystem
function: 简单封装本地文件系统操作、日志操作及其他扩展函数
version: 0.12.0
lastdatetime: 2022-02-04
---

目录

[TOC]

# 引用

## go.mod

```go
go 1.15
```

# 日志

## InitLogger

初始化日志文件。需引用包路径  "hhyt/localsystem/logger"

```go
func NewLogger(tag string, logFilenameWithPath string, loglevel string) 
```

入口参数：

| 参数名              | 类型   | 描述                                                         |
| ------------------- | ------ | ------------------------------------------------------------ |
| tag                 | string | 日志文件的tag，可以指定多个标记以致可以同时存在多个日志logger    |
| logFilenameWithPath | string | 日志文件的完整文件名（带路径）                               |
| loglevel            | string | 日志等级，如果传入空字符串，则默认debug：<br>debug<br>info<br>error<br>warn |

返回值：

无（通过  logger.LogContainer["标签名"] 访问）

示例：

```go
package main

import (
	"hhyt/localsystem"
	"hhyt/localsystem/logger"
	"path"
)

func main() {
	println(localsystem.CurrentDirectory())
	testlog()
}

func testlog() {
	p, _ := localsystem.CurrentDirectory()
    //返回的日志可操作指针可以保存到临时变量，也可以不保存，
	logger.NewLogger("log1",path.Join(p, "test.log"), "")
    //因为包内的Log已经公开，初始化后，可以在需要操作日志的模块中，引用该包，直接调用Log即可
    logger.LogContainer["log1"].Info("test infomation")
}
```



# 文件目录类

## CurrentDirectory

获取当前系统路径，字符串末尾不带 /  。

```go
func CurrentDirectory() (currentpath string, err error)
```

入口参数：

无

返回值：

| 返回变量    | 类型   | 描述                                               |
| ----------- | ------ | -------------------------------------------------- |
| currentpath | string | 返回当前路径                                       |
| err         | error  | 如果操作失败，返回异常，一般不会产生错误，可以忽略 |



## IsDir

测试路径是否为目录。

```go
func IsDir(path string) bool
```

入口参数：

| 参数名 | 类型   | 描述               |
| ------ | ------ | ------------------ |
| path   | string | 需要进行测试的路径 |

返回值：

| 返回变量 | 类型 | 描述                                      |
| -------- | ---- | ----------------------------------------- |
|          | bool | 返回值，true表示为目录，false表示不是目录 |



## IsFile

测试路径是否为文件。

```go
func IsFile(path string) bool
```

入口参数：

| 参数名 | 类型   | 描述               |
| ------ | ------ | ------------------ |
| path   | string | 需要进行测试的路径 |

返回值：

| 返回变量 | 类型 | 描述                                      |
| -------- | ---- | ----------------------------------------- |
|          | bool | 返回值，true表示为文件，false表示不是文件 |



## Exists

测试路径是否存在。

```go
func Exists(path string) bool
```

入口参数：

| 参数名 | 类型   | 描述               |
| ------ | ------ | ------------------ |
| path   | string | 需要进行测试的路径 |

返回值：

| 返回变量 | 类型 | 描述                                  |
| -------- | ---- | ------------------------------------- |
|          | bool | 返回值，true表示存在，false表示不存在 |



## FileNameOnly

测试路径是否存在。

```go
func FileNameOnly(fullFilename string) string
```

入口参数：

| 参数名       | 类型   | 描述                     |
| ------------ | ------ | ------------------------ |
| fullFilename | string | 原始文件名，可以携带路径 |

返回值：

| 返回变量 | 类型   | 描述                           |
| -------- | ------ | ------------------------------ |
|          | string | 返回不带路径与扩展名的主文件名 |

示例：

```go
	p, _ := localsystem.CurrentDirectory()
	filename := path.Join(p, "test.log")
	println(filename)
    //返回结果为： test
	println(localsystem.FileNameOnly(filename))
```



## FileNameWithoutPath

返回文件名及扩展名，去除路径。

```go
func FileNameWithoutPath(fullFilename string) string 
```

入口参数：

| 参数名       | 类型   | 描述                     |
| ------------ | ------ | ------------------------ |
| fullFilename | string | 原始文件名，可以携带路径 |

返回值：

| 返回变量 | 类型   | 描述                         |
| -------- | ------ | ---------------------------- |
|          | string | 返回不带路径的 文件名+扩展名 |



## CopyFile

复制磁盘文件。

```go
func CopyFile(src, dst string) (int64, error)
```

入口参数：

| 参数名 | 类型   | 描述                   |
| ------ | ------ | ---------------------- |
| src    | string | 源文件名。可以携带路径 |
| dst    | string | 目标文件。可以携带路径 |

返回值：

| 返回变量 | 类型  | 描述                          |
| -------- | ----- | ----------------------------- |
|          | int64 | 返回复制文件的字节数          |
|          | error | 如果操作没有发生错误，返回nil |



## CreateMutiDir

创建多级目录。该操作直接将希望存在的目录结果以路径字符串方式输入，如果操作成功，多级目录将会自动建立出来。对于输入路径中某些级别的目录是否存在，忽略不计。

```go
func CreateMutiDir(filePath string) error 
```

入口参数：

| 参数名   | 类型   | 描述             |
| -------- | ------ | ---------------- |
| filePath | string | 希望存在的路径。 |

返回值：

| 返回变量 | 类型  | 描述                          |
| -------- | ----- | ----------------------------- |
|          | error | 如果操作没有发生错误，返回nil |



# 日期时间类

## DiffSecByStr

计算时间差（单位秒）。  公式     t2 -  t1  =时间差秒数（理论应该>0）。第一时间为字符串，第二时间为time，t2-t1=相差秒,第一个时间应该小于第二个时间

```go
func DiffSecByStr(firsttime string, endtime time.Time) int
```

入口参数：

| 参数名    | 类型      | 描述                                 |
| --------- | --------- | ------------------------------------ |
| firsttime | string    | 计算的第一个时间(t1)，即较早的时间。 |
| endtime   | time.Time | 计算的第二个时间(t2)，即较早的时间。 |

返回值：

| 返回变量 | 类型 | 描述     |
| -------- | ---- | -------- |
|          | int  | 相差秒数 |



## DiffSecByStrWithFormat

计算时间差（单位秒）。  公式     ts2 -  ts1  =时间差秒数（理论应该>0），用户规定传入的时间格式。ts1与ts2格式应该相同。第一时间为字符串，第二时间为字符串，第三个为日期格式，t2-t1=相差秒,第一个时间应该小于第二个时间，如果出现错误返回-1

```go
func DiffSecByStrWithFormat(ts1, ts2, formattemp string) int
```

入口参数：

| 参数名 | 类型   | 描述                                 |
| ------ | ------ | ------------------------------------ |
| ts1    | string | 计算的第一个时间(t1)，即较早的时间。 |
| ts2    | string | 计算的第二个时间(t2)，即较早的时间。 |

返回值：

| 返回变量 | 类型 | 描述                    |
| -------- | ---- | ----------------------- |
|          | int  | 相差秒数,如果错误返回-1 |



## DiffSecByStrWithFormatErr

与DiffSecByStrWithFormat相似，增加返回错误信息

```go
func DiffSecByStrWithFormatErr(ts1, ts2, formattemp string) (int, error)
```

入口参数：

| 参数名 | 类型   | 描述                                 |
| ------ | ------ | ------------------------------------ |
| ts1    | string | 计算的第一个时间(t1)，即较早的时间。 |
| ts2    | string | 计算的第二个时间(t2)，即较早的时间。 |

返回值：

| 返回变量 | 类型  | 描述     |
| -------- | ----- | -------- |
|          | int   | 相差秒数 |
|          | error | 错误信息 |



## SecondsToString

seconds为需要计算的秒数,如果为负数，函数会自动现将其调整为正数。模板参数"d:hh:mm:ss"为全格式，可以逐级减项，"d:hh:mm" "hh:mm:ss" "hh:mm" "mm:ss"，最少两项。

**功能提示：**该函数只是先将时间计算成 "d:hh:mm:ss"格式，然后按照需要，进行局部格式的裁剪，并且`d`表示为天数，为该整体格式最高单位，所以只能写一个`d`，其余部分最宽两位

```go
func SecondsToString(seconds int, template string) string
```

入口参数：

| 参数名   | 类型   | 描述                   |
| -------- | ------ | ---------------------- |
| seconds  | int    | 需要计算的秒数         |
| template | string | 返回字符串的格式模板。 |

返回值：

| 返回变量 | 类型   | 描述         |
| -------- | ------ | ------------ |
|          | string | 格式化字符串 |



## SecondsToStringTimelong

将秒数转换成字符串格式，**指定返回模板的最高位为完整计算数**，可能会超过2位，比如mm:ss则表示所有的描述最大单位以分钟计数。 seconds为计算秒数,如果为负数，函数会自动现将其调整为正数。

模板参数"d:hh:mm:ss"为全格式，可以逐级减项，"d:hh:mm" "h:mm:ss" "h:mm" "m:ss"，最少两项。

如果模板截断最低位ss，结果将忽略秒数部分。

```go
func SecondsToStringTimelong(seconds int, template string) string
```

入口参数：

| 参数名   | 类型   | 描述                   |
| -------- | ------ | ---------------------- |
| seconds  | int    | 需要计算的秒数         |
| template | string | 返回字符串的格式模板。 |

返回值：

| 返回变量 | 类型   | 描述         |
| -------- | ------ | ------------ |
|          | string | 格式化字符串 |



## String2Timestamp

将字符串，按照指定布局转换为时间。

```go
func String2Timestamp(layout, value string) (time.Time, error)
```

入口参数：

| 参数名 | 类型   | 描述                          |
| ------ | ------ | ----------------------------- |
| layout | string | 对应value的时间格式字符串模板 |
| value  | string | 待转换的时间字符串            |

返回值：

| 返回变量 | 类型      | 描述       |
| -------- | --------- | ---------- |
|          | time.Time | 转换的时间 |
|          | error     | 错误       |

# 字符串类

## ConcatStrings

通过传入字符串数组，拼接类SQL  IN 条件 字符串。例如：["a","b","c"]  =>  " 'a','b','c' "

```go
func ConcatStrings(data []string, separator string, delimiter string) string
```

入口参数：

| 参数名    | 类型     | 描述                       |
| --------- | -------- | -------------------------- |
| data      | []string | 字符串数组                 |
| separator | string   | 指定间隔符，比如 ,         |
| delimiter | string   | 每个字符串的定界符，比如 ' |

返回值：

| 返回变量 | 类型   | 描述           |
| -------- | ------ | -------------- |
|          | string | 拼接结果字符串 |



## InStrings

检索字符串值（target）是否在字符串数组（str_array）中存在中存在

```go
func InStrings(target string, str_array []string) bool
```

入口参数：

| 参数名    | 类型     | 描述                 |
| --------- | -------- | -------------------- |
| target    | string   | 被检索值             |
| str_array | []string | 最大范围的字符串数组 |

返回值：

| 返回变量 | 类型   | 描述           |
| -------- | ------ | -------------- |
|          | string | 拼接结果字符串 |



## GetLastRune

获得指定数量的结尾字符

```go
func GetLastRune(str string, amount int) string 
```

入口参数：

| 参数名 | 类型   | 描述                     |
| ------ | ------ | ------------------------ |
| str    | string | 被检索字符串             |
| amount | int    | 需要从右边返回的字符数量 |

返回值：

| 返回变量 | 类型   | 描述       |
| -------- | ------ | ---------- |
|          | string | 结果字符串 |



## RemoveLastRune

获得指定数量的结尾字符

```go
func RemoveLastRune(str string, amount int) string
```

入口参数：

| 参数名 | 类型   | 描述                     |
| ------ | ------ | ------------------------ |
| str    | string | 原始字符串               |
| amount | int    | 需要从右边删除的字符数量 |

返回值：

| 返回变量 | 类型   | 描述       |
| -------- | ------ | ---------- |
|          | string | 结果字符串 |

## Float64toString

float64转换成字符串

```go
func Float64toString(f64 float64) string
```

# 算法

## 归并排序

### int

```go
// package:  localsystem.algorithm
data := IntSlice_MergeSort([]int{2, 3, 1})
// Sort方法传入true 升序排列，如果传入false降序排列
got := data.Sort(true)
```

### float64

```go
// package:  localsystem.algorithm
data := Float64Slice_MergeSort([]float64{2.2, 30.3, 1})
// Sort方法传入true 升序排列，如果传入false降序排列
got := data.Sort(true)
```

# 数据结构

## Set4String

为字符串设计的集合

包路径：`localsystem/datastructure/set`

### New

根据一个字符串切片创建一个字符串集合

```go
func New(strSlice []string) Set4String 
```

### Intersection

求两（或多个）个字符串集合的交集。字符串集合方法

```go
func (set1 Set4String) Intersection(sets ...Set4String) Set4String 
```

### Compare

比较两个集合是否所有元素一致

```go
func (set1 Set4String) Compare(set2 Set4String) bool
```

### Union

多个集合合并（并集）

```go
func (set1 Set4String) Union(sets ...Set4String) (Set4String, error) 
```

### IsExist

检查1个或多个字符串在集合中是否存。有一个不存在就返回false，全部存在返回true。如果入参没有传入任何内容，返回false。如果指定的集合没有任何内容返回false。

```go
func (set1 Set4String) IsExist(keys ...string) bool 
```

### Categorizing

将给定的字符串按照在集合中出现、未出现进行分检返回。如果被检索集合（set1）不包含任何元素，则传入检索项全部返回non_exist。如果没有传入参数，则返回的 exist, non_exist分别都为空集合

```go
func (set1 Set4String) Categorizing(keys ...string) (exist, non_exist Set4String)
```

### ToStringSlice

将集合转为字符串切片

```go
func (set1 Set4String) ToStringSlice() []string 
```

### Add

增加元素到集合

```go
func (set1 Set4String) Add(keys ...string)
```

### Remove

从集合中移除指定key

```go
func (set1 Set4String) Remove(keys ...string)
```

