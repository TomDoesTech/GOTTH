# A Modified GoTTH Stack
[![Forked from this GoTTH stack walkthrough](https://img.youtube.com/vi/k00jVJeZxrs/0.jpg)](https://www.youtube.com/watch?v=k00jVJeZxrs)

## Technologies
### Go
Backend

### Tailwind
We are using **NPM** to generate our `css` using tailwind.
```bash
npm install
```
~~To get started with TailWind CSS, make sure you have the correct binary in the root directory.
https://tailwindcss.com/blog/standalone-cli~~

### Templ
Templating
```bash
go install github.com/a-h/templ/cmd/templ@latest
```

Docs: https://templ.guide/

### HTMX
Interactivity

### Postgres
Database

### SQLC
Object Relational Mapping
```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

### Chi
Server Routing

Docs: https://docs.sqlc.dev/en/stable/index.html

## Getting started

The tools we use that are required for development:
* [**make**](https://www.gnu.org/software/make/manual/make.html) - utility
  * install for Windows: `choco install make`
  * install for MacOS: `brew install make`
  * _no need for explanation to linux users_
* [**air**](https://github.com/cosmtrek/air?tab=readme-ov-file#installation) - for development live-reloading

```
git clone https://github.com/markkhoo/GOTTH.git
cd GOTTH
make dev
```

### Environment Variables
Using the [**Go** port of DotEnv](https://github.com/joho/godotenv). Example `.env` file for development:
```bash
DB_HOST="127.0.0.1"
DB_PORT="5432"
DB_DATABASE="sqlc-example"
DB_USERNAME="postgres"
DB_PASSWORD="password"
```

## Makefile
This Makefile is designed to simplify common development tasks for your project. It includes targets for building your Go application, watching and building Tailwind CSS, generating templates, and running your development server using Air.

### Targets:
This target watches the ./static/css/input.css file and automatically rebuilds the Tailwind CSS styles whenever changes are detected.
```bash
make tailwind-watch
```

This target minifies the Tailwind CSS styles by running the tailwindcss command.
```bash
make tailwind-build
```

This target generates templates using the templ command.
```bash
make templ-generate
```

This target generates database controllers in go using the sqlc command.
```bash
make sqlc-generate
```

This target runs the development server using Air, which helps in hot-reloading your Go application during development.
```bash
make dev
```

This target orchestrates the building process by executing the tailwind-build, templ-generate, sqlc-generate, and go build commands sequentially. It creates the binary output in the ./bin/ directory.
```bash
make build
```
