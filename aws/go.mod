module github.com/c-mueller/faas-migration-go/aws

go 1.12

require (
	github.com/aws/aws-lambda-go v1.8.2
	github.com/aws/aws-sdk-go v1.16.36
	github.com/c-mueller/faas-migration-go/core v0.0.0-00010101000000-000000000000
	github.com/cenkalti/backoff v2.1.1+incompatible
	github.com/gofrs/uuid v3.2.0+incompatible
	github.com/google/uuid v1.1.1
	github.com/guregu/dynamo v1.1.0
	github.com/jmespath/go-jmespath v0.0.0-20180206201540-c2b33e8439af
	golang.org/x/net v0.0.0-20190213061140-3a22650c66bd
)

replace github.com/c-mueller/faas-migration-go/core => ../core
