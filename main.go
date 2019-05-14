package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/v25/github"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "{ACCESS_TOKEN}"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	repoUrls := make(map[string]bool)

	repos, _, _ := client.Repositories.List(ctx, "", nil)
	for _, el1 := range repos {
		repoUrls[*el1.URL] = true
	}

	s := fmt.Sprintf("type:pr is:private is:open org:applift ")
	issues, _, _ := client.Search.Issues(ctx, s, nil)
	for _, el2 := range issues.Issues {
		if repoUrls[*el2.RepositoryURL] {
			fmt.Println(*el2.HTMLURL)
		}
	}
}
