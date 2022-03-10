<template>
  <div>
    <el-descriptions class="margin-top" :column="2" size="small" border>
      <el-descriptions-item>
        <template #label>
          <i class="el-icon-user"></i>
          操作
        </template>
        <el-button
          size="small"
          :disabled="runningMod"
          type="danger"
          icon="el-icon-warning-outline"
          @click="displayDialog"
          >重置所有数据</el-button
        >
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label>
          <i class="el-icon-mobile-phone"></i>
          搜索
        </template>
        <el-input
          placeholder="请输入查询内容，回车执行搜索"
          size="small"
          @keyup.enter="doSearch"
          v-model="searchContent"
        ></el-input>
      </el-descriptions-item>
    </el-descriptions>
    <el-scrollbar height="75vh" class="flex-content">
      <el-table
        size="mini"
        highlight-current-row
        border
        :data="operateLogs"
        stripe
        :row-style="rowStyle"
        :cell-style="cellStyle"
      >
        <el-table-column label="序号" width="50px">
          <template v-slot="scope">
            {{ scope.$index + 1 }}
          </template>
        </el-table-column>
        <el-table-column label="用户名" width="100px" prop="username">
        </el-table-column>
        <el-table-column label="IP地址" width="150px" prop="ipaddr">
        </el-table-column>
        <el-table-column label="类型" width="80px" prop="operate_type">
        </el-table-column>
        <el-table-column
          label="操作内容"
          show-overflow-tooltip
          width="230px"
          prop="operate_content"
        >
        </el-table-column>
        <el-table-column label="结果" width="50px" prop="operate_result">
        </el-table-column>
        <el-table-column label="操作时间" width="130px" prop="operate_at">
          <template v-slot="{ row }">
            <span>{{ parseTimeSelf(row.operate_at) }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="错误原因"
          show-overflow-tooltip
          prop="operate_error"
        >
        </el-table-column>
      </el-table>
      <div class="block" v-if="paginationShow">
        <el-pagination
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
          :current-page="currentPage"
          :page-sizes="[30, 50, 80, 100]"
          :page-size="pageSize"
          layout="total, sizes, prev, pager, next, jumper"
          :total="pageTotal"
        >
        </el-pagination>
      </div>
    </el-scrollbar>
    <div>
      <el-dialog title="输入重置码" v-model="dialogFormVisible" width="430px">
        <div class="reset-explain">
          <p>系統会在后台生成一个重置码，可以在服务的日志中看到，如：</p>
          <p style="font-width: bold">
            reset <span style="color: blue">prometheus</span> config key:
            <span style="color: red">6eb96669-2fb1-4758-8776-ddcdb58405c9</span>
          </p>
          <p>输入红色部分的KEY即可重置成出产设置</p>
        </div>
        <el-form
          :model="ruleForm"
          :rules="rules"
          ref="ruleForm"
          label-width="auto"
          class="demo-ruleForm"
          size="mini"
        >
          <el-form-item label="重置码" prop="code">
            <el-input v-model="ruleForm.code"></el-input>
          </el-form-item>
          <el-form-item>
            <div class="dialog-footer">
              <el-button size="small" @click="closeDialog('ruleForm')"
                >关 闭</el-button
              >
              <el-button
                v-if="runningMod === true"
                size="small"
                type="primary"
                icon="el-icon-loading"
                >确 定</el-button
              >
              <el-button
                v-if="runningMod === false"
                size="small"
                type="primary"
                icon="el-icon-refresh"
                @click="doResetSystem('ruleForm')"
                >确 定</el-button
              >
            </div>
          </el-form-item>
        </el-form>
      </el-dialog>
    </div>
  </div>
</template>

<script>
import {
  resetSystem,
  preResetSystem,
  getOperateLog
} from '@/api/empty.js'
export default {
  name: "Empty",
  data () {
    return {
      operateLogs: [],
      paginationShow: true,
      pageSize: 30,
      pageTotal: 0,
      currentPage: 1,
      searchContent: '',
      dialogFormVisible: false,
      runningMod: false,
      ruleForm: {
        code: ''
      },
      rules: {
        code: [
          { required: true, message: '请输入重置码！', trigger: 'blur' }
        ],
      }
    }
  },
  mounted () {
    this.doGetOperateLog()
  },
  methods: {
    doSearch () {
      this.doGetOperateLog()
    },
    doGetOperateLog (getInfo) {
      if (!getInfo) {
        getInfo = {
          'pageNo': this.currentPage,
          'pageSize': this.pageSize,
          'search': this.searchContent
        }
      }
      getOperateLog(getInfo).then(r => {
        this.operateLogs = r.data.data
        this.pageTotal = r.data.totalCount
        this.currentPage = r.data.pageNo
        this.pageSize = r.data.pageSize
        this.paginationShow = false;//让分页隐藏
        this.$nextTick(() => {//重新渲染分页
          this.paginationShow = true;
        });
      }).catch(e => console.log(e))
    },
    displayDialog () {
      preResetSystem().then(r => {
        this.$notify({
          title: '警告',
          message: '已经生成新的重置码！',
          type: 'warning'
        })
        this.dialogFormVisible = true
      })
    },
    closeDialog (formName) {
      this.dialogFormVisible = false
      this.$refs[formName].resetFields()
    },
    doResetSystem (formName) {
      this.runningMod = true
      this.$refs[formName].validate((valid) => {
        if (valid) {
          resetSystem(this.ruleForm).then(r => {
            this.$notify({
              title: '成功',
              message: '重置成功！',
              type: 'success'
            });
            this.runningMod = false
            this.doGetOperateLog()
          }).catch(e => {
            console.log(e)
            this.runningMod = false
            this.doGetOperateLog()
          })
        } else {
          this.runningMod = false
          console.log(`validate: ${valid}`)
        }
      })
    },
    handleSizeChange (val) {
      let getInfo = {
        'pageNo': this.currentPage,
        'pageSize': val,
        'search': this.searchContent
      }
      this.doGetOperateLog(getInfo)
    },
    handleCurrentChange (val) {
      let getInfo = {
        'pageNo': val,
        'pageSize': this.pageSize,
        'search': this.searchContent
      }
      this.doGetOperateLog(getInfo)
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
    }
  }
}
</script>

<style scoped>
.el-pagination {
  text-align: center;
}
.block {
  padding-top: 12px;
}
.operate-area {
  display: flex;
  flex-direction: row;
  justify-content: flex-end;
}
.dialog-footer {
  text-align: right;
  align-content: right;
}
.reset-explain {
  font: 0.8em Arial, Tahoma, Verdana;
  color: #777;
  margin-top: -30px;
  margin-bottom: 12px;
}
</style>