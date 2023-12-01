package main

import "fmt"

func main() {
	s := "Hello world!"
	fmt.Println(falar(s))
}

func falar(s string) string {
	return s + "Teste"
}
