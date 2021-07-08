var e=Object.defineProperty,t=Object.getOwnPropertySymbols,i=Object.prototype.hasOwnProperty,a=Object.prototype.propertyIsEnumerable,l=(t,i,a)=>i in t?e(t,i,{enumerable:!0,configurable:!0,writable:!0,value:a}):t[i]=a,n=(e,n)=>{for(var d in n||(n={}))i.call(n,d)&&l(e,d,n[d]);if(t)for(var d of t(n))a.call(n,d)&&l(e,d,n[d]);return e};import{g as d}from"./jobs.75f86484.js";import{s}from"./request.0c690da8.js";import{p as o}from"./publish.9be40a68.js";import{p as r,b as h,r as c,o as p,c as u,a as g,h as m,F as b,i as f,j as y,t as C,k as V,w as _}from"./vendor.61ee7ef1.js";function I(e){return s({url:"/v1/machine",method:"post",data:e})}function v(e){return s({url:"/v1/machine",method:"put",data:e})}const w={name:"Manager",components:{},data:()=>({machines:[],jobs:[],jobsMap:{},selectTypeValue:{},search:"",pageSize:15,pageTotal:0,currentPage:1,selectOption:"1",searchContent:"",dialogVisible:!1,deleteVisible:{},enterBtnTitle:"创建",selectTableChanged:{},addMechineInfo:{ipAddr:"",jobId:[]},groupNameList:{},addAndContinueDisabled:!1,rules:{ipAddr:[{required:!0,message:"请输入正确的IP地址",validator:function(e,t,i){if(""===t||void 0===t||null==t)i(new Error("请输入正确的IP地址"));else{/^(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])$/.test(t)||""===t?i():i(new Error("请输入正确的IP地址"))}},trigger:"blur"}],jobId:[{required:!0,message:"请选择分组",trigger:["blur","change"]}]}}),created(){},mounted(){this.doGetMechines()},methods:{doAdd(){this.enterBtnTitle="创建",this.dialogVisible=!0,this.addAndContinueDisabled=!1},doGetMechines(e){d().then((t=>{var i;this.jobs=t.data,t.data.forEach((e=>{this.jobsMap[e.id]=e.name})),e||(e={pageNo:this.currentPage,pageSize:this.pageSize,search:this.searchContent,option:this.selectOption}),(i=e,s({url:"/v1/machines",method:"get",params:i})).then((e=>{this.pageTotal=e.data.totalCount,this.currentPage=e.data.pageNo,this.pageSize=e.data.pageSize,this.machines=e.data.data,this.deleteVisible={};let t=0;e.data.data.forEach((e=>{this.selectTypeValue[e.id]=e.job_id,this.deleteVisible[t]=!1,t+=1}))})).catch((e=>{console.log(e)}))})).catch((e=>{console.log(e)}))},handleSelect(e,t,i){if(e)this.selectTableChanged[i.id]=!1;else if(!1!==this.selectTableChanged[i.id]){var a={};a.id=i.id,a.ipaddr=i.ipaddr,a.job_id=this.selectTypeValue[i.id],v(a).then((e=>{this.$notify({title:"成功",message:"更新成功！",type:"success"}),this.selectTableChanged[i.id]=!1,this.selectTypeValue=n({},this.selectTypeValue),this.expandChange(i,i)})).catch((e=>{console.log(e)}))}},tableSelectChange(e,t,i){this.selectTableChanged[i.id]=!0},handleSizeChange(e){let t={pageNo:this.currentPage,pageSize:e,search:this.searchContent,option:this.selectOption};this.doGetMechines(t)},handleCurrentChange(e){let t={pageNo:e,pageSize:this.pageSize,search:this.searchContent,option:this.selectOption};this.doGetMechines(t)},handleClose(e){this.dialogVisible=!1},onSubmitAndContinue(e){this.$refs[e].validate((e=>{if(!e)return!1;{let e={};e.ipaddr=this.addMechineInfo.ipAddr,e.job_id=[parseInt(this.addMechineInfo.jobId)],I(e).then((e=>{this.$notify({title:"成功",message:"创建成功！",type:"success"}),this.doGetMechines()})).catch((e=>{console.log(e)}))}}))},onSubmit(e){this.$refs[e].validate((t=>{if(!t)return!1;{let t={};t.ipaddr=this.addMechineInfo.ipAddr,t.job_id=this.addMechineInfo.jobId,"创建"===this.enterBtnTitle?I(t).then((t=>{this.$notify({title:"成功",message:"创建成功！",type:"success"}),this.doGetMechines(),this.dialogVisible=!1,this.$refs[e].resetFields()})).catch((e=>console.log(e))):(t.id=this.addMechineInfo.id,v(t).then((t=>{this.$notify({title:"成功",message:"更新成功！",type:"success"}),this.selectTypeValue[this.addMechineInfo.id]=this.addMechineInfo.jobId,this.selectTypeValue=n({},this.selectTypeValue),this.dialogVisible=!1,this.$refs[e].resetFields(),this.expandChange({id:this.addMechineInfo.id},{id:this.addMechineInfo.id})})).catch((e=>console.log(e))))}}))},onCancel(e){this.dialogVisible=!1,this.$refs[e].resetFields()},onSearch(){this.doGetMechines()},onPublish(){o().then((e=>{this.$notify({title:"成功",message:"发布成功！",type:"success"})})).catch((e=>{console.log(e)}))},parseTimeSelf(e){var t=new Date(Date.parse(e));return t.toLocaleDateString()+" "+t.toTimeString().split(" ")[0]},doYes(e){var t;(t=e.row.id,s({url:"/v1/machine",method:"delete",params:{id:t}})).then((e=>{this.$notify({title:"成功",message:"删除成功！",type:"success"}),this.doGetMechines()})).catch((e=>{console.log(e)})),this.deleteVisible[e.$index]=!1},doNo(e){this.deleteVisible[e.$index]=!1},doDelete(e){this.deleteVisible[e.$index]=!0},rowStyle:e=>({padding:"0"}),cellStyle:e=>({padding:"0"}),edit(e){this.enterBtnTitle="更新",this.addAndContinueDisabled=!0,this.dialogVisible=!0,this.addMechineInfo={ipAddr:e.row.ipaddr,id:e.row.id},0!==this.selectTypeValue[e.row.id].length?this.addMechineInfo.jobId=this.selectTypeValue[e.row.id]:this.addMechineInfo.jobId=e.row.job_id},expandChange(e,t){if(!t||0===t.length)return;let i=[];this.selectTypeValue[e.id].forEach((e=>{i.push(this.jobsMap[e])})),this.groupNameList[e.id]=i},invocate(e){const t=!this.machines[e.$index].enabled;(function(e){return s({url:"/v1/machine/status",method:"put",data:e})})({id:e.row.id,enabled:t}).then((i=>{this.$notify({title:"成功",message:"更新状态成功！",type:"success"}),this.machines[e.$index].enabled=t,this.machines=[...this.machines]})).catch((e=>console.log(e)))}}},z=_();r("data-v-565c7e16");const M={class:"ipManager-board"},S={class:"do_action"},j={style:{"padding-right":"15px"}},k=y("发布"),T=y("添加"),x=g("span",{type:"warning"},"状态",-1),$=y("编辑"),A=y("禁用"),P=y("启用"),O=g("p",null,"确定删除吗？",-1),N={style:{"text-align":"right",margin:"0"}},D=y("取消"),E=y("确定"),G=y("删除"),U={class:"block"},B=y("创建并继续"),F=y("取消");h();const L=z(((e,t,i,a,l,n)=>{const d=c("el-button"),s=c("el-option"),o=c("el-select"),r=c("el-input"),h=c("el-descriptions-item"),_=c("el-descriptions"),I=c("el-table-column"),v=c("el-tooltip"),w=c("el-tag"),L=c("el-popover"),q=c("el-table"),K=c("el-pagination"),Y=c("el-form-item"),H=c("el-form"),J=c("el-dialog");return p(),u("div",M,[g("div",S,[g("div",j,[g(d,{size:"small",type:"primary",onClick:t[1]||(t[1]=e=>n.onPublish())},{default:z((()=>[k])),_:1}),g(d,{size:"small",type:"success",plain:"",onClick:t[2]||(t[2]=e=>n.doAdd())},{default:z((()=>[T])),_:1})]),g("div",null,[g(r,{size:"small",onKeyup:t[5]||(t[5]=m((e=>n.onSearch()),["enter"])),placeholder:"请输入内容",modelValue:l.searchContent,"onUpdate:modelValue":t[6]||(t[6]=e=>l.searchContent=e),class:"input-with-select"},{prepend:z((()=>[g(o,{class:"searchSelect",modelValue:l.selectOption,"onUpdate:modelValue":t[3]||(t[3]=e=>l.selectOption=e),placeholder:"请选择"},{default:z((()=>[g(s,{label:"IP地址",value:"1"}),g(s,{label:"分组",value:"2"})])),_:1},8,["modelValue"])])),append:z((()=>[g(d,{icon:"el-icon-search",onClick:t[4]||(t[4]=e=>n.onSearch())})])),_:1},8,["modelValue"])])]),g(q,{size:"mini","highlight-current-row":"",border:"",data:l.machines,stripe:"","row-style":n.rowStyle,"cell-style":n.cellStyle,"header-align":"center",onExpandChange:n.expandChange},{default:z((()=>[g(I,{type:"expand"},{default:z((e=>[g(_,{title:"分组列表",size:"mini",column:4,border:""},{default:z((()=>[(p(!0),u(b,null,f(l.groupNameList[e.row.id],((e,t)=>(p(),u(h,{key:t,label:t+1},{default:z((()=>[y(C(e),1)])),_:2},1032,["label"])))),128))])),_:2},1024)])),_:1}),g(I,{label:"序号",width:"50px",align:"center","header-align":"center"},{default:z((e=>[y(C(e.$index+1),1)])),_:1}),g(I,{label:"IP地址",prop:"ipaddr",align:"center","header-align":"center"}),g(I,{label:"分组选项",prop:"job_id",align:"center","header-align":"center"},{default:z((e=>[g(o,{modelValue:l.selectTypeValue[e.row.id],"onUpdate:modelValue":t=>l.selectTypeValue[e.row.id]=t,class:"borderNone","popper-class":"pppselect",onChange:t=>n.tableSelectChange(t,e.$index,e.row),onVisibleChange:t=>n.handleSelect(t,e.$index,e.row),size:"small",multiple:"","collapse-tags":"",placeholder:"请选择"},{default:z((()=>[(p(!0),u(b,null,f(l.jobs,(e=>(p(),u(s,{key:e.id,label:e.name,value:e.id},null,8,["label","value"])))),128))])),_:2},1032,["modelValue","onUpdate:modelValue","onChange","onVisibleChange"])])),_:1}),g(I,{label:"状态",width:"90px",align:"center","header-align":"center"},{header:z((()=>[g(v,{content:"所有机器状态5分钟更新一次",placement:"top"},{default:z((()=>[x])),_:1})])),default:z((({row:e})=>["up"===e.health?(p(),u(v,{key:0,content:"主机正在运行",placement:"top"},{default:z((()=>[g(w,{type:"primary",size:"mini"},{default:z((()=>[y(C(e.health),1)])),_:2},1024)])),_:2},1024)):(p(),u(v,{key:1,content:e.last_error,placement:"top"},{default:z((()=>["down"===e.health?(p(),u(w,{key:0,size:"mini",type:"danger"},{default:z((()=>[y(C(e.health),1)])),_:2},1024)):(p(),u(w,{key:1,type:"info",size:"mini"},{default:z((()=>[y(C(e.health),1)])),_:2},1024))])),_:2},1032,["content"]))])),_:1}),g(I,{label:"最后更新时间",prop:"update_at",align:"center","header-align":"center"},{default:z((({row:e})=>[g("span",null,C(n.parseTimeSelf(e.update_at)),1)])),_:1}),g(I,{label:"操作",align:"center","header-align":"center"},{default:z((e=>[g(d,{type:"primary",plain:"",size:"mini",onClick:t=>n.edit(e)},{default:z((()=>[$])),_:2},1032,["onClick"]),!0===e.row.enabled?(p(),u(d,{key:0,type:"info",plain:"",size:"mini",onClick:t=>n.invocate(e)},{default:z((()=>[A])),_:2},1032,["onClick"])):V("",!0),!1===e.row.enabled?(p(),u(d,{key:1,type:"warning",plain:"",size:"mini",onClick:t=>n.invocate(e)},{default:z((()=>[P])),_:2},1032,["onClick"])):V("",!0),g(L,{visible:l.deleteVisible[e.$index],placement:"top",width:160},{reference:z((()=>[g(d,{size:"mini",type:"danger",plain:"",onClick:t=>n.doDelete(e)},{default:z((()=>[G])),_:2},1032,["onClick"])])),default:z((()=>[O,g("div",N,[g(d,{size:"mini",type:"text",onClick:t=>n.doNo(e)},{default:z((()=>[D])),_:2},1032,["onClick"]),g(d,{type:"primary",size:"mini",onClick:t=>n.doYes(e)},{default:z((()=>[E])),_:2},1032,["onClick"])])])),_:2},1032,["visible"])])),_:1})])),_:1},8,["data","row-style","cell-style","onExpandChange"]),g("div",U,[g(K,{onSizeChange:n.handleSizeChange,onCurrentChange:n.handleCurrentChange,"current-page":l.currentPage,"page-sizes":[15,50,100,200],"page-size":l.pageSize,layout:"total, sizes, prev, pager, next, jumper",total:l.pageTotal},null,8,["onSizeChange","onCurrentChange","current-page","page-size","total"])]),g(J,{title:"新增IP地址",modelValue:l.dialogVisible,"onUpdate:modelValue":t[13]||(t[13]=e=>l.dialogVisible=e),width:"400px",modal:"","before-close":n.handleClose},{default:z((()=>[g("span",null,[g(H,{"label-position":"right",rules:l.rules,ref:"addMechineInfo",model:l.addMechineInfo,"label-width":"90px",size:"small"},{default:z((()=>[g(Y,{label:"IP地址：",prop:"ipAddr"},{default:z((()=>[g(r,{style:{width:"230px"},modelValue:l.addMechineInfo.ipAddr,"onUpdate:modelValue":t[7]||(t[7]=e=>l.addMechineInfo.ipAddr=e),onKeyup:t[8]||(t[8]=m((e=>n.onSubmitAndContinue("addMechineInfo")),["enter"]))},null,8,["modelValue"])])),_:1}),g(Y,{label:"分组名：",prop:"jobId"},{default:z((()=>[g(o,{style:{width:"230px"},modelValue:l.addMechineInfo.jobId,"onUpdate:modelValue":t[9]||(t[9]=e=>l.addMechineInfo.jobId=e),placeholder:"请选择","collapse-tags":"",multiple:""},{default:z((()=>[(p(!0),u(b,null,f(l.jobs,((e,t)=>(p(),u(s,{key:t,label:e.name,value:e.id},null,8,["label","value"])))),128))])),_:1},8,["modelValue"])])),_:1}),g(Y,{size:"small"},{default:z((()=>[g(d,{size:"small",type:"primary",disabled:l.addAndContinueDisabled,onClick:t[10]||(t[10]=e=>n.onSubmitAndContinue("addMechineInfo"))},{default:z((()=>[B])),_:1},8,["disabled"]),g(d,{size:"small",type:"primary",onClick:t[11]||(t[11]=e=>n.onSubmit("addMechineInfo"))},{default:z((()=>[y(C(l.enterBtnTitle),1)])),_:1}),g(d,{size:"small",type:"info",onClick:t[12]||(t[12]=e=>n.onCancel("addMechineInfo"))},{default:z((()=>[F])),_:1})])),_:1})])),_:1},8,["rules","model"])])])),_:1},8,["modelValue","before-close"])])}));w.render=L,w.__scopeId="data-v-565c7e16";export default w;
