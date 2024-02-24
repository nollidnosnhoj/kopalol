.PHONY: tailwind-build
tailwind-build:
	@npx tailwindcss -i ./assets/tailwind.css -o ./assets/dist/css/tailwind.css

.PHONY: tailwind-watch
tailwind-watch:
	@npx tailwindcss -i ./assets/tailwind.css -o ./assets/dist/css/tailwind.css --watch

.PHONY: templ-generate
templ-generate:
	@templ generate

.PHONY: dev
dev:
	@air -c air.toml

.PHONY: build
build:
	@make tailwind-build
	@make templ-generate
	@go build -o ./bin/simplimg ./cmd/main.go