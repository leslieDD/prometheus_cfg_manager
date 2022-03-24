<template>
  <div class="relabel-config-box">
    <div class="btn-action-area">
      <div>
        <span class="explain-words">
        说明：定义prometheus中的<el-tag size="mini" type="warning"
          >relabel_configs</el-tag
        >规则，可以在<el-tag size="mini">JOB组管理</el-tag>中使用</span>
      </div>
      <!-- <div>
        <el-button
          size="mini"
          type="success"
          @click="doPostReLablesCode()"
          icon="el-icon-upload2"
          >提 交</el-button
        >
      </div> -->
    </div>
    <div class="main-content">
      <div class="page_left">
        <div class="do_action">
          <div style="padding-right: 15px">
            <el-button size="small" type="success" icon="el-icon-baseball" plain @click="doAdd()"
              >添加规则</el-button
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
        <el-scrollbar class="card-scrollbar">
          <div class="table-show">
            <el-table
              size="mini"
              highlight-current-row
              border
              :data="ReLabels"
              stripe
              :row-style="rowStyle"
              :cell-style="cellStyle"
            >
              <el-table-column label="序号" width="50px">
                <template v-slot="scope">
                  {{ scope.$index + 1 }}
                </template>
              </el-table-column>
              <el-table-column label="名称" prop="name"> </el-table-column>
              <el-table-column label="更新账号" prop="update_by">
              </el-table-column>
              <el-table-column label="更新时间" prop="update_at">
                <template v-slot="{ row }">
                  <span>{{ parseTimeSelf(row.update_at) }}</span>
                </template>
              </el-table-column>
              <el-table-column label="操作" align="center" width="320px">
                <template v-slot="scope" align="center">
                  <div class="actioneara">
                    <div>
                      <el-button
                        size="mini"
                        type="primary"
                        @click="doEdit(scope)"
                        plain
                        :disabled="editCodeButVisable[scope.row.id]"
                        icon="el-icon-edit"
                        >名称</el-button
                      >
                    </div>
                    <div>
                      <el-button
                        size="mini"
                        type="primary"
                        @click="doEditReLablesCode(scope)"
                        :disabled="editCodeButVisable[scope.row.id]"
                        plain
                        icon="el-icon-edit"
                        >规则</el-button
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
            <div class="block">
              <el-pagination
                @size-change="handleSizeChange"
                @current-change="handleCurrentChange"
                :current-page="currentPage"
                :page-sizes="[18, 25, 30]"
                :page-size="pageSize"
                layout="total, sizes, prev, pager, next, jumper"
                :total="pageTotal"
              >
              </el-pagination>
            </div>
          </div>
        </el-scrollbar>
      </div>
    </div>
    <el-dialog
      :title="dialogEditRuleTitle"
      v-model="dialogEditRuleVisible"
      width="900px"
      modal
      :before-close="handleEditRuleClose"
    >
      <span>
        <el-form
          label-position="right"
          ref="editRuleInfo"
          :model="editRuleInfo"
          label-width="20px"
          size="small"
        >
          <el-form-item label="" prop="rules">
            <el-input
              style="width: 820px"
              :autosize="{ minRows: 20, maxRows: 20 }" type="textarea" placeholder=""
              v-model="editRuleInfo.rules"
            ></el-input>
          </el-form-item>
          <el-form-item size="small">
            <div class="dialog_edit_rule_action">
              <el-button
                size="small"
                type="primary"
                @click="oneditRuleSubmit('editRuleInfo')"
                >更新规则</el-button
              >
              <el-button
                size="small"
                type="info"
                @click="oneditRuleCancel('editRuleInfo')"
                >取消</el-button
              >
            </div>
          </el-form-item>
        </el-form>
      </span>
    </el-dialog>
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
          ref="editRelabelInfo"
          :model="editRelabelInfo"
          label-width="90px"
          size="small"
        >
          <el-form-item label="标签名：" prop="name">
            <el-input
              style="width: 230px"
              v-model="editRelabelInfo.name"
            ></el-input>
          </el-form-item>
          <el-form-item size="small">
            <el-button
              size="small"
              type="primary"
              @click="onSubmit('editRelabelInfo')"
              >{{ buttonTitle }}</el-button
            >
            <el-button
              size="small"
              type="info"
              @click="onCancel('editRelabelInfo')"
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
  getReLabels,
  putReLabels,
  postReLabels,
  deleteReLabels,
  putReLabelsCode,
  enabledRelabelCfg,
  checkViewCodePriv,
} from '@/api/relabel.js'

export default {
  name: 'ymalRelabelEdit',
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
      ReLabels: [],
      editRelabelInfo: {
        'id': 0,
        'name': ''
      },
      editRuleInfo: {
        id: 0,
        rules: '',
      },
      editCodeButVisable: {},
      defaultRuleID: 0,
      pageSize: 18,
      pageTotal: 0,
      currentPage: 1,
      searchContent: '',
      deleteVisible: {},
      dialogVisible: false,
      dialogEditRuleVisible: false,
      buttonTitle: '',
      dialogTitle: '',
      dialogEditRuleTitle: '',
      rules: {
        name: [
          { required: true, message: '请输入正确的名称', validator: validateStr, trigger: ['blur'] }
        ]
      },
      currentID: 0,
    }
  },
  mounted () {
    this.doGetReLabels()
  },
  methods: {
    doAdd () {
      this.editRelabelInfo = {
        'id': 0,
        'name': ''
      }
      this.buttonTitle = '创建'
      this.dialogTitle = '增加重写标签'
      this.dialogVisible = true
    },
    doGetReLabels (getInfo) {
      if (!getInfo) {
        getInfo = {
          'pageNo': this.currentPage,
          'pageSize': this.pageSize,
          'search': this.searchContent
        }
      }
      getReLabels(getInfo).then(
        r => {
          r.data.data.forEach(element => {
            if (element.name === "空规则") {
              this.editCodeButVisable[element.id] = true
            } else {
              this.editCodeButVisable[element.id] = false
            }
          })
          this.ReLabels = r.data.data
          this.pageTotal = r.data.totalCount
          this.currentPage = r.data.pageNo
          this.pageSize = r.data.pageSize
        }
      ).catch(
        e => {
          console.log(e)
        }
      )
    },
    doEdit (data) {
      this.buttonTitle = '更新'
      this.dialogTitle = '编辑标签'
      this.editRelabelInfo.id = data.row.id
      this.editRelabelInfo.name = data.row.name
      this.dialogVisible = true
    },
    handleSizeChange (val) {
      let getInfo = {
        'pageNo': this.currentPage,
        'pageSize': val,
        'search': this.searchContent
      }
      this.doGetReLabels(getInfo)
    },
    handleCurrentChange (val) {
      let getInfo = {
        'pageNo': val,
        'pageSize': this.pageSize,
        'search': this.searchContent
      }
      this.doGetReLabels(getInfo)
    },
    handleClose (done) {
      this.dialogVisible = false
    },
    handleEditRuleClose(){
      this.dialogEditRuleVisible = false
    },
    oneditRuleSubmit(formName){
      this.doPostReLablesCode()
    },
    onSubmit (formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          let postData = {}
          postData['name'] = this.editRelabelInfo.name
          if (this.buttonTitle === '创建') {
            postReLabels(postData).then(
              r => {
                this.$notify({
                  title: '成功',
                  message: '创建成功！',
                  type: 'success'
                });
                this.doGetReLabels()
                this.dialogVisible = false
                this.$refs[formName].resetFields()
              }
            ).catch(
              e => { console.log(e) }
            )
          } else {
            postData['id'] = this.editRelabelInfo.id
            putReLabels(postData).then(
              r => {
                this.$notify({
                  title: '成功',
                  message: '更新成功！',
                  type: 'success'
                });
                this.doGetReLabels()
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
    oneditRuleCancel(formName){
      this.dialogEditRuleVisible = false
      this.$refs[formName].resetFields()
    },
    onSearch () {
      this.doGetReLabels()
    },
    parseTimeSelf (t) {
      var time = new Date(Date.parse(t))
      return time.toLocaleDateString() + ' ' + time.toTimeString().split(' ')[0]
    },
    doYes (scope) {
      deleteReLabels({ id: scope.row.id }).then(
        r => {
          this.$notify({
            title: '成功',
            message: '删除成功！',
            type: 'success'
          });
          this.doGetReLabels()
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
    refresh () {
    },
    doEditReLablesCode (scape) {
      checkViewCodePriv().then(r => {
        this.dialogEditRuleTitle = scape.row.name
        this.currentID = scape.row.id
        this.editRuleInfo = {
          id: scape.row.id,
          rules: scape.row.code,
        }
        this.dialogEditRuleVisible = true
      }).catch(e => console.log(e))
    },
    doPostReLablesCode() {
      const code = {
        id: this.editRuleInfo.id,
        code: this.editRuleInfo.rules,
      }
      putReLabelsCode(code).then(
        r => {
          this.$notify({
            title: '成功',
            message: '提交成功！',
            type: 'success'
          });
          this.doGetReLabels()
        }
      ).catch(
        e => {
          console.log(e)
        }
      )
    },
    invocate (scope) {
      const newStatus = !this.ReLabels[scope.$index].enabled
      const rInfo = {
        id: scope.row.id,
        enabled: newStatus
      }
      enabledRelabelCfg(rInfo).then(r => {
        this.$notify({
          title: '成功',
          message: '更新状态成功！',
          type: 'success'
        });
        this.ReLabels[scope.$index].enabled = newStatus
        this.ReLabels = [...this.ReLabels]
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
  height: 80vh;
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}
.page_left {
  margin-top: 10px;
  width: 100%;
  height: 100%;
}
.page_right {
  width: 44.5%;
  height: 100%;
}
.do_action {
  display: flex;
  flex-direction: row;
  justify-content: flex-end;
  margin-bottom: 5px;
}
.actioneara {
  display: flex;
  /* flex-wrap: nowrap; */
  flex-wrap: wrap;
  justify-content:space-around;
}
.el-pagination {
  text-align: center;
}
.block {
  padding-top: 12px;
}
.yaml-relabel-edit {
  border: 1px solid burlywood;
  margin-top: 10px;
}
.cm-s-monokai {
  height: 100%;
}
.explain-words {
  font: 0.9em Arial, Tahoma, Verdana;
  color: #777;
}
/* .table-show {
  width: 100%;
} */
.btn-action-area {
  display: flex;
  flex-wrap: nowrap;
  justify-content: space-between;
}

.dialog_edit_rule_action {
  display: flex;
  justify-content: right;
  margin-right: 20px;
}
.card-scrollbar {
  height: 90%;
}
</style>
