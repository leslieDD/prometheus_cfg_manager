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
      <el-table-column label="规则" prop="rule" align="center" header-align="center" show-overflow-tooltip>
      </el-table-column>
      <el-table-column label="接口" prop="prometheus" width="300px" align="center" header-align="center"
        show-overflow-tooltip>
      </el-table-column>
      <el-table-column label="执行" prop="run_times" align="center" header-align="center" width="100px">
      </el-table-column>
      <el-table-column label="成功" prop="success_times" align="center" header-align="center" width="100px">
      </el-table-column>
      <el-table-column label="失败" prop="fail_times" align="center" header-align="center" width="100px">
      </el-table-column>
      <el-table-column label="更新账号" prop="update_by" width="80px" align="center" header-align="center"
        show-overflow-tooltip></el-table-column>
      <el-table-column label="更新时间" prop="update_at" width="140px" align="center" header-align="center">
        <template v-slot="{ row }">
          <span>{{ parseTimeSelf(row.update_at) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" width="300px">
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
  </div>
</template>

<script>


export default {
  name: 'Crontab',
  data () {
    return {
      pageSize: 20,
      pageTotal: 0,
      currentPage: 1,
      pageshow: true,
      multipleSelection: [],
      crontabRules: [],
    }
  },
  created () {
  },
  mounted () {
  },
  methods: {
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
    doAdd () {
      getAllPrometheusAPI().then(r => {
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
    onSearch () {

    },
    doEdit () {

    },
    doNo () {

    },
    doYes () {

    },
    doDelete () {

    },
    invocate () {

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
  }
}
</script>

<style scoped>
.crontab-board {
  text-align: center;
}

.do_action {
  display: flex;
  flex-wrap: nowrap;
  justify-content: space-between;
  margin-top: -5px;
  padding-bottom: 8px;
}

.block {
  padding-top: 12px;
}
</style>
