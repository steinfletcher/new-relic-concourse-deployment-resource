package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/steinfletcher/new-relic-concourse-deployment-resource/domain"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func NewNewRelicClient(apiKey string) domain.DeploymentWriter {
	return &cli{
		apiKey:     apiKey,
		httpClient: &http.Client{Timeout: time.Second * 10},
	}
}

type recordDeployment struct {
	Deployment domain.DeploymentInfo `json:"deployment"`
}

type cli struct {
	apiKey     string
	httpClient *http.Client
}

type application struct {
	ID int `json:"id,omitempty"`
}

type getApplicationsResponse struct {
	Applications []application `json:"applications"`
}

func (r *cli) Write(deployment domain.DeploymentInfo, applicationID string) error {
	data, err := json.Marshal(recordDeployment{Deployment: deployment})
	if err != nil {
		panic(err)
	}
	res, err := r.doAuthenticatedRequest(http.MethodPost,
		fmt.Sprintf("https://api.newrelic.com/v2/applications/%s/deployments.json", applicationID),
		bytes.NewReader(data))
	if err != nil {
		return err
	}
	if res.StatusCode >= 400 {
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("failed to unmarshall new relic error response. %v", err)
		}
		return fmt.Errorf("received statusCode %d when creating a deployment. error: %s",
			res.StatusCode, string(data))
	}
	return nil
}

func (r *cli) WriteWithName(deployment domain.DeploymentInfo, applicationName string) error {
	id, err := r.getApplicationID(applicationName)
	if err != nil {
		return err
	}
	return r.Write(deployment, strconv.Itoa(id))
}

func (r *cli) getApplicationID(name string) (int, error) {
	req, err := http.NewRequest(http.MethodGet,
		"https://api.newrelic.com/v2/applications.json", nil)
	if err != nil {
		return -1, err
	}
	req.Header.Add("X-Api-Key", r.apiKey)

	q := req.URL.Query()
	q.Add("filter[name]", name)
	req.URL.RawQuery = q.Encode()

	res, err := r.httpClient.Do(req)
	if err != nil {
		return -1, err
	}
	if res.StatusCode >= 400 {
		return -1, fmt.Errorf("received statusCode: %d when fetching applicationID", res.StatusCode)
	}

	var apps getApplicationsResponse
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return -1, err
	}
	err = json.Unmarshal(data, &apps)
	if err != nil {
		return -1, err
	}
	if len(apps.Applications) == 0 {
		return -1, fmt.Errorf("could not find application with name=%s", name)
	}
	return apps.Applications[0].ID, nil
}

func (r *cli) doAuthenticatedRequest(method string, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Api-Key", r.apiKey)

	if body != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	return r.httpClient.Do(req)
}
