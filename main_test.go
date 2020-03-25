package main

import "testing"

func TestHelloWorld(t *testing.T) {
	if helloworld() != "Hellotest World!!" {
		t.Fatal("Test fail")
	}
}
