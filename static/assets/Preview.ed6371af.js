import{C as e}from"./lint.624d88ac.js";/* empty css                */import"./monokai.8bdeb9d9.js";import"./yaml.111f68d9.js";import{s as t}from"./request.c24bac87.js";import{j as a}from"./js-yaml.039310f8.js";import{y as r,p as s,j as l,o as i,c as o,a as d,k as m}from"./vendor.684f159e.js";import"./index.6f2c4da2.js";window.jsyaml=a;const n={name:"YamlEditor",data:()=>({yamlEditor:!1,value:""}),watch:{value(e){e!==this.yamlEditor.getValue()&&this.yamlEditor.setValue(this.value)}},beforeUnmount(){},mounted(){this.yamlEditor=r(e.fromTextArea(this.$refs.textarea,{mode:"text/x-yaml",indentUnit:2,smartIndent:!0,styleActiveLine:!0,lineNumbers:!0,theme:"darcula",keyMap:"sublime",lineWrapping:!0,foldGutter:!0,gutters:["CodeMirror-linenumbers","CodeMirror-foldgutter","CodeMirror-lint-markers"],lint:!0,fullScreen:!1,matchBrackets:!0,autoCloseBrackets:!0,extraKeys:{F11:e=>{e.setOption("fullScreen",!e.getOption("fullScreen"))},Esc:e=>{e.getOption("fullScreen")&&e.setOption("fullScreen",!1)}},scrollbarStyle:"native"})),this.refresh(),this.yamlEditor.setValue(this.value),this.yamlEditor.on("change",(e=>{this.$emit("changed",e.getValue()),this.$emit("input",e.getValue())})),this.loadYaml()},methods:{loadYaml(){t({url:"/v1/preview",method:"get"}).then((e=>{this.value=e.data,this.yamlEditor.setValue(e.data)})).catch((e=>{console.log(e)}))},destroy(){},refresh(){this.yamlEditor.refresh()}}},c=m();s("data-v-24de4176");const u={class:"yaml-editor-box"},h={class:"yaml-editor"},f={ref:"textarea"};l();const p=c(((e,t,a,r,s,l)=>(i(),o("div",u,[d("div",h,[d("textarea",f,null,512)])]))));n.render=p,n.__scopeId="data-v-24de4176";export default n;
