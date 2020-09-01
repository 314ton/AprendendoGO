package main

import (
	"fmt"
)

func rtf() ( func() ) {
	return func() {
		fmt.Println("Isso é o retorno dos q não foram")
	}
}

func main() {
	def := rtf()
	def()
}