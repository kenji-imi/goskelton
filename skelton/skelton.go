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
	Project string
	User    string
	Dest    string
}

func (c *Config) validate() error {
	// project name
	if len(c.Project) == 0 {
		return fmt.Errorf("[ERROR] Project Name was empty")
	}

	// user name
	if len(c.User) == 0 {
		if os.Getenv("GOSKELTON_USER") != "" {
			c.User = os.Getenv("GOSKELTON_USER")
		} else {
			return fmt.Errorf("[ERROR] User Name was empty")
		}
	}

	// destination
	if len(c.Dest) > 0 {
		c.Dest = strings.TrimRight(c.Project, "/")
	} else if os.Getenv("GOSKELTON_DEST_DIR") != "" {
		c.Dest = os.Getenv("GOSKELTON_DEST_DIR")
	} else {
		c.Dest = "."
	}

	return nil
}

func Run(config *Config) error {
	if err := config.validate(); err != nil {
		return err
	}

	path := config.Dest + "/" + config.Project
	if err := os.Mkdir(path, 0755); err != nil {
		return err
	}
	log.Printf("[INFO] Created %s\n", path)

	out, err := os.Open(path)
	if err != nil {
		return err
	}
	defer out.Close()

	for targetFile, tmplStr := range map[string]string{
		"Makefile":                tmpl.MakefileTmpl,
		"main.go":                 tmpl.MainGoTmpl,
		"src/hello/hello.go":      tmpl.HelloGoTmpl,
		"src/hello/hello_test.go": tmpl.HelloTestGoTmpl,
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
				log.Printf("[INFO] Created %s\n", basedir)
			}
		}

		dest := path + "/" + targetFile
		f, err := os.Create(dest)
		defer f.Close()
		if err != nil {
			return err
		}
		err = t.Execute(f, map[string]interface{}{
			"Project": config.Project,
			"User":    config.User,
			"Path":    path,
		})
		if err != nil {
			return err
		}
		log.Printf("[INFO] Created %s\n", dest)
	}

	return nil
}
