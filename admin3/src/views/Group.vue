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
          <el-scrollbar height="75vh" class="flex-content">
            <el-table
              size="mini"
              highlight-current-row
              border
              :data="ManagerGroup"
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
              <el-table-column label="用户数" prop="user_count" width="150px">
                <template v-slot="{ row }">
                  <el-button
                    size="mini"
                    @click="routeToUser(row)"
                    type="text"
                    >{{ row.user_count }}</el-button
                  >
                </template>
              </el-table-column>
              <el-table-column
                label="最后更新账号"
                align="center"
                prop="update_by"
              ></el-table-column>
              <el-table-column
                label="最后更新时间"
                width="150px"
                prop="update_at"
              >
                <template v-slot="{ row }">
                  <span>{{ parseTimeSelf(row.update_at) }}</span>
                </template>
              </el-table-column>
              <el-table-column label="操作" align="center" width="350px">
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
                        type="primary"
                        @click="doUserList(scope)"
                        plain
                        >用户列表</el-button
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
                      <el-button
                        size="mini"
                        type="primary"
                        @click="doPriv(scope)"
                        plain
                        >权限</el-button
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
      width="400px"
      modal
      :before-close="handleClose"
    >
      <span>
        <el-form
          label-position="right"
          :rules="rules"
          ref="ManagerGroupRef"
          :model="ManagerGroupRef"
          label-width="auto"
          size="small"
        >
          <el-form-item label="组名：" prop="name">
            <el-input
              style="width: 280px"
              v-model="ManagerGroupRef.name"
            ></el-input>
          </el-form-item>
          <el-form-item size="small">
            <el-button
              size="small"
              type="primary"
              @click="onSubmit('ManagerGroupRef')"
              >{{ buttonTitle }}</el-button
            >
            <el-button
              size="small"
              type="info"
              @click="onCancel('ManagerGroupRef')"
              >取消</el-button
            >
          </el-form-item>
        </el-form>
      </span>
    </el-dialog>
    <el-dialog
      :title="'编辑组成员：' + editObject"
      v-model="editUserVisible"
      width="700px"
      modal
      :before-close="editUserClose"
    >
      <el-row type="flex" align="middle" justify="center">
        <el-transfer
          v-model="usersListValue"
          filterable
          :filter-method="filterUserMethod"
          filter-placeholder="请输入关键字"
          :data="usersListData"
          @change="transferChange"
          :titles="['用户列表', '组成员']"
        >
        </el-transfer>
      </el-row>
      <div class="user-list-push-box">
        <el-button
          class="user-list-close-btn"
          type="info"
          size="small"
          @click="editUserClose"
          >关闭</el-button
        >
        <el-button
          class="user-list-push-btn"
          type="warning"
          size="small"
          @click="pushNewUserList"
          >提交</el-button
        >
      </div>
    </el-dialog>
  </div>
</template>

<script>

import {
  getManagerGroup,
  putManagerGroup,
  postManagerGroup,
  deleteManagerGroup,
  enabledManagerGroup,
  getAllUserList,
  getAllUserForGroup,
  updateGroupMember
} from '@/api/manager.js'

export default {
  name: 'ManagerGroup',
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
      ManagerGroupRef: {
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
      paginationShow: true,
      editObject: '',
      editUserVisible: false,
      usersListData: [],
      transferChanged: false,
      usersListValue: [],
      groupInfo: {}
    }
  },
  mounted () {
    if (this.$route.params.group_name) {
      this.searchContent = this.$route.params.group_name
    }
    const pageInfo = this.$store.getters.getPageInfo
    this.pageSize = pageInfo.pageSize
    this.pageNo = pageInfo.pageNo
    this.doGetManagerGroup()
  },
  methods: {
    doAdd () {
      this.ManagerGroupRef = {
        'id': 0,
        'name': ''
      }
      this.buttonTitle = '创建'
      this.dialogTitle = '增加组'
      this.dialogVisible = true
    },
    doGetManagerGroup (getInfo) {
      if (!getInfo) {
        getInfo = {
          'pageNo': this.currentPage,
          'pageSize': this.pageSize,
          'search': this.searchContent
        }
      }
      getManagerGroup(getInfo).then(
        r => {
          this.ManagerGroup = r.data.data
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
      this.ManagerGroupRef.id = data.row.id
      this.ManagerGroupRef.name = data.row.name
      this.dialogVisible = true
    },
    doUserList (scope) {
      this.editObject = scope.row.name
      let usersListData = []
      // usersListData
      getAllUserList().then(r => {
        r.data.forEach(u => {
          usersListData.push({
            label: u.username,
            key: u.id,
            spell: u.username,
            disabled: false
          })
        })
        let usersListValue = []
        this.groupInfo = {
          group_name: scope.row.name,
          group_id: scope.row.id
        }
        getAllUserForGroup(this.groupInfo).then(r => {
          r.data.forEach(u => {
            usersListValue.push(u.id)
          })
          this.usersListData = usersListData
          this.usersListValue = usersListValue
          this.transferChanged = false
          this.editUserVisible = true
        }).catch(e => console.log(e))
      }).catch(e => console.log(e))
    },
    handleSizeChange (val) {
      let getInfo = {
        'pageNo': this.currentPage,
        'pageSize': val,
        'search': this.searchContent
      }
      this.doGetManagerGroup(getInfo)
    },
    handleCurrentChange (val) {
      let getInfo = {
        'pageNo': val,
        'pageSize': this.pageSize,
        'search': this.searchContent
      }
      this.doGetManagerGroup(getInfo)
    },
    handleClose (done) {
      this.dialogVisible = false
    },
    onSubmit (formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          let postData = {}
          postData['name'] = this.ManagerGroupRef.name
          if (this.buttonTitle === '创建' || this.buttonTitle === 'create') {
            postManagerGroup(postData).then(
              r => {
                this.$notify({
                  title: '成功',
                  message: '创建成功！',
                  type: 'success'
                });
                this.doGetManagerGroup()
                this.dialogVisible = false
                this.$refs[formName].resetFields()
              }
            ).catch(
              e => { console.log(e) }
            )
          } else {
            postData['id'] = this.ManagerGroupRef.id
            putManagerGroup(postData).then(
              r => {
                this.$notify({
                  title: '成功',
                  message: '更新成功！',
                  type: 'success'
                });
                this.doGetManagerGroup()
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
      this.doGetManagerGroup()
    },
    parseTimeSelf (t) {
      var time = new Date(Date.parse(t))
      return time.toLocaleDateString() + ' ' + time.toTimeString().split(' ')[0]
    },
    doYes (scope) {
      deleteManagerGroup({ id: scope.row.id }).then(
        r => {
          this.$notify({
            title: '成功',
            message: '删除成功！',
            type: 'success'
          });
          this.doGetManagerGroup()
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
      const newStatus = !this.ManagerGroup[scope.$index].enabled
      const rInfo = {
        id: scope.row.id,
        enabled: newStatus
      }
      enabledManagerGroup(rInfo).then(r => {
        this.$notify({
          title: '成功',
          message: '更新状态成功！',
          type: 'success'
        });
        this.ManagerGroup[scope.$index].enabled = newStatus
        this.ManagerGroup = [...this.ManagerGroup]
      }).catch(e => console.log(e))
    },
    doPriv (scape) {
      this.$store.dispatch('setPageInfo', this.pageSize, this.pageNo)
      this.$router.push({
        name: 'privileges', params: {
          group_name: scape.row.name,
          group_id: scape.row.id
        }
      })
    },
    routeToUser (row) {
      this.$router.push({ name: 'user', params: { group_name: row.name } })
    },
    editUserClose () {
      this.editUserVisible = false
      this.transferChanged = false
    },
    pushNewUserList () {
      const selectUserIDs = [...this.usersListValue]
      updateGroupMember(this.groupInfo, selectUserIDs).then(r => {
        this.$notify({
          title: '成功',
          message: '更新组成员成功！',
          type: 'success'
        })
        this.doGetManagerGroup()
      }).catch(e => console.log(e))
    },
    filterUserMethod (query, item) {
      return item.spell.indexOf(query) > -1
    },
    transferChange (value, direction, movedKeys) {
      this.transferChanged = true
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
.group-board :deep() .el-transfer {
  margin-top: -15px;
  margin-bottom: 20px;
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
}
.group-board :deep() .el-transfer-panel {
  width: 250px;
}
.group-board :deep() .el-transfer__buttons {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}
.group-board :deep() .el-transfer__button {
  padding: 10px 15px;
}
.group-board :deep() .el-transfer__button {
  margin-left: 0px;
}
.user-list-push-box {
  display: float;
  text-align: right;
}
.user-list-push-btn {
  margin-right: 26px;
}
.user-list-close-btn {
  margin-right: 8px;
}
</style>
