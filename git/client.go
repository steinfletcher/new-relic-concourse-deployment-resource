package git

import (
	"fmt"
	"github.com/steinfletcher/new-relic-concourse-deployment-resource/domain"
	"os/exec"
	"strings"
)

func NewDeploymentInfoReader(repoPath string) domain.DeploymentInfoReader {
	return cli{repoPath: repoPath}
}

type cli struct {
	repoPath string
}

func (c cli) Read() (domain.DeploymentInfo, error) {
	sha1, err := c.exec("git", "rev-parse", "HEAD")
	if err != nil {
		return domain.DeploymentInfo{}, err
	}
	sha1 = strings.TrimSpace(sha1)

	author, err := c.exec("git", "--no-pager", "show", "-s", "--format=%ae", sha1)
	if err != nil {
		return domain.DeploymentInfo{}, err
	}

	message, err := c.exec("git", "--no-pager", "show", "-s", "--format=%B", sha1)
	if err != nil {
		return domain.DeploymentInfo{}, err
	}

	return domain.DeploymentInfo{
		User:        strings.TrimSpace(author),
		Description: strings.TrimSpace(message),
		Revision:    sha1,
	}, nil
}

func (c cli) exec(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...) // #nosec
	cmd.Dir = c.repoPath
	cmdOut, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("there was an error running command: '%s' args: '%v' err: %s", command, args, err)
	}
	return string(cmdOut), nil
}
