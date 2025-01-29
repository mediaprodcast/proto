.PHONY: lint generate update-deps check

lint:
	buf lint

generate:
	buf generate

update-deps:
	buf dep update

# Optional: Generate documentation
docs:
	buf generate --docs-out ./docs

# Optional: Check for breaking changes
breaking:
	buf breaking

# Optional: Check for compatibility with a specific version
compatible:
	buf breaking check --against v1.2.3 # Replace with the desired version

clean:
	rm -rf gen

# Run all checks
check: lint generate update-deps breaking compatible
