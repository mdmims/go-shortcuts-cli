package mapping

import (
	"encoding/json"
	"fmt"
	"github.com/mdmims/go-shortcuts-cli/embed"
	"gopkg.in/yaml.v2"
	"log"
	"os/exec"
	"runtime"
)

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}

type github struct {
	Mdmims   string `yaml:"Mdmims" json:"Mdmims"`
	AzureSDK string `yaml:"AzureSDK" json:"AzureSDK"`
	GoSDK    string `yaml:"GoSDK" json:"GoSDK"`
}

type golang struct {
	Docs        string `yaml:"Docs" json:"Docs"`
	StandardLib string `yaml:"StandardLib" json:"StandardLib"`
}

type python struct {
	StandardLib string `yaml:"StandardLib" json:"StandardLib"`
}

func Run(parentOption string, option []string) {
	// parse the incoming split argument slice to extract the elements
	var mainArg string
	if len(option) == 1 {
		mainArg = option[0]
	} else {
		mainArg = option[1]
	}

	// determine which .yml file to access based on input and main/first argument
	yamlFile := parseYaml(parentOption)

	// access the correct struct based on condition and main arg chosen
	var mapper interface{}
	switch parentOption {
	case "github":
		mapper = &github{}
	case "golang":
		mapper = &golang{}
	case "python":
		mapper = &python{}
	}

	// parse and read the yaml definitions
	err := yaml.Unmarshal(yamlFile, mapper)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	// convert the struct to a map
	j, _ := toMap(mapper)

	// parse the map and retrieve the requested value pair and open the website
	selection := j[mainArg]
	str := fmt.Sprintf("%v", selection)
	fmt.Printf("Opening site: %s \n", str)
	openBrowser(str)
}

func parseYaml(s string) []byte {
	parsedPath := fmt.Sprintf("/%s.yml", s)
	// read the binary Config file holding the .yml configs
	binaryConfig := embed.Config
	// extract the specific .yml file needed
	return binaryConfig[parsedPath]
}

// toMap converts input into JSON and then converts that into map[string]string
func toMap(c interface{}) (map[string]interface{}, error) {
	e, err := json.Marshal(&c)
	if err != nil {
		log.Printf("Error building JSON")
	}

	// convert json to map
	var result map[string]interface{}
	err = json.Unmarshal(e, &result)
	return result, nil
}
