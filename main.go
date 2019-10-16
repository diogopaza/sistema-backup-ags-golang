package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	copyFile("diretorio")
}

func copyFile(dst string) error {
	files, err := ioutil.ReadDir("/home/zaptec/Documentos/")
	if err != nil {
		fmt.Println("Erro na leitura do diretorio ", err)
		return nil
	}
	for _, file := range files {
		fmt.Println("Arquivo no diretorio ", file.Name())
	}
	return nil
}
