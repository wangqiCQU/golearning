# map
## map声明
    var map1 map[keytype]valuetype
    var map1 map[string]int
    
        map 传递给函数的代价很小：在 32 位机器上占 4 个字节，64 位机器上占 8 个字节，
    无论实际上存储了多少数据。通过 key 在 map 中寻找值是很快的，比线性查找快得多，
    但是仍然比从数组和切片的索引中直接读取要慢 100 倍；所以如果你很在乎性能的话
    还是建议用切片来解决问题。
    
    map 的初始化：var map1 = make(map[keytype]valuetype)。
    
    或者简写为：map1 := make(map[keytype]valuetype)。
    
    不要使用 new，永远用 make 来构造 map
    
    注意 如果你错误的使用 new() 分配了一个引用对象，你会获得一个空引用的指针，相当于声明了一个未初始化的变量并且取了它的地址：
    
    mapCreated := new(map[string]float32)
    
    
## 测试键值对是否存在
    , isPresent = map1[key1]   isPresent 返回一个 bool 值：如果 key1 存在于 map1，val1 就是 key1 对应的 value 值，
    并且 isPresent为true；如果 key1 不存在，val1 就是一个空值，并且 isPresent 会返回 false。
    
## 删除键
    delete(map key)
    
## map类型切片

  假设我们想获取一个 map 类型的切片，我们必须使用两次 make() 函数，第一次分配切片，
第二次分配 切片中每个 map 元素


##new() 和 make() 的区别

看起来二者没有什么区别，都在堆上分配内存，但是它们的行为不同，适用于不同的类型。

new(T) 为每个新的类型T分配一片内存，初始化为 0 并且返回类型为*T的内存地址：这种方法 返回一个指向类型为 T，值为 0 的地址的指针，它适用于值类型如数组和结构体（参见第 10 章）；它相当于 &T{}。
make(T) 返回一个类型为 T 的初始值，它只适用于3种内建的引用类型：切片、map 和 channel（参见第 8 章，第 13 章）。