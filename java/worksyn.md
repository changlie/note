### merge
```java
import com.huawei.it.util.Futil;

import java.io.File;

public class MergeClassFile {
    public static void main(String[] args) {
        String classPath = "D:\\workspace\\eclipse_ws\\ari\\src\\main\\java\\com\\huawei\\it\\demo\\ari\\btree\\v1";
        String outFile = "d:/MergeClassFileResult_"+System.currentTimeMillis()+".txt";

        StringBuilder res = new StringBuilder();
        res.append("```java").append(System.lineSeparator());
        File f = new File(classPath);
        for(String cf: f.list()){
            String s = Futil.readStr(classPath+"/"+cf);

            res.append(System.lineSeparator());
            res.append(s);
            res.append("\r\n/////////////////////////////////////////////////////////////////////////////////////////////\r\n");
        }
        res.append("```").append(System.lineSeparator());
        Futil.out(outFile, res.toString());
    }
}
```

### split
```java

```
