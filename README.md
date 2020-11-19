# roseSomeApi
api server

## 此api服务主要提供交易处理与oasis_api_server 一块使用
## 其他查询类接口使用https://github.com/Sayplaydiv/oasis_api_server
## 所需包
```
import (
	"context"
	"crypto/rand"
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"github.com/oasisprotocol/oasis-core/go/common/crypto/signature"
	"github.com/oasisprotocol/oasis-core/go/staking/api"
	oasisGrpc "github.com/oasisprotocol/oasis-core/go/common/grpc"
	consensus "github.com/oasisprotocol/oasis-core/go/consensus/api"
	"github.com/oasisprotocol/oasis-core/go/consensus/api/transaction"
	//"github.com/oasisprotocol/ed25519"
	memorySigner "github.com/oasisprotocol/oasis-core/go/common/crypto/signature/signers/memory"
	staking "github.com/oasisprotocol/oasis-core/go/staking/api"
	"google.golang.org/grpc"
	"log"
	"testing"
)
```
## 生成地址
```

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
	testSigner:= memorySigner.NewFromRuntime(pr)

    //签名
	testSigTx, _ := transaction.Sign(testSigner, testTx)
	log.Println("交易hash:",testSigTx.Hash())
	err:=cc.SubmitTxNoWait(context.Background(),testSigTx)
	if err!=nil {
		log.Println(err)
	}
```