<template>
  <div class="login-board">
    <div class="login-title">
      <h2>Prometheus配置管理平台</h2>
    </div>
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <single-svg icon-class="earch" />
          <span> 登 录</span>
        </div>
      </template>
      <el-form
        :model="ruleForm"
        :rules="rules"
        ref="ruleForm"
        label-width="auto"
        class="demo-ruleForm"
        size="small"
      >
        <el-form-item label="用户名：" prop="username">
          <el-input
            clearable
            placeholder="请输入用户名"
            v-model="ruleForm.username"
          ></el-input>
        </el-form-item>
        <el-form-item label="密码：" prop="password">
          <el-input
            clearable
            type="password"
            placeholder="请输入密码"
            v-model="ruleForm.password"
          ></el-input>
        </el-form-item>
        <el-form-item>
          <div class="action-area">
            <el-button @click="doLogin('ruleForm')" type="primary" width="300px"
              >登录</el-button
            >
          </div>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script>

import { login } from '@/api/login.js'
import { setToken } from '@/utils/auth.js'

export default {
  data () {
    return {
      ruleForm: {
        username: '',
        password: ''
      },
      rules: {
        username: [
          { required: true, message: '请输入用户名', trigger: ['blur'] }
        ],
        password: [
          { required: true, message: '请输入密码', trigger: ['blur'] }
        ]
      }
    }
  },
  methods: {
    doLogin (formName) {
      this.$refs[formName].validate((valid) => {
        const logInfo = { ...this.ruleForm }
        if (valid) {
          login(logInfo).then(r => {
            setToken(r.data)
            this.$store.dispatch('setUserInfo', r.data)
            // console.log('token is ', this.$store.getters.token)
            this.$router.push({ name: 'person' })
          }).catch(e => console.log(e))
        }
      })
    }
  }
}
</script>

<style scoped>
.login-board {
  /* border: 1px solid red; */
  width: 400px;
  height: 450px;
  position: absolute;
  left: 50%;
  margin-left: -200px;
  top: 50%;
  margin-top: -280px;
}
.login-title {
  text-align: center;
  margin-bottom: 30px;
}
.action-area {
  display: flex;
  justify-content: space-between;
}
</style>