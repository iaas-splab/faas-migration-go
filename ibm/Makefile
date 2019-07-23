.PHONY: build clean deploy


build:
	env GOOS=linux CGO_ENABLED=0 GOARCH=amd64 GO111MODULES=on go build -v -a -o bin/put/exec functions/put.go
	env GOOS=linux CGO_ENABLED=0 GOARCH=amd64 GO111MODULES=on go build -v -a -o bin/get/exec functions/get.go
	env GOOS=linux CGO_ENABLED=0 GOARCH=amd64 GO111MODULES=on go build -v -a -o bin/del/exec functions/delete.go
	env GOOS=linux CGO_ENABLED=0 GOARCH=amd64 GO111MODULES=on go build -v -a -o bin/lst/exec functions/list.go
	env GOOS=linux CGO_ENABLED=0 GOARCH=amd64 GO111MODULES=on go build -v -a -o bin/done/exec functions/done.go

package: build
	cd bin/put  && zip exec.zip exec
	cd bin/get  && zip exec.zip exec
	cd bin/del  && zip exec.zip exec
	cd bin/lst  && zip exec.zip exec
	cd bin/done && zip exec.zip exec

clean:
	rm -rvf ./bin

deploy: clean package
	cd bin/put && ibmcloud fn action create todo-put --native exec.zip
	cd bin/get && ibmcloud fn action create todo-get --native exec.zip
	cd bin/del && ibmcloud fn action create todo-del --native exec.zip
	cd bin/lst && ibmcloud fn action create todo-lst --native exec.zip
	cd bin/done && ibmcloud fn action create todo-done --native exec.zip

update: clean package
	cd bin/put && ibmcloud fn action update todo-put --native exec.zip
	cd bin/get && ibmcloud fn action update todo-get --native exec.zip
	cd bin/del && ibmcloud fn action update todo-del --native exec.zip
	cd bin/lst && ibmcloud fn action update todo-lst --native exec.zip
	cd bin/done && ibmcloud fn action update todo-done --native exec.zip

destroy:
	ibmcloud fn action delete todo-put
	ibmcloud fn action delete todo-get
	ibmcloud fn action delete todo-del
	ibmcloud fn action delete todo-lst
	ibmcloud fn action delete todo-done