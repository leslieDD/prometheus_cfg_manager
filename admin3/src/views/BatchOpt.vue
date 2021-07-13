<template>
  <div class="batch-box">
    <div class="batch-box-desc">
      <el-tag style="primary">导入格式说明：</el-tag>

      <el-descriptions title="自定义样式列表" :column="1" border>
        <el-descriptions-item
          label="用户名"
          label-align="right"
          align="center"
          label-class-name="my-label"
          class-name="my-content"
          width="150px"
          >kooriookami</el-descriptions-item
        >
        <el-descriptions-item label="手机号" label-align="right" align="center"
          >18100000000</el-descriptions-item
        >
        <el-descriptions-item label="居住地" label-align="right" align="center"
          >苏州市</el-descriptions-item
        >
        <el-descriptions-item label="备注" label-align="right" align="center">
          <el-tag size="small">学校</el-tag>
        </el-descriptions-item>
        <el-descriptions-item
          label="联系地址"
          label-align="right"
          align="center"
          >江苏省苏州市吴中区吴中大道 1188 号</el-descriptions-item
        >
      </el-descriptions>
    </div>
    <div class="batch-box-data">
      <div class="do_action">
        <div style="padding-right: 15px">
          <el-button size="small" type="warning" plain>转入本地文件</el-button>
          <el-button size="small" type="success" plain>提交</el-button>
        </div>
        <div>
          <el-input
            size="small"
            @keyup.enter="onSearch()"
            placeholder="请输入内容"
            v-model="searchContent"
            class="input-with-select"
          >
            <template #prepend>
              <el-select
                class="searchSelect"
                v-model="selectOption"
                placeholder="请选择"
              >
                <el-option label="IP地址" value="1"></el-option>
                <el-option label="分组" value="2"></el-option>
              </el-select>
            </template>
            <template #append>
              <el-button icon="el-icon-search" @click="onSearch()"></el-button>
            </template>
          </el-input>
        </div>
      </div>
      <el-table
        size="mini"
        highlight-current-row
        border
        :data="machines"
        stripe
        :row-style="rowStyle"
        :cell-style="cellStyle"
        header-align="center"
        @expand-change="expandChange"
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
          label="IP地址"
          prop="ipaddr"
          align="center"
          header-align="center"
        >
        </el-table-column>
        <el-table-column label="操作" align="center" header-align="center">
          <template v-slot="scope">
            <el-button type="primary" plain size="mini" @click="edit(scope)"
              >编辑</el-button
            >
            <el-button
              v-if="scope.row.enabled === true"
              type="info"
              plain
              size="mini"
              @click="invocate(scope)"
              >禁用</el-button
            >
            <el-button
              v-if="scope.row.enabled === false"
              type="warning"
              plain
              size="mini"
              @click="invocate(scope)"
              >启用</el-button
            >
            <el-popover
              :visible="deleteVisible[scope.$index]"
              placement="top"
              :width="160"
            >
              <p>确定删除吗？</p>
              <div style="text-align: right; margin: 0">
                <el-button size="mini" type="text" @click="doNo(scope)"
                  >取消</el-button
                >
                <el-button type="primary" size="mini" @click="doYes(scope)"
                  >确定</el-button
                >
              </div>
              <template #reference>
                <el-button
                  size="mini"
                  type="danger"
                  plain
                  @click="doDelete(scope)"
                  >删除</el-button
                >
              </template>
            </el-popover>
          </template>
        </el-table-column>
      </el-table>
      <div class="block">
        <el-pagination
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
          :current-page="currentPage"
          :page-sizes="[15, 20, 30, 50]"
          :page-size="pageSize"
          layout="total, sizes, prev, pager, next, jumper"
          :total="pageTotal"
        >
        </el-pagination>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  data () {

  },
}
</script>

<style scoped>
.batch-box {
  display: flex;
  justify-content: flex-start;
  flex-wrap: nowrap;
  height: 85vh;
}
.batch-box-desc {
  width: 600px;
}
.batch-box-data {
  width: 100%;
}
.do_action {
  display: flex;
  justify-content: flex-end;
  flex-wrap: nowrap;
  margin-bottom: 5px;
}
.block {
  padding-top: 12px;
}
</style>