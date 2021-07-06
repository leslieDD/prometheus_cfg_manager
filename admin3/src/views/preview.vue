<template>
  <div class="yaml-editor-box">
    <div class="yaml-editor">
      <textarea ref="textarea" />
    </div>
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

import { preview } from '@/api/preview'
import * as js_yaml from "js-yaml"

// window.jsyaml = require('js-yaml') // 引入js-yaml为codemirror提高语法检查核心支持
window.jsyaml = js_yaml

export default {
  name: 'YamlEditor',
  // eslint-disable-next-line vue/require-prop-types
  //   props: ['value'],
  data () {
    return {
      yamlEditor: false,
      value: ''
    }
  },
  watch: {
    value (value) {
      const editorValue = this.yamlEditor.getValue()
      if (value !== editorValue) {
        this.yamlEditor.setValue(this.value)
      }
    }
  },
  beforeUnmount () {
    // this.yamlEditor.destroy();
  },
  mounted () {
    this.yamlEditor = markRaw(CodeMirror.fromTextArea(this.$refs.textarea, {
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
      },
      scrollbarStyle: 'native'
    }))
    this.refresh()
    this.yamlEditor.setValue(this.value)

    this.yamlEditor.on('change', (cm) => {
      this.$emit('changed', cm.getValue())
      this.$emit('input', cm.getValue())
    })
    this.loadYaml()
  },
  methods: {
    // getValue () {
    //   return this.yamlEditor.getValue()
    // },
    loadYaml () {
      preview().then(
        r => {
          this.value = r.data
          this.yamlEditor.setValue(r.data)
          //   this.$notify({
          //     title: '成功',
          //     message: '配置加载成功！',
          //     type: 'success'
          //   });
        }
      ).catch(
        e => { console.log(e) }
      )
    },
    destroy () {
      // this.yamlEditor.destroy()
    },
    refresh () {
      this.yamlEditor.refresh()
    }
  }
}
</script>

<style scoped>
.yaml-editor-box {
  border: 1px dashed burlywood;
}
.yaml-editor {
  height: 85vh;
  position: relative;
}
.yaml-editor :deep() .CodeMirror {
  /* height: auto; */
  height: 85vh;
  /* height: 100%; */
  min-height: 300px;
}
.yaml-editor :deep() .CodeMirror-scroll {
  min-height: 300px;
}
.yaml-editor :deep() .cm-s-rubyblue span.cm-string {
  color: #f08047;
}
</style>
