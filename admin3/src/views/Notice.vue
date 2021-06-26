<template>
  <div class="monitor-board">
    <div class="tree-container">
      <el-scrollbar class="card-scrollbar">
        <div class="tree-box">
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
                <el-input
                  size="mini"
                  v-if="data.display === true"
                  placeholder="输入名称"
                  v-model="titleFromShowMe"
                  @keyup.enter="receiveEnter"
                  @keyup.esc="receiveEsc"
                  clearable
                ></el-input>
                <el-tag
                  v-show="data.showMe !== true"
                  v-if="data.level === 1 && data.display !== true"
                  size="small"
                  type="success"
                  effect="plain"
                  @click="closeMenu"
                  @contextmenu.prevent.native="openMenu($event, data, node)"
                  >{{ node.label }}</el-tag
                >
                <el-tag
                  v-show="data.showMe !== true && data.display !== true"
                  v-if="data.level === 2"
                  size="small"
                  type="warning"
                  effect="plain"
                  @click="closeMenu"
                  @contextmenu.prevent.native="openMenu($event, data, node)"
                  >{{ node.label }}</el-tag
                >
                <el-tag
                  v-show="data.showMe !== true && data.display !== true"
                  v-if="data.level === 3"
                  size="small"
                  type="info"
                  effect="plain"
                  @click="closeMenu"
                  @contextmenu.prevent.native="openMenu($event, data, node)"
                  >{{ node.label }}</el-tag
                >
                <el-tag
                  v-show="data.showMe !== true && data.display !== true"
                  v-if="data.level === 4"
                  size="small"
                  type="info"
                  effect="plain"
                  @click="closeMenu"
                  @contextmenu.prevent.native="openMenu($event, data, node)"
                  >{{ node.label }}</el-tag
                >
              </span>
            </template>
          </el-tree>
        </div>
      </el-scrollbar>
    </div>
    <div class="node-content">
      <div class="node-content-top">
        <span class="node-content-path">
          <el-tag size="small" class="path-all" type="success">路径：</el-tag>
          <el-tag
            class="path-list path-all"
            size="small"
            type="warning"
            effect="plain"
            >{{ labelPath.join(" > ") }}</el-tag
          >
        </span>
        <span><el-button type="primary" size="mini">发布</el-button></span>
      </div>
      <div class="node-content-edit">
        <el-scrollbar class="card-scrollbar-right">
          <div class="rule-edit-box">
            <RuleEdit
              ref="ruleEditRef"
              @fatherMethod="fatherMethod"
              :query="queryInfo"
            ></RuleEdit>
          </div>
        </el-scrollbar>
      </div>
    </div>
    <ul
      v-show="visible"
      :style="{ left: left + 'px', top: top + 'px' }"
      class="contextmenu"
    >
      <li>
        <el-button
          v-bind:disabled="menuAddDisabled"
          size="mini"
          type="text"
          icon="el-icon-plus"
          @click="append(null)"
          >添加子节点</el-button
        >
      </li>
      <li>
        <el-button
          v-bind:disabled="menuRenameDisabled"
          size="mini"
          type="text"
          icon="el-icon-edit-outline"
          @click="edit(null, null)"
          >重命名此节点</el-button
        >
      </li>
      <li>
        <el-button
          v-bind:disabled="menuDelDisabled"
          size="mini"
          type="text"
          icon="el-icon-delete"
          @click="remove(null, null)"
          >删除此节点</el-button
        >
      </li>
    </ul>
  </div>
</template>

<script>
import RuleEdit from '@/views/RuleEdit.vue'
// import Leslie from '@/views/Leslie.vue'
import { getTree, createTreeNode, updateTreeNode, removeTreeNode } from '@/api/monitor.js'
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
      menuData: null,
      menuNode: null,
      titleFromShowMe: '',
      menuAddDisabled: true,
      menuDelDisabled: true,
      menuRenameDisabled: true,
      labelPath: []
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
          this.data = [...r.data]
        }
      ).catch(
        e => { console.log(e) }
      )
    },
    handleNodeClick (data) {
      this.labelPath = data.path
      if (data.level !== 4) {
        this.$refs.ruleEditRef.resetForm()
        this.$refs.ruleEditRef.setFormDisable()
        return false
      }
      this.$refs.ruleEditRef.setFormEnable()
      if (data.label === '[must rename me]') {
        return false
      }
      const queryInfo = {
        id: data.id,
        label: data.label,
        level: data.level,
        is_new: data.is_new,
        parent: data.parent
      }
      this.$refs.ruleEditRef.doGetRuleDetail(queryInfo)
    },
    append (data) {
      if (!data) {
        data = this.menuData
      }
      const newChild = {
        id: id++,
        level: data.level + 1,
        label: '[must rename me]',
        children: [],
        is_new: true,
        parent: data.id,
        path: [...data.path]
      }
      newChild.path.push(newChild.label)
      if (!data.children) {
        data.children = []
      }
      data.children.push(newChild);
      this.data = [...this.data]
    },

    remove (node, data) {
      if (!node) {
        node = this.menuNode
      }
      if (!data) {
        data = this.menuData
      }
      this.$confirm('是否确定删除？', '确认信息', {
        distinguishCancelAndClose: true,
        confirmButtonText: '确定',
        cancelButtonText: '放弃'
      }).then(() => {
        if (data.is_new) {
          const parent = node.parent;
          const children = parent.data.children || parent.data;
          const index = children.findIndex(d => d.id === data.id);
          children.splice(index, 1);
          this.data = [...this.data]
          return true
        }
        const nodeInfo = {
          id: this.menuData.id,
          label: this.titleFromShowMe,
          level: this.menuData.level,
          parent: this.menuData.parent
        }
        removeTreeNode(nodeInfo).then(
          r => {
            this.$notify({
              title: '成功',
              message: '删除节点成功！',
              type: 'success'
            })
            const parent = node.parent;
            const children = parent.data.children || parent.data;
            const index = children.findIndex(d => d.id === data.id);
            children.splice(index, 1);
            this.data = [...this.data]
          }
        ).catch(
          e => { console.log(e) }
        )
      }).catch(() => { })
    },

    edit (node, data) {
      if (!node) {
        node = this.menuNode
      }
      if (!data) {
        data = this.menuData
      }
      data.display = true
      this.titleFromShowMe = data.label
      //   this.data = [...this.data]
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
      if (data.path) {
        this.labelPath = data.path
      }
      this.doGetTree()
      //   this.data.forEach(element => {
      //     if (!element.children || element.children.length === 0) {
      //       return
      //     }
      //     element.children.forEach(children => {
      //       if (!children.children || children.children.length === 0) {
      //         return
      //       }
      //       children.children.forEach(final => {
      //         if (final.id === data.id) {
      //           final.label = data.alert
      //         }
      //       })
      //     })
      //   })
      //   this.data = [...this.data]
    },
    openMenu (e, data, node) {
      if (data.level === 1) {
        this.menuDelDisabled = true
        this.menuRenameDisabled = true
        this.menuAddDisabled = false
      } else if (data.level === 2 || data.level === 3) {
        if (data.label === '[must rename me]') {
          this.menuAddDisabled = true
        } else {
          this.menuAddDisabled = false
        }
        this.menuDelDisabled = false
        this.menuRenameDisabled = false
      } else if (data.level === 4) {
        if (data.label === '[must rename me]') {
          this.menuRenameDisabled = false
        } else {
          this.menuRenameDisabled = true
        }
        this.menuDelDisabled = false
        this.menuAddDisabled = true
      } else {
        this.menuDelDisabled = true
        this.menuRenameDisabled = true
        this.menuAddDisabled = true
      }
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

      this.top = e.clientY - 40 // fix 位置bug
      this.left = this.left + 20
      //   this.top = e.clientY - 60 // fix 位置bug
      this.menuData = data.valueOf()
      this.menuNode = node.valueOf()
      this.visible = true
    },
    receiveEsc () {
      this.titleFromShowMe = ''
      this.menuData.display = false
      this.data = [...this.data]
    },
    receiveEnter () {
      if (this.titleFromShowMe === '[must rename me]') {
        this.$notify({
          title: '错误',
          message: '请输入正确的名称！',
          type: 'error'
        })
        return false
      }
      const nodeInfo = {
        id: this.menuData.id,
        label: this.titleFromShowMe,
        level: this.menuData.level,
        parent: this.menuData.parent
      }
      if (this.menuData.is_new) {
        createTreeNode(nodeInfo).then(
          r => {
            this.$notify({
              title: '成功',
              message: '创建新节点成功！',
              type: 'success'
            })
            this.menuData.label = this.titleFromShowMe
            this.titleFromShowMe = ''
            this.menuData.display = false
            // this.data = [...this.data]
            this.doGetTree()
            // this.handleNodeClick()
          }
        ).catch(
          e => { console.log(e) }
        )
      } else {
        updateTreeNode(nodeInfo).then(
          r => {
            this.$notify({
              title: '成功',
              message: '更新节点成功！',
              type: 'success'
            })
            this.menuData.label = this.titleFromShowMe
            this.menuData.display = false
            this.titleFromShowMe = ''
            this.doGetTree()
          }
        ).catch(
          e => { console.log(e) }
        )
      }
    },
    closeMenu () {
      this.visible = false
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
.node-content {
  width: 100%;
  display: flex;
  flex-direction: column;
}
.tree-box {
  min-height: 150vh;
}
.rule-edit-box {
  min-height: 90vh;
}
.node-content-top {
  display: flex;
}
.node-content-path {
  display: flex;
  flex-direction: row;
}
.path-list {
  width: 600px;
  margin-right: 10px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
.path-all {
  height: 28px;
}
.card-scrollbar {
  height: 85vh;
  width: 350px;
}
.card-scrollbar-right {
  height: 83vh;
  /* width: 100%; */
}
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
