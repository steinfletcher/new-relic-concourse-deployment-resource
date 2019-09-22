package git

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	reader := NewDeploymentInfoReader("..")

	info, err := reader.Read()

	assert.NoError(t, err)
	assert.NotEmpty(t, info.Revision)
	assert.NotEmpty(t, info.Description)
	assert.NotEmpty(t, info.User)
}
