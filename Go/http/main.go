package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

//interfaces send us down the right path but not neccesarily the right code

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	//bs := make([]byte, 99999) //necessary to initialise a byte slice w large size to store stuff from read function
	//resp.Body.Read(bs)
	//fmt.Println(string(bs)) //reader interface
	//or
	lw := logWriter{}
	io.Copy(lw, resp.Body) //writer interface //writes to terminal, http file etc.etc.
	//func Copy(dst Writer, src Reader) (written int64, err error)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}

//logWriter now implements Write function
//this func instead of actual write// actually doesn't write anything
