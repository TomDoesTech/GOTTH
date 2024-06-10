IMAGE_ID ?= app
NETWORK ?= app

.PHONY: tailwind-watch
tailwind-watch:
	./tailwindcss -i ./static/css/input.css -o ./static/css/style.css --watch

.PHONY: tailwind-build
tailwind-build:
	./tailwindcss -i ./static/css/input.css -o ./static/css/style.min.css --minify

.PHONY: templ-generate
templ-generate:
	templ generate

.PHONY: templ-watch
templ-watch:
	templ generate --watch
	
.PHONY: dev
dev:
	go build -o ./tmp/$(APP_NAME) ./cmd/$(APP_NAME)/main.go && air

.PHONY: start
start:
	go run cmd/*.go

.PHONY: create-network
create-network:
	docker network ls | grep -q -w $(NETWORK) || docker network create -d bridge $(NETWORK)

.PHONY: build-image
build-image:
	make create-network
	docker images | grep $(NETWORK) -q  && [ "${FORCE}" != "true" ] \
		&& echo "\nImage exists! Skipping... (use FORCE=true to force the image to rebuild)" \
		|| docker build --target development -f ./Dockerfile . -t $(NETWORK)

.PHONY: build
build:
	make tailwind-build
	make templ-generate
	go build -ldflags "-X main.Environment=production" -o ./bin/$(APP_NAME) ./cmd/$(APP_NAME)/main.go

.PHONY: start
start:
	go run cmd/*.go

.PHONY: down
down:
	docker compose -p $(NETWORK) down

.PHONY: up
up:
	docker compose -p $(NETWORK) up

.PHONY: vet
vet:
	go vet ./...

.PHONY: staticcheck
staticcheck:
	staticcheck ./...

.PHONY: test
test:
	  go test -race -v -timeout 30s ./...