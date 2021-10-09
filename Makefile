
SetReconciliation-build:
	@echo "Building SetReconciliation Docker Image"	
	docker build -t set -f Dockerfile .

SetReconciliation-run:
	@echo "Running Single SetReconciliation Docker Container"
	docker run -p 8080:8080 -d set

build:
	@echo "Building SetReconciliation Server"	
	go build -o bin/SetReconciliation main.go

fmt:
	@echo "go fmt SetReconciliation Server"	
	go fmt ./...

vet:
	@echo "go vet SetReconciliation Server"	
	go vet ./...

lint:
	@echo "go lint SetReconciliation Server"	
	golint ./...

test:
	@echo "Testing SetReconciliation Server"	
	go test -v --cover ./...

shellcheck:
	@echo "shellcheck SetReconciliation Scripts"
	shellcheck scripts/*.sh
