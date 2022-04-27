package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type SingleOrArray[T any] []T

func (s *SingleOrArray[T]) UnmarshalYAML(value *yaml.Node) error {
	var v T
	if err := value.Decode(&v); err == nil {
		*s = []T{v}
		return nil
	}

	var arr []T
	if err := value.Decode(&arr); err == nil {
		*s = arr
		return nil
	}

	return fmt.Errorf("element must be of type %T or %T", v, arr)
}

func main() {
	var data struct {
		Foo SingleOrArray[string] `yaml:"foo"`
	}

	file, err := os.Open("./data.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if err := yaml.NewDecoder(file).Decode(&data); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Items: %v\n", data.Foo)
}
