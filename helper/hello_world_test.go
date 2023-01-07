package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Bad")
	if result != "Hello Bad" {
		t.Errorf("HelloWorld() = %s; want Hello Bad", result)
	}
}

func TestBad(t *testing.T) {
	result := HelloWorld("Bad")
	if result != "Hello Bad" {
		t.Errorf("HelloWorld() = %s; want Hello Bad", result)
	}
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Eko")
	assert.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
	fmt.Println("TestHelloWorld with Assert Done")
}
