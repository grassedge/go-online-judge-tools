package atcoder

import (
	"net/url"
	"regexp"

	"github.com/PuerkitoBio/purell"
)

type AtCoderProblemType int
const (
	Default AtCoderProblemType = iota
	Beta
	Old
)
type AtCoderProblem struct {
	ContestId string
	ProblemId string
	problemType AtCoderProblemType
	lang string
}

type option func (p *AtCoderProblem) error
func Type(t AtCoderProblemType) option {
	return func (p *AtCoderProblem) error {
		p.problemType = t
		return nil
	}
}

func NewAtCoderProblemFromUrl(u *url.URL) (*AtCoderProblem, error) {
	normalized := purell.NormalizeURL(
		u,
		purell.FlagRemoveTrailingSlash|purell.FlagRemoveDotSegments,
	)
	normalizedUrl, err := url.Parse(normalized)
	if err != nil {
		return nil, err
	}
	path := normalizedUrl.Path
	r := regexp.MustCompile(`^/contests/([\w\-_]+)/tasks/([\w\-_]+)$`)
	group := r.FindSubmatch([]byte(path))

	contestId := string(group[1])
	problemId := string(group[2])

	return &AtCoderProblem{
		ContestId: contestId,
		ProblemId: problemId, // NOTE: AtCoder calls this as "task_screen_name"
		problemType: Default,
	}, nil
}
