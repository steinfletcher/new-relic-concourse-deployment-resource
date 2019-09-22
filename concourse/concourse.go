package concourse

import "errors"

// A Source is the resource's source configuration.
type Source struct {
	NewRelicAccount string `json:"new_relic_account"`
	NewRelicApiKey  string `json:"new_relic_api_key"`
}

// Metadata are a key-value pair that must be included for in the in and out
// operation responses.
type Metadata struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

// CheckResponse is the output for the check operation.
type CheckResponse []Version

// InResponse is the output for the in operation.
type InResponse struct {
	Version  Version    `json:"version"`
	Metadata []Metadata `json:"metadata"`
}

// OutParams are the parameters that can be configured for the out operation.
type OutParams struct {
	DeploymentInfoFilePath  string `json:"deployment_info_file_path"`
	RepoPath                string `json:"repo_path"`
	NewRelicApplicationName string `json:"new_relic_application_name"`
	NewRelicApplicationID   string `json:"new_relic_application_id"`
}

// OutRequest is in the input for the out operation.
type OutRequest struct {
	Source Source    `json:"source"`
	Params OutParams `json:"params"`
}

// Validate returns an error if any required source parameters are not set
func (s OutRequest) Validate() error {
	if s.Source.NewRelicAccount == "" {
		return errors.New("'new_relic_account' is not set")
	}
	if s.Source.NewRelicApiKey == "" {
		return errors.New("'new_relic_api_key' is not set")
	}
	return nil
}

// OutResponse is the output for the out operation.
type OutResponse struct {
	Version  Version    `json:"version"`
	Metadata []Metadata `json:"metadata"`
}

// Version is the key-value pair that the resource is checking, getting or putting.
type Version map[string]string
