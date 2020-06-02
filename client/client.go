package main

import (
	"common"
	"fmt"
	"net/rpc"
)

func main() {

	// get RPC client by dialing at `rpc.DefaultRPCPath` endpoint
	client, _ := rpc.DialHTTP("tcp", "127.0.0.1:9000") // or `localhost:9000`

	/*--------------*/

	// create pao variable of type `common.User`
	var pao common.User

	// get User by id `1`
	if err := client.Call("Profile.Get", 1, &pao); err != nil {
		fmt.Println("Error:1 Profile.Get()", err)
	} else {
		fmt.Printf("Success:1 User '%s' found with id=1 \n", pao.FullName())
	}

	/*--------------*/

	// add User by id `1`
	if err := client.Call("Profile.Add", common.User{
		ID:        1,
		FirstName: "pao",
		LastName:  "pyo",
	}, &pao); err != nil {
		fmt.Println("Error:2 Profile.Add()", err)
	} else {
		fmt.Printf("Success:2 User '%s' created with id=1 \n", pao.FullName())
	}

}
