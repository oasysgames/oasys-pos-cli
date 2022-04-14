package crypto

import (
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"
)

var createAccountCmd = &cobra.Command{
	Use:   cmdPrefix + "create-account",
	Short: "Create a new account.",
	Run: func(cmd *cobra.Command, args []string) {
		out, err := cmd.Flags().GetString(outputFlag)
		if err != nil {
			panic(err)
		}

		address, key, err := createAccount()
		if err != nil {
			panic(err)
		}

		content := ""
		content += "Address : " + address + "\n"
		content += "Key     : 0x" + key + "\n"

		if out != "" {
			ioutil.WriteFile(out, []byte(content), 0600)
			fmt.Println("created")
		} else {
			fmt.Print(content)
		}
	},
}

func createAccount() (string, string, error) {
	privKey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", err
	}
	pubKey, ok := (privKey.Public()).(*ecdsa.PublicKey)
	if !ok {
		return "", "", errors.New("failed to private key to public key")
	}
	return crypto.PubkeyToAddress(*pubKey).String(), hex.EncodeToString(crypto.FromECDSA(privKey)), nil
}
