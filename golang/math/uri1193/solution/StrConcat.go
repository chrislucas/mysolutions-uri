package main

import (
	"bytes"
	"fmt"
	"strings"
)
/*
https://stackoverflow.com/questions/1760757/how-to-efficiently-concatenate-strings-in-go
https://yourbasic.org/golang/build-append-concatenate-strings-efficiently/
https://golang.org/pkg/strings/#Builder
*/



func concatWithStringsBuilder() {
	stringBuilderII()
}

func stringBuilderI() {
	// strings. Buffer foi adicionado na versao 1.10 Golang
	var buffer strings.Builder
	for i, p  := range [] int {1,2,3} {
		if i == 0 {
			fmt.Fprintf(&buffer, "%d:%d", i, p)
		} else {
			fmt.Fprintf(&buffer, ", %d:%d", i, p)
		}
	}
	var rs = buffer.String()
	fmt.Println(rs)
}

func stringBuilderII() {
	var builder strings.Builder
	for i := 0; i <= 1000; i++ {
		if i == 0 {
			builder.WriteString(fmt.Sprintf("%d", i))
		} else {
			builder.WriteString(fmt.Sprintf(" %d", i))
		}
	}
	fmt.Println(builder.String())
}

func byteBufferI() {
	var buffer  bytes.Buffer
	for i := 0; i < 100; i++ {
		if i == 0 {
			buffer.WriteString(fmt.Sprintf("%d", i))

		} else {
			buffer.WriteString(fmt.Sprintf(" %d", i))
		}
	}
	fmt.Println(buffer.String())
}

func concatWithByteBuffer() {
	byteBufferI()
}

func main() {
	concatWithStringsBuilder()
}
