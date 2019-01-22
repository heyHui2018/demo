package main

import (
	"io"
	"net/http"
)

func ServeHTTP222(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world2222222!"))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func helloHandler1(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world111!\n")
}

//func main() {
//	http.HandleFunc("/test2", ServeHTTP222)
//	http.ListenAndServe(":8088", nil)
//}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello/", helloHandler)
	mux.HandleFunc("/hello/hello/", helloHandler1)
	http.ListenAndServe(":8088", mux)
}
//关于ServeMux：
//1、URL 分为两种，末尾是 /，表示一个子树，后面可以跟其他子路径； 末尾不是 /，表示一个叶子，固定的路径
//2、以 / 结尾的 URL 可以匹配它的任何子路径，比如 /images 会匹配 /images/cute-cat.jpg
//3、采用最长匹配原则，如果有多个匹配，一定采用匹配路径最长的那个进行处理
//4、如果没有找到任何匹配项，会返回 404 错误
//5、会识别和处理 . 和 ..，正确转换成对应的 URL 地址