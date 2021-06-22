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
                effect="dark"
                >{{ node.label }}</el-tag
              >
              <el-tag
                v-if="data.level === 3"
                size="small"
                type="info"
                effect="dark"
                >{{ node.label }}</el-tag
              >
              <span v-if="data.level === 1"> </span>
              <span v-if="data.level === 2">
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
                  title="删除"
                ></i>
              </span>
              <span v-if="data.level === 3">
                <i
                  class="el-icon-minus icon-action"
                  @click="remove(node, data)"
                  title="删除"
                ></i>
                <i
                  class="el-icon-edit icon-action"
                  @click="edit(node, data)"
                  title="删除"
                ></i>
              </span>
            </span>
          </template>
        </el-tree>
      </el-scrollbar>
    </div>
    <div>
      <!-- <RuleEdit :monitorData="{ a: 1 }"></RuleEdit> -->
      <div>
        <Leslie>
          <span slot="diandian">
            <span>Hello!, DianDian</span>
          </span>
        </Leslie>
      </div>
    </div>
  </div>
</template>

<script>
import RuleEdit from '@/views/RuleEdit.vue'
import Leslie from '@/views/Leslie.vue'
let id = 1000;

export default {
  name: "Notice",
  components: {
    RuleEdit: RuleEdit,
    Leslie: Leslie
  },
  data () {
    const data = [
      {
        id: 1,
        label: '监控规则列表',
        level: 1,
        children: [
          {
            id: 2,
            label: '一级 1',
            level: 2,
            children: [{
              id: 5,
              label: '二级 1-1',
              level: 3
            }]
          },
          {
            id: 3,
            label: '一级 2',
            level: 2,
            children: [{
              id: 6,
              label: '二级 2-1',
              level: 3
            }, {
              id: 7,
              label: '二级 2-2',
              level: 3
            }]
          },
          {
            id: 4,
            label: '一级 3',
            level: 2,
            children: [{
              id: 8,
              label: '二级 3-1',
              level: 3
            }, {
              id: 9,
              label: '二级 3-2',
              level: 3
            }]
          }
        ]
      },
    ];
    return {
      data: JSON.parse(JSON.stringify(data))
    }
  },

  methods: {
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
}
.card-scrollbar {
  height: 84vh;
  width: 250px;
}
.icon-action {
  font-size: 5px;
  margin-left: 5px;
  color: green;
}
.tree-container {
  overflow: hidden;
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
