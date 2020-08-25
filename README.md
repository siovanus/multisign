# multisign

## Instruction

### 1.clone repo

```shell
git clone https://github.com/siovanus/multisign.git
```

### 2.build or download latest release

bulid:

```shell
go build main.go
```

or download latest release:

https://github.com/siovanus/multisign/releases

### 3.update config file

```shell
vim config.json
```

content of config.json：

```json
{
  "JsonRpcAddress":"http://polaris2.ont.io:20336",
  "WalletPath": "wallet.dat",
  "GasPrice":2500,
  "GasLimit":20000,
  "PublicKeyList":["030a34dcb075d144df1f65757b85acaf053395bb47b019970607d2d1cdd222525c"],
  "PosList":[1]
}
```

`JsonRpcAddress`：rpc of ontology node，for polaris testnet: `http://polaris1.ont.io:20336`，`http://polaris2.ont.io:20336`，`http://polaris3.ont.io:20336`，`http://polaris4.ont.io:20336`，for mainnet: `http://dappnode1.ont.io:20336`，`http://dappnode2.ont.io:20336`，`http://dappnode3.ont.io:20336`，`http://dappnode4.ont.io:20336`。

`WalletPath`: path of wallet file, this wallet is deposit account

`PublicKeyList`: list of public key want to authorize

`PosList`: list of pos want to authorize

### 4.update public key group file

```shell
vim pubKeysGroup.json
```

content of pubKeysGroup.json：

```json
{
  "PubKeysGroup":[
    ["0253ccfd439b29eca0fe90ca7c6eaa1f98572a054aa2d1d56e72ad96c466107a85","035eb654bad6c6409894b9b42289a43614874c7984bde6b03aaf6fc1d0486d9d45","0281d198c0dd3737a9c39191bc2d1af7d65a44261a8a64d6ef74d63f27cfb5ed92"],
    ["022e911fb5a20b4b2e4f917f10eb92f27d17cad16b916bce8fd2dd8c11ac2878c0","0253719ac66d7cafa1fe49a64f73bd864a346da92d908c19577a003a8a4160b7fa","02765d98bb092962734e365bd436bdc80c5b5991dcf22b28dbb02d3b3cf74d6444"]
  ]
}
```

`PubKeysGroup`: list of multi-sign group, each group consist of several public keys, when user make tx and sign with his own account, this tool will search the group file and find out other publickey of this user's group

### 5.run command line

list of supported command line: 

| command line              | function                                                     |
| ------------------------- | ------------------------------------------------------------ |
| `./main -t MakeAuthorizeTxAndSign` | read input data from config.json to make authorizeForPeer tx and sign it, write signed tx to tx.txt file |
| `./main -t MakeUnAuthorizeTxAndSign` | read input data from config.json to make unAuthorizeForPeer tx and sign it, write signed tx to tx.txt file |
| `./main -t MakeWithdrawTxAndSign` | read input data from config.json to make withdraw tx and sign it, write signed tx to tx.txt file |
| `./main -t Sign` | read raw tx from tx.txt and sign it |
| `./main -t SignAndSend` | read raw tx from tx.txt, sign it and send it to ontology network configured in config.json |

And now you can run your command and input your password if needed.