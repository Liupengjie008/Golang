Golang Map：引用类型，哈希表。键必须是支持相等运算符 ("=="、"!=") 类型， 如 number、string、 pointer、array、struct，以及对应的 interface。值可以是任意类型，没有限制。

map常见操作:
    // 创建map
    var m map[string]string = map[string]string{“hello”: “world”}
    // map初始化
    m = make(map[string]string, 10)
    


    m := map[int]struct {
        name string
        age int 
    }{
        1:  {"user1", 10},  // 可省略元素类型。
        2:  {"user2", 20},
    }
    
    println(m[1].name)
            
    预先给 make 函数一个合理元素数量参数，有助于提升性能。因为事先申请一大块内存，可避免后续操作时频繁扩张。
    
    m := make(map[string]int, 1000)    

    m := map[string]int{
        "a": 1,
    }

    // 插入和更新：
    m["hello"] = 2
    
    // 查找：
    val, ok := m["hello"]

    // 判断 key 是否存在。
    if v, ok := m["a"]; ok {    
        println(v)
    }
    
    println(m["c"])     // 对于不存在的 key，直接返回 \0，不会出错。
    
    m["a"] = 2       // 新增或修改。
    
    // 删除：
    delete(m, "c")      // 删除。如果 key 不存在，不会出错。
    
    // 长度：
    len(m)
    println(len(m))    // 获取键值对数量。cap 无效。
    
    // 遍历：
    for k, v := range m {       // 迭代，可仅返回 key。随机顺序返回，每次都不相同。
        println(k, v)
    }
    
    不能保证迭代返回次序，通常是随机结果，具体和版本实现有关。

    

slice与map操作（slice of map）
    items := make([]map[int][int], 5)
    for i := 0; i < 5; i++ {
        items[i] = make(map[int][int])
    }

map排序：先获取所有key，把key进行排序，再按照排序好的key，进行遍历。
map反转：初始化另外一个map，把key、value互换即可.






从 map 中取回的是一个 value 临时复制品，对其成员的修改是没有任何意义的。
  
type user struct{ name string }

m := map[int]user{  // 当 map 因扩张而重新哈希时，各键值项存储位置都会发生改变。 因此，map 被设计成 not addressable。 类似 m[1].name 这种期望透过原 value 指针修改成员的行为自然会被禁 。

    1: {"user1"},
}

m[1].name = "Tom"   // Error: cannot assign to m[1].name


正确做法是完整替换 value 或使用指针。

u := m[1]
u.name = "Tom"
m[1] = u    // 替换 value。

m2 := map[int]*user{
    1: &user{"user1"},
}

m2[1].name = "Jack"     // 返回的是指针复制品。透过指针修改原对象是允许的。


可以在迭代时安全删除键值。但如果期间有新增操作，那么就不知道会有什么意外了。

for i := 0; i < 5; i++ {
    m := map[int]string{
        0:  "a", 1:  "a", 2:  "a", 3:  "a", 4:  "a",
        5:  "a", 6:  "a", 7:  "a", 8:  "a", 9:  "a",
    }

    for k := range m {
        m[k+k] = "x"
        delete(m, k)
    }

    fmt.Println(m)
}

输出:     
//每次输出都会变化
map[36:x 28:x 32:x 2:x 8:x 10:x 12:x]
map[12:x 6:x 16:x 28:x 4:x 10:x 72:x]
map[12:x 14:x 16:x 18:x 20:x]
map[18:x 10:x 14:x 4:x 6:x 16:x 24:x]
map[12:x 16:x 4:x 40:x 14:x 18:x]


