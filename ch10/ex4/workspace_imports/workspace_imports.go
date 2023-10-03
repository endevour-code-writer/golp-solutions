package workspace_imports

import (
	"encoding/json"
	"os/exec"
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

// func makeListResult(packageName string)

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
