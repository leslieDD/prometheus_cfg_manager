<template>
  <div class="jobs-labels-board">
    <div class="label-and-action">
      <div>
        <el-tag type="warning" effect="dark" @change="clickElTag"
          >编辑分组：{{ jobInfo.name }} [ {{ jobInfo.count }} ]</el-tag
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
      @expand-change="expandChange"
    >
      <el-table-column type="expand">
        <template #default="props">
          <el-descriptions title="IP列表" size="mini" :column="3" border>
            <el-descriptions-item
              v-for="(ipaddr, index) in ipsAndLabels[props.row.id].ips"
              :key="index"
              :label="index + 1"
              >{{ ipaddr }}</el-descriptions-item
            >
          </el-descriptions>
          <el-descriptions title="标签列表" size="mini" :column="3" border>
            <el-descriptions-item
              v-for="(item, index) in ipsAndLabels[props.row.id].labels"
              :key="index"
              :label="item.key"
              >{{ item.value }}
            </el-descriptions-item>
          </el-descriptions>
        </template>
      </el-table-column>
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
                >编辑IP</el-button
              >
            </div>
            <div>
              <el-button
                size="mini"
                type="primary"
                @click="doEditLabelList(scope)"
                plain
                >编辑标签</el-button
              >
            </div>
            <div>
              <el-button
                v-if="scope.row.enabled === true"
                size="mini"
                type="info"
                @click="invocate(scope)"
                plain
                >禁用</el-button
              >
              <el-button
                v-if="scope.row.enabled === false"
                size="mini"
                type="warning"
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
              :type="buttonType"
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
    <el-dialog
      :title="'编辑IP列表：' + editObject"
      v-model="editIPVisible"
      width="700px"
      modal
      :before-close="editIPClose"
    >
      <el-row type="flex" align="middle" justify="center">
        <el-transfer
          v-model="editIPValue"
          filterable
          :filter-method="filterIPMethod"
          filter-placeholder="请输入关键字"
          :data="editIPData"
          :titles="['分组IP地址池', '子分组IP列表']"
        >
        </el-transfer>
      </el-row>
      <div class="ip-list-push-box">
        <el-button
          class="ip-list-push-btn"
          type="warning"
          size="small"
          @click="pushNewIPList"
          >提交</el-button
        >
      </div>
    </el-dialog>
    <el-dialog
      :title="'编辑标签列表：' + editObject"
      v-model="editLabelsVisible"
      width="800px"
      modal
      :before-close="editLabelsClose"
    >
      <div class="add-label-form">
        <el-form
          size="mini"
          :inline="true"
          :rules="glRules"
          ref="addLablesForm"
          :model="addNewGroupLabels"
          class="demo-form-inline"
        >
          <el-form-item label="标签：" prop="key">
            <div class="add-labels-select-box">
              <el-select
                v-model="addNewGroupLabels.key"
                filterable
                allow-create
                default-first-option
                placeholder="请选择或者输入"
              >
                <el-option
                  v-for="item in defaultLabels"
                  :key="item.id"
                  :label="item.label"
                  :value="item.label"
                >
                </el-option>
              </el-select>
            </div>
          </el-form-item>
          <el-form-item label="标签值：" prop="value">
            <div class="add-labels-input-box">
              <el-input v-model="addNewGroupLabels.value"></el-input>
            </div>
          </el-form-item>
          <el-form-item>
            <div class="add-labels-btn-box">
              <el-button
                type="primary"
                @click="onAddNewLable('addLablesForm')"
                >{{ labelsBtnTitle }}</el-button
              >
              <el-button
                type="danger"
                @click="onResetAddNewLable('addLablesForm')"
                >重置</el-button
              >
            </div>
          </el-form-item>
        </el-form>
      </div>
      <div class="status-edit-area">
        <div>
          <el-alert
            effect="dark"
            :title="alertTitle"
            type="warning"
            :closable="false"
            show-icon
          >
          </el-alert>
        </div>
        <div class="search-label-input">
          <el-input
            size="small"
            placeholder="请输入内容"
            @keyup.enter="onLablesSearch()"
            v-model="glSearchContent"
          >
            <template #append>
              <el-button
                size="small"
                @click="onLablesSearch()"
                icon="el-icon-search"
              ></el-button>
            </template>
          </el-input>
        </div>
      </div>
      <div>
        <el-table
          size="mini"
          highlight-current-row
          border
          stripe
          :row-style="rowStyle"
          :cell-style="cellStyle"
          header-align="center"
          :data="labelsData"
        >
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
          <el-table-column property="key" label="标签名"></el-table-column>
          <el-table-column property="value" label="标签值"></el-table-column>
          <el-table-column
            label="最后更新"
            prop="update_at"
            align="center"
            header-align="center"
          >
            <template v-slot="{ row }">
              <span>{{ parseTimeSelf(row.update_at) }}</span>
            </template>
          </el-table-column>
          <el-table-column label="操作" align="center">
            <template v-slot="scope" align="center">
              <div class="actioneara">
                <div>
                  <el-button
                    size="mini"
                    type="primary"
                    @click="doLabelsEdit(scope)"
                    plain
                    >编辑</el-button
                  >
                </div>
                <div>
                  <el-button
                    size="mini"
                    type="info"
                    v-if="scope.row.enabled === true"
                    @click="invocateLabels(scope)"
                    plain
                    >禁用</el-button
                  >
                  <el-button
                    size="mini"
                    type="warning"
                    v-if="scope.row.enabled === false"
                    @click="invocateLabels(scope)"
                    plain
                    >启用</el-button
                  >
                </div>
                <div>
                  <el-popover
                    :visible="deleteLabelsVisible[scope.$index]"
                    placement="top"
                  >
                    <p>确定删除吗？</p>
                    <div style="text-align: right; margin: 0">
                      <el-button
                        size="mini"
                        type="text"
                        @click="doLablesNo(scope)"
                        >取消</el-button
                      >
                      <el-button
                        type="primary"
                        size="mini"
                        @click="doLabelsYes(scope)"
                        >确定</el-button
                      >
                    </div>
                    <template #reference>
                      <el-button
                        size="mini"
                        type="danger"
                        plain
                        @click="doLabelsDelete(scope)"
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
            @size-change="glHandleSizeChange"
            @current-change="glHandleCurrentChange"
            :current-page="glCurrentPage"
            :page-sizes="[10]"
            :page-size="glPageSize"
            layout="total, sizes, prev, pager, next, jumper"
            :total="glPageTotal"
          >
          </el-pagination>
        </div>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
  addJobGroup,
  getJobGroup,
  putJobGroup,
  getJobMachines,
  getJobMachinesForGroup,
  putJobMachinesForGroup,
  delJobGroup,
  getGroupLabels,
  putGroupLabels,
  delGroupLabels,
  getAllIPAndLabels,
  enabledJobGroup,
  enabledJobGroupLables
} from '@/api/labelsJob.js'
import { getDefaultLabels } from '@/api/monitor.js'
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
      deleteLabelsVisible: {},
      dialogVisible: false,
      buttonTitle: '',
      buttonType: 'primary',
      dialogTitle: '',
      subGroupData: {},
      editIPVisible: false,
      editIPValue: [],  //穿梭框绑定的数据，选定到右侧框中的数据项的value组成的数组
      editIPValueMap: {},
      editIPData: [],  //Transfer 的数据源	array[{ key, label, disabled }]
      editIPDataMap: {},
      editObject: '',
      jobGroupIDCurrent: 0,
      editLabelsVisible: false,
      defaultLabels: [],
      labelsData: [],
      glPageTotal: 0,
      glPageSize: 10,
      glCurrentPage: 1,
      addNewGroupLabels: {},
      glSearchContent: '',
      alertTitle: '当前状态：可以添加新标签',
      labelsBtnTitle: '添加',
      ipsAndLabels: {},
      ipsAndLabelsDone: {},
      needtoUpdate: false,
      rules: {
        name: [
          { required: true, message: '请输入正确的分组名称', validator: validateStr, trigger: ['blur'] }
        ]
      },
      glRules: {
        key: [
          { required: true, message: '请选择或者输入分组名称', trigger: ['blur', 'change'] }
        ]
        // value: [
        //   { required: true, message: '请输入正确的标签值', trigger: ['blur'] }
        // ]
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
    pushNewIPList () {
      let newValues = []
      this.editIPValue.forEach(machinedID => {
        const v = this.editIPValueMap[machinedID]
        let newV = {
          machines_id: machinedID,
          ipaddr: this.editIPDataMap[machinedID].ipaddr
        }
        if (v) {
          newV.id = v.id
        } else {
          newV.id = 0
        }
        newValues.push(newV)
      })
      // const jobID = this.jobInfo.id
      putJobMachinesForGroup(this.jobGroupIDCurrent, newValues).then(r => {
        this.$notify({
          title: '成功',
          message: '更新子组成功！',
          type: 'success'
        })
        this.needtoUpdate = true
      }).catch(e => console.log(e))
    },
    pushNewLablesList () {

    },
    filterIPMethod (query, item) {
      // return item.spell.indexOf(query) > -1;
      return true
    },
    doEditIPList (scope) {
      const jobID = this.jobInfo.id
      const gid = scope.row.id
      getJobMachinesForGroup(gid).then(r => {
        let editIPValue = []
        this.editIPValueMap = {}
        r.data.forEach(item => {
          editIPValue.push(item.machines_id)
          this.editIPValueMap[item.machines_id] = item  // 不会出现重复
        })
        // this.editIPValue = editIPValue
        let editIPData = []
        this.editIPDataMap = {}
        getJobMachines(jobID).then(r => {
          r.data.forEach(item => {
            let defVal = {
              label: item.ipaddr,
              key: item.id,
              spell: item.ipaddr,
              disabled: false
            }
            // if (this.editIPValueMap[item.id]) {
            //   defVal.disabled = true
            // }
            editIPData.push(defVal)
            this.editIPDataMap[item.id] = item
          })
          this.editIPData = editIPData
          this.editIPValue = editIPValue
          this.editObject = scope.row.name
          this.jobGroupIDCurrent = gid
          this.editIPVisible = true
        }).catch(e => {
          console.log(e)
        })
        // this.editIPVisible = true
      }).catch(e => { console.log(e) })
    },
    editIPClose (scope) {
      if (this.needtoUpdate) {
        this.doGetSubGroup()
      }
      this.editIPVisible = false
    },
    editLabelsClose () {
      if (this.needtoUpdate) {
        this.doGetSubGroup()
      }
      this.editLabelsVisible = false
    },
    doEditLabelList (scope) {
      const jobID = this.jobInfo.id
      const gid = scope.row.id
      const getInfo = {
        'pageNo': this.glCurrentPage,
        'pageSize': this.glPageSize,
        'search': this.glSearchContent,
        id: gid
      }
      getDefaultLabels().then(r => {
        this.defaultLabels = r.data
        getGroupLabels(getInfo).then(r => {
          // 有两个一样的地方要更新
          this.deleteLabelsVisible = {}
          let n = 0
          r.data.data.forEach(item => {
            this.deleteLabelsVisible[n] = false
            n += 1
          })
          this.glPageTotal = r.data.totalCount
          this.glCurrentPage = r.data.pageNo
          this.glPageSize = r.data.pageSize
          this.labelsData = r.data.data
          this.jobGroupIDCurrent = gid
          this.editLabelsVisible = true
          this.editObject = scope.row.name
        }).catch(e => console.log(e))
      }).catch(e => console.log(e))
    },
    doGetGroupLabels (getInfo) {
      if (!getInfo) {
        getInfo = {
          'pageNo': this.glCurrentPage,
          'pageSize': this.glPageSize,
          'search': this.glSearchContent,
          'id': this.jobGroupIDCurrent
        }
      }
      getGroupLabels(getInfo).then(r => {
        // 有两个一样的地方要更新
        this.deleteLabelsVisible = {}
        let n = 0
        r.data.data.forEach(item => {
          this.deleteLabelsVisible[n] = false
          n += 1
        })
        this.glPageTotal = r.data.totalCount
        this.glCurrentPage = r.data.pageNo
        this.glPageSize = r.data.pageSize
        this.labelsData = r.data.data
        this.jobGroupIDCurrent = gid
        this.editLabelsVisible = true
      }).catch(e => console.log(e))
    },
    doAddSubGroupShow () {
      this.buttonTitle = '创建'
      this.buttonType = 'primary'
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
    onAddNewLable (formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          const gid = this.jobGroupIDCurrent
          let newGroupLabels = { ...this.addNewGroupLabels, job_group_id: this.jobGroupIDCurrent }
          putGroupLabels(gid, newGroupLabels).then(
            r => {
              this.$notify({
                title: '成功',
                message: '创建标签成功！',
                type: 'success'
              })
              this.doGetGroupLabels()
              this.needtoUpdate = true
            }
          ).catch(e => console.log(e))
        }
      })
    },
    onResetAddNewLable (formName) {
      this.$refs[formName].resetFields()
      this.alertTitle = '当前状态：可以添加新标签'
      this.addNewGroupLabels = {}
      this.labelsBtnTitle = '添加'
      this.buttonType = 'primary'
    },
    doEdit (scope) {
      this.buttonTitle = '更新'
      this.buttonType = 'warning'
      this.dialogTitle = '更新子分组'
      this.subGroupData = { ...scope.row }
      this.dialogVisible = true
    },
    doLabelsEdit (scope) {
      this.alertTitle = `当前编辑对象：${scope.row.key}，序号：${scope.$index + 1}`
      this.addNewGroupLabels = {
        key: scope.row.key,
        value: scope.row.value,
        id: scope.row.id
      }
      this.labelsBtnTitle = '更新'
    },
    handleSizeChange (val) {
      let getInfo = {
        'pageNo': this.currentPage,
        'pageSize': val,
        'search': this.searchContent
      }
      this.doGetSubGroup(getInfo)
    },
    glHandleSizeChange (val) {
      let getInfo = {
        'pageNo': this.glCurrentPage,
        'pageSize': val,
        'search': this.glSearchContent,
        'id': this.jobGroupIDCurrent
      }
      this.doGetGroupLabels(getInfo)
    },
    handleCurrentChange (val) {
      let getInfo = {
        'pageNo': val,
        'pageSize': this.pageSize,
        'search': this.searchContent
      }
      this.doGetSubGroup(getInfo)
    },
    glHandleCurrentChange (val) {
      let getInfo = {
        'pageNo': val,
        'pageSize': this.glPageSize,
        'search': this.glSearchContent,
        'id': this.jobGroupIDCurrent
      }
      this.doGetGroupLabels(getInfo)
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
      if (isNaN(this.jobInfo.id)) {
        return
      }
      getInfo.id = this.jobInfo.id
      getJobGroup(getInfo).then(
        r => {
          let n = 0
          this.ipsAndLabels = {}
          this.ipsAndLabelsDone = {}
          r.data.data.forEach(element => {
            this.deleteVisible[n] = false
            this.ipsAndLabels[element.id] = {}
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
                message: '更新子组成功！',
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
    doLabelsYes (scope) {
      const delInfo = {
        gid: this.jobGroupIDCurrent,
        id: scope.row.id
      }
      delGroupLabels(delInfo).then(r => {
        this.$notify({
          title: '成功',
          message: '删除标签成功！',
          type: 'success'
        })
        this.doGetGroupLabels()
      }).catch(e => {
        console.log(e)
      })
      this.deleteLabelsVisible[scope.$index] = false
    },
    doNo (scope) {
      this.deleteVisible[scope.$index] = false
    },
    doLablesNo (scope) {
      this.deleteLabelsVisible[scope.$index] = false
    },
    doDelete (scope) {
      this.deleteVisible[scope.$index] = true
    },
    doLabelsDelete (scope) {
      this.deleteLabelsVisible[scope.$index] = true
    },
    onSearch () {
      this.doGetSubGroup()
    },
    onLablesSearch () {
      this.doGetGroupLabels()
    },
    clickElTag (checked) {
      console.log('checked')
      this.$router.push({ name: 'jobs' })
    },
    expandChange (row, expandRows) {
      if (!expandRows || expandRows.length === 0) {
        return
      }
      if (this.ipsAndLabelsDone[row.id]) {
        return
      }
      const query = {
        'job_id': row.jobs_id,
        'group_id': row.id
      }
      getAllIPAndLabels(query).then(r => {
        this.ipsAndLabels[row.id] = r.data
        this.ipsAndLabelsDone[row.id] = true
      }).catch(e => console.log(e))
    },
    invocate (scope) {
      const newStatus = !this.subGroups[scope.$index].enabled
      const gInfo = {
        id: scope.row.id,
        enabled: newStatus
      }
      enabledJobGroup(gInfo).then(r => {
        this.$notify({
          title: '成功',
          message: '更新状态成功！',
          type: 'success'
        });
        this.subGroups[scope.$index].enabled = newStatus
        this.subGroups = [...this.subGroups]
      }).catch(e => console.log(e))
    },
    invocateLabels (scope) {
      const newStatus = !this.labelsData[scope.$index].enabled
      const lInfo = {
        id: scope.row.id,
        enabled: newStatus
      }
      enabledJobGroupLables(lInfo).then(r => {
        this.$notify({
          title: '成功',
          message: '更新状态成功！',
          type: 'success'
        });
        this.labelsData[scope.$index].enabled = newStatus
        this.labelsData = [...this.labelsData]
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
.label-and-action :deep() .el-input__inner {
  width: 230px;
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
.add-label-form {
  margin-top: -18px;
}
.add-label-form :deep() .el-form-item:nth-child(3) {
  margin-right: 0px;
}
.search-label-input :deep() .el-input__inner {
  display: float;
  float: right;
  width: 250px;
  /* text-align: right; */
}
.status-edit-area {
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
  justify-content: space-between;
}
/* .el-form:last-child {
  text-align: left;
} */
/* .borderNone :deep() .el-input__inner {
  border: none;
  background: transparent;
} */
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
.jobs-labels-board :deep() .el-transfer {
  margin-top: -15px;
  margin-bottom: 20px;
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
}
.jobs-labels-board :deep() .el-transfer-panel {
  width: 250px;
}
.jobs-labels-board :deep() .el-transfer__buttons {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}
.jobs-labels-board :deep() .el-transfer__button {
  padding: 10px 15px;
}
.jobs-labels-board :deep() .el-transfer__button {
  margin-left: 0px;
}
.ip-list-push-box {
  display: float;
  text-align: right;
}
.ip-list-push-btn {
  margin-right: 26px;
}
.add-labels-input-box :deep() .el-input__inner {
  width: 256px;
}
.add-labels-select-box {
  margin-right: 10px;
}
.add-labels-select-box :deep() .el-input__inner {
  width: 200px;
}
</style>
