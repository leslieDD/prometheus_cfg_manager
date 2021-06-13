<template>
  <div class="yaml-editor">
    <textarea ref="textarea" />
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
import { preview } from '@/api/preview'

window.jsyaml = require('js-yaml') // 引入js-yaml为codemirror提高语法检查核心支持

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
  mounted () {
    this.yamlEditor = CodeMirror.fromTextArea(this.$refs.textarea, {
      lineNumbers: true, // 显示行号
      mode: 'text/x-yaml', // 语法model
      gutters: ['CodeMirror-lint-markers'], // 语法检查器
      theme: 'monokai', // 编辑器主题
      lint: true, // 开启语法检查
      smartIndent: true,
      readOnly: true,
      scrollbarStyle: 'native'
    })

    this.yamlEditor.setValue(this.value)
    // this.yamlEditor.on('change', (cm) => {
    //   this.$emit('changed', cm.getValue())
    //   this.$emit('input', cm.getValue())
    // })
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
          this.$message({
            showClose: true,
            message: '配置加载成功！',
            type: 'success'
          })
        }
      ).catch(
        e => { console.log(e) }
      )
    }
  }
}
</script>

<style scoped>
.yaml-editor{
  height: 85vh;
  position: relative;
}
.yaml-editor >>> .CodeMirror {
  /* height: auto; */
  height: 85vh;
  /* height: 100%; */
  min-height: 300px;
}
.yaml-editor >>> .CodeMirror-scroll{
  min-height: 300px;
}
.yaml-editor >>> .cm-s-rubyblue span.cm-string {
  color: #F08047;
}
</style>
