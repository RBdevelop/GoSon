package apiconv

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestReadSourceFile(t *testing.T) {
	data, err := readSourceFile("serviceA.json")
	if err != nil {
		t.Errorf("error reading source file: %v", err)
	}
	spew.Dump(data)
}
