package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 rot13Reader) Read(b []byte) (int, error) {
	c := make([]byte, len(b))
	n, oErr := r13.r.Read(c)

	for i := range c {
		newVal := c[i] + 13
		if (c[i] <= 90 && newVal > 90) || newVal > 122 {
			newVal -= 26
		}
		b[i] = newVal
	}
	return n, oErr
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}

	io.Copy(os.Stdout, &r)
}
