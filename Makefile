.PHONY: test
.DEFAULT: test

test:
	@echo "$@"
	go test -v
