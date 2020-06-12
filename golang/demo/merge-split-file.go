```golang
package main

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "regexp"
    "strings"
)

func main() {
    mergefile()
    splitfile()
}

func splitfile() {
    src := "d:/setfile"
    dest := "d:/pt-unpkg"

    if _, err := os.Stat(dest); os.IsNotExist(err) {
        os.MkdirAll(dest, 0666)
    }

    bs, _ := ioutil.ReadFile(src)


    re, err := regexp.Compile("(?s:#{5}([^#]+)#{5}([^~]+?)~{20})")
    fmt.Println(err)
    arrs := re.FindAllStringSubmatch(string(bs), -1)
    for _, arr := range arrs {
        var fname, fcontent string
        for i, item := range arr {
            if i == 1 {
                fname = item
            }
            if i == 2 {
                fcontent = item
            }
        }
        abspath := filepath.Join(dest, fname)
        mkdir(abspath)
        contentBuf := []byte(fcontent)
        contentBuf = bytes.ReplaceAll(contentBuf, []byte{0}, []byte{'~'}) // 反转义
        ioutil.WriteFile(abspath, contentBuf, 0666)
    }

}

func mkdir(path string) {
    dir := filepath.Dir(path)
    if dir == "" {
        return
    }
    if _, err := os.Stat(dir); os.IsNotExist(err) {
        os.MkdirAll(dir, 0666)
    }
}

func mergefile() {
    srcDir := `D:\software\golangLocalRepo\src\changlie\pithy\`
    destFile := `d:/setfile`

    walker4read(srcDir, destFile, srcDir)
}

func walker4read(src, dest, root string) {
    fs, _ := ioutil.ReadDir(src)
    for _, f := range fs {
        fname := f.Name()
        path := filepath.Join(src, fname)
        if f.IsDir() {
            if strings.HasPrefix(fname, ".") {
                continue
            }
            walker4read(path, dest, root)
        } else {
            bs, _ := ioutil.ReadFile(path)
            bs = bytes.ReplaceAll(bs, []byte{'~'}, []byte{0}) // 转义
            var buf bytes.Buffer
            buf.WriteString("#####")
            buf.WriteString(getRelativePath(root, path))
            buf.WriteString("#####")
            buf.Write(bs)
            buf.WriteString("~~~~~~~~~~~~~~~~~~~~\r\n")
            destFile, _ := os.OpenFile(dest, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
            destFile.Write(buf.Bytes())
            destFile.Sync()
            destFile.Close()
        }
    }
}

func getRelativePath(rootDir, abspath string) string {
    rootDirLen := len(rootDir)
    return abspath[rootDirLen:]
}
```
