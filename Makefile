.PHONY: tailwind-build
tailwind-build:
	@echo "Building tailwind.css"
	@npx tailwindcss -i ./assets/tailwind.css -o ./assets/dist/css/tailwind.css

.PHONY: tailwind-watch
tailwind-watch:
	@echo "Watching tailwind.css"
	@npx tailwindcss -i ./assets/tailwind.css -o ./assets/dist/css/tailwind.css --watch

.PHONY: templ-generate
templ-generate:
	@echo "Generating templ templates"
	@templ generate

.PHONY: dev
dev:
	@air -c air.toml

.PHONY: esbuild
esbuild:
    @echo "Building esbuild"
    @node esbuild.mjs

.PHONY: build
build:
    @make esbuild
	@make tailwind-build
	@make templ-generate
	@echo "Building go binary"
	@go build -o ./bin/kopalol ./main.go

.PHONY: build-lite
build-lite:
    @make templ-generate
    @echo "Building go binary"
	@go build -o ./bin/kopalol ./main.go

.PHONY: clean
clean:
	@echo "Clean up bin, node_modules, tmp folder"
	@go clean ./...
	@rm -rf ./bin ./node_modules ./tmp
	@echo "Clean up uploads folder"
	@rm -rf ./uploads

.PHONY: install
install:
	@echo "Install node_modules and go modules"
	@npm install
	@go mod download
	@echo "Create uploads folder"
	@mkdir -p ./uploads && chmod 777 ./uploads && touch ./uploads/.keep
