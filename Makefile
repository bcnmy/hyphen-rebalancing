IMAGE=arbitrage

.PHONY: help
help: # Display this help
	@awk 'BEGIN{FS=":.*#";printf "Usage:\n  make <target>\n\nTargets:\n"}/^[a-zA-Z_-]+:.*?#/{printf"  %-10s %s\n",$$1,$$2}' $(MAKEFILE_LIST)

.PHONY: build
build: # Build docker image
	docker build -t $(IMAGE) .

.PHONY: run
run: build # Run bot
	docker run -v "$$(pwd)/config.yaml":/arbitrage/config.yaml $(IMAGE)
