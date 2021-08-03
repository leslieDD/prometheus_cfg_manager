<template>
  <div class="priv-board">
    <el-scrollbar>
      <el-table
        :data="privTableData"
        :span-method="objectSpanMethod"
        border
        style="width: 100%"
        :row-style="rowStyle"
        :cell-style="cellStyle"
      >
        <el-table-column prop="page_nice_name" label="页面" width="120px">
        </el-table-column>
        <el-table-column prop="sub_page_nice_name" label="子页面" width="120px">
        </el-table-column>
        <el-table-column label="权限选项">
          <template #default="{ row }">
            <el-checkbox
              v-for="item in row.func_list"
              :key="item.id"
              v-model="item.checked"
              >{{ item.func_nice_name }}</el-checkbox
            >
          </template>
        </el-table-column>
      </el-table>
    </el-scrollbar>
  </div>
</template>

<script>

import { getPriv } from '@/api/priv.js'

export default {
  data () {
    return {
      privTableData: [],
      tableRowSpan: {},
      groupInfo: {}
    };
  },
  mounted () {
    this.current = ''
    this.getPriv()
  },
  methods: {
    objectSpanMethod ({ row, column, rowIndex, columnIndex }) {
      if (columnIndex === 0) {
        if (this.current !== row.page_name) {
          this.current = row.page_name
          const dataResp = {
            rowspan: this.tableRowSpan[row.page_name],
            colspan: 1
          }
          return dataResp
        } else {
          return {
            rowspan: 0,
            colspan: 0
          };
        }
      }
    },

    getPriv () {
      const groupInfo = {
        id: 1
      }
      getPriv(groupInfo).then(r => {
        let tableRowSpan = {}
        this.tableRowSpanDone = {}
        r.data.forEach(item => {
          if (tableRowSpan[item.page_name]) {
            tableRowSpan[item.page_name] += 1
          } else {
            tableRowSpan[item.page_name] = 1
          }
        })
        this.tableRowSpan = tableRowSpan
        this.privTableData = r.data
      }).catch(e => console.log(e))
    },
    rowStyle (row) {
      let rs = {
        'padding': '1px'
      }
      return rs
    },
    cellStyle (column) {
      let cs = {
        'padding': '1px'
      }
      return cs
    },
  }
};
</script>

<style scoped>
.priv-board {
  height: 100%;
}
</style>