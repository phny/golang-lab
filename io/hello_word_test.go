package iotest

import (
	"fmt"
	"testing"
)

func Hello() {
	fmt.Println("Hello world!")
}

func TestHello(t *testing.T) {
	Hello()
}
