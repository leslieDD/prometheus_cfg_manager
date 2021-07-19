<template>
  <div class="monitor-board">
    <el-card class="box-card" shadow="never">
      <template #header>
        <div class="card-header">
          <div class="node-content-top">
            <div class="node-content-path">
              <el-breadcrumb separator="/">
                <el-breadcrumb-item
                  v-for="(pName, index) in labelPath"
                  :key="index"
                  >{{ pName }}</el-breadcrumb-item
                >
              </el-breadcrumb>
            </div>
            <div>
              <el-button
                v-if="pushing === false"
                type="primary"
                icon="el-icon-upload"
                size="mini"
                @click="doRulesPublish"
                >发布</el-button
              >
              <el-button
                v-if="pushing === true"
                type="primary"
                icon="el-icon-loading"
                size="mini"
                >发布</el-button
              >
            </div>
          </div>
        </div>
      </template>
      <div class="node-content">
        <div class="tree-container">
          <el-scrollbar noresize class="card-scrollbar flex-content">
            <div class="tree-box">
              <el-tree
                :default-expanded-keys="expandedList"
                node-key="code"
                class="tree"
                :data="treeData"
                :indent="0"
                :expand-on-click-node="true"
                highlight-current
                icon-class="none"
                @node-click="handleNodeClick"
                @node-contextmenu="handleNodeRightClick"
                @node-expand="nodeExpand"
                @node-collapse="nodeCollapse"
                @keyup.esc="closeMenu"
              >
                <template #default="{ node, data }">
                  <div>
                    <single-svg v-if="data.level === 1" icon-class="root" />
                    <single-svg
                      v-if="data.level === 2"
                      icon-class="footprint_unknow"
                    />
                    <single-svg
                      v-if="data.level === 3"
                      icon-class="footprint_chongwu"
                    />
                    <single-svg
                      v-if="data.level === 4"
                      icon-class="footprint_person"
                    />
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
                      <el-badge
                        v-show="data.showMe !== true"
                        v-if="data.level === 1 && data.display !== true"
                        type="success"
                        :value="data.children ? data.children.length : 0"
                        class="item"
                      >
                        <el-tag
                          size="small"
                          type="success"
                          effect="plain"
                          @click="closeMenu"
                          @keyup.esc="closeMenu"
                          @contextmenu.prevent.native="
                            openMenu($event, data, node)
                          "
                          >{{ node.label }}</el-tag
                        >
                      </el-badge>
                      <el-badge
                        v-show="data.showMe !== true && data.display !== true"
                        v-if="data.level === 2"
                        type="success"
                        :value="data.children ? data.children.length : 0"
                        class="item"
                      >
                        <el-tag
                          size="small"
                          type="warning"
                          effect="plain"
                          @click="closeMenu"
                          @keyup.esc="closeMenu"
                          @contextmenu.prevent.native="
                            openMenu($event, data, node)
                          "
                          >{{ node.label }}</el-tag
                        >
                      </el-badge>
                      <el-badge
                        v-show="data.showMe !== true && data.display !== true"
                        v-if="data.level === 3"
                        type="success"
                        :value="data.children ? data.children.length : 0"
                        class="item"
                      >
                        <el-tag
                          size="small"
                          type="info"
                          effect="plain"
                          @click="closeMenu"
                          @keyup.esc="closeMenu"
                          @contextmenu.prevent.native="
                            openMenu($event, data, node)
                          "
                          >{{ node.label }}</el-tag
                        >
                      </el-badge>
                      <el-tag
                        v-show="data.showMe !== true && data.display !== true"
                        v-if="data.level === 4"
                        size="small"
                        type="info"
                        effect="plain"
                        @click="closeMenu"
                        @keyup.esc="closeMenu"
                        @contextmenu.prevent.native="
                          openMenu($event, data, node)
                        "
                        >{{ node.label }}</el-tag
                      >
                    </span>
                  </div>
                </template>
              </el-tree>
            </div>
          </el-scrollbar>
        </div>
        <div class="node-content-edit">
          <el-scrollbar
            v-show="showEditArea === true"
            class="card-scrollbar-right"
          >
            <div class="rule-edit-box">
              <RuleEdit
                ref="ruleEditRef"
                @fatherMethod="fatherMethod"
                :query="queryInfo"
              ></RuleEdit>
            </div>
          </el-scrollbar>
          <div class="empty-board" v-show="showEditArea !== true">
            <el-empty description="请选择一个规则（叶子节点）"></el-empty>
          </div>
        </div>
      </div>
    </el-card>
    <div>
      <ul
        v-show="visible"
        :style="{ left: left + 'px', top: top + 'px' }"
        class="contextmenu"
        @keyup.esc="closeMenu"
      >
        <li>
          <el-button
            v-bind:disabled="menuAddDisabled"
            size="mini"
            icon="el-icon-plus"
            @click="append(null)"
            @keyup.esc="closeMenu"
            >{{ btnTitleAppend }}</el-button
          >
        </li>
        <!-- <el-divider class="divider-class"></el-divider> -->
        <li>
          <el-button
            v-bind:disabled="menuRenameDisabled"
            size="mini"
            icon="el-icon-edit-outline"
            @click="edit(null, null)"
            @keyup.esc="closeMenu"
            >重命名此节点</el-button
          >
        </li>
        <li>
          <el-button
            v-bind:disabled="menuImportFromFile"
            size="mini"
            icon="el-icon-document-add"
            @click="importRules(null, null)"
            @keyup.esc="closeMenu"
            >从文件导入规则</el-button
          >
        </li>
        <!-- <el-divider class="divider-class"></el-divider> -->
        <li>
          <el-button
            v-bind:disabled="menuAddDisabled"
            size="mini"
            icon="el-icon-notebook-2"
            @click="expand(null)"
            @keyup.esc="closeMenu"
            >展开所有节点</el-button
          >
        </li>
        <!-- <el-divider class="divider-class"></el-divider> -->
        <li>
          <el-button
            size="mini"
            icon="el-icon-open"
            @click="enabled(null, null)"
            @keyup.esc="closeMenu"
            >启用此分支</el-button
          >
        </li>
        <li>
          <el-button
            size="mini"
            icon="el-icon-turn-off"
            @click="disabled(null, null)"
            @keyup.esc="closeMenu"
            >禁用此分支</el-button
          >
        </li>
        <!-- <el-divider class="divider-class"></el-divider> -->
        <li>
          <el-button
            v-bind:disabled="menuDelDisabled"
            size="mini"
            icon="el-icon-delete"
            @click="remove(null, null)"
            @keyup.esc="closeMenu"
            >删除此节点</el-button
          >
        </li>
      </ul>
    </div>
    <div class="dialog-area">
      <el-dialog
        :title="'为组：' + ruleGroupName + ' 导入规则'"
        v-model="importDialogDisplay"
        width="800px"
        :before-close="handleDialogClose"
      >
        <el-descriptions class="margin-top" :column="1" size="mini" border>
          <el-descriptions-item>
            <template #label>
              <!-- <i class="el-icon-user"></i> -->
              说明
            </template>
            导入的数据，必须是符合yaml语法，文件必须是如下“样例”展示的格式，
            顶层为数组格式，数组元素个数不限制
          </el-descriptions-item>
          <el-descriptions-item>
            <template #label>样例</template>
            <div style="width: 650px">
              <el-scrollbar height="300px" class="flex-content">
                <pre v-highlight="importCodeExample"><code></code></pre>
              </el-scrollbar>
            </div>
          </el-descriptions-item>
          <el-descriptions-item>
            <template #label>选择文件</template>
            <el-upload
              class="upload-demo"
              ref="uploadYamlRef"
              action="/v1/tree/upload/file/yaml"
              :auto-upload="false"
              accept="text/plain"
              :file-list="fileList"
              :show-file-list="true"
              :http-request="submitUpload"
              :before-upload="importBefore"
              type="file"
            >
              <template #trigger>
                <el-button size="small" type="primary">导入文件</el-button>
              </template>
              <el-button
                v-if="importPushing === false"
                style="margin-left: 10px"
                size="small"
                type="success"
                icon="el-icon-upload"
                @click="doSubmitUpload"
                >上传文件到服务器</el-button
              >
              <el-button
                v-if="importPushing === true"
                style="margin-left: 10px"
                size="small"
                type="success"
                icon="el-icon-loading"
                @click="doSubmitUpload"
                >上传文件到服务器</el-button
              >
              <template #tip>
                <div class="el-upload__tip">文件格式：*.txt, *.yml, *.yaml</div>
              </template>
            </el-upload>
            <!-- <el-upload
              class="upload-demo"
              ref="upload"
              action="/v1/update/ssss"
              accept="text/plain"
              :file-list="fileList"
              :auto-upload="false"
              :before-upload="importBefore"
              :http-request="submitUpload"
            >
              <template #trigger>
                <el-button size="small" type="primary">选取文件</el-button>
              </template>
              <el-button
                style="margin-left: 10px"
                size="small"
                type="success"
                @click="doSubmitUpload"
                >上传到服务器</el-button
              >
              <template #tip>
                <div class="el-upload__tip">
                  只能上传 jpg/png 文件，且不超过 500kb
                </div>
              </template>
            </el-upload> -->
          </el-descriptions-item>
          <el-descriptions-item>
            <template #label>结果</template>
            <el-scrollbar height="60px">
              <div v-if="showError">
                <!-- <pre v-highlight="error"><code></code></pre> -->
              </div>
            </el-scrollbar>
          </el-descriptions-item>
        </el-descriptions>
      </el-dialog>
    </div>
  </div>
</template>

<script>
import RuleEdit from '@/views/RuleEdit.vue'
import md5 from 'js-md5';
import {
  getTree,
  createTreeNode,
  updateTreeNode,
  removeTreeNode,
  rulesPublish,
  disableTreeNode,
  uploadYamlFile
} from '@/api/monitor.js'
let menuid = 1000;

export default {
  name: "Notice",
  components: {
    RuleEdit: RuleEdit
  },
  data () {
    return {
      treeData: [],
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
      menuImportFromFile: true,
      importDialogDisplay: false,
      importCodeExample: `  - alert: PrometheusTargetScrapingSlow
    expr: prometheus_target_interval_length_seconds{quantile="0.9"} > 60
    for: 5m
    labels:
      severity: warning
    annotations:
      summary: Prometheus target scraping slow (instance {{ $labels.instance }})
      description: "Prometheus is scraping exporters slowly\\n  VALUE = {{ $value }}\\n  LABELS = {{ $labels }}"

  - alert: PrometheusLargeScrape
    expr: increase(prometheus_target_scrapes_exceeded_sample_limit_total[10m]) > 10
    for: 5m
    labels:
      severity: warning
    annotations:
      summary: Prometheus large scrape (instance {{ $labels.instance }})
      description: "Prometheus has many scrapes that exceed the sample limit\\n  VALUE = {{ $value }}\\n  LABELS = {{ $labels }}"`,
      fileList: [],
      importPushing: false,
      error: '',
      showError: true,
      yamlRawObj: null,
      ruleGroupName: '',
      labelPath: [],
      expandedList: [],
      currentMode: '',
      pushing: false,
      btnTitleAppend: '',
      showEditArea: false
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
          this.treeData = [...r.data]
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
        this.showEditArea = false
        return false
      }
      this.$refs.ruleEditRef.setFormEnable()
      if (data.label === '[must rename me]') {
        this.showEditArea = false
        return false
      }
      this.showEditArea = true
      const queryInfo = {
        id: data.id,
        label: data.label,
        level: data.level,
        is_new: data.is_new,
        parent: data.parent
      }
      this.$refs.ruleEditRef.doGetRuleDetail(queryInfo)
    },
    handleNodeRightClick (event, data, node, nodeComp) {
      this.handleNodeClick(data)
    },
    append (data) {
      if (!data) {
        data = this.menuData
      }
      let newChild = {
        id: menuid++,
        level: data.level + 1,
        label: '[must rename me]',
        children: [],
        is_new: true,
        parent: data.id,
        path: [...data.path]
      }
      newChild.path.push(newChild.label)
      const newPathStr = newChild.path.join('_')
      newChild.code = md5(`${newChild.level}_${newChild.id}_${newPathStr}`)

      if (!data.children) {
        data.children = []
      }
      data.children.push(newChild);
      this.treeData = [...this.treeData]
    },
    disabled (data) {
      if (!data) {
        data = this.menuData
      }
      const disableData = {
        enabled: false,
        id: data.id,
        level: data.level
      }
      this.$confirm('是否确定禁用此分支以下的所有规则？', '确认信息', {
        distinguishCancelAndClose: true,
        confirmButtonText: '确定',
        cancelButtonText: '放弃'
      }).then(() => {
        disableTreeNode(disableData).then(r => {
          this.$notify({
            title: '成功',
            message: '禁用成功！',
            type: 'success'
          })
        })
      }).catch(e => console.log(e))
    },
    enabled (data) {
      if (!data) {
        data = this.menuData
      }
      const disableData = {
        enabled: true,
        id: data.id,
        level: data.level
      }
      this.$confirm('是否确定启用此分支以下的所有规则？', '确认信息', {
        distinguishCancelAndClose: true,
        confirmButtonText: '确定',
        cancelButtonText: '放弃'
      }).then(() => {
        disableTreeNode(disableData).then(r => {
          this.$notify({
            title: '成功',
            message: '启用成功！',
            type: 'success'
          })
        })
      }).catch(e => console.log(e))
    },
    remove (node, data) {
      if (!node) {
        node = this.menuNode
      }
      if (!data) {
        data = this.menuData
      }
      let remindTxt = ''
      if (data.level === 1) {
        remindTxt = '将会清空所有监控规则，包括文件、监控组、规则。是否确定删除？'
      } else if (data.level === 2) {
        remindTxt = `删除文件为: ${data.label}，及此文件内的所有监控规则。是否确定删除？`
      } else if (data.level === 3) {
        remindTxt = `删除监控组为: ${data.label}，及此组内的所有监控规则。是否确定删除？`
      } else if (data.level === 4) {
        remindTxt = `删除监控规则项为: ${data.label}。是否确定删除？`
      } else {
        return false
      }
      this.$confirm(remindTxt, '确认信息', {
        distinguishCancelAndClose: true,
        confirmButtonText: '确定',
        cancelButtonText: '放弃'
      }).then(() => {
        // console.log('data and node is: ', data, node)
        // return false
        if (data.is_new) {
          const parent = node.parent;
          const children = parent.data.children || parent.data;
          // const index = children.findIndex(d => d.id === data.id);
          const index = children.findIndex(d => d.code === data.code);
          children.splice(index, 1);
          this.treeData = [...this.treeData]
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
            // const index = children.findIndex(d => d.id === data.id);
            const index = children.findIndex(d => d.code === data.code);
            if (children[index].level !== 1) {
              children.splice(index, 1);
            } else {
              children[index].children = []
            }
            this.treeData = [...this.treeData]
          }
        ).catch(
          e => { console.log(e) }
        )
      }).catch(() => { })
    },
    edit (node, data) {
      if (this.currentMode === 'rename') {
        this.$notify({
          title: '错误',
          message: '请先完成节点的重新命名！',
          type: 'error'
        })
        return false
      }
      if (!node) {
        node = this.menuNode
      }
      if (!data) {
        data = this.menuData
      }
      data.display = true
      this.currentMode = 'rename'
      this.titleFromShowMe = data.label
      //   this.treeData = [...this.treeData]
    },
    importRules () {
      this.ruleGroupName = this.menuData.label
      this.importDialogDisplay = true
    },
    handleDialogClose () {
      this.importDialogDisplay = false
    },
    doImportCancel () {
      this.importDialogDisplay = false
    },
    doImportSubmit () {
      this.importDialogDisplay = false
    },
    importBefore (file) {
      console.log('before-upload importBefore')
      const isLt10M = file.size / 1024 / 1024 < 10
      if (!isLt10M) {
        this.error = '导入的文件不能超过10m'
        return false
      }
      return true
    },
    // importFileChange (file) {
    //   //清除文件对象
    //   this.$refs.upload.clearFiles()
    //   // 取出上传文件的对象，在其它地方也可以使用
    //   this.yamlRawObj = file.raw
    //   // 重新手动赋值filstList， 免得自定义上传成功了, 
    //   // 而fileList并没有动态改变， 这样每次都是上传一个对象
    //   this.fileList = [{ name: file.name, url: file.url }]
    // },
    doSubmitUpload () {
      // console.log('this.$refs.upload.submit()', 'this.$refs.upload.submit()')
      // console.log('ref =>', this.$refs.upload)
      this.$refs.uploadYamlRef.submit();
      // this.$refs.uploadYamlRef.onSuccess()
      // this.$refs.uploadYamlRef.onError()
    },
    submitUpload (param) {
      console.log('submitUpload', 'submitUpload')
      this.importPushing = true
      const formData = new FormData()
      formData.append('file', param.file)
      uploadYamlFile(formData).then(r => {
        // console.log('上传图片成功')
        this.error = '文件上传成功，并且已经被服务正确解析'
        // this.$refs.upload.onSuccess()
        // param.onSuccess()  // 上传成功的图片会显示绿色的对勾
        // 但是我们上传成功了图片， fileList 里面的值却没有改变，还好有on-change指令可以使用
        this.importPushing = false
      }).catch(e => {
        this.error = e.toString()
        // console.log('图片上传失败')
        // this.$refs.upload.onError()
        // param.onError()
        this.importPushing = false
      })
      // this.$refs.upload.submit();
    },
    expand (nodes) {
      const expandedList = []
      this.expandRecycle(null, expandedList)
      this.expandedList = expandedList
    },
    expandRecycle (nodes, expandedList) {
      if (!nodes) {
        nodes = this.treeData
      }
      if (nodes.length === 0) {
        return
      }
      nodes.forEach(x => {
        if (x.children && x.children.length !== 0) {
          this.expandRecycle(x.children, expandedList)
        } else {
          expandedList.push(x.code)
        }
      })
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
    },
    openMenu (e, data, node) {
      if (this.currentMode === 'rename') {
        this.$notify({
          title: '错误',
          message: '请先完成节点的重新命名！',
          type: 'error'
        })
        return false
      }
      if (data.level === 1) {
        this.menuDelDisabled = false
        this.menuRenameDisabled = true
        this.menuAddDisabled = false
        this.menuImportFromFile = true
        this.btnTitleAppend = '添加文件'
      } else if (data.level === 2) {
        if (data.label === '[must rename me]') {
          this.menuAddDisabled = true
        } else {
          this.menuAddDisabled = false
        }
        this.menuDelDisabled = false
        this.menuRenameDisabled = false
        this.menuImportFromFile = true
        this.btnTitleAppend = '添加组'
      } else if (data.level === 3) {
        if (data.label === '[must rename me]') {
          this.menuAddDisabled = true
        } else {
          this.menuAddDisabled = false
        }
        this.menuDelDisabled = false
        this.menuRenameDisabled = false
        this.menuImportFromFile = false
        this.btnTitleAppend = '添加规则'
      } else if (data.level === 4) {
        if (data.label === '[must rename me]') {
          this.menuRenameDisabled = false
        } else {
          this.menuRenameDisabled = true
        }
        this.menuDelDisabled = false
        this.menuAddDisabled = true
        this.menuImportFromFile = true
        this.btnTitleAppend = '添加规则'
      } else {
        this.$notify({
          title: '错误',
          message: '未选中节点，请重试',
          type: 'error'
        })
        return false
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
      this.menuData = data
      this.menuNode = node
      this.visible = true
    },
    receiveEsc () {
      this.titleFromShowMe = ''
      this.menuData.display = false
      this.currentMode = 'normal'
      this.treeData = [...this.treeData]
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
            this.currentMode = 'normal'
            // this.treeData = [...this.treeData]
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
            this.currentMode = 'normal'
            this.doGetTree()
          }
        ).catch(
          e => { console.log(e) }
        )
      }
    },
    closeMenu () {
      this.visible = false
    },
    doRulesPublish () {
      this.pushing = true
      rulesPublish().then(
        r => {
          this.$notify({
            title: '成功',
            message: '发布成功！',
            type: 'success'
          })
          this.pushing = false
        }
      ).catch(
        e => {
          console.log(e)
          this.pushing = false
        }
      )
    },
    nodeExpand (data) {
      this.expandedList.push(data.code); // 在节点展开是添加到默认展开数组
    },
    nodeCollapse (data) {
      this.expandedList.splice(this.expandedList.indexOf(data.code), 1); // 收起时删除数组里对应选项
    }
  }
};
</script>

// 以下为scss，记得去掉scoped，或者使用/deep/
<style scoped>
.monitor-board {
  display: flex;
  flex-direction: column;
  height: 85vh;
  width: 100%;
  /* border: 1px dashed burlywood; */
}
.monitor-board :deep() .el-card__header {
  padding: 5px 10px 5px 10px;
}
.node-content {
  width: 900px;
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
  z-index: 99;
}
.node-content-edit {
  z-index: 98;
  align-content: center;
  text-align: center;
}
.empty-board {
  width: 800px;
  height: 100%;
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
  align-content: center;
}
.rule-edit-box {
  min-height: 82vh;
}
.node-content-top {
  display: flex;
  justify-content: space-between;
  flex-direction: row;
  flex-wrap: nowrap;
  align-items: center;
}
.node-content-path {
  widows: 950px;
  max-width: 950px;
  overflow: hidden;
  text-overflow: ellipsis;
  overflow-block: hidden;
  white-space: nowrap;
}
.card-scrollbar {
  height: 76vh;
  width: 300px;
}
.tree-box {
  overflow: none;
  min-height: 150vh;
  min-width: 150%;
}
.card-scrollbar-right {
  height: 75vh;
  width: 100%;
}
.icon-action {
  font-size: 5px;
  margin-left: 5px;
  color: green;
}
.divider-class {
  margin: 0px 0px 0px 0px;
}

.dialog-action {
  align-content: right;
  text-align: right;
}

.dialog-area :deep() .el-dialog {
  margin: 20px auto !important;
}
.dialog-area :deep() .el-dialog__body {
  padding-top: 5px;
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

.contextmenu {
  margin: 0;
  background: rgb(241, 236, 238);
  z-index: 3000;
  position: absolute;
  list-style-type: none;
  padding: 0px 0;
  border-radius: 4px;
  font-size: 12px;
  font-weight: 400;
  color: #333;
  box-shadow: 2px 2px 3px 0 rgba(0, 0, 0, 0.3);
}
.contextmenu > li {
  margin: 0;
  padding: 0px 0px 0px 0px;
  cursor: pointer;
  align-content: left;
  text-align: left;
}
.contextmenu :deep() .el-button {
  padding: 5px 10px 5px 10px;
  width: 150px;
  text-align: left;
}
.custom-tree-node :deep() .el-badge__content {
  margin-top: 11px;
}
.contextmenu > li:hover {
  background: rgb(224, 212, 185);
}
</style>
