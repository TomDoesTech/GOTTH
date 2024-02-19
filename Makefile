.PHONY: tailwind-watch
tailwind-watch:
	./tailwindcss -i ./static/css/input.css -o ./static/css/style.css --watch

.PHONY: tailwind-build
tailwind-build:
	./tailwindcss -i ./static/css/input.css -o ./static/css/style.css --minify

.PHONY: templ-generate
templ-generate:
	@if command -v templ > /dev/null; then \
	    templ generate; \
	else \
	    read -p "'Templ' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/go-gof/tmpl/cmd/templ@latest; \
	        templ generate; \
	    else \
	        echo "You chose not to install templ. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

.PHONY: sqlc-generate
sqlc-generate:
	@if command -v sqlc > /dev/null; then \
	    sqlc generate; \
	else \
	    read -p "'Sqlc' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/kyleconroy/sqlc/cmd/sqlc@latest; \
	        sqlc generate; \
	    else \
	        echo "You chose not to install sqlc. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

.PHONY: dev
dev:
	@if command -v air > /dev/null; then \
	    go build -o ./tmp/$(APP_NAME) ./cmd/$(APP_NAME)/main.go && air; \
	    echo "Watching...";\
	else \
	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/cosmtrek/air@latest; \
	        go build -o ./tmp/$(APP_NAME) ./cmd/$(APP_NAME)/main.go && air; \
	        echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

.PHONY: build
build:
	make tailwind-build && make templ-generate && make sqlc-generate && go build -o ./bin/$(APP_NAME) ./cmd/$(APP_NAME)/main.go