import{s as e}from"./request.d1dca4b5.js";import{f as t}from"./relabel.01b25169.js";import{a as l}from"./srv.20227253.js";import{p as a,j as i,r as o,o as s,c as d,a as n,q as r,w as p,l as h,t as u,F as c,n as b,k as g}from"./vendor.02dbbba7.js";import"./index.a2263c8a.js";const m={name:"JobsDefault",data:()=>({jobs:[],addJobInfo:{name:"",port:0,display_order:1,relabel_id:0},relabelList:[],pageSize:15,pageTotal:0,currentPage:1,searchContent:"",deleteVisible:{},dialogVisible:!1,buttonTitle:"",dialogTitle:"",rules:{name:[{required:!0,message:"请输入正确的分组名称",validator:function(e,t,l){if(""===t||void 0===t||null==t)l(new Error("请输入正确名称"));else{/\s+/.test(t)?l(new Error("名称中不允许有空字符")):l()}},trigger:["blur"]}],port:[{required:!0,type:"number",min:0,max:65535,message:"请输入端口号[>=0, <=65535]",trigger:["change"]}],relabel_id:[{required:!0,type:"number",min:1,message:"请选择重写标签",trigger:["change"]}]},publishMode:!1}),created(){},mounted(){this.doGetJobs()},methods:{doAdd(){t().then((e=>{this.relabelList=e.data,this.addJobInfo={name:"",port:0,display_order:1},this.buttonTitle="创建",this.dialogTitle="增加IP分组",this.dialogVisible=!0})).catch((e=>{console.log(e)}))},doGetJobs(t){var l;t||(t={pageNo:this.currentPage,pageSize:this.pageSize,search:this.searchContent}),(l=t,e({url:"/v1/jobs/default/split",method:"get",params:l})).then((e=>{let t=0;e.data.data.forEach((e=>{this.deleteVisible[t]=!1,t+=1})),this.jobs=e.data.data,this.pageTotal=e.data.totalCount,this.currentPage=e.data.pageNo,this.pageSize=e.data.pageSize})).catch((e=>{console.log(e)}))},doEdit(e){t().then((t=>{this.relabelList=t.data,this.buttonTitle="更新",this.dialogTitle="编辑IP分组",this.addJobInfo.id=e.row.id,this.addJobInfo.name=e.row.name,this.addJobInfo.port=e.row.port,this.addJobInfo.relabel_id=e.row.relabel_id,this.addJobInfo.relabel_name=e.row.relabel_name,this.dialogVisible=!0})).catch((e=>{console.log(e)}))},handleSizeChange(e){let t={pageNo:this.currentPage,pageSize:e,search:this.searchContent};this.doGetJobs(t)},handleCurrentChange(e){let t={pageNo:e,pageSize:this.pageSize,search:this.searchContent};this.doGetJobs(t)},handleClose(e){this.dialogVisible=!1},onSubmit(t){this.$refs[t].validate((l=>{if(!l)return!1;{let l={};l.name=this.addJobInfo.name,l.port=this.addJobInfo.port,l.display_order=this.addJobInfo.display_order,l.relabel_id=this.addJobInfo.relabel_id,"创建"===this.buttonTitle?(a=l,e({url:"/v1/job/default",method:"post",data:a})).then((e=>{this.$notify({title:"成功",message:"创建成功！",type:"success"}),this.doGetJobs(),this.dialogVisible=!1,this.$refs[t].resetFields()})).catch((e=>{console.log(e)})):(l.id=this.addJobInfo.id,function(t){return e({url:"/v1/job/default",method:"put",data:t})}(l).then((e=>{this.$notify({title:"成功",message:"更新成功！",type:"success"}),this.doGetJobs(),this.dialogVisible=!1,this.$refs[t].resetFields()})).catch((e=>{console.log(e)})))}var a}))},onCancel(e){this.dialogVisible=!1,this.$refs[e].resetFields()},onSearch(){this.doGetJobs()},parseTimeSelf(e){var t=new Date(Date.parse(e));return t.toLocaleDateString()+" "+t.toTimeString().split(" ")[0]},doYes(t){var l;(l=t.row.id,e({url:"/v1/job/default",method:"delete",params:{id:l}})).then((e=>{this.$notify({title:"成功",message:"删除成功！",type:"success"}),this.doGetJobs()})).catch((e=>{console.log(e)})),this.deleteVisible[t.$index]=!1},doNo(e){this.deleteVisible[e.$index]=!1},doDelete(e){this.deleteVisible[e.$index]=!0},rowStyle:e=>({padding:"0"}),cellStyle:e=>({padding:"0"}),doUp(e){return 0===this.jobs.length?(this.$notify({title:"错误",message:"没有有效的数据！",type:"error"}),!1):1===e.row.display_order?(this.$notify({title:"警告",message:"已经是顶层！",type:"warning"}),!1):void this.swapAction({id:e.row.id,action:"up",display_order:e.row.display_order})},doDown(e){if(0===this.jobs.length)return this.$notify({title:"错误",message:"没有有效的数据！",type:"error"}),!1;return this.jobs[this.jobs.length-1].id===e.row.id&&this.jobs.length!==this.pageSize||this.currentPage*this.pageSize===this.pageTotal?(this.$notify({title:"警告",message:"已经是底层！",type:"warning"}),!1):void this.swapAction({id:e.row.id,action:"down",display_order:e.row.display_order})},swapAction(e){swapJob(e).then((t=>{"up"===e.action?this.$notify({title:"成功",message:"提升成功！",type:"success"}):this.$notify({title:"成功",message:"下降成功！",type:"success"}),this.doGetJobs()})).catch((e=>{console.log(e)}))},publishJobsfunc(){this.publishMode=!0,e({url:"/v1/job/default/publish",method:"post"}).then((e=>{this.$notify({title:"成功",message:"发布成功！",type:"success"}),this.publishMode=!1})).catch((e=>{this.publishMode=!1,console.log(e)}))},restartServer(){l().then((e=>{this.$notify({title:"成功",message:"重新加载成功！",type:"success"})})).catch((e=>{console.log(e)}))},invocate(t){const l=!this.jobs[t.$index].enabled,a={id:t.row.id,enabled:l};var i;(i=a,e({url:"/v1/job/default/status",method:"put",data:i})).then((e=>{this.$notify({title:"成功",message:"更新状态成功！",type:"success"}),this.jobs[t.$index].enabled=l,this.jobs=[...this.jobs]})).catch((e=>console.log(e)))}}},f=g();a("data-v-2fe4e7a0");const y={class:"main-board"},_={class:"btn-action-area"},w={class:"explain-words"},v=h("说明：定义包含"),C=h("所有IP地址"),z=h("的分组"),J={class:"do_action"},S={style:{"padding-right":"15px"}},V=h("让Prometheus服务重新加载配置"),I=h("发布"),k=h("发布"),j=h("添加默认组"),x={class:"actioneara"},$=h("编辑"),T=h("禁用"),P=h("启用"),G=n("p",null,"确定删除吗？",-1),D={style:{"text-align":"right",margin:"0"}},M=h("取消"),N=h("确定"),U=h("删除"),q={class:"block"},A=h("取消");i();const E=f(((e,t,l,a,i,g)=>{const m=o("el-tag"),E=o("el-button"),L=o("el-input"),F=o("el-table-column"),Y=o("el-popover"),K=o("el-table"),R=o("el-pagination"),B=o("el-form-item"),H=o("el-option"),O=o("el-select"),Q=o("el-form"),W=o("el-dialog");return s(),d("div",y,[n("div",_,[n("div",null,[n("span",w,[v,n(m,{size:"mini",type:"warning"},{default:f((()=>[C])),_:1}),z])]),n("div",J,[n("div",S,[n(E,{size:"small",type:"warning",onClick:t[1]||(t[1]=e=>g.restartServer())},{default:f((()=>[V])),_:1}),!1===i.publishMode?(s(),d(E,{key:0,size:"small",icon:"el-icon-upload",type:"primary",onClick:t[2]||(t[2]=e=>g.publishJobsfunc())},{default:f((()=>[I])),_:1})):r("",!0),!0===i.publishMode?(s(),d(E,{key:1,icon:"el-icon-loading",size:"small",type:"primary",onClick:t[3]||(t[3]=t=>e.publishJobsRunning())},{default:f((()=>[k])),_:1})):r("",!0),n(E,{size:"small",type:"success",plain:"",onClick:t[4]||(t[4]=e=>g.doAdd())},{default:f((()=>[j])),_:1})]),n("div",null,[n("div",null,[n(L,{size:"small",placeholder:"请输入内容",onKeyup:t[6]||(t[6]=p((e=>g.onSearch()),["enter"])),modelValue:i.searchContent,"onUpdate:modelValue":t[7]||(t[7]=e=>i.searchContent=e)},{append:f((()=>[n(E,{size:"small",onClick:t[5]||(t[5]=e=>g.onSearch()),icon:"el-icon-search"})])),_:1},8,["modelValue"])])])])]),n(K,{size:"mini","highlight-current-row":"",border:"",data:i.jobs,stripe:"","row-style":g.rowStyle,"cell-style":g.cellStyle},{default:f((()=>[n(F,{label:"序号",width:"50px"},{default:f((e=>[h(u(e.$index+1),1)])),_:1}),n(F,{label:"分组名称",prop:"name"}),n(F,{label:"端口号",width:"90px",prop:"port"}),n(F,{label:"IP数",width:"80px",prop:"ip_count"}),n(F,{label:"重写标签",prop:"relabel_name"}),n(F,{label:"最新更新账号",prop:"update_by"}),n(F,{label:"最后更新时间",prop:"update_at"},{default:f((({row:e})=>[n("span",null,u(g.parseTimeSelf(e.update_at)),1)])),_:1}),n(F,{label:"操作",align:"center",width:"200px"},{default:f((e=>[n("div",x,[n("div",null,[n(E,{size:"mini",type:"primary",onClick:t=>g.doEdit(e),plain:""},{default:f((()=>[$])),_:2},1032,["onClick"])]),n("div",null,[!0===e.row.enabled?(s(),d(E,{key:0,size:"mini",type:"info",onClick:t=>g.invocate(e),plain:""},{default:f((()=>[T])),_:2},1032,["onClick"])):r("",!0),!1===e.row.enabled?(s(),d(E,{key:1,size:"mini",type:"warning",onClick:t=>g.invocate(e),plain:""},{default:f((()=>[P])),_:2},1032,["onClick"])):r("",!0)]),n("div",null,[n(Y,{visible:i.deleteVisible[e.$index],placement:"top"},{reference:f((()=>[n(E,{size:"mini",type:"danger",plain:"",onClick:t=>g.doDelete(e)},{default:f((()=>[U])),_:2},1032,["onClick"])])),default:f((()=>[G,n("div",D,[n(E,{size:"mini",type:"text",onClick:t=>g.doNo(e)},{default:f((()=>[M])),_:2},1032,["onClick"]),n(E,{type:"primary",size:"mini",onClick:t=>g.doYes(e)},{default:f((()=>[N])),_:2},1032,["onClick"])])])),_:2},1032,["visible"])])])])),_:1})])),_:1},8,["data","row-style","cell-style"]),n("div",q,[n(R,{onSizeChange:g.handleSizeChange,onCurrentChange:g.handleCurrentChange,"current-page":i.currentPage,"page-sizes":[10,15,30,50],"page-size":i.pageSize,layout:"total, sizes, prev, pager, next, jumper",total:i.pageTotal},null,8,["onSizeChange","onCurrentChange","current-page","page-size","total"])]),n(W,{title:i.dialogTitle,modelValue:i.dialogVisible,"onUpdate:modelValue":t[13]||(t[13]=e=>i.dialogVisible=e),width:"400px",modal:"","before-close":g.handleClose},{default:f((()=>[n("span",null,[n(Q,{"label-position":"right",rules:i.rules,ref:"addJobInfo",model:i.addJobInfo,"label-width":"90px",size:"small"},{default:f((()=>[n(B,{label:"分组名：",prop:"name"},{default:f((()=>[n(L,{style:{width:"230px"},modelValue:i.addJobInfo.name,"onUpdate:modelValue":t[8]||(t[8]=e=>i.addJobInfo.name=e)},null,8,["modelValue"])])),_:1}),n(B,{label:"端口号",prop:"port"},{default:f((()=>[n(L,{type:"number",style:{width:"230px"},modelValue:i.addJobInfo.port,"onUpdate:modelValue":t[9]||(t[9]=e=>i.addJobInfo.port=e),modelModifiers:{number:!0}},null,8,["modelValue"])])),_:1}),n(B,{label:"重写标签",prop:"relabel_id"},{default:f((()=>[n(O,{modelValue:i.addJobInfo.relabel_id,"onUpdate:modelValue":t[10]||(t[10]=e=>i.addJobInfo.relabel_id=e),filterable:"","allow-create":"","default-first-option":"",placeholder:"请选择",style:{width:"230px"}},{default:f((()=>[(s(!0),d(c,null,b(i.relabelList,(e=>(s(),d(H,{key:e.id,label:e.name,value:e.id},null,8,["label","value"])))),128))])),_:1},8,["modelValue"])])),_:1}),n(B,{size:"small"},{default:f((()=>[n(E,{size:"small",type:"primary",onClick:t[11]||(t[11]=e=>g.onSubmit("addJobInfo"))},{default:f((()=>[h(u(i.buttonTitle),1)])),_:1}),n(E,{size:"small",type:"info",onClick:t[12]||(t[12]=e=>g.onCancel("addJobInfo"))},{default:f((()=>[A])),_:1})])),_:1})])),_:1},8,["rules","model"])])])),_:1},8,["title","modelValue","before-close"])])}));m.render=E,m.__scopeId="data-v-2fe4e7a0";export default m;