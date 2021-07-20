<template>
  <div class="main-box">
    <div class="main-content">
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
    <div class="yaml-viewer">
      <textarea ref="textareaRV" />
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
import 'codemirror/mode/yaml/yaml'

import { markRaw } from 'vue'

// import { ruleView } from '@/api/preview'
import * as js_yaml from "js-yaml"

window.jsyaml = js_yaml


import { loadRuleFileContent, loadAllRulesFile } from '@/api/publish'

export default {
  name: 'ymalViewer',
  data () {
    return {
      data: [],
      defaultProps: {
        children: 'children',
        label: 'label'
      },
      yamlViewer: false,
      value: ''
    }
  },
  methods: {
    doLoadAllRulesFiles () {
      loadAllRulesFile().then(
        r => {
          this.data = r.data
        }
      ).catch(
        e => { console.log(e) }
      )
    },
    doloadRuleFileContent (info) {
      loadRuleFileContent(info).then(
        r => {
          this.yamlViewer.setValue(r.data)
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
      this.doloadRuleFileContent(fInfo)
    },
    refresh () {
      this.yamlViewer.refresh()
    }
  },
  beforeUnmount () {
    // this.yamlViewer.destroy();
  },
  mounted () {
    this.yamlViewer = markRaw(CodeMirror.fromTextArea(this.$refs.textareaRV, {
      mode: 'text/x-yaml', // 语法model
      indentUnit: 2, // 缩进单位，默认2
      smartIndent: true, // 是否智能缩进
      // 显示行号
      styleActiveLine: true,
      lineNumbers: true,
      // 设置主题
      theme: "darcula",
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
    this.yamlViewer.refresh()
    // editor.setSize('width','height');
    // this.yamlViewer.setSize('835px', '83vh');
    this.yamlViewer.setValue(this.value)
    this.yamlViewer.on('change', (cm) => {
      this.$emit('changed', cm.getValue())
      this.$emit('input', cm.getValue())
    })
    this.doLoadAllRulesFiles()
  },
};
</script>

<style scoped>
.main-box {
  /* border: 1px solid burlywood; */
  border: 1px dashed burlywood;
  /* box-shadow: -10px 0px 10px red, 0px -10px 10px black, 10px 0px 10px green,
    0px 10px 10px blue; */
  /* box-shadow: burlywood 5px 5px 5px 5px; */
  display: flex;
  flex-direction: row;
  justify-content: flex-start;
  height: 85vh;
}
.main-content {
  width: 250px;
  height: 84vh;
}
/* .card-scrollbar {
  height: 84vh;
  width: 250px;
} */
.yaml-viewer {
  /* position: relative; */
  width: 820px;
}
.yaml-viewer :deep() .CodeMirror {
  height: 100%;
  min-height: 300px;
  width: 100%;
}
.yaml-viewer :deep() .CodeMirror-scroll {
  min-height: 300px;
}
.yaml-viewer :deep() .cm-s-rubyblue span.cm-string {
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
