package controller

import (
	"backend/pkg"
	"fmt"

	"github.com/gin-gonic/gin"
)

// 保存了电子产品上链与查询的函数

func Uplink(c *gin.Context) {
	// 与userID不一样，取ID从第二位作为溯源码，即18位，雪花ID的结构如下（转化为10进制共19位）：
	// +--------------------------------------------------------------------------+
	// | 1 Bit Unused | 41 Bit Timestamp |  10 Bit NodeID  |   12 Bit Sequence ID |
	// +--------------------------------------------------------------------------+
	farmer_traceability_code := pkg.GenerateID()[1:]
	args := buildArgs(c, farmer_traceability_code)
	if args == nil {
		return
	}
	res, err := pkg.ChaincodeInvoke("Uplink", args)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "uplink failed" + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":              200,
		"message":           "uplink success",
		"txid":              res,
		"traceability_code": args[1],
	})
}

// 获取电子产品的上链信息
func GetProductinfo(c *gin.Context) {
	res, err := pkg.ChaincodeQuery("GetProductinfo", c.PostForm("traceability_code"))
	if err != nil {
		c.JSON(200, gin.H{
			"message": "查询失败：" + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "query success",
		"data":    res,
	})

}

// 获取用户的电子产品ID列表
func GetProductList(c *gin.Context) {
	userID, _ := c.Get("userID")
	res, err := pkg.ChaincodeQuery("GetProductList", userID.(string))
	if err != nil {
		c.JSON(200, gin.H{
			"message": "query failed" + err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "query success",
		"data":    res,
	})
}

// 获取所有的电子产品信息
func GetAllProductInfo(c *gin.Context) {
	res, err := pkg.ChaincodeQuery("GetAllProductInfo", "")
	fmt.Print("res", res)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "query failed" + err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "query success",
		"data":    res,
	})
}

// 获取电子产品上链历史
// func (s *SmartContract) GetProductHistory(ctx contractapi.TransactionContextInterface, traceability_code string) ([]HistoryQueryResult, error) {
func GetProductHistory(c *gin.Context) {
	res, err := pkg.ChaincodeQuery("GetProductHistory", c.PostForm("traceability_code"))
	if err != nil {
		c.JSON(200, gin.H{
			"message": "query failed" + err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "query success",
		"data":    res,
	})
}

func buildArgs(c *gin.Context, farmer_traceability_code string) []string {
	var args []string
	userID, _ := c.Get("userID")
	userType, _ := pkg.ChaincodeQuery("GetUserType", userID.(string))
	args = append(args, userID.(string))
	fmt.Print(userID)
	// 种植户不需要输入溯源码，其他用户需要，通过雪花算法生成ID
	if userType == "工厂" {
		args = append(args, farmer_traceability_code)
	} else {
		// 检查溯源码是否正确
		res, err := pkg.ChaincodeQuery("GetProductinfo", c.PostForm("traceability_code"))
		if res == "" || err != nil || len(c.PostForm("traceability_code")) != 18 {
			c.JSON(200, gin.H{
				"message": "请检查溯源码是否正确!!",
			})
			return nil
		} else {
			args = append(args, c.PostForm("traceability_code"))
		}
	}
	args = append(args, c.PostForm("arg1"))
	args = append(args, c.PostForm("arg2"))
	args = append(args, c.PostForm("arg3"))
	args = append(args, c.PostForm("arg4"))
	args = append(args, c.PostForm("arg5"))
	args = append(args, c.PostForm("arg6"))
	return args
}
