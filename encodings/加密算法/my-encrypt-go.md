
```golang

package main

import (
  "fmt"
  "math/rand"
  "os"
  "time"
  "path/filepath"
  // "reflect"
  "strings"
  "io/ioutil"
  hash "hash1"
  "bytes"
  // "strconv"
)

func main() {
  if len(os.Args)<3 || !exist(os.Args, "-d", "-e") || !exist(os.Args, "-s", "-f") || !exist(os.Args, "-key") {
    base := os.Args[0]
    baseUrl := filepath.Dir(base)
    cmd := base[len(baseUrl)+1:]
    fmt.Printf("Usage: %v (-d | -e) -key=key4encrypt (-s=string4input | -f=file4input)....\n", cmd)
    return
  }
  arg := parseArgs(os.Args[1:])

  if arg.encrypt {
    doEncrypt(arg)
  } else {
    doDecrypt(arg)
  }
}

func doEncrypt(arg info) {
  inputBuf := getInput(arg)
  inputLen := len(inputBuf)
  keyBuf, randomArr := getKeyBuf4Encrypt(arg.key)
  keyLen := len(arg.key)
  
  // log := make([]string, inputLen, inputLen)
  outBuf := make([]byte, inputLen, inputLen)
  for _, key := range keyBuf {
    for i, b := range inputBuf {
      keyIndex := i % keyLen
      outBuf[i] = b ^ key[keyIndex]
      // if keysIndex == 0 {
      //   log[i] += fmt.Sprintf(" %v:", strconv.Itoa(i))
      // }
      // log[i] += fmt.Sprintf("  %s|", string(key[keyIndex]))
    }
  }

  res := make([]byte, 0, keyLen + len(outBuf) + 1)
  res = append(res, byte(keyLen))
  res = append(res, randomArr...)
  res = append(res, outBuf...)

  res = hash.Base64Encode2Byte(res)

  // printLog(log)

  if arg.is4File() {
    res = insertLineSeperator(res)
    ioutil.WriteFile(arg.file, res, 0666)
  }else{
    fmt.Println("-|"+string(res)+"|-")
  }
}

func insertLineSeperator(src []byte) []byte {
  lineLen := 128
  srcLen := len(src)
  seperatorCount := srcLen / lineLen
  if srcLen % lineLen > 0 {
    seperatorCount ++
  }
  res := make([]byte, 0, srcLen+seperatorCount)
  for i:=1; i<=seperatorCount; i++ {
    startIndex := (i-1)*lineLen
    endIndex := i*lineLen
    if i==seperatorCount {
      endIndex = srcLen
    }
    if i>1 {
      res = append(res, '\n')
    }
    res = append(res, src[startIndex:endIndex]...)
  }
  return res
}

func printLog(msgs []string) {
  for _, msg := range msgs {
    fmt.Println("msg:", msg)
  }
}


func doDecrypt(arg info) {
  
  inputBuf := getInput(arg)
  inputBuf = clearLineSeperator(inputBuf)
  inputBuf = hash.Base64Decode2Byte(inputBuf)
  inputBuf, randomArr := parseEncryptStr(inputBuf)
  keyBytes := []byte(arg.key)
  keyBuf := getKeyBuf(keyBytes)
  randomKeyArr := getRandomKeyArr(keyBytes, randomArr)
  keyBuf = append(keyBuf, randomKeyArr)
  keyLen := len(keyBytes)
  msgLen := len(inputBuf)


  // log := make([]string, msgLen, msgLen)
  res := make([]byte, msgLen, msgLen)
  for _, key := range keyBuf {
    for i, b := range inputBuf {
      keyIndex := i % keyLen
      res[i] = b ^ key[keyIndex]
      
      // if keysIndex == 0 {
      //   log[i] += fmt.Sprintf(" %v:", strconv.Itoa(i))
      // }
      // log[i] += fmt.Sprintf("  %s|", string(key[keyIndex]))
    }
  }
  
  // printLog(log)

  if arg.is4File() {
    ioutil.WriteFile(arg.file, res, 0666)
  }else{
    fmt.Println("-|"+string(res)+"|-")
  }
}

func clearLineSeperator(src []byte) []byte {
  return bytes.ReplaceAll(src, []byte("\n"), []byte(""))
}

func parseEncryptStr(src []byte) ([]byte, []byte) {
  randomArr := src[1:src[0]+1]
  msgBuf := src[src[0]+1:]
  return msgBuf, randomArr
}

func getInput(arg info) []byte {
  var buf []byte
  if arg.is4File() {
    buf, _ = ioutil.ReadFile(arg.file)
  }else{
    buf = []byte(arg.str)
  }
  return buf
}

type info struct {
  encrypt bool
  key string
  file string
  str string
}

func (this *info) is4File() bool {
  return len(this.file) > 0
}

func (this *info) isHandleCurrentDir() bool {
  return this.file == "."
}

func getKeyBuf4Encrypt(key string) ([][]byte, []byte) {
  keyBytes := []byte(key)
  res := getKeyBuf(keyBytes)

  randomArr := getRandomArr(len(key))
  randomKeyArr := getRandomKeyArr(keyBytes, randomArr)
  res = append(res, randomKeyArr)

  return res, randomArr
}

func getRandomKeyArr(keyBytes, randomArr []byte) []byte {
  keyLen := len(keyBytes)
  randomKeyArr := make([]byte, 0, keyLen)
  for i, index := range randomArr {
    if i>= keyLen {
      break
    }
    randomKeyArr = append(randomKeyArr, keyBytes[index])
  }
  return randomKeyArr
}

func getKeyBuf(keyBytes []byte) ([][]byte) {
  length := len(keyBytes)
  
  var res [][]byte
  res = append(res, keyBytes)

  tmp1 := make([]byte, 0, length)
  for i:=length-1; i>=0; i-- {
    tmp1 = append(tmp1, keyBytes[i])
  }
  res = append(res, tmp1)
  return res
}

func parseArgs(args []string) info {
  var arg info
  for _, item := range args {
    if strings.HasPrefix(item, "-e") {
      arg.encrypt = true
    } else if strings.HasPrefix(item, "-d") {
      arg.encrypt = false
    } else if strings.HasPrefix(item, "-key") {
      if len(item) <= 5 {
        panic("args key can't be empty!")
      }
      arg.key = item[5:]
    } else if strings.HasPrefix(item, "-s") {
      arg.str = item[3:]
    } else if strings.HasPrefix(item, "-f") {
      if len(item) == 2 {
        arg.file = "."
      } else {
        arg.file = item[3:]
      }
    }
  }
  return arg
}

func exist(src []string, target ...string) bool {
  for _, item := range src {
    for _, matchItem := range target {
      if strings.HasPrefix(item, matchItem) {
        return true
      }
    }
  }
  return false
}


func getRandomArr(length int) []byte {
  set := make(map[int]struct{})
  void := struct{}{}
  rand.Seed(time.Now().UnixNano())
  arr := make([]byte, 0, length)
  for len(set) < length {
    item := rand.Intn(length)
    if _, ok := set[item]; !ok {
      arr = append(arr, byte(item))
      set[item] = void
    }
  }
  return arr
}

```
