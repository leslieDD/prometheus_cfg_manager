<template>
  <div class="priv-board">
    <el-scrollbar>
      <el-table
        size="mini"
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
          <template #default="{ row }">
            <div class="sub-page-class">
              <span>{{ row.sub_page_nice_name + " " }}</span
              ><el-checkbox
                v-model="
                  checkBoxSelect[`${row.page_name}_${row.sub_page_name}`]
                "
                @change="checkboxChange($event, row)"
              ></el-checkbox>
            </div>
          </template>
        </el-table-column>
        <el-table-column label="权限列表" scoped-slot>
          <template #header>
            <span>{{ "权限列表" }}</span>
            <el-divider direction="vertical"></el-divider>
            <el-button size="mini" @click="selectAll(true)">全选</el-button>
            <el-button size="mini" @click="selectAll(false)">全不选</el-button>
            <el-button size="mini" type="primary" @click="save()"
              >保存</el-button
            >
            <el-button size="mini" type="warning" @click="reload()"
              >重新加载</el-button
            >
            <el-button size="mini" type="info" @click="goBack()"
              >返回</el-button
            >
            <span class="show-group-name"
              ><el-tag size="small"
                >当前组名：{{ groupInfo.group_name }}</el-tag
              ></span
            >
          </template>
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

import { getPriv, savePriv } from '@/api/priv.js'
import { h } from "vue"

export default {

  data () {
    return {
      privTableData: [],
      tableRowSpan: {},
      groupInfo: {},
      checkBoxSelect: {}
    };
  },
  mounted () {
    this.current = ''
    if (this.$route.params.group_name) {
      this.groupInfo.group_name = this.$route.params.group_name
    }
    if (this.$route.params.group_name) {
      this.groupInfo.group_id = parseInt(this.$route.params.group_id)
    }
    this.getPriv(false)
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

    getPriv (notice) {
      const groupInfo = {
        ...this.groupInfo
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
          this.checkBoxSelect[`${item.page_name}_${item.sub_page_name}`] = undefined
        })
        this.tableRowSpan = tableRowSpan
        this.privTableData = r.data
        if (notice) {
          this.$notify({
            title: '成功',
            message: '重新加载成功！',
            type: 'success'
          });
        }
      }).catch(e => console.log(e))
    },
    rowStyle (row) {
      let rs = {
        'padding': '2px'
      }
      return rs
    },
    cellStyle (column) {
      let cs = {
        'padding': '2px'
      }
      if (column.columnIndex === 0) {
        cs['background-color'] = '#F1FAFA'
      } else if (column.columnIndex === 1) {
        cs['background-color'] = '#F1FAFA'
      }
      return cs
    },
    checkboxChange (event, row) {
      if (event) {
        row.func_list.forEach(r => {
          r.checked = true
        })
      } else {
        row.func_list.forEach(r => {
          r.checked = false
        })
      }
    },
    allCheckBoxChange () {
      console.log('allCheckBoxChange')
    },
    renderHeader ({ column, $index }) { // h即为cerateElement的简写，具体可看vue官方文档
      return (
        h('span', [
          h('el-checkbox', {
            props: {

            }
          }),
          h('span', column.label)
        ])
      )
    },
    selectAll (value) {
      let checked
      if (value) {
        checked = true
      } else {
        checked = false
      }
      this.privTableData.forEach(d => {
        d.func_list.forEach(l => {
          l.checked = checked
        })
      })
      for (const key in this.checkBoxSelect) {
        this.checkBoxSelect[key] = checked
      }
    },
    save () {
      const groupInfo = {
        ...this.groupInfo
      }
      savePriv(groupInfo, [...this.privTableData]).then(r => {
        this.$notify({
          title: '成功',
          message: '更新成功！',
          type: 'success'
        });
      }).catch(e => console.log(e))
    },
    reload () {
      this.getPriv(true)
    },
    goBack () {
      this.$router.push({ name: 'group' })
    }
  }
};
</script>

<style scoped>
.priv-board {
  height: 100%;
}
.sub-page-class {
  display: flex;
  justify-content: space-between;
  flex-wrap: nowrap;
  flex-direction: row;
}
.show-group-name {
  float: right;
}
</style>