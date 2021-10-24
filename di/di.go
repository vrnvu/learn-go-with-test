package di

import (
	"fmt"
	"io"
)

func Greet(writer io.Writer, msg string) {
	fmt.Fprintf(writer, "Hello, %s", msg)
}
