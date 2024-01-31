.PHONY: tailwind-build
tailwind-watch:
	./tailwindcss -i ./static/css/input.css -o ./static/css/style.css --watch

.PHONY: tailwind-build
 tailwind-build:
	./tailwindcss -i ./static/css/input.css -o ./static/css/style.css --minify