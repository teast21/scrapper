package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

type Deposit struct {
	Pubkey                 string
	Withdrawal_credentials string
	Signature              string
	Deposit_data_root      string
}

func main() {

	var deposit []Deposit

	fptr := flag.String("fpath", "test.txt", "file path of deposit information")
	flag.Parse()
	Data := []byte(*fptr)
	/*
				`
		    [{"pubkey": "pk1",
		    "withdrawal_credentials": "wc1",
		    "amount": 32000000000, "signature":
		    "s1", "deposit_message_root": "dmr",
		    "deposit_data_root": "dd1",
		    "fork_version": "00000000", "network_name": "mainnet", "deposit_cli_version": "2.2.0"},
		    {"pubkey": "pk2",
		    "withdrawal_credentials": "wc2",
		    "amount": 32000000000,
		    "signature": "s2", "deposit_message_root": "b1bc5324ee3d1653e95d4320877107274e0fc426371758acab49f26bca1c6b27",
		    "deposit_data_root": "dd2", "fork_version": "00000000", "network_name": "mainnet",
		    "deposit_cli_version": "2.2.0"}
		    ]`)
	*/
	err := json.Unmarshal(Data, &deposit)

	if err != nil {
		fmt.Println(err)
	}
	f, err2 := os.Create("dep.txt")

	if err2 != nil {
		log.Fatal(err2)
	}

	defer f.Close()

	//fmt.Println(len(deposit))

	for i := range deposit {
		f.WriteString(deposit[i].Pubkey)
	}
	f.WriteString("\n")
	for i := range deposit {
		f.WriteString(deposit[i].Withdrawal_credentials)
	}
	f.WriteString("\n")
	for i := range deposit {
		f.WriteString(deposit[i].Signature)
	}
	f.WriteString("\n")
	for i := range deposit {
		f.WriteString(deposit[i].Deposit_data_root)
	}
	fmt.Println("done")
}
