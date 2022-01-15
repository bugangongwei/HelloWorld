GO_FILES=`go list ./... | grep -v -E "mock"`

ci-test:
	go test $(GO_FILES) -v -race -coverprofile .testCoverage.txt
