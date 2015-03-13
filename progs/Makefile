
.PHONY: getdependencies clean all

getdependencies:
	go get gopkg.in/yaml.v2

validator: common.go validator.go
	go build validator.go common.go

statistics: common.go statistics.go
	go build statistics.go common.go

clean:
	rm validator statistics

all: validator statistics
