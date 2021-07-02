<template>
  <div class="main-board">
    <div class="base-config-menu">
      <el-tabs
        type="border-card"
        v-model="activeTabName"
        @tab-click="handleTabClick"
      >
        <el-tab-pane label="IP管理" name="ipManager">
          <router-view name="ipManager"></router-view>
        </el-tab-pane>
        <el-tab-pane label="分组管理" name="jobs">
          <router-view name="jobs"></router-view>
        </el-tab-pane>
        <el-tab-pane label="基本配置" name="baseConfig">
          <router-view name="baseConfig"></router-view>
        </el-tab-pane>
        <el-tab-pane label="告警管理" name="noticeManager">
          <router-view name="noticeManager"></router-view>
        </el-tab-pane>
        <el-tab-pane label="配置预览" name="preview">
          <router-view name="preview"></router-view>
        </el-tab-pane>
        <el-tab-pane label="分组预览" name="ftree">
          <router-view name="ftree"></router-view>
        </el-tab-pane>
        <el-tab-pane label="告警规则预览" name="ruleView">
          <router-view name="ruleView"></router-view>
        </el-tab-pane>
      </el-tabs>
    </div>
  </div>
</template>

<script>
export default {
  name: "Menu",
  data () {
    return {
      activeTabName: 'ipManager'
    }
  },
  mounted () {
    this.$router.push({ name: 'ipManager' })
  },
  methods: {
    handleTabClick (tab, event) {
      if (tab.instance.props.name === 'ipManager') {
        // this.doGetMechines()
        this.$router.push({ name: 'ipManager' })
      } else if (tab.instance.props.name === 'jobs') {
        // this.$refs.groupManagerRef.doGetJobs()
        this.$router.push({ name: 'jobs' })
      } else if (tab.instance.props.name === 'preview') {
        // this.$refs.previewRef.loadYaml()
        this.$router.push({ name: 'preview' })
      } else if (tab.instance.props.name === 'ftree') {
        // this.$refs.ftreeRef.doLoadAllFiles()
        this.$router.push({ name: 'ftree' })
      } else if (tab.instance.props.name === 'noticeManager') {
        // this.$refs.noticeManagerRef.doGetTree()
        this.$router.push({ name: 'noticeManager' })
      } else if (tab.instance.props.name === 'ruleView') {
        // this.$refs.ruleViewRef.doLoadAllRulesFiles()
        this.$router.push({ name: 'ruleView' })
      } else if (tab.instance.props.name === 'baseConfig') {
        // this.$router.push({ name: 'baseLabels' })
        this.$router.push({ name: 'baseConfig' })
      }
    }
  }
}
</script>

<style scoped>
.main-board {
  padding: 0;
  max-width: 1100px;
  margin: 0 auto;
}
el-tabs {
  padding: 0px;
}
</style>