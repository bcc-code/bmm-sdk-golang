# Determine the current version from the latest git tag
CURRENT_VERSION=$(shell git describe --tags --abbrev=0)
# Extract the major, minor, and patch components
CURRENT_MAJOR=$(shell echo $(CURRENT_VERSION) | cut -d. -f1 | cut -dv -f2)
CURRENT_MINOR=$(shell echo $(CURRENT_VERSION) | cut -d. -f2)
CURRENT_PATCH=$(shell echo $(CURRENT_VERSION) | cut -d. -f3)

# Compute the next patch version
NEXT_PATCH=$(shell echo $$(($(CURRENT_PATCH) + 1)))
NEXT_VERSION=v$(CURRENT_MAJOR).$(CURRENT_MINOR).$(NEXT_PATCH)

# Help command
help:
	@echo "Makefile for tagging with the next patch version"
	@echo ""
	@echo "Usage:"
	@echo "  make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  help    Show this help message"
	@echo "  tag     Tag the repository with the next patch version"


tag:
	@echo "Current version: $(CURRENT_VERSION)"
	@echo "Next version: $(NEXT_VERSION)"
	@git tag -a $(NEXT_VERSION) -m "Release $(NEXT_VERSION)"

.PHONY: tag help
