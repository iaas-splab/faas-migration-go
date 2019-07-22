.PHONY: build clean deploy


build:
	env GOOS=linux GO111MODULES=on go build -v -a -ldflags="-s -w" -o bin/put functions/put.go
	env GOOS=linux GO111MODULES=on go build -v -a -ldflags="-s -w" -o bin/get functions/get.go
	env GOOS=linux GO111MODULES=on go build -v -a -ldflags="-s -w" -o bin/del functions/delete.go
	env GOOS=linux GO111MODULES=on go build -v -a -ldflags="-s -w" -o bin/lst functions/list.go
	env GOOS=linux GO111MODULES=on go build -v -a -ldflags="-s -w" -o bin/done functions/done.go

package: build
	cd bin && cp put exec && zip put.zip exec && rm exec
	cd bin && cp get exec && zip get.zip exec && rm exec
	cd bin && cp del exec && zip del.zip exec && rm exec
	cd bin && cp lst exec && zip lst.zip exec && rm exec
	cd bin && cp done exec && zip done.zip exec && rm exec

clean:
	rm -rvf ./bin

deploy: clean package
	ibmcloud fn action create todo-put --native bin/put.zip
	ibmcloud fn action create todo-get --native bin/get.zip
	ibmcloud fn action create todo-del --native bin/del.zip
	ibmcloud fn action create todo-lst --native bin/lst.zip
	ibmcloud fn action create todo-done --native bin/done.zip

update: clean package
	ibmcloud fn action update todo-put --native bin/put.zip
	ibmcloud fn action update todo-get --native bin/get.zip
	ibmcloud fn action update todo-del --native bin/del.zip
	ibmcloud fn action update todo-lst --native bin/lst.zip
	ibmcloud fn action update todo-done --native bin/done.zip

destroy:
	ibmcloud fn action delete todo-put
	ibmcloud fn action delete todo-get
	ibmcloud fn action delete todo-del
	ibmcloud fn action delete todo-lst
	ibmcloud fn action delete todo-done