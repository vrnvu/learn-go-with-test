package di

import (
	"bytes"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "world")

	got := buffer.String()
	want := "Hello, world"

	assert.Equal(t, want, got)
}

func TestGreetHttp(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		Greet(w, "world")
	}
	http.ListenAndServe(":8000", http.HandlerFunc(handler))
}

func TestGreetStdout(t *testing.T) {
	Greet(os.Stdout, "world")
}
