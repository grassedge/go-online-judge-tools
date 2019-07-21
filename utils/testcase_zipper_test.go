package utils

import (
	"reflect"
	"testing"

	"github.com/grassedge/go-online-judge-tools/types"
)

func TestSampleZipper(t *testing.T) {
	zipper := NewSampleZipper()
	{
		actual := zipper.Get()
		if len(actual) != 0 {
			t.Errorf("got: %v\nwant: %v", len(actual), 0)
		}
	}
	zipper.Add("0 0 0", "input1")
	zipper.Add("1 1 1", "output1")
	{
		actual := zipper.Get()
		expected := []*types.TestCase{&types.TestCase{
			Name: "Sample-1",
			InputName: "input1",
			InputData: "0 0 0",
			OutputName: "output1",
			OutputData: "1 1 1",
		}}
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("got: %v\nwant: %v", actual, expected)
		}
	}
}
