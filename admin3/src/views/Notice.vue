<template>
  <div class="monitor-board">
    <div class="tree-container">
      <el-scrollbar class="card-scrollbar">
        <el-tree
          class="tree"
          :data="data"
          :indent="0"
          node-key="id"
          default-expand-all
          :expand-on-click-node="false"
          @node-click="handleNodeClick"
        >
          <template #default="{ node, data }">
            <span class="custom-tree-node">
              <el-tag
                v-if="data.level === 1"
                size="small"
                type="success"
                effect="dark"
                >{{ node.label }}</el-tag
              >
              <el-tag
                v-if="data.level === 2"
                size="small"
                type="warning"
                plain
                >{{ node.label }}</el-tag
              >
              <el-tag
                v-if="data.level === 3 || data.level === 4"
                size="small"
                type="info"
                plain
                >{{ node.label }}</el-tag
              >
              <!-- <span v-if="data.level === 1"> </span>
              <span v-if="data.level === 2 || data.level === 3">
                <i
                  class="el-icon-plus icon-action"
                  @click="append(data)"
                  title="增加"
                ></i>
                <i
                  class="el-icon-minus icon-action"
                  @click="remove(node, data)"
                  title="删除"
                ></i>
                <i
                  class="el-icon-edit icon-action"
                  @click="edit(node, data)"
                  title="编辑"
                ></i>
              </span>
              <span v-if="data.level === 4">
                <i
                  class="el-icon-minus icon-action"
                  @click="remove(node, data)"
                  title="删除"
                ></i>
                <i
                  class="el-icon-edit icon-action"
                  @click="edit(node, data)"
                  title="编辑"
                ></i>
              </span> -->
            </span>
          </template>
        </el-tree>
      </el-scrollbar>
    </div>
    <div class="node-content-edit">
      <el-scrollbar class="card-scrollbar-right">
        <RuleEdit ref="ruleEditRef" :query="queryInfo"></RuleEdit>
      </el-scrollbar>
    </div>
  </div>
</template>

<script>
import RuleEdit from '@/views/RuleEdit.vue'
// import Leslie from '@/views/Leslie.vue'
import { getTree, getRuleDetail } from '@/api/monitor.js'
let id = 1000;

export default {
  name: "Notice",
  components: {
    RuleEdit: RuleEdit
  },
  data () {
    return {
      data: [],
      queryInfo: {}
    }
  },
  mounted () {
    this.doGetTree()
  },
  methods: {
    doGetTree () {
      getTree().then(
        r => {
          this.data = r.data
        }
      ).catch(
        e => { console.log(e) }
      )
    },
    handleNodeClick (data) {
      if (data.level !== 4) {
        return false
      }
      this.queryInfo = {
        id: data.id,
        label: data.label,
        level: data.level
      }
      this.$refs.ruleEditRef.doGetRuleDetail(this.queryInfo)
    },
    append (data) {
      const newChild = { id: id++, level: 3, label: 'testtest', children: [] };
      if (!data.children) {
        data.children = []
      }
      data.children.push(newChild);
      this.data = [...this.data]
    },

    remove (node, data) {
      const parent = node.parent;
      const children = parent.data.children || parent.data;
      const index = children.findIndex(d => d.id === data.id);
      children.splice(index, 1);
      this.data = [...this.data]
    },

    edit (node, data) {
      const parent = node.parent;
      const children = parent.data.children || parent.data;
      const index = children.findIndex(d => d.id === data.id);
      children.splice(index, 1);
      this.data = [...this.data]
    },

    renderContent (h, { node, data, store }) {
      return h("span", {
        class: "custom-tree-node"
      }, h("span", null, node.label), h("span", null, h("a", {
        onClick: () => this.append(data)
      }, "Append "), h("a", {
        onClick: () => this.remove(node, data)
      }, "Delete")));
    }
  }
};
</script>

// 以下为scss，记得去掉scoped，或者使用/deep/
<style scoped>
.monitor-board {
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
  height: 85vh;
  width: 100%;
}
.node-content-edit {
  width: 100%;
  /* border: 1px solid royalblue; */
}
.card-scrollbar {
  height: 85vh;
  width: 250px;
}
/* .card-scrollbar-right {
  height: 84vh;
  width: 100%;
} */
.icon-action {
  font-size: 5px;
  margin-left: 5px;
  color: green;
}
.tree-container {
  /* overflow: hidden; */
  /* border: 1px solid rgb(222, 222, 224); */
}
.tree :deep() .el-tree-node {
  position: relative;
  padding-left: 0;
}

.tree :deep() .el-tree-node__children {
  padding-left: 16px;
}

.tree :deep() .el-tree-node :last-child:before {
  height: 12px;
}

.treev :deep() .el-tree > .el-tree-node:before {
  border-left: none;
}

.tree-container :deep() .el-tree > .el-tree-node:after {
  border-top: none;
}

.tree :deep() .el-tree-node:before {
  content: "";
  left: -4px;
  position: absolute;
  right: auto;
  border-width: 1px;
  border-left: 1px dashed #4386c6;
  bottom: 0px;
  height: 100%;
  top: 0px;
  width: 1px;
}

.tree :deep() .el-tree-node:after {
  content: "";
  left: -4px;
  position: absolute;
  right: auto;
  border-width: 1px;
  border-top: 1px dashed #4386c6;
  height: 20px;
  top: 12px;
  width: 10px;
}
</style>
