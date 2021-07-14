<template>
  <div class="batch-box">
    <div class="batch-box-desc">
      <el-tag style="primary">导入格式说明：</el-tag>
      <el-descriptions title="自定义样式列表" :column="1" border>
        <el-descriptions-item
          label="用户名"
          label-align="right"
          align="center"
          label-class-name="my-label"
          class-name="my-content"
          width="150px"
          >kooriookami</el-descriptions-item
        >
        <el-descriptions-item label="手机号" label-align="right" align="center"
          >18100000000</el-descriptions-item
        >
        <el-descriptions-item label="居住地" label-align="right" align="center"
          >苏州市</el-descriptions-item
        >
        <el-descriptions-item label="备注" label-align="right" align="center">
          <el-tag size="small">学校</el-tag>
        </el-descriptions-item>
        <el-descriptions-item
          label="联系地址"
          label-align="right"
          align="center"
          >江苏省苏州市吴中区吴中大道 1188 号</el-descriptions-item
        >
      </el-descriptions>
    </div>
    <div class="batch-box-data">
      <div class="ctl_btn_area">
        <div>
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
            <el-upload
              class="upload-demo"
              ref="upload"
              action=""
              :auto-upload="false"
              accept=".csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet, application/vnd.ms-excel"
              :file-list="fileList"
              :show-file-list="false"
              :on-change="importExcel"
              type="file"
            >
              <template #trigger>
                <el-button size="small" type="primary">选取文件</el-button>
              </template>
              <el-button
                style="margin-left: 10px"
                size="small"
                type="success"
                @click="submitUpload"
                >上传到服务器</el-button
              >
              <template #tip>
                <div class="el-upload__tip">文件格式：xlxs, xls, cvs</div>
              </template>
            </el-upload>
            <!-- <el-button size="small" type="success" plain>提交</el-button> -->
          </div>
          <div>
            <el-input
              size="small"
              placeholder="请输入内容"
              v-model="searchContent"
            >
              <template #append>搜索</template>
            </el-input>
          </div>
        </div>
      </div>
      <el-table
        size="mini"
        highlight-current-row
        border
        :data="uploadIPsSplit"
        stripe
        :row-style="rowStyle"
        :cell-style="cellStyle"
        header-align="center"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="40"> </el-table-column>
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
          label="操作"
          align="center"
          header-align="center"
          width="180px"
        >
          <template v-slot="scope">
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
      <div class="block">
        <el-pagination
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
          :current-page="currentPage"
          :page-sizes="[15, 20, 30, 50]"
          :page-size="pageSize"
          layout="total, sizes, prev, pager, next, jumper"
          :total="pageTotal"
        >
        </el-pagination>
      </div>
    </div>
  </div>
</template>
<script>
import XLSX from "xlsx";
export default {
  data () {
    return {
      fileList: [],
      uploadIPs: [],
      uploadIPsSplit: [],
      deleteVisible: {},
      currentPage: 1,
      pageSize: 15,
      pageTotal: 0,
      searchContent: '',
      multipleSelection: []
    }
  },
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
          type: 'warning'
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
              ipaddr: each
            })
            deleteVisible[n] = false
          })
          this.uploadIPs = uploadIPs
          this.deleteVisible = deleteVisible
          this.uploadIPsSplit = uploadIPs.slice(0, this.pageSize)
          this.pageTotal = uploadIPs.length
        }
      });
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
      this.tableData(this.currentPage, this.pageSize)
    },
    handleCurrentChange (val) {
      this.currentPage = val
      this.tableData(this.currentPage, this.pageSize)
    },
    tableData (currPage, pageSize) {
      if (currPage === 0) {
        currPage = 0
      } else {
        currPage -= 1
      }
      this.uploadIPsSplit = this.uploadIPs.slice(currPage * pageSize, currPage * pageSize + pageSize)
      this.pageTotal = this.uploadIPs.length
    },
    filterIPMethod (query, ipaddr) {
      return ipaddr.indexOf(query) > -1;
    },
    doYes (scope) {
      const index = this.uploadIPs.findIndex(d => d.ipaddr === scope.row.ipaddr)
      this.uploadIPs.splice(index, 1);
      this.tableData(this.currentPage, this.pageSize)
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
    doBatchDel () { },
    submitUpload () {
      this.$refs.upload.submit();
    },
    onSearch () { },
    handleRemove (file, fileList) {
      console.log(file, fileList);
    },
    handlePreview (file) {
      console.log(file);
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
  text-align: center;
}
.do_action {
  display: flex;
  justify-content: flex-end;
  flex-wrap: nowrap;
  margin-bottom: 5px;
}
.block {
  padding-top: 12px;
}
</style>