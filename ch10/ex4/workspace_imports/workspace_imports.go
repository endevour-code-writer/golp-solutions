package workspace_imports

import (
	"fmt"
	"os"
	"os/exec"
)

const (
	GO_LIST_JSON = "-json"
)

type listResult struct {
}

func GetPackageImportsByPackageName(packageName string) ([]byte, error) {
	pkgMetadata, err := getGoListDataByFlag(packageName, GO_LIST_JSON)

	fmt.Println(string(pkgMetadata))
	os.Exit(21)

	if err != nil {
		return nil, err
	}

	// err = json.Unmarshal(pkgMetadata, v)
	// // jsonDecoder := json.NewDecoder(strings.NewReader(string(pkgMetadata)))

	// fmt.Println(err)
	// os.Exit(25)
	return pkgMetadata, nil

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
