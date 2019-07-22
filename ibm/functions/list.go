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
		return
	}

	data := map[string]interface{}{
		"items": items,
	}

	res, _ := json.Marshal(data)

	fmt.Println(string(res))
}
