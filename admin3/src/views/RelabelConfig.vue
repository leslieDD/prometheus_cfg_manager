<template>
  <div class="relabel-config-box">
    <div class="main-content">
      <div>
        <div class="btn-action-area">
          <div>
            <span class="explain-words">
              说明：定义prometheus中的<el-tag size="mini" type="warning"
                >relabel_configs</el-tag
              >规则，可以在<el-tag size="mini">JOB组管理</el-tag>中使用</span
            >
          </div>
          <div class="do_action">
            <div style="padding-right: 15px">
              <el-button size="small" type="success" plain @click="doAdd()"
                >添加规则</el-button
              >
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
        </div>
        <div class="table-show">
          <el-table
            size="mini"
            highlight-current-row
            border
            :data="ReLabels"
            stripe
            :row-style="rowStyle"
            :cell-style="cellStyle"
          >
            <el-table-column label="序号" width="50px">
              <template v-slot="scope">
                {{ scope.$index + 1 }}
              </template>
            </el-table-column>
            <el-table-column label="名称" prop="name"> </el-table-column>
            <el-table-column label="最后更新时间" prop="update_at">
              <template v-slot="{ row }">
                <span>{{ parseTimeSelf(row.update_at) }}</span>
              </template>
            </el-table-column>
            <el-table-column label="操作" align="center" width="400px">
              <template v-slot="scope" align="center">
                <div class="actioneara">
                  <div>
                    <el-button
                      size="mini"
                      type="primary"
                      @click="doEdit(scope)"
                      plain
                      :disabled="editCodeButVisable[scope.row.id]"
                      >编辑名称</el-button
                    >
                  </div>
                  <div>
                    <el-button
                      size="mini"
                      type="primary"
                      @click="doEditReLablesCode(scope)"
                      :disabled="editCodeButVisable[scope.row.id]"
                      plain
                      >编辑规则</el-button
                    >
                  </div>
                  <div>
                    <el-button
                      size="mini"
                      type="warning"
                      :disabled="postCodeButVisable[scope.row.id]"
                      @click="doPostReLablesCode(scope)"
                      >提交规则</el-button
                    >
                  </div>
                  <div>
                    <el-button
                      size="mini"
                      type="info"
                      v-if="scope.row.enabled === true"
                      @click="invocate(scope)"
                      plain
                      :disabled="editCodeButVisable[scope.row.id]"
                      >禁用</el-button
                    >
                    <el-button
                      size="mini"
                      type="warning"
                      v-if="scope.row.enabled === false"
                      @click="invocate(scope)"
                      plain
                      :disabled="editCodeButVisable[scope.row.id]"
                      >启用</el-button
                    >
                  </div>
                  <div>
                    <el-popover
                      :visible="deleteVisible[scope.$index]"
                      placement="top"
                    >
                      <p>确定删除吗？</p>
                      <div style="text-align: right; margin: 0">
                        <el-button size="mini" type="text" @click="doNo(scope)"
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
          <div class="block">
            <el-pagination
              @size-change="handleSizeChange"
              @current-change="handleCurrentChange"
              :current-page="currentPage"
              :page-sizes="[10]"
              :page-size="pageSize"
              layout="total, sizes, prev, pager, next, jumper"
              :total="pageTotal"
            >
            </el-pagination>
          </div>
        </div>
      </div>
      <div class="yaml-relabel-edit">
        <textarea ref="textareaREdit" />
      </div>
    </div>
    <el-dialog
      :title="dialogTitle"
      v-model="dialogVisible"
      width="400px"
      modal
      :before-close="handleClose"
    >
      <span>
        <el-form
          label-position="right"
          :rules="rules"
          ref="editRelabelInfo"
          :model="editRelabelInfo"
          label-width="90px"
          size="small"
        >
          <el-form-item label="标签名：" prop="name">
            <el-input
              style="width: 230px"
              v-model="editRelabelInfo.name"
            ></el-input>
          </el-form-item>
          <el-form-item size="small">
            <el-button
              size="small"
              type="primary"
              @click="onSubmit('editRelabelInfo')"
              >{{ buttonTitle }}</el-button
            >
            <el-button
              size="small"
              type="info"
              @click="onCancel('editRelabelInfo')"
              >取消</el-button
            >
          </el-form-item>
        </el-form>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import CodeMirror from "codemirror";
import "codemirror/lib/codemirror.css";
// JSON代码高亮需要由JavaScript插件支持
import "codemirror/mode/javascript/javascript.js";
// 选择IDEA主题样式，还有其他很多主题可选
import "codemirror/theme/idea.css";
// 支持使用Sublime快捷键
import "codemirror/keymap/sublime.js";
// 搜索功能的依赖
import "codemirror/addon/dialog/dialog.js";
import "codemirror/addon/dialog/dialog.css";
// 支持搜索功能
import "codemirror/addon/search/search";
import "codemirror/addon/search/searchcursor.js";
// 支持各种代码折叠
import "codemirror/addon/fold/foldgutter.css";
import "codemirror/addon/fold/foldcode.js";
import "codemirror/addon/fold/foldgutter.js";
import "codemirror/addon/fold/brace-fold.js";
import "codemirror/addon/fold/comment-fold.js";
// 支持代码区域全屏功能
import "codemirror/addon/display/fullscreen.css";
import "codemirror/addon/display/fullscreen.js";
// 支持括号自动匹配
import "codemirror/addon/edit/matchbrackets.js";
import "codemirror/addon/edit/closebrackets.js";
// 支持代码自动补全
import "codemirror/addon/hint/show-hint.css";
import "codemirror/addon/hint/show-hint.js";
import "codemirror/addon/hint/anyword-hint.js";
// 行注释
import "codemirror/addon/comment/comment.js";
// JSON错误检查
import "codemirror/addon/lint/lint.css";
import "codemirror/addon/lint/lint.js";
import 'codemirror/addon/lint/yaml-lint'
import 'codemirror/theme/monokai.css'
import 'codemirror/mode/yaml/yaml'

import { markRaw } from 'vue'

import * as js_yaml from "js-yaml"
window.jsyaml = js_yaml

import {
  getReLabels,
  putReLabels,
  postReLabels,
  deleteReLabels,
  putReLabelsCode,
  enabledRelabelCfg,
  checkViewCodePriv,
} from '@/api/relabel.js'

export default {
  name: 'ymalRelabelEdit',
  data () {
    function validateStr (rule, value, callback) {
      if (value === '' || typeof value === 'undefined' || value == null) {
        callback(new Error('请输入正确名称'))
      } else {
        const reg = /\s+/
        if (reg.test(value)) {
          callback(new Error('名称中不允许有空字符'))
        } else {
          callback()
        }
      }
    }
    return {
      yamlRelabelEdit: false,
      value: '',
      ReLabels: [],
      editRelabelInfo: {
        'id': 0,
        'name': ''
      },
      postCodeButVisable: {},
      editCodeButVisable: {},
      defaultRuleID: 0,
      pageSize: 10,
      pageTotal: 0,
      currentPage: 1,
      searchContent: '',
      deleteVisible: {},
      dialogVisible: false,
      buttonTitle: '',
      dialogTitle: '',
      rules: {
        name: [
          { required: true, message: '请输入正确的名称', validator: validateStr, trigger: ['blur'] }
        ]
      }
    }
  },
  mounted () {
    this.yamlRelabelEdit = markRaw(CodeMirror.fromTextArea(this.$refs.textareaREdit, {
      mode: 'text/x-yaml', // 语法model
      indentUnit: 2, // 缩进单位，默认2
      smartIndent: true, // 是否智能缩进
      // 显示行号
      styleActiveLine: true,
      lineNumbers: true,
      // 设置主题
      theme: "idea",
      // 绑定sublime快捷键
      keyMap: "sublime",
      // 开启代码折叠
      lineWrapping: false,
      foldGutter: true,
      gutters: [
        "CodeMirror-linenumbers",
        "CodeMirror-foldgutter",
        "CodeMirror-lint-markers",
      ],
      // CodeMirror-lint-markers是实现语法报错功能
      lint: true,

      // 全屏模式
      fullScreen: false,

      // 括号匹配
      matchBrackets: true,
      autoCloseBrackets: true,
      extraKeys: {
        F11: (cm) => {
          cm.setOption("fullScreen", !cm.getOption("fullScreen"));
        },
        Esc: (cm) => {
          if (cm.getOption("fullScreen")) {
            cm.setOption("fullScreen", false);
          }
        },
      },
      scrollbarStyle: 'native'
    }))
    this.yamlRelabelEdit.refresh()
    this.yamlRelabelEdit.setSize('1120px', '50vh');
    this.yamlRelabelEdit.setValue(this.value)

    this.yamlRelabelEdit.on('change', (cm) => {
      this.$emit('changed', cm.getValue())
      this.$emit('input', cm.getValue())
    })
    this.doGetReLabels()
  },
  methods: {
    doAdd () {
      this.editRelabelInfo = {
        'id': 0,
        'name': ''
      }
      this.buttonTitle = '创建'
      this.dialogTitle = '增加重写标签'
      this.dialogVisible = true
    },
    doGetReLabels (getInfo) {
      if (!getInfo) {
        getInfo = {
          'pageNo': this.currentPage,
          'pageSize': this.pageSize,
          'search': this.searchContent
        }
      }
      getReLabels(getInfo).then(
        r => {
          this.postCodeButVisable = {}
          r.data.data.forEach(element => {
            this.postCodeButVisable[element.id] = true
            if (element.name === "空规则") {
              this.editCodeButVisable[element.id] = true
              // this.defaultRuleID = element.id
            } else {
              this.editCodeButVisable[element.id] = false
            }
          })
          this.ReLabels = r.data.data
          this.pageTotal = r.data.totalCount
          this.currentPage = r.data.pageNo
          this.pageSize = r.data.pageSize
        }
      ).catch(
        e => {
          console.log(e)
        }
      )
    },
    doEdit (data) {
      this.buttonTitle = '更新'
      this.dialogTitle = '编辑标签'
      this.editRelabelInfo.id = data.row.id
      this.editRelabelInfo.name = data.row.name
      this.dialogVisible = true
    },
    handleSizeChange (val) {
      let getInfo = {
        'pageNo': this.currentPage,
        'pageSize': val,
        'search': this.searchContent
      }
      this.doGetReLabels(getInfo)
    },
    handleCurrentChange (val) {
      let getInfo = {
        'pageNo': val,
        'pageSize': this.pageSize,
        'search': this.searchContent
      }
      this.doGetReLabels(getInfo)
    },
    handleClose (done) {
      this.dialogVisible = false
    },
    onSubmit (formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          let postData = {}
          postData['name'] = this.editRelabelInfo.name
          if (this.buttonTitle === '创建') {
            postReLabels(postData).then(
              r => {
                this.$notify({
                  title: '成功',
                  message: '创建成功！',
                  type: 'success'
                });
                this.doGetReLabels()
                this.dialogVisible = false
                this.$refs[formName].resetFields()
              }
            ).catch(
              e => { console.log(e) }
            )
          } else {
            postData['id'] = this.editRelabelInfo.id
            putReLabels(postData).then(
              r => {
                this.$notify({
                  title: '成功',
                  message: '更新成功！',
                  type: 'success'
                });
                this.doGetReLabels()
                this.dialogVisible = false
                this.$refs[formName].resetFields()
              }
            ).catch(
              e => { console.log(e) }
            )
          }
        } else {
          return false
        }
      })
    },
    onCancel (formName) {
      this.dialogVisible = false
      this.$refs[formName].resetFields()
    },
    onSearch () {
      this.doGetReLabels()
    },
    parseTimeSelf (t) {
      var time = new Date(Date.parse(t))
      return time.toLocaleDateString() + ' ' + time.toTimeString().split(' ')[0]
    },
    doYes (scope) {
      deleteReLabels({ id: scope.row.id }).then(
        r => {
          this.$notify({
            title: '成功',
            message: '删除成功！',
            type: 'success'
          });
          this.doGetReLabels()
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
      // if (row.rowIndex % 2 === 0) {
      //   rs['background'] = '#f2eada'
      // }
      return rs
    },
    cellStyle (column) {
      let cs = {
        'padding': '0'
      }
      return cs
    },
    refresh () {
      this.yamlRelabelEdit.refresh()
    },
    doEditReLablesCode (scape) {
      checkViewCodePriv().then(r => {
        Object.keys(this.postCodeButVisable).forEach(key => {
          if (key !== scape.row.id) {
            this.postCodeButVisable[key] = true
          }
        })
        this.yamlRelabelEdit.setValue(scape.row.code)
        this.postCodeButVisable[scape.row.id] = false
      }).catch(e => console.log(e))
    },
    doPostReLablesCode (scape) {
      const code = {
        id: scape.row.id,
        code: this.yamlRelabelEdit.getValue()
      }
      putReLabelsCode(code).then(
        r => {
          this.$notify({
            title: '成功',
            message: '提交成功！',
            type: 'success'
          });
          this.doGetReLabels()
        }
      ).catch(
        e => {
          console.log(e)
        }
      )
    },
    invocate (scope) {
      const newStatus = !this.ReLabels[scope.$index].enabled
      const rInfo = {
        id: scope.row.id,
        enabled: newStatus
      }
      enabledRelabelCfg(rInfo).then(r => {
        this.$notify({
          title: '成功',
          message: '更新状态成功！',
          type: 'success'
        });
        this.ReLabels[scope.$index].enabled = newStatus
        this.ReLabels = [...this.ReLabels]
      }).catch(e => console.log(e))
    }
  }
  // beforeUnmount () {
  //   this.yamlRelabelEdit.destroy();
  // }
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
  justify-content: flex-end;
  margin-bottom: 5px;
}
.actioneara {
  display: flex;
  flex-wrap: nowrap;
  justify-content: space-around;
}
.el-pagination {
  text-align: center;
}
.block {
  padding-top: 12px;
}
.yaml-relabel-edit {
  border: 1px solid burlywood;
  margin-top: 10px;
}
.cm-s-monokai {
  height: 100%;
}
.explain-words {
  font: 0.9em Arial, Tahoma, Verdana;
  color: #777;
}
.btn-action-area {
  display: flex;
  flex-wrap: nowrap;
  justify-content: space-between;
}
</style>
