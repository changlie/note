```golang
bodyBuffer := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuffer)

	fileWriter, _ := bodyWriter.CreateFormFile("files", "filexxx.txt")
	file, _ := os.Open("file.txt")
	defer file.Close()
	io.Copy(fileWriter, file)

	fileWriter2, _ := bodyWriter.CreateFormFile("files", "code.txt")
	file2, _ := os.Open("test.go")
	defer file2.Close()
	io.Copy(fileWriter2, file2)

	// other form data
	extraParams := map[string]string{
		"title":       "My Document",
		"author":      "Matt Aimonetti",
		"description": "A document with all the Go programming language secrets",
	}
	for key, value := range extraParams {
		_ = bodyWriter.WriteField(key, value)
	}


	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, _ := http.Post("http://localhost:12345/hello?id=998", contentType, bodyBuffer)
	defer resp.Body.Close()

	resp_body, _ := ioutil.ReadAll(resp.Body)

	log.Println(resp.Status)
	log.Println(string(resp_body))
```