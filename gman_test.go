// gman_test.go
package main

import (
	"testing"
)

func TestGman(t *testing.T) {
	_, err := Gman("colamone")
	if err != nil {
		t.Error(err)
		return
	}
}
func TestShowList(t *testing.T) {
	_, err := ShowList("colamone")
	if err != nil {
		t.Error(err)
		return
	}
}
func TestOpenRepo(t *testing.T) {
	_, err := OpenRepo("kurehajime/gman")
	if err != nil {
		t.Error(err)
		return
	}
}
