package roper_test

import (
	"github.com/cdornsife/roper"
	"reflect"
	"testing"
)

var controlTestStruct *TestStruct

func init() {
	controlTestStruct = &TestStruct{
		One: 1,
		Two: "two",
	}

	controlTestStruct.Three = append(controlTestStruct.Three, struct {
		A string `json:"a"`
		B string `json:"b"`
		C string `json:"c"`
	}{"a", "b", "c"})

}

// TestStruct represents the data in the file
type TestStruct struct {
	One   int    `json:"one"`
	Two   string `json:"two"`
	Three []struct {
		A string `json:"a"`
		B string `json:"b"`
		C string `json:"c"`
	} `json:"three"`
}

// TestRoper_JSON load and test tests/test.json
func TestRoper_JSON(t *testing.T) {
	doIt("tests/test.json", t)
}

// TestRoper_YAMLFile load and test tests/test.yaml
func TestRoper_YAML(t *testing.T) {
	doIt("tests/test.yaml", t)
}

// TestRoper_YMLFile load and test tests/test.yml
func TestRoper_YML(t *testing.T) {
	doIt("tests/test.yml", t)
}

// TestRoper_YMLFile load and test tests/test.yml
func TestRoper_URL(t *testing.T) {
	doIt("https://raw.githubusercontent.com/applariat/roper/master/tests/test.yaml", t)
}

// TestRoper_YMLFile load and test tests/test.yml
func TestRoper_STDIN(t *testing.T) {
	// TODO:
	//doIt("https://raw.githubusercontent.com/applariat/roper/master/tests/test.yaml", t)
}


func doIt(file string, t *testing.T) {

	var out = new(TestStruct)

	err := roper.Unmarshal(file, &out)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(out, controlTestStruct) {
		t.Fatal("results don't match")
	}

}
