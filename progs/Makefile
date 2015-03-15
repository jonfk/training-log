
.PHONY: getdependencies clean all statistics

getdependencies:
	go get gopkg.in/yaml.v2

validator: validator.go
	go build validator.go

statistics:
	$(MAKE) -C statistics all

clean:
	rm validator
	$(MAKE) -C statistics clean

all: validator statistics
