### HttpUrlConnection 其实java内置的一个http协议的客户端实现。

实例获取
```java
URL url = new URL(path);
HttpURLConnection openConnection = (HttpURLConnection) url.openConnection();
```

### 相关方法
1. 设置请求方法
```java
public void setRequestMethod(String method)
```

2. 设置请求头 
```java
public void setRequestProperty(String key, String value)
```

3. 获取请求头
```java
public String getRequestProperty(String key)
public Map<String,List<String>> getRequestProperties()
```

4. 获取响应头
```java
public String getHeaderField(String name)
public Map<String,List<String>> getHeaderFields()
public int getHeaderFieldInt(String name, int Default)
public long getHeaderFieldLong(String name, long Default)
```

5. 获取响应体 http response body
```java
public InputStream getInputStream()
```

6. 获取响应状态码
```java
public int getResponseCode()
```

7. 获取响应头首行的状态码说明信息。如，HTTP/1.0 200 OK中OK; "HTTP/1.0 404 Not Found"中"Not Found"
```java
public String getResponseMessage()
```


