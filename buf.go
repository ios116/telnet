package main

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
)

type Writer int
func (*Writer) Write(p []byte) (n int, err error) {
	fmt.Printf("Write: %q\n", p)
	return 0, errors.New("boom!")
}
func main() {
	reader := strings.NewReader("hello ")
	bw := bufio.NewReader(reader)
	fmt.Println(bw)

}
