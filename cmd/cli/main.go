package main

import "github.com/wizedkyle/artifactsmmo/v2/internal/cmd"

func main() {
	root := cmd.NewCmdRoot()
	root.Execute()
}
