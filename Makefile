# Generate statik/ from assets/
.PHONY: gen
gen:
	statik -src ./assets

.PHONY: test1
test1: gen
	go build -o bin/a.out \
	&& ./bin/a.out
