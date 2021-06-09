<template>
  <dir>
    <el-table :data="machines" style="width: 100%">
    <el-table-column label="序号" width="48px">
      <template slot-scope="scope">
          {{scope.$index+1}}
      </template>
    </el-table-column>
    <el-table-column label="IP地址" prop="ipaddr">
    </el-table-column>
    <el-table-column label="分组选项" prop="job_id">
      <template slot-scope="scope">
        <el-select v-model="selectTypeValue[scope.row.ipaddr]" @change="handleSelect(selectTypeValue[scope.row.ipaddr])" size="small" multiple placeholder="请选择">
          <el-option
            v-for="item in jobs"
            :key="item.id"
            :label="item.name"
            :value="item.id">
          </el-option>
        </el-select>
      </template>
    </el-table-column>
    <el-table-column align="right">
      <template slot-scope="scope">
        <el-button size="mini" type="danger" @click="handleDelete(scope.$index, scope.row)">删除</el-button>
      </template>
    </el-table-column>
  </el-table>
  </dir>
</template>

<script>
import { getJobs } from '@/api/jobs'
import { getMachines } from '@/api/machines'

export default {
  name: 'Manager',
  data () {
    return {
      machines: [],
      jobs: [],
      'selectTypeValue': [],
      'search': ''
    }
  },
  created () {
    this.doGetJobs()
    this.doGetMachines()
  },
  methods: {
    doGetList () {

    },
    doGetJobs () {
      getJobs().then(
        r => {
          this.jobs = r.data
        }
      ).catch(
        e => {
          console.log(e)
        }
      )
    },
    doGetMachines () {
      getMachines().then(
        r => {
          this.machines = r.data.data
          r.data.data.forEach(element => {
            this.selectTypeValue[element.ipaddr] = element.job_id
            // this.selectTypeValue = element.job_id
          })
        }
      ).catch(
        e => {
          console.log(e)
        }
      )
    },
    handleDelete (index, row) {},
    handleSelect (dataId, menuId) {
      console.log(dataId, menuId)
      // console.log(this.selectTypeValue[row.ipaddr])
      this.$set(this.selectTypeValue, dataId, menuId)
      // console.log(this.selectTypeValue)
    }
  }
}
</script>

<style scoped>
</style>
