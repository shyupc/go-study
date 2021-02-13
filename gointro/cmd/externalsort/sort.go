package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/shyupc/go-study/gointro/pipeline"
)

func main() {
	p := createNetworkPipeline("large.in", 800000000, 4)
	writeToFile(p, "large.out")
	printFile("large.out")
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipeline.ReaderSource(file, -1)
	count := 0
	for v := range p {
		fmt.Println(v)
		count++
		if count >= 100 {
			break
		}
	}
}

func writeToFile(p <-chan int, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	pipeline.WriteSink(writer, p)
}

func createPipeline(filename string, fileSize, chuckCount int) <-chan int {
	pipeline.Init()
	chuckSize := fileSize / chuckCount
	var sortResult []<-chan int

	for i := 0; i < chuckCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}

		_, _ = file.Seek(int64(i*chuckSize), 0)

		source := pipeline.ReaderSource(bufio.NewReader(file), chuckSize)
		sortResult = append(sortResult, pipeline.InMemSort(source))
	}

	return pipeline.MergeN(sortResult...)
}

func createNetworkPipeline(filename string, fileSize, chuckCount int) <-chan int {
	pipeline.Init()
	chuckSize := fileSize / chuckCount
	var sortAddr []string

	for i := 0; i < chuckCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}

		_, _ = file.Seek(int64(i*chuckSize), 0)

		source := pipeline.ReaderSource(bufio.NewReader(file), chuckSize)
		addr := fmt.Sprintf(":%d", 7000+i)
		pipeline.NetworkSink(addr, pipeline.InMemSort(source))
		sortAddr = append(sortAddr, addr)
	}

	var sortResult []<-chan int
	for _, addr := range sortAddr {
		sortResult = append(sortResult, pipeline.NetworkSource(addr))
	}

	return pipeline.MergeN(sortResult...)
}
