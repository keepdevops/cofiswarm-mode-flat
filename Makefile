ROLE := mode-flat
.PHONY: build test test-standalone-layout test-mode-config-gate
build:
	go build -o bin/cofiswarm-mode-flat ./cmd/cofiswarm-mode-flat
test: build test-standalone-layout test-mode-config-gate
test-standalone-layout:
	./test/scripts/assert-layout.sh $(ROLE)
test-mode-config-gate:
	./test/scripts/test-mode-config-gate.sh $(ROLE)
