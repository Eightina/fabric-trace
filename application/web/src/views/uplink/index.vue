<template>
  <div class="uplink-container">
    <div style="color:#909399;margin-bottom: 30px">
      当前用户：{{ name }};
      用户角色: {{ userType }}
    </div>
    <div>
      <el-form ref="form" :model="tracedata" label-width="80px" size="mini" style="">
        <el-form-item v-show="userType!='工厂'&userType!='消费者'" label="溯源码:" style="width: 300px" label-width="120px">
          <el-input v-model="tracedata.traceability_code" />
        </el-form-item>

        <div v-show="userType=='工厂'">
          <el-form-item label="工厂名称与ID:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Factory_input.Fac_factoryNameAndID" />
          </el-form-item>
          <el-form-item label="物理地址:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Factory_input.Fac_physicalAddress" />
          </el-form-item>
          <el-form-item label="工厂联系方式:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Factory_input.Fac_contactNumber" />
          </el-form-item>
          <el-form-item label="产品型号:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Factory_input.Fac_productModel" />
          </el-form-item>
          <el-form-item label="出厂日期:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Factory_input.Fac_productionDate" />
          </el-form-item>
          <el-form-item label="质检信息:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Factory_input.Fac_qualityCheck" />
          </el-form-item>
        </div>
        <div v-show="userType=='经销商'">
          <el-form-item label="经销商名称与ID:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Retailer_input.Ret_retailerNameAndID" />
          </el-form-item>
          <el-form-item label="经销商地址:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Retailer_input.Ret_address" />
          </el-form-item>
          <el-form-item label="营业执照编号:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Retailer_input.Ret_licenseNumber" />
          </el-form-item>
          <el-form-item label="入仓时间:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Retailer_input.Ret_warehousingTime" />
          </el-form-item>
          <el-form-item label="售出时间:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Retailer_input.Ret_saleTime" />
          </el-form-item>
          <el-form-item label="售出标价:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Retailer_input.Ret_salePrice" />
          </el-form-item>
        </div>
        <div v-show="userType=='平台'">
          <el-form-item label="姓名:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Platform_input.Pla_platformNameAndID" />
          </el-form-item>
          <el-form-item label="年龄:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Platform_input.Pla_productPage" />
          </el-form-item>
          <el-form-item label="联系电话:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Platform_input.Pla_licenseNumber" />
          </el-form-item>
          <el-form-item label="车牌号:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Platform_input.Pla_orderID" />
          </el-form-item>
          <el-form-item label="运输记录:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Platform_input.Pla_paymentMethod" />
          </el-form-item>
          <el-form-item label="交付方式:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Platform_input.Pla_deliveryMode" />
          </el-form-item>
        </div>
        <div v-show="userType=='物流'">
          <el-form-item label="物流公司名称与ID:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Logistic_input.Log_logisticsNameAndID" />
          </el-form-item>
          <el-form-item label="营业执照编号:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Logistic_input.Log_licenseNumber" />
          </el-form-item>
          <el-form-item label="物流单号:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Logistic_input.Log_trackingNumber" />
          </el-form-item>
          <el-form-item label="物流类型:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Logistic_input.Log_logisticsType" />
          </el-form-item>
          <el-form-item label="派送联系方式:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Logistic_input.Log_deliveryContact" />
          </el-form-item>
          <el-form-item label="物流送达记录:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Logistic_input.Log_deliveryRecord" />
          </el-form-item>
        </div>
        <div v-show="userType=='售后'">
          <el-form-item label="售后服务中心名称ID:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Logistic_input.Ser_serviceCenterID" />
          </el-form-item>
          <el-form-item label="服务中心地址:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Logistic_input.Ser_address" />
          </el-form-item>
          <el-form-item label="联系方式:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Logistic_input.Ser_contactNumber" />
          </el-form-item>
          <el-form-item label="售后原因:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Logistic_input.Ser_reason" />
          </el-form-item>
          <el-form-item label="售后类型:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Logistic_input.Ser_type" />
          </el-form-item>
          <el-form-item label="售后结果:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Logistic_input.Ser_result" />
          </el-form-item>
        </div>
      </el-form>
      <span slot="footer" style="color: gray;" class="dialog-footer">
        <el-button v-show="userType != '消费者'" type="primary" plain style="margin-left: 220px;" @click="submittracedata()">提 交</el-button>
      </span>
      <span v-show="userType == '消费者'" slot="footer" style="color: gray;" class="dialog-footer">
        消费者无需录入，可直接使用溯源功能~
      </span>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { uplink } from '@/api/trace'

export default {
  name: 'Uplink',
  data() {
    return {
      tracedata: {
        traceability_code: '',
        Factory_input: {
          Fac_factoryNameAndID: '',
          Fac_physicalAddress: '',
          Fac_contactNumber: '',
          Fac_productModel: '',
          Fac_productionDate: '',
          Fac_qualityCheck: ''
        },
        Retailer_input: {
          Ret_retailerNameAndID: '',
          Ret_address: '',
          Ret_licenseNumber: '',
          Ret_warehousingTime: '',
          Ret_saleTime: '',
          Ret_salePrice: ''
        },
        Platform_input: {
          Pla_platformNameAndID: '',
          Pla_productPage: '',
          Pla_licenseNumber: '',
          Pla_orderID: '',
          Pla_paymentMethod: '',
          Pla_deliveryMode: ''
        },
        Logistic_input: {
          Log_logisticsNameAndID: '',
          Log_licenseNumber: '',
          Log_trackingNumber: '',
          Log_logisticsType: '',
          Log_deliveryContact: '',
          Log_deliveryRecord: ''
        },
        Service_input: {
          Ser_serviceCenterID: '',
          Ser_address: '',
          Ser_contactNumber: '',
          Ser_reason: '',
          Ser_type: '',
          Ser_result: ''
        }
      },
      loading: false
    }
  },
  computed: {
    ...mapGetters([
      'name',
      'userType'
    ])
  },
  methods: {
    submittracedata() {
      console.log(this.tracedata)
      const loading = this.$loading({
        lock: true,
        text: '数据上链中...',
        spinner: 'el-icon-loading',
        background: 'rgba(0, 0, 0, 0.7)'
      })
      var formData = new FormData()
      formData.append('traceability_code', this.tracedata.traceability_code)
      // 根据不同的用户给arg1、arg2、arg3..赋值,
      switch (this.userType) {
        case '工厂':
          formData.append('arg1', this.tracedata.Factory_input.Fac_factoryNameAndID)
          formData.append('arg2', this.tracedata.Factory_input.Fac_physicalAddress)
          formData.append('arg3', this.tracedata.Factory_input.Fac_contactNumber)
          formData.append('arg4', this.tracedata.Factory_input.Fac_productModel)
          formData.append('arg5', this.tracedata.Factory_input.Fac_productionDate)
          formData.append('arg6', this.tracedata.Factory_input.Fac_qualityCheck)
          break
        case '经销商':
          formData.append('arg1', this.tracedata.Retailer_input.Ret_retailerNameAndID)
          formData.append('arg2', this.tracedata.Retailer_input.Ret_address)
          formData.append('arg3', this.tracedata.Retailer_input.Ret_licenseNumber)
          formData.append('arg4', this.tracedata.Retailer_input.Ret_warehousingTime)
          formData.append('arg5', this.tracedata.Retailer_input.Ret_saleTime)
          formData.append('arg6', this.tracedata.Retailer_input.Ret_salePrice)
          break
        case '平台':
          formData.append('arg1', this.tracedata.Platform_input.Pla_platformNameAndID)
          formData.append('arg2', this.tracedata.Platform_input.Pla_productPage)
          formData.append('arg3', this.tracedata.Platform_input.Pla_licenseNumber)
          formData.append('arg4', this.tracedata.Platform_input.Pla_orderID)
          formData.append('arg5', this.tracedata.Platform_input.Pla_paymentMethod)
          formData.append('arg6', this.tracedata.Platform_input.Pla_deliveryMode)
          break
        case '物流':
          formData.append('arg1', this.tracedata.Logistic_input.Log_logisticsNameAndID)
          formData.append('arg2', this.tracedata.Logistic_input.Log_licenseNumber)
          formData.append('arg3', this.tracedata.Logistic_input.Log_trackingNumber)
          formData.append('arg4', this.tracedata.Logistic_input.Log_logisticsType)
          formData.append('arg5', this.tracedata.Logistic_input.Log_deliveryContact)
          formData.append('arg6', this.tracedata.Logistic_input.Log_deliveryRecord)
          break
        case '售后':
          formData.append('arg1', this.tracedata.Service_input.Ser_serviceCenterID)
          formData.append('arg2', this.tracedata.Service_input.Ser_address)
          formData.append('arg3', this.tracedata.Service_input.Ser_contactNumber)
          formData.append('arg4', this.tracedata.Service_input.Ser_reason)
          formData.append('arg5', this.tracedata.Service_input.Ser_type)
          formData.append('arg6', this.tracedata.Service_input.Ser_result)
          break
      }
      uplink(formData).then(res => {
        if (res.code === 200) {
          loading.close()
          this.$message({
            message: '上链成功，交易ID：' + res.txid + '\n溯源码：' + res.traceability_code,
            type: 'success'
          })
        } else {
          loading.close()
          this.$message({
            message: '上链失败',
            type: 'error'
          })
        }
      }).catch(err => {
        loading.close()
        console.log(err)
      })
    }
  }
}

</script>

<style lang="scss" scoped>
.uplink {
  &-container {
    margin: 30px;
  }
  &-text {
    font-size: 30px;
    line-height: 46px;
  }
}
</style>
