<template>
  <dir class="main-board">
    <el-tabs
      type="border-card"
      v-model="activeTabName"
      @tab-click="handleTabClick"
    >
      <el-tab-pane label="IP管理" name="ipManager">
        <div class="do_action">
          <div style="padding-right: 15px">
            <el-button size="small" type="primary" @click="onPublish()"
              >发布</el-button
            >
            <el-button size="small" type="success" plain @click="doAdd()"
              >添加</el-button
            >
          </div>
          <div>
            <el-input
              size="small"
              @keyup.enter.native="onSearch()"
              placeholder="请输入内容"
              v-model="searchContent"
              class="input-with-select"
            >
              <el-select
                class="searchSelect"
                v-model="selectOption"
                slot="prepend"
                placeholder="请选择"
              >
                <el-option label="IP地址" value="1"></el-option>
                <el-option label="分组" value="2"></el-option>
              </el-select>
              <el-button
                slot="append"
                icon="el-icon-search"
                @click="onSearch()"
              ></el-button>
            </el-input>
          </div>
        </div>
        <el-table
          size="mini"
          highlight-current-row
          border
          :data="machines"
          stripe
          :row-style="rowStyle"
          :cell-style="cellStyle"
        >
          <el-table-column label="序号" width="50px">
            <template slot-scope="scope">
              {{ scope.$index + 1 }}
            </template>
          </el-table-column>
          <el-table-column label="IP地址" prop="ipaddr"> </el-table-column>
          <el-table-column label="分组选项" prop="job_id">
            <template slot-scope="scope">
              <el-select
                v-model="selectTypeValue[scope.row.id]"
                class="borderNone"
                popper-class="pppselect"
                @change="handleSelect(scope.$index, scope.row)"
                size="small"
                placeholder="请选择"
              >
                <el-option
                  v-for="item in jobs"
                  :key="item.id"
                  :label="item.display_order + ' ' + item.name"
                  :value="item.id"
                >
                </el-option>
              </el-select>
            </template>
          </el-table-column>
          <el-table-column label="最后更新时间" prop="update_at">
            <template slot-scope="{ row }">
              <span>{{ parseTimeSelf(row.update_at) }}</span>
            </template>
          </el-table-column>
          <el-table-column label="操作" align="center">
            <template slot-scope="scope">
              <el-popover
                trigger="click"
                :ref="`popover-${scope.$index}`"
                placement="top"
                width="160"
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
                <el-button size="mini" slot="reference" type="danger" plain
                  >删除</el-button
                >
              </el-popover>
            </template>
          </el-table-column>
        </el-table>
        <div class="block">
          <el-pagination
            @size-change="handleSizeChange"
            @current-change="handleCurrentChange"
            :current-page="currentPage"
            :page-sizes="[15, 50, 100, 200]"
            :page-size="pageSize"
            layout="total, sizes, prev, pager, next, jumper"
            :total="pageTotal"
          >
          </el-pagination>
        </div>
      </el-tab-pane>
      <el-tab-pane label="分组管理" name="groupManager">
        <componentJob ref="groupManager"></componentJob>
      </el-tab-pane>
      <el-tab-pane label="配置预览" name="preview">
        <previewer ref="preview"></previewer>
      </el-tab-pane>
      <el-tab-pane label="分组预览" name="gpreview">
        <gboard ref="gpreview"></gboard>
      </el-tab-pane>
    </el-tabs>
    <el-dialog
      title="新增IP地址"
      :visible.sync="dialogVisible"
      width="400px"
      modal
      :before-close="handleClose"
    >
      <span>
        <el-form
          label-position="right"
          :rules="rules"
          ref="addMechineInfo"
          :model="addMechineInfo"
          label-width="90px"
          size="small"
        >
          <el-form-item label="IP地址：" prop="ipAddr">
            <el-input
              style="width: 230px"
              v-model="addMechineInfo.ipAddr"
            ></el-input>
          </el-form-item>
          <el-form-item label="分组名：" prop="jobId">
            <el-select
              style="width: 230px"
              v-model="addMechineInfo.jobId"
              placeholder="请选择"
            >
              <el-option
                v-for="(item, index) in jobs"
                :key="index"
                :label="item.display_order + ' ' + item.name"
                :value="item.id"
              ></el-option>
            </el-select>
          </el-form-item>
          <el-form-item size="small">
            <el-button
              size="small"
              type="primary"
              @click="onSubmitAndContinue('addMechineInfo')"
              >创建并继续</el-button
            >
            <el-button
              size="small"
              type="primary"
              @click="onSubmit('addMechineInfo')"
              >创建</el-button
            >
            <el-button
              size="small"
              type="info"
              @click="onCancel('addMechineInfo')"
              >取消</el-button
            >
          </el-form-item>
        </el-form>
      </span>
    </el-dialog>
  </dir>
</template>

<script>
import componentJob from '@/components/jobs'
import previewer from '@/components/preview'
import gboard from '@/components/gboard'
import { getJobs } from '@/api/jobs'
import { getMachines, postMachine, deleteMachine, putMachine } from '@/api/machines'
import { publish } from '@/api/publish'

export default {
  name: 'Manager',
  components: {
    componentJob: componentJob,
    previewer: previewer,
    gboard: gboard
  },
  data () {
    function validateIP (rule, value, callback) {
      if (value === '' || typeof value === 'undefined' || value == null) {
        callback(new Error('请输入正确的IP地址'))
      } else {
        const reg = /^(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])$/
        if ((!reg.test(value)) && value !== '') {
          callback(new Error('请输入正确的IP地址'))
        } else {
          callback()
        }
      }
    }
    return {
      machines: [],
      jobs: [],
      jobsMap: {},
      selectTypeValue: {},
      search: '',
      pageSize: 15,
      pageTotal: 0,
      currentPage: 1,
      selectOption: '1',
      searchContent: '',
      dialogVisible: false,
      activeTabName: 'ipManager',
      addMechineInfo: {
        ipAddr: '',
        jobId: ''
      },
      rules: {
        ipAddr: [
          { required: true, message: '请输入正确的IP地址', validator: validateIP, trigger: 'blur' }
        ],
        jobId: [
          { required: true, message: '请选择分组', trigger: ['blur', 'change'] }
        ]
      }
    }
  },
  created () {
    this.doGetMechines()
  },
  methods: {
    doAdd () {
      this.dialogVisible = true
    },
    doGetJobs () {
      getJobs().then(
        r => {
          this.jobs = r.data
          r.data.forEach(e => {
            this.jobsMap[e.id] = e.name
          })
        }
      ).catch(
        e => {
          console.log(e)
        }
      )
    },
    doGetMechines (getInfo) {
      getJobs().then(
        r => {
          this.jobs = r.data
          r.data.forEach(e => {
            this.jobsMap[e.id] = e.name
          })
          if (!getInfo) {
            getInfo = {
              'pageNo': this.currentPage,
              'pageSize': this.pageSize,
              'search': this.searchContent,
              'option': this.selectOption
            }
          }
          getMachines(getInfo).then(
            r => {
              this.pageTotal = r.data.totalCount
              this.currentPage = r.data.pageNo
              this.pageSize = r.data.pageSize
              this.machines = r.data.data
              r.data.data.forEach(element => {
                this.selectTypeValue[element.id] = this.jobsMap[element.job_id[0]]
              })
            }
          ).catch(
            e => {
              console.log(e)
            }
          )
        }
      ).catch(
        e => {
          console.log(e)
        }
      )
    },
    handleSelect (index, row) {
      var updateData = {}
      updateData['id'] = row.id
      updateData['ipaddr'] = row.ipaddr
      updateData['job_id'] = [this.selectTypeValue[row.id]]
      putMachine(updateData).then(
        r => {
          this.$message({
            showClose: true,
            message: '更新成功！',
            type: 'success'
          })
          this.selectTypeValue = { ...this.selectTypeValue }
        }
      ).catch(
        e => {
          console.log(e)
        }
      )
    },
    handleSizeChange (val) {
      let getInfo = {
        'pageNo': this.currentPage,
        'pageSize': val,
        'search': this.searchContent,
        'option': this.selectOption
      }
      this.doGetMechines(getInfo)
    },
    handleCurrentChange (val) {
      let getInfo = {
        'pageNo': val,
        'pageSize': this.pageSize,
        'search': this.searchContent,
        'option': this.selectOption
      }
      this.doGetMechines(getInfo)
    },
    handleClose (done) {
      this.dialogVisible = false
    },
    onSubmitAndContinue (formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          let postData = {}
          postData['ipaddr'] = this.addMechineInfo.ipAddr
          postData['job_id'] = [parseInt(this.addMechineInfo.jobId)]
          postMachine(postData).then(
            r => {
              this.$message({
                showClose: true,
                message: '创建成功！',
                type: 'success'
              })
              this.doGetMechines()
              this.$refs[formName].resetFields()
            }
          ).catch(
            e => {
              console.log(e)
            }
          )
        } else {
          return false
        }
      })
    },
    onSubmit (formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          let postData = {}
          postData['ipaddr'] = this.addMechineInfo.ipAddr
          postData['job_id'] = [parseInt(this.addMechineInfo.jobId)]
          postMachine(postData).then(
            r => {
              this.$message({
                showClose: true,
                message: '创建成功！',
                type: 'success'
              })
              this.doGetMechines()
              this.dialogVisible = false
              this.$refs[formName].resetFields()
            }
          ).catch(
            e => {
              console.log(e)
            }
          )
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
      this.doGetMechines()
    },
    onPublish () {
      publish().then(
        r => {
          this.$message({
            showClose: true,
            message: '发布成功！',
            type: 'success'
          })
        }
      ).catch(
        e => {
          console.log(e)
        }
      )
    },
    parseTimeSelf (t) {
      var time = new Date(Date.parse(t))
      // return time.toLocaleString()
      return time.toLocaleDateString() + ' ' + time.toTimeString().split(' ')[0]
    },
    doYes (scope) {
      deleteMachine(scope.row.id).then(
        r => {
          this.$message({
            showClose: true,
            message: '删除成功！',
            type: 'success'
          })
          this.doGetMechines()
        }
      ).catch(
        e => {
          console.log(e)
        }
      )
      scope._self.$refs[`popover-${scope.$index}`].doClose()
    },
    doNo (scope) {
      scope._self.$refs[`popover-${scope.$index}`].doClose()
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
    handleTabClick (tab, event) {
      if (tab.name === 'ipManager') {
        this.doGetMechines()
      } else if (tab.name === 'groupManager') {
        this.$refs.groupManager.doGetJobs()
      } else if (tab.name === 'preview') {
        this.$refs.preview.loadYaml()
      }
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
.borderNone >>> .el-input__inner {
  border: none;
  background: transparent;
}
</style>
