package main

import (
	"encoding/json"
	"fmt"
	"github.com/steinfletcher/new-relic-concourse-deployment-resource/concourse"
	"github.com/steinfletcher/new-relic-concourse-deployment-resource/domain"
	"github.com/steinfletcher/new-relic-concourse-deployment-resource/file"
	"github.com/steinfletcher/new-relic-concourse-deployment-resource/git"
	"github.com/steinfletcher/new-relic-concourse-deployment-resource/http"
	"log"
	"os"
)

func main() {
	buildSourcesPath := os.Args[1]

	var input *concourse.OutRequest
	if err := json.NewDecoder(os.Stdin).Decode(&input); err != nil {
		log.Fatalln(fmt.Sprintf("error reading stdin: %v", err))
	}

	if err := input.Validate(); err != nil {
		log.Fatalln(fmt.Sprintf("input validation failed: %v", err))
	}

	newRelicClient := http.NewNewRelicClient(input.Source.NewRelicApiKey)
	var deploymentInfoReader domain.DeploymentInfoReader
	if input.Params.DeploymentInfoFilePath == "" {
		deploymentInfoReader = git.NewDeploymentInfoReader(
			fmt.Sprintf("%s/%s", buildSourcesPath, input.Params.RepoPath))
	} else {
		deploymentInfoReader = file.NewDeploymentInfoReader(input.Params.DeploymentInfoFilePath)
	}

	deploymentInfo, err := domain.RecordDeployment(
		deploymentInfoReader,
		newRelicClient,
		input.Params.NewRelicApplicationName,
		input.Params.NewRelicApplicationID)
	if err != nil {
		log.Fatalln(fmt.Errorf("failed to record deployment: %+v", err))
	}

	response := concourse.OutResponse{
		Version: concourse.Version{"ver": "static"},
		Metadata: []concourse.Metadata{
			{Name: "revision", Value: deploymentInfo.Revision},
			{Name: "author", Value: deploymentInfo.User},
			{Name: "description", Value: deploymentInfo.Description},
		},
	}
	if err := json.NewEncoder(os.Stdout).Encode(response); err != nil {
		log.Fatalln(fmt.Sprintf("error writing response to stdout. %v", err))
	}
}
