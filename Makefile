
SetReconciliation-build:
	@echo "Building SetReconciliation Docker Image"	
	docker build -t set -f Dockerfile .

SetReconciliation-run:
	@echo "Running Single SetReconciliation Docker Container"
	docker run -p 8080:8080 -d set

provision:
	@echo "Provisioning Set Cluster"	
	bash scripts/provision.sh

e2e:
	@echo "Running E2E Testing On Set Cluster"	
	bash scripts/tests.sh

info:
	echo "Set Cluster Nodes"
	docker ps | grep 'set'
	docker network ls | grep set_network

clean:
	@echo "Cleaning Set Cluster"
	docker ps -a | awk '$$2 ~ /set/ {print $$1}' | xargs -I {} docker rm -f {}
	docker network rm set_network

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
