### 对二进制文件进行加码解码

```java

import java.io.File;
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.FileOutputStream;
import java.io.IOException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class BinaryFileConvertByBase64 {
    static void merge() throws Exception {
        long startTime = System.currentTimeMillis();

        String srcDir = "D:/bakup/BCompare-3.3/BCompare/";
        String destFile = "D:/bakup/BCompare-3.3/BCompare.exe";

        File dir = new File(srcDir);
        File[] files = dir.listFiles();
        List<byte[]> list = new ArrayList<>(files.length);
        int total = 0;
        for (File f : files) {
            System.out.println(f.getAbsolutePath());
            byte[] bytes = readBytes(f);
            int countLF = countLF(bytes);
            total += bytes.length - countLF;
            list.add(bytes);
        }
        byte[] buf = new byte[total];
        int currentLen = 0;
        for (byte[] bytes : list) {
            int len = bytes.length;

            int startIndex = 0;
            for (int i = 0; i < len; i++) {
                boolean isEnd = i == len - 1;
                byte b = bytes[i];
                if (b != '\n' && !isEnd) {
                    continue;
                }

                int copyLen = isEnd ? len - startIndex : i - startIndex;
                System.arraycopy(bytes, startIndex, buf, currentLen, copyLen);

                startIndex = i + 1;
                currentLen += copyLen;
            }
        }
        buf = Base64.base64Decode(buf);

        save(destFile, buf);

        System.out.println("merge finish! spend: " + (System.currentTimeMillis() - startTime) + "ms");
    }

    static int countLF(byte[] bytes) {
        int count = 0;
        for (byte b : bytes) {
            if (b == '\n') {
                count++;
            }
        }
        return count;
    }

    static void split() throws Exception {
        long startTime = System.currentTimeMillis();

        String srcFile = "d:/bakup/BCompare-3.3/BCompare-zh-3.3.3.14128.exe";
        String destDir = "D:/bakup/BCompare-3.3/BCompare/";

        int block = 1024 * 1024;
        byte[] buf = readBytes(srcFile);

        String base64Str = Base64.base64Encode(buf);

        byte[] base64Bytes = base64Str.getBytes();

        int length = base64Bytes.length;
        int blockCount = length / block;
        boolean isNotDivide = false;
        if (length % block > 0) {
            isNotDivide = true;
            blockCount++;
        }

        for (int i = 0; i < blockCount; i++) {
            boolean isEnd = i == blockCount - 1;
            int startIndex = block * i;
            int copyLen = block;
            if (isEnd && isNotDivide) {
                copyLen = length % block;
            }
            byte[] blockArr = new byte[copyLen];
            System.arraycopy(base64Bytes, startIndex, blockArr, 0, copyLen);

            String outFile = destDir + i;

            byte[] res = addLineSeperator(blockArr);
            save(outFile, res);
        }

        System.out.println("split finish! spend: " + (System.currentTimeMillis() - startTime) + "ms");
    }

    static byte[] readBytes(String srcFile) throws FileNotFoundException, IOException {
        return readBytes(new File(srcFile));
    }

    static byte[] readBytes(File srcFile) throws FileNotFoundException, IOException {
        FileInputStream fis = new FileInputStream(srcFile);
        int total = fis.available();
        byte[] buf = new byte[total];
        fis.read(buf);
        fis.close();
        return buf;
    }

    static void save(String outFile, byte[] buf) throws Exception {
        FileOutputStream fos = new FileOutputStream(outFile);
        fos.write(buf);
        fos.flush();
        fos.close();
    }

    static void save(String outFile, byte[] buf, int length) throws Exception {
        FileOutputStream fos = new FileOutputStream(outFile);
        fos.write(buf, 0, length);
        fos.flush();
        fos.close();
    }

    static byte[] addLineSeperator(byte[] src) {
        int length = src.length;
        int lineLen = 128;
        int blockCount = length / lineLen;
        boolean isNotDivide = false;
        if (length % lineLen > 0) {
            isNotDivide = true;
            blockCount++;
        }
        int accumulate = blockCount - 1;
        int currentLen = 0;
        byte[] res = new byte[length + accumulate];
        for (int i = 0; i < blockCount; i++) {
            boolean isEnd = i == blockCount - 1;
            int startIndex = lineLen * i;
            int copyLen = lineLen;
            if (isEnd && isNotDivide) {
                copyLen = length % lineLen;
            }
            System.arraycopy(src, startIndex, res, currentLen, copyLen);
            if (!isEnd) {
                res[currentLen + copyLen] = '\n';
            }
            currentLen += copyLen + 1;
        }
        return res;
    }

    public static void main(String[] args) throws Exception {
        split();
        merge();
        // testAddLineSeperator();
        // testArrayCopy();
    }

    static void testArrayCopy() {
        String[] str = {"a", "b", "c", "d", "f", "g", "h"};
        String[] sub = new String[4];
        System.arraycopy(str, 2, sub, 0, 4);
        System.out.println(Arrays.toString(sub));
    }

    static void testAddLineSeperator() {
        String raw = "iVBOR w0KGg oAAAA NSUhE UgAAA ZAAAA GQBAM AAABy kSv";
        byte[] bytes = raw.getBytes();
        System.out.println("init size: " + bytes.length);
        byte[] res = addLineSeperator(bytes);
        System.out.println(new String(res));
    }
}

```
