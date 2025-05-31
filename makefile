run:
	go run cmd/api/main.go

dev:
	go run cmd/api/main.go --logLevel=debug -bePersistent

build:
	go build cmd/api/main.go
	
enrich:
	go test -count=1 cmd/test/enricher_test.go