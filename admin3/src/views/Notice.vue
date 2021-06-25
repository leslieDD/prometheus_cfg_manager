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
                effect="plain"
                @contextmenu.prevent.native="openMenu($event, data)"
                >{{ node.label }}</el-tag
              >
              <el-tag
                v-if="data.level === 2"
                size="small"
                type="warning"
                effect="plain"
                @contextmenu.prevent.native="openMenu($event, data)"
                >{{ node.label }}</el-tag
              >
              <el-tag
                v-if="data.level === 3"
                size="small"
                type="info"
                effect="plain"
                @contextmenu.prevent.native="openMenu($event, data)"
                >{{ node.label }}</el-tag
              >
              <el-tag
                v-if="data.level === 4"
                size="small"
                type="info"
                effect="plain"
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
        <RuleEdit
          ref="ruleEditRef"
          @fatherMethod="fatherMethod"
          :query="queryInfo"
        ></RuleEdit>
      </el-scrollbar>
    </div>
    <ul
      v-show="visible"
      :style="{ left: left + 'px', top: top + 'px' }"
      class="contextmenu"
    >
      <li>
        <el-button size="mini" type="text" icon="el-icon-plus"
          >添加子节点</el-button
        >
      </li>
      <li>
        <el-button size="mini" type="text" icon="el-icon-delete"
          >删除此节点</el-button
        >
      </li>
      <li>
        <el-button size="mini" type="text" icon="el-icon-edit-outline"
          >重命名此节点</el-button
        >
      </li>
    </ul>
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
      queryInfo: {},
      visible: false,
      top: 0,
      left: 0,
      menuData: null
    }
  },
  mounted () {
    this.doGetTree()
  },
  watch: {
    visible (value) {
      if (value) {
        document.body.addEventListener('click', this.closeMenu)
      } else {
        document.body.removeEventListener('click', this.closeMenu)
      }
    }
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
    },
    fatherMethod (data, aType) {
      this.data.forEach(element => {
        if (!element.children || element.children.length === 0) {
          return
        }
        element.children.forEach(children => {
          if (!children.children || children.children.length === 0) {
            return
          }
          children.children.forEach(final => {
            if (final.id === data.id) {
              final.label = data.alert
            }
          })

        })
      })
      this.data = [...this.data]
    },
    openMenu (e, data) {
      const menuMinWidth = 105
      const offsetLeft = this.$el.getBoundingClientRect().left // container margin left
      const offsetWidth = this.$el.offsetWidth // container width
      const maxLeft = offsetWidth - menuMinWidth // left boundary
      const left = e.clientX - offsetLeft // 15: margin right

      if (left > maxLeft) {
        this.left = maxLeft
      } else {
        this.left = left
      }

      this.top = e.clientY - 60 // fix 位置bug
      this.visible = true
      this.menuData = data
    },
    closeMenu () {
      this.visible = false
      this.menuData = null
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
/* .tree-container {
  overflow: hidden;
  border: 1px solid rgb(222, 222, 224);
} */
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

.contextmenu {
  margin: 0;
  background: #fff;
  z-index: 3000;
  position: absolute;
  list-style-type: none;
  padding: 5px 0;
  border-radius: 4px;
  font-size: 12px;
  font-weight: 400;
  color: #333;
  box-shadow: 2px 2px 3px 0 rgba(0, 0, 0, 0.3);
}
.contextmenu > li {
  margin: 0;
  padding: 5px 15px;
  cursor: pointer;
}
.contextmenu > li:hover {
  background: #eee;
}
</style>
