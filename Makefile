.PHONY: run codegen

codegen:
	rm -rf mocks || true
	mockgen -source=internal/service/service.go -destination=mocks/services_mock.go -package=mocks
	mockgen -source=internal/repository/repository.go -destination=mocks/repo_mock.go -package=mocks

run:
	docker-compose up --build
