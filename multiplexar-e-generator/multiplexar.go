package main

import (
	"fmt"
	"time"
)

// falar é uma função que simula o padrão generator. Basicamente a partir de um valor é gerado um chan e retornado
func falar(pessoa string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			c <- fmt.Sprintf("%s falando: %d", pessoa, i)
		}
	}()

	return c
}

// juntar é uma função que simula o padrão multiplexar. Basicamente a partir de dois channels é retornado o primeiro a ser inserido nos atributos de entradas.
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-entrada1:
				c <- s
			case s := <-entrada2:
				c <- s
			}
		}
	}()

	return c
}

func main() {
	c := juntar(falar("João"), falar("Maria"))

	fmt.Println(<-c + " | " + <-c)
	fmt.Println(<-c + " | " + <-c)
	fmt.Println(<-c + " | " + <-c)
}
