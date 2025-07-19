.PHONY: reset_tags grs

reset_tags:
	git tag -l | xargs git tag -d

grs:
	cd src && go run *.go