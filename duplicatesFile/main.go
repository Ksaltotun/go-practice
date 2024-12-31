package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	fileUniq := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "standardIO", fileUniq)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				continue
			}
			countLines(f, counts, arg, fileUniq)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%s: %d\n", line, n)
		}
	}
	fmt.Printf("Файлы, с повторяющимися строками:\n")
	for name, n := range fileUniq {

		if n > 1 {
			fmt.Printf("%s: %d\n", name, n)
		}
	}
}

func countLines(f *os.File, counts map[string]int, name string, fileUniq map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		_, ok := counts[input.Text()]
		if ok {
			fileUniq[name]++
		}
	}
}
