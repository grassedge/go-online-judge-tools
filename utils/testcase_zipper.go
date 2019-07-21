package utils

import (
	"fmt"
	"log"
	"regexp"

	"github.com/grassedge/go-online-judge-tools/types"
)

type SampleZipper struct {
	testCases []*types.TestCase
	dangling *struct {
		name string
		content string
	}
}

func NewSampleZipper() *SampleZipper {
	return &SampleZipper{
		testCases: []*types.TestCase{},
		dangling: nil,
	}
}

func (z *SampleZipper) Add(content string, name string) {
	if z.dangling == nil {
		r1 := regexp.MustCompile(`(?i)output`)
		r2 := regexp.MustCompile(`出力`)
		if r1.MatchString(name) || r2.MatchString(name) {
			// TODO: logging
		}
		z.dangling = &struct {
			name string
			content string
		}{
			name,
			content,
		}
	} else {
		r1 := regexp.MustCompile(`(?i)input`)
		r2 := regexp.MustCompile(`入力`)
		if r1.MatchString(name) || r2.MatchString(name) {
			// TODO: logging
		}
		index := len(z.testCases)
		inputName := z.dangling.name
		inputContent := z.dangling.content
		z.testCases = append(z.testCases, &types.TestCase{
			Name: fmt.Sprintf("Sample-%d", index + 1),
			InputName: inputName,
			InputData: inputContent,
			OutputName: name,
			OutputData: content,
		})
		z.dangling = nil
	}
}

func (z *SampleZipper) Get() []*types.TestCase {
	if z.dangling != nil {
		log.Fatal(fmt.Sprintf("dangling sample string: %s", z.dangling.name))
	}
	return z.testCases
}
