<template>
  <div class="user-board">
    <div class="main-content">
      <div>
        <div class="do_action">
          <div style="padding-right: 15px">
            <el-button size="small" type="success" plain @click="doAdd()"
              >添加用户</el-button
            >
          </div>
          <div>
            <el-input
              size="small"
              placeholder="请输入内容"
              @keyup.enter="onSearch()"
              v-model="searchContent"
            >
              <template #append>
                <el-button
                  size="small"
                  @click="onSearch()"
                  icon="el-icon-search"
                ></el-button>
              </template>
            </el-input>
          </div>
        </div>
        <div class="table-show">
          <el-scrollbar height="75vh" class="flex-content">
            <el-table
              size="mini"
              highlight-current-row
              border
              :data="ManagerUser"
              stripe
              :row-style="rowStyle"
              :cell-style="cellStyle"
            >
              <el-table-column label="序号" width="50px">
                <template v-slot="scope">
                  {{ scope.$index + 1 }}
                </template>
              </el-table-column>
              <el-table-column label="登录名" prop="username">
              </el-table-column>
              <el-table-column
                label="密码"
                show-overflow-tooltip
                prop="password"
              >
              </el-table-column>
              <el-table-column label="姓名" prop="nice_name"> </el-table-column>
              <el-table-column label="手机" prop="phone"> </el-table-column>
              <el-table-column label="所属组" prop="group_name">
                <template v-slot="{ row }">
                  <el-button
                    size="mini"
                    @click="routeToGroup(row)"
                    type="text"
                    >{{ row.group_name }}</el-button
                  >
                </template>
              </el-table-column>
              <el-table-column label="最后更新账号" prop="update_by">
              </el-table-column>
              <el-table-column label="最后更新时间" prop="update_at">
                <template v-slot="{ row }">
                  <span>{{ parseTimeSelf(row.update_at) }}</span>
                </template>
              </el-table-column>
              <el-table-column label="操作" align="center" width="200px">
                <template v-slot="scope" align="center">
                  <div class="actioneara">
                    <div>
                      <el-button
                        size="mini"
                        type="primary"
                        @click="doEdit(scope)"
                        plain
                        :disabled="editCodeButVisable[scope.row.id]"
                        >编辑</el-button
                      >
                    </div>
                    <div>
                      <el-button
                        size="mini"
                        type="info"
                        v-if="scope.row.enabled === true"
                        @click="invocate(scope)"
                        plain
                        :disabled="editCodeButVisable[scope.row.id]"
                        >禁用</el-button
                      >
                      <el-button
                        size="mini"
                        type="warning"
                        v-if="scope.row.enabled === false"
                        @click="invocate(scope)"
                        plain
                        :disabled="editCodeButVisable[scope.row.id]"
                        >启用</el-button
                      >
                    </div>
                    <div>
                      <el-popover
                        :visible="deleteVisible[scope.$index]"
                        placement="top"
                      >
                        <p>确定删除吗？</p>
                        <div style="text-align: right; margin: 0">
                          <el-button
                            size="mini"
                            type="text"
                            @click="doNo(scope)"
                            >取消</el-button
                          >
                          <el-button
                            type="primary"
                            size="mini"
                            @click="doYes(scope)"
                            >确定</el-button
                          >
                        </div>
                        <template #reference>
                          <el-button
                            size="mini"
                            type="danger"
                            plain
                            @click="doDelete(scope)"
                            :disabled="editCodeButVisable[scope.row.id]"
                            >删除</el-button
                          >
                        </template>
                      </el-popover>
                    </div>
                  </div>
                </template>
              </el-table-column>
            </el-table>
            <div class="block" v-if="paginationShow">
              <el-pagination
                @size-change="handleSizeChange"
                @current-change="handleCurrentChange"
                :current-page="currentPage"
                :page-sizes="[20, 30, 50]"
                :page-size="pageSize"
                layout="total, sizes, prev, pager, next, jumper"
                :total="pageTotal"
              >
              </el-pagination>
            </div>
          </el-scrollbar>
        </div>
      </div>
    </div>
    <el-dialog
      :title="dialogTitle"
      v-model="dialogVisible"
      width="510px"
      modal
      :before-close="handleClose"
    >
      <span>
        <el-form
          label-position="right"
          :rules="rules"
          ref="ManagerUserRef"
          :model="ManagerUserRef"
          label-width="auto"
          size="small"
        >
          <el-form-item label="用户登录名：" prop="username">
            <el-input
              style="width: 340px"
              v-model="ManagerUserRef.username"
            ></el-input>
          </el-form-item>
          <el-form-item label="登录密码：" prop="password">
            <el-input
              type="password"
              style="width: 340px"
              v-model="ManagerUserRef.password"
            ></el-input>
          </el-form-item>
          <el-form-item label="用户姓名：" prop="nice_name">
            <el-input
              style="width: 340px"
              v-model="ManagerUserRef.nice_name"
            ></el-input>
          </el-form-item>
          <el-form-item label="手机：" prop="phone">
            <el-input
              style="width: 340px"
              v-model="ManagerUserRef.phone"
            ></el-input>
          </el-form-item>
          <el-form-item label="所属组：" prop="group_id">
            <el-select
              v-model="ManagerUserRef.group_id"
              class="borderNone"
              popper-class="pppselect"
              size="small"
              collapse-tags
              placeholder="请选择"
              style="width: 340px"
            >
              <el-option
                style="width: 340px"
                v-for="item in ManagerGroup"
                :key="item.id"
                :label="item.name"
                :value="item.id"
              >
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item size="small">
            <el-button
              size="small"
              type="primary"
              @click="onSubmit('ManagerUserRef')"
              >{{ buttonTitle }}</el-button
            >
            <el-button
              size="small"
              type="info"
              @click="onCancel('ManagerUserRef')"
              >取消</el-button
            >
          </el-form-item>
        </el-form>
      </span>
    </el-dialog>
  </div>
</template>

<script>

import {
  getManagerUser,
  putManagerUser,
  postManagerUser,
  deleteManagerUser,
  enabledManagerUser,
  getManangerGroupEnabled
} from '@/api/manager.js'

export default {
  name: 'ManagerUser',
  data () {
    function validateStr (rule, value, callback) {
      if (value === '' || typeof value === 'undefined' || value == null) {
        callback(new Error('请输入正确名称'))
      } else {
        const reg = /\s+/
        if (reg.test(value)) {
          callback(new Error('名称中不允许有空字符'))
        } else {
          callback()
        }
      }
    }
    return {
      value: '',
      ManagerGroup: [],
      ManagerUser: [],
      ManagerUserRef: {
        id: 0,
        username: '',
        password: '',
        group_id: null,
        phone: ''
      },
      editCodeButVisable: {},
      pageSize: 20,
      pageTotal: 0,
      currentPage: 1,
      searchContent: '',
      deleteVisible: {},
      dialogVisible: false,
      buttonTitle: '',
      dialogTitle: '',
      rules: {
        username: [
          { required: true, message: '请输入正确的登录名', validator: validateStr, trigger: ['blur'] }
        ],
        password: [
          { required: true, message: '请输入密码', trigger: ['blur'] }
        ],
        nice_name: [
          { required: true, message: '请输入正确的用户名称', validator: validateStr, trigger: ['blur'] }
        ],
        group_id: [
          { required: true, message: '请选择组', trigger: ['blur'] }
        ]
      },
      paginationShow: true
    }
  },
  mounted () {
    if (this.$route.params.group_name) {
      this.searchContent = this.$route.params.group_name
    }
    this.doGetManagerUser()
  },
  methods: {
    doAdd () {
      this.ManagerUserRef = {
        id: 0,
        username: '',
        password: '',
        nice_name: '',
        group_id: null,
        phone: ''
      }
      getManangerGroupEnabled().then(r => {
        this.ManagerGroup = r.data
        this.buttonTitle = '创建'
        this.dialogTitle = '增加用户'
        this.dialogVisible = true
      }).catch(e => console.log(e))
    },
    doGetManagerUser (getInfo) {
      if (!getInfo) {
        getInfo = {
          'pageNo': this.currentPage,
          'pageSize': this.pageSize,
          'search': this.searchContent
        }
      }
      getManagerUser(getInfo).then(
        r => {
          this.ManagerUser = r.data.data
          this.pageTotal = r.data.totalCount
          this.currentPage = r.data.pageNo
          this.pageSize = r.data.pageSize
          this.paginationShow = false;//让分页隐藏
          this.$nextTick(() => {//重新渲染分页
            this.paginationShow = true;
          });
        }
      ).catch(
        e => {
          console.log(e)
        }
      )
    },
    doEdit (scope) {
      getManangerGroupEnabled().then(r => {
        this.ManagerGroup = r.data
        this.buttonTitle = '更新'
        this.dialogTitle = '编辑用户'
        this.ManagerUserRef = { ...scope.row }
        this.dialogVisible = true
      }).catch(e => console.log(e))
    },
    handleSizeChange (val) {
      let getInfo = {
        'pageNo': this.currentPage,
        'pageSize': val,
        'search': this.searchContent
      }
      this.doGetManagerUser(getInfo)
    },
    handleCurrentChange (val) {
      let getInfo = {
        'pageNo': val,
        'pageSize': this.pageSize,
        'search': this.searchContent
      }
      this.doGetManagerUser(getInfo)
    },
    handleClose (done) {
      this.dialogVisible = false
    },
    onSubmit (formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          const postData = { ...this.ManagerUserRef }
          if (this.buttonTitle === '创建' || this.buttonTitle === 'create') {
            postManagerUser(postData).then(
              r => {
                this.$notify({
                  title: '成功',
                  message: '创建成功！',
                  type: 'success'
                });
                this.doGetManagerUser()
                this.dialogVisible = false
                this.$refs[formName].resetFields()
              }
            ).catch(
              e => { console.log(e) }
            )
          } else {
            // postData['id'] = this.ManagerUserRef.id
            putManagerUser(postData).then(
              r => {
                this.$notify({
                  title: '成功',
                  message: '更新成功！',
                  type: 'success'
                });
                this.doGetManagerUser()
                this.dialogVisible = false
                this.$refs[formName].resetFields()
              }
            ).catch(
              e => { console.log(e) }
            )
          }
        } else {
          return false
        }
      })
    },
    onCancel (formName) {
      this.dialogVisible = false
      this.$refs[formName].resetFields()
    },
    onSearch () {
      this.doGetManagerUser()
    },
    parseTimeSelf (t) {
      var time = new Date(Date.parse(t))
      return time.toLocaleDateString() + ' ' + time.toTimeString().split(' ')[0]
    },
    doYes (scope) {
      deleteManagerUser({ id: scope.row.id }).then(
        r => {
          this.$notify({
            title: '成功',
            message: '删除成功！',
            type: 'success'
          });
          this.doGetManagerUser()
        }
      ).catch(
        e => {
          console.log(e)
        }
      )
      this.deleteVisible[scope.$index] = false
    },
    doNo (scope) {
      this.deleteVisible[scope.$index] = false
    },
    doDelete (scope) {
      this.deleteVisible[scope.$index] = true
    },
    rowStyle (row) {
      let rs = {
        'padding': '0'
      }
      return rs
    },
    cellStyle (column) {
      let cs = {
        'padding': '0'
      }
      return cs
    },
    invocate (scope) {
      const newStatus = !this.ManagerUser[scope.$index].enabled
      const rInfo = {
        id: scope.row.id,
        enabled: newStatus
      }
      enabledManagerUser(rInfo).then(r => {
        this.$notify({
          title: '成功',
          message: '更新状态成功！',
          type: 'success'
        });
        this.ManagerUser[scope.$index].enabled = newStatus
        this.ManagerUser = [...this.ManagerUser]
      }).catch(e => console.log(e))
    },
    routeToGroup (row) {
      this.$router.push({ name: 'group', params: { group_name: row.group_name } })
    }
  }
};
</script>

<style scoped>
.main-box {
  height: 85vh;
}
.main-content {
  width: 100%;
  height: 40vh;
}
.do_action {
  display: flex;
  flex-direction: row;
  justify-content: flex-end;
  margin-bottom: 5px;
}
.actioneara {
  display: flex;
  flex-wrap: nowrap;
  justify-content: space-between;
}
.el-pagination {
  text-align: center;
}
.block {
  padding-top: 12px;
}
.cm-s-monokai {
  height: 100%;
}
</style>
