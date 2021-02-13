package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/shyupc/go-study/gointro/pipeline"
)

const (
	filename = "small.in"
	n        = 64
)

func main() {
	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	p := pipeline.RandomSource(n)
	writer := bufio.NewWriter(file)
	pipeline.WriteSink(writer, p)
	writer.Flush()

	file, err = os.Open(filename)
	p = pipeline.ReaderSource(bufio.NewReader(file), -1)

	count := 0
	for v := range p {
		fmt.Println(v)
		count++
		if count >= 100 {
			break
		}
	}

}

func mergeDemo() {
	p := pipeline.Merge(
		pipeline.InMemSort(
			pipeline.ArraySource(3, 2, 6, 7, 4)),
		pipeline.InMemSort(
			pipeline.ArraySource(7, 4, 0, 3, 2, 13, 8)))

	for v := range p {
		fmt.Println(v)
	}
}
