package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://golang.org/",
		"http://google.com",
		"http://www.amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}
	for l := range c { 
		// A função abaixo funciona como uma callback
		// os parametros na goFunc e no () abaixo são importantes para manejar a memória e evitar
		// que os valores apontem para 
		go func(link string) { 
			time.Sleep(3 * time.Second)
			checkLink(link, c)
		}(l)
	}
	// for i := 0; i < len(links); i++ {
	// 	// fmt.Println(<- c) // le a mensagem enviada
	// 	go checkLink(<- c)
	// }
}


func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil { 
		fmt.Println(link, "might be down")
		c <- link
		return
	}
	fmt.Println(link, "is up")
	c <- link
}