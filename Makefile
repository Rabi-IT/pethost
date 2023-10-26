server:
	go run cmd/server/main.go

setup_test:
	go run cmd/test_setup/test_setup.go

test:
	go clean -testcache
	go run gotest.tools/gotestsum@latest --format testname -- ./... -p 1 -v

generate_mock:
	mockery --all --output fixtures/mocks