.PHONY: reset_tags grs lint \
	test_jwt test_jwt_encode test_jwt_decode \
	test_tasks_crud \
	test

reset_tags:
	git tag -l | xargs git tag -d

grs:
	cd src && go run *.go

lint:
	cd src && golangci-lint run

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

test: test_tasks_crud test_jwt
	cd test && hurl --test entry.hurl
	echo "Complete"