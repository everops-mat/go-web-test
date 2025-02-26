default:
	@echo "Choose a target"
	@echo "  clean"
	@echo "  build-docker (build docker image)"
	@echo "  run-docker (run docker image)"
	@echo "  pre-commit (install pre-commit)"
	@echo "  pre-commit-update (update pre-commit)"
	@echo "  run-pre-commit (run pre-commit on all files"

clean:
	@rm -rf *~

build-docker:
	@docker build -t go-web-test -f Dockerfile .

run-docker: build-docker
	@docker run --rm --name go-web-test -p 9991:9991 go-web-test

pre-commit:
	@pre-commit install

pre-commit-update:
	@pre-commit autoupdate

run-pre-commit:
	@pre-commit run --all-files

.PHONY: clean build-docker run-docker pre-commit pre-commit-update
.PHONY: run-pre-commit
