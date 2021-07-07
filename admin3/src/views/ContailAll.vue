<template>
  <div class="main-board">
    <div class="do_action">
      <div style="padding-right: 15px">
        <el-button size="small" type="warning" @click="restartServer()"
          >重新加载配置（Reload）</el-button
        >
        <el-button size="small" type="primary" @click="publishJobsfunc()"
          >发布</el-button
        >
        <el-button size="small" type="success" plain @click="doAdd()"
          >添加</el-button
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
    <el-table
      size="mini"
      highlight-current-row
      border
      :data="jobs"
      stripe
      :row-style="rowStyle"
      :cell-style="cellStyle"
    >
      <el-table-column label="序号" width="50px">
        <template v-slot="scope">
          {{ scope.$index + 1 }}
        </template>
      </el-table-column>
      <el-table-column label="分组名称" prop="name"> </el-table-column>
      <el-table-column label="端口号" width="90px" prop="port">
      </el-table-column>
      <el-table-column label="IP数" width="80px" prop="count">
      </el-table-column>
      <el-table-column label="重写标签" prop="relabel_name"> </el-table-column>
      <el-table-column label="最后更新时间" prop="update_at">
        <template v-slot="{ row }">
          <span>{{ parseTimeSelf(row.update_at) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" width="200px">
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
          ref="addJobInfo"
          :model="addJobInfo"
          label-width="90px"
          size="small"
        >
          <el-form-item label="分组名：" prop="name">
            <el-input style="width: 230px" v-model="addJobInfo.name"></el-input>
          </el-form-item>
          <el-form-item label="端口号" prop="port">
            <el-input
              type="number"
              style="width: 230px"
              v-model.number="addJobInfo.port"
            ></el-input>
          </el-form-item>
          <el-form-item label="重写标签" prop="relabel_id">
            <el-select
              v-model="addJobInfo.relabel_id"
              filterable
              allow-create
              default-first-option
              placeholder="请选择"
              style="width: 230px"
            >
              <el-option
                v-for="item in relabelList"
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
              @click="onSubmit('addJobInfo')"
              >{{ buttonTitle }}</el-button
            >
            <el-button size="small" type="info" @click="onCancel('addJobInfo')"
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
  getDefJobsWithSplitPage,
  postDefJob,
  putDefJob,
  deleteDefJob,
  enabledDefaultJob
} from '@/api/defJobs.js'
import { publishJobs } from '@/api/jobs.js'
import { getAllReLabels } from '@/api/relabel.js'
import { restartSrv } from '@/api/srv'

export default {
  name: 'JobsDefault',
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
      jobs: [],
      addJobInfo: {
        'name': '',
        'port': 0,
        'display_order': 1,
        'relabel_id': 0
      },
      relabelList: [],
      pageSize: 15,
      pageTotal: 0,
      currentPage: 1,
      searchContent: '',
      deleteVisible: {},
      dialogVisible: false,
      buttonTitle: '',
      dialogTitle: '',
      rules: {
        name: [
          { required: true, message: '请输入正确的分组名称', validator: validateStr, trigger: ['blur'] }
        ],
        port: [
          { required: true, type: 'number', min: 0, max: 65535, message: '请输入端口号[>=0, <=65535]', trigger: ['change'] }
        ],
        relabel_id: [
          { required: true, type: 'number', min: 1, message: '请选择重写标签', trigger: ['change'] }
        ]
        // display_order: [
        //   { type: 'number', min: 1, message: '请输入排序号[>=1]', trigger: ['blur'] }
        // ]
      }
    }
  },
  created () {
    // this.doGetJobs()
  },
  mounted () {
    this.doGetJobs()
  },
  methods: {
    doAdd () {
      getAllReLabels().then(r => {
        this.relabelList = r.data
        this.addJobInfo = {
          'name': '',
          'port': 0,
          'display_order': 1
        }
        this.buttonTitle = '创建'
        this.dialogTitle = '增加IP分组'
        this.dialogVisible = true
      }).catch(e => {
        console.log(e)
      })
    },
    doGetJobs (getInfo) {
      if (!getInfo) {
        getInfo = {
          'pageNo': this.currentPage,
          'pageSize': this.pageSize,
          'search': this.searchContent
        }
      }
      getDefJobsWithSplitPage(getInfo).then(
        r => {
          let n = 0
          r.data.data.forEach(element => {
            this.deleteVisible[n] = false
            n += 1
          })
          this.jobs = r.data.data
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
      getAllReLabels().then(r => {
        this.relabelList = r.data
        this.buttonTitle = '更新'
        this.dialogTitle = '编辑IP分组'
        this.addJobInfo.id = data.row.id
        this.addJobInfo.name = data.row.name
        this.addJobInfo.port = data.row.port
        this.addJobInfo.relabel_id = data.row.relabel_id
        this.addJobInfo.relabel_name = data.row.relabel_name
        //   this.addJobInfo.display_order = data.row.display_order
        this.dialogVisible = true
      }).catch(e => {
        console.log(e)
      })
    },
    handleSizeChange (val) {
      let getInfo = {
        'pageNo': this.currentPage,
        'pageSize': val,
        'search': this.searchContent
      }
      this.doGetJobs(getInfo)
    },
    handleCurrentChange (val) {
      let getInfo = {
        'pageNo': val,
        'pageSize': this.pageSize,
        'search': this.searchContent
      }
      this.doGetJobs(getInfo)
    },
    handleClose (done) {
      this.dialogVisible = false
    },
    onSubmit (formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          let postData = {}
          postData['name'] = this.addJobInfo.name
          postData['port'] = this.addJobInfo.port
          postData['display_order'] = this.addJobInfo.display_order
          postData['relabel_id'] = this.addJobInfo.relabel_id
          if (this.buttonTitle === '创建') {
            postDefJob(postData).then(
              r => {
                this.$notify({
                  title: '成功',
                  message: '创建成功！',
                  type: 'success'
                });
                this.doGetJobs()
                this.dialogVisible = false
                this.$refs[formName].resetFields()
              }
            ).catch(
              e => { console.log(e) }
            )
          } else {
            postData['id'] = this.addJobInfo.id
            putDefJob(postData).then(
              r => {
                this.$notify({
                  title: '成功',
                  message: '更新成功！',
                  type: 'success'
                });
                this.doGetJobs()
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
      this.doGetJobs()
    },
    parseTimeSelf (t) {
      var time = new Date(Date.parse(t))
      return time.toLocaleDateString() + ' ' + time.toTimeString().split(' ')[0]
    },
    doYes (scope) {
      deleteDefJob(scope.row.id).then(
        r => {
          this.$notify({
            title: '成功',
            message: '删除成功！',
            type: 'success'
          });
          this.doGetJobs()
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
    doUp (data) {
      if (this.jobs.length === 0) {
        this.$notify({
          title: '错误',
          message: '没有有效的数据！',
          type: 'error'
        })
        return false
      }
      if (data.row.display_order === 1) {
        this.$notify({
          title: '警告',
          message: '已经是顶层！',
          type: 'warning'
        })
        return false
      }
      this.swapAction({
        id: data.row.id,
        action: 'up',
        display_order: data.row.display_order
      })
    },
    doDown (data) {
      if (this.jobs.length === 0) {
        this.$notify({
          title: '错误',
          message: '没有有效的数据！',
          type: 'error'
        })
        return false
      }
      let lastObj = this.jobs[this.jobs.length - 1]
      if (lastObj.id === data.row.id && this.jobs.length !== this.pageSize) {
        this.$notify({
          title: '警告',
          message: '已经是底层！',
          type: 'warning'
        })
        return false
      }
      if (this.currentPage * this.pageSize === this.pageTotal) {
        this.$notify({
          title: '警告',
          message: '已经是底层！',
          type: 'warning'
        })
        return false
      }
      this.swapAction({
        id: data.row.id,
        action: 'down',
        display_order: data.row.display_order
      })
    },
    swapAction (swapInfo) {
      swapJob(swapInfo).then(
        r => {
          if (swapInfo.action === 'up') {
            this.$notify({
              title: '成功',
              message: '提升成功！',
              type: 'success'
            });
          } else {
            this.$notify({
              title: '成功',
              message: '下降成功！',
              type: 'success'
            });
          }
          this.doGetJobs()
        }
      ).catch(
        e => {
          console.log(e)
        }
      )
    },
    publishJobsfunc () {
      publishJobs().then(
        r => {
          this.$notify({
            title: '成功',
            message: '发布成功！',
            type: 'success'
          });
        }
      ).catch(
        e => { console.log(e) }
      )
    },
    restartServer () {
      restartSrv().then(
        r => {
          this.$notify({
            title: '成功',
            message: '重新加载成功！',
            type: 'success'
          });
        }
      ).catch(
        e => { console.log(e) }
      )
    },
    invocate (scope) {
      const newStatus = !this.jobs[scope.$index].enabled
      const jInfo = {
        id: scope.row.id,
        enabled: newStatus
      }
      enabledDefaultJob(jInfo).then(r => {
        this.$notify({
          title: '成功',
          message: '更新状态成功！',
          type: 'success'
        });
        this.jobs[scope.$index].enabled = newStatus
        this.jobs = [...this.jobs]
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
  max-width: 900px;
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
</style>
