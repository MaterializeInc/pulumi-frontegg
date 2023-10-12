# Maintainer instructions

## Updating the tf provider
1. modify go.mod to update the provider

2. run `go get github.com/MaterializeInc/pulumi-frontegg/provider`

3. run `go mod tidy`

4. validate we can build locally `go build -o bin/pulumi-tfgen-frontegg ./cmd/pulumi-tfgen-frontegg`

5. create a pr `bump frontegg tf provider (<version>)`


## New Release

1. Make changes.

2. Commit changes and push to GitHub.

3. If changes introduced in PR, ensure that you tagging commit on main after merge.
   ```
   git fetch origin
   git checkout main
   git reset origin/main
   ````

3. Choose a version.
   - The major and minor version should match the upstream pulumi terraform provider version.
   - Find the newest version tag of the frontegg-pulumi provider.
   - If the major and minor version match of the new version match the existing tag,
     increment the patch version by one, other wise set the patch version to 1.


4. Create the tag.

    ```
    version=vA.B.C
    git tag -a $version -m $version
    git push
    git push origin $version
    ```

