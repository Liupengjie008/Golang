表达式
根据调 者不同， 法分为两种表现形式: instance.method(args...) ---> <type>.func(instance, args...)
前者称为 method value，后者 method expression。
两者都可像普通函数那样赋值和传参，区别在于 method value 绑定实例，  method expression 则须显式传参。

 
  type User struct {
    id   int
name string }
func (self *User) Test() {
    fmt.Printf("%p, %v\n", self, self)
}
56

 
func main() {
    u := User{1, "Tom"}
    u.Test()
    mValue := u.Test
    mValue()
    mExpression := (*User).Test
    mExpression(&u)
// 隐式传递 receiver
// 显式传递 receiver
 }
输出:
需要注意，method value 会复制 receiver。
 0x210230000, &{1 Tom}
0x210230000, &{1 Tom}
0x210230000, &{1 Tom}
 type User struct {
    id   int
name string }
func (self User) Test() {
    fmt.Println(self)
}
func main() {
    u := User{1, "Tom"}
    mValue := u.Test
    u.id, u.name = 2, "Jack"
    u.Test()
mValue() }
//  即复制 receiver，因为不是指针类型，不受后续修改影响。
输出:
在汇编层 ，method value 和闭包的实现 式相同，实际返回 FuncVal 类型对象。 FuncVal { method_address, receiver_copy }
可依据 法集转换 method expression，注意 receiver 类型的差异。
 {2 Jack}
{1 Tom}
 57

 输出:
将 法 "还原" 成函数，就容易理解下 的代码了。

type User struct {
    id   int
name string }
func (self *User) TestPointer() {
    fmt.Printf("TestPointer: %p, %v\n", self, self)
}
func (self User) TestValue() {
    fmt.Printf("TestValue: %p, %v\n", &self, self)
}
func main() {
    u := User{1, "Tom"}
    fmt.Printf("User: %p, %v\n", &u, u)
 }
mv := User.TestValue
mv(u)
mp := (*User).TestPointer
mp(&u)
mp2 := (*User).TestValue
mp2(&u)
// *User  法集包含 TestValue。
// 签名变为 func TestValue(self *User)。 // 实际依然是 receiver value copy。
 User       : 0x210231000, {1 Tom}
TestValue  : 0x210231060, {1 Tom}
TestPointer: 0x210231000, &{1 Tom}
TestValue  : 0x2102310c0, {1 Tom}
 type Data struct{}
func (Data) TestValue()    {}
func (*Data) TestPointer() {}
func main() {
    var p *Data = nil
    p.TestPointer()
    (*Data)(nil).TestPointer()  // method value
    (*Data).TestPointer(nil)    // method expression
// p.TestValue()            // invalid memory address or nil pointer dereference
58

     // (Data)(nil).TestValue()  // cannot convert nil to type Data
    // Data.TestValue(nil)      // cannot use nil as type Data in function argument
}