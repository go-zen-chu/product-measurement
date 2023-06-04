package schedule

import (
	"errors"
	"fmt"
	"log"

	jira "github.com/andygrunwald/go-jira"
	"github.com/go-zen-chu/product-measurement/internal/config"
)

type JiraBoardList struct {
	// 保存したいデータを書く
	Values []JiraBoard
}

type JiraBoard struct {
	ID int
	Type string
}

func ImportJira(config *config.Config) error {
	if config.Jira == nil {
		return errors.New("JIRA config is empty")
	}
	for _, jcnf := range config.Jira.JiraConfig {
		var jc *jira.Client
		var err error
		if jcnf.JiraAuth != nil {
			switch jcnf.JiraAuth.Method {
			case "basic":
				tp := jira.BasicAuthTransport{
					Username: jcnf.JiraAuth.User,
					Password: jcnf.JiraAuth.Password,
				}
				jc, err = jira.NewClient(tp.Client(), jcnf.Endpoint)
				if err != nil {
					return fmt.Errorf("create new jira client: %w", err)
				}
			default:
				return fmt.Errorf("unsupported method: %s", jcnf.JiraAuth.Method)
			}
		}
		bl, _, err := jc.Board.GetAllBoards(nil)
		if err != nil {
			return err
		}
		for _, b := range bl.Values {
			b.ID
		}


		log.Printf("%+v", bl)
		issue, _, err := jc.Issue.Get("TEST-1", nil)
		if err != nil {
			return fmt.Errorf("get issue: %w", err)
		}
		fmt.Printf("%s: %+v\n", issue.Key, issue.Fields.Summary)
	}
	return nil
}
