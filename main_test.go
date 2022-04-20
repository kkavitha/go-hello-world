package main

import (
    "testing"

)

func TestHelloWorld(t *testing.T) {
    want := "hello world"
    msg := returnString()
    if want != msg {
        t.Fatalf(`Got %q, want match for %#q, nil`, msg, want)
    }
}
