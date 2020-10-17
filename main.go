package main

import (
	"fmt"
	"github.com/mdmims/go-shortcuts-cli/commands"
	"github.com/mdmims/go-shortcuts-cli/mapping"
	"os"
	"regexp"
	"strings"

	prompt "github.com/c-bata/go-prompt"
)

//go:generate go run embed/embed.go

var shellCommands commands.Commands = commands.New()

var commandExpression = regexp.MustCompile(`(?P<command>github|golang|python)`)

var Input string

func getRegexGroups(text string) map[string]string {
	if !commandExpression.Match([]byte(text)) {
		return nil
	}

	match := commandExpression.FindStringSubmatch(text)
	result := make(map[string]string)
	for i, name := range commandExpression.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}
	return result
}

func completer(d prompt.Document) []prompt.Suggest {
	word := d.GetWordBeforeCursor()

	group := getRegexGroups(d.Text)
	if group != nil {
		// store the value of the incoming command
		Input = group["command"]

		// retrieve the SubCommands to display based on the input command
		return shellCommands.GetSubSuggestions(Input)
	}

	return prompt.FilterHasPrefix(shellCommands.GetMainSuggestions(), word, true)
}

func main() {
	fmt.Println("Please use `exit` to exit this program.")
	for {
		coCommand := prompt.Input(">>> shortcuts ",
			completer,
			prompt.OptionTitle("Code Orange Links"),
			prompt.OptionSelectedDescriptionTextColor(prompt.DarkGray),
			prompt.OptionInputTextColor(prompt.Green),
			prompt.OptionPrefixBackgroundColor(prompt.DarkGray),
			prompt.OptionPreviewSuggestionTextColor(prompt.Green),
			prompt.OptionSuggestionBGColor(prompt.DarkGray))

		splitCommands := strings.Split(coCommand, " ")
		if splitCommands[0] == "exit" {
			os.Exit(0)
		}

		// enforce at most 2 input args
		if len(splitCommands) > 2 {
			fmt.Println("Too many commands")
			continue
		}

		// validate input commands and then map the input to the correct yaml configurations
		// additionally, serve the corresponding website through the client default browser
		firstCommand := getRegexGroups(Input)
		if firstCommand != nil {
			mapping.Run(Input, splitCommands)
		} else {
			fmt.Println("Unknown command")
		}

	}
}
