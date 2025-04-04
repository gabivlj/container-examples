package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"

	"github.com/klauspost/compress/zstd"
)

var logs bytes.Buffer

func main() {
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write(logs.Bytes())
		})

		http.ListenAndServe(":8002", nil)
	}()

	// TODO: simplify once we have half-open sockets fixed

	controller, err := net.Listen("tcp", "0.0.0.0:8003")
	if err != nil {
		panic(err)
	}

	ld, err := net.Listen("tcp", "0.0.0.0:8001")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ld.Accept()
		if err != nil {
			panic(err)
		}

		output, err := controller.Accept()
		if err != nil {
			panic(err)
		}

		writer, err := zstd.NewWriter(output, zstd.WithEncoderConcurrency(1))
		if err != nil {
			fmt.Fprintln(&logs, "error new writer:", err)
			conn.Close()
			return
		}

		defer output.Close()
		n, err := io.Copy(writer, conn)
		if err != nil {
			fmt.Fprintln(&logs, "error new copy:", err)
			conn.Close()
			return
		}

		fmt.Fprintln(&logs, "Written", n)
		writer.Close()
		conn.Close()
		output.Close()
	}
}
