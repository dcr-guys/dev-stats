package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/fguisso/dev-stats/conn"
	"github.com/google/go-github/v29/github"
)

func main() {
	cfg, err := loadConfig()
	if err != nil {
		log.Printf("%v", err)
		return
	}
	ctx := context.Background()
	gc := conn.InitClient(cfg.GithubToken, ctx)

	var opt = &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{PerPage: 10},
	}
	var allRepos []*github.Repository
	for _, user := range cfg.Users {
		for {
			repos, resp, err := gc.Repositories.ListByOrg(ctx, user, opt)
			if err != nil {
				log.Printf("%v", err)
				return
			}
			allRepos = append(allRepos, repos...)
			if resp.NextPage == 0 {
				break
			}
			opt.ListOptions.Page = resp.NextPage
		}
	}
	pages := 0
	for _, repo := range allRepos {
		if pages == 10 {
			return
		}
		fmt.Printf("\nUser: %v\nRepo: %v\n",
			*repo.Owner.Login, *repo.Name)
		fmt.Printf(" Forks: %v\n Open Issues: %v\n Stars: %v\n",
			*repo.ForksCount, *repo.OpenIssuesCount,
			*repo.StargazersCount)
		stats, _, err := gc.Repositories.ListContributorsStats(ctx,
			fmt.Sprintf("%v", *repo.Owner.Login), fmt.Sprintf("%v", *repo.Name))
		if err != nil {
			log.Printf("%v", err)
			time.Sleep(time.Duration(15) * time.Second)
			stats, _, _ = gc.Repositories.ListContributorsStats(ctx,
				fmt.Sprintf("%v", *repo.Owner.Login), fmt.Sprintf("%v", *repo.Name))
		}
		fmt.Printf(" Contributors: %v\n", len(stats))
		pages++
	}
}
