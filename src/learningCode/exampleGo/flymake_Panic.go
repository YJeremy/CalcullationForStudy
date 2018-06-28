package main

import "os"

func main() {
	panics("a problem")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
