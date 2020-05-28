//go:generate go run gen.go
package main

import (
	"fmt"
	"github.com/golangee/bundle"
	"github.com/golangee/http"
	"github.com/golangee/openapi-client/async"
	v3 "github.com/golangee/openapi/v3"
	"github.com/golangee/reflectplus"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	err := generate()
	if err != nil {
		panic(err)
	}
}

func generate() error {
	modDir, err := ModRootDir()
	if err != nil {
		return err
	}

	pkg := reflectplus.Must(reflectplus.Generate(modDir))
	// load into current reflectplus context
	reflectplus.AddPackage(*pkg)

	doc, err := makeOpenAPI(modDir, pkg)
	if err != nil {
		return err
	}

	if err := makeOpenAPIClient(doc); err != nil {
		return err
	}

	if err := makeBundle(); err != nil {
		return err
	}

	return nil
}

func makeBundle() error {
	return bundle.Embed(bundle.Options{
		TargetDir:   "internal/resources",
		PackageName: "resources",
		Include:     []string{"doc/openapi/apidoc.json"},
	})
}

func makeOpenAPIClient(doc *v3.Document) error {
	return async.Generate([]byte(doc.String()), async.Options{
		TargetDir:     "webapp/internal/client",
		TargetPackage: "client",
	})
}

func makeOpenAPI(dir string, pkg *reflectplus.Package) (*v3.Document, error) {
	doc := v3.NewDocument()
	doc.Info.Version = "tbd" //TODO depends on commit?
	doc.Info.Title = "mercurius"
	doc.Servers = []v3.Server{
		{Url: "http://localhost:" + strconv.Itoa(8080)}, //TODO depends on runtime
	}

	err := http.MakeDoc(doc, pkg.AllStructs())
	if err != nil {
		return nil, fmt.Errorf("failed to create OpenAPI specification: %w", err.Error())
	}

	err = ioutil.WriteFile(filepath.Join(dir, "doc", "openapi", "apidoc.json"), []byte(doc.String()), os.ModePerm)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

// ModRootDir returns the root directory of current module. If the current working directory is not a module
// returns an error.
func ModRootDir() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	root := cwd
	for {
		stat, err := os.Stat(filepath.Join(root, "go.mod"))
		if err == nil && stat.Mode().IsRegular() {
			return root, nil
		}
		root = filepath.Dir(root)
		if root == "/" || root == "." {
			return "", fmt.Errorf("%s is not withing a go module", cwd)
		}
	}
}
