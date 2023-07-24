package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {
}

func main() {

	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// fmt.Println(resp)

	// slice of type byte
	// make(Type, size), built-in function used to create and initialize certain composite types, specifically slices, maps, and channels

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs) //response body will be taken and read into byte slice
	// fmt.Println(string(bs))

	lw := logWriter{}

	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {

	// return 1, nil
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))

	return len(bs), nil
}
