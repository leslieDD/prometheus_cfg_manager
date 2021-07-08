<template>
  <div class="main-board">
    <div class="base-config-menu">
      <el-tabs
        type="border-card"
        v-model="activeTabName"
        @tab-click="handleTabClick"
      >
        <el-tab-pane label="IP管理" name="ipManager">
          <router-view name="ipManager" v-slot="{ Component }">
            <transition name="slide-fade">
              <component :is="Component" />
            </transition>
          </router-view>
        </el-tab-pane>
        <el-tab-pane label="JOB组管理" name="jobs">
          <router-view name="jobs" v-slot="{ Component }">
            <transition name="slide-fade">
              <component :is="Component" />
            </transition>
          </router-view>
        </el-tab-pane>
        <el-tab-pane label="基本配置" name="baseConfig">
          <router-view name="baseConfig" v-slot="{ Component }">
            <transition name="slide-fade">
              <component :is="Component" />
            </transition>
          </router-view>
        </el-tab-pane>
        <el-tab-pane label="告警管理" name="noticeManager">
          <router-view name="noticeManager" v-slot="{ Component }">
            <transition name="slide-fade">
              <component :is="Component" />
            </transition>
          </router-view>
        </el-tab-pane>
        <el-tab-pane label="配置预览" name="preview">
          <router-view name="preview" v-slot="{ Component }">
            <transition name="slide-fade">
              <component :is="Component" />
            </transition>
          </router-view>
        </el-tab-pane>
        <el-tab-pane label="分组预览" name="ftree">
          <router-view name="ftree" v-slot="{ Component }">
            <transition name="slide-fade">
              <component :is="Component" />
            </transition>
          </router-view>
        </el-tab-pane>
        <el-tab-pane label="告警规则预览" name="ruleView">
          <router-view name="ruleView" v-slot="{ Component }">
            <transition name="slide-fade">
              <component :is="Component" />
            </transition>
          </router-view>
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
      activeTabName: 'ipManager',
      transitionName: 'slide-left'
    }
  },
  watch: {
    '$route' (to, from) {
      // console.log('现在路由:',to.path.split('/')[1],'来自路由:',from.path.split('/')[1],'现在的动画:',this.transitionName)
      const toDepth = to.path.split('/').length
      const fromDepth = from.path.split('/').length
      this.transitionName = toDepth < fromDepth ? 'slide-right' : 'slide-left'
    }
  },
  mounted () {
    this.$router.push({ name: 'ipManager' })
  },
  methods: {
    handleTabClick (tab, event) {
      if (tab.instance.props.name === 'ipManager') {
        this.$router.push({ name: 'ipManager' })
      } else if (tab.instance.props.name === 'jobs') {
        this.$router.push({ name: 'jobs' })
      } else if (tab.instance.props.name === 'preview') {
        this.$router.push({ name: 'preview' })
      } else if (tab.instance.props.name === 'ftree') {
        this.$router.push({ name: 'ftree' })
      } else if (tab.instance.props.name === 'noticeManager') {
        this.$router.push({ name: 'noticeManager' })
      } else if (tab.instance.props.name === 'ruleView') {
        this.$router.push({ name: 'ruleView' })
      } else if (tab.instance.props.name === 'baseConfig') {
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
.slide-fade-enter-active {
  transition: all 0.3s ease-out;
}

/* .slide-fade-leave-active {
  transition: all 0.8s cubic-bezier(1, 0.5, 0.8, 1);
} */

.slide-fade-enter-from,
.slide-fade-leave-to {
  /* transition: all 0.5s ease-out; */
  transform: translateX(30px);
  opacity: 0;
}
</style>