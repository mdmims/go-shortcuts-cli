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
			{Text: "github", Description: "Github repositories"},
			{Text: "golang", Description: "Golang resources"},
			{Text: "python", Description: "Python resources"},
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
