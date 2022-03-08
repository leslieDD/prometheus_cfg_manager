<template>
  <div class="user-board">
    <div class="main-content">
      <div>
        <div class="do_action">
          <div style="padding-right: 15px">
            <!-- <el-button size="small" type="success" plain @click="doAdd()"
              >添加用户</el-button
            > -->
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
              <el-table-column label="会话Token" prop="username">
              </el-table-column>
              <el-table-column label="登录名" prop="nice_name"> </el-table-column>
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
              <el-table-column
                label="最后更新时间"
                width="135px"
                prop="update_at"
              >
                <template v-slot="{ row }">
                  <span>{{ parseTimeSelf(row.update_at) }}</span>
                </template>
              </el-table-column>
              <el-table-column label="操作" align="center" width="200px">
                <template v-slot="scope" align="center">
                  <div class="actioneara">
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
  justify-content: center;
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
