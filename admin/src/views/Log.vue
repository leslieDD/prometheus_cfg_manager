<template>
  <div>
    <el-descriptions class="margin-top" :column="3" size="small" border>
      <el-descriptions-item>
        <template #label>
          <i class="el-icon-user"></i>
          要记录的日志类型
        </template>
        <el-select size="small" v-model="levelPool" clearable placeholder="请选择" multiple collapse-tags
          @visible-change="doPutRecodeSetting" @clear="doPutRecodeSetting" @change="selectChange">
          <el-option v-for="item in allLevelPool" :key="item.id" :label="item.label" :value="item.id">
          </el-option>
        </el-select>
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label>
          <i class="el-icon-user"></i>
          请空所有日志
        </template>
        <el-button size="small" :disabled="runningMod" type="danger" @click="doClearSystem">请空所有日志</el-button>
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label>
          <i class="el-icon-mobile-phone"></i>
          搜索
        </template>
        <el-input placeholder="请输入查询内容，回车执行搜索" size="small" @keyup.enter="doSearch" v-model="searchContent"></el-input>
      </el-descriptions-item>
    </el-descriptions>
    <!-- <el-scrollbar height="78vh" class="flex-content"> -->
    <el-table size="mini" highlight-current-row border :data="systemLogs" stripe :row-style="rowStyle"
      :cell-style="cellStyle" height="78vh">
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
      <el-table-column label="操作内容" show-overflow-tooltip width="230px" prop="operate_content">
      </el-table-column>
      <el-table-column label="结果" width="50px" prop="operate_result">
      </el-table-column>
      <el-table-column label="操作时间" width="150px" prop="operate_at">
        <template v-slot="{ row }">
          <span>{{ parseTimeSelf(row.operate_at) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="错误原因" show-overflow-tooltip prop="operate_error">
      </el-table-column>
    </el-table>
    <!-- </el-scrollbar> -->
    <div class="block" v-if="paginationShow">
      <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange" :current-page="currentPage"
        :page-sizes="[25, 50, 80, 100]" :page-size="pageSize" layout="total, sizes, prev, pager, next, jumper"
        :total="pageTotal">
      </el-pagination>
    </div>
  </div>
</template>

<script>
import {
  getSystemReocdeSetting,
  putSystemReocdeSetting,
  getSystemLog,
  clearSystemLog
} from '@/api/log.js'
export default {
  name: "Log",
  data () {
    return {
      systemLogs: [],
      paginationShow: true,
      pageSize: 25,
      pageTotal: 0,
      currentPage: 1,
      searchContent: '',
      runningMod: false,
      levelPool: [],
      allLevelPool: [],
      selectChanged: false
    }
  },
  mounted () {
    this.dogetSystemReocdeSetting()
    this.doGetSystemLog()
  },
  methods: {
    doSearch () {
      this.doGetSystemLog()
    },
    doGetSystemLog (getInfo) {
      if (!getInfo) {
        getInfo = {
          'pageNo': this.currentPage,
          'pageSize': this.pageSize,
          'search': this.searchContent
        }
      }
      getSystemLog(getInfo).then(r => {
        this.systemLogs = r.data.data
        this.pageTotal = r.data.totalCount
        this.currentPage = r.data.pageNo
        this.pageSize = r.data.pageSize
        this.paginationShow = false;//让分页隐藏
        this.$nextTick(() => {//重新渲染分页
          this.paginationShow = true;
        });
      }).catch(e => console.log(e))
    },
    doClearSystem () {
      this.runningMod = true
      clearSystemLog().then(r => {
        this.$notify({
          title: '成功',
          message: '清理成功！',
          type: 'success',
          duration: 1000,
        });
        this.runningMod = false
        this.doGetSystemLog()
      }).catch(e => {
        console.log(e)
        this.runningMod = false
      })
    },
    dogetSystemReocdeSetting () {
      getSystemReocdeSetting().then(r => {
        let levelPool = []
        r.data.forEach(x => {
          if (x.selected) {
            levelPool.push(x.id)
          }
        })
        this.levelPool = levelPool
        this.allLevelPool = r.data
      })
    },
    doPutRecodeSetting (action) {
      if (!this.selectChanged) {
        return
      }
      if (!action) {
        putSystemReocdeSetting(this.levelPool).then(r => {
          this.selectChanged = false
          this.$notify({
            title: '成功',
            message: '更新成功！',
            type: 'success',
            duration: 1000,
          });
        })
      }
    },
    selectChange () {
      this.selectChanged = true
    },
    handleSizeChange (val) {
      let getInfo = {
        'pageNo': this.currentPage,
        'pageSize': val,
        'search': this.searchContent
      }
      this.doGetSystemLog(getInfo)
    },
    handleCurrentChange (val) {
      let getInfo = {
        'pageNo': val,
        'pageSize': this.pageSize,
        'search': this.searchContent
      }
      this.doGetSystemLog(getInfo)
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
  padding-top: 6px;
}

.system-area {
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