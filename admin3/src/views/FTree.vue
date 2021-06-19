
<template>
  <div class="main-box">
    <el-scrollbar class="card-scrollbar">
      <el-tree
        :data="data"
        :props="defaultProps"
        accordion
        default-expand-all
        :indent="2"
        @node-click="handleNodeClick"
      >
      </el-tree>
    </el-scrollbar>
    <div class="yaml-editor">
      <textarea ref="textarea" />
    </div>
  </div>
</template>

<script>
import CodeMirror from 'codemirror'
import 'codemirror/addon/lint/lint.css'
import 'codemirror/lib/codemirror.css'
import 'codemirror/theme/monokai.css'
import 'codemirror/mode/yaml/yaml'
import 'codemirror/addon/lint/lint'
import 'codemirror/addon/lint/yaml-lint'
import * as js_yaml from "js-yaml"
import { loadFileContent, loadAllFile } from '@/api/publish'

// window.jsyaml = require('js-yaml') // 引入js-yaml为codemirror提高语法检查核心支持
window.jsyaml = js_yaml

export default {
  data () {
    return {
      data: [],
      defaultProps: {
        children: 'children',
        label: 'label'
      },
      yamlEditor: false,
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
          this.yamlEditor.setValue(r.data)
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
  mounted () {
    this.yamlEditor = CodeMirror.fromTextArea(this.$refs.textarea, {
      lineNumbers: true, // 显示行号
      mode: 'text/x-yaml', // 语法model
      gutters: ['CodeMirror-lint-markers'], // 语法检查器
      theme: 'monokai', // 编辑器主题
      lint: true, // 开启语法检查
      smartIndent: true,
      readOnly: false,
      scrollbarStyle: 'native'
    })

    this.yamlEditor.setValue(this.value)
    // this.yamlEditor.on('change', (cm) => {
    //   this.$emit('changed', cm.getValue())
    //   this.$emit('input', cm.getValue())
    // })
  },
};
</script>

<style>
.main-box {
  display: flex;
  height: 85vh;
}
.card-scrollbar {
  height: 400px;
  width: 250px;
}
.yaml-editor {
  position: relative;
  width: 100%;
}
.yaml-editor :deep() .CodeMirror {
  height: 100%;
  min-height: 300px;
  width: 100%;
}
.yaml-editor :deep() .CodeMirror-scroll {
  min-height: 300px;
}
.yaml-editor :deep() .cm-s-rubyblue span.cm-string {
  color: #f08047;
}
.cm-s-monokai {
  height: 100%;
}
</style>
