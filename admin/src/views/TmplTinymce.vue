<template>
  <div ref="tinymcebox" class="app-content" style="height: 100%">
    <editor ref="tinymceObj" :value="content" @showValue="showValue"></editor>
  </div>
</template>

<script>
import Editor from "@/components/ddedit/imcoder-tinymce.vue";

export default {
  components: { Editor },
  data () {
    return {
      content: "",
      height: 0,
      firstName: ""
    };
  },
  watch: {
    content (newVal, oldVal) {
      console.log(newVal, oldVal)
    }
  },
  mounted () {
    // let height = this.$refs.tinymcebox.clientHeight
    // console.log('mounted', height)
    // this.$refs.tinymceObj.mountTinymce()
    // setInterval(this.show, 2000)
  },
  methods: {
    showValue (v) {
      // this.content = new Date().toString()
      // console.log('content =>', v)
      this.firstName = v
      // this.content = new Date().toString()
    }
  }
}
</script>
<style lang="scss" scoped>
.Tinymce_box {
  display: flex;
}
.editor-content {
  margin-left: 30px;
  flex-grow: 1;
  border: 2px dashed #f1f1f1;
  padding: 0 20px;
}
h3 {
  color: #808080;
}
</style>