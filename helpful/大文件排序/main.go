package main

import (
	"bufio"
	"encoding/binary"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"sort"
	"strconv"
	"time"
)

//生成一个随机数文件，假设文件很大
//从文件中分块地读取数据到内存，进行各个结点的内部排序
//归并排序得到最终的排序结果写入文件中

// 数据源来自 ArraySource
func ArraySource(a ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, i := range a {
			out <- i
		}
		close(out)
	}()
	return out
}

// 内部排序
func InMemSort(in <-chan int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		a := []int{}
		for i := range in {
			a = append(a, i)
		}
		sort.Ints(a)
		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}

// 归并排序
func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		v1, ok1 := <-in1
		v2, ok2 := <-in2
		for ok1 || ok2 {
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1
				v1, ok1 = <-in1
			} else {
				out <- v2
				v2, ok2 = <-in2
			}
		}
		close(out)
	}()
	return out
}

// 随机数生成
func RandomSource(count int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Int()
		}
		close(out)
	}()
	return out
}

// 读写数据
func ReaderSource(reader io.Reader, chunkSize int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		buffer := make([]byte, 8)
		bytesRead := 0
		for {
			n, err := reader.Read(buffer)
			bytesRead += n
			if n > 0 {
				out <- int(binary.BigEndian.Uint64(buffer))
			}
			if err != nil || (chunkSize != -1 && bytesRead >= chunkSize) {
				break
			}
		}
		close(out)
	}()
	time.Sleep(time.Second)
	return out
}

func WriteSink(writer io.Writer, in <-chan int) {
	for i := range in {
		buffer := make([]byte, 8)
		binary.BigEndian.PutUint64(buffer, uint64(i))
		writer.Write(buffer)
	}
}

func main() {
	// 写文件
	const filename = "small.in"
	const n = 64
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	randomSource := RandomSource(n)
	writer := bufio.NewWriter(file)
	WriteSink(writer, randomSource)
	writer.Flush()

	// 读文件
	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	source := ReaderSource(bufio.NewReader(file), -1)
	count := 0
	for v := range source {
		log.Println(v)
		count++
		if count >= 100 {
			break
		}
	}
	file.Close()

	// 排序文件，并写新文件
	log.Println("-------------")
	filenameout := "small.out"
	pipeline := createPipeline(filename, 512, 4)
	writeToFile(pipeline, filenameout)
	printFile(filenameout)

	// 网络版
	log.Println("-------Net------")
	networkPipeline := createNetworkPipeline("small.in", 512, 4)
	writeToFile(networkPipeline, "small-net.out")
	printFile("small-net.out")

}

// 递归节点组，调用递归实现2路归并
func MergeN(inputs ...<-chan int) <-chan int {
	if len(inputs) == 1 {
		return inputs[0]
	}
	m := len(inputs) / 2
	return Merge(MergeN(inputs[:m]...), MergeN(inputs[m:]...))
}

func createPipeline(filename string, fileSize, chunkCount int) <-chan int {
	chunkSize := fileSize / chunkCount
	sortResults := []<-chan int{}
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	for i := 0; i < chunkCount; i++ {
		file.Seek(int64(i*chunkSize), 0)
		source := ReaderSource(bufio.NewReader(file), chunkSize)
		sortResults = append(sortResults, InMemSort(source))
	}
	return MergeN(sortResults...)
}

func writeToFile(p <-chan int, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	WriteSink(writer, p)
	writer.Flush()
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	source := ReaderSource(file, -1)
	count := 0
	for i := range source {
		log.Println(i)
		count++
		if count >= 100 {
			break
		}
	}
}

func NetworkSink(addr string, in <-chan int) {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	go func() {
		defer listen.Close()
		accept, err2 := listen.Accept()
		if err2 != nil {
			panic(err2)
		}
		defer accept.Close()
		writer := bufio.NewWriter(accept)
		WriteSink(writer, in)
		writer.Flush()
	}()
}

func NetworkSource(addr string) <-chan int {
	out := make(chan int)
	go func() {
		dial, err := net.Dial("tcp", addr)
		if err != nil {
			panic(err)
		}
		source := ReaderSource(bufio.NewReader(dial), -1)
		for i := range source {
			out <- i
		}
		close(out)
	}()
	return out
}

func createNetworkPipeline(filename string, fileSize, chunkCount int) <-chan int {
	chunkSize := fileSize/chunkCount
	sortAddr := []string{}
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	for i:=0;i<chunkCount;i++{
		file.Seek(int64(i*chunkSize),0)
		source:= ReaderSource(bufio.NewReader(file), chunkSize)
		addr:=":"+strconv.Itoa(7000+i)
		NetworkSink(addr, InMemSort(source))
		sortAddr = append(sortAddr, addr)
	}
	sortResults := []<-chan int{}
	for _, s := range sortAddr {
		sortResults = append(sortResults, NetworkSource(s))
	}
	return MergeN(sortResults...)
}
