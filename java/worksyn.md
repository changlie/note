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
import java.io.*;

public class SplitClassFile extends Object{

    public static void main(String[] args) {
        String projectRootPath = System.getProperty("user.dir");
        String sourceDir = projectRootPath+"/src/main/java";

        String pkgDir = null;
        String className = null;
        StringBuilder buf = new StringBuilder();

        try(BufferedReader r = new BufferedReader(new FileReader("/home/changlie/ws/git/note/algorithm/BTree/java-v2.md"))){

            String line = null;
            while ((line = r.readLine())!=null){
                if(className==null && "".equals(line)){
                    continue;
                }
                if(line.startsWith("```")){
                    continue;
                }
                if(pkgDir == null && line.startsWith("package")){
                    String tmp = line.substring("package".length()).trim();
                    pkgDir = sourceDir+"/"+ tmp.substring(0, tmp.length()-1).trim().replaceAll("\\.", "/");
                    System.out.println("pkgDir: "+pkgDir);
                    File f = new File(pkgDir);
                    if(!f.exists()) f.mkdirs();
                }
                boolean isclass = line.startsWith("public class");
                boolean isInterface = line.startsWith("public interface");
                if(className == null && (isclass || isInterface)){
                    int startIndex = isclass ? "public class".length() : "public interface".length();
                    String tmp = line.substring(startIndex).trim();
                    tmp = tmp.substring(0, tmp.length()-1).trim();
                    int implIndex = tmp.indexOf("implements");
                    if(implIndex>1){
                        tmp = tmp.substring(0, implIndex).trim();
                    }
                    int extIndex = tmp.indexOf("extends");
                    if(extIndex>1){
                        tmp = tmp.substring(0, extIndex);
                    }
                    className = pkgDir+"/"+tmp+".java";
                    System.out.println("className: "+className);
                }
                if(line.equals("/////////////////////////////////////////////////////////////////////////////////////////////")){
                   BufferedWriter w = new BufferedWriter(new FileWriter(className));
                   w.write(buf.toString());
                   w.flush();
                   w.close();
                   className = null;
                   buf.delete(0, buf.length());
                   continue;
                }

                buf.append(line).append(System.lineSeparator());

            }
        }catch (Exception e){
            e.printStackTrace();
        }
    }
}
```
