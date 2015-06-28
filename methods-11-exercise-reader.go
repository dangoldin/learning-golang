package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (mr MyReader) Read(a[] byte) (int, error) {
     for i, _ := range a {
     	 a[i] = byte('A')
	 }
	 return len(a), nil
}

func main() {
     reader.Validate(MyReader{})
}
