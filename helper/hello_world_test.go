package helper

import "testing"

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
