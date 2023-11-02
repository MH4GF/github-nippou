package lib_test

import (
	"context"
	"testing"

	"github.com/MH4GF/github-nippou/v4/lib"
	"github.com/google/go-github/github"
)

func TestFormatAll(t *testing.T) {
	issue := github.Issue{
		State:   github.String("closed"),
		Title:   github.String("イベントを取得できないことがある"),
		User:    &github.User{Login: github.String("MH4GF")},
		HTMLURL: github.String("https://github.com/MH4GF/github-nippou/issues/1"),
	}
	pr := github.PullRequest{
		State:   github.String("closed"),
		Title:   github.String("Bundle Update on 2015-10-04"),
		User:    &github.User{Login: github.String("deppbot")},
		HTMLURL: github.String("https://github.com/MH4GF/github-nippou/pull/31"),
		Merged:  github.Bool(true),
	}
	lines := lib.Lines{
		lib.NewLineByIssue("MH4GF/github-nippou", issue),
		lib.NewLineByPullRequest("MH4GF/github-nippou", pr),
	}
	settings := lib.Settings{}
	settings.Init("", "")

	ctx := context.Background()
	f := lib.NewFormat(ctx, nil, settings, false)

	result, err := f.All(lines)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	expected := `
### masutaka/github-nippou

* [イベントを取得できないことがある](https://github.com/masutaka/github-nippou/issues/1) by @[masutaka](https://github.com/masutaka) **closed!**
* [Bundle Update on 2015-10-04](https://github.com/masutaka/github-nippou/pull/31) by @[deppbot](https://github.com/deppbot) **merged!**
`

	if result != expected {
		t.Errorf("unexpected result: got %q, want %q", result, expected)
	}
}
