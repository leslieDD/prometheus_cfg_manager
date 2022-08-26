<template>
  <div class="batch-box">
    <div class="batch-box-desc">
      <div class="return-back">
        <el-button size="small" type="primary" @click="goBack" plain>返回之前页面</el-button>
      </div>
      <div class="push-options">
        <el-card class="box-card" shadow="never">
          <template #header>
            <div class="card-header">
              <span>导入选项：</span>
            </div>
          </template>
          <el-form ref="form" :model="pushOption" label-width="auto" size="mini">
            <el-form-item label="当遇到错误，是否继续导入：">
              <el-tooltip :content="'当前: ' + ignoreErrStatus" placement="top">
                <el-switch v-model="pushOption.ignoreErr" active-color="#13ce66" inactive-color="#ff4949"
                  :active-value="true" :inactive-value="false">
                </el-switch>
              </el-tooltip>
            </el-form-item>
            <el-form-item label="为导入IP指定所属组：">
              <el-select v-model="selectTypeValue" class="borderNone" popper-class="pppselect"
                @change="selectChange($event)" size="small" multiple collapse-tags placeholder="请选择">
                <el-option v-for="item in jobs" :key="item.id" :label="item.name" :value="item.id">
                </el-option>
              </el-select>
            </el-form-item>
          </el-form>
        </el-card>
      </div>
      <div>
        <el-card class="box-card" shadow="never">
          <template #header>
            <div class="card-header">
              <span>导入表格格式如下：</span>
            </div>
          </template>
          <table class="pure-table">
            <thead>
              <tr>
                <th class="pure-table-th1">
                  <single-svg icon-class="excel" />
                </th>
                <th class="pure-table-th2">A</th>
                <th class="pure-table-th3">B</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(ipaddr, index) in exampleData" :key="index">
                <td class="pure-table-td1">{{ index + 1 }}</td>
                <td class="pure-table-td2">{{ ipaddr }}</td>
                <td class="pure-table-td2"></td>
              </tr>
            </tbody>
          </table>
        </el-card>
      </div>
    </div>
    <div class="split-line" />
    <div class="batch-box-data">
      <div class="ctl_btn_area">
        <div>
          <el-button icon="el-icon-lightning" size="small" type="info" plain @click="doBatchDel()">删除选中项</el-button>
        </div>
        <div class="do_action">
          <div style="padding-right: 15px">
            <el-upload class="upload-demo" ref="upload" action="" :auto-upload="false"
              accept=".csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet, application/vnd.ms-excel"
              :file-list="fileList" :show-file-list="false" :on-change="importExcel" type="file">
              <template #trigger>
                <el-button size="small" type="primary">导入文件</el-button>
              </template>
              <el-button v-if="pushing === false" style="margin-left: 10px" size="small" type="success"
                icon="el-icon-upload" @click="submitUpload">上传数据到服务器</el-button>
              <el-button v-if="pushing === true" style="margin-left: 10px" size="small" type="success"
                icon="el-icon-loading">上传数据到服务器</el-button>
              <template #tip>
                <div class="el-upload__tip">文件格式：xlxs, xls, cvs</div>
              </template>
            </el-upload>
          </div>
          <div>
            <el-input size="small" placeholder="请输入内容" @keyup.enter="onSearch()" v-model="searchContent">
              <template #append>
                <el-button size="small" @click="onSearch()" icon="el-icon-search"></el-button>
              </template>
            </el-input>
          </div>
        </div>
      </div>
      <div class="descriptions-display">
        <el-descriptions :column="4" size="mini">
          <el-descriptions-item label="IP总数：">
            <el-tag size="mini" type="success">{{
                ipTongJi.total
            }}</el-tag>
          </el-descriptions-item>
          <el-descriptions-item label="导入成功数：">
            <el-tag size="mini" type="success">{{
                ipTongJi.success
            }}</el-tag>
          </el-descriptions-item>
          <el-descriptions-item label="导入失败数：">
            <el-tag size="mini" type="danger">{{
                ipTongJi.fail
            }}</el-tag>
          </el-descriptions-item>
          <el-descriptions-item label="未操作数：">
            <el-tag size="mini" type="info">{{
                ipTongJi.noaction
            }}</el-tag>
          </el-descriptions-item>
        </el-descriptions>
      </div>
      <el-table size="mini" highlight-current-row border :data="uploadIPsSplit" stripe :row-style="rowStyle"
        :cell-style="cellStyle" header-align="center" @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="40"> </el-table-column>
        <el-table-column label="序号" width="50px" align="center" header-align="center">
          <template v-slot="scope">
            {{ scope.$index + 1 }}
          </template>
        </el-table-column>
        <el-table-column label="IP地址" prop="ipaddr" align="center" header-align="center" width="140px">
        </el-table-column>
        <el-table-column label="IP池" prop="import_in_pool" align="center" header-align="center" width="50px">
        </el-table-column>
        <el-table-column label="JOB数" prop="import_in_job_num" align="center" header-align="center" width="60px">
        </el-table-column>
        <el-table-column label="错误信息" prop="import_error" align="center" header-align="center" show-overflow-tooltip>
        </el-table-column>
        <el-table-column label="操作" align="center" header-align="center" width="80px">
          <template v-slot="scope">
            <el-popover :visible="deleteVisible[scope.$index]" placement="top" :width="160">
              <p>确定删除吗？</p>
              <div style="text-align: right; margin: 0">
                <el-button size="mini" type="text" @click="doNo(scope)">取消</el-button>
                <el-button type="primary" size="mini" @click="doYes(scope)">确定</el-button>
              </div>
              <template #reference>
                <el-button size="mini" type="danger" plain @click="doDelete(scope)">删除</el-button>
              </template>
            </el-popover>
          </template>
        </el-table-column>
      </el-table>
      <div class="block">
        <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange" :current-page="currentPage"
          :page-sizes="[17, 30, 50, 100]" :page-size="pageSize" layout="total, sizes, prev, pager, next, jumper"
          :total="pageTotal">
        </el-pagination>
      </div>
    </div>
    <el-backtop>
      <div style="
           {
            height: 100%;
            width: 100%;
            background-color: #f2f5f6;
            box-shadow: 0 0 6px rgba(0, 0, 0, 0.12);
            text-align: center;
            line-height: 40px;
            color: #1989fa;
          }
        ">
        UP
      </div>
    </el-backtop>
  </div>
</template>
<script>
import XLSX from "xlsx";
import { getJobs } from '@/api/jobs.js'
import { uploadMachines } from '@/api/upload.js'
export default {
  data () {
    return {
      fileList: [],
      uploadIPs: [],
      uploadIPsSplit: [],
      deleteVisible: {},
      currentPage: 1,
      pageSize: 17,
      pageTotal: 0,
      searchContent: '',
      multipleSelection: [],
      pushOption: {
        ignoreErr: false
      },
      // titles: { ignoreErrStatus: '未选中' },
      // ignoreErrStatus: '未选中',
      exampleData: [
        '标题，内容不限',
        '192.168.100.1',
        '192.168.100.2',
        '192.168.100.3',
        '192.168.100.4',
        '192.168.100.5',
        '192.168.100.6',
        '192.168.100.7',
        '192.168.100.8',
        '......',
        ''
      ],
      jobs: [],
      selectTypeValue: [],
      ipTongJi: {
        total: 0,
        success: 0,
        fail: 0,
        noaction: 0
      },
      pushing: false
    }
  },
  mounted () {
    this.doGetJobs()
  },
  // computed: 中的key，不能在data里面定义
  // watched: 中的key,在data是有定义的
  computed: {
    ignoreErrStatus () {
      if (this.pushOption.ignoreErr === true) {
        return '继续导入'
      } else {
        return '停止退出'
      }
    }
  },
  // watch: {
  //   pushOption: {
  //     handler () {
  //       if (this.pushOption.ignoreErr === true) {
  //         // this.titles.ignoreErrStatus = '选中'
  //         this.ignoreErrStatus = '选中'
  //       } else {
  //         // this.titles.ignoreErrStatus = '未选中'
  //         this.ignoreErrStatus = '未选中'
  //       }
  //     },
  //     deep: true
  //   }
  // },
  methods: {
    importExcel (file) {
      const types = file.name.split(".")[1];
      const fileType = ["xlsx", "xlc", "xlm", "xls", "xlt", "cvs"].some(
        item => item === types
      );
      if (!fileType) {
        this.$notify({
          title: '警告',
          message: '文件类型不对！',
          type: 'warning',
          duration: 1000,
        })
        return false
      }
      this.file2Xce(file).then(tabJson => {
        if (tabJson && tabJson.length > 0) {
          let sets = new Set()
          let uploadIPs = []
          tabJson[0].sheet.forEach(ipField => {
            for (let key in ipField) {
              sets.add(ipField[key])
            }
          })
          let deleteVisible = {}
          let n = 0
          sets.forEach(each => {
            uploadIPs.push({
              id: n,
              ipaddr: each,
              import_in_pool: false,
              import_in_job_num: 0,
              import_error: ''
            })
            deleteVisible[n] = false
            n += 1
          })
          this.uploadIPs = uploadIPs
          this.deleteVisible = deleteVisible
          this.uploadIPsSplit = uploadIPs.slice(0, this.pageSize)
          this.pageTotal = uploadIPs.length
          this.ipTongJi.total = this.uploadIPs.length
          this.doGetJobs()
        }
      });
    },
    doGetJobs () {
      getJobs().then(r => {
        this.jobs = r.data
      }).catch(e => console.log(e))
    },
    file2Xce (file) {
      return new Promise(function (resolve, reject) {
        const reader = new FileReader();
        reader.onload = function (e) {
          const data = e.target.result;
          this.wb = XLSX.read(data, {
            type: "binary"
          });
          const result = [];
          this.wb.SheetNames.forEach(sheetName => {
            result.push({
              sheetName: sheetName,
              sheet: XLSX.utils.sheet_to_json(this.wb.Sheets[sheetName])
            });
          });
          resolve(result);
        };
        reader.readAsBinaryString(file.raw);
      });
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
    handleSizeChange (val) {
      this.pageSize = val
      this.tableData()
    },
    handleCurrentChange (val) {
      this.currentPage = val
      this.tableData()
    },
    tableData () {
      let currPage = this.currentPage
      let pageSize = this.pageSize
      if (currPage === 0) {
        currPage = 0
      } else {
        currPage -= 1
      }
      this.ipTongJi.total = this.uploadIPs.length
      if (this.searchContent === '') {
        this.uploadIPsSplit = this.uploadIPs.slice(currPage * pageSize, currPage * pageSize + pageSize)
        this.pageTotal = this.uploadIPs.length
        return
      }
      let searchUploadIPs = this.uploadIPs.filter(x => x.ipaddr.indexOf(this.searchContent) > -1)
      this.pageTotal = searchUploadIPs.length
      this.uploadIPsSplit = searchUploadIPs.slice(currPage * pageSize, currPage * pageSize + pageSize)
    },
    filterIPMethod (query, ipaddr) {
      return ipaddr.filter(query) > -1;
    },
    doYes (scope) {
      const index = this.uploadIPs.findIndex(d => d.ipaddr === scope.row.ipaddr)
      this.uploadIPs.splice(index, 1);
      this.tableData()
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
      let searchUploadIPs = this.uploadIPs.filter(x => {
        const result = this.multipleSelection.some(
          id => id === x.id
        );
        return !result
      })
      this.uploadIPs = searchUploadIPs
      this.tableData()
    },
    submitUpload () {
      this.doUploadMachines()
    },
    onSearch () {
      this.tableData()
    },
    selectChange (event) { },
    goBack () {
      let queryInfo = {
        currentPage: 0,
        pageSize: 0
      }
      if (this.$route.params.currentPage) {
        queryInfo.currentPage = this.$route.params.currentPage
      }
      if (this.$route.params.pageSize) {
        queryInfo.pageSize = this.$route.params.pageSize
      }
      this.$router.push({ name: 'ipManager', params: queryInfo })
    },
    doUploadMachines () {
      if (this.uploadIPs.length === 0) {
        return false
      }
      this.pushing = true
      let uploadIPs = []
      this.uploadIPs.forEach(i => {
        if (i.import_in_pool) {
          return
        }
        uploadIPs.push(i)
      })
      if (uploadIPs.length === 0) {
        this.$notify({
          title: '警告',
          message: '列表中已经没有可提交的IP地址！',
          type: 'warning',
          duration: 1000,
        })
        return false
      }
      let uploadData = {
        opts: {
          ignore_err: this.pushOption.ignoreErr,
        },
        jobs_id: this.selectTypeValue,
        machines: uploadIPs
      }
      uploadMachines(uploadData).then(r => {
        let n = 0
        let deleteVisible = {}
        // let machines = []
        r.data.machines.forEach(each => {
          each.id = n
          deleteVisible[n] = false
          n += 1
        })
        this.uploadIPs = r.data.machines
        this.ipTongJi = r.data.tongji
        this.deleteVisible = deleteVisible
        this.tableData()
        this.$notify({
          title: '警告',
          message: 'IP列表已提交，请参考表格中显示的结果！',
          type: 'warning',
          duration: 1000,
        })
        this.pushing = false
      }).catch(e => {
        this.pushing = false
        console.log(e)
      })
    }
  }
}
</script>

<style scoped>
.batch-box {
  display: flex;
  justify-content: flex-start;
  flex-wrap: nowrap;
  min-height: 85vh;
}

.ctl_btn_area {
  display: flex;
  justify-content: space-between;
  flex-wrap: nowrap;
}

.batch-box-desc {
  width: 600px;
}

.batch-box-data {
  width: 100%;
}

.do_action {
  display: flex;
  justify-content: flex-end;
  flex-wrap: nowrap;
  margin-bottom: 5px;
}

.block {
  padding-top: 12px;
  text-align: center;
}

.return-back {
  margin-bottom: 27px;
}

.split-line {
  margin-left: 8px;
  margin-right: 8px;
  min-height: 85vh;
}

.push-options {
  margin-bottom: 15px;
}

.table-example {
  border: 1px;
}

body {
  margin: 10px;
}

table {
  border-collapse: collapse;
  border-spacing: 0;
}

td,
th {
  border: 1px solid #cbcbcb;
  display: inline-block;
  height: 20px;
}

.pure-table {
  font-size: 15px;
  border: 1px solid lightblue;
  align-content: center;
  text-align: center;
}

.pure-table-td1 {
  align-content: center;
  text-align: center;
  background-color: #f8fbef;
  width: 30px;
}

.pure-table-td2,
.pure-table-th2,
.pure-table-th3 {
  width: 130px;
}

.pure-table-th1,
.pure-table-th2,
.pure-table-th3 {
  display: inline-block;
  font-size: 15px;
  /* border: 1px solid lightblue; */
  align-content: center;
  text-align: center;
  background-color: #f8fbef;
}

.pure-table-th1 {
  width: 30px;
}
</style>