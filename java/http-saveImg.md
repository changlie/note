```java
//okhttpclient
private static void saveImg2(String url) throws IOException {
    OkHttpClient client = new OkHttpClient();
    Request request = new Request.Builder().url(url).get().build();
    Response response = client.newCall(request).execute();

    InputStream inputStream = response.body().byteStream();
    FileUtils.copyInputStreamToFile(inputStream, new File("d:/gogo1.jpg"));
}

//httpclient
private static void saveImg1(String url) throws IOException {
    HttpClient client = HttpClientBuilder.create().build();
    HttpGet httpGet = new HttpGet(url);
    HttpResponse response = client.execute(httpGet);

    InputStream inputStream = response.getEntity().getContent();
    FileUtils.copyInputStreamToFile(inputStream, new File("d:/abc11.jpg"));
}
```
