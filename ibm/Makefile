CLOUDANT_INSTANCE_NAME := todo-cloudant-db
CLOUDANT_ACCESS_KEY_NAME := access-key
CLOUDANT_INSTANCE_PLAN := Lite

GOFLAGS := -v

.PHONY: build clean deploy

init:
	ibmcloud service create cloudantNoSQLDB $(CLOUDANT_INSTANCE_PLAN) $(CLOUDANT_INSTANCE_NAME) || true
	ibmcloud service key-create $(CLOUDANT_INSTANCE_NAME) $(CLOUDANT_ACCESS_KEY_NAME) || true

build_creds: PASSWORD := `cat bin/raw_creds.json | jq -r '.password'`
build_creds: USERNAME := `cat bin/raw_creds.json | jq -r '.username'`
build_creds: HOSTNAME := `cat bin/raw_creds.json | jq -r '.url'`
build_creds: init
	mkdir bin || true
	ibmcloud service key-show todo-cloudant-db $(CLOUDANT_ACCESS_KEY_NAME) | tail -n +5 | jq > bin/raw_creds.json
	cp cloudant_creds.template.json cloudant_creds.json
	jq ".cloudant_password = \"$(PASSWORD)\"" cloudant_creds.json | sponge cloudant_creds.json
	jq ".cloudant_username = \"$(USERNAME)\"" cloudant_creds.json | sponge cloudant_creds.json
	jq ".cloudant_host = \"$(HOSTNAME)\"" cloudant_creds.json | sponge cloudant_creds.json

build:
	env GOOS=linux CGO_ENABLED=0 GOARCH=amd64 GO111MODULES=on go build $(GOFLAGS) -o bin/put/exec functions/put.go
	env GOOS=linux CGO_ENABLED=0 GOARCH=amd64 GO111MODULES=on go build $(GOFLAGS) -o bin/get/exec functions/get.go
	env GOOS=linux CGO_ENABLED=0 GOARCH=amd64 GO111MODULES=on go build $(GOFLAGS) -o bin/del/exec functions/delete.go
	env GOOS=linux CGO_ENABLED=0 GOARCH=amd64 GO111MODULES=on go build $(GOFLAGS) -o bin/lst/exec functions/list.go
	env GOOS=linux CGO_ENABLED=0 GOARCH=amd64 GO111MODULES=on go build $(GOFLAGS) -o bin/done/exec functions/done.go

package: build
	cd bin/put  && zip exec.zip exec
	cd bin/get  && zip exec.zip exec
	cd bin/del  && zip exec.zip exec
	cd bin/lst  && zip exec.zip exec
	cd bin/done && zip exec.zip exec

clean:
	rm -rvf ./bin

deploy_actions: clean package build_creds
	cd bin/put && ibmcloud fn action create todo-put --web true --native exec.zip -a final true -P ../../cloudant_creds.json
	cd bin/get && ibmcloud fn action create todo-get --web true --native exec.zip -a final true -P ../../cloudant_creds.json
	cd bin/del && ibmcloud fn action create todo-del --web true --native exec.zip -a final true -P ../../cloudant_creds.json
	cd bin/lst && ibmcloud fn action create todo-lst --web true --native exec.zip -a final true -P ../../cloudant_creds.json
	cd bin/done && ibmcloud fn action create todo-done --web true --native exec.zip -a final true -P ../../cloudant_creds.json

deploy_api:
	ibmcloud fn api create /todo /lst get todo-lst --response-type http
	ibmcloud fn api create /todo /get get todo-get --response-type http
	ibmcloud fn api create /todo /put post todo-put --response-type http
	ibmcloud fn api create /todo /del post todo-del --response-type http
	ibmcloud fn api create /todo /done post todo-done --response-type http
	ibmcloud fn api list --full

create_database:
	ibmcloud fn action invoke Bluemix_$(CLOUDANT_INSTANCE_NAME)_$(CLOUDANT_ACCESS_KEY_NAME)/create-database -r -p dbname items || true

delete_database:
	ibmcloud fn action invoke Bluemix_$(CLOUDANT_INSTANCE_NAME)_$(CLOUDANT_ACCESS_KEY_NAME)/delete-database -r -p dbname items || true

deploy: deploy_actions deploy_api create_database

update: clean package
	cd bin/put && ibmcloud fn action update todo-put --native exec.zip -a final true -P ../../cloudant_creds.json
	cd bin/get && ibmcloud fn action update todo-get --native exec.zip -a final true -P ../../cloudant_creds.json
	cd bin/del && ibmcloud fn action update todo-del --native exec.zip -a final true -P ../../cloudant_creds.json
	cd bin/lst && ibmcloud fn action update todo-lst --native exec.zip -a final true -P ../../cloudant_creds.json
	cd bin/done && ibmcloud fn action update todo-done --native exec.zip -a final true -P ../../cloudant_creds.json

destroy:
	ibmcloud fn action delete todo-put
	ibmcloud fn action delete todo-get
	ibmcloud fn action delete todo-del
	ibmcloud fn action delete todo-lst
	ibmcloud fn action delete todo-done
	ibmcloud fn api delete /todo 
