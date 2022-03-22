package main

import (
	"fmt"
	"github.com/openshift-online/ocm-cli/pkg/config"
	"github.com/openshift-online/ocm-cli/pkg/ocm"
)

func main() {

	config, err := config.Load()
	if err != nil {
		fmt.Printf("failed to load OCM config: %v", err)
	}
	if config == nil {
		fmt.Printf("not logged in, please run the 'ocm login' command")
	}

	connection, err := ocm.NewConnection().Config(config).Build()
	if err != nil {
		fmt.Printf("failed to create OCM connection: %v", err)
	}
	defer connection.Close()

	url := connection.URL()
	fmt.Printf("URL: %s\n", url)
}