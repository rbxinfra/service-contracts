package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.rbx.com/roblox/service-contracts-generator/configuration"
	"github.rbx.com/roblox/service-contracts-generator/flags"
	"github.rbx.com/roblox/service-contracts-generator/templates"
)

var applicationName string
var buildMode string
var commitSha string

// Pre-setup, runs before main.
func init() {
	flags.SetupFlags(applicationName, buildMode, commitSha)
}

// Main entrypoint.
func main() {
	if *flags.HelpFlag {
		flag.Usage()

		return
	}

	configurationMap, err := configuration.Parse()
	if err != nil {
		panic(err)
	}

	files, err := templates.ParseForConfiguration(configurationMap)
	if err != nil {
		panic(err)
	}

	projectPaths := []string{}

	for fileName, fileContents := range files {
		fileNameOut := path.Join(*flags.OutputDirectoryFlag, fileName)
		pathName := path.Dir(fileNameOut)

		if err := os.MkdirAll(pathName, os.ModePerm); err != nil {
			panic(err)
		}

		file, err := os.Create(fileNameOut)
		if err != nil {
			panic(err)
		}

		defer file.Close()

		_, err = file.WriteString(fileContents)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Wrote file: %s\n", fileNameOut)

		projectPaths = append(projectPaths, fileName)
	}

	if *flags.CreateSolution {
		solutionPath := path.Join(*flags.OutputDirectoryFlag, "Protos.sln")

		if err := os.Remove(solutionPath); err != nil && !os.IsNotExist(err) {
			panic(err)
		}

		createSolutionCommand := exec.Command("dotnet", "new", "sln", "--name", "Protos")
		createSolutionCommand.Dir = *flags.OutputDirectoryFlag
		stdout, err := createSolutionCommand.Output()

    	if err != nil {
        	fmt.Println(err.Error())

        	return
    	}

    	fmt.Println(string(stdout))

		addToSolutionCommand := exec.Command("dotnet", "sln", "add")
		addToSolutionCommand.Args = append(addToSolutionCommand.Args, projectPaths...)
		addToSolutionCommand.Dir = *flags.OutputDirectoryFlag
		
		stdout, err = addToSolutionCommand.Output()

    	if err != nil {
        	fmt.Println(err.Error())

        	return
    	}

    	fmt.Println(string(stdout))
	}
}
