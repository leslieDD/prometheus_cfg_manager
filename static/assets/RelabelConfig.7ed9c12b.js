import{C as e}from"./lint.ec343b19.js";import"./monokai.66f7bbd2.js";import"./yaml.181bea6f.js";import{y as t,p as l,j as i,r as a,o as s,c as o,a as d,w as n,l as r,t as c,q as h,k as b}from"./vendor.185638a5.js";import{j as u}from"./js-yaml.039310f8.js";import{g as p,p as m,a as g,d as f,c as C,b as y,e as R}from"./relabel.c4ab4393.js";import"./request.f1ddaf6e.js";import"./index.755618d3.js";window.jsyaml=u;const V={name:"ymalRelabelEdit",data:()=>({yamlRelabelEdit:!1,value:"",ReLabels:[],editRelabelInfo:{id:0,name:""},postCodeButVisable:{},editCodeButVisable:{},defaultRuleID:0,pageSize:10,pageTotal:0,currentPage:1,searchContent:"",deleteVisible:{},dialogVisible:!1,buttonTitle:"",dialogTitle:"",rules:{name:[{required:!0,message:"请输入正确的名称",validator:function(e,t,l){if(""===t||void 0===t||null==t)l(new Error("请输入正确名称"));else{/\s+/.test(t)?l(new Error("名称中不允许有空字符")):l()}},trigger:["blur"]}]}}),mounted(){this.yamlRelabelEdit=t(e.fromTextArea(this.$refs.textareaREdit,{mode:"text/x-yaml",indentUnit:2,smartIndent:!0,styleActiveLine:!0,lineNumbers:!0,theme:"idea",keyMap:"sublime",lineWrapping:!1,foldGutter:!0,gutters:["CodeMirror-linenumbers","CodeMirror-foldgutter","CodeMirror-lint-markers"],lint:!0,fullScreen:!1,matchBrackets:!0,autoCloseBrackets:!0,extraKeys:{F11:e=>{e.setOption("fullScreen",!e.getOption("fullScreen"))},Esc:e=>{e.getOption("fullScreen")&&e.setOption("fullScreen",!1)}},scrollbarStyle:"native"})),this.yamlRelabelEdit.refresh(),this.yamlRelabelEdit.setSize("1120px","50vh"),this.yamlRelabelEdit.setValue(this.value),this.yamlRelabelEdit.on("change",(e=>{this.$emit("changed",e.getValue()),this.$emit("input",e.getValue())})),this.doGetReLabels()},methods:{doAdd(){this.editRelabelInfo={id:0,name:""},this.buttonTitle="创建",this.dialogTitle="增加重写标签",this.dialogVisible=!0},doGetReLabels(e){e||(e={pageNo:this.currentPage,pageSize:this.pageSize,search:this.searchContent}),p(e).then((e=>{this.postCodeButVisable={},e.data.data.forEach((e=>{this.postCodeButVisable[e.id]=!0,"空规则"===e.name?this.editCodeButVisable[e.id]=!0:this.editCodeButVisable[e.id]=!1})),this.ReLabels=e.data.data,this.pageTotal=e.data.totalCount,this.currentPage=e.data.pageNo,this.pageSize=e.data.pageSize})).catch((e=>{console.log(e)}))},doEdit(e){this.buttonTitle="更新",this.dialogTitle="编辑标签",this.editRelabelInfo.id=e.row.id,this.editRelabelInfo.name=e.row.name,this.dialogVisible=!0},handleSizeChange(e){let t={pageNo:this.currentPage,pageSize:e,search:this.searchContent};this.doGetReLabels(t)},handleCurrentChange(e){let t={pageNo:e,pageSize:this.pageSize,search:this.searchContent};this.doGetReLabels(t)},handleClose(e){this.dialogVisible=!1},onSubmit(e){this.$refs[e].validate((t=>{if(!t)return!1;{let t={};t.name=this.editRelabelInfo.name,"创建"===this.buttonTitle?m(t).then((t=>{this.$notify({title:"成功",message:"创建成功！",type:"success"}),this.doGetReLabels(),this.dialogVisible=!1,this.$refs[e].resetFields()})).catch((e=>{console.log(e)})):(t.id=this.editRelabelInfo.id,g(t).then((t=>{this.$notify({title:"成功",message:"更新成功！",type:"success"}),this.doGetReLabels(),this.dialogVisible=!1,this.$refs[e].resetFields()})).catch((e=>{console.log(e)})))}}))},onCancel(e){this.dialogVisible=!1,this.$refs[e].resetFields()},onSearch(){this.doGetReLabels()},parseTimeSelf(e){var t=new Date(Date.parse(e));return t.toLocaleDateString()+" "+t.toTimeString().split(" ")[0]},doYes(e){f({id:e.row.id}).then((e=>{this.$notify({title:"成功",message:"删除成功！",type:"success"}),this.doGetReLabels()})).catch((e=>{console.log(e)})),this.deleteVisible[e.$index]=!1},doNo(e){this.deleteVisible[e.$index]=!1},doDelete(e){this.deleteVisible[e.$index]=!0},rowStyle:e=>({padding:"0"}),cellStyle:e=>({padding:"0"}),refresh(){this.yamlRelabelEdit.refresh()},doEditReLablesCode(e){C().then((t=>{Object.keys(this.postCodeButVisable).forEach((t=>{t!==e.row.id&&(this.postCodeButVisable[t]=!0)})),this.yamlRelabelEdit.setValue(e.row.code),this.postCodeButVisable[e.row.id]=!1})).catch((e=>console.log(e)))},doPostReLablesCode(e){const t={id:e.row.id,code:this.yamlRelabelEdit.getValue()};y(t).then((e=>{this.$notify({title:"成功",message:"提交成功！",type:"success"}),this.doGetReLabels()})).catch((e=>{console.log(e)}))},invocate(e){const t=!this.ReLabels[e.$index].enabled,l={id:e.row.id,enabled:t};R(l).then((l=>{this.$notify({title:"成功",message:"更新状态成功！",type:"success"}),this.ReLabels[e.$index].enabled=t,this.ReLabels=[...this.ReLabels]})).catch((e=>console.log(e)))}}},w=b();l("data-v-2a87865a");const v={class:"relabel-config-box"},z={class:"main-content"},S={class:"btn-action-area"},k={class:"explain-words"},_=r(" 说明：定义prometheus中的"),x=r("relabel_configs"),L=r("规则，可以在"),E=r("JOB组管理"),$=r("中使用"),B={class:"do_action"},T={style:{"padding-right":"15px"}},I=r("添加规则"),j={class:"table-show"},G={class:"actioneara"},N=r("编辑名称"),P=r("编辑规则"),D=r("提交规则"),O=r("禁用"),A=r("启用"),F=d("p",null,"确定删除吗？",-1),M={style:{"text-align":"right",margin:"0"}},U=r("取消"),q=r("确定"),K=r("删除"),Y={class:"block"},J={class:"yaml-relabel-edit"},W={ref:"textareaREdit"},H=r("取消");i();const Q=w(((e,t,l,i,b,u)=>{const p=a("el-tag"),m=a("el-button"),g=a("el-input"),f=a("el-table-column"),C=a("el-popover"),y=a("el-table"),R=a("el-pagination"),V=a("el-form-item"),Q=a("el-form"),X=a("el-dialog");return s(),o("div",v,[d("div",z,[d("div",null,[d("div",S,[d("div",null,[d("span",k,[_,d(p,{size:"mini",type:"warning"},{default:w((()=>[x])),_:1}),L,d(p,{size:"mini"},{default:w((()=>[E])),_:1}),$])]),d("div",B,[d("div",T,[d(m,{size:"small",type:"success",plain:"",onClick:t[1]||(t[1]=e=>u.doAdd())},{default:w((()=>[I])),_:1})]),d("div",null,[d(g,{size:"small",placeholder:"请输入内容",onKeyup:t[3]||(t[3]=n((e=>u.onSearch()),["enter"])),modelValue:b.searchContent,"onUpdate:modelValue":t[4]||(t[4]=e=>b.searchContent=e)},{append:w((()=>[d(m,{size:"small",onClick:t[2]||(t[2]=e=>u.onSearch()),icon:"el-icon-search"})])),_:1},8,["modelValue"])])])]),d("div",j,[d(y,{size:"mini","highlight-current-row":"",border:"",data:b.ReLabels,stripe:"","row-style":u.rowStyle,"cell-style":u.cellStyle},{default:w((()=>[d(f,{label:"序号",width:"50px"},{default:w((e=>[r(c(e.$index+1),1)])),_:1}),d(f,{label:"名称",prop:"name"}),d(f,{label:"最新后新账号",prop:"update_by"}),d(f,{label:"最后更新时间",prop:"update_at"},{default:w((({row:e})=>[d("span",null,c(u.parseTimeSelf(e.update_at)),1)])),_:1}),d(f,{label:"操作",align:"center",width:"400px"},{default:w((e=>[d("div",G,[d("div",null,[d(m,{size:"mini",type:"primary",onClick:t=>u.doEdit(e),plain:"",disabled:b.editCodeButVisable[e.row.id]},{default:w((()=>[N])),_:2},1032,["onClick","disabled"])]),d("div",null,[d(m,{size:"mini",type:"primary",onClick:t=>u.doEditReLablesCode(e),disabled:b.editCodeButVisable[e.row.id],plain:""},{default:w((()=>[P])),_:2},1032,["onClick","disabled"])]),d("div",null,[d(m,{size:"mini",type:"warning",disabled:b.postCodeButVisable[e.row.id],onClick:t=>u.doPostReLablesCode(e)},{default:w((()=>[D])),_:2},1032,["disabled","onClick"])]),d("div",null,[!0===e.row.enabled?(s(),o(m,{key:0,size:"mini",type:"info",onClick:t=>u.invocate(e),plain:"",disabled:b.editCodeButVisable[e.row.id]},{default:w((()=>[O])),_:2},1032,["onClick","disabled"])):h("",!0),!1===e.row.enabled?(s(),o(m,{key:1,size:"mini",type:"warning",onClick:t=>u.invocate(e),plain:"",disabled:b.editCodeButVisable[e.row.id]},{default:w((()=>[A])),_:2},1032,["onClick","disabled"])):h("",!0)]),d("div",null,[d(C,{visible:b.deleteVisible[e.$index],placement:"top"},{reference:w((()=>[d(m,{size:"mini",type:"danger",plain:"",onClick:t=>u.doDelete(e),disabled:b.editCodeButVisable[e.row.id]},{default:w((()=>[K])),_:2},1032,["onClick","disabled"])])),default:w((()=>[F,d("div",M,[d(m,{size:"mini",type:"text",onClick:t=>u.doNo(e)},{default:w((()=>[U])),_:2},1032,["onClick"]),d(m,{type:"primary",size:"mini",onClick:t=>u.doYes(e)},{default:w((()=>[q])),_:2},1032,["onClick"])])])),_:2},1032,["visible"])])])])),_:1})])),_:1},8,["data","row-style","cell-style"]),d("div",Y,[d(R,{onSizeChange:u.handleSizeChange,onCurrentChange:u.handleCurrentChange,"current-page":b.currentPage,"page-sizes":[10],"page-size":b.pageSize,layout:"total, sizes, prev, pager, next, jumper",total:b.pageTotal},null,8,["onSizeChange","onCurrentChange","current-page","page-size","total"])])])]),d("div",J,[d("textarea",W,null,512)])]),d(X,{title:b.dialogTitle,modelValue:b.dialogVisible,"onUpdate:modelValue":t[8]||(t[8]=e=>b.dialogVisible=e),width:"400px",modal:"","before-close":u.handleClose},{default:w((()=>[d("span",null,[d(Q,{"label-position":"right",rules:b.rules,ref:"editRelabelInfo",model:b.editRelabelInfo,"label-width":"90px",size:"small"},{default:w((()=>[d(V,{label:"标签名：",prop:"name"},{default:w((()=>[d(g,{style:{width:"230px"},modelValue:b.editRelabelInfo.name,"onUpdate:modelValue":t[5]||(t[5]=e=>b.editRelabelInfo.name=e)},null,8,["modelValue"])])),_:1}),d(V,{size:"small"},{default:w((()=>[d(m,{size:"small",type:"primary",onClick:t[6]||(t[6]=e=>u.onSubmit("editRelabelInfo"))},{default:w((()=>[r(c(b.buttonTitle),1)])),_:1}),d(m,{size:"small",type:"info",onClick:t[7]||(t[7]=e=>u.onCancel("editRelabelInfo"))},{default:w((()=>[H])),_:1})])),_:1})])),_:1},8,["rules","model"])])])),_:1},8,["title","modelValue","before-close"])])}));V.render=Q,V.__scopeId="data-v-2a87865a";export default V;
