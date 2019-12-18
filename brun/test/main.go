package main

import (
	"account/services"
	"encoding/json"
	"fmt"
)

func main() {

	d, e := json.Marshal(&services.AccountTransferDTO{})
	fmt.Println(e)
	fmt.Println(string(d))
}
