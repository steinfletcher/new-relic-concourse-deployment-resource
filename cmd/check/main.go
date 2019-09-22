package main

import (
	"encoding/json"
	"fmt"
	"github.com/steinfletcher/new-relic-concourse-deployment-resource/concourse"
	"log"
	"os"
)

func main() {
	if err := json.NewEncoder(os.Stdout).Encode(concourse.CheckResponse{}); err != nil {
		log.Fatalln(fmt.Errorf("error performing 'check': %v", err))
	}
}
