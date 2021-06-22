<template>
  <div>
    <el-form ref="form" :model="form" label-width="80px">
      <el-form-item label="告警名称">
        <span slot="label">
          <el-tooltip content="xxxxxxxx" placement="bottom">
            <i class="el-icon-question"></i>
          </el-tooltip>
          <span>xxxxxxxxxxxxxxxxxxxx{{ label2 }}xxxxxxxxxxxxx</span>
        </span>
        <el-input v-model="form.name"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">立即创建</el-button>
        <el-button>取消</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { ref, reactive } from 'vue'

export default ({
  prop: {
    monitorData: {
      type: Object,
      default () {
        return { message: 'hello' }
      }
    }
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
