
package main

import (
    "fmt"
    "time"
    "math/rand"
)

func find(arquivos chan string, diretorio string) { 

    for arq in diretorio{
        if(arq isFile){
            if(arq.byte() % 2 == 0){
                arquivos <- diretorio
            }
        }else{
            find(arquivos, arq.children)
        }

    }
	joi

}
func write(arquivos chan string){
    for i in arquivos {
        if(i.byte % 2 == 0){
            print(i)
        }
    }
}
func main() {
    var root string 
    fmt.Scan(&root)

    arquivos := make(chan string)

    go find(arquivos, root)
    go write(arquivos)
    <- arquivos

}

func find(root string, filesChan chan string, cont int) {
	if cont == 0 {
		defer close(filesChan)
	}

	files, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		addr := root + "/" + f.Name()
		if f.IsDir() {
			find(addr, filesChan, cont+1)
		} else {
			filesChan <- addr
		}
	}
}