<template>
  <div class="group-board">
    <div class="main-content">
      <div>
        <div class="do_action">
          <div style="padding-right: 15px">
            <el-button size="small" type="success" plain @click="doAdd()"
              >添加组</el-button
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
          <el-table
            size="mini"
            highlight-current-row
            border
            :data="UserGroup"
            stripe
            :row-style="rowStyle"
            :cell-style="cellStyle"
          >
            <el-table-column label="序号" width="50px">
              <template v-slot="scope">
                {{ scope.$index + 1 }}
              </template>
            </el-table-column>
            <el-table-column label="组名" prop="name"> </el-table-column>
            <el-table-column label="用户数" prop="user_count">
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
                        <el-button size="mini" type="text" @click="doNo(scope)"
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
        </div>
      </div>
    </div>
    <el-dialog
      :title="dialogTitle"
      v-model="dialogVisible"
      width="400px"
      modal
      :before-close="handleClose"
    >
      <span>
        <el-form
          label-position="right"
          :rules="rules"
          ref="userGroupRef"
          :model="userGroupRef"
          label-width="auto"
          size="small"
        >
          <el-form-item label="组名：" prop="name">
            <el-input
              style="width: 280px"
              v-model="userGroupRef.name"
            ></el-input>
          </el-form-item>
          <el-form-item size="small">
            <el-button
              size="small"
              type="primary"
              @click="onSubmit('userGroupRef')"
              >{{ buttonTitle }}</el-button
            >
            <el-button
              size="small"
              type="info"
              @click="onCancel('userGroupRef')"
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
  getUserGroup,
  putUserGroup,
  postUserGroup,
  deleteUserGroup,
  enabledUserGroup
} from '@/api/usergroup.js'

export default {
  name: 'UserGroup',
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
      UserGroup: [],
      userGroupRef: {
        'id': 0,
        'name': '',
        'user_count': 0
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
        name: [
          { required: true, message: '请输入正确的名称', validator: validateStr, trigger: ['blur'] }
        ]
      },
      paginationShow: true
    }
  },
  mounted () {
    this.doGetUserGroup()
  },
  methods: {
    doAdd () {
      this.userGroupRef = {
        'id': 0,
        'name': ''
      }
      this.buttonTitle = '创建'
      this.dialogTitle = '增加组'
      this.dialogVisible = true
    },
    doGetUserGroup (getInfo) {
      if (!getInfo) {
        getInfo = {
          'pageNo': this.currentPage,
          'pageSize': this.pageSize,
          'search': this.searchContent
        }
      }
      getUserGroup(getInfo).then(
        r => {
          this.UserGroup = r.data.data
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
    doEdit (data) {
      this.buttonTitle = '更新'
      this.dialogTitle = '编辑组名'
      this.userGroupRef.id = data.row.id
      this.userGroupRef.name = data.row.name
      this.dialogVisible = true
    },
    handleSizeChange (val) {
      let getInfo = {
        'pageNo': this.currentPage,
        'pageSize': val,
        'search': this.searchContent
      }
      this.doGetUserGroup(getInfo)
    },
    handleCurrentChange (val) {
      let getInfo = {
        'pageNo': val,
        'pageSize': this.pageSize,
        'search': this.searchContent
      }
      this.doGetUserGroup(getInfo)
    },
    handleClose (done) {
      this.dialogVisible = false
    },
    onSubmit (formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          let postData = {}
          postData['name'] = this.userGroupRef.name
          if (this.buttonTitle === '创建' || this.buttonTitle === 'create') {
            postUserGroup(postData).then(
              r => {
                this.$notify({
                  title: '成功',
                  message: '创建成功！',
                  type: 'success'
                });
                this.doGetUserGroup()
                this.dialogVisible = false
                this.$refs[formName].resetFields()
              }
            ).catch(
              e => { console.log(e) }
            )
          } else {
            postData['id'] = this.userGroupRef.id
            putUserGroup(postData).then(
              r => {
                this.$notify({
                  title: '成功',
                  message: '更新成功！',
                  type: 'success'
                });
                this.doGetUserGroup()
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
      this.doGetUserGroup()
    },
    parseTimeSelf (t) {
      var time = new Date(Date.parse(t))
      return time.toLocaleDateString() + ' ' + time.toTimeString().split(' ')[0]
    },
    doYes (scope) {
      deleteUserGroup({ id: scope.row.id }).then(
        r => {
          this.$notify({
            title: '成功',
            message: '删除成功！',
            type: 'success'
          });
          this.doGetUserGroup()
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
      const newStatus = !this.UserGroup[scope.$index].enabled
      const rInfo = {
        id: scope.row.id,
        enabled: newStatus
      }
      enabledUserGroup(rInfo).then(r => {
        this.$notify({
          title: '成功',
          message: '更新状态成功！',
          type: 'success'
        });
        this.UserGroup[scope.$index].enabled = newStatus
        this.UserGroup = [...this.UserGroup]
      }).catch(e => console.log(e))
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
