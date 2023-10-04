package workspace_imports

import (
	"encoding/json"
	"os/exec"

	"golang.org/x/exp/slices"
)

const (
	GO_LIST_JSON = "-json"
	IMPORTS      = "Imports"
)

type packageImports struct {
	Imports []string
}

func GetPackageImportsByPackageName(packageName string) ([]string, error) {
	var pkgImports packageImports

	pkgMetadata, err := getGoListDataByFlag(packageName, GO_LIST_JSON)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(pkgMetadata, &pkgImports)

	if err != nil {
		return nil, err
	}

	return pkgImports.Imports, nil
}

func GetAllTransitivelDependWorkspacePackages(pkgName string) ([]string, error) {
	var allPkgNames []string
	srcPkgImports, err := GetPackageImportsByPackageName(pkgName)
	iteration := 0

	if err != nil {
		return nil, err
	}

	allPackageNames, err := collectPackageNames(srcPkgImports, allPkgNames, iteration)

	if err != nil {
		return nil, err
	}

	return allPackageNames, nil
}

func collectPackageNames(currentPkgNames, allPackageNames []string, iteration int) ([]string, error) {
	for _, importPkgName := range currentPkgNames {
		if slices.Contains(allPackageNames, importPkgName) {
			continue
		}
		iteration++
		allPackageNames = append(allPackageNames, importPkgName)
		descendantPkgNames, err := GetPackageImportsByPackageName(importPkgName)

		if err != nil {
			return nil, err
		}

		if len(descendantPkgNames) == 0 {
			continue
		}

		allPackageNames, err = collectPackageNames(descendantPkgNames, allPackageNames, iteration)

		if err != nil {
			return nil, err
		}
	}

	return allPackageNames, nil
}

func getGoListDataByFlag(packageName, flag string) ([]byte, error) {
	args := []string{"list"}

	if flag != "" {
		args = append(args, flag)
	}

	if packageName != "" {
		args = append(args, packageName)
	}

	pkgMetadata, err := exec.Command("go", args...).Output()

	if err != nil {
		return nil, err
	}

	return pkgMetadata, nil
}
