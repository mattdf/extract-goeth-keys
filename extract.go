package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("Usage: extract <key-file>")
		os.Exit(1)
	}

	file, e := ioutil.ReadFile(os.Args[1])

	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter password: ")
	pass, _ := reader.ReadString('\n')

	var jsontype accounts.Key
	json.Unmarshal(file, &jsontype)

	key, err := accounts.DecryptKey(file, strings.Trim(pass, "\r\n"))

	if err != nil {
		fmt.Printf("Key Error: %v", err)
		os.Exit(1)
	}

	fmt.Printf("Address: %x\n", key.Address)
	fmt.Printf("PubKey Curve Points: %x, %x\n", key.PrivateKey.PublicKey.X, key.PrivateKey.PublicKey.Y)
	fmt.Printf("PrivateKey: %x\n", key.PrivateKey.D)

}
