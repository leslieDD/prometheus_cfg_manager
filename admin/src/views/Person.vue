<template>
  <div class="person-board">
    <div id="box">
      <div class="meBox">
        <!-- 头像 -->
        <!-- <div class="headPhoto">
        </div> -->
        <div class="headPhoto">
          <!-- <el-image :src="src"></el-image> -->
          <!-- <img src="/imgs/xiaohei.jpg" /> -->
        </div>
        <div class="meBox-title">
          <p>您好，{{ userInfo.nice_name }}</p>
          <div class="fg"></div>
        </div>
        <div class="meBox-text">
          <el-descriptions title="" :column="1" size="">
            <el-descriptions-item>
              <template #label>
                <span class="cell-item">
                  <i class="el-icon-user"></i>
                  登录账号
                </span>
              </template>
              {{ userInfo.username }}
            </el-descriptions-item>
            <el-descriptions-item>
              <template #label>
                <span class="cell-item">
                  <i class="el-icon-mobile-phone"></i>
                  手机号
                </span>
              </template>{{
                  userInfo.phone
              }}
            </el-descriptions-item>
            <el-descriptions-item>
              <template #label>
                <span class="cell-item">
                  <i class="el-icon-office-building"></i>
                  所属组
                </span>
              </template>{{
                  userInfo.group_name
              }}
            </el-descriptions-item>
            <el-descriptions-item>
              <template #label>
                <span class="cell-item">
                  <i class="el-icon-alarm-clock"></i>
                  注册时间
                </span>
              </template>{{
                  parseTimeSelf(userInfo.register_time)
              }}
            </el-descriptions-item>
            <el-descriptions-item>
              <template #label>
                <span class="cell-item">
                  <i class="el-icon-alarm-clock"></i>
                  登录时间
                </span>
              </template>{{ parseTimeSelf(userInfo.login_time) }}
            </el-descriptions-item>
          </el-descriptions>
        </div>
        <div class="meBox-Button">
          <a href="javascript:void(0)" @click="changePassword()">更改密码</a>
          <a href="javascript:void(0)" @click="quitLog()">退出登录</a>
        </div>
      </div>
      <div id="cmdBox">
        <div class="cmd">
          <div class="click">
            <div class="red"></div>
            <div class="yellow"></div>
            <div class="green"></div>
          </div>
          <div class="title">
            <span>root - bash</span>
          </div>
          <div class="cmdText">
            <span style="color: rgb(0, 190, 0)">root@monitor:~</span>
            <span style="color: blue"># </span>
            <span style="color: rgb(39, 39, 39)">./tianqi.sh</span>
            <br />
            <iframe scrolling="no" src="https://tianqiapi.com/api.php?style=tz&skin=pitaya" frameborder="0" width="100%"
              height="30" allowtransparency="true"></iframe>
            <br />
            <span style="color: rgb(0, 190, 0)">root@monitor:~</span>
            <span style="color: blue"># </span>
            <span style="color: rgb(39, 39, 39)">uptime</span>
            <p>{{ uptime }}</p>
            <br />
            <span style="color: rgb(0, 190, 0)">root@monitor:~</span>
            <span style="color: blue"># </span>
            <span style="color: rgb(39, 39, 39)">cat ./quotes_by_famous_people.txt</span>
            <p v-for="(words, index) in sayWhat" :key="index">{{ words }}</p>
            <br />
            <span style="color: rgb(0, 190, 0)">root@monitor:~</span>
            <span style="color: blue"># </span>
            <span style="color: rgb(39, 39, 39)">sudo rm -rf /过去的自己/*</span>
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
            <span style="color: rgb(0, 190, 0)">root@monitor:~</span>
            <span style="color: blue"># </span>
            <span style="color: rgb(39, 39, 39)">awk '/menu/' ./prometheus.yml</span>
            <p class="prompt-text">点击下面菜单进入：</p>
            <dl>
              <dt v-if="menuShow.show_menu_prometheus_cfg_manager === true">
                <el-button type="text" @click="goToConfig">
                  <single-svg icon-class="xianrenzhang"></single-svg>Prometheus配置
                </el-button>
              </dt>
              <dt class="menu_zero_padding" v-if="menuShow.show_menu_administrator_cfg_manager === true">
                <single-svg icon-class="yezhi"></single-svg>
                <el-button type="text" @click="goToAdmin">用户及权限管理</el-button>
              </dt>
              <dt class="menu_zero_padding" v-if="menuShow.show_menu_hls === true">
                <single-svg icon-class="yezhi"></single-svg>
                <el-button type="text" @click="goToHLS">直播</el-button>
              </dt>
              <dt v-if="Object.keys(menuShow).length === 0">
                <div class="no_menu_show">
                  <el-tag type="warning" size="small">
                    <single-svg icon-class="shuzhi"></single-svg>没有可供显示的菜单
                  </el-tag>
                </div>
              </dt>
            </dl>
            <span style="color: rgb(0, 190, 0)">root@monitor:~</span>
            <span style="color: blue"># </span>
          </div>
        </div>
      </div>
    </div>
    <el-dialog title="更新我的密码" v-model="dialogVisible" width="450px" :before-close="handleClose">
      <el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="auto" class="demo-ruleForm" size="small">
        <el-form-item label="请输出旧密码" prop="old_pwd">
          <el-input placeholder="请输入旧密码" v-model="ruleForm.old_pwd" type="password"></el-input>
        </el-form-item>
        <el-form-item label="新密码" prop="new_pwd1">
          <el-input placeholder="请输入新密码" v-model="ruleForm.new_pwd1" type="password"></el-input>
        </el-form-item>
        <el-form-item label="再次输入新密码" prop="new_pwd2">
          <el-input placeholder="请输入新密码" v-model="ruleForm.new_pwd2" type="password"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button @click="handleClose">取 消</el-button>
          <el-button type="primary" @click="submitChangePassword('ruleForm')">确 定</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>
  </div>
</template>

<script>
// import { removeToken } from '@/utils/auth.js'
import { removeStorageUserInfo } from '@/utils/localStorage.js'
// import store from '@/store/index.js'
import '@/assets/css/term.css'
import { loadTxt, chgPasswd, logout, getMenuPriv } from '@/api/person.js'
export default {
  data () {
    var validatePass2 = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请再次输入密码'))
      } else if (value !== this.ruleForm.new_pwd1) {
        callback(new Error('两次输入密码不一致!'))
      } else {
        callback()
      }
    }
    return {
      src: '/imgs/xiaohei.jpg',
      uptime: 'Linux 4.15.0-150-generic Unknow SMP Sat Jul 3 13:37:31 UTC 2021 x86_64 x86_64 x86_64 GNU/Linux',
      programmer_said: ['选择你能够承担的，承担你已经选择的'],
      sayWhat: [],
      timer: null,
      userInfo: {
        username: '',
        nice_name: '',
        phone: '',
        group_name: '',
        register_time: '',
        login_time: ''
      },
      dialogVisible: false,
      ruleForm: {
        old_pwd: '',
        new_pwd1: '',
        new_pwd2: ''
      },
      rules: {
        old_pwd: [
          { required: true, message: '请输入旧密码', trigger: 'blur' }
        ],
        new_pwd1: [
          { required: true, message: '请输入新密码', trigger: 'blur' },
          // { pattern: /^(?![a-zA-Z]+$)(?![A-Z0-9]+$)(?![A-Z\W_]+$)(?![a-z0-9]+$)(?![a-z\W_]+$)(?![0-9\W_]+$)[a-zA-Z0-9\W_]{8,30}$/, message: '密码为数字，小写字母，大写字母，特殊符号 至少包含三种，长度为 8 - 30位' }
        ],
        new_pwd2: [
          { required: true, validator: validatePass2, trigger: 'blur' }
        ],
      },
      userID: 0,
      menuShow: {
        show_menu_prometheus_cfg_manager: false,
        show_menu_administrator_cfg_manager: false,
        show_menu_hls: false,
      }
    }
  },
  mounted () {
    this.initUserInfo()
    this.goMenu()
    this.switchSay()
    if (this.timer === null) {
      this.setTimeer()
    }
    this.loadProgramerSay(this)
  },
  beforeUnmount () {
    clearInterval(this.timer)
    this.timer = null
  },
  methods: {
    initUserInfo () {
      if (this.$store.getters.userInfo) {
        this.userID = this.$store.getters.userID
        const userInfo = this.$store.getters.userInfo
        this.userInfo = {
          id: this.$store.getters.userID,
          username: userInfo.username,
          nice_name: userInfo.nice_name,
          phone: userInfo.phone,
          group_name: userInfo.group_name,
          register_time: userInfo.create_at,
          login_time: userInfo.session.login_at
        }
      }
    },
    goMenu () {
      getMenuPriv().then(r => {
        this.menuShow = r.data
      }).catch(e => {
        console.log(e)
      })
    },
    goToConfig () {
      this.$router.push({ name: 'menu' })
    },
    goToAdmin () {
      this.$router.push({ name: 'admin' })
    },
    goToHLS () {
      this.$router.push({ name: 'live' })
    },
    loadProgramerSay () {
      loadTxt().then(r => {
        this.programmer_said = r.data.programer_say
        this.uptime = r.data.uname
        this.switchSay()
      }).catch(e => console.log(e))
    },
    setTimeer () {
      this.timer = setInterval(() => {
        this.switchSay();//你所加载数据的方法
      }, 1000 * 15)
    },
    switchSay () {
      const sayWhat = this.programmer_said[Math.floor(Math.random() * this.programmer_said.length)]
      // this.sayWhat = this.fixedLengthFormatString(sayWhat, 28)
      // console.log('sayWhat', sayWhat, this.sayWhat)
      this.sayWhat = [sayWhat]
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
    },
    parseTimeSelf (t) {
      var time = new Date(Date.parse(t))
      return time.toLocaleDateString() + ' ' + time.toTimeString().split(' ')[0]
    },
    changePassword () {
      this.dialogVisible = true
    },
    quitLog () {
      this.$confirm('退出账号, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        logout().then(r => {
          this.$notify({
            title: '成功',
            message: '退出成功！',
            type: 'success',
            duration: 1000,
          });
          this.$store.dispatch('resetToken')
          // removeToken()
          removeStorageUserInfo()
          // location.reload()
          this.$router.push({ name: 'login' })
        })
      }).catch(() => {
        this.$notify({
          title: '消息',
          message: '取消退出',
          type: 'info',
          duration: 1000,
        });
        // this.$message({
        //   type: 'info',
        //   message: '取消退出'
        // });
      });
    },
    handleClose () {
      this.dialogVisible = false
    },
    submitChangePassword (formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          const postData = { ...this.ruleForm }
          chgPasswd(postData).then(r => {
            this.$notify({
              title: '成功',
              message: '密码更新成功！',
              type: 'success',
              duration: 1000,
            });
            this.dialogVisible = false
          }).catch(e => console.log(e))
        }
      })
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

.menu_zero_padding :deep() .el-button--text {
  padding-top: 0px;
}

.no_menu_show {
  padding: 10px;
}

.prompt-text {
  /* font-size: 14px; */
  font-weight: bold;
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