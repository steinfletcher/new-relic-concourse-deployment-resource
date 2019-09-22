package http

import (
	"github.com/steinfletcher/apitest"
	"github.com/steinfletcher/new-relic-concourse-deployment-resource/domain"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestWriteDeployment_WithName(t *testing.T) {
	defer apitest.NewStandaloneMocks(
		apitest.NewMock().
			Get("https://api.newrelic.com/v2/applications.json").
			Header("X-Api-Key", "12345").
			Query("filter[name]", "dave-app").
			RespondWith().
			Status(http.StatusOK).
			Body(`{"applications": [{"id": 15234241}]}`).
			End(),
		apitest.NewMock().
			Post("https://api.newrelic.com/v2/applications/15234241/deployments.json").
			Header("X-Api-Key", "12345").
			Body(`{"deployment":{"revision":"1","description":"changes","user":"dave"}}`).
			RespondWith().
			Status(http.StatusOK).
			End(),
	).End()()

	client := NewNewRelicClient("12345")
	info := domain.DeploymentInfo{
		Revision:    "1",
		Description: "changes",
		User:        "dave",
	}

	err := client.WriteWithName(info, "dave-app")

	assert.NoError(t, err)
}

func TestWriteDeployment_WithID(t *testing.T) {
	defer apitest.NewStandaloneMocks(
		apitest.NewMock().
			Post("https://api.newrelic.com/v2/applications/15234241/deployments.json").
			Header("X-Api-Key", "12345").
			Body(`{"deployment":{"revision":"1","description":"changes","user":"dave"}}`).
			RespondWith().
			Status(http.StatusOK).
			End(),
	).End()()

	client := NewNewRelicClient("12345")
	info := domain.DeploymentInfo{
		Revision:    "1",
		Description: "changes",
		User:        "dave",
	}

	err := client.Write(info, "15234241")

	assert.NoError(t, err)
}
