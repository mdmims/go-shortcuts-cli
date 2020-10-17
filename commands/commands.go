package commands

import (
	"github.com/c-bata/go-prompt"
)

type Commands struct {
	MainSuggestions []prompt.Suggest
	SubSuggestions  map[string][]prompt.Suggest
}

func New() Commands {
	return Commands{
		MainSuggestions: []prompt.Suggest{
			{Text: "confluence", Description: "Confluence documentation for EDAP"},
			{Text: "coretech", Description: "CoreTech releated resources"},
			{Text: "github", Description: "Github EDAP repoistories"},
			{Text: "jira", Description: "JIRA releated stories and release pages"},
			{Text: "rps", Description: "RPS client related links"},
		},
		SubSuggestions: map[string][]prompt.Suggest{
			"github": {
				prompt.Suggest{Text: "Mdmims", Description: "Personal Github Repo"},
			},
			"go": {
				prompt.Suggest{Text: "Docs", Description: "Golang documentation"},
				prompt.Suggest{Text: "StandardLib", Description: "Golang standard library docs"},
			},
			"python": {
				prompt.Suggest{Text: "StandardLib", Description: "Python standard library docs"},
			},
		},
	}
}
func (c *Commands) GetMainSuggestions() []prompt.Suggest {
	return c.MainSuggestions
}

func (c *Commands) GetSubSuggestions(key string) []prompt.Suggest {
	sub := c.SubSuggestions
	return sub[key]
}
