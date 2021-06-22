<template>
  <div class="main-box">
    <div>
      <el-scrollbar class="card-scrollbar">
        <el-tree
          :data="data"
          :props="defaultProps"
          accordion
          default-expand-all
          :indent="0"
          @node-click="handleNodeClick"
          class="tree"
        >
        </el-tree>
      </el-scrollbar>
    </div>
    <div class="json-editor">
      <textarea ref="textareaGroup" />
    </div>
  </div>
</template>

<script>
// 引入CodeMirror和基础样式
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
// 需要依赖全局的jsonlint，不是很优雅
import "codemirror/addon/lint/json-lint.js";
// 引入jsonlint
import jsonlint from "jsonlint-mod";
window.jsonlint = jsonlint;

import { markRaw } from 'vue'


import { loadFileContent, loadAllFile } from '@/api/publish'

// window.jsjson = js_json

export default {
  name: 'JsonEditor',
  data () {
    return {
      data: [],
      defaultProps: {
        children: 'children',
        label: 'label'
      },
      jsonEditor: false,
      value: ''
    }
  },
  methods: {
    doLoadAllFiles () {
      loadAllFile().then(
        r => {
          this.data = r.data
        }
      ).catch(
        e => { console.log(e) }
      )
    },
    doloadFileContent (info) {
      loadFileContent(info).then(
        r => {
          this.jsonEditor.setValue(r.data)
        }
      ).catch(
        e => { console.log(e) }
      )
    },
    handleNodeClick (data, node, args3) {
      if (!data.path) {
        return false
      }
      const fInfo = {
        label: data.label,
        path: data.path
      }
      this.doloadFileContent(fInfo)
    }
  },
  beforeUnmount () {
    this.jsonEditor.destroy();
  },
  mounted () {
    this.jsonEditor = markRaw(CodeMirror.fromTextArea(this.$refs.textareaGroup, {
      mode: "application/json",
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
      lineWrapping: true,
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
      }
    }))
    this.jsonEditor.setValue('')
    this.jsonEditor.on('change', (cm) => {
      this.$emit('changed', cm.getValue())
      this.$emit('input', cm.getValue())
    })
  },
};
</script>

<style scoped>
.main-box {
  display: flex;
  flex-direction: row;
  height: 85vh;
}
.card-scrollbar {
  height: 84vh;
  width: 250px;
}
.json-editor {
  position: relative;
  width: 100%;
}
.json-editor :deep() .CodeMirror {
  height: 100%;
  min-height: 300px;
  width: 100%;
}
.json-editor :deep() .CodeMirror-scroll {
  min-height: 300px;
}
.json-editor :deep() .cm-s-rubyblue span.cm-string {
  color: #f08047;
}
.cm-s-monokai {
  height: 100%;
}
/* .tree-container {
  overflow: hidden;
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
</style>
