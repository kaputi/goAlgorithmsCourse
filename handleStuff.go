package main

import (
	"testing"
)

func ExpectGet(t *testing.T, expect any, get any) {
	if expect != get {
		t.Errorf("expected: %v, got: %v", expect, get)
	}
}

func HandleErr(t *testing.T, err error) {
	if err != nil {
		t.Errorf("error %v", err)
	}
}
