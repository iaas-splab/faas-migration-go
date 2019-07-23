package main

import (
	"encoding/json"
	"fmt"
	"github.com/c-mueller/faas-migration-go/core"
	"github.com/c-mueller/faas-migration-go/ibm"
)

func main() {
	repo := ibm.NewCloudantRepository()

	items, err := core.List(repo)
	if err != nil {
		fmt.Printf("Execution Failed: Error %s\n", err.Error())
		obj, _ := json.Marshal(ibm.Obejct{
			"statuscode": 500,
			"body":       fmt.Sprintf("Server Error: %s", err.Error()),
		})
		fmt.Println(string(obj))
		return
	}

	res, _ := json.Marshal(ibm.Obejct{
		"statuscode": 200,
		"headers": ibm.Obejct{
			"Content-Type": "application/json",
		},
		"body": ibm.Obejct{
			"items": items,
		},
	})

	fmt.Println(string(res))
}
