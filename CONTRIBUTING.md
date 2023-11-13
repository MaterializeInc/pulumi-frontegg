# Contributing

## Developing the provider

The Pulumi provider is dependent on the [Frontegg Terraform provider](https://github.com/frontegg/terraform-provider-frontegg/releases) Any new features or changes should be applied there.

### Updating Terraform provider

To update the Terraform provider, 
1. change the version of `github.com/frontegg/terraform-provider-frontegg`in the [provider/go.mod](provider/go.mod) 
2. cd to `provider/` and run `go mod tidy`.
3. validate local build `make build`
4. create a pr titled `bump frontegg tf provider (<version>)`

**validating with an existing pulumi project**
1. `make build`
1. `cp bin/pulumi-resource-frontegg $GOPATH/bin`
1. in a pulumi project `. venv/bin/activate`
1. `. venv/bin/activate` 
1. `venv/bin/pip install -e pulumi-frontegg/sdk/python`
1. `pulumi up`

### Updating the schema

To update the schema for the Pulumi provider (after the Terraform version has been updated) run `make tfgen` to automatically generate any new schema or property changes.

If a new resource or data source has been added, the reference will need to be added to [provider/resources.go]. After it has been added, you can run `make tfgen`.

## Cutting a release

1. To
To cut a new release of the provider, create a new tag and push that tag. This will trigger a Github Action to generate the artifacts and build the configured SDKs.

```bash
git fetch origin
git checkout master
git reset origin/master
git tag vX.Y.Z
git push origin vX.Y.Z
```


**Note choosing a version:**
  - in general just bump the version to the next available patch version
  - if the upstream has changed a major or minor version bump that here
  - if there are breaking changes bump a major version here

