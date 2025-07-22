.PHONY: reset_tags grs lint \
	test_tasks_crud

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

test: test_tasks_crud #Run all hurl tests
	echo "Complete"