package mian

import (
	"io"
	"fmt"
	"bytes"
	"net/http"
	"net/textproto"
	"mime/multipart"
)

/*
因multipart.CreateFormFile函数写死了file的Content-Type为application/octet-stream,本demo重写此函数以解决服务端不接受octet-stream类型的file的问题
 */

func main() {
	//获取file可通过接口GetFile,也可通过os.Open.
	//file, fileHeader, err := c.GetFile("file")
	//fh, err := os.Open("xx.go")

	b := bytes.NewBuffer(nil)
	n, err := io.Copy(b, file)
	if err != nil {
		fmt.Println(err)
	}

	fileName := ""

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	fileWriter, err := CreateFormFile("body", fileName, "image/jpeg", bodyWriter)
	if err != nil {
		fmt.Println(err)
	}
	n2, err := fileWriter.Write(b.Bytes())
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Boundary()
	bodyWriter.Close()
	resp, err := http.Post(url, contentType, bodyBuf)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	//do something with resp
}

func CreateFormFile(fieldname, filename, contentType string, w *multipart.Writer) (io.Writer, error) {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
			escapeQuotes(fieldname), escapeQuotes(filename)))
	h.Set("Content-Type", contentType)
	return w.CreatePart(h)
}

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")
