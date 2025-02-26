default:
	@echo "Choose a target"
	@echo "  clean"
	@echo "  build-docker (docker image)"
	@echo "  run-docker (docker image)"

clean:
	@rm -rf *~

build-docker:
	@docker build -t go-web-test -f Dockerfile .

run-docker: build-docker
	@docker run --rm --name go-web-test -p 9991:9991 go-web-test

.PHONY: clean build-docker run-docker
