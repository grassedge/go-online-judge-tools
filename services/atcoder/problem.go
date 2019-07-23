package atcoder

import (
	"fmt"
	"net/url"
	"regexp"

	"github.com/PuerkitoBio/purell"
)

type AtCoderProblem struct {
	ContestId string
	ProblemId string
	lang string
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
	}, nil
}

func (p *AtCoderProblem) GetUrl() string {
	url := fmt.Sprintf("https://atcoder.jp/contests/%s/tasks/%s", p.ContestId, p.ProblemId)
	if p.lang != "" {
		url += fmt.Sprintf("?lang=%s", p.lang)
	}
	return url
}

func (p *AtCoderProblem) GetService() *AtCoderService {
	return &AtCoderService{}
}

func (p *AtCoderProblem) GetContest() *AtCoderContest {
	return &AtCoderContest{
		ContestId: p.ContestId,
	}
}
