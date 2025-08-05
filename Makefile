.PHONY: reset_tags update clean grs lint \
	test_jwt test_jwt_encode test_jwt_decode \
	test_cors \
	test_tasks_crud \
	test_src \
	test

reset_tags:
	git tag -l | xargs git tag -d

update:
	go get -t -u ./...

clean: lint update
	swag init -g src/*.go
	rm -rf tmp

grs: #No hot-reload, re-build swagger docs
	swag init -g src/*.go
	go run src/*.go

lint:
	golangci-lint run ./src
	golangci-lint run ./test/srctest

define run_test
	cd test/$(1) && hurl --test *.hurl
endef

test_tasks_crud:
	$(call run_test,tasks_crud)

test_jwt_encode:
	cd test/jwt_auth && hurl --test encode.hurl

test_jwt_decode:
	cd test/jwt_auth && sh test_decode.sh

test_jwt: test_jwt_encode test_jwt_decode

test_cors:
	$(call run_test,cors)

test_src:
	cd test/srctest && go test

test: test_tasks_crud test_jwt test_cors test_src
	cd test && hurl --test entry.hurl
	echo "Complete"