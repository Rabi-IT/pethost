setup_test:
	go run cmd/test_setup.go

generate_mock:
	mockery --all --output fixtures/mocks