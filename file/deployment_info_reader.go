package file

import (
	"encoding/json"
	"github.com/steinfletcher/new-relic-concourse-deployment-resource/domain"
	"io/ioutil"
)

type reader struct {
	deploymentInfoPath string
}

func NewDeploymentInfoReader(path string) domain.DeploymentInfoReader {
	return reader{deploymentInfoPath: path}
}

func (r reader) Read() (domain.DeploymentInfo, error) {
	data, err := ioutil.ReadFile(r.deploymentInfoPath)
	if err != nil {
		return domain.DeploymentInfo{}, err
	}
	var info domain.DeploymentInfo
	err = json.Unmarshal(data, &info)
	if err != nil {
		return domain.DeploymentInfo{}, err
	}
	return info, nil
}
