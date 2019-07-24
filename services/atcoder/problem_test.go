package atcoder

import (
	"github.com/grassedge/go-online-judge-tools/types"

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

func TestAtCoderProblem_DownloadSampleCases(t *testing.T) {
	t.Run("contest: agc114, problem: agc114_c", func(t *testing.T) {
		u, _ := url.Parse("https://atcoder.jp/contests/abc114/tasks/abc114_c")
		problem, _ := NewAtCoderProblemFromUrl(u)
		actual := problem.DownloadSampleCases()
		expected := []*types.TestCase{
			&types.TestCase{
				Name: "Sample-1",
				InputName: "入力例 1",
				InputData: "575\n",
				OutputName: "出力例 1",
				OutputData: "4\n",
			},
			&types.TestCase{
				Name: "Sample-2",
				InputName: "入力例 2",
				InputData: "3600\n",
				OutputName: "出力例 2",
				OutputData: "13\n",
			},
			&types.TestCase{
				Name: "Sample-3",
				InputName: "入力例 3",
				InputData: "999999999\n",
				OutputName: "出力例 3",
				OutputData: "26484\n",
			},
		}
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("got: %v\nwant: %v", actual, expected)
		}
	})
}
