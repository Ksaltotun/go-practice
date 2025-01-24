package main

import "fmt"

func main() {
	for i, r := range "аоом klfkls SSSSЫ ЫЫЫЫЫ" {
		fmt.Printf("\a%d\t%q\t%d\n", i, r, r)
	}
}
