package atcoder

import (
	"net/url"
	"reflect"
	"testing"
)

func TestNewAtCoderProblemFromUrl(t *testing.T) {
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

func TestAtCoderProblem_GetUrl(t *testing.T) {
	t.Run("no lang", func(t *testing.T) {
		problem := &AtCoderProblem{
			ContestId: "agc030",
			ProblemId: "agc030_c",
		}
		actual := problem.GetUrl()
		expected := "https://atcoder.jp/contests/agc030/tasks/agc030_c"
		if actual != expected {
			t.Errorf("got: %v\nwant: %v", actual, expected)
		}
	})

	t.Run("lang is en", func(t *testing.T) {
		problem := &AtCoderProblem{
			ContestId: "agc030",
			ProblemId: "agc030_c",
			lang     : "en",
		}
		actual := problem.GetUrl()
		expected := "https://atcoder.jp/contests/agc030/tasks/agc030_c?lang=en"
		if actual != expected {
			t.Errorf("got: %v\nwant: %v", actual, expected)
		}
	})
}
