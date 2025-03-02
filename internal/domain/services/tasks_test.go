package services

import (
	"testing"
)

func SayHello_PrintsName(t *testing.T) {
	want := "Hello"
	msg := sayHelloWorld()
	if want != msg {
		t.Fatalf("Expected %s but got %s", want, msg)
	}
}
