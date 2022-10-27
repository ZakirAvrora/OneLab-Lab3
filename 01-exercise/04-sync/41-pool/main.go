package main

import (
	"bytes"
	"io"
	"os"
	"sync"
	"time"
)

//TODO: create pool of bytes.Buffers which can be reused.

var pool = sync.Pool{
	New: func() any {
		return &bytes.Buffer{}
	},
}

func log(w io.Writer, val string) {
	b := pool.Get().(*bytes.Buffer)
	defer pool.Put(b)

	b.WriteString(time.Now().Format("15:04:05"))
	b.WriteString(" : ")
	b.WriteString(val)
	b.WriteString("\n")

	w.Write(b.Bytes())
	b.Reset()
}

func main() {
	log(os.Stdout, "debug-string1")
	log(os.Stdout, "debug-string2")
}
