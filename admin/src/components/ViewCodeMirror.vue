<template>
  <div class="tmpl-editor">
    <textarea ref="textareaGo" />
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
// import 'codemirror/addon/lint/go-lint'
import 'codemirror/theme/monokai.css'
// import 'codemirror/mode/yaml/yaml'
// import 'codemirror/mode/django/django.js'
// import '@/assets/codemirror/mode/overlay.js'
import 'codemirror/addon/scroll/simplescrollbars.css'
import 'codemirror/addon/scroll/simplescrollbars.js'

import { markRaw } from 'vue'
import * as js_yaml from "js-yaml"

// window.jsyaml = require('js-yaml') // 引入js-yaml为codemirror提高语法检查核心支持
window.jsyaml = js_yaml

export default {
  name: 'ViewCodeMirror',
  // eslint-disable-next-line vue/require-prop-types
  //   props: ['value'],
  props: {
    value: {
      type: String,
      default: ""
    }
  },
  data () {
    return {
      structView: false
    }
  },
  watch: {
    value (value) {
      const editorValue = this.structView.getValue()
      if (value !== editorValue) {
        this.refresh()
        this.structView.setValue(this.value)
      }
    }
  },
  beforeUnmount () {
    // this.structView.destroy();
  },
  mounted () {
    this.structView = markRaw(CodeMirror.fromTextArea(this.$refs.textareaGo, {
      mode: 'text/x-go', // 语法model
      indentUnit: 2, // 缩进单位，默认2
      smartIndent: true, // 是否智能缩进
      // 显示行号
      styleActiveLine: true,
      lineNumbers: true,
      readOnly: true,
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
    this.structView.setSize('100%', '250px');
    this.refresh()
    this.structView.setValue(this.value)

    // this.structView.on('change', (cm) => {
    //   this.$emit('changed', cm.getValue())
    //   this.$emit('input', cm.getValue())
    // })
  },
  methods: {
    refresh () {
      this.structView.refresh()
    }
  },

}
</script>

<style scoped>
.tmpl-editor {
  position: relative;
  border: 1px solid burlywood;
}
</style>
