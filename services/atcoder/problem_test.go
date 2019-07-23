package atcoder

import (
	"net/url"
	"reflect"
	"testing"
)

func TestAtCoderProblemNewFromUrl(t *testing.T) {
	u, _ := url.Parse("https://atcoder.jp/contests/agc030/tasks/agc030_c")
	problem, _ := NewAtCoderProblemFromUrl(u)
	{
		actual := problem.ContestId
		expected := "agc030"
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("got: %v\nwant: %v", actual, expected)
		}
	}
	{
		actual := problem.ProblemId
		expected := "agc030_c"
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("got: %v\nwant: %v", actual, expected)
		}
	}
}
