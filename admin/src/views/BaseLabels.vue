<template>
  <div class="main-board">
    <div class="btn-action-area">
      <div>
        <span class="explain-words">
          说明：定义的标签（key/value）可以在<el-tag size="mini">JOB组管理</el-tag>和<el-tag size="mini">告警管理</el-tag>中使用</span>
      </div>
      <div class="do_action">
        <div style="padding-right: 15px">
          <el-button size="small" type="success" icon="el-icon-baseball" plain @click="doAdd()">添加标签</el-button>
        </div>
        <div>
          <div>
            <el-input size="small" placeholder="请输入内容" v-model="searchContent" @keyup.enter="onSearch()">
              <template #append>
                <el-button size="small" @click="onSearch()" icon="el-icon-search"></el-button>
              </template>
            </el-input>
          </div>
        </div>
      </div>
    </div>
    <el-table size="mini" highlight-current-row border :data="baseLabels" stripe :row-style="rowStyle"
      :cell-style="cellStyle">
      <el-table-column label="序号" width="50px">
        <template v-slot="scope">
          {{ scope.$index + 1 }}
        </template>
      </el-table-column>
      <el-table-column label="标签名称" prop="label" show-overflow-tooltip> </el-table-column>
      <el-table-column label="更新账号" prop="update_by" show-overflow-tooltip> </el-table-column>
      <el-table-column label="更新时间" prop="update_at">
        <template v-slot="{ row }">
          <span>{{ parseTimeSelf(row.update_at) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" width="210px">
        <template v-slot="scope" align="center">
          <div class="actioneara">
            <div>
              <el-button size="mini" type="primary" @click="doEdit(scope)" plain>编辑</el-button>
            </div>
            <div>
              <el-button size="mini" type="info" v-if="scope.row.enabled === true" @click="invocate(scope)" plain>禁用
              </el-button>
              <el-button size="mini" type="warning" v-if="scope.row.enabled === false" @click="invocate(scope)" plain>启用
              </el-button>
            </div>
            <div>
              <el-popover :visible="deleteVisible[scope.$index]" placement="top">
                <p>确定删除吗？</p>
                <div style="text-align: right; margin: 0">
                  <el-button size="mini" type="text" @click="doNo(scope)">取消</el-button>
                  <el-button type="primary" size="mini" @click="doYes(scope)">确定</el-button>
                </div>
                <template #reference>
                  <el-button size="mini" type="danger" plain @click="doDelete(scope)">删除</el-button>
                </template>
              </el-popover>
            </div>
          </div>
        </template>
      </el-table-column>
    </el-table>
    <div class="block">
      <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange" :current-page="currentPage"
        :page-sizes="[10, 15, 30, 50]" :page-size="pageSize" layout="total, sizes, prev, pager, next, jumper"
        :total="pageTotal">
      </el-pagination>
    </div>
    <el-dialog :title="dialogTitle" v-model="dialogVisible" width="400px" modal :before-close="handleClose">
      <span>
        <el-form label-position="right" :rules="rules" ref="addLabelInfo" :model="addLabelInfo" label-width="90px"
          size="small">
          <el-form-item label="标签名：" prop="label">
            <el-input style="width: 230px" v-model="addLabelInfo.label"></el-input>
          </el-form-item>
          <el-form-item size="small">
            <el-button size="small" type="primary" @click="onSubmit('addLabelInfo')">{{ buttonTitle }}</el-button>
            <el-button size="small" type="info" @click="onCancel('addLabelInfo')">取消</el-button>
          </el-form-item>
        </el-form>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import {
  getBaseLabels,
  putBaseLabels,
  postBaseLabels,
  deleteBaseLabels,
  enabledBaseLabels
} from '@/api/base.js'

export default {
  name: 'baseLabels',
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
      baseLabels: [],
      addLabelInfo: {
        'id': 0,
        'label': ''
      },
      pageSize: 15,
      pageTotal: 0,
      currentPage: 1,
      searchContent: '',
      deleteVisible: {},
      dialogVisible: false,
      buttonTitle: '',
      dialogTitle: '',
      rules: {
        label: [
          { required: true, message: '请输入正确的标签名称', validator: validateStr, trigger: ['blur'] }
        ]
        // display_order: [
        //   { type: 'number', min: 1, message: '请输入排序号[>=1]', trigger: ['blur'] }
        // ]
      }
    }
  },
  created () {
    // this.doGetBaseLabels()
  },
  mounted () {
    this.doGetBaseLabels()
  },
  methods: {
    doAdd () {
      this.addLabelInfo = {
        'label': '',
        'port': 0,
        'display_order': 1
      }
      this.buttonTitle = '创建'
      this.dialogTitle = '增加标签'
      this.dialogVisible = true
    },
    doGetBaseLabels (getInfo) {
      if (!getInfo) {
        getInfo = {
          'pageNo': this.currentPage,
          'pageSize': this.pageSize,
          'search': this.searchContent
        }
      }
      getBaseLabels(getInfo).then(
        r => {
          this.baseLabels = r.data.data
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
      this.addLabelInfo.id = data.row.id
      this.addLabelInfo.label = data.row.label
      this.dialogVisible = true
    },
    handleSizeChange (val) {
      let getInfo = {
        'pageNo': this.currentPage,
        'pageSize': val,
        'search': this.searchContent
      }
      this.doGetBaseLabels(getInfo)
    },
    handleCurrentChange (val) {
      let getInfo = {
        'pageNo': val,
        'pageSize': this.pageSize,
        'search': this.searchContent
      }
      this.doGetBaseLabels(getInfo)
    },
    handleClose (done) {
      this.dialogVisible = false
    },
    onSubmit (formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          let postData = {}
          postData['label'] = this.addLabelInfo.label
          if (this.buttonTitle === '创建') {
            postBaseLabels(postData).then(
              r => {
                this.$notify({
                  title: '成功',
                  message: '创建成功！',
                  type: 'success',
                  duration: 1000,
                });
                this.doGetBaseLabels()
                this.dialogVisible = false
                this.$refs[formName].resetFields()
              }
            ).catch(
              e => { console.log(e) }
            )
          } else {
            postData['id'] = this.addLabelInfo.id
            putBaseLabels(postData).then(
              r => {
                this.$notify({
                  title: '成功',
                  message: '更新成功！',
                  type: 'success',
                  duration: 1000,
                });
                this.doGetBaseLabels()
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
      this.doGetBaseLabels()
    },
    parseTimeSelf (t) {
      var time = new Date(Date.parse(t))
      return time.toLocaleDateString() + ' ' + time.toTimeString().split(' ')[0]
    },
    doYes (scope) {
      deleteBaseLabels({ id: scope.row.id }).then(
        r => {
          this.$notify({
            title: '成功',
            message: '删除成功！',
            type: 'success',
            duration: 1000,
          });
          this.doGetBaseLabels()
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
      // if (row.rowIndex % 2 === 0) {
      //   rs['background'] = '#f2eada'
      // }
      return rs
    },
    cellStyle (column) {
      let cs = {
        'padding': '0'
      }
      return cs
    },
    invocate (scope) {
      const newStatus = !this.baseLabels[scope.$index].enabled
      const bInfo = {
        id: scope.row.id,
        enabled: newStatus
      }
      enabledBaseLabels(bInfo).then(r => {
        this.$notify({
          title: '成功',
          message: '更新状态成功！',
          type: 'success',
          duration: 1000,
        });
        this.baseLabels[scope.$index].enabled = newStatus
        this.baseLabels = [...this.baseLabels]
      }).catch(e => console.log(e))
    }
  }
}
</script>

<style scoped>
.do_action {
  text-align: right;
  margin-top: -5px;
  padding-bottom: 8px;
}

el-dialog {
  padding: 0px;
}

.do_action>div {
  display: inline-block;
}

.main-board {
  padding: 0;
  /* max-width: 900px; */
  margin: 0 auto;
}

.searchSelect {
  width: 90px;
}

.el-input__inner {
  width: 130px;
}

.el-pagination {
  text-align: center;
}

.block {
  padding-top: 12px;
}

el-tabs {
  padding: 0px;
}

.el-form:last-child {
  text-align: left;
}

.borderNone :deep() .el-input__inner {
  border: none;
  background: transparent;
}

.actioneara {
  display: flex;
  flex-wrap: nowrap;
  justify-content: space-around;
}

.change_order_button {
  display: flex;
  flex-wrap: nowrap;
  justify-content: center;
}

.change_order_button .el-button {
  padding: auto 0px;
  width: 20px;
}

.explain-words {
  font: 0.9em Arial, Tahoma, Verdana;
  color: #777;
}

.btn-action-area {
  display: flex;
  flex-wrap: nowrap;
  justify-content: space-between;
}
</style>
