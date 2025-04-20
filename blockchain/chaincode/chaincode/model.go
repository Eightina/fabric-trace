package chaincode

/*
定义用户结构体
用户ID
用户类型
实名认证信息哈希,包括用户注册的姓名、身份证号、手机号、注册平台同意协议签名的哈希
电子产品列表
*/
type User struct {
	UserID       string   `json:"userID"`
	UserType     string   `json:"userType"`
	RealInfoHash string   `json:"realInfoHash"`
	ProductList    []*Product `json:"productList"`
}

/*
定义电子产品结构体
溯源码
工厂输入
经销商输入
平台输入
物流输入
售后输入
*/
type Product struct {
	Traceability_code string         `json:"traceability_code"`
	Factory_input     Factory_input  `json:"factory_input"`
	Retailer_input    Retailer_input `json:"retailer_input"`
	Platform_input    Platform_input `json:"platform_input"`
    Logistic_input    Logistic_input `json:"logistic_input"`
	Service_input     Service_input  `json:"service_input"`
}

// HistoryQueryResult structure used for handling result of history query
type HistoryQueryResult struct {
	Record    *Product `json:"record"`
	TxId      string `json:"txId"`
	Timestamp string `json:"timestamp"`
	IsDelete  bool   `json:"isDelete"`
}

/*
工厂输入
工厂名称与ID
物理地址
工厂联系方式
产品型号
出厂日期
质检信息
*/
type Factory_input struct {
    Fac_factoryNameAndID string `json:"fac_factoryNameAndID"` // 工厂名称与ID
    Fac_physicalAddress  string `json:"fac_physicalAddress"`  // 物理地址
    Fac_contactNumber    string `json:"fac_contactNumber"`    // 联系方式
    Fac_productModel     string `json:"fac_productModel"`     // 产品型号
    Fac_productionDate   string `json:"fac_productionDate"`   // 出厂日期
    Fac_qualityCheck     string `json:"fac_qualityCheck"`     // 质检信息
    Fac_Txid             string `json:"fac_txid"`             // 交易ID
    Fac_Timestamp        string `json:"fac_timestamp"`        // 时间戳
}

/*
经销商输入
经销商ID与名称
经销商地址
营业执照编号
入仓时间
售出时间
售出标价
*/
type Retailer_input struct {
    Ret_retailerNameAndID string `json:"ret_retailerNameAndID"` // 经销商ID与名称
    Ret_address           string `json:"ret_address"`          // 经销商地址
    Ret_licenseNumber     string `json:"ret_licenseNumber"`    // 营业执照编号
    Ret_warehousingTime   string `json:"ret_warehousingTime"`  // 入仓时间
    Ret_saleTime          string `json:"ret_saleTime"`         // 售出时间
    Ret_salePrice         string `json:"ret_salePrice"`        // 售出标价
    Ret_Txid              string `json:"ret_txid"`             // 交易ID
    Ret_Timestamp         string `json:"ret_timestamp"`        // 时间戳
}

/*
平台输入
平台名称ID（唯一标识符）
产品页面
营业执照编号
订单号
支付方式
交付方式
*/
type Platform_input struct {
    Pla_platformNameAndID    string `json:"pla_platformNameAndID"`     // 平台名称与ID
    Pla_productPage   string `json:"pla_productPage"`    // 产品页面
    Pla_licenseNumber string `json:"pla_licenseNumber"`  // 营业执照编号
    Pla_orderID       string `json:"pla_orderID"`        // 订单号
    Pla_paymentMethod string `json:"pla_paymentMethod"`  // 支付方式
    Pla_deliveryMode  string `json:"pla_deliveryMode"`   // 交付方式
    Pla_Txid          string `json:"pla_txid"`           // 交易ID
    Pla_Timestamp     string `json:"pla_timestamp"`      // 时间戳
}

/*
物流输入
物流公司名称与ID
物流单号
物流类型
派送联系方式
物流送达记录
*/
type Logistic_input struct {
    Log_logisticsNameAndID string `json:"log_logisticsNameAndID"` // 物流公司名称与ID
    Log_licenseNumber      string `json:"log_licenseNumber"`      // 营业执照编号
    Log_trackingNumber     string `json:"log_trackingNumber"`     // 物流单号
    Log_logisticsType      string `json:"log_logisticsType"`      // 物流类型
    Log_deliveryContact    string `json:"log_deliveryContact"`    // 派送联系方式
    Log_deliveryRecord     string `json:"log_deliveryRecord"`     // 物流送达记录
    Log_Txid               string `json:"log_txid"`               // 交易ID
    Log_Timestamp          string `json:"log_timestamp"`          // 时间戳
}

/*
售后输入
售后服务中心名称ID（唯一标识符）
服务中心地址
联系方式
售后原因
售后类型
售后结果
*/
type Service_input struct {
    Ser_serviceCenterID string `json:"ser_serviceCenterID"` // 售后服务中心名称ID
    Ser_address         string `json:"ser_address"`         // 服务中心地址
    Ser_contactNumber   string `json:"ser_contactNumber"`   // 联系方式
    Ser_reason          string `json:"ser_reason"`          // 售后原因
    Ser_type            string `json:"ser_type"`            // 售后类型
    Ser_result          string `json:"ser_result"`          // 售后结果
    Ser_Txid            string `json:"ser_txid"`            // 交易ID
    Ser_Timestamp       string `json:"ser_timestamp"`       // 时间戳
}