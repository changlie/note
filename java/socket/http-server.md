
<details>
<summary>解析http请求数据</summary>

```java
import java.io.IOException;
import java.io.InputStream;
import java.io.UnsupportedEncodingException;
import java.util.*;

public class HttpRequest implements HttpMessage{
	private InputStream inputStream;
	private byte[] buffer=new byte[2048];
	private String requestLine;
	private Map<String,String> headersMap=new HashMap<>();
	private int contentLength;
	private boolean headerEnd;
	private StringBuffer header=new StringBuffer();
	private List<Byte> entity;
	
	public HttpRequest(InputStream inputStream) {
		this.inputStream=inputStream;
	}
	
	public void parseRequest() throws IOException{
		int bufferSize=inputStream.read(buffer);
		for(int index=0;index<bufferSize;index++){
			byte byt = buffer[index];
			char c=(char) buffer[index];
			if(c==LF){
				int length=header.length();
				if(header.charAt(length-1)==CR){
					if(requestLine==null){//第一个CRLF是请求行
						requestLine=header.toString();
					}else{
						int crlf=header.lastIndexOf(CRLF);
						String headerLine=header.substring(crlf+2,length-1);
						if(!"".equals(headerLine)){//请求头
							putHeader(headerLine);
						}else{//以CRLF结尾的空行标识请求头的结束
							headerEnd=true;
							if(headersMap.get(CONTENT_LENGTH)!=null){
								contentLength=Integer.parseInt(headersMap.get(CONTENT_LENGTH));
							}
						}
					}
				}
			}
			if(contentLength!=0){
				if(entity==null){//请求头的最后一个LF，不放入实体部分
					header.append(c);
					entity= new ArrayList<>(contentLength);
				}else{
					entity.add(byt);
//					entity.append(c);
				}
			}else{
				header.append(c);
			}
		}
		if(headerEnd){//如果请求头读取完毕
			while(contentLength>0&&contentLength>entity.size()){
				bufferSize=inputStream.read(buffer);
				for(int i=0; i<bufferSize; i++){
					entity.add(buffer[i]);
				}
//				entity.append(new String(buffer,0,bufferSize, DEFAULT_CHARSET));
			}
		}else{
			parseRequest();
		}
	}
	
	private void putHeader(String header){
		int index=header.indexOf(":");
		String key=header.substring(0,index);
		String value=header.charAt(index+1)==' '?header.substring(index+2):header.substring(index+1);
		headersMap.put(key, value);
	}
	
	public String getHeader() {
		return header.toString();
	}
	
	public String getEntity() throws UnsupportedEncodingException {
		if(entity==null){
			return null;
		}

		byte[] buf = new byte[entity.size()];

		for(int i=0; i<entity.size(); i++){
			buf[i] = entity.get(i);
		}

		return new String(buf, DEFAULT_CHARSET);
	}
	
	public String getUri(){
		return requestLine.split(" ")[1];
	}
}

interface HttpMessage {

	/** 回车 */
	char CR='\r';
	/** 换行 */
	char LF='\n';
	/** 回车换行 */
	String CRLF="\r\n";

	/** 请求body长度 头名称 */
	String CONTENT_LENGTH="Content-Length";

	/** 默认编码 */
	String DEFAULT_CHARSET = "utf-8";

}

```

</details>


<details>
<summary>解析form-data 数据</summary>

```java
private static void parseHttpBody() {
    String boundary = "----WebKitFormBoundaryaK4dtoQlBHMyMqDk";
    String content = Futil.readStr("d:/body.txt");

    String regex = boundary+"(?<part>.+?)--";
    Pattern compile = Pattern.compile(regex, Pattern.MULTILINE | Pattern.DOTALL | Pattern.CASE_INSENSITIVE);
    Matcher matcher = compile.matcher(content);
    while (matcher.find()) {
        String bodyPart = matcher.group("part").trim();
        handleBodyPart(bodyPart);
    }
}

private static void handleBodyPart(String bodyPart) {
    String lineSeperator = "\r\n";
    int firstLineEndIndex = bodyPart.indexOf(lineSeperator);
    String firstLine = bodyPart.substring(0, firstLineEndIndex);
    boolean isFilePart = firstLine.contains("filename");
    System.out.println(firstLine);
    System.out.println(isFilePart);

    String regex = "name=\"(?<name>[^\"]+)\"";
    Pattern compile = Pattern.compile(regex);
    Matcher matcher = compile.matcher(firstLine);

    if(matcher.find()){
        String name = matcher.group("name");
        System.out.println(name);
    }
    System.out.println("value:");
    int valStartIndex = 0;
    if(isFilePart){
        int nameEndIndex = bodyPart.indexOf(lineSeperator, firstLineEndIndex+lineSeperator.length());
        valStartIndex = nameEndIndex+lineSeperator.length()*2;
    }else{
        valStartIndex = firstLineEndIndex+lineSeperator.length()*2;
    }
    String value = bodyPart.substring(valStartIndex);
    System.out.println(value);
    System.out.println("||||||||||||||||||||||||||||||||||||||");

    System.out.println("-=--=--=--=--=--=--=--=--=--=--=-");
}
```

</details>
