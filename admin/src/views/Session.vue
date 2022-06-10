<template>
  <div class="user-board">
    <div class="main-content">
      <div>
        <div class="do_action">
          <div style="padding-right: 15px">
            <el-form
              ref="sessionExpire"
              :model="sessionExpire"
              label-width="170px"
              class="demo-ruleForm"
              size="small"
              :inline="true"
            >
              <el-form-item
                label="会话过期时间(分钟)"
                prop="expire"
                :rules="[
                  { required: true, message: '会话过期字段必须填写' },
                  { type: 'number', message: '会话过期字段必须得是一个数字' },
                ]"
              >
                <div class="expire_class">
                  <el-input
                    v-model.number="sessionExpire.expire"
                    type="text"
                    autocomplete="off"
                    style="width: 120px;"
                  />
                  <el-button type="primary" @click="submitSessionExpire()">提交</el-button>
                </div>
              </el-form-item>
            </el-form>
          </div>
          <div>
            <el-input
              size="small"
              placeholder="请输入内容"
              @keyup.enter="onSearch()"
              v-model="searchContent"
            >
              <template #append>
                <el-button
                  size="small"
                  @click="onSearch()"
                  icon="el-icon-search"
                ></el-button>
              </template>
            </el-input>
          </div>
        </div>
        <div class="table-show">
          <el-scrollbar height="75vh" class="flex-content">
            <el-table
              size="mini"
              highlight-current-row
              border
              :data="sessions"
              stripe
              :row-style="rowStyle"
              :cell-style="cellStyle"
            >
              <el-table-column label="序号" width="50px">
                <template v-slot="scope">
                  {{ scope.$index + 1 }}
                </template>
              </el-table-column>
              <el-table-column label="会话Token" prop="token" width="260px">
              </el-table-column>
              <el-table-column label="登录地址" prop="ipaddr" width="270px">
              </el-table-column>
              <el-table-column label="登录名" prop="username">
                <template v-slot="{ row }">
                  <el-button
                    size="mini"
                    @click="routeToUser(row)"
                    type="text"
                    >{{ row.username }}</el-button
                  >
                </template>
              </el-table-column>
              <el-table-column label="所属组" prop="group_name">
                <template v-slot="{ row }">
                  <el-button
                    size="mini"
                    @click="routeToGroup(row)"
                    type="text"
                    >{{ row.group_name }}</el-button
                  >
                </template>
              </el-table-column>
              <el-table-column
                label="更新时间"
                width="135px"
                prop="update_at"
              >
                <template v-slot="{ row }">
                  <span>{{ parseTimeSelf(row.update_at) }}</span>
                </template>
              </el-table-column>
              <el-table-column label="操作" align="center" width="200px">
                <template v-slot="scope" align="center">
                  <div class="actioneara">
                    <div>
                      <el-popover
                        :visible="deleteVisible[scope.$index]"
                        placement="top"
                      >
                        <p>确定删除吗？</p>
                        <div style="text-align: right; margin: 0">
                          <el-button
                            size="mini"
                            type="text"
                            @click="doNo(scope)"
                            >取消</el-button
                          >
                          <el-button
                            type="primary"
                            size="mini"
                            @click="doYes(scope)"
                            >确定</el-button
                          >
                        </div>
                        <template #reference>
                          <el-button
                            size="mini"
                            type="danger"
                            plain
                            @click="doDelete(scope)"
                            :disabled="editCodeButVisable[scope.row.id]"
                            >删除</el-button
                          >
                        </template>
                      </el-popover>
                    </div>
                  </div>
                </template>
              </el-table-column>
            </el-table>
            <div class="block" v-if="paginationShow">
              <el-pagination
                @size-change="handleSizeChange"
                @current-change="handleCurrentChange"
                :current-page="currentPage"
                :page-sizes="[20, 30, 50]"
                :page-size="pageSize"
                layout="total, sizes, prev, pager, next, jumper"
                :total="pageTotal"
              >
              </el-pagination>
            </div>
          </el-scrollbar>
        </div>
      </div>
    </div>
  </div>
</template>

<script>

import {
  getSession,
  delSession,
  updateSessionParams,
  getSessionParams
} from '@/api/session.js'

export default {
  name: 'ManagerUser',
  data () {
    return {
      sessions: [],
      editCodeButVisable: {},
      pageSize: 20,
      pageTotal: 0,
      currentPage: 1,
      searchContent: '',
      deleteVisible: {},
      paginationShow: true,
      sessionExpire: {
        expire: 0
      }
    }
  },
  mounted () {
    this.doGetSession()
    this.doGetSessionParams()
  },
  methods: {
    doGetSession (getInfo) {
      if (!getInfo) {
        getInfo = {
          'pageNo': this.currentPage,
          'pageSize': this.pageSize,
          'search': this.searchContent
        }
      }
      getSession(getInfo).then(r=>{
        this.sessions = r.data.data
        this.pageTotal = r.data.totalCount
        this.currentPage = r.data.pageNo
        this.pageSize = r.data.pageSize
        this.paginationShow = false;//让分页隐藏
        this.$nextTick(() => {//重新渲染分页
          this.paginationShow = true;
        });
      }).catch(e=>{
        console.log(e)
      })
    },
    handleSizeChange (val) {
      let getInfo = {
        'pageNo': this.currentPage,
        'pageSize': val,
        'search': this.searchContent
      }
      this.doGetSession(getInfo)
    },
    handleCurrentChange (val) {
      let getInfo = {
        'pageNo': val,
        'pageSize': this.pageSize,
        'search': this.searchContent
      }
      this.doGetSession(getInfo)
    },
    onSearch () {
      this.doGetSession()
    },
    parseTimeSelf (t) {
      var time = new Date(Date.parse(t))
      return time.toLocaleDateString() + ' ' + time.toTimeString().split(' ')[0]
    },
    doYes (scope) {
      delSession({ id: scope.row.id }).then(
        r => {
          this.$notify({
            title: '成功',
            message: '删除成功！',
            type: 'success'
          });
          this.doGetSession()
        }
      ).catch(
        e => {
          console.log(e)
        }
      )
      this.deleteVisible[scope.$index] = false
    },
    doNo (scope) {
      this.deleteVisible[scope.$index] = false
    },
    doDelete (scope) {
      this.deleteVisible[scope.$index] = true
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
    routeToGroup (row) {
      this.$router.push({ name: 'group', params: { group_name: row.group_name } })
    },
    routeToUser(row){
      this.$router.push({ name: 'user', params: { group_name: row.username } })
    },
    submitSessionExpire(){
      // console.log(this.sessionExpire)
      var data = {...this.sessionExpire}
      updateSessionParams(data).then(r=>{
          this.$notify({
            title: '成功',
            message: '提交成功！',
            type: 'success'
          });
      }).catch(e=>console.log(e))
    },
    doGetSessionParams(){
      getSessionParams().then(r=>{
        r.data.forEach(element => {
          if (element['key'] === 'session_expire') {
            this.sessionExpire = {
              expire: element['value']
            }
          }
        });
      })
    }
  }
};
</script>

<style scoped>
.main-box {
  height: 85vh;
}
.main-content {
  width: 100%;
  height: 40vh;
}
.do_action {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  margin-bottom: 5px;
}

.expire_class {
  display: flex;
  flex-wrap: nowrap;
  justify-content: left;
}
.actioneara {
  display: flex;
  flex-wrap: nowrap;
  justify-content: center;
}
.el-pagination {
  text-align: center;
}
.block {
  padding-top: 12px;
}
.cm-s-monokai {
  height: 100%;
}
</style>
