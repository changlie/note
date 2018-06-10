
## 记录一下 golang的反射

```go
type Person struct {
	id int
	Name string
	Age  int
}
func (p Person) GetName() string {
	return p.Name
}

func (p Person) SetName(s string) {
	p.Name = s
}

func (this Person) GetAge() int {
	return this.Age
}

func (p *Person) SetAge(age int) {
	p.Age = age
}

func (this *Person) tostring() {
	fmt.Printf("id:%v, name:%v, age:%v \n", this.id, this.Name, this.Age)
}

func main(){
	obj := Person{0, "mobile", 10086}
	// 类型信息获取
	t := reflect.TypeOf(obj)      //reflect.TypeOf()：要求传入的参数是一个对象的值, 否则为空
	fmt.Println("obj's type is ",t.Name(), "-->", t.Kind())

	// 值操作
	v := reflect.ValueOf(&obj).Elem()    //reflect.ValueOf()：要求传入的参数是一个对象的引用, 否则会报错；而且必须调Elem()后才有意义（好操蛋）
	k := v.Type() // 参数值对象获取类型对象

	fmt.Println("\n get field value:-------------------------")
	// 遍历 obj 内的成员变量（成员须是公有的，否则尝试获取值时， 会报错）
	for i := 0; i < v.NumField(); i++ {
		key := k.Field(i)
		val := v.Field(i)
		var value interface{}
		if val.CanInterface() { //CanInterface(): 判断该成员变量是否能被获取值
			value = val.Interface()
		}
		fmt.Println("line",i, "--> ", key.Name, val.Type(), value)
	}

	fmt.Println("\n change field value:-----------------------------")
	//改变成员变量的值
	f := v.FieldByName("Name")
	if f.CanSet() { //判断成员变量能否被直接改变。
		f.Set(reflect.ValueOf("Changelie777"))
	}
	fmt.Println(obj.Name)

	//获取成员方法，并执行。（这个有点绕，共有两种情况）
	fmt.Println("\n get method way1:----------------------------------")
	//1. 当成员方法 与对象 建立关系时（这种情况，方法改变不了成员变量的值）， 通过对象的指针就可以拿到
	v0 := reflect.ValueOf(&obj).Elem()
	k0 := v0.Type()
	for i := 0; i < v0.NumMethod(); i++ {
		key := k0.Method(i)
		val := v0.Method(i)
		fmt.Println(key.Name, val.Type(), val.Interface())
	}
	//结果：有SetName(), GetName() 两个 方法

	fmt.Println("\n get method way2:----------------------------------")
	//2. 当成员方法 与对象的指针 建立关系时（这种情况，方法可以改变成员变量的值）， 就要通过对象的指针 的指针(×双重指针) 才可以拿到
	temp := &obj
	v1 := reflect.ValueOf(&temp).Elem()
	k1 := v1.Type()
	for i := 0; i < v1.NumMethod(); i++ {
		key := k1.Method(i)
		val := v1.Method(i)
		fmt.Println(key.Name, val.Type(), val.Interface())
	}
	// 结果：只有SetAge()方法

	fmt.Println("\n call Method:-----------------------")
	//根据方法名调用方法
	m := v1.MethodByName("SetAge")
	fmt.Println("method: SetAge,", m)
	m.Call([]reflect.Value{reflect.ValueOf(111)})

	//打印，obj的最终值
	fmt.Println("\n final:----------------------------")
	obj.tostring()
}

```
