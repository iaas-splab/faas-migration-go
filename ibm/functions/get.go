package main

import (
	"encoding/json"
	"fmt"
	"github.com/c-mueller/faas-migration-go/core"
	"github.com/c-mueller/faas-migration-go/ibm"
	"os"
)

func main() {
	arg := os.Args[1]

	fmt.Println(arg)
	var obj core.IDRequest
	json.Unmarshal([]byte(arg), &obj)
	var pobj ibm.Obejct
	json.Unmarshal([]byte(arg), &pobj)
	repo := ibm.NewCloudantRepository(pobj)

	item, err := core.Get(obj,repo)
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
		"body": item,
	})

	fmt.Println(string(res))
}
