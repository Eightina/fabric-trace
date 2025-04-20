package chaincode

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// 定义合约结构体
type SmartContract struct {
	contractapi.Contract
}

// 注册用户
func (s *SmartContract) RegisterUser(ctx contractapi.TransactionContextInterface, userID string, userType string, realInfoHash string) error {
	user := User{
		UserID:       userID,
		UserType:     userType,
		RealInfoHash: realInfoHash,
		ProductList:    []*Product{},
	}
	userAsBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState(userID, userAsBytes)
	if err != nil {
		return err
	}
	return nil
}

// 电子产品上链，传入用户ID、电子产品上链信息
func (s *SmartContract) Uplink(ctx contractapi.TransactionContextInterface, userID string, traceability_code string, arg1 string, arg2 string, arg3 string, arg4 string, arg5 string, arg6 string) (string, error) {
	// 获取用户类型
	userType, err := s.GetUserType(ctx, userID)
	if err != nil {
		return "", fmt.Errorf("failed to get user type: %v", err)
	}

	// 通过溯源码获取电子产品的上链信息
	ProductAsBytes, err := ctx.GetStub().GetState(traceability_code)
	if err != nil {
		return "", fmt.Errorf("failed to read from world state: %v", err)
	}
	// 将电子产品的信息转换为Product结构体
	var product Product
	if ProductAsBytes != nil {
		err = json.Unmarshal(ProductAsBytes, &product)
		if err != nil {
			return "", fmt.Errorf("failed to unmarshal product: %v", err)
		}
	}

	//获取时间戳,修正时区
	txtime, err := ctx.GetStub().GetTxTimestamp()
	if err != nil {
		return "", fmt.Errorf("failed to read TxTimestamp: %v", err)
	}
	timeLocation, _ := time.LoadLocation("Asia/Shanghai") // 选择你所在的时区
	time := time.Unix(txtime.Seconds, 0).In(timeLocation).Format("2006-01-02 15:04:05")

	// 获取交易ID
	txid := ctx.GetStub().GetTxID()
	// 给电子产品信息中加上溯源码
	product.Traceability_code = traceability_code
	// 不同用户类型的上链的参数不一致
	switch userType {
	// 工厂
	case "工厂":
		// 将传入的电子产品上链信息转换为Factory_input结构体
		product.Factory_input.Fac_factoryNameAndID = arg1 // 工厂名称与ID
		product.Factory_input.Fac_physicalAddress = arg2  // 物理地址
		product.Factory_input.Fac_contactNumber = arg3    // 工厂联系方式
		product.Factory_input.Fac_productModel = arg4     // 产品型号
		product.Factory_input.Fac_productionDate = arg5   // 出厂日期
		product.Factory_input.Fac_qualityCheck = arg6     // 质检信息
		product.Factory_input.Fac_Txid = txid             // 交易ID
		product.Factory_input.Fac_Timestamp = time        // 时间戳
	
	// 经销商
	case "经销商":
		// 将传入的电子产品上链信息转换为Retailer_input结构体
		product.Retailer_input.Ret_retailerNameAndID = arg1 // 经销商名称与ID
		product.Retailer_input.Ret_address = arg2           // 经销商地址
		product.Retailer_input.Ret_licenseNumber = arg3     // 营业执照编号
		product.Retailer_input.Ret_warehousingTime = arg4   // 入仓时间
		product.Retailer_input.Ret_saleTime = arg5          // 售出时间
		product.Retailer_input.Ret_salePrice = arg6         // 售出标价
		product.Retailer_input.Ret_Txid = txid              // 交易ID
		product.Retailer_input.Ret_Timestamp = time         // 时间戳
	
	// 平台
	case "平台":
		// 将传入的电子产品上链信息转换为Platform_input结构体
		product.Platform_input.Pla_platformNameAndID = arg1     // 平台名称ID
		product.Platform_input.Pla_productPage = arg2    // 产品页面
		product.Platform_input.Pla_licenseNumber = arg3  // 营业执照编号
		product.Platform_input.Pla_orderID = arg4        // 订单号
		product.Platform_input.Pla_paymentMethod = arg5  // 支付方式
		product.Platform_input.Pla_deliveryMode = arg6   // 交付方式
		product.Platform_input.Pla_Txid = txid           // 交易ID
		product.Platform_input.Pla_Timestamp = time      // 时间戳
	
	// 物流
	case "物流":
		// 将传入的电子产品上链信息转换为Logistic_input结构体
		product.Logistic_input.Log_logisticsNameAndID = arg1 // 物流公司名称与ID
		product.Logistic_input.Log_licenseNumber = arg2      // 营业执照编号
		product.Logistic_input.Log_trackingNumber = arg3     // 物流单号
		product.Logistic_input.Log_logisticsType = arg4      // 物流类型
		product.Logistic_input.Log_deliveryContact = arg5    // 派送联系方式
		product.Logistic_input.Log_deliveryRecord = arg6     // 物流送达记录
		product.Logistic_input.Log_Txid = txid               // 交易ID
		product.Logistic_input.Log_Timestamp = time          // 时间戳
	
	// 售后
	case "售后":
		// 将传入的电子产品上链信息转换为Service_input结构体
		product.Service_input.Ser_serviceCenterID = arg1 // 售后服务中心名称ID
		product.Service_input.Ser_address = arg2         // 服务中心地址
		product.Service_input.Ser_contactNumber = arg3   // 联系方式
		product.Service_input.Ser_reason = arg4          // 售后原因
		product.Service_input.Ser_type = arg5            // 售后类型
		product.Service_input.Ser_result = arg6          // 售后结果
		product.Service_input.Ser_Txid = txid            // 交易ID
		product.Service_input.Ser_Timestamp = time       // 时间戳
	}

	//将电子产品的信息转换为json格式
	productAsBytes, err := json.Marshal(product)
	if err != nil {
		return "", fmt.Errorf("failed to marshal product: %v", err)
	}
	//将电子产品的信息存入区块链
	err = ctx.GetStub().PutState(traceability_code, productAsBytes)
	if err != nil {
		return "", fmt.Errorf("failed to put product: %v", err)
	}
	//将电子产品存入用户的电子产品列表
	err = s.AddProduct(ctx, userID, &product)
	if err != nil {
		return "", fmt.Errorf("failed to add product to user: %v", err)

	}

	return txid, nil
}

// 添加电子产品到用户的电子产品列表
func (s *SmartContract) AddProduct(ctx contractapi.TransactionContextInterface, userID string, product *Product) error {
	userBytes, err := ctx.GetStub().GetState(userID)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if userBytes == nil {
		return fmt.Errorf("the user %s does not exist", userID)
	}
	// 将返回的结果转换为User结构体
	var user User
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		return err
	}
	user.ProductList = append(user.ProductList, product)
	userAsBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState(userID, userAsBytes)
	if err != nil {
		return err
	}
	return nil
}

// 获取用户类型
func (s *SmartContract) GetUserType(ctx contractapi.TransactionContextInterface, userID string) (string, error) {
	userBytes, err := ctx.GetStub().GetState(userID)
	if err != nil {
		return "", fmt.Errorf("failed to read from world state: %v", err)
	}
	if userBytes == nil {
		return "", fmt.Errorf("the user %s does not exist", userID)
	}
	// 将返回的结果转换为User结构体
	var user User
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		return "", err
	}
	return user.UserType, nil
}

// 获取用户信息
func (s *SmartContract) GetUserInfo(ctx contractapi.TransactionContextInterface, userID string) (*User, error) {
	userBytes, err := ctx.GetStub().GetState(userID)
	if err != nil {
		return &User{}, fmt.Errorf("failed to read from world state: %v", err)
	}
	if userBytes == nil {
		return &User{}, fmt.Errorf("the user %s does not exist", userID)
	}
	// 将返回的结果转换为User结构体
	var user User
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		return &User{}, err
	}
	return &user, nil
}

// 获取电子产品的上链信息
func (s *SmartContract) GetProductinfo(ctx contractapi.TransactionContextInterface, traceability_code string) (*Product, error) {
	ProductAsBytes, err := ctx.GetStub().GetState(traceability_code)
	if err != nil {
		return &Product{}, fmt.Errorf("failed to read from world state: %v", err)
	}

	// 将返回的结果转换为Product结构体
	var product Product
	if ProductAsBytes != nil {
		err = json.Unmarshal(ProductAsBytes, &product)
		if err != nil {
			return &Product{}, fmt.Errorf("failed to unmarshal product: %v", err)
		}
		if product.Traceability_code != "" {
			return &product, nil
		}
	}
	return &Product{}, fmt.Errorf("the product %s does not exist", traceability_code)
}

// 获取用户的电子产品ID列表
func (s *SmartContract) GetProductList(ctx contractapi.TransactionContextInterface, userID string) ([]*Product, error) {
	userBytes, err := ctx.GetStub().GetState(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if userBytes == nil {
		return nil, fmt.Errorf("the user %s does not exist", userID)
	}
	// 将返回的结果转换为User结构体
	var user User
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		return nil, err
	}
	return user.ProductList, nil
}

// 获取所有的电子产品信息
func (s *SmartContract) GetAllProductInfo(ctx contractapi.TransactionContextInterface) ([]Product, error) {
	productListAsBytes, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	defer productListAsBytes.Close()
	var products []Product
	for productListAsBytes.HasNext() {
		queryResponse, err := productListAsBytes.Next()
		if err != nil {
			return nil, err
		}
		var product Product
		err = json.Unmarshal(queryResponse.Value, &product)
		if err != nil {
			return nil, err
		}
		// 过滤非电子产品的信息
		if product.Traceability_code != "" {
			products = append(products, product)
		}
	}
	return products, nil
}

// 获取电子产品上链历史
func (s *SmartContract) GetProductHistory(ctx contractapi.TransactionContextInterface, traceability_code string) ([]HistoryQueryResult, error) {
	log.Printf("GetAssetHistory: ID %v", traceability_code)

	resultsIterator, err := ctx.GetStub().GetHistoryForKey(traceability_code)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var records []HistoryQueryResult
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var product Product
		if len(response.Value) > 0 {
			err = json.Unmarshal(response.Value, &product)
			if err != nil {
				return nil, err
			}
		} else {
			product = Product{
				Traceability_code: traceability_code,
			}
		}

		timestamp, err := ptypes.Timestamp(response.Timestamp)
		if err != nil {
			return nil, err
		}
		// 指定目标时区
		targetLocation, err := time.LoadLocation("Asia/Shanghai")
		if err != nil {
			return nil, err
		}

		// 将时间戳转换到目标时区
		timestamp = timestamp.In(targetLocation)
		// 格式化时间戳为指定格式的字符串
		formattedTime := timestamp.Format("2006-01-02 15:04:05")

		record := HistoryQueryResult{
			TxId:      response.TxId,
			Timestamp: formattedTime,
			Record:    &product,
			IsDelete:  response.IsDelete,
		}
		records = append(records, record)
	}

	return records, nil
}
