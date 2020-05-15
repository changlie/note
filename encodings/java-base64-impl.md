```java

import java.io.FileOutputStream;
import java.io.RandomAccessFile;

public class Base64 {
    // 第一步，将每三个字节作为一组，一共是24个二进制位。
    // 第二步，将这24个二进制位分为四组，每个组有6个二进制位。
    // 第三步，在每组前面加两个00，扩展成32个二进制位，即四个字节。
    // 第四步，根据下表，得到扩展后的每个字节的对应符号，这就是Base64的编码值。
    static byte[] base64 =
        {'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V',
            'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q',
            'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '+', '/',};

    static byte token = '=';

    static byte getVal(byte b) {
        int len = base64.length;
        for (int i = 0; i < len; i++) {
            if (base64[i] == b) {
                return (byte) i;
            }
        }
        if (token == b) {
            return 0;
        }
        throw new RuntimeException("not found origin value map: " + ((char) b));
    }

    static String toBinaryString4Base64(String s) {
        String res = "";
        int length = s.length();
        for (int i = 0; i < length; i++) {
            if (i > 0) {
                res += ", ";
            }
            res += toBinaryString(getVal((byte) s.charAt(i)));
        }
        return res;
    }

    static String toBinaryStringRAW(String s) {
        String res = "";
        int length = s.length();
        for (int i = 0; i < length; i++) {
            if (i > 0) {
                res += ", ";
            }
            res += toBinaryString((byte) s.charAt(i));
        }
        return res;
    }

    static String toBinaryString(byte b) {
        return String.format("%8s", Integer.toBinaryString(b & 0xff)).replaceAll(" ", "0");
    }

    static String fileToBase64(String filePath) throws Exception {
        RandomAccessFile file = new RandomAccessFile(filePath, "r");
        int len = (int) file.length();
        byte[] buf = new byte[len];
        file.read(buf, 0, len);

        String res = base64Encode(buf);
        file.close();
        return res;
    }

    static void base64ToFile(String base64, String outFile) throws Exception {
        byte[] buf = base64Decode(base64);
        FileOutputStream fos = new FileOutputStream(outFile);
        fos.write(buf);
        fos.flush();
        fos.close();
    }

    static byte[] base64Decode(String string) {
        byte[] buf = base64ToBytes(string);
        int len = buf.length;
        byte[] res = new byte[len / 4 * 3];
        for (int i = 0; i < len; i += 4) {
            int resIndex = i / 4 * 3;

            byte b1 = buf[i];
            byte b2 = buf[i + 1];
            byte b3 = buf[i + 2];
            byte b4 = buf[i + 3];

            res[resIndex] = (byte) ((b1 << 2 | ((b2 >> 4) & 0x03)) & 0xff);
            res[resIndex + 1] = (byte) ((b2 << 4 | ((b3 >> 2) & 0x0f)) & 0xff);
            res[resIndex + 2] = (byte) ((b3 << 6 | b4) & 0xff);
        }
        return res;
    }

    static byte[] base64Decode(byte[] buf) {
        buf = base64ToBytes(buf);
        int len = buf.length;
        byte[] res = new byte[len / 4 * 3];
        for (int i = 0; i < len; i += 4) {
            int resIndex = i / 4 * 3;

            byte b1 = buf[i];
            byte b2 = buf[i + 1];
            byte b3 = buf[i + 2];
            byte b4 = buf[i + 3];

            res[resIndex] = (byte) ((b1 << 2 | ((b2 >> 4) & 0x03)) & 0xff);
            res[resIndex + 1] = (byte) ((b2 << 4 | ((b3 >> 2) & 0x0f)) & 0xff);
            res[resIndex + 2] = (byte) ((b3 << 6 | b4) & 0xff);
        }
        return res;
    }

    static byte[] base64ToBytes(byte[] buf) {
        int length = buf.length;
        for (int i = 0; i < length; i++) {
            buf[i] = getVal(buf[i]);
        }
        return buf;
    }

    static byte[] base64ToBytes(String s) {
        int length = s.length();
        byte[] buf = new byte[length];
        for (int i = 0; i < length; i++) {
            buf[i] = getVal((byte) s.charAt(i));
        }
        return buf;
    }

    static String base64Encode(byte[] buf) {
        int len = buf.length;
        StringBuilder res = new StringBuilder();
        for (int i = 0; i < len;) {
            int nextIndex = i + 3;
            boolean isEnd = nextIndex >= len;

            if (!isEnd || (isEnd && (len - 1) % 3 == 2)) {
                byte b1 = (byte) (buf[i] >> 2 & 0x3f);
                byte b2 = (byte) ((buf[i] << 4 | (buf[i + 1] >> 4 & 0x0f)) & 0x3f);
                byte b3 = (byte) ((buf[i + 1] << 2 | (buf[i + 2] >> 6 & 0x03)) & 0x3f);
                byte b4 = (byte) (buf[i + 2] & 0x3f);

                res.append((char) base64[b1]).append((char) base64[b2]).append((char) base64[b3]).append(
                    (char) base64[b4]);
            } else {
                if ((len - 1) % 3 == 1) {
                    // a）二个字节的情况：将这二个字节的一共16个二进制位，按照上面的规则，转成三组，最后一组除了前面加两个0以外，后面也要加两个0。
                    // 这样得到一个三位的Base64编码，再在末尾补上一个"="号。
                    byte b1 = (byte) (buf[i] >> 2 & 0x3f);
                    byte b2 = (byte) ((buf[i] << 4 | (buf[i + 1] >> 4 & 0x0f)) & 0x3f);
                    byte b3 = (byte) (buf[i + 1] << 2 & 0x3c);
                    res.append((char) base64[b1]).append((char) base64[b2]).append((char) base64[b3]).append(
                        (char) token);
                } else if ((len - 1) % 3 == 0) {
                    // b）一个字节的情况：将这一个字节的8个二进制位，按照上面的规则转成二组，最后一组除了前面加二个0以外，后面再加4个0。
                    // 这样得到一个二位的Base64编码，再在末尾补上两个"="号。
                    byte b1 = (byte) (buf[i] >> 2 & 0x3f);
                    byte b2 = (byte) (buf[i] << 4 & 0x30);
                    res.append((char) base64[b1]).append((char) base64[b2]).append((char) token).append((char) token);
                } else {
                    throw new RuntimeException("unknow Exception! " + i);
                }
            }

            i = nextIndex;
        }
        return res.toString();
    }

    public static void main(String[] args) throws Exception {
        // toBase64Test2("man");
        // System.out.println("============");
        // toBase64Test2("DB");
        // System.out.println("============");
        // toBase64Test2("K");
        // System.out.println("============");
        // batchBase64DecodeTest();

        // String string = fileToBase64("D:\\ator.png");
        // System.out.println(string);
        // base64ToFile(string, "d:/gaodao.png");
        System.out.println((byte) '\n');
        System.out.println("hello".getBytes().length);
        System.out.println("hello".getBytes("utf-8").length);
        System.out.println("hello".getBytes("utf-16").length);

        String s = base64Encode("hello".getBytes());
        System.out.println(s);
        System.out.println(s.length());
    }

    static void batchBase64DecodeTest() {
        base64DecodeTest("changlie");
        System.out.println("=================");
        base64DecodeTest("JS");
        System.out.println("=================");
        base64DecodeTest("Q");
    }

    private static void base64DecodeTest(String raw) {
        String base64_ = base64Encode(raw.getBytes());
        System.out.println("base64: " + base64_);
        byte[] bytes = base64Decode(base64_);

        System.out.println(new String(bytes));
    }

    static void toBase64Test2(String src) {
        System.out.println(toBinaryStringRAW(src));
        String s1 = base64Encode(src.getBytes());
        System.out.println(s1);
        String s2 = toBinaryString4Base64(s1);
        System.out.println(s2);
    }

    static void toBase64Test1() {
        // System.out.println(toBinaryString((byte) 0b00001010));
        // System.out.println(toBinaryString((byte) 0b01101010));
        System.out.println(Integer.toBinaryString(0b00111111 & 0xff));
        System.out.println(Integer.toBinaryString(0x3f & 0xff));
        byte[] buf = "man".getBytes();
        System.out.println(buf.length);
        String s = toBinaryString(buf[0]);
        s += ", " + toBinaryString(buf[1]);
        s += ", " + toBinaryString(buf[2]);
        System.out.println("bit show: " + s);
        int i = 0;

        byte b1 = (byte) (buf[i] >> 2 & 0x3f);
        byte b2 = (byte) ((buf[i] << 4 | (buf[i + 1] >> 4 & 0x0f)) & 0x3f);
        byte b3 = (byte) ((buf[i + 1] << 2 | (buf[i + 2] >> 6 & 0x03)) & 0x3f);
        byte b4 = (byte) (buf[i + 2] & 0x3f);
        String s1 = toBinaryString(b1);
        s1 += ", " + toBinaryString(b2);
        s1 += ", " + toBinaryString(b3);
        s1 += ", " + toBinaryString(b4);
        System.out.println("after split: " + s1);
    }
}

```
