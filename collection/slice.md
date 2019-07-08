# 数组和slice


## 数组声明
    格式为：
        var identifier [len]type     var arr2 [5]int   a := [...]string{"a", "b", "c", "d"}
        var identifier = ([len]type)  var arr1 = new([5]int)
        两者区别：
            arr1 的类型是 *[5]int，而 arr2的类型是 [5]int。
    
## 切片声明
   格式为：
       声明切片的格式是： var identifier []type（不需要说明长度）。一个切片在未初始化之前默认为 nil，长度为 0。
       切片的初始化格式是：var slice1 []type = arr1[start:end]
       
       一个由数字 1、2、3 组成的切片可以这么生成：s := [3]int{1,2,3}[:](注: 应先用s := [3]int{1, 2, 3}生成数组,
        再使用s[:]转成切片) 甚至更简单的 s := []int{1,2,3}。