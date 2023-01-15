package schedule

import (
	"fmt"

	jira "github.com/andygrunwald/go-jira"
	"github.com/go-zen-chu/product-measurement/internal/config"
)

func ImportJira(config *config.Config) error {
	for _, jcnf := range config.Jira.JiraConfig {
		jc, _ := jira.NewClient(nil, jcnf.Endpoint)
		issue, _, _ := jc.Issue.Get("", nil)
		fmt.Printf("%s: %+v\n", issue.Key, issue.Fields.Summary)
	}
	return nil
}
