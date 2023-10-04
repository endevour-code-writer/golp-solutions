package main

import (
	"GoTheProgrammingLanguage/ch10/ex4/workspace_imports"
	"fmt"
	"os"
)

func main() {
	var packageName string
	args := os.Args[1:]

	if len(args) > 0 {
		packageName = args[0]
	}

	pkgImports, err := workspace_imports.GetAllTransitivelDependWorkspacePackages(packageName)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(pkgImports)
	fmt.Println(len(pkgImports))
}
