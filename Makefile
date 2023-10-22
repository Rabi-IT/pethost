server:
	go run cmd/server/main.go

setup_test:
	go run cmd/test_setup/test_setup.go

generate_mock:
	mockery --all --output fixtures/mocks