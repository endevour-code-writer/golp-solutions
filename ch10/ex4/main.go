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

	pkgImportsViaDeps, err := workspace_imports.GetAllTransitivelDependWorkspacePackagesViaDeps(packageName)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("imports via recursions:\n %v\n", pkgImports)
	fmt.Printf("number imports via recursions:\n %v\n", len(pkgImports))
	fmt.Printf("imports via deps:\n %v\n", pkgImportsViaDeps)
	fmt.Printf("number imports via deps:\n %v\n", len(pkgImportsViaDeps))
}
