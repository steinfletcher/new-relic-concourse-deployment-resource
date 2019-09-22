package domain

//go:generate mockgen -source deployment.go -destination ../mocks/deployment.go -package mocks

type DeploymentInfo struct {
	Revision    string `json:"revision"`
	Description string `json:"description"`
	User        string `json:"user"`
}

type DeploymentWriter interface {
	Write(deployment DeploymentInfo, applicationID string) error
	WriteWithName(deployment DeploymentInfo, applicationName string) error
}

type DeploymentInfoReader interface {
	Read() (DeploymentInfo, error)
}

func RecordDeployment(
	reader DeploymentInfoReader,
	writer DeploymentWriter,
	applicationName string,
	applicationID string,
) (DeploymentInfo, error) {
	info, err := reader.Read()
	if err != nil {
		return DeploymentInfo{}, err
	}

	if applicationID != "" {
		return info, writer.Write(info, applicationID)
	}
	return info, writer.WriteWithName(info, applicationName)
}
