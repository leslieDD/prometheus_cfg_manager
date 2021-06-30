<template>
  <div ref="tinymcebox" class="app-content" style="height: 100%">
    <editor ref="tinymceObj" v-model="content" :height="height"></editor>
  </div>
</template>

<script>
import Editor from "@/components/ddedit/imcoder-tinymce.vue";

export default {
  components: { Editor },
  data () {
    return {
      content: "",
      height: 0
    };
  },
  watch: {
    content (val) {
      console.log(val, this.content)
    }
  },
  mounted () {
    let height = this.$refs.tinymcebox.clientHeight
    console.log('mounted', height)
    this.$refs.tinymceObj.mountTinymce()
  }
  // mounted () {
  //   this.height = this.$refs.tinymcebox.clientHeight
  //   // let width = this.$refs.tinymcebox.clientWidth
  // }
};
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