package main

import (
	"fmt"
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Example struct {
	Debug      bool
	Port       int
	User       string
	Users      []string
	Rate       float32
	Timeout    time.Duration
	ColorCodes map[string]int
}

func PrintExample() {
	var ex Example
	err := envconfig.Process("myapp", &ex)
	if err != nil {
		log.Fatal(err.Error())
	}
	format := "Debug: %v\nPort: %d\nUser: %s\nRate: %f\nTimeout: %s\n"
	_, err = fmt.Printf(format, ex.Debug, ex.Port, ex.User, ex.Rate, ex.Timeout)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Users:")
	for _, u := range ex.Users {
		fmt.Printf("  %s\n", u)
	}

	fmt.Println("Color codes:")
	for k, v := range ex.ColorCodes {
		fmt.Printf("  %s: %d\n", k, v)
	}
}

type Specification struct {
	ManualOverride1         string `envconfig:"manual_override_1"`
	DefaultVar              string `default:"foobar"`
	RequiredVar             string `required:"true"`
	IgnoredVar              string `ignored:"true"`
	AutoSplitVar            string `split_words:"true"`
	RequiredAndAutoSplitVar string `required:"true" split_words:"true"`
}

func PrintSpecification() {
	var spec Specification
	err := envconfig.Process("myapp", &spec)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("default: %v\n", spec.DefaultVar)
	fmt.Printf("required_split: %v\n", spec.RequiredAndAutoSplitVar)
	fmt.Printf("ignored_var: %v\n", spec.IgnoredVar)
}

func main() {
	PrintExample()
	PrintSpecification()
}
