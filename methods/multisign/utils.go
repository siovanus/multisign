package multisign

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"github.com/siovanus/multisign/config"
	"github.com/ontio/ontology-crypto/keypair"
	sdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology/core/types"
	"github.com/ontio/ontology/smartcontract/service/native/governance"
	"github.com/ontio/ontology/smartcontract/service/native/utils"
	"os"
)

var OntIDVersion = byte(0)

const (
	TX_PATH = "tx.txt"
	M       = 2
)

func makeAuthorizeTxAndSign(ontSdk *sdk.OntologySdk, signer *sdk.Account, publicKeyList []string, posList []uint32) bool {
	pubKeys, err := findPubKeys(signer.PublicKey)
	if err != nil {
		fmt.Println("findPubKeys error:", err)
		return false
	}
	address, err := types.AddressFromMultiPubKeys(pubKeys, 2)
	if err != nil {
		fmt.Println("types.AddressFromMultiPubKeys error:", err)
		return false
	}
	params := &governance.AuthorizeForPeerParam{
		Address:        address,
		PeerPubkeyList: publicKeyList,
		PosList:        posList,
	}
	method := "authorizeForPeer"
	contractAddress := utils.GovernanceContractAddress
	tx, err := ontSdk.Native.NewNativeInvokeTransaction(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
		OntIDVersion, contractAddress, method, []interface{}{params})
	if err != nil {
		fmt.Println("NewNativeInvokeTransaction error:", err)
		return false
	}
	err = ontSdk.MultiSignToTransaction(tx, M, pubKeys, signer)
	if err != nil {
		fmt.Println("ontSdk.MultiSignToTransaction error:", err)
		return false
	}
	itx, err := tx.IntoImmutable()
	if err != nil {
		fmt.Println("tx.IntoImmutable error:", err)
		return false
	}
	err = writeTxFile(itx.ToArray())
	if err != nil {
		fmt.Println("writeTxFile error:", err)
		return false
	}
	return true
}

func sign(ontSdk *sdk.OntologySdk, signer *sdk.Account) bool {
	tx, err := signTx(ontSdk, signer)
	if err != nil {
		fmt.Println("signTx error:", err)
		return false
	}
	itx, err := tx.IntoImmutable()
	if err != nil {
		fmt.Println("tx.IntoImmutable error:", err)
		return false
	}
	err = writeTxFile(itx.ToArray())
	if err != nil {
		fmt.Println("writeTxFile error:", err)
		return false
	}
	return true
}

func signAndSend(ontSdk *sdk.OntologySdk, signer *sdk.Account) bool {
	tx, err := signTx(ontSdk, signer)
	if err != nil {
		fmt.Println("signTx error:", err)
		return false
	}
	itx, err := tx.IntoImmutable()
	if err != nil {
		fmt.Println("tx.IntoImmutable error:", err)
		return false
	}
	err = writeTxFile(itx.ToArray())
	if err != nil {
		fmt.Println("writeTxFile error:", err)
		return false
	}
	txHash, err := ontSdk.SendTransaction(tx)
	if err != nil {
		fmt.Println("ontSdk.SendTransaction error:", err)
		return false
	}
	fmt.Println("signAndSend success, txHash is:", txHash.ToHexString())
	return true
}

func readTxFile() ([]byte, error) {
	f, err := os.Open(TX_PATH)
	if err != nil {
		return nil, fmt.Errorf("os.Open error:%s", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	if scanner.Scan() {
		s := scanner.Text()
		r, err := hex.DecodeString(s)
		if err != nil {
			return nil, fmt.Errorf("hex.DecodeString error:%s", err)
		}
		return r, nil
	}
	return nil, fmt.Errorf("scanner.Scan() failed")
}

func writeTxFile(data []byte) error {
	f, err := os.Create(TX_PATH)
	if err != nil {
		return fmt.Errorf("os.Create error: %v", err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	w.WriteString(hex.EncodeToString(data))
	w.Flush()
	return nil
}

func findPubKeys(pubKey keypair.PublicKey) ([]keypair.PublicKey, error) {
	p := hex.EncodeToString(keypair.SerializePublicKey(pubKey))
	for _, pubKeys := range config.DefPubKeysGroup.PubKeysGroup {
		for _, v := range pubKeys {
			if v == p {
				r := make([]keypair.PublicKey, 0)
				for _, s := range pubKeys {
					sBytes, err := hex.DecodeString(s)
					if err != nil {
						return nil, fmt.Errorf("hex.DecodeString error:%s", err)
					}
					pk, err := keypair.DeserializePublicKey(sBytes)
					if err != nil {
						return nil, fmt.Errorf("keypair.DeserializePublicKey error:%s", err)
					}
					r = append(r, pk)
				}
				return r, nil
			}
		}
	}
	return nil, fmt.Errorf("findPubKeys error: can no find any record")
}

func signTx(ontSdk *sdk.OntologySdk, signer *sdk.Account) (*types.MutableTransaction, error) {
	pubKeys, err := findPubKeys(signer.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("findPubKeys error:%s", err)
	}
	raw, err := readTxFile()
	if err != nil {
		return nil, fmt.Errorf("readTxFile error:%s", err)
	}
	itx, err := types.TransactionFromRawBytes(raw)
	if err != nil {
		return nil, fmt.Errorf("types.TransactionFromRawBytes error:%s", err)
	}
	tx, err := itx.IntoMutable()
	if err != nil {
		return nil, fmt.Errorf("itx.IntoMutable error:%s", err)
	}
	err = ontSdk.MultiSignToTransaction(tx, M, pubKeys, signer)
	if err != nil {
		return nil, fmt.Errorf("ontSdk.MultiSignToTransaction error:%s", err)
	}
	return tx, nil
}
