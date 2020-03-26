<details>
	<summary>公共类</summary>

```java
/**
 * 用于临时记录 multipart/form-data 类型的httpbody的  bodypart信息
 */
public class NameValue {
	public String name;
	public Object value;

	public String fileName;

	boolean isFinish(){
		return name != null && value != null;
	}
}


/**
 * 用于记录上传文件信息。
 */
public class FileInfo{
	private String name;
	private byte[] data;

	public FileInfo() {
	}

	public FileInfo(String name, byte[] data) {
		this.name = name;
		this.data = data;
	}

	public String getName() {
		return name;
	}

	public void setName(String name) {
		this.name = name;
	}

	public byte[] getData() {
		return data;
	}

	public void setData(byte[] data) {
		this.data = data;
	}

	public void saveFile(String path){
		File f = new File(path);
		try (FileOutputStream os = new FileOutputStream(f)) {
			os.write(data, 0, data.length);
			os.flush();
		}catch (Exception e){
			e.printStackTrace();
		}
	}
}


/**
 * 解决数据用的常量定义集合
 */
public interface HttpMessage {

	/** 回车 */
	char CR='\r';
	/** 换行 */
	char LF='\n';
	/** 回车换行 */
	String CRLF="\r\n";

	/** 请求body长度(头名称) */
	String CONTENT_LENGTH="Content-Length";

	/** 请求body参数类型(头名称) */
	String CONTENT_TYPE = "Content-Type";

	/** 默认编码 */
	String DEFAULT_CHARSET = "utf-8";
//	String DEFAULT_CHARSET = "ISO-8859-1";

	/** utf8编码 */
	String CHARSET_UTF8 = "utf-8";

	/** form表单带文件请求类型 */
	String MULTIPART_FORM_DATA = "multipart/form-data";


	default String getValByUTF8(String origin){
		try {
			byte[] bytes = origin.getBytes(DEFAULT_CHARSET);
			return new String(bytes, CHARSET_UTF8);
		} catch (Exception e) {
			e.printStackTrace();
		}
		return origin;
	}

}
```

</details>


<details>
<summary>http请求数据解析类</summary>

```java
/**
 * http请求数据解析类
 */
public class HttpRequest implements HttpMessage{
	private InputStream inputStream = null;

	/** 请求头的第一行，包含请求方法，请求url, http协议 */
	private String requestLine = null;

	private int contentLength = 0; // 请求体的长度
	private boolean headerEnd = false; // 标记请求头的位置。

	/** 请求头键值对 */
	private Map<String,String> headersMap=new HashMap<>();
	/** 请求头缓冲区 */
	private StringBuffer header=new StringBuffer();
	/** 请求体缓冲区 */
	private List<Byte> entity = null;

	private Map<String, Object> httpArgs = new HashMap<>();
	
	public HttpRequest(InputStream inputStream) {
		this.inputStream=inputStream;
	}
	
	public void parseRequest() throws IOException{
		byte[] buffer=new byte[2048]; //临时缓冲区。

		// 获取请求头数据。
		while(!headerEnd){
			// 逐个逐个字节读取，以获取请求头数据。
			int bufferSize=inputStream.read(buffer);
			for(int index=0;index<bufferSize;index++){
				byte byt = buffer[index];
				char c=(char) buffer[index];
				int length=header.length();

				// 以回车符，换行符分析，归类请求数据。
				// 并从请求头缓冲区解析出每个请求头name,value保存至一个map.
				// 标记出请求头结束位置。
				if(c==LF && header.charAt(length-1)==CR){
					doParseRequest(length);
				}

				// 收集请求头数据至请求头缓冲区，并处理请求头，请求体的边界的数据
				sortoutByte(c, byt);
			}
		}

		//如果请求头读取完毕，并且存在请求体，则继续读取请求体数据
		while(contentLength>0 && contentLength>entity.size()){
			int bufferSize=inputStream.read(buffer);

			for(int i=0; i<bufferSize; i++){
				entity.add(buffer[i]);
			}
		}
	}

	/**
	 * 以回车符，换行符分析，归类请求数据。<br>
	 * 并从请求头缓冲区解析出每个请求头name,value保存至一个map.<br>
	 * 标记出请求头结束位置。
 	 * @param length 当前请求头缓冲区的长度
	 */
	private void doParseRequest(int length) {
		if(requestLine==null){//第一个CRLF是请求行
			requestLine=header.toString();
			return;
		}

		// 从请求头缓冲区解析出每个请求头name,value保存至一个map
		int crlf = header.lastIndexOf(CRLF);
		String headerLine=header.substring(crlf+2,length-1);
		if(!"".equals(headerLine)){//请求头
			putHeader(headerLine);
			return;
		}

		//请求数据中，第一次出现以CRLF结尾的空行, 意味着请求头已结束
		headerEnd=true;
		if(headersMap.get(CONTENT_LENGTH)!=null){
			contentLength=Integer.parseInt(headersMap.get(CONTENT_LENGTH));
		}
	}

	/**
	 * 收集请求头数据至请求头缓冲区，并处理请求头，请求体的边界的数据
	 * @param c 字符 由字节强转得来，为ASCII时不会出错。
	 * @param byt 请求数据的单字节数据
	 */
	private void sortoutByte(char c, byte byt) {
		// 收集请求头数据。
		if(!headerEnd){
			header.append(c);
			return;
		}

		// 处理请求头，请求体的边界
		if(entity==null){//请求头的最后一个LF，不放入实体部分
			header.append(c);
			entity= new ArrayList<>(contentLength);
		}else{
			// 收集请求体的数据
			entity.add(byt);
		}
	}

	/**
	 * 解析请求参数。
	 * @throws UnsupportedEncodingException
	 */
	public void parseParams() throws UnsupportedEncodingException {
		String contentType = getHeader(CONTENT_TYPE);
		if(contentType.startsWith(MULTIPART_FORM_DATA)){
			String[] split = contentType.split(";");
			String boundary = split[1].split("=")[1];
//			String content = getEntity();
//			FormDataParser p = new FormDataParser(content, boundary);

			byte[] r = getRawEntity();
			FormDataBytesParser p = new FormDataBytesParser(r, boundary);
			Map<String, Object> params = p.doParse();
			httpArgs.putAll(params);
		}
	}

	public String getParamString(String name){
		Object val = httpArgs.get(name);
		if(val==null) return null;
		return val.toString();
	}

	public FileInfo getParamFile(String name){
		Object val = httpArgs.get(name);
		if(val==null) return null;
		return (FileInfo) val;
	}

	public Map<String, Object> getHttpArgs(){
		return httpArgs;
	}
	
	private void putHeader(String header){
		int index=header.indexOf(":");
		String key=header.substring(0,index);
		String value=header.charAt(index+1)==' '?header.substring(index+2):header.substring(index+1);
		headersMap.put(key, value);
	}

	public String getHeader(String name){
		String val = headersMap.get(name);
		if(val == null){
			return null;
		}
		return val.trim();
	}
	
	public String getHeaders() {
		return header.toString();
	}
	
	public String getEntity() throws UnsupportedEncodingException {
		if(getRawEntity()==null){
			return null;
		}

		return new String(getRawEntity(), DEFAULT_CHARSET);
	}

	public byte[] getRawEntity(){
		if(entity==null){
			return null;
		}
		int len = entity.size();

		byte[] buf = new byte[len];

		for(int i = 0; i< len; i++){
			buf[i] = entity.get(i);
		}
		return buf;
	}
	
	public String getUri(){
		return requestLine.split(" ")[1];
	}
}
```

</details>


<details>
<summary>multipart/form-data 类型的httpbody的解析类</summary>

```java
/**
 *  multipart/form-data 类型的httpbody的解析类
 */
public class FormDataBytesParser implements HttpMessage {
	/** 包含整个http body 的 字节数组 */
	private byte[] bytes;
	/** 请求头的约定好的键值对分隔符。 */
	String boundary;

	/** 键值对之间的分隔符 */
	String nameValueSeperator;
	/** 键值对起始token */
	String nameValueEnd;
	/** 键值对结束token */
	String nameValueStart;

	/** 用于记录bodyPart的起始索引 */
	List<Integer> nameValueStartIndexStack = new ArrayList<>();
	/** 用于记录键值对的value的起始索引 */
	List<Integer> valueStartIndexStack = new ArrayList<>();

	/** 相应起始索引获取失败时，返回的值 */
	int NONE = -1;

	/** 本解析器，解析后的结果集 */
	Map<String, Object> params = new HashMap<>();

	public FormDataBytesParser(byte[] bytes, String boundary) {
		this.bytes = bytes;
		this.boundary = boundary;
		this.nameValueStart = boundary+CRLF;
		this.nameValueSeperator = CRLF+CRLF;
		this.nameValueEnd = CRLF+"--";
	}

	/**
	 * 开始执行对httpbody 的解析任务。
	 * @return 一个map
	 * @throws UnsupportedEncodingException
	 */
	public Map<String, Object> doParse() throws UnsupportedEncodingException {
		int len = bytes.length;


		NameValue nv = null;
		for(int i=0; i<len; i++){
			char c = (char) bytes[i];
			if(i<nameValueStart.length()){
				continue;
			}
			if(LF == c){
				if(tailMatchNameValueStart(i)){
					// 标记键值对的起始索引
					nameValueStartIndexStack.add(i+1);
					nv = new NameValue();
				}else if(tailMatchCRLF(i) && lastStartIndex() != NONE && isNew(nv)){
					// 获取键值对的键信息。
					int lastStartIndex = lastStartIndex();
					int lineLength = i-1 - lastStartIndex;
					String nameValueFirstLine = new String(bytes, lastStartIndex, lineLength, CHARSET_UTF8);

					nv.name = getName(nameValueFirstLine);
					if(nameValueFirstLine.contains("filename=\"")){
						nv.fileName = getFileName(nameValueFirstLine);
					}
				}else if(tailMatchSeperator(i) && hasName(nv)){
					// 标记键值对的值的起始索引
					valueStartIndexStack.add(i+1);
				}
			}else if('-' == c && tailMatchEnd(i) && lastValueStartIndex() != NONE){
				// 获取键值对的值信息，把解析出来的键值对信息保存到结果集
				// 并且置空临时对象NameValue nv
				int valueStartIndex = lastValueStartIndex();
				int valueEndIndex = i - (nameValueEnd.length()-1);
				byte[] valBytes = getBytes(this.bytes, valueStartIndex, valueEndIndex);
				if(nv.fileName != null){
					FileInfo f = new FileInfo(nv.fileName, valBytes);
					params.put(nv.name, f);
					continue;
				}
				params.put(nv.name, new String(valBytes, CHARSET_UTF8));
				nv = null;
			}
		}
		return params;
	}

	/**
	 * 把换行符，回车符显示出来
	 * @param raw
	 * @return
	 */
	String convert(String raw){
		StringBuilder res = new StringBuilder();
		int length = raw.length();
		for(int i=0; i<length; i++){
			char c = raw.charAt(i);
			if(CR == c){
				res.append("\\r");
			}else if(LF == c){
				res.append("\\n");
			}else{
				res.append(c);
			}
		}
		return res.toString();
	}

	/**
	 * 根据给定的起始索引与结束索引，获取一个子 字节数组
	 * @param bytes
	 * @param startIndex
	 * @param endIndex
	 * @return
	 */
	private byte[] getBytes(byte[] bytes, int startIndex, int endIndex) {
		int len = endIndex-startIndex;
		byte[] res = new byte[len];
		for(int i=0, j=startIndex; i<len; i++, j++){
			res[i] = bytes[j];
		}
		return res;
	}

	/**
	 * NameValue 临时对象的成员变量name是否已被初始化
	 * @param nv
	 * @return
	 */
	private boolean hasName(NameValue nv) {
		return nv != null && nv.name != null;
	}

	/**
	 * NameValue 临时对象是否为刚初始化的状态
	 * @param nv
	 * @return
	 */
	private boolean isNew(NameValue nv) {
		return nv != null && nv.name == null;
	}


	/**
	 * 获取获取键值对的键名称。
	 * @param content
	 * @return
	 */
	private String getName(String content) {
		String regex = "name=\"(?<name>[^\"]+)\"";
		Pattern compile = Pattern.compile(regex);
		Matcher matcher = compile.matcher(content);
		if(matcher.find()){
			return matcher.group("name").trim();
		}else {
			return null;
		}
	}

	/**
	 * 获取上传文件名称。
	 * @param line
	 * @return
	 */
	private String getFileName(String line){
		String regex = "filename=\"(?<name>[^\"]+)\"";
		Pattern compile = Pattern.compile(regex);
		Matcher matcher = compile.matcher(line);
		if(matcher.find()){
			return matcher.group("name");
		}else {
			return null;
		}
	}

	/**
	 * 获取当前键值对的值的起始索引
	 * @return
	 */
	private int lastValueStartIndex() {
		int len = valueStartIndexStack.size();
		if(len<1 || len != nameValueStartIndexStack.size()){
			return NONE;
		}
		return valueStartIndexStack.get(len-1);
	}

	/**
	 * 获取当前键值对的起始索引
	 * @return
	 */
	private int lastStartIndex() {
		int len = nameValueStartIndexStack.size();
		if(len>0){
			return nameValueStartIndexStack.get(len-1);
		}

		return NONE;
	}

	/**
	 * 当前位置是否为键值对的结束位置
	 * @param i
	 * @return
	 */
	private boolean tailMatchEnd(int i) {
		return tailMatch(i, nameValueEnd);
	}

	/**
	 * 当前位置是否为键值对的键与值的分隔符位置
	 * @param i
	 * @return
	 */
	private boolean tailMatchSeperator(int i) {
		return tailMatch(i, nameValueSeperator);
	}

	/**
	 * 当前位置是否匹配一个回车换行字符串
	 * @param endIndex
	 * @return
	 */
	private boolean tailMatchCRLF(int endIndex) {
		return tailMatch(endIndex, CRLF);
	}

	/**
	 * 当前位置是否为键值对的起始位置。
	 * @param endIndex
	 * @return
	 */
	private boolean tailMatchNameValueStart(int endIndex) {
		return tailMatch(endIndex, nameValueStart);
	}

	/**
	 * 从当前位置往后推(token.length()-1)个位置是否匹配token
	 * @param endIndex 当前位置
	 * @param token 标记字符串
	 * @return
	 */
	private boolean tailMatch(int endIndex, String token) {
		int startIndex = endIndex-(token.length()-1);
		for(int i=startIndex, j=0; i<= endIndex; i++, j++){
			char c = (char) bytes[i];
			if(c != token.charAt(j)){
				return false;
			}
		}
		return true;
	}
}
```

</details>
