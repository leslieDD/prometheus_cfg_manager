<template>
  <div class="image-area">
    <div>
      <span class="vintage">是 不 是 偷 偷 改 URL 了 呢</span>
    </div>
    <div>
      <img src="/imgs/404.png" alt="404, page no found" />
    </div>
    <div class="switch-btn">
      <div class="btn-one">
        <el-button @click="goBackBefore">返回上一页</el-button>
      </div>
      <div class="btn-between"></div>
      <div class="btn-two">
        <el-button @click="goBackMainPage">返回首页</el-button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data () {
    return {}
  },
  methods: {
    goBackBefore () {
      this.$router.go(-1)
    },
    goBackMainPage () {
      this.$router.push({ name: "person" })
    }
  }
}
</script>

<style scoped>
.image-area {
  padding: 0 0;
  margin: 10% auto;
  height: 200px;
  width: 1100px;
  /* overflow: hidden; */
  align-content: center;
  text-align: center;
  align-items: center;
}
.vintage {
  background: #eee
    url(data:image/gif;base64,iVBORw0KGgoAAAANSUhEUgAAAAQAAAAECAIAAAAmkwkpAAAAHklEQVQImWNkYGBgYGD4//8/A5wF5SBYyAr+//8PAPOCFO0Q2zq7AAAAAElFTkSuQmCC)
    repeat;
  text-shadow: 5px -5px #a8a8a8, 4px -4px white;
  font-weight: bold;
  -webkit-text-fill-color: transparent;
  -webkit-background-clip: text;
  font-size: 25px;
}

.switch-btn {
  display: flex;
  justify-content: center;
  align-items: center;
  align-content: center;
}
.btn-between {
  width: 100px;
}
</style>