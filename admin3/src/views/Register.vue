<template>
  <div class="login-board">
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <span> 注 册</span>
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
        <el-form-item label="用户账号：" prop="username">
          <el-input
            clearable
            placeholder="请输入登录账号"
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
        <el-form-item label="用户姓名：" prop="nice_name">
          <el-input
            clearable
            placeholder="请输入用户姓名"
            v-model="ruleForm.nice_name"
          ></el-input>
        </el-form-item>
        <el-form-item label="手机号：" prop="phone">
          <el-input
            clearable
            placeholder="请输入手机号"
            v-model="ruleForm.phone"
          ></el-input>
        </el-form-item>
        <el-form-item>
          <div class="action-area">
            <el-button
              @click="doRegister('ruleForm')"
              type="warning"
              width="300px"
              >&nbsp;&nbsp;注&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;册&nbsp;&nbsp;</el-button
            >
            <el-button @click="doLogin('ruleForm')" type="primary" width="300px"
              >&nbsp;&nbsp;登&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;录&nbsp;&nbsp;</el-button
            >
          </div>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script>

import { register } from '@/api/login.js'

export default {
  data () {
    return {
      ruleForm: {
        username: '',
        nice_name: '',
        password: ''
      },
      rules: {
        username: [
          { required: true, message: '请输入用户账号', trigger: ['blur'] }
        ],
        nice_name: [
          { required: true, message: '请输入用户名称', trigger: ['blur'] }
        ],
        password: [
          { required: true, message: '请输入密码', trigger: ['blur'] }
        ]
      },
      userInfo: null,
    }
  },
  methods: {
    doRegister (formName) {
      this.$refs[formName].validate((valid) => {
        const logInfo = { ...this.ruleForm }
        if (valid) {
          register(logInfo).then(r => {
            this.userInfo = r.data
            this.$notify({
              title: '成功',
              message: '注册成功！',
              type: 'success'
            });
          }).catch(e => console.log(e))
        }
      })
    },
    doLogin () {
      this.$router.push({ name: 'login' })
    }
  }
}
</script>

<style scoped>
.login-board {
  /* border: 1px solid red; */
  width: 500px;
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
  margin-top: 10px;
  margin-left: -30px;
  margin-right: 20px;
  display: flex;
  justify-content: space-between;
}
</style>