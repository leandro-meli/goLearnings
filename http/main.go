package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logWriter{}
	io.Copy(lw, resp.Body)
	// bs := make([]byte, 99999) // A função read vai ler até preencher todo byteslice, por isso inicializar com muito espaço
	fmt.Println(resp.StatusCode)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))
}

func (logWriter) Write(bs []byte) (int, error) { // receiver, parameters, return
	fmt.Println("We just wrote", len(bs), " bytes")

	return len(bs), nil //A função write retorna a quantidade de bytes processadas
}
