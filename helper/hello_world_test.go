package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Eko")
	require.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
	fmt.Println("TestHelloWorld with Require Done")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Mac OS")
	}

	result := HelloWorld("Eko")
	require.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
}

func TestSubTest(t *testing.T) {
	t.Run("Eko", func(t *testing.T) {
		result := HelloWorld("Eko")
		require.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
	})
	t.Run("Kurniawan", func(t *testing.T) {
		result := HelloWorld("Kurniawan")
		require.Equal(t, "Hello Kurniawan", result, "Result must be 'Hello Kurniawan'")
	})
	t.Run("Khannedy", func(t *testing.T) {
		result := HelloWorld("Khannedy")
		require.Equal(t, "Hellosss Khannedy", result, "Result must be 'Hello Khannedy'")
	})
}
