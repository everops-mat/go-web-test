SERVER_CODE := cmd/server/main.go internal/logger/logger.go \
internal/signals/signals.go internal/sayings/eo_sayings.go \
internal/handlers/handlers.go internal/handlers/auth_middleware.go

default:
	@echo "Choose a target"
	@echo "  clean"
	@echo "  build-docker (build docker image)"
	@echo "  run-docker (run docker image)"
	@echo "  pre-commit (install pre-commit)"
	@echo "  pre-commit-update (update pre-commit)"
	@echo "  run-pre-commit (run pre-commit on all files"
	@echo "  go-web-test (build go-web-test locally)"

clean:
	@rm -rf *~ go-web-test
	@find . -type f -not -path './.git/*' -name '*~' | xargs rm -f

build-docker:
	@docker build -t go-web-test -f Dockerfile .

run-docker: build-docker
	@docker run --rm --name go-web-test -p 9991:8080 go-web-test

pre-commit:
	@pre-commit install

pre-commit-update: pre-commit
	@pre-commit autoupdate

run-pre-commit: pre-commit
	@pre-commit run --all-files

go-web-test: $(SERVER_CODE)
	@go build -o go-web-test ./cmd/server

.PHONY: clean build-docker run-docker pre-commit pre-commit-update
.PHONY: run-pre-commit
