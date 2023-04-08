## Map Summary
+ hashtable是一种
+ map由多个无序的key-value对组成的集合
+ 所有的key都是不同的，通过给定的key在常数时间内检索、更新、删除对应的value
+ 一个map是一个哈希表的引用
+ map的所有的key和所有的value都是一样的类型
+ key可以通过 == 进行比较
## 创建Map
+使用make初始化
```go
var mapdemo = make(map[string]int)
```
+ 直接声明
```go
var mapdemo = map[string]int {
    "str1" : 11,
    "str2" : 12
}
```
等价于
```go
var mapdemo = make(map[string]int)
mapdemo["str1"] = 11
mapdemo["str2"] = 12
```
+ 创建一个空的map

```go
var mapdemo = map[string]int {}
```

## 访问map中的元素
+ 直接使用key访问

```go
var mapdemo = map[string]int {
    "str1" : 11,
    "str2" : 13,
}

fmt.Println(mapdemo["str1"])
a := mapdemo["str1"] + 1;
b := 
```

+ 遍历map中的元素

```go
mapdemo := make(map[string]int, 10)
for i := 0 ; i < 10 ; i++ {
    key := "str" + i
    value := i * 10 
    mapdemo[key] = value
}
for key, value := range 
```
>虽然map是动态增加的，也可以初始化容量
>当key-value对的个数超过map的容量时，map的大小会自动加1

```go
var map1 = make(map[string]int, 10)
```



## Map的空值问题
见MapDemo