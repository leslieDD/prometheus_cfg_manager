<template>
  <div class="jobs-board">
    <div class="do_action">
      <div class="action-btn-del">
        <el-button
          icon="el-icon-lightning"
          size="small"
          type="info"
          plain
          @click="doBatchDel()"
          >删除选中</el-button
        >
        <el-button
          icon="el-icon-warning-outline"
          size="small"
          type="info"
          plain
          @click="doBatchDisable()"
          >禁用选中</el-button
        >
        <el-button
          icon="el-icon-warning-outline"
          size="small"
          type="success"
          plain
          @click="doBatchEnable()"
          >启用选中</el-button
        >
      </div>
      <div class="do-action-search">
        <span style="padding-right: 15px">
          <el-button size="small" type="success" icon="el-icon-baseball" plain @click="doAdd()"
            >添加组</el-button
          >
        </span>
        <span>
          <el-input
            size="small"
            placeholder="请输入内容"
            @keyup.enter="onSearch()"
            v-model="searchContent"
            style="width: 300px"
          >
            <template #append>
              <el-button
                size="small"
                @click="onSearch()"
                icon="el-icon-search"
              ></el-button>
            </template>
          </el-input>
        </span>
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
      @expand-change="expandChange"
      @selection-change="handleSelectionChange"
      @cell-click="clickcell"
    >
      <el-table-column type="selection" width="40"> </el-table-column>
      <el-table-column type="expand">
        <template #default="props">
          <el-descriptions title="IP列表" size="mini" :column="4" border>
            <el-descriptions-item
              v-for="(item, index) in jobAllIPList[props.row.id]"
              :key="index"
              :label="index + 1"
              >{{ item.ipaddr }}
                <span v-if="item.position.error"><el-tag type="danger" size="mini" >{{item.position.error}}</el-tag></span>
                <span v-else-if="item.position.isp"><el-tag size="mini" type="info">{{item.position.country}}{{item.position.province}}{{item.position.city}}{{item.position.county}}</el-tag><el-tag type="warning" size="mini" >{{item.position.isp}}</el-tag></span>
                <span v-else><el-tag size="mini" type="info">{{item.position.country}}{{item.position.province}}{{item.position.city}}{{item.position.county}}{{item.position.isp}}</el-tag></span>
            </el-descriptions-item>
          </el-descriptions>
        </template>
      </el-table-column>
      <el-table-column label="序号" width="50px">
        <template v-slot="scope">
          {{ scope.$index + 1 }}
        </template>
      </el-table-column>
      <el-table-column
        label="JOB组名称"
        prop="name"
        show-overflow-tooltip
      >
      </el-table-column>
      <el-table-column label="端口号" width="90px" prop="port">
        <template #header>
          <el-tooltip
            content="如果端口号为0，在生成配置时忽略端口号，只保留IP地址"
            placement="top"
          >
            <span type="warning"
              >端口号
              <i
                style="font-size: 13px; color: #0081ff"
                class="el-icon-warning"
              ></i
            ></span>
          </el-tooltip>
        </template>
      </el-table-column>
      <el-table-column label="IP数" width="55px">
        <template v-slot="scope">
          <el-button size="mini" type="text" @click="doEditLabelsJob(scope)">{{
            scope.row.ip_count
          }}</el-button>
        </template>
      </el-table-column>
      <el-table-column label="加黑" width="55px">
        <template v-slot="scope">
          <span>{{scope.row.black_count}}</span>
        </template>
      </el-table-column>
      <el-table-column label="子组数" width="65px">
        <template v-slot="scope">
          <el-button size="mini" type="text" @click="doEditLabelsJob(scope)">{{
            scope.row.group_count
          }}</el-button>
        </template>
      </el-table-column>
      <el-table-column
        label="重写标签"
        prop="relabel_name"
        show-overflow-tooltip
      >
        <template #header>
          <el-tooltip
            content="在‘基本配置’选项卡下的‘标签重写’下定义"
            placement="top"
          >
            <span type="warning"
              >重写标签
              <i
                style="font-size: 13px; color: #0081ff"
                class="el-icon-warning"
              ></i
            ></span>
          </el-tooltip>
        </template>
      </el-table-column>
      <!-- <el-table-column label="排序" width="75px" prop="display_order">
      </el-table-column>
      <el-table-column label="调整排序" width="70px">
        <template v-slot="scope">
          <div class="change_order_button">
            <el-button
              size="mini"
              type="primary"
              icon="el-icon-top"
              @click="doUp(scope)"
              plain
            ></el-button>
            <el-button
              size="mini"
              type="primary"
              icon="el-icon-bottom"
              @click="doDown(scope)"
              plain
            ></el-button>
          </div>
        </template>
      </el-table-column> -->
      <el-table-column
        label="最后更新账号"
        prop="update_by"
        width="100px"
      ></el-table-column>
      <el-table-column label="最后更新时间" prop="update_at" width="160px">
        <template v-slot="{ row }">
          <span>{{ parseTimeSelf(row.update_at) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" width="500px">
        <template v-slot="scope" align="center">
          <div class="actioneara">
            <div>
              <el-button size="mini" type="primary" @click="doEdit(scope)" plain
                >编辑</el-button
              >
            </div>
            <div>
              <el-button size="mini" type="primary" @click="doEditPool(scope)" plain
                >编辑IP列表</el-button
              >
            </div>
            <div>
              <el-button size="mini" type="primary" @click="doPool(scope)" plain
                >IP池</el-button
              >
            </div>
            <div>
              <el-button size="mini" type="success" @click="doUpdateSubGroup(scope)" plain
                >更新子组</el-button
              >
            </div>
            <div>
              <el-button size="mini" type="info" @click="doBlack(scope)" plain
                >黑名单</el-button
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
    <div class="block" v-if="pageshow">
      <el-pagination
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :current-page="currentPage"
        :page-sizes="[10, 15, 20, 30, 50]"
        :page-size="pageSize"
        layout="total, sizes, prev, pager, next, jumper"
        :total="pageTotal"
      >
      </el-pagination>
    </div>
    <el-dialog
      :title="dialogTitle"
      v-model="dialogVisible"
      width="450px"
      modal
      :before-close="handleClose"
    >
      <span>
        <el-form
          label-position="right"
          :rules="rules"
          ref="addJobInfo"
          :model="addJobInfo"
          label-width="110px"
          size="small"
        >
          <el-form-item label="分组名：" prop="name">
            <el-input style="width: 250px" v-model="addJobInfo.name"></el-input>
          </el-form-item>
          <el-form-item label="端口号：" prop="port">
            <template #label>
              <span>端口号</span>
              <el-tooltip
                content="如果端口号为0，在生成配置时忽略端口号，只保留IP地址"
                placement="top"
                style="diaplay: inline"
              >
                <span
                  ><i style="font-size: 13px" class="el-icon-warning"
                    >：</i
                  ></span
                >
              </el-tooltip>
            </template>
            <el-input
              type="number"
              style="width: 250px"
              v-model.number="addJobInfo.port"
            ></el-input>
          </el-form-item>
          <el-form-item label="重写标签：" prop="relabel_id">
            <el-select
              v-model="addJobInfo.relabel_id"
              filterable
              allow-create
              default-first-option
              placeholder="请选择"
              style="width: 250px"
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
    <el-dialog
      :title="'IP池 - ' + editIPDialog"
      v-model="editIPVisible"
      width="700px"
      modal
      :before-close="clearIPClose"
    >
      <el-row type="flex" align="middle" justify="center">
        <el-transfer
          v-model="currentIPValue"
          filterable
          :filter-method="filterIPMethod"
          filter-placeholder="请输入关键字"
          :data="allIPData"
          @change="transferChange"
          :titles="['IP池', 'Job组IP列表']"
        >
        </el-transfer>
      </el-row>
      <div class="ip-list-push-box">
        <el-button
          class="ip-list-close-btn"
          type="info"
          size="small"
          @click="clearIPClose"
          >关闭</el-button
        >
        <el-button
          class="ip-list-push-btn"
          type="warning"
          size="small"
          @click="updateJobIPList"
          >提交</el-button
        >
      </div>
    </el-dialog>
    <el-dialog
      :title="'JOB组内编辑 - ' + editIPInJobDialog"
      v-model="editIPInJobVisible"
      width="700px"
      modal
      :before-close="clearIPInJobClose"
    >
      <el-row type="flex" align="middle" justify="center">
        <el-transfer
          v-model="currentIPInJobValue"
          filterable
          :filter-method="filterIPMethod"
          filter-placeholder="请输入关键字"
          :data="allIPInJobData"
          @change="transferInJobChange"
          :titles="['JOB组IP池', '加黑IP列表']"
        >
        </el-transfer>
      </el-row>
      <div class="ip-list-push-box">
        <el-button
          class="ip-list-close-btn"
          type="info"
          size="small"
          @click="clearIPInJobClose"
          >关闭</el-button
        >
        <el-button
          class="ip-list-push-btn"
          type="warning"
          size="small"
          @click="updateJobIPInJobList"
          >提交</el-button
        >
      </div>
    </el-dialog>
    <el-dialog
      :title="editIPListDialog"
      v-model="editIPListVisible"
      width="900px"
      modal
      :before-close="clearIPListClose"
      custom-class="editIPListDialog"
    >
    <span><el-tag type="warning">IP以半角分号（;）隔开，只会添加在“IP管理”中已经有的IP地址</el-tag></span>
      <div>
        <el-input
          type="textarea"
          :autosize="{ minRows: 20, maxRows: 20}"
          placeholder="请输入内容"
          v-model="ipListConetnt">
        </el-input>
      </div>
      <div class="ip-list-push-box">
        <el-button
          class="ip-list-close-btn2"
          type="info"
          size="small"
          @click="clearIPListClose"
          >关 闭</el-button
        >
        <el-button
          class="ip-list-push-btn2"
          type="warning"
          size="small"
          @click="doUpdateAction"
          >提 交</el-button
        >
      </div>
    </el-dialog>
    <div class="dialog-relabel-ruleshow">
      <el-dialog
        :title="'规则名称：' + currentJobReWriteRule +'，规则内容：'"
        v-model="dialogJobReWriteRuleVisible"
        width="700px"
      >
        <!-- <div style="width: 550px">
          <el-scrollbar height="300px">
            <div v-if="showError">
            <pre v-highlight="error"><code></code></pre>
          </div>
            <JsonView :json="reWriteRuleInfo"></JsonView>
          </el-scrollbar>
        </div> -->
        <div style="width: 670px">
          <el-scrollbar height="300px" class="flex-content">
            <pre v-highlight="reWriteRuleInfo"><code></code></pre>
          </el-scrollbar>
        </div>
      </el-dialog>
    </div>
  </div>
</template>

<script>
import {
  getJobsWithSplitPage,
  postJob,
  putJob,
  deleteJob,
  swapJob,
  publishJobs,
  enabledJob,
  updateIPForJob,
  updateIPV2ForJob,
  updateIPInJobForJob,
  getAllReLabels,
  updateSubGroup,
  batchDeleteJob,
  batchEnableJob,
  batchDisableJob,
} from '@/api/jobs.js'
import { allIPList } from '@/api/machines.js'
import { restartSrv } from '@/api/srv'
import { 
  getJobMachines, 
  getJobMachinesBlack, 
  putJobMachinesBlack,
} from '@/api/labelsJob.js'

import {getReLabel} from '@/api/relabel.js'

export default {
  name: 'Jobs',
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
      pageSize: 20,
      pageTotal: 0,
      currentPage: 1,
      searchContent: '',
      deleteVisible: {},
      dialogVisible: false,
      buttonTitle: '',
      dialogTitle: '',
      jobAllIPList: {},
      currentIPValue: [],
      currentIPInJobValue: [],
      allIPData: [],
      allIPDataMap: {},
      allIPInJobData: [],
      allIPInJobDataMap: {},
      currentJobId: 0,
      editIPVisible: false,
      editIPListVisible: false,
      editIPInJobVisible: false,
      transferChanged: false,
      transferInJobChanged: false,
      editIPDialog: '未设置',
      editIPListDialog: '未设置',
      ipListConetnt: '',
      publishMode: false,
      pageshow: true,
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
      },
      flags: '',
      multipleSelection: [],
      editIPInJobDialog: '',
      editIPInJobVisible: false,
      currentJobReWriteRule: '',
      reWriteRuleInfo: null,
      dialogJobReWriteRuleVisible: false,
    }
  },
  created () {
    // this.doGetJobs()
  },
  mounted () {
    if (this.$route.params.currentPage) {
      this.currentPage = parseInt(this.$route.params.currentPage)
    }
    if (this.$route.params.pageSize) {
      this.pageSize = parseInt(this.$route.params.pageSize)
    }
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
    doEditLabelsJob (scope) {
      const jobInfo = {
        ...scope.row,
        currentPage: this.currentPage,
        pageSize: this.pageSize
      }
      this.$router.push({ name: 'labelsJobs', params: jobInfo })
    },
    doGetJobs (getInfo) {
      if (!getInfo) {
        getInfo = {
          'pageNo': this.currentPage,
          'pageSize': this.pageSize,
          'search': this.searchContent
        }
      }
      getJobsWithSplitPage(getInfo).then(
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
          this.pageshow = false;//让分页隐藏
          this.$nextTick(function () {//重新渲染分页
            this.pageshow = true;
          });
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
            postJob(postData).then(
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
            putJob(postData).then(
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
      deleteJob(scope.row.id).then(
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
    handleSelectionChange (val) {
      this.multipleSelection = []
      val.forEach(each => {
        this.multipleSelection.push(each.id)
      })
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
        batchDeleteJob(this.multipleSelection).then(r => {
          this.$notify({
            title: '成功',
            message: '删除所选项成功！',
            type: 'success'
          });
          this.doGetJobs()
        }).catch(e => console.log(e))
      }).catch(e => console.log(e))
    },
    doBatchDisable(){
      if (this.multipleSelection.length === 0) {
        this.$notify({
          title: '警告',
          message: '未选中任何项！',
          type: 'warning'
        });
        return false
      }
      this.$confirm('是否确定禁用？', '确认信息', {
        distinguishCancelAndClose: true,
        confirmButtonText: '确定',
        cancelButtonText: '放弃'
      }).then(_ => {
        batchDisableJob(this.multipleSelection).then(r => {
          this.$notify({
            title: '成功',
            message: '禁用所选项成功！',
            type: 'success'
          });
          this.doGetJobs()
        }).catch(e => console.log(e))
      }).catch(e => console.log(e))
    },
    doBatchEnable(){
      if (this.multipleSelection.length === 0) {
        this.$notify({
          title: '警告',
          message: '未选中任何项！',
          type: 'warning'
        });
        return false
      }
      this.$confirm('是否确定启用？', '确认信息', {
        distinguishCancelAndClose: true,
        confirmButtonText: '确定',
        cancelButtonText: '放弃'
      }).then(_ => {
        batchEnableJob(this.multipleSelection).then(r => {
          this.$notify({
            title: '成功',
            message: '启用所选项成功！',
            type: 'success'
          });
          this.doGetJobs()
        }).catch(e => console.log(e))
      }).catch(e => console.log(e))
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
    publishJobsRunning () {
      this.$notify({
        title: '警告',
        message: '正在发布，请稍等...',
        type: 'warning'
      });
    },
    publishJobsfunc () {
      this.publishMode = true
      publishJobs().then(
        r => {
          this.$notify({
            title: '成功',
            message: '发布成功！',
            type: 'success'
          });
          this.publishMode = false
          this.doGetJobs()
        }
      ).catch(
        e => {
          console.log(e)
          this.publishMode = false
        }
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
    expandChange (row, expandRows) {
      if (!expandRows || expandRows.length === 0) {
        return
      }
      if (this.jobAllIPList[row.id]) {
        return
      }
      getJobMachines(row.id).then(r => {
        this.jobAllIPList[row.id] = r.data
      }).catch(e => console.log(e))
    },
    filterIPMethod (query, item) {
      return item.spell.indexOf(query) > -1;
    },
    clearIPClose (scope) {
      if (this.transferChanged) {
        this.doGetJobs()
      }
      this.editIPVisible = false
    },
    clearIPInJobClose(){
      if (this.transferInJobChanged) {
        this.doGetJobs()
      }
      this.editIPInJobVisible = false
    },
    clearIPListClose(scope){
      this.editIPListVisible = false
      this.flags = ''
    },
    doEditPool(scope){
      getJobMachines(scope.row.id).then(r => {
        let currentIPValue = []
        r.data.forEach(item => {
          currentIPValue.push(item.ipaddr)
        })
        this.ipListConetnt = currentIPValue.join(';')
        this.currentJobId = scope.row.id
        this.editIPListDialog = '编辑IP列表 - ' + scope.row.name
        this.editIPListVisible = true
      }).catch(
        e => console.log(e)
      )
    },
    doPool (scope) {
      getJobMachines(scope.row.id).then(r => {
        let currentIPValue = []
        r.data.forEach(item => {
          currentIPValue.push(item.id)
        })
        let allIPData = []
        allIPList().then(r => {
          r.data.forEach(item => {
            let defVal = {
              label: item.ipaddr,
              key: item.id,
              spell: item.ipaddr,
              disabled: false
            }
            allIPData.push(defVal)
          })
          this.currentJobId = scope.row.id
          this.allIPData = allIPData
          this.currentIPValue = currentIPValue
          this.transferChanged = false
          this.editIPDialog = scope.row.name
          this.editIPVisible = true
        }).catch(e => console.log(e))
      }).catch(e => console.log(e))
    },
    doUpdateSubGroup(scope){
      updateSubGroup(scope.row.id).then(r=>{
        this.$notify({
          title: '成功',
          message: '整理成功！',
          type: 'success'
        })
        this.doGetJobs()
      }).catch(e=>console.log(e))
    },
    doBlack(scope){
      let currentIPBlackedValue = []
      let allIPInJobData = []
      getJobMachines(scope.row.id).then(r => {
        r.data.forEach(item => {
          let defVal = {
            label: item.ipaddr,
            key: item.id,
            spell: item.ipaddr,
            disabled: false
          }
          allIPInJobData.push(defVal)
          if (item.blacked === 1) {
            currentIPBlackedValue.push(item.id)
          }
        })
        this.currentJobId = scope.row.id
        this.allIPInJobData = allIPInJobData
        this.currentIPInJobValue = currentIPBlackedValue
        this.transferInJobChanged = false
        this.editIPInJobDialog = scope.row.name
        this.editIPInJobVisible = true
      }).catch(e => console.log(e))
    },
    updateJobIPList () {
      let clearInfo = {
        job_id: this.currentJobId,
        machines_ids: []
      }
      this.currentIPValue.forEach(item => {
        if (item) {
          clearInfo.machines_ids.push(item)
        }
      })
      updateIPForJob(clearInfo).then(r => {
        this.$notify({
          title: '成功',
          message: '整理成功！',
          type: 'success'
        })
        this.doGetJobs()
      }).catch(e => console.log(e))
    },
    updateJobIPInJobList(){
      let clearInfo = {
        job_id: this.currentJobId,
        machines_ids: []
      }
      this.currentIPInJobValue.forEach(item => {
        if (item) {
          clearInfo.machines_ids.push(item)
        }
      })
      updateIPInJobForJob(clearInfo).then(r => {
        this.$notify({
          title: '成功',
          message: '整理成功！',
          type: 'success'
        })
        this.doGetJobs()
      }).catch(e => console.log(e))
    },
    doUpdateAction(){
      if(this.flags === 'black') {
        this.updateJobIPListV2Black()
      } else {
        this.updateJobIPListV2()
      }
    },
    updateJobIPListV2(){
      updateIPV2ForJob(this.ipListConetnt, this.currentJobId).then(r => {
        this.$notify({
          title: '成功',
          message: '整理成功！',
          type: 'success'
        })
        this.doGetJobs()
      }).catch(e => console.log(e))
    },
    updateJobIPListV2Black(){
      putJobMachinesBlack(this.currentJobId, this.ipListConetnt).then(r => {
        this.$notify({
          title: '成功',
          message: '整理黑名单成功！',
          type: 'success'
        })
        this.doGetJobs()
      }).catch(e => console.log(e))
    },
    transferChange (value, direction, movedKeys) {
      this.transferChanged = true
    },
    invocate (scope) {
      const newStatus = !this.jobs[scope.$index].enabled
      const jInfo = {
        id: scope.row.id,
        enabled: newStatus
      }
      enabledJob(jInfo).then(r => {
        this.$notify({
          title: '成功',
          message: '更新状态成功！',
          type: 'success'
        });
        this.jobs[scope.$index].enabled = newStatus
        this.jobs = [...this.jobs]
      }).catch(e => console.log(e))
    },
    clickcell(row, column, cell, event){
      if (column.property === 'relabel_name') {
        getReLabel({id: row.relabel_id}).then(r=>{
          this.currentJobReWriteRule = r.data.name
          this.reWriteRuleInfo = r.data.code
          this.dialogJobReWriteRuleVisible = true
        }).catch(e=>console.log(e))
      }
    },
  }
}
</script>

<style scoped>
.do_action {
  display: flex;
  flex-wrap: nowrap;
  justify-content: space-between;
  margin-top: -5px;
  padding-bottom: 8px;
}
.do-action-search {
  display: inline-block;
  /* display: flex;
  flex-wrap: nowrap;
  justify-content: flex-end; */
  /* margin-top: -5px;
  padding-bottom: 8px; */
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
  justify-content: space-between;
}
.change_order_button {
  display: flex;
  flex-wrap: nowrap;
}
.change_order_button .el-button {
  width: 20px;
}
.change_order_button .el-button:nth-child(1) {
  margin: 0px 0px 0px -10px;
}
.change_order_button .el-button:nth-child(2) {
  margin: 0px 0px 0px 0px;
}
.jobs-board :deep() .el-transfer {
  margin-top: -15px;
  margin-bottom: 20px;
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
}
.jobs-board :deep() .el-transfer-panel {
  width: 250px;
}
.jobs-board :deep() .el-transfer__buttons {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}
.jobs-board :deep() .el-transfer__button {
  padding: 10px 15px;
}
.jobs-board :deep() .el-transfer__button {
  margin-left: 0px;
}
.ip-list-push-box {
  display: float;
  text-align: right;
}
.ip-list-push-btn {
  /* margin-right: 20px; */
}
.ip-list-close-btn {
  margin-right: 8px;
}
.ip-list-push-btn2 {
  margin-top: 7px;
  /* margin-right: 26px; */
}
.ip-list-close-btn2 {
  margin-top: 7px;
  margin-right: 8px;
}
.form-inline-label {
  display: flex;
  flex-direction: row;
  margin-left: 0px;
}
:deep() .editIPListDialog > div.el-dialog__body {
  padding-top: 6px;
}
.dialog-relabel-ruleshow :deep() .el-dialog__body {
  padding-top: 5px;
  /* padding-bottom: 5px; */
}
</style>
