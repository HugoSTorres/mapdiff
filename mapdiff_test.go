package mapdiff

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func Test_CompareTwoEqualObjects(t *testing.T) {
	p := "golden/equal.golden"

	x := map[string]interface{}{
		"a": 1,
		"b": true,
		"c": "hello",
	}

	y := map[string]interface{}{
		"a": 1,
		"b": true,
		"c": "hello",
	}

	result := Compare(x, y)
	if !result.Equal {
		t.Fatalf("Expected result equality to be true. Got: %v", result.Equal)
	}

	raw, err := ioutil.ReadFile(p)
	if err != nil {
		t.Fatalf("Error reading response file at %v: %v", p, err)
	}

	expected := bytes.NewBuffer(raw).String()
	actual := result.Diff

	if expected != actual {
		t.Fatalf("Expected: \n\n%v\n\nGot: \n\n%v\n\n", expected, actual)
	}
}
