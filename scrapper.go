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
	err := json.Unmarshal(Data, &deposit)

	if err != nil {
		fmt.Println(err)
	}
	f, err2 := os.Create("dep.txt")

	if err2 != nil {
		log.Fatal(err2)
	}

	defer f.Close()

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
