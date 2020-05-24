### 字符相关


[strings官方文档](https://golang.org/pkg/strings)

<details>
  <summary> strings Index </summary>

```golang
func Compare(a, b string) int
func Contains(s, substr string) bool
func ContainsAny(s, chars string) bool
func ContainsRune(s string, r rune) bool
func Count(s, substr string) int
func EqualFold(s, t string) bool
func Fields(s string) []string
func FieldsFunc(s string, f func(rune) bool) []string
func HasPrefix(s, prefix string) bool
func HasSuffix(s, suffix string) bool
func Index(s, substr string) int
func IndexAny(s, chars string) int
func IndexByte(s string, c byte) int
func IndexFunc(s string, f func(rune) bool) int
func IndexRune(s string, r rune) int
func Join(elems []string, sep string) string
func LastIndex(s, substr string) int
func LastIndexAny(s, chars string) int
func LastIndexByte(s string, c byte) int
func LastIndexFunc(s string, f func(rune) bool) int
func Map(mapping func(rune) rune, s string) string
func Repeat(s string, count int) string
func Replace(s, old, new string, n int) string
func ReplaceAll(s, old, new string) string
func Split(s, sep string) []string
func SplitAfter(s, sep string) []string
func SplitAfterN(s, sep string, n int) []string
func SplitN(s, sep string, n int) []string
func Title(s string) string
func ToLower(s string) string
func ToLowerSpecial(c unicode.SpecialCase, s string) string
func ToTitle(s string) string
func ToTitleSpecial(c unicode.SpecialCase, s string) string
func ToUpper(s string) string
func ToUpperSpecial(c unicode.SpecialCase, s string) string
func ToValidUTF8(s, replacement string) string
func Trim(s string, cutset string) string
func TrimFunc(s string, f func(rune) bool) string
func TrimLeft(s string, cutset string) string
func TrimLeftFunc(s string, f func(rune) bool) string
func TrimPrefix(s, prefix string) string
func TrimRight(s string, cutset string) string
func TrimRightFunc(s string, f func(rune) bool) string
func TrimSpace(s string) string
func TrimSuffix(s, suffix string) string
type Builder
    func (b *Builder) Cap() int
    func (b *Builder) Grow(n int)
    func (b *Builder) Len() int
    func (b *Builder) Reset()
    func (b *Builder) String() string
    func (b *Builder) Write(p []byte) (int, error)
    func (b *Builder) WriteByte(c byte) error
    func (b *Builder) WriteRune(r rune) (int, error)
    func (b *Builder) WriteString(s string) (int, error)
type Reader
    func NewReader(s string) *Reader
    func (r *Reader) Len() int
    func (r *Reader) Read(b []byte) (n int, err error)
    func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)
    func (r *Reader) ReadByte() (byte, error)
    func (r *Reader) ReadRune() (ch rune, size int, err error)
    func (r *Reader) Reset(s string)
    func (r *Reader) Seek(offset int64, whence int) (int64, error)
    func (r *Reader) Size() int64
    func (r *Reader) UnreadByte() error
    func (r *Reader) UnreadRune() error
    func (r *Reader) WriteTo(w io.Writer) (n int64, err error)
type Replacer
    func NewReplacer(oldnew ...string) *Replacer
    func (r *Replacer) Replace(s string) string
    func (r *Replacer) WriteString(w io.Writer, s string) (n int, err error)
```

</details>

[strconv官方文档](https://golang.org/pkg/strconv/)

<details>
  <summary> strconv Index </summary>

```golang
func AppendBool(dst []byte, b bool) []byte
func AppendFloat(dst []byte, f float64, fmt byte, prec, bitSize int) []byte
func AppendInt(dst []byte, i int64, base int) []byte
func AppendQuote(dst []byte, s string) []byte
func AppendQuoteRune(dst []byte, r rune) []byte
func AppendQuoteRuneToASCII(dst []byte, r rune) []byte
func AppendQuoteRuneToGraphic(dst []byte, r rune) []byte
func AppendQuoteToASCII(dst []byte, s string) []byte
func AppendQuoteToGraphic(dst []byte, s string) []byte
func AppendUint(dst []byte, i uint64, base int) []byte
func Atoi(s string) (int, error)
func CanBackquote(s string) bool
func FormatBool(b bool) string
func FormatFloat(f float64, fmt byte, prec, bitSize int) string
func FormatInt(i int64, base int) string
func FormatUint(i uint64, base int) string
func IsGraphic(r rune) bool
func IsPrint(r rune) bool
func Itoa(i int) string
func ParseBool(str string) (bool, error)
func ParseFloat(s string, bitSize int) (float64, error)
func ParseInt(s string, base int, bitSize int) (i int64, err error)
func ParseUint(s string, base int, bitSize int) (uint64, error)
func Quote(s string) string
func QuoteRune(r rune) string
func QuoteRuneToASCII(r rune) string
func QuoteRuneToGraphic(r rune) string
func QuoteToASCII(s string) string
func QuoteToGraphic(s string) string
func Unquote(s string) (string, error)
func UnquoteChar(s string, quote byte) (value rune, multibyte bool, tail string, err error)
type NumError
    func (e *NumError) Error() string
    func (e *NumError) Unwrap() error
```
</details>


