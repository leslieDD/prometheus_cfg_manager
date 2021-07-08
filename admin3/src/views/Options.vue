<template>
  <el-descriptions
    class="margin-top"
    title="选项编辑"
    :column="1"
    size="medium"
    border
  >
    <!-- <template #extra>
      <el-button type="primary" size="small">操作</el-button>
    </template> -->
    <el-descriptions-item>
      <template #label>
        <!-- <i class="el-icon-user"></i> -->
        在发布JOB组时，针对没有配置任何子组的JOB组，是否为此JOB组的所有IP生成无标签子组
      </template>
      <el-tooltip
        :content="'当前:' + publish_at_null_subgroup_title"
        placement="top"
      >
        <el-switch
          v-model="options.publish_at_null_subgroup"
          active-color="#13ce66"
          inactive-color="#ff4949"
          @change="doPutOptions($event, 'publish_at_null_subgroup')"
          active-value="true"
          inactive-value="false"
        >
        </el-switch>
      </el-tooltip>
    </el-descriptions-item>
    <el-descriptions-item>
      <template #label>
        <!-- <i class="el-icon-mobile-phone"></i> -->
        在发布JOB组时，针对有配置子组的JOB组，是否为此JOB组的未分组IP生成无标签子组
      </template>
      <el-tooltip
        :content="'当前:' + publish_at_remain_subgroup_title"
        placement="top"
      >
        <el-switch
          v-model="options.publish_at_remain_subgroup"
          active-color="#13ce66"
          inactive-color="#ff4949"
          @change="doPutOptions($event, 'publish_at_remain_subgroup')"
          active-value="true"
          inactive-value="false"
        >
        </el-switch>
      </el-tooltip>
    </el-descriptions-item>
  </el-descriptions>
</template>

<script>
import { getOptions, putOptions } from '@/api/options.js'
export default {
  data () {
    return {
      options: {
        publish_at_null_subgroup: 'true',
        publish_at_remain_subgroup: 'true'
      }
    }
  },
  computed: {
    publish_at_null_subgroup_title: function () {
      if (this.options.publish_at_null_subgroup === 'true') {
        return '开启'
      } else {
        return '关闭'
      }
    },
    publish_at_remain_subgroup_title: function () {
      if (this.options.publish_at_remain_subgroup === 'true') {
        return '开启'
      } else {
        return '关闭'
      }
    }
  },
  mounted () {
    this.doGetOptions()
  },
  methods: {
    doGetOptions () {
      // 变化后的值
      getOptions().then(r => {
        this.options = r.data
      }).catch(e => console.log(e))
    },
    doPutOptions (event, optKey) {
      // 变化后的值
      putOptions({ [optKey]: event }).then(r => {
        this.$notify({
          title: '成功',
          message: '更新成功！',
          type: 'success'
        });
      }).catch(e => console.log(e))
    }
  }
}
</script>

<style scoped>
</style>