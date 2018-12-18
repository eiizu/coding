
.PHONY: build
build: ## Build
	@echo "::: building app"
	go build -o ./bin/app
	@echo "::: build done"

.PHONY: run
run: build ## Run application 
	@echo "::: running app"
	./bin/app