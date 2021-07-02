<template>
  <div class="jobs-labels-board">
    <div class="label-and-action">
      <div>
        <el-tag type="success" effect="dark"
          >{{ jobInfo.name }} [ {{ jobInfo.count }} ]</el-tag
        >
      </div>
      <div class="do_action">
        <div style="padding-right: 15px">
          <el-button size="small" type="primary" @click="publishJobsfunc()"
            >发布此分组</el-button
          >
          <el-button
            size="small"
            type="success"
            plain
            @click="doAddSubGroupShow()"
            >添加子组</el-button
          >
        </div>
        <div>
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
      </div>
    </div>

    <el-table
      size="mini"
      highlight-current-row
      border
      :data="subGroups"
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
      <el-table-column label="IP数" prop="ip_count"> </el-table-column>
      <el-table-column label="标签数" prop="labels_count"> </el-table-column>
      <el-table-column label="最后更新时间" prop="update_at">
        <template v-slot="{ row }">
          <span>{{ parseTimeSelf(row.update_at) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="350px" align="center">
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
                type="primary"
                @click="doEditIPList(scope)"
                plain
                >编辑IP列表</el-button
              >
            </div>
            <div>
              <el-button
                size="mini"
                type="primary"
                @click="doEditLabelList(scope)"
                plain
                >编辑标签列表</el-button
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
          ref="subGroupData"
          :model="subGroupData"
          label-width="90px"
          size="small"
        >
          <el-form-item label="子组名：" prop="name">
            <el-input
              style="width: 230px"
              v-model="subGroupData.name"
            ></el-input>
          </el-form-item>
          <el-form-item size="small">
            <el-button
              size="small"
              type="primary"
              @click="onSubmit('subGroupData')"
              >{{ buttonTitle }}</el-button
            >
            <el-button
              size="small"
              type="info"
              @click="onCancel('subGroupData')"
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
  addJobGroup,
  getJobGroup,
  putJobGroup,
  getJobMachines,
  delJobGroup,
  getGroupLabels,
  putGroupLabels,
  getGroupMachine,
  putGroupMachine
} from '@/api/labelsJob.js'
import { restartSrv } from '@/api/srv'

export default {
  name: 'JobsLabels',
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
      subGroups: [],
      jobInfo: {},
      relabelList: [],
      pageSize: 15,
      pageTotal: 0,
      currentPage: 1,
      searchContent: '',
      deleteVisible: {},
      dialogVisible: false,
      buttonTitle: '',
      dialogTitle: '',
      subGroupData: {},
      rules: {
        name: [
          { required: true, message: '请输入正确的分组名称', validator: validateStr, trigger: ['blur'] }
        ]
      }
    }
  },
  created () {
    // this.doGetJobs()
  },
  mounted () {
    this.jobInfo = this.$route.params
    this.jobInfo.id = parseInt(this.jobInfo.id, 10)
    this.doGetSubGroup()
  },
  methods: {
    doAddSubGroupShow () {
      this.buttonTitle = '创建'
      this.dialogTitle = '创建新的子分组'
      this.dialogVisible = true
    },
    doAddSubGroup () {
      const jobSubGroupInfo = {}
      addJobGroup(jobSubGroupInfo).then(r => {
        this.subGroup = r.data
        this.$notify({
          title: '成功',
          message: '创建子组成功！',
          type: 'success'
        })
      }).catch(e => {
        console.log(e)
      })
    },
    doEdit (scope) {
      this.buttonTitle = '更新'
      this.dialogTitle = '更新子分组'
      this.subGroupData = { ...scope.row }
      this.dialogVisible = true
    },
    handleSizeChange (val) {
      let getInfo = {
        'pageNo': this.currentPage,
        'pageSize': val,
        'search': this.searchContent
      }
      this.doGetSubGroup(getInfo)
    },
    handleCurrentChange (val) {
      let getInfo = {
        'pageNo': val,
        'pageSize': this.pageSize,
        'search': this.searchContent
      }
      this.doGetSubGroup(getInfo)
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
    parseTimeSelf (t) {
      var time = new Date(Date.parse(t))
      return time.toLocaleDateString() + ' ' + time.toTimeString().split(' ')[0]
    },
    doGetSubGroup (getInfo) {
      if (!getInfo) {
        getInfo = {
          'pageNo': this.currentPage,
          'pageSize': this.pageSize,
          'search': this.searchContent
        }
      }
      getInfo.id = this.jobInfo.id
      getJobGroup(getInfo).then(
        r => {
          let n = 0
          r.data.data.forEach(element => {
            this.deleteVisible[n] = false
            n += 1
          })
          this.subGroups = r.data.data
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
    onSubmit (formName) {
      const createDate = {
        ...this.subGroupData,
        'jobs_id': this.jobInfo.id
      }
      this.$refs[formName].validate((valid) => {
        if (valid) {
          if (this.buttonTitle === '创建') {
            addJobGroup(createDate).then(r => {
              this.$notify({
                title: '成功',
                message: '创建子组成功！',
                type: 'success'
              })
              this.doGetSubGroup()
              this.dialogVisible = false
              this.$refs[formName].resetFields()
            }).catch(e => {
              console.log(e)
            })
          } else {
            putJobGroup(createDate).then(r => {
              this.$notify({
                title: '成功',
                message: '创建子组成功！',
                type: 'success'
              })
              this.doGetSubGroup()
              this.dialogVisible = false
              this.$refs[formName].resetFields()
            }).catch(e => {
              console.log(e)
            })
          }
        }
      })
    },
    handleClose (done) {
      this.dialogVisible = false
    },
    onCancel (formName) {
      this.subGroupData = {}
      this.$refs[formName].resetFields()
      this.dialogVisible = false
    },
    doYes (scope) {
      const delInfo = {
        jobs_id: this.jobInfo.id,
        id: scope.row.id
      }
      delJobGroup(delInfo).then(r => {
        this.$notify({
          title: '成功',
          message: '删除子组成功！',
          type: 'success'
        })
        this.doGetSubGroup()
      }).catch(e => {
        console.log(e)
      })
      this.deleteVisible[scope.$index] = false
    },
    doNo (scope) {
      this.deleteVisible[scope.$index] = false
    },
    doDelete (scope) {
      this.deleteVisible[scope.$index] = true
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
.label-and-action {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  flex-wrap: nowrap;
}
el-dialog {
  padding: 0px;
}
.do_action > div {
  display: inline-block;
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
</style>
