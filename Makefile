.PHONY: reset_tags grs test lint

reset_tags:
	git tag -l | xargs git tag -d

grs:
	cd src && go run *.go

test: #Run all hurl tests
	cd src/test && \
	hurl --test *.hurl

lint:
	cd src && golangci-lint run