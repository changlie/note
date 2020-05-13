> 转  [Base64的编码与解码](https://developer.mozilla.org/zh-CN/docs/Web/API/WindowBase64/Base64_encoding_and_decoding#The_.22Unicode_Problem.22)


Unicode 问题
由于 DOMString 是16位编码的字符串，所以如果有字符超出了8位ASCII编码的字符范围时，在大多数的浏览器中对Unicode字符串调用 window.btoa 将会造成一个 Character Out Of Range 的异常。有很多种方法可以解决这个问题：

- the first method consists in encoding JavaScript's native UTF-16 strings directly into base64 (fast, portable, clean)
- the second method consists in converting JavaScript's native UTF-16 strings to UTF-8 and then encode the latter into base64 (relatively fast, portable, clean)
- the third method consists in encoding JavaScript's native UTF-16 strings directly into base64 via binary strings (very fast, relatively portable, very compact)
- the fourth method consists in escaping the whole string (with UTF-8, see encodeURIComponent) and then encode it (portable, non-standard)
- the fifth method is similar to the second method, but uses third party libraries

###　Solution #1 – JavaScript's UTF-16 => base64

A very fast and widely useable way to solve the unicode problem is by encoding JavaScript native UTF-16 strings directly into base64. Please visit the URL data:text/plain;charset=utf-16;base64,OCY5JjomOyY8Jj4mPyY= for a demonstration (copy the data uri, open a new tab, paste the data URI into the address bar, then press enter to go to the page). This method is particularly efficient because it does not require any type of conversion, except mapping a string into an array. The following code is also useful to get an ArrayBuffer from a Base64 string and/or viceversa (see below).
```javascript
"use strict";

/*\
|*|
|*|  Base64 / binary data / UTF-8 strings utilities (#1)
|*|
|*|  https://developer.mozilla.org/en-US/docs/Web/API/WindowBase64/Base64_encoding_and_decoding
|*|
|*|  Author: madmurphy
|*|
\*/

/* Array of bytes to base64 string decoding */

function b64ToUint6 (nChr) {

  return nChr > 64 && nChr < 91 ?
      nChr - 65
    : nChr > 96 && nChr < 123 ?
      nChr - 71
    : nChr > 47 && nChr < 58 ?
      nChr + 4
    : nChr === 43 ?
      62
    : nChr === 47 ?
      63
    :
      0;

}

function base64DecToArr (sBase64, nBlockSize) {

  var
    sB64Enc = sBase64.replace(/[^A-Za-z0-9\+\/]/g, ""), nInLen = sB64Enc.length,
    nOutLen = nBlockSize ? Math.ceil((nInLen * 3 + 1 >>> 2) / nBlockSize) * nBlockSize : nInLen * 3 + 1 >>> 2, aBytes = new Uint8Array(nOutLen);

  for (var nMod3, nMod4, nUint24 = 0, nOutIdx = 0, nInIdx = 0; nInIdx < nInLen; nInIdx++) {
    nMod4 = nInIdx & 3;
    nUint24 |= b64ToUint6(sB64Enc.charCodeAt(nInIdx)) << 18 - 6 * nMod4;
    if (nMod4 === 3 || nInLen - nInIdx === 1) {
      for (nMod3 = 0; nMod3 < 3 && nOutIdx < nOutLen; nMod3++, nOutIdx++) {
        aBytes[nOutIdx] = nUint24 >>> (16 >>> nMod3 & 24) & 255;
      }
      nUint24 = 0;
    }
  }

  return aBytes;
}

/* Base64 string to array encoding */

function uint6ToB64 (nUint6) {

  return nUint6 < 26 ?
      nUint6 + 65
    : nUint6 < 52 ?
      nUint6 + 71
    : nUint6 < 62 ?
      nUint6 - 4
    : nUint6 === 62 ?
      43
    : nUint6 === 63 ?
      47
    :
      65;

}

function base64EncArr (aBytes) {

  var eqLen = (3 - (aBytes.length % 3)) % 3, sB64Enc = "";

  for (var nMod3, nLen = aBytes.length, nUint24 = 0, nIdx = 0; nIdx < nLen; nIdx++) {
    nMod3 = nIdx % 3;
    /* Uncomment the following line in order to split the output in lines 76-character long: */
    /*
    if (nIdx > 0 && (nIdx * 4 / 3) % 76 === 0) { sB64Enc += "\r\n"; }
    */
    nUint24 |= aBytes[nIdx] << (16 >>> nMod3 & 24);
    if (nMod3 === 2 || aBytes.length - nIdx === 1) {
      sB64Enc += String.fromCharCode(uint6ToB64(nUint24 >>> 18 & 63), uint6ToB64(nUint24 >>> 12 & 63), uint6ToB64(nUint24 >>> 6 & 63), uint6ToB64(nUint24 & 63));
      nUint24 = 0;
    }
  }

  return  eqLen === 0 ?
      sB64Enc
    :
      sB64Enc.substring(0, sB64Enc.length - eqLen) + (eqLen === 1 ? "=" : "==");

}
```
Tests
```javascript
var myString = "☸☹☺☻☼☾☿";

/* Part 1: Encode `myString` to base64 using native UTF-16 */

var aUTF16CodeUnits = new Uint16Array(myString.length);
Array.prototype.forEach.call(aUTF16CodeUnits, function (el, idx, arr) { arr[idx] = myString.charCodeAt(idx); });
var sUTF16Base64 = base64EncArr(new Uint8Array(aUTF16CodeUnits.buffer));

/* Show output */

alert(sUTF16Base64); // "OCY5JjomOyY8Jj4mPyY="

/* Part 2: Decode `sUTF16Base64` to UTF-16 */

var sDecodedString = String.fromCharCode.apply(null, new Uint16Array(base64DecToArr(sUTF16Base64, 2).buffer));

/* Show output */

alert(sDecodedString); // "☸☹☺☻☼☾☿"
```
