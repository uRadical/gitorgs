package main

import (
	_ "embed"
	"fmt"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/AlecAivazis/survey/v2"
	"github.com/tcnksm/go-gitconfig"
)

//go:embed gitconfig.tmpl
var gitConfigTemplate string

//go:embed includeif.tmpl
var includeIfTemplate string

type ConfigData struct {
	Email      string
	SigningKey string
	Key        string
	Dir        string
	Org        string
}

func availableSshKeys() []string {
	files, err := os.ReadDir(os.ExpandEnv("$HOME/.ssh"))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	var keyNames []string
	for _, file := range files {
		name := file.Name()
		if !strings.HasSuffix(name, ".pub") && !file.IsDir() && name != "config" && name != "known_hosts" {
			keyNames = append(keyNames, name)
		}
	}
	return keyNames
}

func defaultEmail() string {
	defaultEmail, err := gitconfig.Global("user.email")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return defaultEmail
}

func writeGitConfig(d *ConfigData) {
	tmpl, err := template.New("gitconfig").Parse(gitConfigTemplate)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	outputFile, err := os.Create(".gitconfig")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer outputFile.Close()

	err = tmpl.Execute(outputFile, d)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func globalGitConfigIncludeIf(d *ConfigData) {
	fmt.Println("Run the following to update your global gitconfig:")
	tmpl, err := template.New("gitconfig").Parse(includeIfTemplate)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	err = tmpl.Execute(os.Stdout, d)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func main() {
	d := &ConfigData{}

	var qs = []*survey.Question{
		{
			Name: "email",
			Prompt: &survey.Input{
				Message: "Email:",
				Default: defaultEmail(),
			},
		},
	}

	err := survey.Ask(qs, d)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	prompt := &survey.Select{
		Message: "SSH key:",
		Options: availableSshKeys(),
	}

	var s string
	err = survey.AskOne(prompt, &s)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	d.Key = s

	qs = []*survey.Question{
		{
			Name: "SigningKey",
			Prompt: &survey.Input{
				Message: "Signing Key:",
				Help:    "use gpg --list-keys to see your available keys, then paste the key ID here",
			},
		},
	}
	err = survey.Ask(qs, d)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	wd, _ := os.Getwd()
	d.Dir = wd
	d.Org = path.Base(wd)

	writeGitConfig(d)

	globalGitConfigIncludeIf(d)
}
