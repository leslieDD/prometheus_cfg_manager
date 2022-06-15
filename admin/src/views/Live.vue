<template>
  <div class="admin-board">
    <div class="person-center-btn">
      <div>
        <el-button size="mini" type="warning" @click="goBackPerson">
          <single-svg icon-class="pig" />个人中心</el-button
        >
      </div>
    </div>
    <el-tabs type="border-card">
      <el-tab-pane label="直播测试">
        <template #label>
          <span>
            <!-- <i class="iconfont prometheusshuiguo1"></i> -->
            <icon-svg icon-class="promgtshuiguo9" />
            直播测试</span
          >
        </template>
        <div class="base-config-board">
          <div class="base-config-menu">
            <el-menu
              default-active="FLV"
              class="el-menu-vertical-demo"
              @open="handleOpen"
              @close="handleClose"
              @select="handleSelect"
            >
              <el-menu-item index="HLS">
                <i class="el-icon-potato-strips"></i>
                <template #title>HLS</template>
              </el-menu-item>
              <el-menu-item index="FLV">
                <i class="el-icon-milk-tea"></i>
                <template #title>FLV</template>
              </el-menu-item>
              <el-menu-item index="RTMP">
                <i class="el-icon-chicken"></i>
                <template #title>RTMP</template>
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
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script>
export default {
  name: "Admin",
  data () {
    return {

    }
  },
  mounted () {
    this.$router.push({ name: 'flv' })
  },
  methods: {
    handleOpen (key, keyPath) {
      console.log(key, keyPath);
    },
    handleClose (key, keyPath) {
      console.log(key, keyPath);
    },
    handleSelect (key, keyPath) {
      if (key === 'HLS') {
        this.$router.push({ name: 'hls' })
      } else if (key === 'FLV') {
        this.$router.push({ name: 'flv' })
      } else if (key === 'RTMP') {
        this.$router.push({ name: 'rtmp' })
      }
    },
    goBackPerson () {
      this.$router.push({ name: 'person' })
    }
  }
}
</script>

<style scoped>
.admin-board {
  /* position: absolute; */
  padding: 0;
  max-width: 99%;
  margin: 0 auto;
}
.base-config-menu {
  width: 300px;
  /* margin-right: 10px; */
}
.base-config-router {
  width: 100%;
  height: 89vh;
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