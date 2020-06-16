`encoding/json`主要功能api
```golang
func Marshal(v interface{}) ([]byte, error)   // obj to json
func Unmarshal(data []byte, v interface{}) error // json to obj
func Valid(data []byte) bool // is json ?
type Decoder 
    func NewDecoder(r io.Reader) *Decoder
    func (dec *Decoder) Decode(v interface{}) error //  file > json to obj
type Encoder
    func NewEncoder(w io.Writer) *Encoder
    func (enc *Encoder) Encode(v interface{}) error // obj to json > file
```

<details>
  <summary> 把一个对象转成json字符串 </summary>

```golang
func main() {
	type Message struct {
		Name string
		Body string
		Time int64
	}

	m := Message{"Alice", "Hello", 1294706395881547000}
	bs, err := json.Marshal(m)
	fmt.Printf("%s, err: %v\n", bs, err)

	m1 := map[string]interface{}{
		"id": 10010,
		"name": "IT DEP",
		"duty": "coding",
		"is old": false,
	}
	bs, err = json.Marshal(m1)
	fmt.Printf("%s, err: %v\n", bs, err)
}
```

</details>


<details>
  <summary> 把一个json字符串转成对象 </summary>

```golang
type Message struct {
	Name string
	Body string
	Time int64
}

func (this Message) String() string {
	return fmt.Sprintf(`{name:%v, body:%v, time:%v}`, this.Name, this.Body, this.Time)
}

func main() {
	s1 := `{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`
	s2 := `{"duty":"coding","id":10010,"is old":false,"name":"IT DEP","body":"it's html","time":10086}`
	var m1, m2 Message
	err1 := json.Unmarshal([]byte(s1), &m1)
	err2 := json.Unmarshal([]byte(s2), &m2)
	fmt.Printf("%v, err: %v \n", m1, err1)
	fmt.Printf("%v, err: %v \n", m2, err2)

	fmt.Println("seperator line--------------")

	var f interface{}
	var m3 map[string]interface{}
	m4 := make(map[string]interface{})
	err1 = json.Unmarshal([]byte(s2), &f)
	err2 = json.Unmarshal([]byte(s2), &m3)
	err3 := json.Unmarshal([]byte(s2), &m4)
	fmt.Printf("%v, type: %T, err: %v \n", f, f, err1)
	fmt.Printf("%v, type: %T, err: %v \n", m3, m3, err2)
	fmt.Printf("%v, type: %T, err: %v \n", m4, m4, err3)
}
```

</details>

<details>
  <summary> 从文件读取一个json字符串，并转成obj </summary>

demo1
```golang
func demo1() {
const jsonStream = `
	[
		{"Name": "Ed", "Text": "Knock knock."},
		{"Name": "Sam", "Text": "Who's there?"},
		{"Name": "Ed", "Text": "Go fmt."},
		{"Name": "Sam", "Text": "Go fmt who?"},
		{"Name": "Ed", "Text": "Go fmt yourself!"}
	]
`
	type Message struct {
		Name, Text string
	}
	dec := json.NewDecoder(strings.NewReader(jsonStream))

	// read open bracket
	t, err := dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)

	// while the array contains values
	for dec.More() {
		var m Message
		// decode an array value (Message)
		err := dec.Decode(&m)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v: %v\n", m.Name, m.Text)
	}

	// read closing bracket
	t, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)
}
```
demo2
```golang
func jsonTest1() {
	const jsonStream = `
	{
	"Name": "changlie",
	"Age": 22,
	"Addr": "gd sz"
	}
`
	dec := json.NewDecoder(strings.NewReader(jsonStream))

	var u User
	err := dec.Decode(&u)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(u)
	fmt.Println("-------------")
	var res interface{}
	err = dec.Decode(&res)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(res)
}

type User struct {
	Addr string
	Age int
	Name string
}

func (this User) String() string {
	return fmt.Sprintf(`{"name":"%v", "age":%v, "addr":"%v"}`, this.Name, this.Age, this.Addr)
}
```

</details>


<details>
  <summary> 把一个对象转成json字符串, 并写出到文件 </summary>

```golang
func main() {
	type User1 struct {
		Addr string
		Age int
		Name string
	}

	u := User1{"China SiChuan", 18, "Penda AC"}

	// 结果保存至文件
	f, _ := os.OpenFile("d:/json1.txt", os.O_RDWR | os.O_CREATE, 0666)
	err := json.NewEncoder(f).Encode(u)
	if err != nil {
		fmt.Println(err.Error())
	}

	 // 编码结果暂存到 buffer
	 bytes3 := new(bytes.Buffer)
	 err = json.NewEncoder(bytes3).Encode(u)
	 if err == nil {
			 fmt.Print("json.NewEncoder 编码结果: ", string(bytes3.Bytes()))
	 }
}
```

</details>

<details>
  <summary> func (*Decoder) Token </summary>

```golang
func main() {
	const jsonStream = `
	{"Message": "Hello", "Array": [1, 2, 3], "Null": null, "Number": 1.234}
`
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		t, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%T: %v", t, t)
		if dec.More() {
			fmt.Printf(" (more)")
		}
		fmt.Printf("\n")
	}
}
```

</details>

<details>
  <summary> json标签 </summary>

```golang
func main() {
	type Message struct {
		Name string `json:"name"`
		Body string `json:"content"`
		Time int64 `json:"createtime"`
	}

	m := Message{"Alice", "Hello", 1294706395881547000}
	bs, err := json.Marshal(m)
	fmt.Printf("%T, %s, err: %v\n", m, bs, err)

	bs1 := []byte(`{"name":"tome","body":"js","time":9527, "content":"broadCast", "createtime":123456798}`)
	var msg Message
	err = json.Unmarshal(bs1, &msg)
	fmt.Printf("%v, err: %v \n", msg, err)
}
```

</details>
