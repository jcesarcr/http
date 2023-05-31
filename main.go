package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error " + err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.Status)
	fmt.Println(resp.StatusCode)
	//fmt.Println(resp.Header)
	fmt.Println(resp.Body)

	//bs := make([]byte, 99999)
	//resp.Body.Read(bs)
	//fmt.Println(string(bs))

	lw := logWriter{}

	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes ", len(bs))
	return len(bs), nil
}
