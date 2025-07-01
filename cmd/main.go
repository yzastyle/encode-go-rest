package main

import (
	"fmt"
	"log"

	"github.com/yzastyle/encode-go-rest/internal"
)

func main() {
	appConfig, err := internal.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	fmt.Println("Application Name:", appConfig.Name)
}
