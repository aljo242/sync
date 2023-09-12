#!/usr/bin/make -f
DOCKER := $(shell which docker)

###############################################################################
###                                Protobuf                                 ###
###############################################################################

BUF_VERSION=1.21.0

proto-all: proto-gen

proto-gen: proto-format
	@echo "🤖 Generating code from protobuf..."
	@$(DOCKER) run --rm --volume "$(PWD)":/workspace --workdir /workspace \
		sync-proto sh ./proto/generate.sh
	@echo "✅ Completed code generation!"

proto-lint:
	@echo "🤖 Running protobuf linter..."
	@$(DOCKER) run --volume "$(PWD)":/workspace --workdir /workspace \
		bufbuild/buf:$(BUF_VERSION) lint
	@echo "✅ Completed protobuf linting!"

proto-format:
	@echo "🤖 Running protobuf format..."
	@$(DOCKER) run --volume "$(PWD)":/workspace --workdir /workspace \
		bufbuild/buf:$(BUF_VERSION) format -w
	@echo "✅ Completed protobuf format!"

proto-breaking-check:
	@echo "🤖 Running protobuf breaking check against develop branch..."
	@$(DOCKER) run --volume "$(PWD)":/workspace --workdir /workspace \
		bufbuild/buf:$(BUF_VERSION) breaking --against '.git#branch=develop'
	@echo "✅ Completed protobuf breaking check!"

proto-setup:
	@echo "🤖 Setting up protobuf environment..."
	@$(DOCKER) build --rm --tag sync-proto:latest --file proto/Dockerfile .
	@echo "✅ Setup protobuf environment!"