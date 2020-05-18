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

type 
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
}
```
输出
```
p type: main.dog
dog(tom) eat bone
p type: *main.cat
cat(jerry) eat fish
```

##### 错误的使用
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
