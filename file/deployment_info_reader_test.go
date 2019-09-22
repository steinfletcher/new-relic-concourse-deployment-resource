package file_test

import (
	"github.com/steinfletcher/new-relic-concourse-deployment-resource/domain"
	"github.com/steinfletcher/new-relic-concourse-deployment-resource/file"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReader(t *testing.T) {
	tests := map[string]struct {
		path         string
		expectedErr  string
		expectedInfo domain.DeploymentInfo
	}{
		"reads deployment info": {
			path: "testdata/deployment_info.json",
			expectedInfo: domain.DeploymentInfo{
				Revision:    "1",
				Description: "changes",
				User:        "dave",
			},
		},
		"deployment info file does not exist": {
			path:         "testdata/dontexist.json",
			expectedErr:  "no such file or directory",
			expectedInfo: domain.DeploymentInfo{},
		},
		"json decoding error": {
			path:         "testdata/not_json.txt",
			expectedErr:  "invalid character",
			expectedInfo: domain.DeploymentInfo{},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			reader := file.NewDeploymentInfoReader(test.path)

			info, err := reader.Read()

			if test.expectedErr != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), test.expectedErr)
			}
			assert.Equal(t, test.expectedInfo, info)
		})
	}
}
