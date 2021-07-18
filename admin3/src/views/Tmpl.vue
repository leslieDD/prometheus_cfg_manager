<template>
  <div class="tmpl-box">
    <div class="control-butn">
      <el-button size="small" plain @click="viewGoStruct" icon="el-icon-view"
        >查看数组结构</el-button
      >
      <el-button
        size="small"
        type="danger"
        plain
        @click="doReload(true)"
        icon="el-icon-refresh"
        >重新加载</el-button
      >
      <el-button
        size="small"
        type="warning"
        plain
        @click="undo"
        icon="el-icon-folder-checked"
        >撤消</el-button
      >
      <el-button
        size="small"
        type="warning"
        plain
        @click="redo"
        icon="el-icon-refresh"
        >回退</el-button
      >
      <el-button
        size="small"
        type="primary"
        @click="doSave"
        icon="el-icon-folder-checked"
        >保存数据</el-button
      >
    </div>
    <div class="tmpl-editor">
      <textarea ref="textarea" />
    </div>
    <el-dialog
      title="查看数据结构"
      v-model="dialogVisible"
      width="30%"
      custom-class="custom-dialog-class"
    >
      <div>
        <ViewCode :value="structData"></ViewCode>
      </div>
      <template #footer>
        <span class="dialog-footer">
          <el-button size="small" type="primary" @click="enterDialog"
            >确定</el-button
          >
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import CodeMirror from "codemirror";
import "codemirror/lib/codemirror.css";
// JSON代码高亮需要由JavaScript插件支持
import "codemirror/mode/javascript/javascript.js";
// 选择IDEA主题样式，还有其他很多主题可选
// import "codemirror/theme/idea.css";
import "codemirror/theme/darcula.css"
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
// import 'codemirror/mode/yaml/yaml'
import 'codemirror/mode/django/django.js'
// import '@/assets/codemirror/mode/overlay.js'
import 'codemirror/addon/scroll/simplescrollbars.css'
import 'codemirror/addon/scroll/simplescrollbars.js'

import { markRaw } from 'vue'

// import { preview } from '@/api/preview'
import { getProTmpl, putProTmpl } from '@/api/tmpl.js'
import ViewCode from '@/components/ViewCodeMirror.vue'
import { getGoStruct } from '@/api/tmpl.js'
import * as js_yaml from "js-yaml"

// window.jsyaml = require('js-yaml') // 引入js-yaml为codemirror提高语法检查核心支持
window.jsyaml = js_yaml

export default {
  name: 'tmplEditor',
  // eslint-disable-next-line vue/require-prop-types
  //   props: ['value'],
  components: {
    ViewCode
  },
  data () {
    return {
      tmplEditor: false,
      value: '',
      tmplChangled: false,
      dialogVisible: false,
      structData: ''
    }
  },
  watch: {
    value (value) {
      const editorValue = this.tmplEditor.getValue()
      if (value !== editorValue) {
        this.tmplEditor.setValue(this.value)
      }
    }
  },
  beforeUnmount () {
    // this.tmplEditor.destroy();
  },
  mounted () {
    this.tmplEditor = markRaw(CodeMirror.fromTextArea(this.$refs.textarea, {
      mode: 'text/x-django', // 语法model
      indentUnit: 2, // 缩进单位，默认2
      smartIndent: true, // 是否智能缩进
      // 显示行号
      styleActiveLine: true,
      lineNumbers: true,
      // 设置主题
      // theme: "idea",
      theme: "darcula",
      // 绑定sublime快捷键
      keyMap: "sublime",
      // 开启代码折叠
      lineWrapping: false,   // 横向滚动条
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
      // scrollbarStyle: 'overlay'
      // scrollbarSytle: 'native',
      scrollbarStyle: 'simple'
    }))
    // editor.setSize('width','height');
    // this.tmplEditor.setSize('835px', '83vh');
    this.refresh()
    this.tmplEditor.setValue(this.value)

    this.tmplEditor.on('change', (cm) => {
      this.$emit('changed', cm.getValue())
      this.$emit('input', cm.getValue())
      this.tmplChangled = true
    })
    this.tmplEditor.on('update', (cm) => {
      this.tmplChangled = true
    })
    this.loadTmpl(false)
  },
  methods: {
    loadTmpl (useHand) {
      getProTmpl().then(
        r => {
          this.value = r.data.tmpl
          this.tmplEditor.setValue(this.value)
          if (useHand) {
            this.$notify({
              title: '成功',
              message: '加载成功！',
              type: 'success'
            })
          }
          this.tmplChangled = false
        }
      ).catch(
        e => { console.log(e) }
      )
    },
    doReload (useHand) {
      if (this.tmplChangled) {
        this.$confirm('文本已经更改，是否丢弃？', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          this.loadTmpl(true)
        }).catch(() => {
          this.$notify.info({
            title: '提示',
            message: '放弃重新加载！',
            showClose: false
          });
        });
      } else {
        this.loadTmpl(true)
      }
    },
    doSave () {
      const tmplVal = {
        "tmpl": this.tmplEditor.getValue()
      }
      putProTmpl(tmplVal).then(
        r => {
          this.$notify({
            title: '成功',
            message: '保存成功！',
            type: 'success'
          })
        }
      ).catch(
        e => { console.log(e) }
      )
    },
    undo () {
      this.tmplEditor.undo()
    },
    redo () {
      this.tmplEditor.redo()
    },
    destroy () {
      // this.tmplEditor.destroy()
    },
    refresh () {
      this.tmplEditor.refresh()
    },
    viewGoStruct () {
      getGoStruct().then(r => {
        this.structData = r.data
        this.dialogVisible = true
      }).catch(e => {
        console.log(e)
      })
    },
    enterDialog () {
      this.dialogVisible = false
    }
  },

}
</script>

<style scoped>
.tmpl-editor {
  position: relative;
  border: 1px solid burlywood;
}
.control-butn {
  display: flex;
  flex-direction: row;
  justify-content: flex-end;
  margin-bottom: 3px;
  margin-top: -10px;
}
.custom-class {
  padding-top: 5px;
  padding-bottom: 5px;
}
.tmpl-box :deep() .el-dialog__body {
  padding-top: 5px;
  padding-bottom: 5px;
}
</style>
