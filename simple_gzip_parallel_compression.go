package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"sync"
	"strings"
)

func main() {
	var wg sync.WaitGroup //Wait groups don't need to be initialised
	var i int = -1        //Because we want to reference i outside the for loop, we declare it here
	var file string
	for i, file = range os.Args[1:] {
		wg.Add(1)
		go func(filename string) {
			compress(filename)
			wg.Done()
		}(file)
	}
	wg.Wait()
	fmt.Printf("Compressed %d files\n", i+1)
}

func compress(filename string) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(filename + ".gz")
	if err != nil {
		return err
	}
	defer out.Close()

	gzout := gzip.NewWriter(out)
	_, err = io.Copy(gzout, in)
	gzout.Close()

	return err
}

func deCompress(filename string) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(strings.TrimSuffix(filename, ".gz"))
	if err != nil {
		return err
	}
	defer out.Close()

	gzout, err := gzip.NewReader(out)
	if err != nil {
		return err
	}
	_, err = io.Copy(gzout, in)
	gzout.Close()

	return err
}
