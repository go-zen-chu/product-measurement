package importer

import (
	"fmt"
	"log"

	"github.com/andygrunwald/go-jira"
	"github.com/go-zen-chu/product-measurement/internal/config"
)

type UseCaseImporter interface {
	UseCaseImportAll() error
}

type useCaseImporter struct {
	cnf *config.Config
}

func NewUseCaseImporter(cnf *config.Config) UseCaseImporter {
	return &useCaseImporter{
		cnf: cnf,
	}
}

// Import from all datasources
func (uci *useCaseImporter) UseCaseImportAll() error {
	if uci.cnf.DataSources.Jira != nil {
		for _, jcnf := range uci.cnf.DataSources.Jira.JiraConfigs {
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
			iss, _, err := jc.Issue.Search(fmt.Sprintf("project = %s AND issueType = Epic AND Status != Done", jcnf.Project), nil)
			if err != nil {
				return fmt.Errorf("search issue: %w", err)
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
	}
	return nil
}
