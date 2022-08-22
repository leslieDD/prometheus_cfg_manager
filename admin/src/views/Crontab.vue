<template>
  <div class="crontab-board">
    <div class="do_action">
      <div class="action-btn-del">
        <el-button icon="el-icon-lightning" size="small" type="info" plain @click="doBatchDel()">删除选中</el-button>
        <el-button icon="el-icon-warning-outline" size="small" type="info" plain @click="doBatchDisable()">禁用选中
        </el-button>
        <el-button icon="el-icon-warning-outline" size="small" type="success" plain @click="doBatchEnable()">启用选中
        </el-button>
      </div>
      <div class="do-action-search">
        <span style="padding-right: 15px">
          <el-button v-if="pushing === false" type="primary" plain icon="el-icon-upload" size="small"
            @click="doRulesPublish">发布</el-button>
          <el-button v-if="pushing === true" type="primary" plain icon="el-icon-loading" size="small">发布</el-button>
        </span>
        <span style="padding-right: 15px">
          <el-button size="small" type="success" icon="el-icon-baseball" plain @click="doAdd()">添加任务</el-button>
        </span>
        <span>
          <el-input size="small" placeholder="请输入内容" @keyup.enter="onSearch()" v-model="searchContent"
            style="width: 300px">
            <template #append>
              <el-button size="small" @click="onSearch()" icon="el-icon-search"></el-button>
            </template>
          </el-input>
        </span>
      </div>
    </div>
    <el-table size="mini" highlight-current-row border :data="crontabRules" stripe :row-style="rowStyle"
      :cell-style="cellStyle" @selection-change="handleSelectionChange" @cell-click="clickcell">
      <el-table-column type="selection" width="40"> </el-table-column>
      <el-table-column label="序号" width="50px">
        <template v-slot="scope">
          {{ scope.$index + 1 }}
        </template>
      </el-table-column>
      <el-table-column label="名称" width="200px" prop="name" align="center" header-align="center" show-overflow-tooltip>
      </el-table-column>
      <el-table-column label="规则" prop="rule" align="center" header-align="center">
        <template v-slot="{ row }">
          <el-tooltip placement="top">
            <template #content>
              <!-- <div class="rule-box" v-html="ToBreak(row.rule)"></div> -->
              <div class="rule-box">{{ row.rule }}</div>
            </template>
            <div class="oneLine">{{ row.rule }}</div>
          </el-tooltip>
        </template>
      </el-table-column>
      <el-table-column label="接口" prop="api_id" width="150px" align="center" header-align="center"
        show-overflow-tooltip>
        <template v-slot="{ row }">
          <span>{{ getApiName(row.api_id) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="执行周期" prop="exec_cycle" align="center" header-align="center" width="150px"
        show-overflow-tooltip>
      </el-table-column>
      <el-table-column label="执行次数" prop="run_times" align="center" header-align="center" width="75px">
      </el-table-column>
      <el-table-column label="成功次数" prop="success_times" align="center" header-align="center" width="75px">
      </el-table-column>
      <el-table-column label="失败次数" prop="fail_times" align="center" header-align="center" width="75px">
      </el-table-column>
      <el-table-column label="更新账号" prop="update_by" width="100px" align="center" header-align="center"
        show-overflow-tooltip></el-table-column>
      <el-table-column label="更新时间" prop="update_at" width="140px" align="center" header-align="center">
        <template v-slot="{ row }">
          <span>{{ parseTimeSelf(row.update_at) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" width="250px">
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
    <div class="block" v-if="pageshow">
      <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange" :current-page="currentPage"
        :page-sizes="[15, 20, 30, 50]" :page-size="pageSize" layout="total, sizes, prev, pager, next, jumper"
        :total="pageTotal">
      </el-pagination>
    </div>
    <el-dialog :title="dialogTitle" v-model="dialogVisible" width="700px" modal :before-close="handleClose">
      <span>
        <el-form label-position="right" :rules="addCronRuleInfoRule" ref="addCronRuleInfo" :model="addCronRuleInfo"
          label-width="80px" size="small">
          <el-form-item label="名称：" prop="name">
            <el-input style="width: 550px" v-model="addCronRuleInfo.name"></el-input>
          </el-form-item>
          <el-form-item label="周期：" prop="exec_cycle">
            <div class="crontab-edit">
              <el-popover v-model:visible="cronPopover" width="700px" trigger="manual">
                <vue3Cron @change="changeCron" @close="togglePopover(false)" max-height="400px" i18n="cn"></vue3Cron>
                <template #reference>
                  <el-input @focus="togglePopover(true)" v-model="addCronRuleInfo.exec_cycle" style="width: 250px"
                    placeholder="* * * * * ? *">
                  </el-input>
                </template>
              </el-popover>
              <!-- <el-button size="small" width="150px" type="primary" @click="cronPopover = true">设置任务定时执行周期</el-button> -->
              <!-- <el-input style="width: 200px" v-model="addCronRuleInfo.exec_cycle"></el-input><span> 秒
            </span><span>{{ convertData(addCronRuleInfo.exec_cycle) }}</span> -->
            </div>
          </el-form-item>
          <el-form-item label="规则：" prop="rule">
            <el-input type="textarea" :autosize="{ minRows: 10, maxRows: 10 }" style="width: 550px"
              v-model="addCronRuleInfo.rule"></el-input>
          </el-form-item>
          <el-form-item label="接口：" prop="api_id">
            <el-select v-model="addCronRuleInfo.api_id" style="width: 350px" collapse-tags placeholder="Select"
              size="small">
              <el-option v-for="item in cronApiList" :key="item.id" :label="item.name" :value="item.id">
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="状态：" prop="enabled">
            <el-switch v-model="addCronRuleInfo.enabled" active-color="#13ce66" inactive-color="#ff4949"></el-switch>
          </el-form-item>
          <el-form-item size="small">
            <el-button size="small" type="primary" @click="onSubmit('addCronRuleInfo')">{{ buttonTitle }}</el-button>
            <el-button size="small" type="info" @click="onCancel('addCronRuleInfo')">取消</el-button>
          </el-form-item>
        </el-form>
      </span>
    </el-dialog>
  </div>
</template>

<script>

import vue3Cron from '@/components/vue3-cron.vue'

import {
  getCronRulesWithSplitPage,
  postCronRule,
  putCronRule,
  deleteCronRule,
  enabledCronRule,
  getAllBaseCronApi,
  batchDeleteCronRules,
  batchEnableCronRules,
  batchDisableCronRules,
  cronRulesPublish,
} from '@/api/cron.js'

export default {
  components: { vue3Cron },
  name: 'Crontab',
  data () {
    return {
      pageSize: 20,
      pageTotal: 0,
      currentPage: 1,
      pageshow: true,
      multipleSelection: [],
      crontabRules: [],
      searchContent: '',
      deleteVisible: {},
      dialogVisible: false,
      dialogTitle: '',
      buttonTitle: '',
      cronApiList: [],
      cronApiMap: {},
      // batchJobSelect: '',
      addCronRuleInfo: {
        'id': 0,
        'name': '',
        'rule': '',
        'api_id': null,
        'enabled': true,
        'exec_cycle': "",
      },
      addCronRuleInfoRule: {
        name: [
          { required: true, message: '请输入名称', trigger: ['blur'] }
        ],
        rule: [
          { required: true, message: '请输入监控规则', trigger: ['blur'] }
        ],
        api_id: [
          { required: true, message: '请选择接口地址', trigger: ['blur'] }
        ],
        enabled: [
          { required: true, message: '请设置状态', trigger: ['blur'] }
        ],
        exec_cycle: [
          { required: true, message: '请输入执行周期', trigger: ['blur'] }
        ]
      },
      cronPopover: false,
      pushing: false,
    }
  },
  created () {
  },
  mounted () {
    if (this.$route.params.currentPage) {
      this.currentPage = parseInt(this.$route.params.currentPage)
    }
    if (this.$route.params.pageSize) {
      this.pageSize = parseInt(this.$route.params.pageSize)
    }
    this.doGetCronRules()
  },
  methods: {
    doGetCronRules (getInfo) {
      getAllBaseCronApi().then(r => {
        this.cronApiList = r.data
        this.convertcronApiToMap(r.data)
        if (!getInfo) {
          getInfo = {
            'pageNo': this.currentPage,
            'pageSize': this.pageSize,
            'search': this.searchContent
          }
        }
        getCronRulesWithSplitPage(getInfo).then(
          r => {
            let n = 0
            r.data.data.forEach(element => {
              this.deleteVisible[n] = false
              n += 1
            })
            this.crontabRules = r.data.data
            this.pageTotal = r.data.totalCount
            this.currentPage = r.data.pageNo
            this.pageSize = r.data.pageSize
            this.pageshow = false;//让分页隐藏
            this.$nextTick(function () {//重新渲染分页
              this.pageshow = true;
            });
          }
        ).catch(e => console.log(e))
      }).catch(e => console.log(e))
    },
    convertcronApiToMap (list) {
      let apiMap = {}
      list.forEach(each => {
        apiMap[each.id] = each
      })
      this.cronApiMap = apiMap
    },
    getApiName (apiID) {
      if (this.cronApiMap[apiID]) {
        return this.cronApiMap[apiID].name
      } else {
        return "unknow"
      }
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
        batchDeleteCronRules(this.multipleSelection).then(r => {
          this.$notify({
            title: '成功',
            message: '删除所选项成功！',
            type: 'success'
          });
          this.doGetCronRules()
        }).catch(e => console.log(e))
      }).catch(e => console.log(e))
    },
    ToBreak (val) {
      return val.replaceAll('\n', '<br />')
    },
    doRulesPublish () {
      this.pushing = true
      cronRulesPublish().then(r => {
        this.$notify({
          title: '成功',
          message: '发布成功！',
          type: 'success'
        });
        this.pushing = false
      }).catch(e => {
        console.log(e)
        this.pushing = false
      })
    },
    convertData (val) {
      var reg = /^[\d]+$/;
      if (!reg.test(val)) {
        return "请输入数字"
      }
      let newVal = val
      return this.secondsFormat(newVal)
    },
    secondsFormat (s) {
      var day = Math.floor(s / (24 * 3600)) // Math.floor()向下取整
      var hour = Math.floor((s - day * 24 * 3600) / 3600)
      var minute = Math.floor((s - day * 24 * 3600 - hour * 3600) / 60)
      var second = s - day * 24 * 3600 - hour * 3600 - minute * 60
      return day + "天" + hour + "时" + minute + "分" + second + "秒"
    },
    togglePopover (val) {
      this.cronPopover = val
    },
    changeCron (val) {
      if (typeof val !== 'string') {
        return false
      }
      this.addCronRuleInfo.exec_cycle = val
    },
    doBatchDisable () {
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
        batchDisableCronRules(this.multipleSelection).then(r => {
          this.$notify({
            title: '成功',
            message: '禁用所选项成功！',
            type: 'success'
          });
          this.doGetCronRules()
        }).catch(e => console.log(e))
      }).catch(e => console.log(e))
    },
    doBatchEnable () {
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
        batchEnableCronRules(this.multipleSelection).then(r => {
          this.$notify({
            title: '成功',
            message: '启用所选项成功！',
            type: 'success'
          });
          this.doGetCronRules()
        }).catch(e => console.log(e))
      }).catch(e => console.log(e))
    },
    doAdd () {
      getAllBaseCronApi().then(r => {
        this.cronApiList = r.data
        this.convertcronApiToMap(r.data)
        this.addCronRuleInfo = {
          'id': 0,
          'name': '',
          'rule': '',
          'api_id': null,
          'enabled': true,
          'exec_cycle': "",
        }
        this.buttonTitle = '创建'
        this.dialogTitle = '增加定时任务'
        this.dialogVisible = true
      }).catch(e => {
        console.log(e)
      })
    },
    onSearch () {
      this.doGetCronRules()
    },
    doEdit (data) {
      getAllBaseCronApi().then(r => {
        this.cronApiList = r.data
        this.convertcronApiToMap(r.data)
        this.buttonTitle = '更新'
        this.dialogTitle = '编辑IP分组'
        this.addCronRuleInfo = {
          'id': data.row.id,
          'name': data.row.name,
          'rule': data.row.rule,
          'api_id': data.row.api_id,
          'enabled': data.row.enabled,
          'exec_cycle': data.row.exec_cycle,
        }
        this.dialogVisible = true
      }).catch(e => {
        console.log(e)
      })
    },
    doYes (scope) {
      deleteCronRule(scope.row.id).then(
        r => {
          this.$notify({
            title: '成功',
            message: '删除成功！',
            type: 'success'
          });
          this.doGetCronRules()
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
    invocate (scope) {
      const newStatus = !this.crontabRules[scope.$index].enabled
      const jInfo = {
        id: scope.row.id,
        enabled: newStatus
      }
      enabledCronRule(jInfo).then(r => {
        this.$notify({
          title: '成功',
          message: '更新状态成功！',
          type: 'success'
        });
        this.crontabRules[scope.$index].enabled = newStatus
        this.crontabRules = [...this.crontabRules]
      }).catch(e => console.log(e))
    },
    onSubmit (formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          let postData = {}
          postData['name'] = this.addCronRuleInfo.name
          postData['rule'] = this.addCronRuleInfo.rule
          postData['api_id'] = this.addCronRuleInfo.api_id
          postData['enabled'] = this.addCronRuleInfo.enabled
          postData['exec_cycle'] = this.addCronRuleInfo.exec_cycle
          if (this.buttonTitle === '创建') {
            postCronRule(postData).then(
              r => {
                this.$notify({
                  title: '成功',
                  message: '创建成功！',
                  type: 'success'
                });
                this.doGetCronRules()
                this.dialogVisible = false
                this.$refs[formName].resetFields()
              }
            ).catch(
              e => { console.log(e) }
            )
          } else {
            postData['id'] = this.addCronRuleInfo.id
            putCronRule(postData).then(
              r => {
                this.$notify({
                  title: '成功',
                  message: '更新成功！',
                  type: 'success'
                });
                this.doGetCronRules()
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
    handleSizeChange (val) {
      let getInfo = {
        'pageNo': this.currentPage,
        'pageSize': val,
        'search': this.searchContent
      }
      this.doGetCronRules(getInfo)
    },
    handleCurrentChange (val) {
      let getInfo = {
        'pageNo': val,
        'pageSize': this.pageSize,
        'search': this.searchContent
      }
      this.doGetCronRules(getInfo)
    },
    handleSelectionChange (val) {
      this.multipleSelection = []
      val.forEach(each => {
        this.multipleSelection.push(each.id)
      })
    },
    clickcell (row, column, cell, event) {
      if (column.property === 'relabel_name') {
        getReLabel({ id: row.relabel_id }).then(r => {
          this.currentJobReWriteRule = r.data.name
          this.reWriteRuleInfo = r.data.code
          this.dialogJobReWriteRuleVisible = true
        }).catch(e => console.log(e))
      }
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
    parseTimeSelf (t) {
      var time = new Date(Date.parse(t))
      return time.toLocaleDateString() + ' ' + time.toTimeString().split(' ')[0]
    },
    handleClose () {
      this.dialogVisible = false
      this.cronPopover = false
      this.$refs[formName].resetFields()
    }
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
  text-align: center;
}

.actioneara {
  display: flex;
  flex-wrap: nowrap;
  justify-content: space-around;
}

.block {
  padding-top: 12px;
}

.oneLine {
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
}

.rule-box {
  max-width: 400px;
}

.cron {
  width: 400px;
  margin: 0 auto;
  margin-top: 100px;
}

.cron>h1 {
  font-size: 50px;
  text-align: center;
}

.crontab-edit {
  display: flex;
  flex-wrap: nowrap;
  justify-content: left;
}
</style>
