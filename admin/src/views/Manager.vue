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
        <el-button
          icon="el-icon-download"
          size="small"
          type="text"
          @click="doOutputAllIP()"
          >导出所有IP</el-button
        >
      </div>
      <div class="do_action">
        <div style="padding-right: 15px">
          <el-button size="small" icon="el-icon-basketball" type="warning" plain @click="doBatchAddDomainWeb()"
            >批量添加域名</el-button
          >
          <el-button size="small" icon="el-icon-football" type="warning" plain @click="doBatchAddWeb()"
            >批量添加IP</el-button
          >
          <el-button size="small" icon="el-icon-soccer" type="warning" plain @click="doBatchAddFile()"
            >批量添加IP(文件)</el-button
          >
          <el-button size="small" icon="el-icon-baseball" type="success" plain @click="doAdd()"
            >添加IP</el-button
          >
          <el-button
            v-if="pushing === false"
            size="small"
            icon="el-icon-upload"
            type="primary"
            plain
            @click="onUpdatePosition()"
            >更新所有IP位置</el-button
          >
          <el-button
            v-if="pushing === true"
            size="small"
            icon="el-icon-loading"
            type="primary"
            plain
            >更新所有IP位置</el-button
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
                @change="selectOptionChange"
              >
                <el-option label="所有" value="all"></el-option>
                <el-option label="启用" value="enable"></el-option>
                <el-option label="禁用" value="disable"></el-option>
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
      @cell-click="clickcell"
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
        label="所在机房"
        prop="idc_name"
        align="center"
        header-align="center"
        show-overflow-tooltip
      >
      </el-table-column>
      <el-table-column
        label="线路"
        prop="line_name"
        align="center"
        header-align="center"
        show-overflow-tooltip
      >
      </el-table-column>
      <el-table-column
        label="定位"
        prop="position"
        align="center"
        header-align="center"
        width="260px"
        show-overflow-tooltip
      >
        <template v-slot="scope">
          <span v-if="scope.row.position.error"><el-tag type="danger" size="mini" >{{scope.row.position.error}}</el-tag></span>
          <span v-else-if="scope.row.position.isp"><el-tag size="mini" type="info">{{scope.row.position.country}}{{scope.row.position.province}}{{scope.row.position.city}}{{scope.row.position.county}}</el-tag><el-tag type="warning" size="mini" >{{scope.row.position.isp}}</el-tag></span>
          <span v-else><el-tag size="mini" type="info">{{scope.row.position.country}}{{scope.row.position.province}}{{scope.row.position.city}}{{scope.row.position.county}}{{scope.row.position.isp}}</el-tag></span>
        </template>
      </el-table-column>
      <el-table-column
        label="分组选项"
        prop="jobs_id"
        align="center"
        header-align="center"
        width="180px"
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
        width="120px"
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
              >状态详情</el-button
            >
            <el-button
              v-else
              type="text"
              size="mini"
              icon="el-icon-sunny"
              @click="displayIPStatusDialog(row)"
              >状态详情</el-button
            >
          </el-tooltip>
        </template>
      </el-table-column>
      <el-table-column
        label="最后更新账号"
        prop="update_by"
        align="center"
        header-align="center"
        width="100px"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="最后更新时间"
        prop="update_at"
        align="center"
        header-align="center"
        width="140px"
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
        :page-sizes="[15, 20, 50, 100, 200]"
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
    <div class="dialog-ip-position">
      <el-dialog
        :title="'IP位置信息：' + currentIPPosition"
        v-model="dialogIpPositionVisible"
        width="600px"
      >
        <div style="width: 550px">
          <el-scrollbar height="300px">
            <!-- <div v-if="showError">
            <pre v-highlight="error"><code></code></pre>
          </div> -->
            <JsonView :json="jsonPositionInfo"></JsonView>
          </el-scrollbar>
        </div>
      </el-dialog>
    </div>
    <div>
      <el-dialog
        title="批量导入IP地址"
        v-model="dialogBatchImportVisible"
        width="900px"
      >
        <div class="prompt-message"><el-tag size="mini" type="warning">IP或者网段之间，使用英文分号（;）分隔开或者单独一行，自动跳过已经存在的IP地址，如：</el-tag><br><el-tag size="mini">192.168.1.0/24;10.10.10.1;172.16.1.1~172.16.2.1;172.16.1.1-172.16.2.1</el-tag></div>
        <el-form :model="batchImportIPAddr" :rules="webBatchRules" ref="batchImportIPAddr">
          <el-form-item prop="content">
            <el-input 
              :autosize="{ minRows: 15, maxRows: 15 }" 
              type="textarea" 
              size="small" 
              placeholder="" 
              v-model="batchImportIPAddr.content" 
              autocomplete="off">
            </el-input>
          </el-form-item>
        </el-form>
        <div>
          <span>为以上新添加的地址分配到指定的组：</span>
          <el-select
            v-model="batchTypeSelect"
            multiple
            collapse-tags
            style="margin-left: 20px"
            placeholder="Select"
            size="small"
          >
            <el-option
              v-for="item in jobs"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            >
            </el-option>
          </el-select>
        </div>
        <template #footer>
          <span class="dialog-footer">
            <el-button size="small" @click="dialogBatchImportVisible = false">关 闭</el-button>
            <el-button
              v-if="import_pushing === false"
              style="margin-left: 10px"
              size="small"
              type="success"
              icon="el-icon-upload"
              @click="doBatchAddWebSubmit('batchImportIPAddr')"
              >确 定</el-button
            >
            <el-button
              v-if="import_pushing === true"
              style="margin-left: 10px"
              size="small"
              type="success"
              icon="el-icon-loading"
              >确 定</el-button
            >
            <!-- <el-button size="small" type="primary" @click="doBatchAddWebSubmit('batchImportIPAddr')"
              >确 定</el-button
            > -->
          </span>
        </template>
      </el-dialog>
    </div>
    <div>
      <el-dialog
        title="批量导入域名"
        v-model="dialogBatchImportDomainVisible"
        width="900px"
      >
        <div class="prompt-message"><el-tag type="warning">域名之间使用英文分号(;)隔开，也可以包括IP，服务器不对导入的内容做检查</el-tag></div>
        <el-form :model="batchImportDomain" :rules="webBatchDomainRules" ref="batchImportDomain">
          <el-form-item prop="content">
            <el-input 
              :autosize="{ minRows: 15, maxRows: 15 }" 
              type="textarea" 
              size="small" 
              placeholder="" 
              v-model="batchImportDomain.content" 
              autocomplete="off">
            </el-input>
          </el-form-item>
        </el-form>
        <div>
          <span>为以上新添加的地址分配到指定的组：</span>
          <el-select
            v-model="batchDomainTypeSelect"
            multiple
            collapse-tags
            style="margin-left: 20px"
            placeholder="Select"
            size="small"
          >
            <el-option
              v-for="item in jobs"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            >
            </el-option>
          </el-select>
        </div>
        <template #footer>
          <span class="dialog-footer">
            <el-button size="small" @click="dialogBatchImportDomainVisible = false">关 闭</el-button>
            <el-button
              v-if="import_domain_pushing === false"
              style="margin-left: 10px"
              size="small"
              type="success"
              icon="el-icon-upload"
              @click="doBatchDomainAddWebSubmit('batchImportDomain')"
              >确 定</el-button
            >
            <el-button
              v-if="import_domain_pushing === true"
              style="margin-left: 10px"
              size="small"
              type="success"
              icon="el-icon-loading"
              >确 定</el-button
            >
            <!-- <el-button size="small" type="primary" @click="doBatchDomainAddWebSubmit('batchImportDomain')"
              >确 定</el-button
            > -->
          </span>
        </template>
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
  batchDeleteMachine,
  batchEnableMachine,
  batchDisableMachine,
  updatePosition,
  batchImportIpAddrsWeb,
  batchImportDomainWeb,
  outputAllIP
} from '@/api/machines'
import { publish } from '@/api/publish'
import JsonView from '@/components/JsonView.vue'

import { getStorageVal, setStorageVal } from '@/utils/localStorage.js'

export default {
  name: 'Manager',
  components: {
    JsonView: JsonView
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
    };
    return {
      machines: [],
      jobs: [],
      jobsMap: {},
      selectTypeValue: {},
      batchTypeSelect: [],
      batchDomainTypeSelect: [],
      search: '',
      pageSize: 20,
      pageTotal: 0,
      currentPage: 1,
      selectOption: 'enable',
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
      dialogIpPositionVisible: false,
      dialogBatchImportVisible: false,
      dialogBatchImportDomainVisible: false,
      currentIPStatus: '',
      currentIPPosition: '',
      ipStatusList: {},
      needWarning: {},
      jsonPositionInfo: null,
      batchIpaddrs: '',
      batchImportIPAddr: {
        content: '',
      },
      batchImportDomain: {
        content: '',
      },
      formLabelWidth: '80px',
      webBatchRules: {
        content: [
          { required: true, message: '请输入有效的IP地址或者网段，支持V4及V6', trigger: ['blur'] }
        ],
      },
      webBatchDomainRules: {
        content: [
          { required: true, message: '请输入内容', trigger: ['blur'] }
        ],
      },
      import_pushing: false,
      import_domain_pushing: false,
    }
  },
  created () {
    // this.doGetMechines()
  },
  mounted () {
    let selectOption = getStorageVal('displayIPType')
    if (selectOption === null) {
      this.selectOption = 'enable'
    } else {
      this.selectOption = selectOption
    }
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
    onUpdatePosition(){
      this.pushing = true
      updatePosition().then(r=>{
        this.doGetMechines()
        this.pushing = false
      }).catch(e=>{
        console.log(e)
        this.pushing = false
      })
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
    doBatchAddFile () {
      const params = {
        currentPage: this.currentPage,
        pageSize: this.pageSize
      }
      this.$router.push({ name: 'BatchOpt', params: params })
    },
    doBatchAddWeb(){
      getJobs().then(
        r => {
          this.jobs = r.data
          r.data.forEach(e => {
            this.jobsMap[e.id] = e.name
          })
          this.import_pushing = false
          this.dialogBatchImportVisible = true
        }
      ).catch(
        e => {
          console.log(e)
        }
      )
    },
    doBatchAddDomainWeb(){
      getJobs().then(
        r => {
          this.jobs = r.data
          r.data.forEach(e => {
            this.jobsMap[e.id] = e.name
          })
          this.import_domain_pushing = false
          this.dialogBatchImportDomainVisible = true
        }
      ).catch(
        e => {
          console.log(e)
        }
      )
    },
    doBatchAddWebSubmit(formName){
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.import_pushing = true
          var jobIDSelect = []
          for(var k in this.batchTypeSelect) {
            jobIDSelect.push(this.batchTypeSelect[k])
          }
          batchImportIpAddrsWeb({content: this.batchImportIPAddr.content, jobs_id: jobIDSelect}).then(r=>{
            this.$notify({
              title: '成功',
              message: '批量添加成功！',
              type: 'success'
            });
            this.doGetMechines()
            this.import_pushing = false
          }).catch(e=>console.log(e))
        } else {
          return false
        }
      })
    },
    doBatchDomainAddWebSubmit(formName){
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.import_domain_pushing = true
          var jobIDSelect = []
          for(var k in this.batchDomainTypeSelect) {
            jobIDSelect.push(this.batchDomainTypeSelect[k])
          }
          batchImportDomainWeb({content: this.batchImportDomain.content, jobs_id: jobIDSelect}).then(r=>{
            this.$notify({
              title: '成功',
              message: '批量添加成功！',
              type: 'success'
            });
            this.doGetMechines()
            this.import_domain_pushing = false
          }).catch(e=>console.log(e))
        } else {
          return false
        }
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
        batchDisableMachine(this.multipleSelection).then(r => {
          this.$notify({
            title: '成功',
            message: '禁用所选项成功！',
            type: 'success'
          });
          this.doGetMechines()
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
        batchEnableMachine(this.multipleSelection).then(r => {
          this.$notify({
            title: '成功',
            message: '启用所选项成功！',
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
    },
    clickcell(row, column, cell, event){
      if (column.property === 'position') {
        this.currentIPPosition = row.ipaddr
        this.jsonPositionInfo = row.position
        this.dialogIpPositionVisible = true
      }
    },
    selectChange (event) {},
    doOutputAllIP(){
      outputAllIP().then(r=>{
        let data = r.data.data;  //csv数据
        let filename = r.data.name  //导出的文件名
        let type = "";                      //头部数据类型
        let file = new Blob(["\ufeff" + data], { type: type });
        if (window.navigator.msSaveOrOpenBlob)
            // IE10+
            window.navigator.msSaveOrOpenBlob(file, filename);
          else {
            // Others
            let a = document.createElement("a"),
                url = URL.createObjectURL(file);
            a.href = url;
            a.download = filename;
            document.body.appendChild(a);
            a.click();
            setTimeout(function() {
              document.body.removeChild(a);
              window.URL.revokeObjectURL(url);
            }, 0);
          }
      })
    },
    selectOptionChange(val){
      setStorageVal('displayIPType', val)
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
.borderNone :deep() .el-input__inner {
  width: 130px;
  height: 28px;
  margin-top: 0px;
  margin-bottom: 0px;
  padding-top: 0px;
  padding-bottom: 0px;
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
  width: 160px;
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
.prompt-message {
  margin-top: -22px;
}
</style>
