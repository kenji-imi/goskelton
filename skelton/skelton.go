package skelton

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
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
	} else if os.Getenv("GOSKELTON_HOME") != "" {
		c.Dest = os.Getenv("GOSKELTON_HOME")
	} else {
		c.Dest = "."
		// c.Dest = "/Users/kimai/go/src/github.com/kenji-imi"
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

	for _, targetFile := range []string{
		"Makefile",
		"main.go",
		"src/hello/hello.go",
	} {
		var srcfd *os.File
		var dstfd *os.File

		src := "./template/" + targetFile
		if srcfd, err = os.Open(src); err != nil {
			return err
		}
		defer srcfd.Close()

		basedir := path + "/" + filepath.Dir(targetFile)
		if basedir != "" {
			if _, err := os.Stat(basedir); err != nil {
				if err := os.MkdirAll(basedir, 0755); err != nil {
					return err
				}
			}
		}

		dst := path + "/" + targetFile
		if dstfd, err = os.Create(dst); err != nil {
			return err
		}
		defer dstfd.Close()

		if _, err := io.Copy(dstfd, srcfd); err != nil {
			panic(err)
		}
	}

	return nil
}
