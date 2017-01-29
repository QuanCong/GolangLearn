package main

import (
	"net"
	"log"
	"io"
	"os"
)

func main() {


	connection, err := net.Dial("tcp", "localhost:9090")
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()
	mustCopy(os.Stdout, connection)

}

func mustCopy(dst io.Writer, src io.Reader) {

	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
