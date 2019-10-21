packages = \
	./gql \

.PHONY: test
test:
	@$(foreach package,$(packages), \
		set -e; \
		go test -coverprofile $(package)/cover.out -covermode=count $(package);)

test-clear:
	@echo "Cleaning test output..."
	@$(foreach package,$(packages), \
		rm -rf *.out; \
		rm -rf *.xml; \
		rm -rf ./**/*.out; \
		rm -rf ./**/*.xml;)

.PHONY: cover
cover: test
	@echo "mode: count" > cover-all.out
	@$(foreach package,$(packages), \
		tail -n +2 $(package)/cover.out >> cover-all.out;)
	@gocover-cobertura < cover-all.out > cover-cobertura.xml

.PHONY: show
show:
	@echo "Launching web browser to show overall coverage..."
	@go tool cover -html=cover-all.out
