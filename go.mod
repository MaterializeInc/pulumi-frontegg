module github.com/MaterializeInc/pulumi-frontegg

go 1.16

replace github.com/frontegg/terraform-provider-frontegg => github.com/frontegg/terraform-provider-frontegg v0.2.35

replace (
	github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0
	github.com/hashicorp/terraform-exec => github.com/hashicorp/terraform-exec v0.17.2
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20220725190814-23001ad6ec03
)

require (
	github.com/frontegg/terraform-provider-frontegg v0.0.0-00010101000000-000000000000
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.26.1
	github.com/pulumi/pulumi/sdk/v3 v3.36.0
	golang.org/x/mod v0.5.0
)
