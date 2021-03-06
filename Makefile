_GOPATH 			:= $(PWD)/../../../..
export GOPATH := $(_GOPATH)
GENERATED_TYPES := bool string int int8 int16 int32 int64 uint uint8 uint16 uint32 uint64 byte rune float32 float64 complex64 complex128
# on mac use different option for -i
UNAME_S = $(shell uname -s)
ifeq ($(UNAME_S),Darwin)
	SED_OPT := -i ''
else
	SED_OPT := -i
endif

.PHONY: all
all: codegen code

ipv6code:
	cp template/tree_v4.go template/tree_v6_generated.go
	sed ${SED_OPT} -e 's/TreeV4/TreeV6/g' template/tree_v6_generated.go
	sed ${SED_OPT} -e 's/treeNodeV4/treeNodeV6/g' template/tree_v6_generated.go
	sed ${SED_OPT} -e 's/IPv4Address/IPv6Address/g' template/tree_v6_generated.go

codegen: ipv6code $(addprefix codegen-,$(GENERATED_TYPES))

codegen-%:
	@echo "** generating $* tree"
	mkdir -p "./${*}_tree"
	cp -pa template/*.go "./${*}_tree"
	rm -f ./${*}_tree/*_test.go
	rm -f ./${*}_tree/types.go
	( cd "${*}_tree" && sed ${SED_OPT} "s/GeneratedType/${*}/g" *.go )
	( cd "${*}_tree" && sed ${SED_OPT} "s/package template/package ${*}_tree/g" *.go )

.PHONY: clean
clean:
	rm -rf *_tree
	rm -f template/tree_v6_generated.go

.PHONY: code
code:
	go build `go list ./... | grep -v /vendor/ `

.PHONY: test
test:
	go test -v `go list ./... | grep -v /vendor/ `
