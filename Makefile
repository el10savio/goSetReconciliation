
SetReconciliation-build:
	@echo "Building SetReconciliation Docker Image"	
	docker build -t setreconciliation -f Dockerfile .

SetReconciliation-run:
	@echo "Running Single SetReconciliation Docker Container"
	docker run -p 8080:8080 -d setreconciliation

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
	@echo "Testing SetReconciliation"	
	go test -v --cover ./...
