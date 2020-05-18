### 一. 实例     
公共部分
```golang
type animal interface {
	eat()
}

type dog struct {
	name string
}
func (this dog) eat() {
	fmt.Println("dog("+this.name+") eat bone")
}

type cat struct {
	name string
}
func (this *cat) eat() {
	fmt.Println("cat("+this.name+") eat fish")
}

type lion struct {
	name string 
}
func (this *lion) setName(name string) {
	this.name = name
}
func (this lion) eat() {
	fmt.Println("lion("+this.name+") eat rabbit!")
}
```

##### 正确的使用
```golang
func implTest() {
	var p animal 
	p = dog{"tom"}
	fmt.Println("p type:", reflect.TypeOf(p))
	p.eat()
	p = &cat{"jerry"}
	fmt.Println("p type:", reflect.TypeOf(p))
	p.eat()
	l := &lion{}
	l.setName("tomasi")
	fmt.Println("l type:", reflect.TypeOf(l))
	l.eat()
}
```
输出
```
p type: main.dog
dog(tom) eat bone
p type: *main.cat
cat(jerry) eat fish
l type: *main.lion
lion(tomasi) eat rabbit!
```

##### 错误的使用1
```golang
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
##### 错误的使用2
```golang
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

### 二. 概念
value method
```golang
func (this dog) eat() {
	fmt.Println("dog("+this.name+") eat bone")
}
```
pointer method
```golang
func (this *cat) eat() {
	fmt.Println("cat("+this.name+") eat fish")
}
```
value receiver
```golang
dog{"tom"}
```
pointer receiver
```golang
&cat{"jerry"}
```

### 三. 结论
1. interface即可以是value receiver也可以是pointer receiver
2. value method 可以被 pointer和value 对象调用，而pointer method 只能被 pointer 对象调用
```
原因是：pointer method会修改对象的值，而value method不会，所以如果在一个value对象上调用pointer method，编译器会对原来的值做一份拷贝(参考函数传参规范)，并在拷贝后的值上执行函数，那么如果函数有修改原receiver的值，则修改的行为都发生在拷贝的值上，而不会影响原值，这个错误很隐蔽不容易被调试发现，因此go决定放弃这个错误发生的可能性，直接不支持pointer method被value对象调用。
```
