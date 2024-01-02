VERSION ?= $(patsubst v%,%,$(shell git describe))


all: bin/pulumi-resource-frontegg  bin/pulumi-tfgen-frontegg schema python-sdk

bin/pulumi-resource-frontegg: cmd/pulumi-resource-frontegg/schema.json
	go build -o bin/pulumi-resource-frontegg ./cmd/pulumi-resource-frontegg

bin/pulumi-tfgen-frontegg: cmd/pulumi-tfgen-frontegg/*.go go.sum
	go build -o bin/pulumi-tfgen-frontegg ./cmd/pulumi-tfgen-frontegg

cmd/pulumi-resource-frontegg/schema.json: bin/pulumi-tfgen-frontegg
	bin/pulumi-tfgen-frontegg $(VERSION) schema --out ./cmd/pulumi-resource-frontegg

schema: cmd/pulumi-resource-frontegg/schema.json

python-sdk: bin/pulumi-tfgen-frontegg
	rm -rf sdk
	bin/pulumi-tfgen-frontegg $(VERSION) python
	cp README.md sdk/python/
	cd sdk/python/ && \
		sed -i.bak -e "s/VERSION = .*/VERSION = '$(VERSION)'/g" -e "s/PLUGIN_VERSION = .*/PLUGIN_VERSION = '$(VERSION)'/g" setup.py && \
		rm setup.py.bak
