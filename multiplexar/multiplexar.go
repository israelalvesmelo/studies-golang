package main

import (
	"fmt"

	"github.com/israelalvesmelo/estudos-golang/generators/html"
)

// Recebe um chan de leitura e adicionar o valor no chan destino que pode ser alterado
func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

// Multiplexar - misturar (mensagens) num canal
func multiplexar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := multiplexar(
		html.Titulo("https://www.google.com", "https://www.youtube.com"),
		html.Titulo("https://www.amazon.com", "https://www.mercadolivre.com"),
	)

	fmt.Println(<-c, " | ", <-c)
	fmt.Println(<-c, " | ", <-c)
}
