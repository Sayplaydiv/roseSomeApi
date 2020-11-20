# roseSomeApi
api server

## 此api服务主要提供交易处理与oasis_api_server 一块使用
## 其他查询类接口使用https://github.com/Sayplaydiv/oasis_api_server
## 所需包
```
import (
    "crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/oasisprotocol/oasis-core/go/common/crypto/signature"
	memorySigner "github.com/oasisprotocol/oasis-core/go/common/crypto/signature/signers/memory"
	"github.com/oasisprotocol/oasis-core/go/common/quantity"
	"github.com/oasisprotocol/oasis-core/go/consensus/api/transaction"
	"github.com/oasisprotocol/oasis-core/go/staking/api"
	staking "github.com/oasisprotocol/oasis-core/go/staking/api"
	"net/http"
	"strconv"
	"time"
	rose_resp "walletapiserver_go/walletsdk/rose_client/models/response"
)
```
## 生成地址
```
    //公私钥对生成
pu,pr,_:=ed25519.GenerateKey(rand.Reader)

	log.Println(hex.EncodeToString(pu))
	log.Println(hex.EncodeToString(pr))
	//publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	pk:=signature.NewPublicKey(hex.EncodeToString(pu))
	//log.Println(pk.String())
	address:=api.NewAddress(pk)
	log.Println(address.String())
```

## 交易

```
    //获取nonce
	nonce,_:=cc.GetSignerNonce(context.Background(),&nonceParam)

    //构建交易
	testTx := transaction.NewTransaction(nonce, nil, staking.MethodTransfer, &staking.Transfer{})

    //地址类型转换string-->addres
    var toAddress staking.Address
	toAddress.UnmarshalText([]byte(stringAddress))
    
 
    //私钥转换
	privaByte,err:=hex.DecodeString(req.PrivateKey)
	if err!=nil {
		return "", err
	}
	testSigner:= memorySigner.NewFromRuntime(privaByte)

   //签名
	testSigTx, _ := transaction.Sign(testSigner, testTx)
	log.Println("交易hash:",testSigTx.Hash())
	err:=cc.SubmitTxNoWait(context.Background(),testSigTx)
	if err!=nil {
		log.Println(err)
	}
```