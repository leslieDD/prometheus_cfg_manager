<template>
  <div class="tree-container">
    <el-tree class="tree" :data="data" :indent="0"> </el-tree>
  </div>
</template>

<script>
let id = 1000;

export default {
  data () {
    const data = [{
      id: 1,
      label: '一级 1',
      is_child: false,
      children: [{
        id: 4,
        label: '二级 1-1',
        is_child: false,
        children: [{
          id: 9,
          label: '三级 1-1-1',
          is_child: true
        }, {
          id: 10,
          label: '三级 1-1-2',
          is_child: true
        }]
      }]
    }, {
      id: 2,
      label: '一级 2',
      is_child: false,
      children: [{
        id: 5,
        label: '二级 2-1',
        is_child: true
      }, {
        id: 6,
        label: '二级 2-2',
        is_child: true
      }]
    }, {
      id: 3,
      label: '一级 3',
      is_child: false,
      children: [{
        id: 7,
        label: '二级 3-1',
        is_child: true
      }, {
        id: 8,
        label: '二级 3-2',
        is_child: true
      }]
    }];
    return {
      data: JSON.parse(JSON.stringify(data)),
    }
  },

  methods: {
    append (data) {
      const newChild = { id: id++, label: 'testtest', children: [] };
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
