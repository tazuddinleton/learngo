package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	fmt.Println(s)
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (n int, err error) {
	n, s := r.r.Read(b)
	bb := []byte{}

	for i := 0; i < len(b); i++ {
		var v = b[i]

		if (v >= 65 && v <= 77) || (v >= 97 && v <= 109) {
			v += 13
		} else if (v > 77 && v <= 90) || (v > 109 && v <= 122) {
			v -= 13
		}
		bb = append(bb, v)
	}

	copy(b, bb)
	return n, s
}

func runInfiniteStream() {
	r := MyReader{}

	b := make([]byte, 8)

	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v text: %v\n", n, err, b, string(b))
		if err == io.EOF {
			break
		}
	}
}

type MyReader struct{}

// Returns infinite stream of 'A'
func (r MyReader) Read(b []byte) (int, error) {
	n := copy(b, []byte{65})
	return n, nil
}

func readChunks() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)

	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v text: %v\n", n, err, b, string(b))
		if err == io.EOF {
			break
		}
	}
}
