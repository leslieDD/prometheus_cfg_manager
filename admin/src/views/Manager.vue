<template>
  <div class="ipManager-board">
    <div class="action-btn-area">
      <div class="action-btn-del">
        <el-button
          icon="el-icon-lightning"
          size="small"
          type="info"
          plain
          @click="doBatchDel()"
          >删除选中项</el-button
        >
      </div>
      <div class="do_action">
        <div style="padding-right: 15px">
          <!-- <el-button
            v-if="pushing === false"
            size="small"
            icon="el-icon-upload"
            type="primary"
            @click="onPublish()"
            >发布【file_sd_configs】</el-button
          > -->
          <!-- <el-button
            v-if="pushing === true"
            size="small"
            icon="el-icon-loading"
            type="primary"
            >发布【file_sd_configs】</el-button
          > -->
          <el-button size="small" type="warning" plain @click="doBatchAdd()"
            >批量添加</el-button
          >
          <el-button size="small" type="success" plain @click="doAdd()"
            >添加IP</el-button
          >
        </div>
        <div>
          <el-input
            size="small"
            @keyup.enter="onSearch()"
            placeholder="请输入内容"
            v-model="searchContent"
            class="input-with-select"
          >
            <template #prepend>
              <el-select
                class="searchSelect"
                v-model="selectOption"
                placeholder="请选择"
              >
                <el-option label="IP地址" value="1"></el-option>
                <el-option label="分组" value="2"></el-option>
              </el-select>
            </template>
            <template #append>
              <el-button icon="el-icon-search" @click="onSearch()"></el-button>
            </template>
          </el-input>
        </div>
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
      header-align="center"
      @expand-change="expandChange"
      @selection-change="handleSelectionChange"
      @cell-mouse-enter="cellMouseEnter"
    >
      <el-table-column type="selection" width="40"> </el-table-column>
      <el-table-column type="expand">
        <template #default="props">
          <el-descriptions title="分组列表" size="mini" :column="4" border>
            <el-descriptions-item
              v-for="(ipaddr, index) in groupNameList[props.row.id]"
              :key="index"
              :label="index + 1"
              >{{ ipaddr }}</el-descriptions-item
            >
          </el-descriptions>
        </template>
      </el-table-column>
      <el-table-column
        label="序号"
        width="50px"
        align="center"
        header-align="center"
      >
        <template v-slot="scope">
          {{ scope.$index + 1 }}
        </template>
      </el-table-column>
      <el-table-column
        label="IP地址"
        prop="ipaddr"
        align="center"
        header-align="center"
      >
      </el-table-column>
      <el-table-column
        label="分组选项"
        prop="jobs_id"
        align="center"
        header-align="center"
        width="230px"
      >
        <template v-slot="scope">
          <el-select
            v-model="selectTypeValue[scope.row.id]"
            class="borderNone"
            popper-class="pppselect"
            @change="tableSelectChange($event, scope.$index, scope.row)"
            @visible-change="handleSelect($event, scope.$index, scope.row)"
            size="small"
            multiple
            collapse-tags
            placeholder="请选择"
          >
            <el-option
              v-for="item in jobs"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            >
            </el-option>
          </el-select>
        </template>
      </el-table-column>
      <el-table-column
        label="状态"
        width="150px"
        align="center"
        header-align="center"
      >
        <template #header>
          <el-tooltip content="所有机器状态5分钟更新一次" placement="top">
            <span type="warning"
              >状态
              <i
                style="font-size: 13px; color: #0081ff"
                class="el-icon-warning"
              ></i
            ></span>
          </el-tooltip>
        </template>
        <template v-slot="{ row }">
          <el-tooltip placement="left" effect="light">
            <template #content>
              <table class="pure-table">
                <tbody>
                  <tr
                    class="pure-table-tr"
                    v-for="(value, key, index) in ipStatusList"
                    :key="index"
                  >
                    <td class="pure-table-td">{{ key }}</td>
                    <td class="pure-table-td">{{ value }}</td>
                  </tr>
                </tbody>
              </table>
            </template>

            <el-button
              v-if="needWarning[row.ipaddr] === true"
              class="el-button-color"
              icon="el-icon-warning"
              type="text"
              size="mini"
              @click="displayIPStatusDialog(row)"
              >查看状态详情</el-button
            >
            <el-button
              v-else
              type="text"
              size="mini"
              icon="el-icon-sunny"
              @click="displayIPStatusDialog(row)"
              >查看状态详情</el-button
            >
          </el-tooltip>
        </template>
      </el-table-column>
      <el-table-column
        label="最后更新账号"
        prop="update_by"
        align="center"
        header-align="center"
      ></el-table-column>
      <el-table-column
        label="最后更新时间"
        prop="update_at"
        align="center"
        header-align="center"
      >
        <template v-slot="{ row }">
          <span>{{ parseTimeSelf(row.update_at) }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="操作"
        align="center"
        header-align="center"
        width="210px"
      >
        <template v-slot="scope">
          <el-button type="primary" plain size="mini" @click="edit(scope)"
            >编辑</el-button
          >
          <el-button
            v-if="scope.row.enabled === true"
            type="info"
            plain
            size="mini"
            @click="invocate(scope)"
            >禁用</el-button
          >
          <el-button
            v-if="scope.row.enabled === false"
            type="warning"
            plain
            size="mini"
            @click="invocate(scope)"
            >启用</el-button
          >
          <el-popover
            :visible="deleteVisible[scope.$index]"
            placement="top"
            :width="160"
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
        </template>
      </el-table-column>
    </el-table>
    <div class="block" v-if="pageshow">
      <el-pagination
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :current-page="currentPage"
        :page-sizes="[15, 50, 100, 200]"
        :page-size="pageSize"
        layout="total, sizes, prev, pager, next, jumper"
        :total="pageTotal"
        ref="pagination"
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
          ref="addMechineInfo"
          :model="addMechineInfo"
          label-width="90px"
          size="small"
        >
          <el-form-item label="IP地址：" prop="ipAddr">
            <el-input
              style="width: 230px"
              v-model="addMechineInfo.ipAddr"
              @keyup.enter="onSubmitAndContinue('addMechineInfo')"
            ></el-input>
          </el-form-item>
          <el-form-item label="分组名：" prop="jobId">
            <el-select
              style="width: 230px"
              v-model="addMechineInfo.jobId"
              placeholder="请选择"
              collapse-tags
              multiple
            >
              <el-option
                v-for="(item, index) in jobs"
                :key="index"
                :label="item.name"
                :value="item.id"
              ></el-option>
            </el-select>
          </el-form-item>
          <el-form-item size="small">
            <el-button
              size="small"
              type="primary"
              :disabled="addAndContinueDisabled"
              @click="onSubmitAndContinue('addMechineInfo')"
              >创建并继续</el-button
            >
            <el-button
              size="small"
              type="primary"
              @click="onSubmit('addMechineInfo')"
              >{{ enterBtnTitle }}</el-button
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
    <div class="dialog-ip-status">
      <el-dialog
        :title="'IP状态信息：' + currentIPStatus"
        v-model="dialogIpStatusVisible"
        width="900px"
      >
        <el-table :data="ipStatusData" size="mini">
          <el-table-column
            width="200px"
            property="job"
            label="分组"
          ></el-table-column>
          <el-table-column width="100px" label="状态">
            <template v-slot="{ row }">
              <el-tag v-if="row.health === 'up'" size="mini" type="success">{{
                row.health
              }}</el-tag>
              <el-tag
                v-else-if="row.health === 'down'"
                size="mini"
                type="danger"
                >{{ row.health }}</el-tag
              >
              <el-tag v-else type="info" size="mini">{{ row.health }}</el-tag>
            </template>
          </el-table-column>
          <el-table-column
            property="last_error"
            label="错误信息"
          ></el-table-column>
        </el-table>
      </el-dialog>
    </div>
  </div>
</template>

<script>
import { getJobs } from '@/api/jobs'
import {
  getMachines,
  postMachine,
  deleteMachine,
  putMachine,
  enabledMachine,
  batchDeleteMachine
} from '@/api/machines'
import { publish } from '@/api/publish'

export default {
  name: 'Manager',
  components: {
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
      deleteVisible: {},
      enterBtnTitle: '创建',
      selectTableChanged: {},
      addMechineInfo: {
        ipAddr: '',
        jobId: []
      },
      groupNameList: {},
      addAndContinueDisabled: false,
      dialogTitle: '新增IP地址',
      rules: {
        ipAddr: [
          { required: true, message: '请输入正确的IP地址', validator: validateIP, trigger: 'blur' }
        ],
        // jobId: [
        //   { required: true, message: '请选择分组', trigger: ['blur', 'change'] }
        // ]
      },
      multipleSelection: [],
      pageshow: false,
      pushing: false,
      editObj: null,
      ipStatusData: [],
      dialogIpStatusVisible: false,
      currentIPStatus: '',
      ipStatusList: {},
      needWarning: {},
    }
  },
  created () {
    // this.doGetMechines()
  },
  mounted () {
    if (this.$route.params.currentPage) {
      this.currentPage = parseInt(this.$route.params.currentPage)
    }
    if (this.$route.params.pageSize) {
      this.pageSize = parseInt(this.$route.params.pageSize)
    }
    this.doGetMechines()
  },
  methods: {
    doAdd () {
      this.enterBtnTitle = '创建'
      this.dialogTitle = '新增IP地址'
      this.dialogVisible = true
      this.addAndContinueDisabled = false
      this.addMechineInfo = {}
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
              this.machines = r.data.data
              this.pageTotal = r.data.totalCount
              this.currentPage = r.data.pageNo
              this.pageSize = r.data.pageSize
              this.pageshow = false;//让分页隐藏
              this.$nextTick(() => {//重新渲染分页
                this.pageshow = true;
              });
              this.deleteVisible = {}
              let needWarning = {}
              let n = 0
              r.data.data.forEach(element => {
                this.selectTypeValue[element.id] = element.jobs_id
                this.deleteVisible[n] = false
                needWarning[element.ipaddr] = false
                const other = element.msrv_status.findIndex((item, index) => {
                  return item.health !== 'up';
                })
                if (other !== -1) {
                  needWarning[element.ipaddr] = true
                }
                n += 1
              })
              this.needWarning = needWarning
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
    displayIPStatusDialog (row) {
      this.ipStatusData = row.msrv_status
      this.currentIPStatus = row.ipaddr
      this.dialogIpStatusVisible = true
    },
    cellMouseEnter (row, column, cell, event) {
      if (column.label !== '状态' && column.label !== 'status') {
        return
      }
      let ipStatusList = {}
      row.msrv_status.forEach(x => {
        if (ipStatusList[x.health] !== undefined) {
          ipStatusList[x.health] += 1
        } else {
          ipStatusList[x.health] = 1
        }
      })
      this.ipStatusList = ipStatusList
    },
    handleSelect (visible, index, row) {
      if (visible) {
        this.selectTableChanged[row.id] = false // 下拉框打开的时候，设置为false，刚打开，所以没有变化
        return
      }
      if (this.selectTableChanged[row.id] === false) {
        return
      }
      var updateData = {}
      updateData['id'] = row.id
      updateData['ipaddr'] = row.ipaddr
      updateData['jobs_id'] = this.selectTypeValue[row.id]
      putMachine(updateData).then(
        r => {
          this.$notify({
            title: '成功',
            message: '更新成功！',
            type: 'success'
          });
          this.selectTableChanged[row.id] = false
          this.selectTypeValue = { ...this.selectTypeValue }
          this.expandChange(row, row)
        }
      ).catch(
        e => {
          console.log(e)
        }
      )
    },
    tableSelectChange (visible, index, row) {
      this.selectTableChanged[row.id] = true
    },
    handleSizeChange (val) {
      // this.pageSize = val
      let getInfo = {
        'pageNo': this.currentPage,
        'pageSize': val,
        'search': this.searchContent,
        'option': this.selectOption
      }
      this.doGetMechines(getInfo)
    },
    handleCurrentChange (val) {
      // this.currentPage = val
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
          postData['jobs_id'] = this.addMechineInfo.jobId
          postMachine(postData).then(
            r => {
              this.$notify({
                title: '成功',
                message: '创建成功！',
                type: 'success'
              });
              this.doGetMechines()
              //   this.$refs[formName].resetFields()
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
          postData['jobs_id'] = this.addMechineInfo.jobId
          if (this.enterBtnTitle === '创建') {
            postMachine(postData).then(
              r => {
                this.$notify({
                  title: '成功',
                  message: '创建成功！',
                  type: 'success'
                });
                this.doGetMechines()
                this.dialogVisible = false
                this.$refs[formName].resetFields()
              }
            ).catch(e => console.log(e))
          } else {
            postData['id'] = this.addMechineInfo.id
            putMachine(postData).then(
              r => {
                this.$notify({
                  title: '成功',
                  message: '更新成功！',
                  type: 'success'
                });
                // this.doGetMechines()
                this.selectTypeValue[this.addMechineInfo.id] = this.addMechineInfo.jobId
                this.selectTypeValue = { ...this.selectTypeValue }
                this.dialogVisible = false
                this.editObj.row.jobs_id = this.addMechineInfo.jobId
                this.editObj.row.ipaddr = this.addMechineInfo.ipAddr
                this.$refs[formName].resetFields()
                this.expandChange({ id: this.addMechineInfo.id }, { id: this.addMechineInfo.id })
              }
            ).catch(e => console.log(e))
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
      this.doGetMechines()
    },
    onPublish () {
      this.pushing = true
      publish().then(
        r => {
          this.$notify({
            title: '成功',
            message: '发布成功！',
            type: 'success'
          });
          this.pushing = false
        }
      ).catch(
        e => {
          console.log(e)
          this.pushing = false
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
          this.$notify({
            title: '成功',
            message: '删除成功！',
            type: 'success'
          });
          this.doGetMechines()
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
    edit (scope) {
      this.editObj = scope
      this.enterBtnTitle = '更新'
      this.dialogTitle = '编辑IP地址'
      this.addAndContinueDisabled = true
      this.dialogVisible = true
      this.addMechineInfo = {
        ipAddr: scope.row.ipaddr,
        id: scope.row.id
      }
      if (this.selectTypeValue[scope.row.id].length !== 0) {
        this.addMechineInfo.jobId = this.selectTypeValue[scope.row.id]
      } else {
        this.addMechineInfo.jobId = scope.row.jobs_id
      }
    },
    expandChange (row, expandRows) {
      if (!expandRows || expandRows.length === 0) {
        return
      }
      let groupName = []
      this.selectTypeValue[row.id].forEach(i => {
        groupName.push(this.jobsMap[i])
      })
      this.groupNameList[row.id] = groupName
    },
    invocate (scope) {
      const newStatus = !this.machines[scope.$index].enabled
      const mInfo = {
        id: scope.row.id,
        enabled: newStatus
      }
      enabledMachine(mInfo).then(r => {
        this.$notify({
          title: '成功',
          message: '更新状态成功！',
          type: 'success'
        });
        this.machines[scope.$index].enabled = newStatus
        this.machines = [...this.machines]
      }).catch(e => console.log(e))
    },
    doBatchAdd () {
      const params = {
        currentPage: this.currentPage,
        pageSize: this.pageSize
      }
      this.$router.push({ name: 'BatchOpt', params: params })
    },
    doBatchDel () {
      if (this.multipleSelection.length === 0) {
        this.$notify({
          title: '警告',
          message: '未选中任何项！',
          type: 'warning'
        });
        return false
      }
      this.$confirm('是否确定删除？', '确认信息', {
        distinguishCancelAndClose: true,
        confirmButtonText: '确定',
        cancelButtonText: '放弃'
      }).then(_ => {
        batchDeleteMachine(this.multipleSelection).then(r => {
          this.$notify({
            title: '成功',
            message: '删除所选项成功！',
            type: 'success'
          });
          this.doGetMechines()
        }).catch(e => console.log(e))
      }).catch(e => console.log(e))
    },
    handleSelectionChange (val) {
      this.multipleSelection = []
      val.forEach(each => {
        this.multipleSelection.push(each.id)
      })
    }
  }
}
</script>

<style scoped>
.action-btn-area {
  display: flex;
  justify-content: space-between;
  flex-wrap: nowrap;
}
.action-btn-del {
  margin-top: -5px;
}
.do_action {
  text-align: right;
  margin-top: -5px;
  padding-bottom: 8px;
}
.have-board {
  border: 1px solid burlywood;
}
el-dialog {
  padding: 0px;
}
.do_action > div {
  display: inline-block;
}
/* .main-board {
  padding: 0;
  max-width: 1100px;
  margin: 0 auto;
} */
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
.el-form:last-child {
  text-align: left;
}
.borderNone :deep() .el-input__inner {
  border: none;
  background: transparent;
  width: 225px;
}
.dialog-ip-status :deep() .el-dialog__body {
  padding-top: 5px;
}
.pure-table {
  font-family: verdana, arial, sans-serif;
  font-size: 11px;
  color: #333333;
  border-width: 1px;
  border-color: #999999;
  border-collapse: collapse;
}
.pure-table-tr {
  background: #b5cfd2;
  border-width: 1px;
  padding: 5px;
  border-style: solid;
  border-color: #999999;
}
.pure-table-td {
  background: #dcddc0;
  border-width: 1px;
  padding: 5px;
  border-style: solid;
  border-color: #999999;
}
.el-button-color {
  color: red;
}
</style>
