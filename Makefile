.PHONY: clean

wire:
	cd ./cmd/api && wire

clean:
	rm -rf ./pkg/api

pkg/api:
	./scripts/openapi-to-go.sh