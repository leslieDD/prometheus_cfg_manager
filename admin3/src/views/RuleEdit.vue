<template>
  <el-card class="box-card" :body-style="{ padding: '0px 0px 0px 0px' }">
    <div class="box-member">
      <el-form
        label-position="top"
        ref="form"
        :model="formData"
        label-width="180px"
        size="mini"
      >
        <el-form-item>
          <template #label>
            <span>告警名称(alert)：</span>
            <el-tooltip
              content="警报规则的名称。alert: <string>"
              placement="top"
            >
              <i class="el-icon-question"></i>
            </el-tooltip>
          </template>
          <el-input v-model="formData.alert"></el-input>
        </el-form-item>
        <el-form-item>
          <template #label>
            <span>表达式(expr)：</span>
            <el-tooltip
              content="使用PromQL表达式完成的警报触发条件，用于计算是否有满足触发条件。expr: <string>"
              placement="top"
            >
              <i class="el-icon-question"></i>
            </el-tooltip>
          </template>
          <el-input v-model="formData.expr"></el-input>
        </el-form-item>
        <el-form-item>
          <template #label>
            <span>状态转换等待时间(for)：</span>
            <el-tooltip
              content="带有for子句的警报触发以后首先会先转换成 Pending 状态，然后在转换为 Firing 状态。这里需要俩个周期才能触发警报条件，如果没有设置 for 子句，会直接 Inactive 状态转换成 Firing状态，然后触发警报，发送给 Receiver 设置的通知人。for: <string>，如：for: 5m，(m代表分钟)"
              placement="top"
              popper-class="filed-explain"
            >
              <i class="el-icon-question"></i>
            </el-tooltip>
          </template>
          <el-input v-model="formData.for"></el-input>
        </el-form-item>
        <el-form-item>
          <template #label>
            <span>标签(labels)：</span>
            <el-tooltip
              content="自定义标签，允许自行定义标签附加在警报上。<lable_name>: <label_value>"
              placement="top"
            >
              <i class="el-icon-question"></i>
            </el-tooltip>
          </template>
          <div>
            <div v-for="data in formData.labels" :key="data.id">
              <el-card class="box-card-two" :body-style="{ padding: '1px 0px 1px 0px' }">
                <el-select
                  v-model="data.key"
                  :key="data.id"
                  filterable
                  allow-create
                  default-first-option
                  placeholder="请选择"
                  style="width: 20%"
                >
                  <el-option :label="data.key" :value="data.key"></el-option>
                  <el-option
                    v-for="item in defaultLabels"
                    :key="item.id"
                    :label="item.label"
                    :value="item.label"
                  >
                  </el-option>
                </el-select>
                <el-input
                  style="width: 80%"
                  v-model="data.value"
                  :key="data.id"
                ></el-input>
              </el-card>
            </div>
          </div>
        </el-form-item>
        <el-form-item>
          <template #label>
            <span>注释(annotations)：</span>
            <el-tooltip
              content="用来设置有关警报的一组描述信息，其中包括自定义的标签，以及expr计算后的值。<lable_name>: <tmpl_string>"
              placement="top"
              popper-class="filed-explain"
            >
              <i class="el-icon-question"></i>
            </el-tooltip>
          </template>
          <div>
            <div v-for="data in formData.annotations" :key="data.id">
              <el-card class="box-card-two" :body-style="{ padding: '1px 0px 1px 0px' }">
                <el-select
                  v-model="data.key"
                  :key="data.id"
                  allow-create
                  filterable
                  default-first-option
                  placeholder="请选择"
                  style="width: 20%"
                >
                  <el-option :label="data.key" :value="data.key"></el-option>
                  <el-option
                    v-for="item in defaultLabels"
                    :key="item.id"
                    :label="item.label"
                    :value="item.label"
                  >
                  </el-option>
                </el-select>
                <el-input
                  style="width: 80%"
                  v-model="data.value"
                  :key="data.id"
                ></el-input>
              </el-card>
            </div>
          </div>
        </el-form-item>
        <el-form-item align="right">
          <el-button style="margin-right: 20px" type="primary" @click="onSubmit"
            >更新</el-button
          >
        </el-form-item>
      </el-form>
    </div>
  </el-card>
</template>

<script>
import KeyValue from '@/components/KeyValue.vue'
import { getRuleDetail, getDefaultLabels } from '@/api/monitor.js'

export default ({
  props: {
    query: {
      type: Object,
      default: null
    }
  },
  components: {
    KeyValue
  },
  data() {
    return {
      formData: {
        alert: 'x',
        expr: 'x',
        for: 'x',
        labels: [],
        annotations: []
      },
      defaultLabels: []
    }
  },
  methods: {
    doGetRuleDetail (info) {
      if (!info) {
        return false
      }
      getDefaultLabels().then(
        r => {
          this.defaultLabels = r.data
        }
      ).catch(
        e => { console.log(e) }
      )
      getRuleDetail(info).then(
        r => {
          this.formData = r.data
        }
      ).catch(
        e => { console.log(e) }
      )
    },
    onSubmit () {

    }
  }
})
</script>


<style>
.box-member {
  width: 100%;
}
.box-card {
  /* width: 98%; */
  padding: 0px 0px;
}
.el-form-item__label {
  padding-bottom: 0px !important;
}
</style>