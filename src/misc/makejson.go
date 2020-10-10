package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Println("Test")
	var usr map[string]string
	fmt.Println(usr)
	usr = make(map[string]string)

	var userName string
	var userAddress string

	fmt.Println("Enter your name: ")
	userName, userNameErr := in.ReadString('\n')

	if userNameErr != nil {
		fmt.Println(userNameErr)
		return
	}

	fmt.Println("Enter your address: ")
	userAddress, addrErr := in.ReadString('\n')

	if addrErr != nil {
		fmt.Println(addrErr)
		return
	}

	usr["name"] = userName
	usr["address"] = userAddress

	fmt.Println(usr)

	byteArr, err := json.Marshal(usr)

	fmt.Println(byteArr)

	if err != nil {
		fmt.Println(err)
		return
	}

	println(usr)

}
