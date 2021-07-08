var e=Object.defineProperty,t=Object.defineProperties,l=Object.getOwnPropertyDescriptors,a=Object.getOwnPropertySymbols,i=Object.prototype.hasOwnProperty,s=Object.prototype.propertyIsEnumerable,o=(t,l,a)=>l in t?e(t,l,{enumerable:!0,configurable:!0,writable:!0,value:a}):t[l]=a,d=(e,t)=>{for(var l in t||(t={}))i.call(t,l)&&o(e,l,t[l]);if(a)for(var l of a(t))s.call(t,l)&&o(e,l,t[l]);return e},n=(e,a)=>t(e,l(a));import{p as r,a as u,g as p,b as h,c as b,d as c,e as g,f,h as m,i as y,j as C,k as L,l as w}from"./labelsJob.cd5c3520.js";import{g as G}from"./monitor.f7261b60.js";import"./request.0c690da8.js";import{p as S,b as V,r as _,o as z,c as k,a as v,j as P,t as I,h as j,F as D,i as x,k as T,w as $}from"./vendor.61ee7ef1.js";const N={name:"JobsLabels",data:()=>({subGroups:[],jobInfo:{},relabelList:[],pageSize:15,pageTotal:0,currentPage:1,searchContent:"",deleteVisible:{},deleteLabelsVisible:{},dialogVisible:!1,buttonTitle:"",buttonType:"primary",dialogTitle:"",subGroupData:{},editIPVisible:!1,editIPValue:[],editIPValueMap:{},editIPData:[],editIPDataMap:{},editObject:"",jobGroupIDCurrent:0,editLabelsVisible:!1,defaultLabels:[],labelsData:[],glPageTotal:0,glPageSize:10,glCurrentPage:1,addNewGroupLabels:{},glSearchContent:"",alertTitle:"当前状态：可以添加新标签",labelsBtnTitle:"添加",ipsAndLabels:{},ipsAndLabelsDone:{},needtoUpdate:!1,rules:{name:[{required:!0,message:"请输入正确的分组名称",validator:function(e,t,l){if(""===t||void 0===t||null==t)l(new Error("请输入正确名称"));else{/\s+/.test(t)?l(new Error("名称中不允许有空字符")):l()}},trigger:["blur"]}]},glRules:{key:[{required:!0,message:"请选择或者输入分组名称",trigger:["blur","change"]}]}}),created(){},mounted(){this.jobInfo=this.$route.params,this.jobInfo.id=parseInt(this.jobInfo.id,10),this.doGetSubGroup()},methods:{pushNewIPList(){let e=[];this.editIPValue.forEach((t=>{const l=this.editIPValueMap[t];let a={machines_id:t,ipaddr:this.editIPDataMap[t].ipaddr};a.id=l?l.id:0,e.push(a)})),r(this.jobGroupIDCurrent,e).then((e=>{this.$notify({title:"成功",message:"更新子组成功！",type:"success"}),this.needtoUpdate=!0})).catch((e=>console.log(e)))},pushNewLablesList(){},filterIPMethod:(e,t)=>!0,doEditIPList(e){const t=this.jobInfo.id,l=e.row.id;u(l).then((a=>{let i=[];this.editIPValueMap={},a.data.forEach((e=>{i.push(e.machines_id),this.editIPValueMap[e.machines_id]=e}));let s=[];this.editIPDataMap={},p(t).then((t=>{t.data.forEach((e=>{let t={label:e.ipaddr,key:e.id,spell:e.ipaddr,disabled:!1};s.push(t),this.editIPDataMap[e.id]=e})),this.editIPData=s,this.editIPValue=i,this.editObject=e.row.name,this.jobGroupIDCurrent=l,this.editIPVisible=!0})).catch((e=>{console.log(e)}))})).catch((e=>{console.log(e)}))},editIPClose(e){this.needtoUpdate&&this.doGetSubGroup(),this.editIPVisible=!1},editLabelsClose(){this.needtoUpdate&&this.doGetSubGroup(),this.editLabelsVisible=!1},doEditLabelList(e){this.jobInfo.id;const t=e.row.id,l={pageNo:this.glCurrentPage,pageSize:this.glPageSize,search:this.glSearchContent,id:t};G().then((a=>{this.defaultLabels=a.data,h(l).then((l=>{this.deleteLabelsVisible={};let a=0;l.data.data.forEach((e=>{this.deleteLabelsVisible[a]=!1,a+=1})),this.glPageTotal=l.data.totalCount,this.glCurrentPage=l.data.pageNo,this.glPageSize=l.data.pageSize,this.labelsData=l.data.data,this.jobGroupIDCurrent=t,this.editLabelsVisible=!0,this.editObject=e.row.name})).catch((e=>console.log(e)))})).catch((e=>console.log(e)))},doGetGroupLabels(e){e||(e={pageNo:this.glCurrentPage,pageSize:this.glPageSize,search:this.glSearchContent,id:this.jobGroupIDCurrent}),h(e).then((e=>{this.deleteLabelsVisible={};let t=0;e.data.data.forEach((e=>{this.deleteLabelsVisible[t]=!1,t+=1})),this.glPageTotal=e.data.totalCount,this.glCurrentPage=e.data.pageNo,this.glPageSize=e.data.pageSize,this.labelsData=e.data.data,this.jobGroupIDCurrent=gid,this.editLabelsVisible=!0})).catch((e=>console.log(e)))},doAddSubGroupShow(){this.buttonTitle="创建",this.buttonType="primary",this.dialogTitle="创建新的子分组",this.dialogVisible=!0},doAddSubGroup(){b({}).then((e=>{this.subGroup=e.data,this.$notify({title:"成功",message:"创建子组成功！",type:"success"})})).catch((e=>{console.log(e)}))},onAddNewLable(e){this.$refs[e].validate((e=>{if(e){const e=this.jobGroupIDCurrent;let t=n(d({},this.addNewGroupLabels),{job_group_id:this.jobGroupIDCurrent});c(e,t).then((e=>{this.$notify({title:"成功",message:"创建标签成功！",type:"success"}),this.doGetGroupLabels(),this.needtoUpdate=!0})).catch((e=>console.log(e)))}}))},onResetAddNewLable(e){this.$refs[e].resetFields(),this.alertTitle="当前状态：可以添加新标签",this.addNewGroupLabels={},this.labelsBtnTitle="添加",this.buttonType="primary"},doEdit(e){this.buttonTitle="更新",this.buttonType="warning",this.dialogTitle="更新子分组",this.subGroupData=d({},e.row),this.dialogVisible=!0},doLabelsEdit(e){this.alertTitle=`当前编辑对象：${e.row.key}，序号：${e.$index+1}`,this.addNewGroupLabels={key:e.row.key,value:e.row.value,id:e.row.id},this.labelsBtnTitle="更新"},handleSizeChange(e){let t={pageNo:this.currentPage,pageSize:e,search:this.searchContent};this.doGetSubGroup(t)},glHandleSizeChange(e){let t={pageNo:this.glCurrentPage,pageSize:e,search:this.glSearchContent,id:this.jobGroupIDCurrent};this.doGetGroupLabels(t)},handleCurrentChange(e){let t={pageNo:e,pageSize:this.pageSize,search:this.searchContent};this.doGetSubGroup(t)},glHandleCurrentChange(e){let t={pageNo:e,pageSize:this.glPageSize,search:this.glSearchContent,id:this.jobGroupIDCurrent};this.doGetGroupLabels(t)},rowStyle:e=>({padding:"0"}),cellStyle:e=>({padding:"0"}),parseTimeSelf(e){var t=new Date(Date.parse(e));return t.toLocaleDateString()+" "+t.toTimeString().split(" ")[0]},doGetSubGroup(e){e||(e={pageNo:this.currentPage,pageSize:this.pageSize,search:this.searchContent}),isNaN(this.jobInfo.id)||(e.id=this.jobInfo.id,g(e).then((e=>{let t=0;this.ipsAndLabels={},this.ipsAndLabelsDone={},e.data.data.forEach((e=>{this.deleteVisible[t]=!1,this.ipsAndLabels[e.id]={},t+=1})),this.subGroups=e.data.data,this.pageTotal=e.data.totalCount,this.currentPage=e.data.pageNo,this.pageSize=e.data.pageSize})).catch((e=>{console.log(e)})))},onSubmit(e){const t=n(d({},this.subGroupData),{jobs_id:this.jobInfo.id});this.$refs[e].validate((l=>{l&&("创建"===this.buttonTitle?b(t).then((t=>{this.$notify({title:"成功",message:"创建子组成功！",type:"success"}),this.doGetSubGroup(),this.dialogVisible=!1,this.$refs[e].resetFields()})).catch((e=>{console.log(e)})):f(t).then((t=>{this.$notify({title:"成功",message:"更新子组成功！",type:"success"}),this.doGetSubGroup(),this.dialogVisible=!1,this.$refs[e].resetFields()})).catch((e=>{console.log(e)})))}))},handleClose(e){this.dialogVisible=!1},onCancel(e){this.subGroupData={},this.$refs[e].resetFields(),this.dialogVisible=!1},doYes(e){const t={jobs_id:this.jobInfo.id,id:e.row.id};m(t).then((e=>{this.$notify({title:"成功",message:"删除子组成功！",type:"success"}),this.doGetSubGroup()})).catch((e=>{console.log(e)})),this.deleteVisible[e.$index]=!1},doLabelsYes(e){const t={gid:this.jobGroupIDCurrent,id:e.row.id};y(t).then((e=>{this.$notify({title:"成功",message:"删除标签成功！",type:"success"}),this.doGetGroupLabels()})).catch((e=>{console.log(e)})),this.deleteLabelsVisible[e.$index]=!1},doNo(e){this.deleteVisible[e.$index]=!1},doLablesNo(e){this.deleteLabelsVisible[e.$index]=!1},doDelete(e){this.deleteVisible[e.$index]=!0},doLabelsDelete(e){this.deleteLabelsVisible[e.$index]=!0},onSearch(){this.doGetSubGroup()},onLablesSearch(){this.doGetGroupLabels()},clickElTag(e){console.log("checked"),this.$router.push({name:"jobs"})},expandChange(e,t){if(!t||0===t.length)return;if(this.ipsAndLabelsDone[e.id])return;const l={job_id:e.jobs_id,group_id:e.id};C(l).then((t=>{this.ipsAndLabels[e.id]=t.data,this.ipsAndLabelsDone[e.id]=!0})).catch((e=>console.log(e)))},invocate(e){const t=!this.subGroups[e.$index].enabled,l={id:e.row.id,enabled:t};L(l).then((l=>{this.$notify({title:"成功",message:"更新状态成功！",type:"success"}),this.subGroups[e.$index].enabled=t,this.subGroups=[...this.subGroups]})).catch((e=>console.log(e)))},invocateLabels(e){const t=!this.labelsData[e.$index].enabled,l={id:e.row.id,enabled:t};w(l).then((l=>{this.$notify({title:"成功",message:"更新状态成功！",type:"success"}),this.labelsData[e.$index].enabled=t,this.labelsData=[...this.labelsData]})).catch((e=>console.log(e)))}}},E=$();S("data-v-5a471ad7");const A={class:"jobs-labels-board"},O={class:"label-and-action"},U={class:"do_action"},M={style:{"padding-right":"15px"}},F=P("发布此分组"),B=P("添加子组"),H={class:"actioneara"},R=P("编辑"),Y=P("编辑IP"),q=P("编辑标签"),J=P("禁用"),K=P("启用"),Q=v("p",null,"确定删除吗？",-1),W={style:{"text-align":"right",margin:"0"}},X=P("取消"),Z=P("确定"),ee=P("删除"),te={class:"block"},le=P("取消"),ae={class:"ip-list-push-box"},ie=P("提交"),se={class:"add-label-form"},oe={class:"add-labels-select-box"},de={class:"add-labels-input-box"},ne={class:"add-labels-btn-box"},re=P("重置"),ue={class:"status-edit-area"},pe={class:"search-label-input"},he={class:"actioneara"},be=P("编辑"),ce=P("禁用"),ge=P("启用"),fe=v("p",null,"确定删除吗？",-1),me={style:{"text-align":"right",margin:"0"}},ye=P("取消"),Ce=P("确定"),Le=P("删除"),we={class:"block"};V();const Ge=E(((e,t,l,a,i,s)=>{const o=_("el-tag"),d=_("el-button"),n=_("el-input"),r=_("el-descriptions-item"),u=_("el-descriptions"),p=_("el-table-column"),h=_("el-popover"),b=_("el-table"),c=_("el-pagination"),g=_("el-form-item"),f=_("el-form"),m=_("el-dialog"),y=_("el-transfer"),C=_("el-row"),L=_("el-option"),w=_("el-select"),G=_("el-alert");return z(),k("div",A,[v("div",O,[v("div",null,[v(o,{type:"warning",effect:"dark",onChange:s.clickElTag},{default:E((()=>[P("编辑分组："+I(i.jobInfo.name)+" [ "+I(i.jobInfo.count)+" ]",1)])),_:1},8,["onChange"])]),v("div",U,[v("div",M,[v(d,{size:"small",type:"primary",onClick:t[1]||(t[1]=t=>e.publishJobsfunc())},{default:E((()=>[F])),_:1}),v(d,{size:"small",type:"success",plain:"",onClick:t[2]||(t[2]=e=>s.doAddSubGroupShow())},{default:E((()=>[B])),_:1})]),v("div",null,[v("div",null,[v(n,{size:"small",placeholder:"请输入内容",onKeyup:t[4]||(t[4]=j((e=>s.onSearch()),["enter"])),modelValue:i.searchContent,"onUpdate:modelValue":t[5]||(t[5]=e=>i.searchContent=e)},{append:E((()=>[v(d,{size:"small",onClick:t[3]||(t[3]=e=>s.onSearch()),icon:"el-icon-search"})])),_:1},8,["modelValue"])])])])]),v(b,{size:"mini","highlight-current-row":"",border:"",data:i.subGroups,stripe:"","row-style":s.rowStyle,"cell-style":s.cellStyle,onExpandChange:s.expandChange},{default:E((()=>[v(p,{type:"expand"},{default:E((e=>[v(u,{title:"IP列表",size:"mini",column:3,border:""},{default:E((()=>[(z(!0),k(D,null,x(i.ipsAndLabels[e.row.id].ips,((e,t)=>(z(),k(r,{key:t,label:t+1},{default:E((()=>[P(I(e),1)])),_:2},1032,["label"])))),128))])),_:2},1024),v(u,{title:"标签列表",size:"mini",column:3,border:""},{default:E((()=>[(z(!0),k(D,null,x(i.ipsAndLabels[e.row.id].labels,((e,t)=>(z(),k(r,{key:t,label:e.key},{default:E((()=>[P(I(e.value),1)])),_:2},1032,["label"])))),128))])),_:2},1024)])),_:1}),v(p,{label:"序号",width:"50px"},{default:E((e=>[P(I(e.$index+1),1)])),_:1}),v(p,{label:"名称",prop:"name"}),v(p,{label:"IP数",prop:"ip_count"}),v(p,{label:"标签数",prop:"labels_count"}),v(p,{label:"最后更新时间",prop:"update_at"},{default:E((({row:e})=>[v("span",null,I(s.parseTimeSelf(e.update_at)),1)])),_:1}),v(p,{label:"操作",width:"350px",align:"center"},{default:E((e=>[v("div",H,[v("div",null,[v(d,{size:"mini",type:"primary",onClick:t=>s.doEdit(e),plain:""},{default:E((()=>[R])),_:2},1032,["onClick"])]),v("div",null,[v(d,{size:"mini",type:"primary",onClick:t=>s.doEditIPList(e),plain:""},{default:E((()=>[Y])),_:2},1032,["onClick"])]),v("div",null,[v(d,{size:"mini",type:"primary",onClick:t=>s.doEditLabelList(e),plain:""},{default:E((()=>[q])),_:2},1032,["onClick"])]),v("div",null,[!0===e.row.enabled?(z(),k(d,{key:0,size:"mini",type:"info",onClick:t=>s.invocate(e),plain:""},{default:E((()=>[J])),_:2},1032,["onClick"])):T("",!0),!1===e.row.enabled?(z(),k(d,{key:1,size:"mini",type:"warning",onClick:t=>s.invocate(e),plain:""},{default:E((()=>[K])),_:2},1032,["onClick"])):T("",!0)]),v("div",null,[v(h,{visible:i.deleteVisible[e.$index],placement:"top"},{reference:E((()=>[v(d,{size:"mini",type:"danger",plain:"",onClick:t=>s.doDelete(e)},{default:E((()=>[ee])),_:2},1032,["onClick"])])),default:E((()=>[Q,v("div",W,[v(d,{size:"mini",type:"text",onClick:t=>s.doNo(e)},{default:E((()=>[X])),_:2},1032,["onClick"]),v(d,{type:"primary",size:"mini",onClick:t=>s.doYes(e)},{default:E((()=>[Z])),_:2},1032,["onClick"])])])),_:2},1032,["visible"])])])])),_:1})])),_:1},8,["data","row-style","cell-style","onExpandChange"]),v("div",te,[v(c,{onSizeChange:s.handleSizeChange,onCurrentChange:s.handleCurrentChange,"current-page":i.currentPage,"page-sizes":[10,15,30,50],"page-size":i.pageSize,layout:"total, sizes, prev, pager, next, jumper",total:i.pageTotal},null,8,["onSizeChange","onCurrentChange","current-page","page-size","total"])]),v(m,{title:i.dialogTitle,modelValue:i.dialogVisible,"onUpdate:modelValue":t[9]||(t[9]=e=>i.dialogVisible=e),width:"400px",modal:"","before-close":s.handleClose},{default:E((()=>[v("span",null,[v(f,{"label-position":"right",rules:i.rules,ref:"subGroupData",model:i.subGroupData,"label-width":"90px",size:"small"},{default:E((()=>[v(g,{label:"子组名：",prop:"name"},{default:E((()=>[v(n,{style:{width:"230px"},modelValue:i.subGroupData.name,"onUpdate:modelValue":t[6]||(t[6]=e=>i.subGroupData.name=e)},null,8,["modelValue"])])),_:1}),v(g,{size:"small"},{default:E((()=>[v(d,{size:"small",type:i.buttonType,onClick:t[7]||(t[7]=e=>s.onSubmit("subGroupData"))},{default:E((()=>[P(I(i.buttonTitle),1)])),_:1},8,["type"]),v(d,{size:"small",type:"info",onClick:t[8]||(t[8]=e=>s.onCancel("subGroupData"))},{default:E((()=>[le])),_:1})])),_:1})])),_:1},8,["rules","model"])])])),_:1},8,["title","modelValue","before-close"]),v(m,{title:"编辑IP列表："+i.editObject,modelValue:i.editIPVisible,"onUpdate:modelValue":t[11]||(t[11]=e=>i.editIPVisible=e),width:"700px",modal:"","before-close":s.editIPClose},{default:E((()=>[v(C,{type:"flex",align:"middle",justify:"center"},{default:E((()=>[v(y,{modelValue:i.editIPValue,"onUpdate:modelValue":t[10]||(t[10]=e=>i.editIPValue=e),filterable:"","filter-method":s.filterIPMethod,"filter-placeholder":"请输入关键字",data:i.editIPData,titles:["分组IP地址池","子分组IP列表"]},null,8,["modelValue","filter-method","data"])])),_:1}),v("div",ae,[v(d,{class:"ip-list-push-btn",type:"warning",size:"small",onClick:s.pushNewIPList},{default:E((()=>[ie])),_:1},8,["onClick"])])])),_:1},8,["title","modelValue","before-close"]),v(m,{title:"编辑标签列表："+i.editObject,modelValue:i.editLabelsVisible,"onUpdate:modelValue":t[19]||(t[19]=e=>i.editLabelsVisible=e),width:"800px",modal:"","before-close":s.editLabelsClose},{default:E((()=>[v("div",se,[v(f,{size:"mini",inline:!0,rules:i.glRules,ref:"addLablesForm",model:i.addNewGroupLabels,class:"demo-form-inline"},{default:E((()=>[v(g,{label:"标签：",prop:"key"},{default:E((()=>[v("div",oe,[v(w,{modelValue:i.addNewGroupLabels.key,"onUpdate:modelValue":t[12]||(t[12]=e=>i.addNewGroupLabels.key=e),filterable:"","allow-create":"","default-first-option":"",placeholder:"请选择或者输入"},{default:E((()=>[(z(!0),k(D,null,x(i.defaultLabels,(e=>(z(),k(L,{key:e.id,label:e.label,value:e.label},null,8,["label","value"])))),128))])),_:1},8,["modelValue"])])])),_:1}),v(g,{label:"标签值：",prop:"value"},{default:E((()=>[v("div",de,[v(n,{modelValue:i.addNewGroupLabels.value,"onUpdate:modelValue":t[13]||(t[13]=e=>i.addNewGroupLabels.value=e)},null,8,["modelValue"])])])),_:1}),v(g,null,{default:E((()=>[v("div",ne,[v(d,{type:"primary",onClick:t[14]||(t[14]=e=>s.onAddNewLable("addLablesForm"))},{default:E((()=>[P(I(i.labelsBtnTitle),1)])),_:1}),v(d,{type:"danger",onClick:t[15]||(t[15]=e=>s.onResetAddNewLable("addLablesForm"))},{default:E((()=>[re])),_:1})])])),_:1})])),_:1},8,["rules","model"])]),v("div",ue,[v("div",null,[v(G,{effect:"dark",title:i.alertTitle,type:"warning",closable:!1,"show-icon":""},null,8,["title"])]),v("div",pe,[v(n,{size:"small",placeholder:"请输入内容",onKeyup:t[17]||(t[17]=j((e=>s.onLablesSearch()),["enter"])),modelValue:i.glSearchContent,"onUpdate:modelValue":t[18]||(t[18]=e=>i.glSearchContent=e)},{append:E((()=>[v(d,{size:"small",onClick:t[16]||(t[16]=e=>s.onLablesSearch()),icon:"el-icon-search"})])),_:1},8,["modelValue"])])]),v("div",null,[v(b,{size:"mini","highlight-current-row":"",border:"",stripe:"","row-style":s.rowStyle,"cell-style":s.cellStyle,"header-align":"center",data:i.labelsData},{default:E((()=>[v(p,{label:"序号",width:"50px",align:"center","header-align":"center"},{default:E((e=>[P(I(e.$index+1),1)])),_:1}),v(p,{property:"key",label:"标签名"}),v(p,{property:"value",label:"标签值"}),v(p,{label:"最后更新",prop:"update_at",align:"center","header-align":"center"},{default:E((({row:e})=>[v("span",null,I(s.parseTimeSelf(e.update_at)),1)])),_:1}),v(p,{label:"操作",align:"center"},{default:E((e=>[v("div",he,[v("div",null,[v(d,{size:"mini",type:"primary",onClick:t=>s.doLabelsEdit(e),plain:""},{default:E((()=>[be])),_:2},1032,["onClick"])]),v("div",null,[!0===e.row.enabled?(z(),k(d,{key:0,size:"mini",type:"info",onClick:t=>s.invocateLabels(e),plain:""},{default:E((()=>[ce])),_:2},1032,["onClick"])):T("",!0),!1===e.row.enabled?(z(),k(d,{key:1,size:"mini",type:"warning",onClick:t=>s.invocateLabels(e),plain:""},{default:E((()=>[ge])),_:2},1032,["onClick"])):T("",!0)]),v("div",null,[v(h,{visible:i.deleteLabelsVisible[e.$index],placement:"top"},{reference:E((()=>[v(d,{size:"mini",type:"danger",plain:"",onClick:t=>s.doLabelsDelete(e)},{default:E((()=>[Le])),_:2},1032,["onClick"])])),default:E((()=>[fe,v("div",me,[v(d,{size:"mini",type:"text",onClick:t=>s.doLablesNo(e)},{default:E((()=>[ye])),_:2},1032,["onClick"]),v(d,{type:"primary",size:"mini",onClick:t=>s.doLabelsYes(e)},{default:E((()=>[Ce])),_:2},1032,["onClick"])])])),_:2},1032,["visible"])])])])),_:1})])),_:1},8,["row-style","cell-style","data"]),v("div",we,[v(c,{onSizeChange:s.glHandleSizeChange,onCurrentChange:s.glHandleCurrentChange,"current-page":i.glCurrentPage,"page-sizes":[10],"page-size":i.glPageSize,layout:"total, sizes, prev, pager, next, jumper",total:i.glPageTotal},null,8,["onSizeChange","onCurrentChange","current-page","page-size","total"])])])])),_:1},8,["title","modelValue","before-close"])])}));N.render=Ge,N.__scopeId="data-v-5a471ad7";export default N;
