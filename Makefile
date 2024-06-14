BUILDER = ./bin/bdjuno-builder

.PHONY: znet-start
znet-start:
	$(BUILDER) znet start --profiles=monitoring

.PHONY: znet-remove
znet-remove:
	$(BUILDER) znet remove

.PHONY: test
test:
	docker run --name bdjuno-test-db -e POSTGRES_USER=bdjuno -e POSTGRES_PASSWORD=password -e POSTGRES_DB=bdjuno -d -p 6433:5432 postgres
	$(BUILDER) test
	docker stop bdjuno-test-db
	docker rm bdjuno-test-db

.PHONY: build
build:
	$(BUILDER) build

.PHONY: images
images:
	$(BUILDER) images

.PHONY: generate
generate:
	$(BUILDER) generate

.PHONY: dependencies
dependencies:
	$(BUILDER) download
