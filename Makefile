ifeq ($(update),yes)
  u=-u
endif

.PHONY: devel-deps
devel-deps:
	go get ${u} github.com/mattn/goveralls
	go get ${u} golang.org/x/lint/golint
	go get ${u} github.com/x-motemen/gobump/cmd/gobump
	go get ${u} github.com/Songmu/ghch/cmd/ghch

.PHONY: test
test:
	go test -v -race -covermode=atomic -coverprofile=coverage.out ./typetalk/...

.PHONY: cover
cover: devel-deps
	goveralls -coverprofile=coverage.out -service=travis-ci

.PHONY: lint
lint: devel-deps
	go vet ./typetalk/...
	golint -set_exit_status ./typetalk/...

.PHONY: bump
bump: devel-deps
	./_tools/bump