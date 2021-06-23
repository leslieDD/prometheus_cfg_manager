<template>
  <el-card class="box-card" :body-style="{ padding: '0px 0px 0px 0px' }">
    <div v-for="data in customerData" :key="data.id">
      <el-select
        v-model="data.key"
        :key="data.id"
        allow-create
        default-first-option
        placeholder="请选择"
      >
        <el-option :label="data.key" :value="data.key"></el-option>
        <el-option
          v-for="(label, lid) in labels"
          :key="lid"
          :label="label"
          :value="label"
        >
        </el-option>
      </el-select>
      <el-input v-model="data.value" :key="data.id"></el-input>
    </div>
  </el-card>
</template>

<script>
import { reactive } from '@vue/reactivity'
export default {
  props: {
    customerData: {
      type: Object,
      default () {
        return [
          {
            id: 1,
            key: "summary1",
            value: "Environment: {{ $labels.env }} Instance: {{ $labels.instance }} is Down ! ! !"
          },
          {
            id: 2,
            key: "summary2",
            value: "Environment: {{ $labels.env }} Instance: {{ $labels.instance }} is Down ! ! !"
          },
          {
            id: 3,
            key: "summary3",
            value: "Environment: {{ $labels.env }} Instance: {{ $labels.instance }} is Down ! ! !"
          }
        ]
      }
    }
  },
  setup (props, context) {
    const labels = reactive([
      "summary",
      "description"
    ])
    return {
      labels
    }
  },
}
</script>

