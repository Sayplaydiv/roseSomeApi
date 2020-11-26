package roseApi

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	oasisGrpc "github.com/oasisprotocol/oasis-core/go/common/grpc"
	consensus "github.com/oasisprotocol/oasis-core/go/consensus/api"
	"github.com/oasisprotocol/oasis-core/go/consensus/api/transaction"
	staking "github.com/oasisprotocol/oasis-core/go/staking/api"
	"google.golang.org/grpc"
	"net/http"
	conf "roseSomeApi/config"
	"strconv"
)

const (
	PATH = "config/conf.ini"
)

func GetStatus(c *gin.Context) {
	//导入配置文件
	configMap := conf.InitConfig(PATH)
	//路径设置
	socket_Path := configMap["socket_Path"]
	conn, err := oasisGrpc.Dial(socket_Path, grpc.WithInsecure())
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"data":     nil,
			"errorMsg": fmt.Errorf("GetStatus oasisGrpc.Dial===", err),
		})
		return
	}
	cc := consensus.NewConsensusClient(conn)

	status, err := cc.GetStatus(context.Background())
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"data":     nil,
			"errorMsg": fmt.Errorf("cc.GetStatus==", err),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":     status,
		"errorMsg": "",
	})
	return
}

func GetSignerNonce(c *gin.Context) {
	address := c.Query("address")
	height := c.Query("height")
	if address != "" && height != "" {
		//导入配置文件
		configMap := conf.InitConfig(PATH)
		//路径设置
		socket_Path := configMap["socket_Path"]
		conn, err := oasisGrpc.Dial(socket_Path, grpc.WithInsecure())
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"data":     0,
				"errorMsg": fmt.Errorf("GetSignerNonce oasisGrpc.Dial===", err),
			})
			return
		}
		cc := consensus.NewConsensusClient(conn)
		var address_0 staking.Address
		address_0.UnmarshalBinary([]byte(address))
		heightInt64, err := strconv.ParseInt(height, 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"data":     0,
				"errorMsg": fmt.Errorf("GetSignerNonce strconv.ParseInt===", err),
			})
			return
		}
		nonceParam := consensus.GetSignerNonceRequest{
			address_0,
			heightInt64,
		}
		nonce, err := cc.GetSignerNonce(context.Background(), &nonceParam)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"data":     0,
				"errorMsg": fmt.Errorf("cc.GetSignerNonce===", err),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data":     nonce,
			"errorMsg": "",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data":     0,
			"errorMsg": "参数类型错误或者缺失，请检查后再发送",
		})
	}
}
func SubmitTxNoWait(c *gin.Context) {
	txHex := c.Query("txHex")

	if txHex != "" {
		//导入配置文件
		configMap := conf.InitConfig(PATH)
		//路径设置
		socket_Path := configMap["socket_Path"]
		conn, err := oasisGrpc.Dial(socket_Path, grpc.WithInsecure())
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"data":     false,
				"errorMsg": fmt.Errorf("SubmitTxNoWait oasisGrpc.Dial===", err),
			})
			return
		}
		cc := consensus.NewConsensusClient(conn)
		var txStr *transaction.SignedTransaction
		err = json.Unmarshal([]byte(txHex), &txStr)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"data":     false,
				"errorMsg": fmt.Errorf("txHex to  txStruct===", err),
			})
			return
		}
		err = cc.SubmitTxNoWait(context.Background(), txStr)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"data":     false,
				"errorMsg": fmt.Errorf("cc.GetSignerNonce===", err),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data":     true,
			"errorMsg": "",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data":     false,
			"errorMsg": "参数类型错误或者缺失，请检查后再发送",
		})
	}
}
