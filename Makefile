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

clean: lint
	rm -rf tmp

grs: #Use air for hot-reload
	go run src/*.go

lint:
	golangci-lint run ./src

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
	cd test/all_src_tests && go test

test: test_tasks_crud test_jwt test_cors test_src
	cd test && hurl --test entry.hurl
	echo "Complete"