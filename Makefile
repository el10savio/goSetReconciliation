
SetReconciliation-build:
	@echo "Building SetReconciliation Docker Image"	
	docker build -t set -f Dockerfile .

SetReconciliation-run:
	@echo "Running Single SetReconciliation Docker Container"
	docker run -p 8080:8080 -d set

provision:
	@echo "Provisioning Set Cluster"	
	bash scripts/provision.sh 2

e2e:
	@echo "Running E2E Testing On Set Cluster"	
	bash scripts/tests.sh

clean:
	@echo "Cleaning Set Cluster"
	bash scripts/teardown.sh

info:
	echo "Set Cluster Nodes"
	docker ps | grep 'set'
	docker network ls | grep set_network

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

golanglintci:
	@echo "golanglintci SetReconciliation Server"	
	docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:v1.42.1 golangci-lint run --out-format tab --enable-all

semgrep:
	@echo "semgrep SetReconciliation Server"	
	docker run --rm -v "$(shell pwd):/src" returntocorp/semgrep --config=auto

lint-dockerfile:
	@echo "lint SetReconciliation Dockerfile"	
	docker run --rm -i hadolint/hadolint < Dockerfile

test:
	@echo "Testing SetReconciliation Server"	
	go test -v --cover ./...

shellcheck:
	@echo "shellcheck SetReconciliation Scripts"
	shellcheck scripts/*.sh

shfmt:
	@echo "shfmt SetReconciliation Scripts"
	shfmt -i 2 -ci -w -l -bn scripts/*.sh

codespell:
	@echo "checking SetReconciliation spellings"
	codespell
