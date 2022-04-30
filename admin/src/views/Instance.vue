<template>
  <div class="instance_box">
    <div class="instance_action">
      <div><el-tag type="warning">功能说明：导入其它prometheus实例数据</el-tag></div>
      <div class="instance_action_btn">
        <el-form :inline="true" size="small" ref="instanceInfo" :model="instanceInfo" class="demo-form-inline-get">
          <el-form-item label="IP地址及端口号">
            <el-input v-model="instanceInfo.addr" placeholder="IP地址及端口号" />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" icon="el-icon-download" @click="onGetInstanceData('instanceInfo')">获取数据</el-button>
          </el-form-item>
        </el-form>
        <el-form :inline="true" size="small" :model="instanceInfo" class="demo-form-inline-put">
          <el-form-item label="" prop="type">
            <el-radio-group v-model="import_type">
              <el-radio label="group_name_only">只导入组名</el-radio>
              <el-radio label="all_ip_only">只导入IP</el-radio>
              <el-radio label="merge_group_ip">合并组</el-radio>
              <el-radio label="replace_group_ip">替换组</el-radio>
              <el-radio label="user_defined">自定义</el-radio>
            </el-radio-group>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" icon="el-icon-upload2" @click="onSubmit">导入数据</el-button>
          </el-form-item>
        </el-form>
      </div>
    </div>
    <div class="instance_content">
      <el-card class="box-card-left" shadow="never">
        <template #header>
          <div class="card-header">
            <span>统计信息</span>
            <span>
              <!-- <span style="margin-right: 10px"><el-checkbox v-model="search_ip" label="同时搜索线路地址池" size="small" /></span> -->
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
        </template>
        <div height="100%" class="box-card-left-body">
          <el-scrollbar class="card-scrollbar">
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
                label="JOB组名"
                prop="import_in_job_num"
                align="center"
                header-align="center"
                show-overflow-tooltip
              >
              </el-table-column>
              <el-table-column
                label="JOB组IP数"
                prop="import_error"
                align="center"
                header-align="center"
                show-overflow-tooltip
                width="100px"
              >
              </el-table-column>
              <el-table-column type="selection" label="选中提交" width="140"></el-table-column>
            </el-table>
          </el-scrollbar>
            <div class="block">
              <el-pagination
                @size-change="handleSizeChange"
                @current-change="handleCurrentChange"
                :current-page="currentPage"
                :page-sizes="[20, 30, 50, 100]"
                :page-size="pageSize"
                layout="total, sizes, prev, pager, next, jumper"
                :total="pageTotal"
              >
              </el-pagination>
            </div>
        </div>
      </el-card>
      <el-card class="box-card-right" shadow="never">
        <template #header>
          <div class="card-header">
            <span>
              实例返回内容
            </span>
          </div>
        </template>
        <div class="card-body">
            <el-scrollbar>
              <JsonView :json="instanceRespData"></JsonView>
            </el-scrollbar>
        </div>
      </el-card>
    </div>
  </div>
</template>

<script>

import  {getInstanceTargets, putInstanceTargets} from '@/api/instance.js'

export default {
  name: 'Instance',
  data () {
    return {
      instanceInfo: {
        addr: ''
      },
      rules: {
        addr: [
          { required: true, message: '请输入正确IP及端口号', trigger: ['blur'] }
        ],
      },
      import_type: "merge_group_ip",
      targetsJobsData: [],
      instanceRespData: {"result": "没有数据"},
      pageSize: 20,
      currentPage: 1,
      pageTotal: 0,
    }
  },
  created () {
  },
  mounted () {
  },
  methods: {
    onGetInstanceData(formName){
      this.$refs[formName].validate((valid) => {
        if (valid) {
          getInstanceTargets(this.instanceInfo).then(r => {
            this.targetsJobsData = r.data.jobs
            this.instanceRespData = r.data.data
          })
        }
      })
    },
    onSearch(){},
    handleSizeChange (val) {
      this.pageSize = val
      this.tableData()
    },
    handleCurrentChange (val) {
      this.currentPage = val
      this.tableData()
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
.instance_box {
  height: 770px;
}
.instance_content {
  display: flex;
  justify-content: space-between;
}
.demo-form-inline-get {
  margin-right: 20px;
}
.instance_action {
  display: flex;
  justify-content: space-between;
}
.instance_action_btn {
  display: flex;
  justify-content: flex-end;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.box-card-left {
  width: 55%;
}

.box-card-left :deep() .el-card__header {
  padding-top: 13px;
  padding-bottom: 12px;
}
.box-card-right {
  width: 45%;
}

:deep() .el-card__body {
  padding: 0px 0px 0px 0px;
}

.block {
  padding-top: 12px;
  text-align: center;
}

</style>
