<template>
  <div class="admin-set-board">
    <el-descriptions class="margin-top" :column="1" size="mini" border>
      <el-descriptions-item>
        <template #label>
          <i class="el-icon-s-custom"></i>
          当在调整用户所属组时，未被设置的用户默认指向的组
        </template>
        <el-select
          size="small"
          v-model="params.default_group"
          clearable
          placeholder="请选择"
          @visible-change="doPutAdminSetting"
          @clear="doPutAdminSetting"
        >
          <el-option
            v-for="item in groupList"
            :key="item.id"
            :label="item.name"
            :value="item.id"
          >
          </el-option>
        </el-select>
      </el-descriptions-item>
    </el-descriptions>
  </div>
</template>

<script>

import { getGroupList, getAdminSetting, putAdminSetting } from '@/api/manager.js'

export default {
  data () {
    return {
      params: {},
      groupList: []
    }
  },
  mounted () {
    this.doGetParams()
  },
  methods: {
    doGetParams () {
      getGroupList().then(r => {
        this.groupList = r.data
        getAdminSetting().then(r => {
          Object.keys(r.data).forEach(k => {
            if (k === 'default_group') {
              var groupIndex = this.groupList.findIndex((item, index) => {
                return item.id.toString() === r.data[k];
              })
              if (groupIndex === -1) {
                this.params[k] = ''
              } else {
                this.params[k] = this.groupList[groupIndex].name
              }
            }
          })
        }).catch(e => console.log(e))
      }).catch(e => console.log(e))
    },
    doPutAdminSetting (action) {
      if (!action) {
        let params = {}
        Object.keys(this.params).forEach(k => {
          params[k] = this.params[k].toString()
        })
        putAdminSetting(params).then(r => {
          this.$notify({
            title: '成功',
            message: '参数提交成功',
            type: 'success'
          })
        }
        ).catch(e => console.log(e))
      }
    }
  }
}
</script>

<style scoped>
.admin-set-board {
  width: 650px;
}
</style>