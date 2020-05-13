> 转： [Base64笔记](http://www.ruanyifeng.com/blog/2008/06/base64.html)

### 1. 定义简介

昨天的《MIME笔记》中提到，MIME主要使用两种编码转换方式***Quoted-printable和Base64***, 将8位的非英语字符转化为7位的ASCII字符。

虽然这样的初衷，是为了满足电子邮件中不能直接使用非ASCII码字符的规定，但是也有其他重要的意义：
```
a）所有的二进制文件，都可以因此转化为可打印的文本编码，使用文本软件进行编辑；   
b）能够对文本进行简单的加密。
```

### 2. Quoted-printable

首先，简单介绍一下Quoted-printable编码转换方式。它主要用于ACSII文本中夹杂少量非ASCII码字符的情况，不适合于转换纯二进制文件。

它规定将每一个8位的字节，转换为3个字符。

第一个字符是"="号，这是固定不变的。

后面二个字符是二个十六进制数，分别代表了这个字节前四位和后四位的数值。

举例来说，ASCII码中"换页键"（form feed）是12，二进制形式是00001100，写成十六进制就是0C，因此它的编码值为"=0C"。"="号的ASCII值是61，二进制形式是00111101，因为它的编码值是"=3D"。除了可打印的ASCII码以外，所有其他字符都必须用这种方式进行转换。

所有可打印的ASCII码字符（十进制值从33到126）都保持原样不变，"="（十进制值61）除外。

### 3. Base64

下面，详细介绍Base64的编码转换方式。

所谓Base64，就是说选出64个字符----小写字母a-z、大写字母A-Z、数字0-9、符号"+"、"/"（再加上作为垫字的"="，实际上是65个字符）----作为一个基本字符集。然后，其他所有符号都转换成这个字符集中的字符。

具体来说，转换方式可以分为四步。

第一步，将每三个字节作为一组，一共是24个二进制位。      
第二步，将这24个二进制位分为四组，每个组有6个二进制位。    
第三步，在每组前面加两个00，扩展成32个二进制位，即四个字节。    
第四步，根据下表，得到扩展后的每个字节的对应符号，这就是Base64的编码值。
```
　　0　A　　17　R　　　34　i　　　51　z

　　1　B　　18　S　　　35　j　　　52　0

　　2　C　　19　T　　　36　k　　　53　1

　　3　D　　20　U　　　37　l　　　54　2

　　4　E　　21　V　　　38　m　　　55　3

　　5　F　　22　W　　　39　n　　　56　4

　　6　G　　23　X　　　40　o　　　57　5

　　7　H　　24　Y　　　41　p　　　58　6

　　8　I　　　25　Z　　　42　q　　　59　7

　　9　J　　26　a　　　43　r　　　60　8

　　10　K　　27　b　　　44　s　　　61　9

　　11　L　　28　c　　　45　t　　　62　+

　　12　M　　29　d　　　46　u　　　63　/

　　13　N　　30　e　　　47　v

　　14　O　　31　f　　　48　w　　　

　　15　P　　32　g　　　49　x

　　16　Q　　33　h　　　50　y
```
因为，Base64将三个字节转化成四个字节，因此Base64编码后的文本，会比原文本大出三分之一左右。

### 4. base64实例

举一个具体的实例，演示英语单词Man如何转成Base64编码。
```
|Text content    |M        |a        |n  |
|ASCII           |77       |97       |110|
|Bit pattern     |01001101 |01100001 |01101110 |
|split 4 group   |00010011 |00010110 |00000101 |00101110 |
|Index           |19       |22       |5        |46|
|Base64-Encoded  |T        |W        |F        |u |
```



第一步，"M"、"a"、"n"的ASCII值分别是77、97、110，对应的二进制值是01001101、01100001、01101110，将它们连成一个24位的二进制字符串010011010110000101101110。

第二步，将这个24位的二进制字符串分成4组，每组6个二进制位：010011、010110、000101、101110。

第三步，在每组前面加两个00，扩展成32个二进制位，即四个字节：00010011、00010110、00000101、00101110。它们的十进制值分别是19、22、5、46。

第四步，根据上表，得到每个值对应Base64编码，即T、W、F、u。

因此，Man的Base64编码就是TWFu。

### 5. 特殊处理

如果字节数不足三，则这样处理：

a）二个字节的情况：将这二个字节的一共16个二进制位，按照上面的规则，转成三组，最后一组除了前面加两个0以外，后面也要加两个0。这样得到一个三位的Base64编码，再在末尾补上一个"="号。

比如，"Ma"这个字符串是两个字节，可以转化成三组00010011、00010110、00010000以后，对应Base64值分别为T、W、E，再补上一个"="号，因此"Ma"的Base64编码就是TWE=。

b）一个字节的情况：将这一个字节的8个二进制位，按照上面的规则转成二组，最后一组除了前面加二个0以外，后面再加4个0。这样得到一个二位的Base64编码，再在末尾补上两个"="号。

比如，"M"这个字母是一个字节，可以转化为二组00010011、00010000，对应的Base64值分别为T、Q，再补上二个"="号，因此"M"的Base64编码就是TQ==。

### 6. 汉字转Base64

再举一个中文的例子，汉字"严"如何转化成Base64编码？

这里需要注意，汉字本身可以有多种编码，比如gb2312、utf-8、gbk等等，每一种编码的Base64对应值都不一样。下面的例子以utf-8为例。

首先，"严"的utf-8编码为E4B8A5，写成二进制就是三字节的"11100100 10111000 10100101"。将这个24位的二进制字符串，按照第3节中的规则，转换成四组一共32位的二进制值"00111001 00001011 00100010 00100101"，相应的十进制数为57、11、34、37，它们对应的Base64值就为5、L、i、l。

所以，汉字"严"（utf-8编码）的Base64值就是5Lil。

### 7. php  Base64

在PHP语言中，有一对专门的函数用于Base64转换：base64_encode()用于编码、base64_decode()用于解码。

这对函数的特点是，它们不管输入文本的编码是什么，都会按照规则进行Base64编码。因此，如果你想得到utf-8编码下的Base64对应值，你就必须自己保证，输入的文本是utf-8编码的。

### 8. js Base64

这一节介绍如何用Javascript语言进行Base64编码。

首先，假定网页的编码是utf-8，我们希望对于同样的字符串，用PHP和Javascript可以得到同样的Base64编码。

这里就会产生一个问题。因为Javascript内部的字符串，都以utf-16的形式进行保存，因此编码的时候，我们首先必须将utf-8的值转成utf-16再编码，解码的时候，则是解码后还需要将utf-16的值转回成utf-8。

网上已经有人写好了现成的Javascript函数：

```javascript
/* utf.js - UTF-8 <=> UTF-16 convertion
*
* Copyright (C) 1999 Masanao Izumo <iz@onicos.co.jp>
* Version: 1.0
* LastModified: Dec 25 1999
* This library is free. You can redistribute it and/or modify it.
*/

/*
* Interfaces:
* utf8 = utf16to8(utf16);
* utf16 = utf8to16(utf8);
*/

function utf16to8(str) {
  var out, i, len, c;

  out = "";
  len = str.length;
  for (i = 0; i < len; i++) {
    c = str.charCodeAt(i);
    if ((c >= 0x0001) && (c <= 0x007F)) {
      out += str.charAt(i);
    } else if (c > 0x07FF) {
      out += String.fromCharCode(0xE0 | ((c >> 12) & 0x0F));
      out += String.fromCharCode(0x80 | ((c >> 6) & 0x3F));
      out += String.fromCharCode(0x80 | ((c >> 0) & 0x3F));
    } else {
      out += String.fromCharCode(0xC0 | ((c >> 6) & 0x1F));
      out += String.fromCharCode(0x80 | ((c >> 0) & 0x3F));
    }
  }
  return out;
}

function utf8to16(str) {
  var out, i, len, c;
  var char2, char3;

  out = "";
  len = str.length;
  i = 0;
  while (i < len) {
    c = str.charCodeAt(i++);
    switch (c >> 4) {
      case 0: case 1: case 2: case 3: case 4: case 5: case 6: case 7:
        // 0xxxxxxx
        out += str.charAt(i - 1);
        break;
      case 12: case 13:
        // 110x xxxx 10xx xxxx
        char2 = str.charCodeAt(i++);
        out += String.fromCharCode(((c & 0x1F) << 6) | (char2 & 0x3F));
        break;
      case 14:
        // 1110 xxxx 10xx xxxx 10xx xxxx
        char2 = str.charCodeAt(i++);
        char3 = str.charCodeAt(i++);
        out += String.fromCharCode(((c & 0x0F) << 12) |
          ((char2 & 0x3F) << 6) |
          ((char3 & 0x3F) << 0));
        break;
    }
  }

  return out;
}
```
上面的这段代码中定义了两个函数，utf16to8()用于将utf-16转成utf-8，utf8to16用于将utf-8转成utf-16。

下面才是真正用于base64编码的函数。
```javascript
/* Copyright (C) 1999 Masanao Izumo <iz@onicos.co.jp>
* Version: 1.0
* LastModified: Dec 25 1999
* This library is free. You can redistribute it and/or modify it.
*/

/*
* Interfaces:
* b64 = base64encode(data);
* data = base64decode(b64);
*/


var base64EncodeChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/";
var base64DecodeChars = new Array(
  -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
  -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
  -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 62, -1, -1, -1, 63,
  52, 53, 54, 55, 56, 57, 58, 59, 60, 61, -1, -1, -1, -1, -1, -1,
  -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14,
  15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, -1, -1, -1, -1, -1,
  -1, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
  41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, -1, -1, -1, -1, -1);

function base64encode(str) {
  var out, i, len;
  var c1, c2, c3;

  len = str.length;
  i = 0;
  out = "";
  while (i < len) {
    c1 = str.charCodeAt(i++) & 0xff;
    if (i == len) {
      out += base64EncodeChars.charAt(c1 >> 2);
      out += base64EncodeChars.charAt((c1 & 0x3) << 4);
      out += "==";
      break;
    }
    c2 = str.charCodeAt(i++);
    if (i == len) {
      out += base64EncodeChars.charAt(c1 >> 2);
      out += base64EncodeChars.charAt(((c1 & 0x3) << 4) | ((c2 & 0xF0) >> 4));
      out += base64EncodeChars.charAt((c2 & 0xF) << 2);
      out += "=";
      break;
    }
    c3 = str.charCodeAt(i++);
    out += base64EncodeChars.charAt(c1 >> 2);
    out += base64EncodeChars.charAt(((c1 & 0x3) << 4) | ((c2 & 0xF0) >> 4));
    out += base64EncodeChars.charAt(((c2 & 0xF) << 2) | ((c3 & 0xC0) >> 6));
    out += base64EncodeChars.charAt(c3 & 0x3F);
  }
  return out;
}

function base64decode(str) {
  var c1, c2, c3, c4;
  var i, len, out;

  len = str.length;
  i = 0;
  out = "";
  while (i < len) {
    /* c1 */
    do {
      c1 = base64DecodeChars[str.charCodeAt(i++) & 0xff];
    } while (i < len && c1 == -1);
    if (c1 == -1)
      break;

    /* c2 */
    do {
      c2 = base64DecodeChars[str.charCodeAt(i++) & 0xff];
    } while (i < len && c2 == -1);
    if (c2 == -1)
      break;

    out += String.fromCharCode((c1 << 2) | ((c2 & 0x30) >> 4));

    /* c3 */
    do {
      c3 = str.charCodeAt(i++) & 0xff;
      if (c3 == 61)
        return out;
      c3 = base64DecodeChars[c3];
    } while (i < len && c3 == -1);
    if (c3 == -1)
      break;

    out += String.fromCharCode(((c2 & 0XF) << 4) | ((c3 & 0x3C) >> 2));

    /* c4 */
    do {
      c4 = str.charCodeAt(i++) & 0xff;
      if (c4 == 61)
        return out;
      c4 = base64DecodeChars[c4];
    } while (i < len && c4 == -1);
    if (c4 == -1)
      break;
    out += String.fromCharCode(((c3 & 0x03) << 6) | c4);
  }
  return out;
}
```
上面代码中的base64encode()用于编码，base64decode()用于解码。

因此，对utf-8字符进行编码要这样写：
```javascript
sEncoded=base64encode(utf16to8(str));
```
然后，解码要这样写：
```javascript
sDecoded=utf8to16(base64decode(sEncoded));
```
