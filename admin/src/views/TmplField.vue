<template>
  <div class="tmpl-field-board">
    <div class="btn-action-area">
      <div>
        <span class="explain-words"
          >说明：这儿定义的字段（key/value）可以在<el-tag size="mini"
            >模块编辑</el-tag
          >中使用，语法请参考golang中的<el-link
            href="https://pkg.go.dev/text/template"
            target="_blank"
            type="warning"
            >text/template</el-link
          ></span
        >
      </div>
      <div class="do_action">
        <div style="padding-right: 15px">
          <el-button size="small" type="success" plain @click="doAdd()"
            >添加模板字段</el-button
          >
        </div>
        <div>
          <div>
            <el-input
              size="small"
              placeholder="请输入内容"
              v-model="searchContent"
              @keyup.enter="onSearch()"
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
      </div>
    </div>
    <el-table
      size="mini"
      highlight-current-row
      border
      :data="tmplFields"
      stripe
      :row-style="rowStyle"
      :cell-style="cellStyle"
    >
      <el-table-column label="序号" width="50px">
        <template v-slot="scope">
          {{ scope.$index + 1 }}
        </template>
      </el-table-column>
      <el-table-column label="模板字段名称" prop="key"> </el-table-column>
      <el-table-column label="模板字段值" prop="value"> </el-table-column>
      <el-table-column label="最后更新账号" prop="update_by"> </el-table-column>
      <el-table-column label="最后更新时间" prop="update_at">
        <template v-slot="{ row }">
          <span>{{ parseTimeSelf(row.update_at) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" width="210px">
        <template v-slot="scope" align="center">
          <div class="actioneara">
            <div>
              <el-button size="mini" type="primary" @click="doEdit(scope)" plain
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
                >禁用</el-button
              >
              <el-button
                size="mini"
                type="warning"
                v-if="scope.row.enabled === false"
                @click="invocate(scope)"
                plain
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
                  <el-button type="primary" size="mini" @click="doYes(scope)"
                    >确定</el-button
                  >
                </div>
                <template #reference>
                  <el-button
                    size="mini"
                    type="danger"
                    plain
                    @click="doDelete(scope)"
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
        :page-sizes="[10, 15, 30, 50]"
        :page-size="pageSize"
        layout="total, sizes, prev, pager, next, jumper"
        :total="pageTotal"
      >
      </el-pagination>
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
          ref="addTmplFields"
          :model="addTmplFields"
          label-width="auto"
          size="small"
        >
          <el-form-item label="字段名称：" prop="key">
            <el-input
              style="width: 250px"
              v-model="addTmplFields.key"
            ></el-input>
          </el-form-item>
          <el-form-item label="字段值：" prop="value">
            <el-input
              style="width: 250px"
              v-model="addTmplFields.value"
            ></el-input>
          </el-form-item>
          <el-form-item size="small">
            <el-button
              size="small"
              type="primary"
              @click="onSubmit('addTmplFields')"
              >{{ buttonTitle }}</el-button
            >
            <el-button
              size="small"
              type="info"
              @click="onCancel('addTmplFields')"
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
  getBaseFields,
  putBaseFields,
  postBaseFields,
  deleteBaseFields,
  enabledBaseFields
} from '@/api/fields.js'

export default {
  name: 'tmplFields',
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
      tmplFields: [],
      addTmplFields: {
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
        key: [
          { required: true, message: '请输入正确的字段名称', validator: validateStr, trigger: ['blur'] }
        ],
        value: [
          { required: true, message: '请输入正确的字段值', validator: validateStr, trigger: ['blur'] }
        ]
        // display_order: [
        //   { type: 'number', min: 1, message: '请输入排序号[>=1]', trigger: ['blur'] }
        // ]
      }
    }
  },
  created () {
    // this.doGetBaseFields()
  },
  mounted () {
    this.doGetBaseFields()
  },
  methods: {
    doAdd () {
      this.addTmplFields = {
        'key': '',
        'value': ''
      }
      this.buttonTitle = '创建'
      this.dialogTitle = '增加字段'
      this.dialogVisible = true
    },
    doGetBaseFields (getInfo) {
      if (!getInfo) {
        getInfo = {
          'pageNo': this.currentPage,
          'pageSize': this.pageSize,
          'search': this.searchContent
        }
      }
      getBaseFields(getInfo).then(
        r => {
          this.tmplFields = r.data.data
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
      this.dialogTitle = '编辑字段'
      this.addTmplFields.id = data.row.id
      this.addTmplFields.key = data.row.key
      this.addTmplFields.value = data.row.value
      this.dialogVisible = true
    },
    handleSizeChange (val) {
      let getInfo = {
        'pageNo': this.currentPage,
        'pageSize': val,
        'search': this.searchContent
      }
      this.doGetBaseFields(getInfo)
    },
    handleCurrentChange (val) {
      let getInfo = {
        'pageNo': val,
        'pageSize': this.pageSize,
        'search': this.searchContent
      }
      this.doGetBaseFields(getInfo)
    },
    handleClose (done) {
      this.dialogVisible = false
    },
    onSubmit (formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          let postData = {}
          postData['key'] = this.addTmplFields.key
          postData['value'] = this.addTmplFields.value
          if (this.buttonTitle === '创建') {
            postData['enabled'] = true
            postBaseFields(postData).then(
              r => {
                this.$notify({
                  title: '成功',
                  message: '创建成功！',
                  type: 'success'
                });
                this.doGetBaseFields()
                this.dialogVisible = false
                this.$refs[formName].resetFields()
              }
            ).catch(
              e => { console.log(e) }
            )
          } else {
            postData['id'] = this.addTmplFields.id
            putBaseFields(postData).then(
              r => {
                this.$notify({
                  title: '成功',
                  message: '更新成功！',
                  type: 'success'
                });
                this.doGetBaseFields()
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
      this.doGetBaseFields()
    },
    parseTimeSelf (t) {
      var time = new Date(Date.parse(t))
      return time.toLocaleDateString() + ' ' + time.toTimeString().split(' ')[0]
    },
    doYes (scope) {
      deleteBaseFields({ id: scope.row.id }).then(
        r => {
          this.$notify({
            title: '成功',
            message: '删除成功！',
            type: 'success'
          });
          this.doGetBaseFields()
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
      const newStatus = !this.tmplFields[scope.$index].enabled
      const bInfo = {
        id: scope.row.id,
        enabled: newStatus
      }
      enabledBaseFields(bInfo).then(r => {
        this.$notify({
          title: '成功',
          message: '更新状态成功！',
          type: 'success'
        });
        this.tmplFields[scope.$index].enabled = newStatus
        this.tmplFields = [...this.tmplFields]
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
.do_action > div {
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
