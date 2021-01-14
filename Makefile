PROJECT_ROOT:= ${PWD}
ESCAPE_FLAGS := "-gcflags -m"
RACE_FLAGS := "-race"
MAIN_PKG := "./cmd"

.PHONY: escape_analysis
escape_analysis:	## Run escape analysis
	$(info Generating compiler escape analysis report:)
	@go run "$(ESCAPE_FLAGS)" "$(MAIN_PKG)"


.PHONY: data_race
data_race:	## Run escape analysis
	$(info Generating data race detection report:)
	@go run "$(RACE_FLAGS)" "$(MAIN_PKG)"
