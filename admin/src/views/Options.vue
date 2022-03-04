<template>
  <div class="option-board">
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
          1）在发布JOB组时，针对没有配置任何子组的JOB组，是否为此JOB组的所有IP生成无标签子组
        </template>
        <el-tooltip
          :content="'当前:' + titles.publish_at_null_subgroup"
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
          2）在发布JOB组时，针对有配置子组的JOB组，是否为此JOB组的未分组的IP添加进默认子组
        </template>
        <el-tooltip
          :content="'当前:' + titles.publish_at_remain_subgroup"
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
      <el-descriptions-item>
        <template #label>
          <!-- <i class="el-icon-mobile-phone"></i> -->
          3）在发布JOB组时，针对没有设置相关联IP的JOB组（IP数0），不在prometheus.yml中生成job项
        </template>
        <el-tooltip
          :content="'当前:' + titles.publish_at_empty_nocreate_file"
          placement="top"
        >
          <el-switch
            v-model="options.publish_at_empty_nocreate_file"
            active-color="#13ce66"
            inactive-color="#ff4949"
            @change="doPutOptions($event, 'publish_at_empty_nocreate_file')"
            active-value="true"
            inactive-value="false"
          >
          </el-switch>
        </el-tooltip>
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label>
          <!-- <i class="el-icon-mobile-phone"></i> -->
          4）在发布JOB组时，同进也更新JOB组对应的IP分组（在'分组预览'中看到的）文件
        </template>
        <el-tooltip
          :content="'当前:' + titles.publish_jobs_also_ips"
          placement="top"
        >
          <el-switch
            v-model="options.publish_jobs_also_ips"
            active-color="#13ce66"
            inactive-color="#ff4949"
            @change="doPutOptions($event, 'publish_jobs_also_ips')"
            active-value="true"
            inactive-value="false"
          >
          </el-switch>
        </el-tooltip>
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label>
          <!-- <i class="el-icon-mobile-phone"></i> -->
          5）在发布JOB组时，同时Reload Prometheus服务
        </template>
        <el-tooltip
          :content="'当前:' + titles.publish_jobs_also_reload_srv"
          placement="top"
        >
          <el-switch
            v-model="options.publish_jobs_also_reload_srv"
            active-color="#13ce66"
            inactive-color="#ff4949"
            @change="doPutOptions($event, 'publish_jobs_also_reload_srv')"
            active-value="true"
            inactive-value="false"
          >
          </el-switch>
        </el-tooltip>
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label>
          <!-- <i class="el-icon-mobile-phone"></i> -->
          6）在发布'IP管理'中发布IP分组文件时，同时Reload
          Prometheus服务（默认Prometheus自动更新）
        </template>
        <el-tooltip
          :content="'当前:' + titles.publish_ips_also_reload_srv"
          placement="top"
        >
          <el-switch
            v-model="options.publish_ips_also_reload_srv"
            active-color="#13ce66"
            inactive-color="#ff4949"
            @change="doPutOptions($event, 'publish_ips_also_reload_srv')"
            active-value="true"
            inactive-value="false"
          >
          </el-switch>
        </el-tooltip>
      </el-descriptions-item>
    </el-descriptions>
  </div>
</template>

<script>
import { getOptions, putOptions } from '@/api/options.js'
export default {
  data () {
    return {
      options: {
        publish_at_null_subgroup: 'false',
        publish_at_remain_subgroup: 'false',
        publish_at_empty_nocreate_file: 'false',
        publish_jobs_also_ips: 'false',
        publish_jobs_also_reload_srv: 'false',
        publish_ips_also_reload_srv: 'false'
      },
      titles: {
        publish_at_null_subgroup: '关闭',
        publish_at_remain_subgroup: '关闭',
        publish_at_empty_nocreate_file: '关闭',
        publish_jobs_also_ips: '关闭',
        publish_jobs_also_reload_srv: '关闭',
        publish_ips_also_reload_srv: '关闭'
      }
    }
  },
  watch: {
    options: {
      handler () {
        let newTitle = {}
        for (let key in this.options) {
          if (this.options[key] === 'true') {
            newTitle[key] = '开启'
          } else {
            newTitle[key] = '关闭'
          }
        }
        this.titles = { ...newTitle }
      },
      deep: true
    }
  },
  computed: {},
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
        // this.$message({
        //   showClose: true,
        //   message: '更新成功！',
        //   type: 'success'
        // })
      }).catch(e => console.log(e))
    }
  }
}
</script>

<style scoped>
.option-board {
  width: 95%;
}
</style>