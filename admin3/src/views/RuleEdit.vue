<template>
  <el-card class="box-card" :body-style="{ padding: '20px 5px 0px 5px' }">
    <div class="box-member">
      <el-form ref="form" :model="form" label-width="180px" size="mini">
        <el-form-item>
          <template #label>
            <span>告警名称(alert)</span>
            <el-tooltip
              content="警报规则的名称。alert: <string>"
              placement="top"
            >
              <i class="el-icon-question"></i>
            </el-tooltip>
          </template>
          <el-input v-model="form.name"></el-input>
        </el-form-item>
        <el-form-item>
          <template #label>
            <span>表达式(expr)</span>
            <el-tooltip
              content="使用PromQL表达式完成的警报触发条件，用于计算是否有满足触发条件。expr: <string>"
              placement="top"
            >
              <i class="el-icon-question"></i>
            </el-tooltip>
          </template>
          <el-input v-model="form.name"></el-input>
        </el-form-item>
        <el-form-item>
          <template #label>
            <span>状态转换等待时间(for)</span>
            <el-tooltip
              content="带有for子句的警报触发以后首先会先转换成 Pending 状态，然后在转换为 Firing 状态。这里需要俩个周期才能触发警报条件，如果没有设置 for 子句，会直接 Inactive 状态转换成 Firing状态，然后触发警报，发送给 Receiver 设置的通知人。for: <string>，如：for: 5m，(m代表分钟)"
              placement="top"
              popper-class="filed-explain"
            >
              <i class="el-icon-question"></i>
            </el-tooltip>
          </template>
          <el-input v-model="form.name"></el-input>
        </el-form-item>
        <el-form-item>
          <template #label>
            <span>标签(labels)</span>
            <el-tooltip
              content="自定义标签，允许自行定义标签附加在警报上。<lable_name>: <label_value>"
              placement="top"
            >
              <i class="el-icon-question"></i>
            </el-tooltip>
          </template>
          <el-input v-model="form.name"></el-input>
        </el-form-item>
        <el-form-item>
          <template #label>
            <span>注释(annotations)</span>
            <el-tooltip
              content="用来设置有关警报的一组描述信息，其中包括自定义的标签，以及expr计算后的值。<lable_name>: <tmpl_string>"
              placement="top"
              popper-class="filed-explain"
            >
              <i class="el-icon-question"></i>
            </el-tooltip>
          </template>
          <KeyValue></KeyValue>
          <!-- <el-input v-model="form.name"></el-input> -->
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onSubmit">更新</el-button>
        </el-form-item>
      </el-form>
    </div>
  </el-card>
</template>

<script>
import { ref, reactive } from 'vue'
import KeyValue from '@/components/KeyValue.vue'

export default ({
  prop: {
    monitorData: {
      type: Object,
      default () {
        return { message: 'hello' }
      }
    }
  },
  components: {
    KeyValue
  },
  setup (props, context) {
    const tmplData = reactive([
      {
        alert: {
          label: "说明",
        }
      }])
    const form = reactive({
      name: '',
      region: '',
      date1: '',
      date2: '',
      delivery: false,
      type: [],
      resource: '',
      desc: ''
    })
    return {
      tmplData,
      form
    }
  },
  methods: {
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
</style>