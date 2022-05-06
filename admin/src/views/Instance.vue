<template>
  <div class="instance_box">
    <div class="instance_action">
      <div class="instance_action_btn_left">
        <div><el-tag type="warning">功能说明：导入其它prometheus实例数据</el-tag></div>
        <el-form :inline="true" size="small" ref="instanceInfo" :model="instanceInfo" class="demo-form-inline-get">
          <el-form-item label="IP地址及端口号">
            <el-input v-model="instanceInfo.instance" placeholder="IP地址及端口号" />
          </el-form-item>
          <el-form-item>
            <el-button v-if="getDataStatus===false" type="primary" plain icon="el-icon-download" @click="onGetInstanceData('instanceInfo')">获取数据</el-button>
            <el-button v-if="getDataStatus===true" type="primary" plain icon="el-icon-loading">获取数据</el-button>
          </el-form-item>
        </el-form>
      </div>
      <div class="instance_action_btn_right">
        <el-form :inline="true" size="small" :model="instanceInfo" class="demo-form-inline-put">
          <el-form-item label="" prop="type">
            <el-radio-group v-model="import_type">
              <el-radio label="group_name_only">只导入组名</el-radio>
              <el-radio label="all_ip_only">只导入IP</el-radio>
              <el-radio label="merge_group_ip">合并组(会导入IP)</el-radio>
              <el-radio label="replace_group_ip">替换组(会导入IP)</el-radio>
              <!-- <el-radio label="user_defined">自定义</el-radio> -->
            </el-radio-group>
          </el-form-item>
          <el-form-item>
            <el-button v-if="putDataStatus===false" type="primary" plain icon="el-icon-upload2" @click="onSubmit">导入数据</el-button>
            <el-button v-if="putDataStatus===true" type="primary" plain icon="el-icon-loading">导入数据</el-button>
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
              :data="currPageTargetsJobsData"
              stripe
              :row-style="rowStyle"
              :cell-style="cellStyle"
              header-align="center"
              @selection-change="handleSelectionChange"
              ref="multipleTableRef"
              :row-key="getRowKeys"
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
              <el-table-column type="expand">
                <template #default="props">
                  <el-descriptions title="IP列表" size="mini" :column="3" border>
                    <el-descriptions-item
                      v-for="(addr, i) in props.row.addrs"
                      :key="i"
                      :label="i+1"
                      >{{ addr }}
                    </el-descriptions-item>
                  </el-descriptions>
                </template>
              </el-table-column>
              <el-table-column
                label="JOB组名"
                prop="name"
                header-align="center"
                show-overflow-tooltip
              >
              </el-table-column>
              <el-table-column
                label="端口号"
                prop="port"
                align="center"
                width="120px"
                header-align="center"
              >
              </el-table-column>
              <el-table-column
                label="IP数"
                prop="addr_count"
                align="center"
                header-align="center"
                width="80px"
              >
              </el-table-column>
              <el-table-column
                label="up"
                prop="up"
                align="center"
                header-align="center"
                width="60px"
              >
              </el-table-column>
              <el-table-column
                label="down"
                prop="down"
                align="center"
                header-align="center"
                show-overflow-tooltip
                width="60px"
              >
              </el-table-column>
              <el-table-column type="selection" :reserve-selection="true" label="选中提交" width="140"></el-table-column>
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
            <el-scrollbar class="json-scrollbar">
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
        instance: ''
      },
      rules: {
        addr: [
          { required: true, message: '请输入正确IP及端口号', trigger: ['blur'] }
        ],
      },
      import_type: "merge_group_ip",
      targetsJobsData: [],
      currPageTargetsJobsData: [],
      instanceRespData: {"result": "没有数据"},
      pageSize: 20,
      currentPage: 1,
      pageTotal: 0,
      searchContent: '',
      uploadIPsSplit: [],
      multipleSelection: [],
      getDataStatus: false,
      putDataStatus: false,
    }
  },
  created () {
  },
  mounted () {
  },
  computed() {

  },
  watch: {
    searchContent: function (newVal, oldVal) {
      this.searchData(newVal)
    }
  },
  methods: {
    onGetInstanceData(formName){
      this.getDataStatus = true
      this.$refs[formName].validate((valid) => {
        if (valid) {
          getInstanceTargets(this.instanceInfo).then(r => {
            this.targetsJobsData = r.data.jobs
            this.instanceRespData = r.data.data
            this.currPageTargetsJobsData = this.targetsJobsData.slice(0, this.pageSize)
            this.pageTotal = this.targetsJobsData.length
            this.getDataStatus = false
          }).catch(e=>{this.getDataStatus = false})
        } else {
          this.getDataStatus = false
        }
      })
    },
    onSubmit(){
      console.log(this.multipleSelection)
      if (this.multipleSelection.length === undefined || this.multipleSelection.length === 0) {
        return
      }
      this.putDataStatus = true
      putInstanceTargets({type: this.import_type, data: this.multipleSelection}).then(r=>{
        this.$notify({
          title: '成功',
          message: '更新成功！',
          type: 'success'
        });
        this.putDataStatus = false
      }).catch(e=>{console.log(e); this.putDataStatus = false})
    },
    onSearch(){
      this.searchData(this.searchContent)
    },
    searchData(newVal){
      let currPage = this.currentPage
      let pageSize = this.pageSize
      if (currPage === 0) {
        currPage = 0
      } else {
        currPage -= 1
      }
      if (newVal === '') {
        this.pageTotal = this.targetsJobsData.length
        this.currPageTargetsJobsData = this.targetsJobsData.slice(currPage * pageSize, currPage * pageSize + pageSize)
      } else {
        let filterData = this.targetsJobsData.filter(x => {
          if (x.name.indexOf(newVal) > -1) {
            return true
          } 
          if (x.port.indexOf(newVal) > -1) {
            return true
          }
          for (var k = 0, length = x.addrs.length; k < length; k++) {
            if(x.addrs[k].indexOf(newVal) > -1) {
              return true
            }
          }
          return false
        })
        this.pageTotal = filterData.length
        this.currPageTargetsJobsData = filterData.slice(currPage * pageSize, currPage * pageSize + pageSize)
      }
    },
    handleSizeChange (val) {
      this.pageSize = val
      this.searchData(this.searchContent)
    },
    handleCurrentChange (val) {
      this.currentPage = val
      this.searchData(this.searchContent)
    },
    handleSelectionChange (selection) {
      this.multipleSelection = selection
    },
    getRowKeys(row){
      return row.name
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

.instance_action {
  display: flex;
  justify-content: space-between;
}
.instance_action_btn {
  display: flex;
  justify-content: space-around;
  flex-wrap: nowrap;
  flex-direction: row;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.box-card-right :deep() .el-card__header {
  padding-top: 16px;
  padding-bottom: 17px;
}

.box-card-left {
  width: 55%;
}

.instance_action_btn_left {
  display: flex;
  justify-content: space-between;
  width: 55%;
}

.box-card-left :deep() .el-card__header {
  padding-top: 13px;
  padding-bottom: 12px;
}
.box-card-right {
  width: 45%;
}

.instance_action_btn_right {
  width: 45%;
  display: flex;
  justify-content: flex-end;
}

:deep() .el-card__body {
  padding: 0px 0px 0px 0px;
}

.block {
  padding-top: 12px;
  text-align: center;
}

.json-scrollbar {
  margin-left: 10px;
  height: 770px;
}

.card-scrollbar {
  height: 600px;
}

</style>
