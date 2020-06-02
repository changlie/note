```golang
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
  if res[newArrLen-1] == 0 && res[newArrLen-2] == 0 {
    return res[:newArrLen-2]
  }else if res[newArrLen-1] == 0 {
    return res[:newArrLen-1]
  }else {
    return res
  }
}

func revertFromBase64(bytes []byte) []byte {
  for i, item := range bytes {
    bytes[i] = getVal(item)
  }
  return bytes
}
```
