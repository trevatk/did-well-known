
deps:
	go mod tidy
	go mod vendor

build:
	docker build -t trevatk/did-well-known:latest .