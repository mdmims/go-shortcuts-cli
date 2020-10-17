package embed

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	embed("config")
}

func embed(name string) {
	out, _ := os.Create("embed/" + name + ".go")
	_, _ = out.Write([]byte("package embed"))
	_, _ = out.Write([]byte("\n\n// Automatically generated file through go generate."))
	_, _ = out.Write([]byte("\n\n// " + strings.Title(name) + ": scaffolding used in 'build' command"))
	_, _ = out.Write([]byte("\nvar " + strings.Title(name) + " = map[string][]byte{\n"))
	_ = filepath.Walk(name,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				// Get the contents of the current file.
				content, _ := ioutil.ReadFile(path)
				// Correct filename of the .gitignore file.
				if strings.HasSuffix(path, "_plenti_replace") {
					path = strings.TrimSuffix(path, "_plenti_replace")
				}
				// Add a key for the filename to the map.
				_, _ = out.Write([]byte("\t\"" + strings.TrimPrefix(path, name) + "\": []byte(`"))
				// Escape the backticks that would break string literals
				escapedContent := strings.Replace(string(content), "`", "`+\"`\"+`", -1)
				// Add the content as the value of the map.
				_, _ = out.Write([]byte(escapedContent))
				// End the specific file entry in the map.
				_, _ = out.Write([]byte("`),\n"))
			}
			return nil
		})
	_, _ = out.Write([]byte("}\n"))
}
