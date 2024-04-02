# GoTTH
A simple, modern stack for building fast web applications.

**Watch walk-through video:**

[![GoTTH stack walkthrough](https://img.youtube.com/vi/k00jVJeZxrs/0.jpg)](https://www.youtube.com/watch?v=k00jVJeZxrs)

* Go - Backend
* Tailwind - CSS
* Templ - Templating
* HTMX - Interactivity

## Technologies
### Tailwind
To generate the Tailwind style sheet, we use the Tailwind binary. To get started with TailWind CSS, make sure you have the correct binary in the root directory. follow the instructions in this guide. Make sure you download the correct binary for your operating system.
https://tailwindcss.com/blog/standalone-cli

### Templ
https://templ.guide/

### Air
Air is required for hot reloading used in ```make dev```
https://github.com/cosmtrek/air


## Makefile
This Makefile is designed to simplify common development tasks for your project. It includes targets for building your Go application, watching and building Tailwind CSS, generating templates, and running your development server using Air.

### Targets:
```bash
make tailwind-watch
```
This target watches the ./static/css/input.css file and automatically rebuilds the Tailwind CSS styles whenever changes are detected.

```
make tailwind-build
```
This target minifies the Tailwind CSS styles by running the tailwindcss command.

```
make templ-watch
```
This target watches for changes to *.templ files and automatically generates them.


```
make templ-generate
```
This target generates templates using the templ command.


```
make dev
```
This target runs the development server using Air, which helps in hot-reloading your Go application during development.

```
make build
```
This target orchestrates the building process by executing the tailwind-build, templ-generate, and go build commands sequentially. It creates the binary output in the ./bin/ directory.

