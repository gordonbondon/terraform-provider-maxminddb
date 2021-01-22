default: testacc

# Run acceptance tests
.PHONY: testacc
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m

.PHONY: tools
tools:
	go install github.com/katbyte/terrafmt
	go install github.com/bflad/tfproviderlint/cmd/tfproviderlint
	go install github.com/bflad/tfproviderdocs
