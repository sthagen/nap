.DEFAULT_GOAL := all
package = nap
version = 2023.9.20

.PHONY: format
format:
	@go fmt

.PHONY: lint
lint:
	@go vet

.PHONY: build
build:
	@echo ... building for linux/amd64, local, and windows/amd64
	@env GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o nap
	@go build -ldflags "-s -w" -o nap-local
	@env GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o nap.exe
	@echo ... smoke test of local app
	@./nap-local v
	@./nap-local h

.PHONY: man
man:
	@mkdir -p share/man/man1
	@pandoc docs/nap.1.md -s -t man -o share/man/man1/nap.1
	@echo ... man share/man/man1/nap.1

.phony: spellcheck
spellcheck:
	@vale --config etc/vale.ini README.md main.go docs/nap.1.md

.PHONY: test
test:
	@go test -cover -coverprofile=coverage.out

.PHONY: testcov
testcov: test
	@go tool cover -func=coverage.out
	@go tool cover -html=coverage.out -o coverage.html
	@echo ... open coverage.html

.PHONY: all
all: format lint build man testcov

.PHONY: release
release: all
	@echo ... collecting release info
	@make name > PUBLICATIO.tmp
	@printf "\n" >> PUBLICATIO.tmp
	@printf "Changes:\n\n" >> PUBLICATIO.tmp
	@python bin/gen_tag_changes.py >> PUBLICATIO.tmp
	@printf "\n" >> PUBLICATIO.tmp
	@printf "\n" >> PUBLICATIO.tmp
	@printf "Fingerprints:\n\n" >> PUBLICATIO.tmp
	@bin/gen_fingerprints.sh nap >> PUBLICATIO.tmp
	@printf "\n" >> PUBLICATIO.tmp
	@bin/gen_fingerprints.sh nap.exe >> PUBLICATIO.tmp
	@printf "\n" >> PUBLICATIO.tmp
	@cat PUBLICATIO.tmp

.PHONY: tag
tag: release
	@echo ... creating and pushing annotated tag
	@git tag -a "v$(version)" -F PUBLICATIO.tmp -s
	@git push all --tags

.PHONY: clean
clean:
	@rm -f coverage.{html,out} nap{'',-local,.exe} share/man/man1/nap.1 PUBLICATIO.tmp

.PHONY: name
name:
	@printf "Release '%s'\n\n" "$$(git-release-name "$$(git rev-parse HEAD)")"
	@printf "%s revision.is(): sha1:%s\n" "-" "$$(git rev-parse HEAD)"
	@printf "%s name.derive(): '%s'\n" "-" "$$(git-release-name "$$(git rev-parse HEAD)")"
	@printf "%s node.id(): '%s'\n" "-" "$$(bin/gen_node_identifier.py)"
