// https://www.alura.com.br/artigos/criando-uma-simples-aplicacao-web-com-go
package main

import "net/http"

func handler(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, "./static/index.html")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
