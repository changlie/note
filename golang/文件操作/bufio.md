```golang
type Reader
    func NewReader(rd io.Reader) *Reader
    func NewReaderSize(rd io.Reader, size int) *Reader
    func (b *Reader) Buffered() int
    func (b *Reader) Discard(n int) (discarded int, err error)
    func (b *Reader) Peek(n int) ([]byte, error)
    func (b *Reader) Read(p []byte) (n int, err error)
    func (b *Reader) ReadByte() (byte, error)
    func (b *Reader) ReadBytes(delim byte) ([]byte, error)
    func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error) // 按行读取文件
    func (b *Reader) ReadRune() (r rune, size int, err error)  // 按字符读取
    func (b *Reader) ReadSlice(delim byte) (line []byte, err error)
    func (b *Reader) ReadString(delim byte) (string, error)
    func (b *Reader) Reset(r io.Reader)
    func (b *Reader) Size() int
    func (b *Reader) UnreadByte() error
    func (b *Reader) UnreadRune() error
    func (b *Reader) WriteTo(w io.Writer) (n int64, err error)
type Scanner
    func NewScanner(r io.Reader) *Scanner
    func (s *Scanner) Buffer(buf []byte, max int)
    func (s *Scanner) Bytes() []byte
    func (s *Scanner) Err() error
    func (s *Scanner) Scan() bool
    func (s *Scanner) Split(split SplitFunc)
    func (s *Scanner) Text() string
type SplitFunc
type Writer
    func NewWriter(w io.Writer) *Writer
    func NewWriterSize(w io.Writer, size int) *Writer
    func (b *Writer) Available() int
    func (b *Writer) Buffered() int
    func (b *Writer) Flush() error  // Flush writes any buffered data to the underlying io.Writer.
    func (b *Writer) ReadFrom(r io.Reader) (n int64, err error)
    func (b *Writer) Reset(w io.Writer)
    func (b *Writer) Size() int
    func (b *Writer) Write(p []byte) (nn int, err error)
    func (b *Writer) WriteByte(c byte) error
    func (b *Writer) WriteRune(r rune) (size int, err error)  // 写出一个字符
    func (b *Writer) WriteString(s string) (int, error)  // 写出一个字符串
```

<details>
	<summary> ReadLine方法的按行读取 </summary>

```golang
input := "Now is the winter of our discontent,\nMade glorious summer by this sun of York.\ngood day"
br := bufio.NewReader(strings.NewReader(input))
for {
	a, _, c := br.ReadLine()
	if c == io.EOF {
			break
	}
	fmt.Println(string(a))
}
```

</details>


> Scanner的  `Split`,`Scan`,`Text` 这三个方法是结合使用的。     
默认的Split函数是`ScanLines`
```golang
// NewScanner returns a new Scanner to read from r.
// The split function defaults to ScanLines.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{
		r:            r,
		split:        ScanLines,
		maxTokenSize: MaxScanTokenSize,
	}
}
```
### 常见Scanner实例
<details>
  <summary> 统计word </summary>

```golang
input := "Spicy jalapeno pastrami ut ham turducken.\n Lorem sed ullamco, leberkas sint short loin strip steak ut shoulder shankle porchetta venison prosciutto turducken swine.\n Deserunt kevin frankfurter tongue aliqua incididunt tri-tip shank nostrud.\n"
scanner := bufio.NewScanner(strings.NewReader(input))
// Set the split function for the scanning operation.
scanner.Split(bufio.ScanWords)
// Count the words.
count := 0
for scanner.Scan() {
    count++
}
if err := scanner.Err(); err != nil {
    fmt.Fprintln(os.Stderr, "reading input:", err)
}
fmt.Printf("%d\n", count)
```

</details>

<details>
  <summary> 统计line </summary>

```golang
input := "Spicy jalapeno pastrami ut ham turducken.\n Lorem sed ullamco, leberkas sint short loin strip steak ut shoulder shankle porchetta venison prosciutto turducken swine.\n Deserunt kevin frankfurter tongue aliqua incididunt tri-tip shank nostrud.\n"

scanner := bufio.NewScanner(strings.NewReader(input))
// Set the split function for the scanning operation.
scanner.Split(bufio.ScanLines)
// Count the words.
count := 0
for scanner.Scan() {
    count++
}
if err := scanner.Err(); err != nil {
    fmt.Fprintln(os.Stderr, "reading input:", err)
}
fmt.Printf("%d\n", count)
```

</details>

<details>
  <summary> 按word读取 </summary>

```golang
	// An artificial input source.
	const input = "Now is the winter of our discontent,\nMade glorious summer by this sun of York.\n"
	scanner := bufio.NewScanner(strings.NewReader(input))
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	// Count the words.
	count := 0
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Printf("%d\n", count)
```

</details>

<details>
  <summary> 按Line读取 </summary>

```golang
	s := "hey\nhi,tom\nwhat's up?\nI'm fine!thanks"
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
```

</details>

<details>
  <summary> 自定义split 1 </summary>

```golang
  // An artificial input source.
  const input = "1234 5678 1234567901234567890"
  scanner := bufio.NewScanner(strings.NewReader(input))
  // Create a custom split function by wrapping the existing ScanWords function.
  split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
    advance, token, err = bufio.ScanWords(data, atEOF)
    if err == nil && token != nil {
      _, err = strconv.ParseInt(string(token), 10, 32)
    }
    return
  }
  // Set the split function for the scanning operation.
  scanner.Split(split)
  // Validate the input
  for scanner.Scan() {
    fmt.Printf("%s\n", scanner.Text())
  }

  if err := scanner.Err(); err != nil {
    fmt.Printf("Invalid input: %s", err)
  }
```

</details>

<details>
  <summary> 自定义split 2 </summary>

```golang
	// Comma-separated list; last entry is empty.
	const input = "1,2,3,4,"
	scanner := bufio.NewScanner(strings.NewReader(input))
	// Define a split function that separates on commas.
	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == ',' {
				return i + 1, data[:i], nil
			}
		}
		if !atEOF {
			return 0, nil, nil
		}
		// There is one final token to be delivered, which may be the empty string.
		// Returning bufio.ErrFinalToken here tells Scan there are no more tokens after this
		// but does not trigger an error to be returned from Scan itself.
		return 0, data, bufio.ErrFinalToken
	}
	scanner.Split(onComma)
	// Scan.
	for scanner.Scan() {
		fmt.Printf("%q ", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
```

</details>



