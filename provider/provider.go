// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package frontegg

import (
	"fmt"
	"unicode"

	tfversion "github.com/MaterializeInc/pulumi-frontegg"
	frontegg "github.com/frontegg/terraform-provider-frontegg/provider"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

const (
	mainPkg = "frontegg"
	mainMod = "index"
)

// makeMember manufactures a type token for the package and the given module and
// type.
func makeMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(mainPkg + ":" + mod + ":" + mem)
}

// makeType manufactures a type token for the package and the given module and
// type.
func makeType(mod string, typ string) tokens.Type {
	return tokens.Type(makeMember(mod, typ))
}

// makeDataSource manufactures a standard resource token given a module and
// resource name.  It automatically uses the main package and names the file by
// simply lower casing the data source's first character.
func makeDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeMember(mod+"/"+fn, res)
}

// makeResource manufactures a standard resource token given a module and
// resource name.  It automatically uses the main package and names the file by
// simply lower casing the resource's first character.
func makeResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeType(mod+"/"+fn, res)
}

// Provider returns additional overlaid schema and metadata associated with the
// provider.
func Provider(version string) tfbridge.ProviderInfo {
	p := shimv2.NewProvider(frontegg.New(tfversion.TFVersion)())

	prov := tfbridge.ProviderInfo{
		P:                 p,
		Name:              "frontegg",
		Description:       "A Pulumi package for creating and managing Frontegg resources.",
		Keywords:          []string{"pulumi", "frontegg"},
		License:           "Apache-2.0",
		Homepage:          "https://github.com/MaterializeInc/pulumi-frontegg",
		Repository:        "https://github.com/pulumi/pulumi-frontegg",
		PluginDownloadURL: fmt.Sprintf("https://github.com/MaterializeInc/pulumi-frontegg/releases/download/v%s/", version),

		Resources: map[string]*tfbridge.ResourceInfo{
			"frontegg_permission":          {Tok: makeResource(mainMod, "Permission")},
			"frontegg_permission_category": {Tok: makeResource(mainMod, "PermissionCategory")},
			"frontegg_role":                {Tok: makeResource(mainMod, "Role")},
			"frontegg_tenant":              {Tok: makeResource(mainMod, "Tenant")},
			"frontegg_user":                {Tok: makeResource(mainMod, "User")},
			"frontegg_webhook":             {Tok: makeResource(mainMod, "Webhook")},
			"frontegg_workspace":           {Tok: makeResource(mainMod, "Workspace")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"frontegg_permission": {Tok: makeDataSource(mainMod, "getPermission")},
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
