> 参考: [官方文档](https://golang.org/pkg/builtin/)


<details>
  <summary> Index </summary>

```golang
func append(slice []Type, elems ...Type) []Type
func cap(v Type) int
func close(c chan<- Type)
func complex(r, i FloatType) ComplexType
func copy(dst, src []Type) int
func delete(m map[Type]Type1, key Type)
func imag(c ComplexType) FloatType
func len(v Type) int
func make(t Type, size ...IntegerType) Type
func new(Type) *Type
func panic(v interface{})
func print(args ...Type)
func println(args ...Type)
func real(c ComplexType) FloatType
func recover() interface{}
type ComplexType
type FloatType
type IntegerType
type Type
type Type1
type bool
type byte
type complex128
type complex64
type error
type float32
type float64
type int
type int16
type int32
type int64
type int8
type rune
type string
type uint
type uint16
type uint32
type uint64
type uint8
type uintptr
```

</details>


### type error
The error built-in interface type is the conventional interface for representing an error condition, with the nil value representing no error.
```golang
type error interface {
    Error() string
}
```

### type rune
rune is an alias for int32 and is equivalent to int32 in all ways. It is used, by convention, to distinguish character values from integer values.
```golang
type rune = int32
```

### nil
nil is a predeclared identifier representing the zero value for a pointer, channel, func, interface, map, or slice type.
```golang
var nil Type // Type must be a pointer, channel, func, interface, map, or slice type
```
