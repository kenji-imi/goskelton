package tmpl

// [projeect_root]/Makefile
const MakefileTmpl = `FMTPKGS=` + "`go list`" + `
VETPKGS=` + "`go list`" + `
LINTPKGS=` + "`go list`" + `

GOTEST= go test -v

.PHONY: init_mod
init_mod:
	go mod init github.com/{{ .User }}/{{ .Project }}

.PHONY: setup
setup:
	go get -u
		golang.org/x/tools/cmd/goimports
		github.com/kisielk/errcheck
		github.com/gcpug/zagane
		github.com/stretchr/testify/assert
	go get -d honnef.co/go/tools/cmd/staticcheck
		cd ${GOPATH}/src/honnef.co/go/tools/cmd/staticcheck && git checkout 2017.2.2
		go install honnef.co/go/tools/cmd/staticcheck


## lint

.PHONY: lint
lint: fmt vet staticcheck errcheck

.PHONY: fmt
fmt:
	goimports -l $(FMTPKGS) | grep -E '.'; test $$? -eq 1
	gofmt -l $(FMTPKGS) | grep -E '.'; test $$? -eq 1

.PHONY: vet
vet:
	go vet $(VETPKGS)

.PHONY: staticcheck
staticcheck:
	staticcheck -checks "SA*" -ignore ` + "`cat .staticcheckignore`" + ` $(LINTPKGS)

.PHONY: errcheck
errcheck:
	errcheck -ignore 'fmt:[FS]?[Pp]rint*' -exclude .errcheckignore $(LINTPKGS)

.PHONY: golint
golint:
	golint src/... | grep -v 'should have comment or be unexported' | grep -v 'comment on exported function'

.PHONY: zagane
zagane:
	zagane src/...


## test

.PHONY: test_unit
test_unit:
	$(GOTEST) ./src/...


## xo

.PHONY: init_xo
init_xo:
	go get -v -u github.com/knq/xo

.PHONY: gen_xo_models
gen_xo_models:
	cd ./xo_sample/; xo pgsql://hoge@localhost:5432/testdb?sslmode=disable -o xo_models


## grpc

GRPC_GATEWAY_REPO=${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis
PROTO_NAME=hello

.PHONY: init_grpc
init_grpc:
	go get -v -u github.com/golang/protobuf/protoc-gen-go
	go get -v -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway


.PHONY: protoc_go
protoc_go:
	protoc -I. \
	-I${GRPC_GATEWAY_REPO} \
	--go_out=plugins=grpc:. \
	pb/${PROTO_NAME}.proto

.PHONY: protoc_gateway_go
protoc_gateway_go:
	protoc -I/usr/local/include -I. -I${GOPATH}/src \
	-I${GRPC_GATEWAY_REPO} \
	--grpc-gateway_out=logtostderr=true:. \
	pb/${PROTO_NAME}.gw.proto

.PHONY: downgrade_protoc # https://qiita.com/go_sagawa/items/5ba0ebb0cf42a629e2e9
downgrade_protoc:
	brew upgrade protobuf
	go get -v -u google.golang.org/grpc
	go get -v -u github.com/golang/protobuf/proto
	go get -v -u github.com/golang/protobuf/protoc-gen-go
	go get -v -u go.pedge.io/protoeasy/cmd/protoeasy
	go get -v -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -v -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

	cd ${GOPATH}/src/github.com/golang/protobuf/protoc-gen-go
	git checkout v1.2.0
	go install
	git checkout master

	cd ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	git checkout v1.5.1
	go install
	git checkout master

.PHONY: pprof
pprof:
	go tool pprof -http=":8888" localhost:8080/debug/pprof/profile
`
