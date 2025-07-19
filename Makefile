.PHONY: reset_tags

reset_tags:
	git tag -l | xargs git tag -d