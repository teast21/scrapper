package main

import (
	"encoding/json"
	"fmt"
)

type Deposit struct {
	pubkey string
	//withdrawal_credentials string
	//signature              string
	//deposit_data_root      string
}

func main() {

	var deposit []Deposit

	Data := []byte(`
    [{"pubkey": "b06a9e28251e0402af1a7249e6be0f27f1eb986ad8fa00e44619efe0a3b963f9302fe109bc198bdb7f279ced3fcdb652", 
    "withdrawal_credentials": "008cab8f5ba17523d30c8e5d48b0dae637d869e7bb9f5e278f4a710b9634ca3e", 
    "amount": 32000000000, "signature": 
    "abaf385e5814d2c852f038352696943d1f61931195724d934ebfedf561b8709cbb05cf2c465b74029f02a8680c5f5b8d091759bb241fff35c0479a9613f58ae8b86d77a738e9373f5f808d8f8d70f6588ca71c3c1182695ee3402e9a3686278c", "deposit_message_root": "afa202331b64e084dfd6f136cd3d63bc43e1c7e37b3135f02e260facb003bf17", 
    "deposit_data_root": "2f05ccfe2458b98999262cca199a82162d4e8238490587cc016e0cdeb2987045", 
    "fork_version": "00000000", "network_name": "mainnet", "deposit_cli_version": "2.2.0"}, 
    {"pubkey": "846e4d7e4d58407a9d9a8a1900e8fb72729fdf75f8fda3edb787921b50de45746765394c06f568e233978c52c04e77a8", 
    "withdrawal_credentials": "0033e43b8d20ab691308d708fb88f2af58a5df82f26ecdf9e54ff5d5acc3cd95", 
    "amount": 32000000000, 
    "signature": "a43b64d6c0ee176445a4bde7af42b0de982edce4a7d1e7b40708fb12206e4069a352a58c760cf13e53dc6e849fefae9b05528d5e1b1e4182e108474189bdc20f59fddedb4624315e67ba74cd74deed5023de7f97c2ef80cb4cc5cc995a074bc4", "deposit_message_root": "b1bc5324ee3d1653e95d4320877107274e0fc426371758acab49f26bca1c6b27", 
    "deposit_data_root": "fb3dc81853fad137f534ca816a97a130e9b162cbaba4571e8f00fc1d96998f64", "fork_version": "00000000", "network_name": "mainnet", 
    "deposit_cli_version": "2.2.0"}
    ]`)

	err := json.Unmarshal(Data, &deposit)

	if err != nil {
		fmt.Println(err)
	}

	for i := range deposit {
		fmt.Println(deposit[i].pubkey)
	}
}
