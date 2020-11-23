# go-shortcuts-cli
***go-shortcuts-cli*** is an interactive cli client that aims to provide helpful resource links to commonly used sites and documents and features fast auto-complete and suggestions for available options. Once a valid command set has been entered ***go-shortcuts-cli*** will automatically open the respective url inside the default web browser.

### Usage:
![](shortcuts_demo.gif)

- Upon selecting a link through the first main option the app will automatically open the link through the configured default web browser.
- Additionally, the URL will be printed to the terminal.

### Requirements:
Go binary distribution installed. (https://golang.org/) Go 1.15+

### Installation:
Clone the repository:
```
git clone https://github.com/mdmims/go-shortcuts-cli.git
```

### Compile binary within project directory:
```
go build
```

### To install globally on OS:

- Will place binary by default into: $GOPATH/bin
- Must have \$GOPATH defined and binaries listed in $PATH
- GOPATH docs: https://golang.org/cmd/go/#hdr-GOPATH_environment_variable

```
go install
```

### Supported commands with page links available:

- [commands.go](commands/commands.go)
- Explicit definitions for URLs defined within respective YAML file within:
  - [config](config)
  
  
### Supported tests (In progress):

- [mapping_test.go](lib/mapping_test.go)