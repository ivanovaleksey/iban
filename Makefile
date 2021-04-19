SWAGGER_VERSION = 2.2.10
API_VERSION ?= v0.1

.PHONY: build
build:
	go build -o bin/api ./cmd/

.PHONY: build-docker
build-docker:
	docker build -t ramone/iban:$(API_VERSION) .

.PHONY: test
test:
	go test -v -count=1 ./...

.PHONY: docs
docs:
	mkdir -p docs

.PHONY: swagger
swagger: docs
	curl -s https://codeload.github.com/swagger-api/swagger-ui/tar.gz/v$(SWAGGER_VERSION) | tar xzv -C docs swagger-ui-$(SWAGGER_VERSION)/dist
	mv -f docs/swagger-ui-$(SWAGGER_VERSION)/dist/* docs/
	rm -rf docs/swagger-ui-$(SWAGGER_VERSION) docs/swagger-ui.js
	sed -i_ "s/swagger-ui\.js/swagger-ui\.min\.js/" docs/index.html
	sed -i_ "s/http:\/\/petstore\.swagger\.io\/v2\///" docs/index.html
	rm -f docs/*_

.PHONY: clean
clean:
	rm -rf docs

.PHONY: generate
generate:
	go generate ./...
