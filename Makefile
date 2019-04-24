all: install
install: go.sum
	GO111MODULE=on go install -tags "$(build_tags)" ./x/nameshake/cmd/nsd
	GO111MODULE=on go install -tags "$(build_tags)" ./x/nameshake/cmd/nscli

go.sum: go.mod
	@echo "--> Ensure dependencides have not been modified"
	GO111MODULE=on go mod verify
