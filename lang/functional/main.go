package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/shyupc/go-study/lang/functional/fib"
)

type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func printFileContents(r io.Reader) {
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fib.Fibonacci()
	printFileContents(intGen(f))
}
