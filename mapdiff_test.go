package mapdiff

import (
	"bytes"
	"io/ioutil"
	"strings"
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

	lines := strings.Split(expected, "\n")

	for _, line := range lines {
		if !strings.Contains(actual, line) {
			t.Fatalf("Expected diff to contain line: \n\n%v\n\nDiff: %v", line, actual)
		}
	}
}

func Test_CompareTwoDifferentObjects(t *testing.T) {
	p := "golden/different.golden"

	x := map[string]interface{}{
		"a": 1,
		"b": true,
		"c": "hello",
		"d": 'a',
	}

	y := map[string]interface{}{
		"a": 2,
		"b": true,
		"c": "hello",
		"e": 2.71,
	}

	result := Compare(x, y)
	if result.Equal {
		t.Fatalf("Expected result equality to be false. Got: %v", result.Equal)
	}

	raw, err := ioutil.ReadFile(p)
	if err != nil {
		t.Fatalf("Error reading response file at %v: %v", p, err)
	}

	expected := bytes.NewBuffer(raw).String()
	actual := result.Diff

	lines := strings.Split(expected, "\n")

	for _, line := range lines {
		if !strings.Contains(actual, line) {
			t.Fatalf("Expected diff to contain line: \n\n%v\n\nDiff: %v", line, actual)
		}
	}
}
