<template>
  <div class="base-config-board">
    <div class="base-config-menu">
      <el-menu
        default-active="1"
        class="el-menu-vertical-demo"
        @open="handleOpen"
        @close="handleClose"
        @select="handleSelect"
      >
        <el-menu-item index="1">
          <i class="el-icon-orange"></i>
          <template #title>选项编辑</template>
        </el-menu-item>
        <el-menu-item index="2">
          <i class="el-icon-grape"></i>
          <template #title>公共标签</template>
        </el-menu-item>
        <el-menu-item index="3">
          <i class="el-icon-watermelon"></i>
          <template #title>标签重写</template>
        </el-menu-item>
        <el-menu-item index="4">
          <i class="el-icon-cherry"></i>
          <template #title>默认分组</template>
        </el-menu-item>
        <el-menu-item index="5">
          <i class="el-icon-apple"></i>
          <template #title>测试配置</template>
        </el-menu-item>
        <el-menu-item index="6">
          <i class="el-icon-pear"></i>
          <template #title>模板编辑</template>
        </el-menu-item>
        <el-menu-item index="7">
          <i class="el-icon-lollipop"></i>
          <template #title>重置</template>
        </el-menu-item>
      </el-menu>
    </div>
    <div class="base-config-router">
      <router-view v-slot="{ Component }">
        <transition name="slide-fade">
          <component :is="Component" />
        </transition>
      </router-view>
      <!-- <router-view /> -->
    </div>
  </div>
</template>

<script>
export default {
  name: "GlobalConfig",
  data () {
    return {

    }
  },
  mounted () {
    this.$router.push({ name: 'options' })
  },
  methods: {
    handleOpen (key, keyPath) {
      console.log(key, keyPath);
    },
    handleClose (key, keyPath) {
      console.log(key, keyPath);
    },
    handleSelect (key, keyPath) {
      if (key === '1') {
        this.$router.push({ name: 'options' })
      } else if (key === '2') {
        this.$router.push({ name: 'baseLabels' })
      } else if (key === '3') {
        this.$router.push({ name: 'reLabels' })
      } else if (key === '4') {
        this.$router.push({ name: 'defaultJobs' })
      } else if (key === '5') {
        this.$router.push({ name: 'checkYml' })
      } else if (key === '6') {
        this.$router.push({ name: 'editPrometheusYml' })
      } else if (key === '7') {
        this.$router.push({ name: 'empty' })
      }
    }
  }
}
</script>

<style scoped>
.base-config-menu {
  width: 300px;
  /* margin-right: 10px; */
}
.base-config-router {
  width: 100%;
  height: 85vh;
}
.base-config-board {
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
  width: 100%;
  /* justify-content: space-between; */
}
.slide-fade-enter-active {
  transition: all 0.3s ease-in;
}

/* .slide-fade-leave-active {
  transition: all 0.8s cubic-bezier(1, 0.5, 0.8, 1);
} */

.slide-fade-enter-from,
.slide-fade-leave-to {
  /* transition: all 0.5s ease-out; */
  transform: translateX(10px);
  opacity: 0;
}
</style>