package skelton

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/kenji-imi/goskelton/tmpl"
)

type Config struct {
	Dest string
	Name string
}

func (c *Config) validate() error {
	if len(c.Name) == 0 {
		msg := "goskelton: Name was empty"
		return fmt.Errorf(msg)
	}
	if len(c.Dest) > 0 {
		c.Dest = strings.TrimRight(c.Dest, "/")
	} else if os.Getenv("GOSKELTON_DESTINATION_HOME") != "" {
		c.Dest = os.Getenv("GOSKELTON_DESTINATION_HOME")
	} else {
		c.Dest = "."
	}
	log.Printf("goskelton: Dest directory is %s.\n", c.Dest)

	return nil
}

func Run(config *Config) error {
	if err := config.validate(); err != nil {
		return err
	}

	path := config.Dest + "/" + config.Name
	if err := os.Mkdir(path, 0755); err != nil {
		return err
	}
	out, err := os.Open(path)
	if err != nil {
		return err
	}
	defer out.Close()

	for targetFile, tmplStr := range map[string]string{
		"Makefile":           tmpl.MakefileTmpl,
		"main.go":            tmpl.MainGoTmpl,
		"src/hello/hello.go": tmpl.HelloGoTmpl,
	} {
		t, err := template.New("").Parse(tmplStr)
		if err != nil {
			return err
		}

		basedir := path + "/" + filepath.Dir(targetFile)
		if basedir != "" {
			if _, err := os.Stat(basedir); err != nil {
				if err := os.MkdirAll(basedir, 0755); err != nil {
					return err
				}
			}
		}

		dest := path + "/" + targetFile
		f, err := os.Create(dest)
		defer f.Close()
		if err != nil {
			return err
		}
		err = t.Execute(f, map[string]interface{}{
			"Path": path,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
