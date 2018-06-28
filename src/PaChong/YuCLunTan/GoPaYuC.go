package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	r, _ := http.Get("http://bbs.fishc.com")
	defer r.Body.Close()
	fmt.Println("Header")
	for k, v := range r.Header {
		fmt.Printf("%s,%s\n", k, v)
	}

	body, _ := ioutil.ReadAll(r.Body)
	writeFile(body)
}

func writeFile(body []byte) {
	userFile := "fishC.html"
	fout, err := os.Create(userFile)
	defer fout.Close()
	if err != nil {
		fmt.Println(userFile, err)
		return

	}
	fout.Write(body)
}
