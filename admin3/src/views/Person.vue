<template>
  <div class="person-board">
    <div id="box">
      <!-- 个人资料卡片 -->
      <div class="meBox">
        <!-- 头像 -->
        <!-- <div class="headPhoto">

        </div> -->
        <div class="headPhoto">
          <!-- <el-image :src="src"></el-image> -->
          <!-- <img src="/imgs/xiaohei.jpg" /> -->
        </div>
        <!-- 介绍 -->
        <div class="meBox-title">
          <p>I'm DianDian</p>
          <div class="fg"></div>
        </div>
        <div class="meBox-text">
          <el-descriptions title="" :column="1" size="">
            <el-descriptions-item label="手机号"
              >18100000000</el-descriptions-item
            >
            <el-descriptions-item label="所属组">苏州市</el-descriptions-item>
            <el-descriptions-item label="注册时间"
              >1986.09.24</el-descriptions-item
            >
            <el-descriptions-item label="登录时间"
              >1986.09.24</el-descriptions-item
            >
          </el-descriptions>
        </div>
        <!-- 两个按钮 -->
        <div class="meBox-Button">
          <a href="http://www.bootstrapmb.com">更改密码</a>
          <a href="http://www.bootstrapmb.com">更换头像</a>
        </div>
      </div>

      <!-- 伪终端介绍 -->
      <div id="cmdBox">
        <!-- 第一个终端 -->
        <div class="cmd">
          <!-- 三个按钮 -->
          <div class="click">
            <div class="red"></div>
            <div class="yellow"></div>
            <div class="green"></div>
          </div>
          <!-- 顶部标题 -->
          <div class="title">
            <span>root - bash</span>
          </div>
          <!-- 终端内文字 -->
          <div class="cmdText">
            <span style="color: rgb(0, 190, 0)">root@monitor</span>
            <span style="color: blue">~</span>
            <span style="color: rgb(39, 39, 39)">./tianqi.sh</span>
            <br />
            <iframe
              scrolling="no"
              src="https://tianqiapi.com/api.php?style=tc&skin=pitaya"
              frameborder="0"
              width="350"
              height="24"
              allowtransparency="true"
            ></iframe>
            <br />
            <span style="color: rgb(0, 190, 0)">root@monitor</span>
            <span style="color: blue">~</span>
            <span style="color: rgb(39, 39, 39)">uname</span>
            <p>{{ uname }}</p>
            <br />
            <span style="color: rgb(0, 190, 0)">root@monitor</span>
            <span style="color: blue">~</span>
            <span style="color: rgb(39, 39, 39)"
              >cat ./quotes_by_famous_people.txt</span
            >
            <p></p>
            <span style="color: rgb(0, 190, 0)">root@monitor</span>
            <span style="color: blue">~</span>
            <span style="color: rgb(39, 39, 39)"
              >sudo rm -rf /过去的自己/*</span
            >
          </div>
        </div>
        <!-- 第二个终端 -->
        <div class="cmd cmd2">
          <!-- 三个按钮 -->
          <div class="click">
            <div class="red"></div>
            <div class="yellow"></div>
            <div class="green"></div>
          </div>
          <!-- 顶部标题 -->
          <div class="title">
            <span>root - bash</span>
          </div>
          <!-- 终端内文字 -->
          <div class="cmdText">
            <span style="color: rgb(0, 190, 0)">root@monitor</span>
            <span style="color: blue">~</span>
            <span style="color: rgb(39, 39, 39)">vim ./prometheus.yml</span>
            <p>点我进入Prometheus的配置管理：</p>
            <ul class="ul">
              <li>
                <el-button type="text" @click="goToConfig"
                  >Prometheus配置</el-button
                >
              </li>
            </ul>
            <span style="color: rgb(0, 190, 0)">root@monitor</span>
            <span style="color: blue">~</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import '@/assets/css/term.css'
import { loadTxt } from '@/api/person.js'
export default {
  data () {
    return {
      src: '/imgs/xiaohei.jpg',
      uname: 'Linux 4.15.0-150-generic Unknow SMP Sat Jul 3 13:37:31 UTC 2021 x86_64 x86_64 x86_64 GNU/Linux',
      programmer_said: [],
      sayWhat: [],
      timer: null
    }
  },
  mounted () {
    this.setTimeer()
    this.loadProgramerSay(this)
  },
  beforeUnmount () {
    clearInterval(this.timer)
  },
  methods: {
    goToConfig () {
      this.$router.push({ name: 'menu' })
    },

    loadProgramerSay () {
      loadTxt().then(r => {
        this.programmer_said = r.data
        this.setTimeer()
      }).catch(e => console.log(e))
    },
    setTimeer () {
      this.timer = setInterval(() => {
        this.switchSay();//你所加载数据的方法
      }, 5000)
    },
    switchSay () {
      if (this.programmer_said.length === 0) {
        return
      }
      const sayWhat = Math.floor(Math.random() * this.programmer_said.length)
      this.sayWhat = this.fixedLengthFormatString(sayWhat, 28)
    },
    fixedLengthFormatString (str, num) {
      if (str == null || str == undefined) return null;
      if (!(/^[0-9]*[1-9][0-9]*$/.test(num))) return null;
      var array = new Array();
      var len = str.length;
      for (var i = 0; i < (len / num); i++) {
        if ((i + 1) * num > len) {
          array.push(str.substring(i * num, len));
        } else {
          array.push(str.substring(i * num, (i + 1) * num));
        }
      }
      return array;
    }
  }
}
</script>

<style scoped>
.person-board {
  padding: 0;
  max-width: 1100px;
  margin: 0 auto;
}
.person-title {
  font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB",
    "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
  font-size: 25px;
}
.person-area {
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
  justify-content: space-between;
}
.person-info {
  width: 40%;
}
.person-action {
  width: 50%;
}
.person-hr {
  width: 30px;
}
.go-to-config {
  float: right;
}
</style>