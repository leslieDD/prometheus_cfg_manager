<template>
  <!-- 富文本 -->
  <div class="tinymce-div" ref="tinymce" style="height: 100%">
    <editor
      v-model="content"
      :init="init"
      :disabled="disabled"
      @click="onClick"
    ></editor>
  </div>
</template>


<script>
import tinymce from "tinymce/tinymce";
import Editor from "@tinymce/tinymce-vue";
// import Editor from "tinymce/tinymce";
import "tinymce/icons/default/icons";
import "tinymce/themes/silver";
import "tinymce/plugins/image";
import "tinymce/plugins/media";
import "tinymce/plugins/table";
import "tinymce/plugins/lists";
import "tinymce/plugins/contextmenu";
import "tinymce/plugins/wordcount";
import "tinymce/plugins/colorpicker";
import "tinymce/plugins/textcolor";
import "tinymce/plugins/preview";
import "tinymce/plugins/code";
import "tinymce/plugins/link";
import "tinymce/plugins/advlist";
import "tinymce/plugins/codesample";
import "tinymce/plugins/hr";
import "tinymce/plugins/fullscreen";
import "tinymce/plugins/textpattern";
import "tinymce/plugins/searchreplace";
import "tinymce/plugins/autolink";
import "tinymce/plugins/directionality";
import "tinymce/plugins/visualblocks";
import "tinymce/plugins/visualchars";
import "tinymce/plugins/template";
import "tinymce/plugins/charmap";
import "tinymce/plugins/nonbreaking";
import "tinymce/plugins/insertdatetime";
import "tinymce/plugins/imagetools";
import "tinymce/plugins/autosave";
import "tinymce/plugins/save"
//编辑器高度自适应,注：plugins里引入此插件时，Init里设置的height将失效
// import "tinymce/plugins/autoresize";
// 扩展插件
import "@/plugins/tinymce/lineheight/plugin.js";
import "@/plugins/tinymce/bdmap/plugin.js";


export default {
  components: {
    Editor
  },
  props: {
    value: {
      type: String,
      default: ""
    },
    height: {
      type: Number,
      default: 700
    },
    disabled: {
      type: Boolean,
      default: false
    },
    plugins: {
      type: [String, Array],
      default:
        "preview searchreplace autolink directionality visualblocks visualchars fullscreen image link media template code codesample table charmap hr nonbreaking insertdatetime advlist lists wordcount imagetools textpattern autosave bdmap lineheight save"
    },
    toolbar: {
      type: [String, Array],
      default:
        "save | code undo redo restoredraft | cut copy paste pastetext | forecolor backcolor bold italic underline strikethrough link codesample | alignleft aligncenter alignright alignjustify outdent indent lineheight formatpainter | \
    styleselect formatselect fontselect fontsizeselect | bullist numlist | blockquote subscript superscript removeformat | \
    table image media charmap hr pagebreak insertdatetime | bdmap fullscreen preview"
    }
  },
  data () {
    return {
      //初始化配置
      firstName: "diandian",
      init: {
        language_url: "/static/tinymce/langs/zh_CN.js",
        language: "zh_CN",
        skin_url: "/static/tinymce/skins/ui/oxide",
        height: "85vh",
        min_height: "85vh",
        max_height: "85vh",
        width: 830,
        max_width: 830,
        min_width: 830,
        toolbar_mode: "wrap",
        plugins: this.plugins,
        toolbar: this.toolbar,
        forced_root_block: 'p',
        content_style: "p {margin: 0px 0; font-famliy: Consolas,Courier New,Cascadia Code}",
        fontsize_formats: "12px 14px 16px 18px 24px 36px 48px 56px 72px",
        font_formats:
          "Courier=Cascadia Code,Consolas, Courier New, monospace;微软雅黑=Microsoft YaHei,Helvetica Neue,PingFang SC,sans-serif;苹果苹方=PingFang SC,Microsoft YaHei,sans-serif;宋体=simsun,serif;仿宋体=FangSong,serif;黑体=SimHei,sans-serif;Arial=arial,helvetica,sans-serif;Arial Black=arial black,avant garde;Book Antiqua=book antiqua,palatino;",
        branding: false,
        theme_advanced_buttons3_add: "save",
        save_enablewhendirty: false,
        style_formats: [
          {
            title: '首行缩进',
            block: 'p',
            styles: { 'text-indent': '2em' }
          },
          {
            title: '行高',
            items: [
              { title: '1', styles: { 'line-height': '1' }, inline: 'span' },
              { title: '1.5', styles: { 'line-height': '1.5' }, inline: 'span' },
              { title: '2', styles: { 'line-height': '2' }, inline: 'span' },
              { title: '2.5', styles: { 'line-height': '2.5' }, inline: 'span' },
              { title: '3', styles: { 'line-height': '3' }, inline: 'span' }
            ]
          }
        ],
        save_onsavecallback: function (v, v2, v3) {
          // tinymce.triggerSave();
          // console.log(tinymce)
          var activeEditor = tinymce.activeEditor;
          var editBody = activeEditor.getBody();
          activeEditor.selection.select(editBody);
          var text = activeEditor.selection.getContent({ 'format': 'text' });
          console.log(text)
        },
        //此处为图片上传处理函数，这个直接用了base64的图片形式上传图片，
        //如需ajax上传可参考https://www.tiny.cloud/docs/configure/file-image-upload/#images_upload_handler
        images_upload_handler: (blobInfo, success, failure) => {
          const img = "data:image/jpeg;base64," + blobInfo.base64();
          success(img);
        }
      },
      content: this.value
    };
  },
  mounted () {
    tinymce.init({});

  },
  methods: {
    onClick () {
    },
    saveContent (value, value2, value3) {
      console.log(value, value2, value3)
    }
  },
  watch: {
    value (newValue) {
      console.log('value', newValue)
      this.content = newValue;
    },
    content (newValue) {
      this.$emit("input", newValue);
      this.$emit('showValue', newValue)
    },
    height (newValue) {
      this.height = newValue
    }
  }
};
</script>
<style scoped>
.tinymce-div {
  height: 80vh;
}
/* .tinymce-div :deep() .tox-tinymce {
  height: 85vh !important;
} */
/* .tinymce-div :deep() .tox-edit-area__iframe {
  height: 70vh;
  min-height: 70vh !important;
  max-height: 70vh !important;
  overflow: scroll;
} */
</style>