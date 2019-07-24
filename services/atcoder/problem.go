package atcoder

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"

	"github.com/grassedge/go-online-judge-tools/types"
	"github.com/grassedge/go-online-judge-tools/utils"
	"github.com/PuerkitoBio/goquery"
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

func (p *AtCoderProblem) DownloadContent() *AtCoderProblemContent {
	resp, err := http.Get(p.GetUrl())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	return NewAtCoderProblemContentFromHTML(resp.Body, p)
}

func (p *AtCoderProblem) DownloadSampleCases() []*types.TestCase {
	return p.DownloadContent().sampleCases
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

type AtCoderProblemContent struct {
	sampleCases []*types.TestCase
	problem *AtCoderProblem
}

func NewAtCoderProblemContentFromHTML(html io.Reader, problem *AtCoderProblem) *AtCoderProblemContent {
	doc, err := goquery.NewDocumentFromReader(html)
	if err != nil {
		log.Fatal(err)
	}
	sampleCases := parseAtCoderProblemContentSampleCases(doc)
	// TODO: parse other content
	return &AtCoderProblemContent{
		sampleCases: sampleCases,
		problem: problem,
	}
}

func parseAtCoderProblemContentSampleCases(doc *goquery.Document) []*types.TestCase {
	zipper := utils.NewSampleZipper()
	doc.Find(".lang-ja h3+pre").Each(func(i int, s *goquery.Selection) {
		zipper.Add(s.Text(), s.Prev().Text())
	})
	doc.Find(".prettyprint").Each(func(i int, s *goquery.Selection) {
		zipper.Add(s.Text(), s.Parent().Prev().Text())
	})
	return zipper.Get()
}
