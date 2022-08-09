package confg

import (
	"gopkg.in/yaml.v2"
	"io"
	"log"
	"os"
)

type conf struct {
	Addr string `yaml:"port"`
}

func New() *conf {
	return &conf{}
}

func (c *conf) Load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	all, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(all, c)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
