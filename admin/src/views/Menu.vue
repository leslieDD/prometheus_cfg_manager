<template>
  <div class="main-board">
    <div class="person-center-btn">
      <div>
        <el-button size="mini" type="danger" @click="goControlCenter">
          <single-svg icon-class="wxctl" />控制中心</el-button
        >
      </div>
      <div><span>&nbsp;&nbsp;</span></div>
      <div>
        <el-button size="mini" type="warning" @click="goBackPerson">
          <single-svg icon-class="pig" />个人中心</el-button
        >
      </div>
    </div>
    <div>
      <el-dialog
        v-model="dialogVisible"
        title="Prometheus服务控制中心"
        width="600px"
        :before-close="handleClose"
      >
        <el-descriptions
          class="margin-top"
          :column="1"
          size="small"
          border
        >

          <el-descriptions-item>
            <template #label>
              <!-- <el-icon><user /></el-icon> -->
              <!-- <icon-svg icon-class="yezi" /> -->
              <single-svg icon-class="prometheus"></single-svg>
              打开Prometheus自带后台
            </template>
            <el-button size="small" type="info" @click="open_prometheus_self" plain>点击执行操作</el-button>
          </el-descriptions-item>

          <el-descriptions-item>
            <template #label>
              <!-- <el-icon><user /></el-icon> -->
              <icon-svg icon-class="yezi" />
              重新生成所有规则
            </template>
            <el-button size="small" type="warning" @click="create" plain>点击执行操作</el-button>
          </el-descriptions-item>
          <el-descriptions-item>
            <template #label>
              <icon-svg icon-class="yezi" />
              重新加载（Reload）配置
            </template>
            <el-button size="small" type="danger" @click="reload" plain>点击执行操作</el-button>
          </el-descriptions-item>
          <el-descriptions-item>
            <template #label>
              <icon-svg icon-class="yezi" />
              重新生成所有规则，并重新加载（Reload）配置
            </template>
            <el-button size="small" type="danger" @click="createAreload" plain>点击执行操作</el-button>
          </el-descriptions-item>
        </el-descriptions>
        <!-- <span>This is a message</span> -->
        <!-- <template #footer>
          <span class="dialog-footer">
            <el-button @click="dialogVisible = false">Cancel</el-button>
            <el-button type="primary" @click="dialogVisible = false"
              >Confirm</el-button
            >
          </span>
        </template> -->
      </el-dialog>
    </div>
    <div class="base-config-menu">
      <el-tabs
        type="border-card"
        v-model="activeTabName"
        @tab-click="handleTabClick"
      >
        <el-tab-pane label="JOB组管理" name="jobs">
          <template #label>
            <span>
              <icon-svg icon-class="promgtxinxianshuiguo" />
              <!-- <i class="iconfont prometheusshuiguopingguo"></i> -->
              JOB组管理</span
            >
          </template>
          <router-view name="jobs" v-slot="{ Component }">
            <transition name="slide-fade">
              <component :is="Component" />
            </transition>
          </router-view>
        </el-tab-pane>
        <el-tab-pane label="IP管理" name="ipManager">
          <template #label>
            <span>
              <icon-svg icon-class="promgtshuiguo5" />
              <!-- <i class="iconfont prometheusshengxianshuiguo"></i> -->
              IP管理</span
            >
          </template>
          <router-view name="ipManager" v-slot="{ Component }">
            <transition name="slide-fade">
              <component :is="Component" />
            </transition>
          </router-view>
        </el-tab-pane>
        <el-tab-pane label="IDC机房" name="idc">
          <template #label>
            <span>
              <icon-svg icon-class="promgtshengxianshuiguo" />
              <!-- <i class="iconfont prometheusshengxianshuiguo"></i> -->
              IDC机房</span
            >
          </template>
          <router-view name="idc" v-slot="{ Component }">
            <transition name="slide-fade">
              <component :is="Component" />
            </transition>
          </router-view>
        </el-tab-pane>
        <el-tab-pane label="基本配置" name="baseConfig">
          <template #label>
            <span>
              <!-- <i class="iconfont prometheusshuiguo1"></i> -->
              <icon-svg icon-class="promgtshuiguo7" />
              基本配置</span
            >
          </template>
          <router-view name="baseConfig" v-slot="{ Component }">
            <transition name="slide-fade">
              <component :is="Component" />
            </transition>
          </router-view>
        </el-tab-pane>
        <el-tab-pane label="告警管理" name="noticeManager">
          <template #label>
            <span>
              <!-- <i class="iconfont prometheushuolongguo"></i> -->
              <icon-svg icon-class="promgtshuiguo3" />
              告警管理</span
            >
          </template>
          <router-view name="noticeManager" v-slot="{ Component }">
            <transition name="slide-fade">
              <component :is="Component" />
            </transition>
          </router-view>
        </el-tab-pane>
        <el-tab-pane label="配置预览" name="preview">
          <template #label>
            <span>
              <!-- <i class="iconfont prometheuskiwi__easyico"></i> -->
              <icon-svg icon-class="promgtshuiguo1" />
              配置预览</span
            >
          </template>
          <router-view name="preview" v-slot="{ Component }">
            <transition name="slide-fade">
              <component :is="Component" />
            </transition>
          </router-view>
        </el-tab-pane>
        <el-tab-pane label="分组预览" name="ftree">
          <template #label>
            <span>
              <icon-svg icon-class="promgtlanmei" />
              <!-- <i class="iconfont prometheusshuiguohuluobu"></i> -->
              分组预览</span
            >
          </template>
          <router-view name="ftree" v-slot="{ Component }">
            <transition name="slide-fade">
              <component :is="Component" />
            </transition>
          </router-view>
        </el-tab-pane>
        <el-tab-pane label="告警规则预览" name="ruleView">
          <template #label>
            <span>
              <!-- <i class="iconfont prometheusshuiguo2"></i> -->
              <icon-svg icon-class="promgtshengnvguo" />
              告警规则预览</span
            > </template
          ><router-view name="ruleView" v-slot="{ Component }">
            <transition name="slide-fade">
              <component :is="Component" />
            </transition>
          </router-view>
        </el-tab-pane>
      </el-tabs>
    </div>
    <el-backtop>
      <div
        style="
           {
            height: 100%;
            width: 100%;
            background-color: #f2f5f6;
            box-shadow: 0 0 6px rgba(0, 0, 0, 0.12);
            text-align: center;
            line-height: 40px;
            color: #1989fa;
          }
        "
      >
        UP
      </div>
    </el-backtop>
  </div>
</template>

<script>

import {ctlCreate, ctlReloadPrometheus, ctlCreateAndReload, getPrometheusUrl} from '@/api/control.js'
import SingleSvg from '../components/singleSvg.vue'

export default {
  name: "Menu",
  data () {
    SingleSvg
    return {
      activeTabName: 'jobs',
      transitionName: 'slide-left',
      scrollBackShow: false,
      dialogVisible: false
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
    window.addEventListener('scroll', this.scrollToTop)
    this.$router.push({ name: 'jobs' })
  },
  activated () {
    console.log('activated')
    this.$router.push({ name: 'ipManager' })
  },
  deactivated () {
    console.log('deactivated')
  },
  destroyed () {
    window.removeEventListener('scroll', this.scrollToTop);
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
      } else if (tab.instance.props.name === 'idc') {
        this.$router.push({ name: 'idc' })
      }
    },
    scrollToTop () {
      let scrollTop = window.pageYOffset || document.documentElement.scrollTop || document.body.scrollTop;
      let browserHeight = window.outerHeight;
      if (scrollTop > browserHeight) {
        this.scrollBackShow = true
      } else {
        this.scrollBackShow = false
      }
    },
    goBackPerson () {
      this.$router.push({ name: 'person' })
    },
    goControlCenter(){
      this.dialogVisible = true
    },
    handleClose(){
      this.dialogVisible = false
    },
    open_prometheus_self(){
      getPrometheusUrl().then(r=>{
        window.open(r.data, '_blank').location;
      }).catch(e=> console.log(e))
    },
    create(){
      ctlCreate().then(
        r => {
          this.$notify({
            title: '成功',
            message: '重新创建配置成功！',
            type: 'success'
          })
        }
      ).catch(
        e => { console.log(e) }
      )
    },
    reload(){
      ctlReloadPrometheus().then(
        r => {
          this.$notify({
            title: '成功',
            message: '重新加载配置成功！',
            type: 'success'
          })
        }
      ).catch(
        e => { console.log(e) }
      )
    },
    createAreload(){
      ctlCreateAndReload().then(
        r => {
          this.$notify({
            title: '成功',
            message: '重新创建并加载配置成功！',
            type: 'success'
          })
        }
      ).catch(
        e => { console.log(e) }
      )
    }
  }
}
</script>

<style scoped>
.main-board {
  /* position: absolute; */
  padding: 0;
  max-width: 99%;
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
.person-center-btn {
  position: relative;
  align-items: center;
  text-align: center;
  display: flex;
  justify-content: center;
  height: 38px;
  top: 0;
  float: right;
  z-index: 100;
}
</style>