package domain

import (
	"testing"
)

func SayHello_PrintsName(t *testing.T) {
	want := "Hello"
	msg := SayHelloWorld()
	if want != msg {
		t.Fatalf("Expected %s but got %s", want, msg)
	}
}
