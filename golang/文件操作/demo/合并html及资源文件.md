<details>
  <summary> base64工具 </summary>

```golang
package hash1

import (
  "fmt"
)

var base64 = [65]byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V',
            'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q',
            'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '+', '/', '='}

const FILLER byte = '='

const FILLER_INDEX byte = 64

func getVal(b byte) byte {
  if FILLER == b {
    return 0
  }

  for i, item := range base64 {
    if item == b {
      return byte(i)
    }
  }
  
  var msg string = fmt.Sprintf("not found origin value map: %v \n", string(b))
  panic(msg)
}

func Base64Encode(bytes []byte) string{
  res := Base64Encode2Byte(bytes)
  return string(res[:])
}

func Base64Encode2Byte(bytes []byte) []byte{
  len := len(bytes)
  var newArrLen int 
  
  if len % 3 > 0 {
    newArrLen = (len / 3 + 1) * 4
  } else {
    newArrLen = (len / 3) * 4
  }
  
  res := make([]byte, 0, newArrLen)
  for i := 0; i<len; i += 3 {
    isEnd := i + 3 >= len
    var b1 byte = bytes[i] >> 2 & 0x3f
    var b2 byte
    var b3 byte
    var b4 byte

    if !isEnd || (isEnd && (len-1) % 3 == 2) {
      b2 = (bytes[i] << 4 | (bytes[i + 1] >> 4 & 0x0f)) & 0x3f
      b3 = (bytes[i + 1] << 2 | (bytes[i + 2] >> 6 & 0x03)) & 0x3f
      b4 = bytes[i + 2] & 0x3f
    }else if isEnd && (len-1) % 3 == 1 {
      b2 = (bytes[i] << 4 | (bytes[i + 1] >> 4 & 0x0f)) & 0x3f
      b3 = (bytes[i + 1] << 2) & 0x3c
      b4 = FILLER_INDEX
    }else if isEnd && (len-1) % 3 == 0 {
      b2 = bytes[i] << 4 & 0x30
      b3 = FILLER_INDEX
      b4 = FILLER_INDEX
    }

    res = append(res, base64[b1], base64[b2], base64[b3], base64[b4])
  }
  return res
}

func Base64Decode(bytes []byte) string {
  res := Base64Decode2Byte(bytes)
  return string(res[:])
}

func Base64Decode2Byte(bytes []byte) []byte {
  doubleFiller, singleFiller := checkTailToken(bytes)
  bytes = revertFromBase64(bytes)
  len := len(bytes)
  newArrLen := len / 4 * 3
  res := make([]byte, newArrLen, newArrLen)
  for i := 0; i < len; i += 4 {
    resIndex := i / 4 * 3

    b1 := bytes[i]
    b2 := bytes[i+1]
    b3 := bytes[i+2]
    b4 := bytes[i+3]

    res[resIndex] = (b1 << 2 | ((b2 >> 4) & 0x03)) & 0xff
    res[resIndex+1] = (b2 << 4 | ((b3 >> 2) & 0x0f)) & 0xff
    res[resIndex+2] = (b3 << 6 | b4) & 0xff
  }

  if doubleFiller {
    return res[:newArrLen-2]
  }

  if singleFiller {
    return res[:newArrLen-1]
  }

  return res
}

func checkTailToken(bytes []byte) (doubleFiller, singleFiller bool) {
  len := len(bytes)
  if bytes[len-1] == FILLER && bytes[len-2] == FILLER {
    doubleFiller = true
    return 
  }
  if bytes[len-1] == FILLER {
    singleFiller = true
    return 
  }
  return
}

func revertFromBase64(bytes []byte) []byte {
  for i, item := range bytes {
    bytes[i] = getVal(item)
  }
  return bytes
}
```

</details>




<details>
  <summary> 主文件 </summary>

```golang
import (
	"os"
	"bufio"
	// "io"
	"fmt"
	"strings"
	"path/filepath"
	"regexp"
	hash "hash1"
	"io/ioutil"
	"bytes"
)

func main() {
	in := "D:/algorithm/btree/B-Tree Visualization.html"
	out := "d:/out.html"
	doMerge(in, out)
}

func doMerge(in, out string) {
	dir := filepath.Dir(in)
	fmt.Println(dir)

	linkRe := regexp.MustCompile(`href="([^"]*)"`)
	jsRe := regexp.MustCompile(`src="([^"]*)"`)

	var outContent bytes.Buffer
	f, _ := os.Open(in)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		
		line = strings.TrimSpace(line)
		
		if !strings.HasPrefix(line, "<link") && !strings.HasPrefix(line, "<script") {
			outContent.WriteString(line+"\n")
			continue
		}

		var url string
		var resArr [][]string
		if strings.HasPrefix(line, "<link") {
			resArr = linkRe.FindAllStringSubmatch(line, -1)
		}else if strings.HasPrefix(line, "<script") {
			resArr = jsRe.FindAllStringSubmatch(line, -1)
		}
		url = resArr[0][1]
		fpath := filepath.Join(dir, url)
		res := readContent(fpath)

		if strings.HasPrefix(line, "<link") {
			afterText := fmt.Sprintf(`href="data:text/css;base64,%s"`, res)
			line = linkRe.ReplaceAllString(line, afterText)
		}else if strings.HasPrefix(line, "<script") {
			afterText := fmt.Sprintf(`src="data:text/javascript;base64,%s"`, res)
			line = jsRe.ReplaceAllString(line, afterText)
		}
		outContent.WriteString(line+"\n")
	}
	ioutil.WriteFile(out, []byte(outContent.String()), 0666)
  
}

func readContent(path string) string {
	bytes, _ := ioutil.ReadFile(path)
	return hash.Base64Encode(bytes)
}
```

</details>
