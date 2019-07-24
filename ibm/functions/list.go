package main

import (
	"encoding/json"
	"fmt"
	"github.com/c-mueller/faas-migration-go/core"
	"github.com/c-mueller/faas-migration-go/ibm"
	"os"
)

func main() {
	fmt.Println(os.Args[1])
	var pobj ibm.Obejct
	json.Unmarshal([]byte(os.Args[1]), &pobj)
	repo := ibm.NewCloudantRepository(pobj)

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
