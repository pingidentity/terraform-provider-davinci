TEST?=$$(go list ./...)
SWEEP_DIR=./internal/sweep
DAVINCI_DIR=./internal/service/davinci
NAMESPACE=pingidentity
PKG_NAME=davinci
BINARY=terraform-provider-${NAME}
VERSION=0.4.8
OS_ARCH=linux_amd64

default: install

tools:
	@echo "==> Installing tools..."
	go generate -tags tools tools/main.go

fmtcheck:
	@echo "==> Formatting Terraform documentation examples with terraform fmt..."
	@terraform fmt -recursive ./examples/

build:
	@echo "==> Building..."
	go mod tidy
	go work vendor
	go build -v .

install: build
	@echo "==> Installing..."
	go install -ldflags="-X main.version=$(VERSION)"

generate: build generateconnectorref fmtcheck
	@echo "==> Generating code..."
	go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

generateconnectorref: build
	@echo "==> Generating connector docs & examples..."
	rm examples/connectors/*.tf || true
	go run github.com/pingidentity/terraform-provider-davinci/dvgenerate/cmd/generate
	
test: build
	@echo "==> Running tests..."
	go test $(TEST) $(TESTARGS) -timeout=5m

testacc: build
	@echo "==> Running acceptance tests..."
	TF_ACC=1 go test $$(go list ./internal/client/...) -v $(TESTARGS) -timeout 120m -parallel 15
	TF_ACC=1 go test $$(go list ./internal/service/...) -v $(TESTARGS) -timeout 120m -parallel 15

sweep: build
	@echo "==> Running sweep..."
	@echo "WARNING: This will destroy infrastructure. Use only in development accounts."
	go test $(SWEEP_DIR) -v -sweep=all $(SWEEPARGS) -timeout 10m

vet:
	@echo "==> Running go vet..."
	@go vet ./... ; if [ $$? -ne 0 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

docscategorycheck:
	@echo "==> Checking for missing category in generated docs..."
	@find ./docs/**/*.md -print | xargs grep "subcategory: \"\""; if [ $$(find ./docs/**/*.md -print | xargs grep "subcategory: \"\"" | wc -l) -ne 0 ]; then \
		echo ""; \
		echo "Documentation check found a blank subcategory for the above files.  Ensure a template is created (./templates) with a subcategory set."; \
		exit 1; \
	fi

depscheck:
	@echo "==> Checking source code with go mod tidy..."
	@go mod tidy
	@git diff --exit-code -- go.mod go.sum || \
		(echo; echo "Unexpected difference in go.mod/go.sum files. Run 'go mod tidy' command or revert any go.mod/go.sum changes and commit."; exit 1)

lint: golangci-lint providerlint importlint tflint terrafmtcheck

golangci-lint:
	@echo "==> Checking source code with golangci-lint..."
	@golangci-lint run ./...

importlint:
	@echo "==> Checking source code with importlint..."
	@impi --local . --scheme stdThirdPartyLocal ./...

providerlint:
	@echo "==> Checking source code with tfproviderlintx..."
	@tfproviderlintx \
		-c 1 \
		-AT001.ignored-filename-suffixes=_data_source_test.go \
		-XR004=false \
		-XS002=false \
		-R001=false \
		./internal/provider/... ./internal/service/...

tflint:
	@echo "==> Checking Terraform code with tflint..."
	@tflint --init

terrafmt:
	@echo "==> Formatting embedded Terraform code with terrafmt..."
	@find ./internal/service -type f -name '*_test.go' \
    | sort -u \
    | xargs -I {} terrafmt -f fmt {}

terrafmtcheck:
	@echo "==> Checking embedded Terraform code with terrafmt..."
	@find ./internal/service -type f -name '*_test.go' \
    | sort -u \
    | xargs -I {} terrafmt diff -f --check --fmtcompat {} ; if [ $$? -ne 0 ]; then \
		echo ""; \
		echo "terrafmt found bad formatting of HCL embedded in the test scripts. Please run "; \
		echo "\"make terrafmt\" before submitting the code for review."; \
		exit 1; \
	fi

fmt: terrafmt fmtcheck

devcheck: build vet tools fmt generate docscategorycheck lint test sweep testacc

.PHONY: tools build generate docscategorycheck test testacc sweep vet fmtcheck depscheck lint golangci-lint importlint providerlint tflint terrafmt terrafmtcheck testacc devcheck fmt
