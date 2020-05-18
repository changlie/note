interface 实例是不允许修改其成员变量的, 所以方法的receiver是pointer receiver时, 说明该struct没有实现相应的interface

详情见下方例子

公共部分
```golang
type animal interface {
	eat()
}
```

### 正确的实现
```golang
type dog struct {
	name string
}
func (this dog) eat() {
	fmt.Println("dog("+this.name+") eat bone")
}
func implTest() {
	var p animal 
	p = dog{"tom"}
	p.eat()
}
```
输出
```
dog(tom) eat bone
```

### 错误的实现
```golang
type cat struct {
	name string
}
func (this *cat) eat() {
	fmt.Println("cat("+this.name+") eat fish")
}
func implTest() {
	var p animal 
	p = cat{"jerry"}
	p.eat()
}
```
输出
```
# command-line-arguments
src\demo\demo.go:75:4: cannot use cat literal (type cat) as type animal in assignment:
        cat does not implement animal (eat method has pointer receiver)
```
