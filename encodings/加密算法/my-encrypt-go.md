
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
)

func main() {
  fmt.Println(os.Args, len(os.Args))
  fmt.Printf(`exist(os.Args, "-d", "-e"): %v %v`, exist(os.Args, "-d", "-e"), "\n")
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
  keyBuf, randomArr := getKeyBuf(arg.key)
  keyLen := len(arg.key)
  
  outBuf := make([]byte, inputLen, inputLen)
  for _, key := range keyBuf {
    for i, b := range inputBuf {
      keyIndex := i % keyLen
      outBuf[i] = b ^ key[keyIndex]
    }
  }

  outBuf = hash.Base64Encode2Byte(outBuf)

  res := make([]byte, 0, keyLen + len(outBuf) + 1)
  res = append(res, []byte(keyLen)...)
  res = append(res, randomArr...)
  res = append(res, outBuf...)

  if arg.is4File() {
    ioutil.WriteFile(arg.file, res, 0666)
  }else{
    fmt.Println(string(res))
  }
}


func doDecrypt(arg info) {

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

func getKeyBuf(key string) ([][]byte, []byte) {
  length := len(key)
  keyBytes := []byte(key)
  var res [][]byte
  res = append(res, keyBytes)

  tmp1 := make([]byte, 0, length)
  for i:=length-1; i>=0; i-- {
    tmp1 = append(tmp1, keyBytes[i])
  }
  res = append(res, tmp1)

  tmp2 := make([]byte, 0, length)
  randomArr := getRandomArr(length)
  for _, index := range randomArr {
    tmp2 = append(tmp2, keyBytes[index])
  }
  res = append(res, tmp2)
  return res, randomArr
}

func parseArgs(args []string) info {
  var arg info
  for _, item := range args {
    if strings.HasPrefix(item, "-e") {
      arg.encrypt = true
    } else if strings.HasPrefix(item, "-d") {
      arg.encrypt = false
    } else if strings.HasPrefix(item, "-key") {
      if len(item) <= 4 {
        panic("args key can't be empty!")
      }
      arg.key = item[4:]
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
