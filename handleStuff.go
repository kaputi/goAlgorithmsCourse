package main

import (
	"testing"
)

func ExpectGetInt(t *testing.T, expect int, get int) {
	if expect != get {
		t.Errorf("expected: %v, got: %v", expect, get)
	}
}

func HandleErr(t *testing.T, err error) {
	if err != nil {
		t.Errorf("error %v", err)
	}
}
