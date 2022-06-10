import{c as e,f as t,h as a,i,j as l,k as s,l as o,u as n}from"./manager.e529774e.js";import{p as r,j as d,r as u,o as p,c as h,a as c,w as g,l as m,t as f,q as b,k as C}from"./vendor.02dbbba7.js";import"./request.d1dca4b5.js";import"./index.a2263c8a.js";const y={name:"ManagerGroup",data:()=>({value:"",ManagerGroup:[],ManagerGroupRef:{id:0,name:"",user_count:0},editCodeButVisable:{},pageSize:20,pageTotal:0,currentPage:1,searchContent:"",deleteVisible:{},dialogVisible:!1,buttonTitle:"",dialogTitle:"",rules:{name:[{required:!0,message:"请输入正确的名称",validator:function(e,t,a){if(""===t||void 0===t||null==t)a(new Error("请输入正确名称"));else{/\s+/.test(t)?a(new Error("名称中不允许有空字符")):a()}},trigger:["blur"]}]},paginationShow:!0,editObject:"",editUserVisible:!1,usersListData:[],transferChanged:!1,usersListValue:[],groupInfo:{}}),mounted(){this.$route.params.group_name&&(this.searchContent=this.$route.params.group_name);const e=this.$store.getters.getPageInfo;this.pageSize=e.pageSize,this.pageNo=e.pageNo,this.doGetManagerGroup()},methods:{doAdd(){this.ManagerGroupRef={id:0,name:""},this.buttonTitle="创建",this.dialogTitle="增加组",this.dialogVisible=!0},doGetManagerGroup(t){t||(t={pageNo:this.currentPage,pageSize:this.pageSize,search:this.searchContent}),e(t).then((e=>{this.ManagerGroup=e.data.data,this.pageTotal=e.data.totalCount,this.currentPage=e.data.pageNo,this.pageSize=e.data.pageSize,this.paginationShow=!1,this.$nextTick((()=>{this.paginationShow=!0}))})).catch((e=>{console.log(e)}))},doEdit(e){this.buttonTitle="更新",this.dialogTitle="编辑组名",this.ManagerGroupRef.id=e.row.id,this.ManagerGroupRef.name=e.row.name,this.dialogVisible=!0},doUserList(e){this.editObject=e.row.name;let i=[];t().then((t=>{t.data.forEach((e=>{i.push({label:e.username,key:e.id,spell:e.username,disabled:!1})}));let l=[];this.groupInfo={group_name:e.row.name,group_id:e.row.id},a(this.groupInfo).then((e=>{e.data.forEach((e=>{l.push(e.id)})),this.usersListData=i,this.usersListValue=l,this.transferChanged=!1,this.editUserVisible=!0})).catch((e=>console.log(e)))})).catch((e=>console.log(e)))},handleSizeChange(e){let t={pageNo:this.currentPage,pageSize:e,search:this.searchContent};this.doGetManagerGroup(t)},handleCurrentChange(e){let t={pageNo:e,pageSize:this.pageSize,search:this.searchContent};this.doGetManagerGroup(t)},handleClose(e){this.dialogVisible=!1},onSubmit(e){this.$refs[e].validate((t=>{if(!t)return!1;{let t={};t.name=this.ManagerGroupRef.name,"创建"===this.buttonTitle||"create"===this.buttonTitle?i(t).then((t=>{this.$notify({title:"成功",message:"创建成功！",type:"success"}),this.doGetManagerGroup(),this.dialogVisible=!1,this.$refs[e].resetFields()})).catch((e=>{console.log(e)})):(t.id=this.ManagerGroupRef.id,l(t).then((t=>{this.$notify({title:"成功",message:"更新成功！",type:"success"}),this.doGetManagerGroup(),this.dialogVisible=!1,this.$refs[e].resetFields()})).catch((e=>{console.log(e)})))}}))},onCancel(e){this.dialogVisible=!1,this.$refs[e].resetFields()},onSearch(){this.doGetManagerGroup()},parseTimeSelf(e){var t=new Date(Date.parse(e));return t.toLocaleDateString()+" "+t.toTimeString().split(" ")[0]},doYes(e){s({id:e.row.id}).then((e=>{this.$notify({title:"成功",message:"删除成功！",type:"success"}),this.doGetManagerGroup()})).catch((e=>{console.log(e)})),this.deleteVisible[e.$index]=!1},doNo(e){this.deleteVisible[e.$index]=!1},doDelete(e){this.deleteVisible[e.$index]=!0},rowStyle:e=>({padding:"0"}),cellStyle:e=>({padding:"0"}),invocate(e){const t=!this.ManagerGroup[e.$index].enabled,a={id:e.row.id,enabled:t};o(a).then((a=>{this.$notify({title:"成功",message:"更新状态成功！",type:"success"}),this.ManagerGroup[e.$index].enabled=t,this.ManagerGroup=[...this.ManagerGroup]})).catch((e=>console.log(e)))},doPriv(e){this.$store.dispatch("setPageInfo",this.pageSize,this.pageNo),this.$router.push({name:"privileges",params:{group_name:e.row.name,group_id:e.row.id}})},routeToUser(e){this.$router.push({name:"user",params:{group_name:e.name}})},editUserClose(){this.editUserVisible=!1,this.transferChanged=!1},pushNewUserList(){const e=[...this.usersListValue];n(this.groupInfo,e).then((e=>{this.$notify({title:"成功",message:"更新组成员成功！",type:"success"}),this.doGetManagerGroup()})).catch((e=>console.log(e)))},filterUserMethod:(e,t)=>t.spell.indexOf(e)>-1,transferChange(e,t,a){this.transferChanged=!0}}},w=C();r("data-v-6c40452e");const V={class:"group-board"},_={class:"main-content"},z={class:"do_action"},G={style:{"padding-right":"15px"}},v=m("添加组"),k={class:"table-show"},S={class:"actioneara"},M=m("编辑"),x=m("用户列表"),$=m("禁用"),U=m("启用"),T=m("权限"),L=c("p",null,"确定删除吗？",-1),R={style:{"text-align":"right",margin:"0"}},j=m("取消"),N=m("确定"),P=m("删除"),D={key:0,class:"block"},I=m("取消"),E={class:"user-list-push-box"},B=m("关闭"),O=m("提交");d();const q=w(((e,t,a,i,l,s)=>{const o=u("el-button"),n=u("el-input"),r=u("el-table-column"),d=u("el-popover"),C=u("el-table"),y=u("el-pagination"),q=u("el-scrollbar"),F=u("el-form-item"),A=u("el-form"),Y=u("el-dialog"),K=u("el-transfer"),H=u("el-row");return p(),h("div",V,[c("div",_,[c("div",null,[c("div",z,[c("div",G,[c(o,{size:"small",type:"success",plain:"",onClick:t[1]||(t[1]=e=>s.doAdd())},{default:w((()=>[v])),_:1})]),c("div",null,[c(n,{size:"small",placeholder:"请输入内容",onKeyup:t[3]||(t[3]=g((e=>s.onSearch()),["enter"])),modelValue:l.searchContent,"onUpdate:modelValue":t[4]||(t[4]=e=>l.searchContent=e)},{append:w((()=>[c(o,{size:"small",onClick:t[2]||(t[2]=e=>s.onSearch()),icon:"el-icon-search"})])),_:1},8,["modelValue"])])]),c("div",k,[c(q,{height:"75vh",class:"flex-content"},{default:w((()=>[c(C,{size:"mini","highlight-current-row":"",border:"",data:l.ManagerGroup,stripe:"","row-style":s.rowStyle,"cell-style":s.cellStyle},{default:w((()=>[c(r,{label:"序号",width:"50px"},{default:w((e=>[m(f(e.$index+1),1)])),_:1}),c(r,{label:"组名",prop:"name"}),c(r,{label:"用户数",prop:"user_count",width:"150px"},{default:w((({row:e})=>[c(o,{size:"mini",onClick:t=>s.routeToUser(e),type:"text"},{default:w((()=>[m(f(e.user_count),1)])),_:2},1032,["onClick"])])),_:1}),c(r,{label:"更新账号",align:"center",prop:"update_by"}),c(r,{label:"更新时间",width:"150px",prop:"update_at"},{default:w((({row:e})=>[c("span",null,f(s.parseTimeSelf(e.update_at)),1)])),_:1}),c(r,{label:"操作",align:"center",width:"350px"},{default:w((e=>[c("div",S,[c("div",null,[c(o,{size:"mini",type:"primary",onClick:t=>s.doEdit(e),plain:"",disabled:l.editCodeButVisable[e.row.id]},{default:w((()=>[M])),_:2},1032,["onClick","disabled"])]),c("div",null,[c(o,{size:"mini",type:"primary",onClick:t=>s.doUserList(e),plain:""},{default:w((()=>[x])),_:2},1032,["onClick"])]),c("div",null,[!0===e.row.enabled?(p(),h(o,{key:0,size:"mini",type:"info",onClick:t=>s.invocate(e),plain:"",disabled:l.editCodeButVisable[e.row.id]},{default:w((()=>[$])),_:2},1032,["onClick","disabled"])):b("",!0),!1===e.row.enabled?(p(),h(o,{key:1,size:"mini",type:"warning",onClick:t=>s.invocate(e),plain:"",disabled:l.editCodeButVisable[e.row.id]},{default:w((()=>[U])),_:2},1032,["onClick","disabled"])):b("",!0)]),c("div",null,[c(o,{size:"mini",type:"primary",onClick:t=>s.doPriv(e),plain:""},{default:w((()=>[T])),_:2},1032,["onClick"])]),c("div",null,[c(d,{visible:l.deleteVisible[e.$index],placement:"top"},{reference:w((()=>[c(o,{size:"mini",type:"danger",plain:"",onClick:t=>s.doDelete(e),disabled:l.editCodeButVisable[e.row.id]},{default:w((()=>[P])),_:2},1032,["onClick","disabled"])])),default:w((()=>[L,c("div",R,[c(o,{size:"mini",type:"text",onClick:t=>s.doNo(e)},{default:w((()=>[j])),_:2},1032,["onClick"]),c(o,{type:"primary",size:"mini",onClick:t=>s.doYes(e)},{default:w((()=>[N])),_:2},1032,["onClick"])])])),_:2},1032,["visible"])])])])),_:1})])),_:1},8,["data","row-style","cell-style"]),l.paginationShow?(p(),h("div",D,[c(y,{onSizeChange:s.handleSizeChange,onCurrentChange:s.handleCurrentChange,"current-page":l.currentPage,"page-sizes":[20,30,50],"page-size":l.pageSize,layout:"total, sizes, prev, pager, next, jumper",total:l.pageTotal},null,8,["onSizeChange","onCurrentChange","current-page","page-size","total"])])):b("",!0)])),_:1})])])]),c(Y,{title:l.dialogTitle,modelValue:l.dialogVisible,"onUpdate:modelValue":t[8]||(t[8]=e=>l.dialogVisible=e),width:"400px",modal:"","before-close":s.handleClose},{default:w((()=>[c("span",null,[c(A,{"label-position":"right",rules:l.rules,ref:"ManagerGroupRef",model:l.ManagerGroupRef,"label-width":"auto",size:"small"},{default:w((()=>[c(F,{label:"组名：",prop:"name"},{default:w((()=>[c(n,{style:{width:"280px"},modelValue:l.ManagerGroupRef.name,"onUpdate:modelValue":t[5]||(t[5]=e=>l.ManagerGroupRef.name=e)},null,8,["modelValue"])])),_:1}),c(F,{size:"small"},{default:w((()=>[c(o,{size:"small",type:"primary",onClick:t[6]||(t[6]=e=>s.onSubmit("ManagerGroupRef"))},{default:w((()=>[m(f(l.buttonTitle),1)])),_:1}),c(o,{size:"small",type:"info",onClick:t[7]||(t[7]=e=>s.onCancel("ManagerGroupRef"))},{default:w((()=>[I])),_:1})])),_:1})])),_:1},8,["rules","model"])])])),_:1},8,["title","modelValue","before-close"]),c(Y,{title:"编辑组成员："+l.editObject,modelValue:l.editUserVisible,"onUpdate:modelValue":t[10]||(t[10]=e=>l.editUserVisible=e),width:"700px",modal:"","before-close":s.editUserClose},{default:w((()=>[c(H,{type:"flex",align:"middle",justify:"center"},{default:w((()=>[c(K,{modelValue:l.usersListValue,"onUpdate:modelValue":t[9]||(t[9]=e=>l.usersListValue=e),filterable:"","filter-method":s.filterUserMethod,"filter-placeholder":"请输入关键字",data:l.usersListData,onChange:s.transferChange,titles:["用户列表","组成员"]},null,8,["modelValue","filter-method","data","onChange"])])),_:1}),c("div",E,[c(o,{class:"user-list-close-btn",type:"info",size:"small",onClick:s.editUserClose},{default:w((()=>[B])),_:1},8,["onClick"]),c(o,{class:"user-list-push-btn",type:"warning",size:"small",onClick:s.pushNewUserList},{default:w((()=>[O])),_:1},8,["onClick"])])])),_:1},8,["title","modelValue","before-close"])])}));y.render=q,y.__scopeId="data-v-6c40452e";export default y;
