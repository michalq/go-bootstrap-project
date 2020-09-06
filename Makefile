.PHONY: clean

clean:
	rm -rf ./pkg/api

pkg/api:
	./scripts/openapi-to-go.sh