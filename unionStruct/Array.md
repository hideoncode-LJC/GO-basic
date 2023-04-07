# Array summary
+ 在GO中，数组是一种类型，即[50]int 与 [10]int 就像是 float 和 int 的区别
+ 数组声明类型之后可以不进行赋值
	+ int 被赋值为0
	+ float
	+ string被赋值为nil
## 定长数组声明
+ 普通声明
```go
var arrayByNormal = [5]int{1, 2, 3, 4, 5}
arrayByNormalOne := [5]int{1, 2, 3, 4, 5}
var arrayNot [20]int
```
+ 索引声明
```go
//注意不要越界，没被索引的位置默认为0 int类型为0
var arrayByIndex = [5]int{1 : 2, 3 : 4}
```
+ 混合声明
```go
//注意不要越界
var arrayByMix = [5]int{1 : 2, 3, 4}
```
## 变长数组
+ 变长数组不代表数据的长度是变长的
+ 数据类型一经定义，数据类型的长度就固定下来的了
+ 变长数组是根据数组的定理来确定数组的长度
+ 变长数组声明同样有普通声明、索引声明、混合声明三种方式
```go
var arrayLazy = [...]int{5, 6, 7, 8, 9}
var arrayPack = [...]int{6 : 10}
var arrayMix = [...]{10 : 11, 12, 13, 14}
```
## new 声明数组
+ new声明数组返回的是地址
```go
var arrayByNew = new([5]int)
```
## 多维数组的声明
+ 多维数组的声明只有第一维能使用省略号
```go
var multiArray = [...][5]int{{1, 2 ,3 ,4 ,5},{2, 2, 2, 2, 2}}
```
## 获取数组的长度
+ len函数
+ cap函数
>对于多维数组来说，len和cap函数只返回第一维度的个数
## 数组的遍历
```go
//一维
for i, v := range arr {
	fmt.Prinf("第%d位置的值为%d\n", i, v)
}
//多维
var multiArray = [...][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}, {13, 14, 15}}  
for i := range multiArray {  
   for j, v := range multiArray[i] {  
      fmt.Printf("第%d行%d列位置的元素的值为%d\n", i, j ,v)  
   }  
}
```
## 数组的应用
+ function传递参数时一般不用数组类型，使用new创造的数组类型，提高效率
## new声明数组和普通的数组类型的比较 
见Array
```go
//普通类型的声明
var a = [5]int{1, 2, 3, 4, 5}  
arr := a  
arr[1] = 1000  
fmt.Printf("a[1]的值 = %d\n", a[1])  
fmt.Printf("arr[1]的值 = %d\n", arr[1])  
//new类型的声明
var a1 = new([5]int)  
arr1 := a1  
arr1[1] = 100  
fmt.Printf("a1[1]的值 = %d\n", a1[1])  
fmt.Printf("arr1[1]的值 = %d\n", arr1[1])
```
+ 输出结果
![[Pasted image 20230102214823.png]]
## 数组类型之间的比较
+ 数组比较必须是同等类型的元素之间的比较
+ 不同长度的数组类型进行比较无法通过编译
+ 同等长度的数组类型比较必须每个位置的元素都相等
+ 数据类型和指针类型也不能进行比较
```go
var a = [5]int{1, 2, 3, 4, 5}  
var b = [5]int{1, 3, 3, 4, 5}  
//var c = [10]int{9 : 10}  
var d = [5]int{1, 2, 3, 4, 5}  
  
if a == b {  
   fmt.Printf("a == b\n")  
}  
  
if a == d {  
   fmt.Printf("a == d\n")  
}
```