# new-relic-concourse-deployment-resource

## Resource Configuration

```yaml
resource_types:
- name: new-relic-deployment
  type: docker-image
  source:
    repository: steinfletcher/new-relic-concourse-deployment-resource
    tag: latest

resources:
  - name: deployment-marker
    type: new-relic-deployment
    source:
      new_relic_account: "account id"
      new_relic_api_key: "api key"
```

### `check`: no operation

### `in`: no operation

### `out`: Creates a new relic deployment marker

#### Parameters

Required:

One of the following:
-	`new_relic_application_name`: The name of the application in New Relic (must be unique). This is used to resolve an application id.
-	`new_relic_application_id`: The id of the application. This takes precedence over `new_relic_application_name`

Optional:

-	`text_file`: File that contains the deployment information in JSON format. This is used when the default git output is not sufficient or when custom behaviour is required. Leave this empty to use the default commit info, which produces the following output
```
{
  "revision": "v0.1.1", // git tag or sha1 if commit not tagged
  "description": "Made some changes", // git commit message
  "user": "a@b.com" // git commit author
}
```

-	`repo_path`: The path of the source code

#### Example

```yaml
jobs:
  - name: job
    public: true
    plan:
      - task: simple-task
        config:
          platform: linux
          image_resource:
            type: registry-image
            source:
              repository: busybox
          run:
            path: echo
            args: ["Hello, world!"]
      - get: repo
      - put: deployment-marker
        params:
          repo_path: repo
          new_relic_application_name: my-login-app
```