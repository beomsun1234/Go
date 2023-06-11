package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Properties struct {
	Config Config
}
type Config struct {
	Addr   string `yaml:"addr"`
	Mobile string `yaml:"mobile"`
}

/*
https://yaml2go.prasadg.dev/   -> convert yaml
*/
func main() {

	y, err := ioutil.ReadFile("beom.yaml")
	if err != nil {
		panic("error")
	}
	p := &Properties{}
	yaml.Unmarshal(y, p)

	fmt.Println(p)

}
