<template>
  <div class="admin-set-board">
    <el-descriptions class="margin-top" :column="1" size="mini" border>
      <el-descriptions-item>
        <template #label>
          <i class="el-icon-s-custom"></i>
          当在调整用户所属组时，未被设置的用户默认指向的组
        </template>
        <el-select
          size="small"
          v-model="params.default_group"
          clearable
          placeholder="请选择"
          @visible-change="doPutAdminSetting"
          @clear="doPutAdminSetting"
          @change="selectChange"
        >
          <el-option
            v-for="item in groupList"
            :key="item.id"
            :label="item.name"
            :value="item.id"
          >
          </el-option>
        </el-select>
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label>
          <i class="el-icon-refresh-left"></i>
          重置用户及权限管理成出产设置
        </template>
        <el-button size="small" type="danger" @click="showResetDialog">
          重置数据
        </el-button>
      </el-descriptions-item>
    </el-descriptions>
    <div>
      <el-dialog title="输入重置码" v-model="dialogFormVisible" width="430px">
        <div class="reset-explain">
          <p>系統会在后台生成一个重置码，可以在服务的日志中看到，如：</p>
          <p style="font-width: bold">
            reset <span style="color: blue">admin</span> config key:
            <span style="color: red">6eb96669-2fb1-4758-8776-ddcdb58405c9</span>
          </p>
          <p>输入红色部分的KEY即可重置成出产设置</p>
        </div>
        <el-form
          :model="ruleForm"
          :rules="rules"
          ref="ruleForm"
          label-width="auto"
          class="demo-ruleForm"
          size="mini"
        >
          <el-form-item label="重置码" prop="code">
            <el-input v-model="ruleForm.code"></el-input>
          </el-form-item>
          <el-form-item>
            <div class="dialog-footer">
              <el-button size="small" @click="closeDialog('ruleForm')"
                >关 闭</el-button
              >
              <el-button
                v-if="runningMod === true"
                size="small"
                type="primary"
                icon="el-icon-loading"
                >确 定</el-button
              >
              <el-button
                v-if="runningMod === false"
                size="small"
                type="primary"
                icon="el-icon-refresh"
                @click="doResetAdmin('ruleForm')"
                >确 定</el-button
              >
            </div>
          </el-form-item>
        </el-form>
      </el-dialog>
    </div>
  </div>
</template>

<script>

import {
  getGroupList,
  getAdminSetting,
  putAdminSetting,
  resetAdmin,
  preResetAdmin
} from '@/api/manager.js'

export default {
  data () {
    return {
      params: {},
      groupList: [],
      dialogFormVisible: false,
      ruleForm: {
        code: ''
      },
      rules: {
        code: [
          { required: true, message: '请输入重置码！', trigger: 'blur' }
        ],
      },
      runningMod: false
    }
  },
  mounted () {
    this.selectChanged = false
    this.doGetParams()
  },
  methods: {
    doGetParams () {
      getGroupList().then(r => {
        this.groupList = r.data
        getAdminSetting().then(r => {
          Object.keys(r.data).forEach(k => {
            if (k === 'default_group') {
              var groupIndex = this.groupList.findIndex((item, index) => {
                return item.id.toString() === r.data[k];
              })
              if (groupIndex === -1) {
                this.params[k] = ''
              } else {
                this.params[k] = this.groupList[groupIndex].name
              }
            }
          })
        }).catch(e => console.log(e))
      }).catch(e => console.log(e))
    },
    doPutAdminSetting (action) {
      if (!this.selectChanged) {
        return
      }
      if (!action) {
        let params = {}
        Object.keys(this.params).forEach(k => {
          params[k] = this.params[k].toString()
        })
        putAdminSetting(params).then(r => {
          this.$notify({
            title: '成功',
            message: '参数提交成功',
            type: 'success'
          })
          this.selectChanged = false
        }
        ).catch(e => console.log(e))
      }
    },
    selectChange () {
      this.selectChanged = true
    },
    showResetDialog () {
      preResetAdmin().then(r => {
        this.$notify({
          title: '警告',
          message: '已经生成新的重置码！',
          type: 'warning'
        })
        this.ruleForm.code = ''
        this.dialogFormVisible = true
      })
    },
    closeDialog (formName) {
      this.dialogFormVisible = false
    },
    doResetAdmin (formName) {
      this.runningMod = true
      this.$refs[formName].validate((valid) => {
        if (valid) {
          resetAdmin(this.ruleForm).then(r => {
            this.$notify({
              title: '成功',
              message: '重置成功！',
              type: 'success'
            });
            this.runningMod = false
            this.doGetOperateLog()
          }).catch(e => {
            console.log(e)
            this.runningMod = false
            this.doGetOperateLog()
          })
        } else {
          this.runningMod = false
          console.log(`validate: ${valid}`)
        }
      })
    }
  }
}
</script>

<style scoped>
.admin-set-board {
  width: 100%;
}
.operate-area {
  display: flex;
  flex-direction: row;
  justify-content: flex-end;
}
.dialog-footer {
  text-align: right;
  align-content: right;
}
.reset-explain {
  font: 0.8em Arial, Tahoma, Verdana;
  color: #777;
  margin-top: -30px;
  margin-bottom: 12px;
}
</style>