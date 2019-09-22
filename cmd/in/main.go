package main

import (
	"encoding/json"
	"fmt"
	"github.com/steinfletcher/new-relic-concourse-deployment-resource/concourse"
	"log"
	"os"
)

func main() {
	if err := json.NewEncoder(os.Stdout).Encode(concourse.InResponse{
		Version: concourse.Version{"ver": "static"},
	}); err != nil {
		log.Fatalln(fmt.Errorf("error writing 'in': %v", err))
	}
}
