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
	    /bin/echo -n "'Templ' is not installed on your machine. Do you want to install it? [Y/n] "; \
			read choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
					GO_VERSION=$$(go version | awk '{print $$3}' | tr -d 'go'); \
					MIN_GO_VERSION=1.20; \
					if [ "$$(printf '%s\n' "$$MIN_GO_VERSION" "$$GO_VERSION" | sort -V | head -n1)" = "$$MIN_GO_VERSION" ]; then \
							echo "Installing templ..."; \
	        		go install github.com/a-h/templ/cmd/templ@latest; \
	        		templ generate; \
					else \
	        		echo "templ requires golang version 1.20 or greater. Exiting..."; \
	        		exit 1; \
					fi
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
			/bin/echo -n "'Sqlc' is not installed on your machine. Do you want to install it? [Y/n] "; \
			read choice; \
			if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
					echo "Installing sqlc..."; \
					GO_VERSION=$$(go version | awk '{print $$3}' | tr -d 'go'); \
					MIN_GO_VERSION=1.17; \
					if [ "$$(printf '%s\n' "$$MIN_GO_VERSION" "$$GO_VERSION" | sort -V | head -n1)" = "$$MIN_GO_VERSION" ]; then \
							go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest; \
							sqlc generate; \
					else \
							go get github.com/sqlc-dev/sqlc/cmd/sqlc; \
							sqlc generate; \
					fi \
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
	    /bin/echo -n "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] "; \
			read choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
					echo "Installing air..."; \
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