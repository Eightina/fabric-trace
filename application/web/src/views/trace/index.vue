<template>
  <div class="trace-container">
    <el-input v-model="input" placeholder="请输入溯源码查询" style="width: 300px;margin-right: 15px;" />
    <el-button type="primary" plain @click="ProductInfo"> 查询 </el-button>
    <el-button type="success" plain @click="AllProductInfo"> 获取所有电子产品信息 </el-button>
    <el-table
      :data="tracedata"
      style="width: 100%"
    >
      <el-table-column type="expand">
        <template slot-scope="props">
          <el-form label-position="left" inline class="demo-table-expand">
            <div><span class="trace-text" style="color: #67C23A;">电子产品信息</span></div>
            <el-form-item label="工厂名称与ID：">
              <span>{{ props.row.factory_input.fac_factoryNameAndID }}</span>
            </el-form-item>
            <el-form-item label="物理地址：">
              <span>{{ props.row.factory_input.fac_physicalAddress }}</span>
            </el-form-item>
            <el-form-item label="联系方式：">
              <span>{{ props.row.factory_input.fac_contactNumber }}</span>
            </el-form-item>
            <el-form-item label="产品型号：">
              <span>{{ props.row.factory_input.fac_productModel }}</span>
            </el-form-item>
            <el-form-item label="出厂日期：">
              <span>{{ props.row.factory_input.fac_productionDate }}</span>
            </el-form-item>
            <el-form-item label="质检信息：">
              <span>{{ props.row.factory_input.fac_qualityCheck }}</span>
            </el-form-item>
            <el-form-item label="交易ID：">
              <span>{{ props.row.factory_input.fac_txid }}</span>
            </el-form-item>
            <el-form-item label="时间戳：">
              <span>{{ props.row.factory_input.fac_timestamp }}</span>
            </el-form-item>
            <div><span class="trace-text" style="color: #409EFF;">经销商信息</span></div>
            <el-form-item label="经销商ID与名称：">
              <span>{{ props.row.retailer_input.ret_retailerNameAndID }}</span>
            </el-form-item>
            <el-form-item label="经销商地址：">
              <span>{{ props.row.retailer_input.ret_address }}</span>
            </el-form-item>
            <el-form-item label="营业执照编号：">
              <span>{{ props.row.retailer_input.ret_licenseNumber }}</span>
            </el-form-item>
            <el-form-item label="入仓时间：">
              <span>{{ props.row.retailer_input.ret_warehousingTime }}</span>
            </el-form-item>
            <el-form-item label="售出时间：">
              <span>{{ props.row.retailer_input.ret_saleTime }}</span>
            </el-form-item>
            <el-form-item label="售出标价：">
              <span>{{ props.row.retailer_input.ret_salePrice }}</span>
            </el-form-item>
            <el-form-item label="交易ID：">
              <span>{{ props.row.retailer_input.ret_txid }}</span>
            </el-form-item>
            <el-form-item label="时间戳">
              <span>{{ props.row.retailer_input.ret_timestamp }}</span>
            </el-form-item>
            <div><span class="trace-text" style="color: #E6A23C;">平台信息</span></div>
            <el-form-item label="平台名称与ID：">
              <span>{{ props.row.platform_input.pla_platformNameAndID }}</span>
            </el-form-item>
            <el-form-item label="产品页面：">
              <span>{{ props.row.platform_input.pla_productPage }}</span>
            </el-form-item>
            <el-form-item label="营业执照编号：">
              <span>{{ props.row.platform_input.pla_licenseNumber }}</span>
            </el-form-item>
            <el-form-item label="订单号：">
              <span>{{ props.row.platform_input.pla_orderID }}</span>
            </el-form-item>
            <el-form-item label="支付方式：">
              <span>{{ props.row.platform_input.pla_paymentMethod }}</span>
            </el-form-item>
            <el-form-item label="交付方式：">
              <span>{{ props.row.platform_input.pla_deliveryMode }}</span>
            </el-form-item>
            <el-form-item label="交易ID：">
              <span>{{ props.row.platform_input.pla_txid }}</span>
            </el-form-item>
            <el-form-item label="时间戳：">
              <span>{{ props.row.platform_input.pla_timestamp }}</span>
            </el-form-item>
            <div><span class="trace-text" style="color: #909399;">物流信息</span></div>
            <el-form-item label="物流公司名称与ID：">
              <span>{{ props.row.logistic_input.log_logisticsNameAndID }}</span>
            </el-form-item>
            <el-form-item label="营业执照编号：">
              <span>{{ props.row.logistic_input.log_licenseNumber }}</span>
            </el-form-item>
            <el-form-item label="物流单号：">
              <span>{{ props.row.logistic_input.log_trackingNumber }}</span>
            </el-form-item>
            <el-form-item label="物流类型：">
              <span>{{ props.row.logistic_input.log_logisticsType }}</span>
            </el-form-item>
            <el-form-item label="派送联系方式：">
              <span>{{ props.row.logistic_input.log_deliveryContact }}</span>
            </el-form-item>
            <el-form-item label="物流送达记录：">
              <span>{{ props.row.logistic_input.log_deliveryRecord }}</span>
            </el-form-item>
            <el-form-item label="交易ID：">
              <span>{{ props.row.logistic_input.log_txid }}</span>
            </el-form-item>
            <el-form-item label="时间戳：">
              <span>{{ props.row.logistic_input.log_timestamp }}</span>
            </el-form-item>
            <div><span class="trace-text" style="color: #909399;">售后信息</span></div>
            <el-form-item label="售后服务中心名称ID：">
              <span>{{ props.row.service_input.ser_serviceCenterID }}</span>
            </el-form-item>
            <el-form-item label="服务中心地址：">
              <span>{{ props.row.service_input.ser_address }}</span>
            </el-form-item>
            <el-form-item label="联系方式：">
              <span>{{ props.row.service_input.ser_contactNumber }}</span>
            </el-form-item>
            <el-form-item label="售后原因：">
              <span>{{ props.row.service_input.ser_reason }}</span>
            </el-form-item>
            <el-form-item label="售后类型：">
              <span>{{ props.row.service_input.ser_type }}</span>
            </el-form-item>
            <el-form-item label="售后结果：">
              <span>{{ props.row.service_input.ser_result }}</span>
            </el-form-item>
            <el-form-item label="交易ID：">
              <span>{{ props.row.service_input.ser_txid }}</span>
            </el-form-item>
            <el-form-item label="时间戳：">
              <span>{{ props.row.service_input.ser_timestamp }}</span>
            </el-form-item>
          </el-form>
        </template>
      </el-table-column>
      <el-table-column
        label="溯源码"
        prop="traceability_code"
      />
      <el-table-column
        label="电子产品名称"
        prop="factory_input.fac_productModel"
      />
      <el-table-column
        label="电子产品出厂时间"
        prop="factory_input.fac_productionDate"
      />
    </el-table>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { getProductinfo, getProductList, getAllProductInfo, getProductHistory } from '@/api/trace'

export default {
  name: 'Trace',
  data() {
    return {
      tracedata: [],
      loading: false,
      input: ''
    }
  },
  computed: {
    ...mapGetters([
      'name',
      'userType'
    ])
  },
  created() {
    getProductList().then(res => {
      this.tracedata = JSON.parse(res.data)
    })
  },
  methods: {
    AllProductInfo() {
      getAllProductInfo().then(res => {
        this.tracedata = JSON.parse(res.data)
      })
    },
    ProductHistory() {
      getProductHistory().then(res => {
        // console.log(res)
      })
    },
    ProductInfo() {
      var formData = new FormData()
      formData.append('traceability_code', this.input)
      getProductinfo(formData).then(res => {
        if (res.code === 200) {
          // console.log(res)
          this.tracedata = []
          this.tracedata[0] = JSON.parse(res.data)
          return
        } else {
          this.$message.error(res.message)
        }
      })
    }
  }
}
</script>

<style lang="scss" scoped>

.demo-table-expand {
    font-size: 0;
  }
  .demo-table-expand label {
    width: 90px;
    color: #99a9bf;
  }
  .demo-table-expand .el-form-item {
    margin-right: 0;
    margin-bottom: 0;
    width: 50%;
  }
.trace {
  &-container {
    margin: 30px;
  }
  &-text {
    font-size: 30px;
    line-height: 46px;
  }
}
</style>
