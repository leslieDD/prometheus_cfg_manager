var e=Object.defineProperty,a=Object.getOwnPropertySymbols,t=Object.prototype.hasOwnProperty,o=Object.prototype.propertyIsEnumerable,l=(a,t,o)=>t in a?e(a,t,{enumerable:!0,configurable:!0,writable:!0,value:o}):a[t]=o,n=(e,n)=>{for(var r in n||(n={}))t.call(n,r)&&l(e,r,n[r]);if(a)for(var r of a(n))o.call(n,r)&&l(e,r,n[r]);return e};import{s as r}from"./request.f1ddaf6e.js";import{E as s,p as c,j as i,r as p,o as u,c as d,a as h,t as m,l as g,F as _,n as f,k as b}from"./vendor.185638a5.js";import"./index.755618d3.js";const k={data:()=>({privTableData:[],tableRowSpan:{},groupInfo:{},checkBoxSelect:{}}),mounted(){this.current="",this.$route.params.group_name&&(this.groupInfo.group_name=this.$route.params.group_name),this.$route.params.group_name&&(this.groupInfo.group_id=parseInt(this.$route.params.group_id)),this.getPriv(!1)},methods:{objectSpanMethod({row:e,column:a,rowIndex:t,columnIndex:o}){if(0===o){if(this.current!==e.page_name){this.current=e.page_name;return{rowspan:this.tableRowSpan[e.page_name],colspan:1}}return{rowspan:0,colspan:0}}},getPriv(e){(function(e){return r({url:"/v1/manager/user/priv",method:"get",params:e})})(n({},this.groupInfo)).then((a=>{let t={};this.tableRowSpanDone={},a.data.forEach((e=>{t[e.page_name]?t[e.page_name]+=1:t[e.page_name]=1,this.checkBoxSelect[`${e.page_name}_${e.sub_page_name}`]=void 0})),this.tableRowSpan=t,this.privTableData=a.data,e&&this.$notify({title:"成功",message:"重新加载成功！",type:"success"})})).catch((e=>console.log(e)))},rowStyle:e=>({padding:"2px"}),cellStyle(e){let a={padding:"2px"};return(0===e.columnIndex||1===e.columnIndex)&&(a["background-color"]="#F1FAFA"),a},checkboxChange(e,a){e?a.func_list.forEach((e=>{e.checked=!0})):a.func_list.forEach((e=>{e.checked=!1}))},allCheckBoxChange(){console.log("allCheckBoxChange")},renderHeader:({column:e,$index:a})=>s("span",[s("el-checkbox",{props:{}}),s("span",e.label)]),selectAll(e){let a;a=!!e,this.privTableData.forEach((e=>{e.func_list.forEach((e=>{e.checked=a}))}));for(const t in this.checkBoxSelect)this.checkBoxSelect[t]=a},save(){(function(e,a){return r({url:"/v1/manager/user/priv",method:"put",params:e,data:a})})(n({},this.groupInfo),[...this.privTableData]).then((e=>{this.$notify({title:"成功",message:"更新成功！",type:"success"})})).catch((e=>console.log(e)))},reload(){this.getPriv(!0)},goBack(){this.$router.push({name:"group"})}}},v=b();c("data-v-00ec0c28");const y={class:"priv-board"},x={class:"sub-page-class"},w=h("span",null,m("权限列表"),-1),S=g("全选"),$=g("全不选"),C=g("保存"),I=g("重新加载"),j=g("返回"),B={class:"show-group-name"};i();const V=v(((e,a,t,o,l,n)=>{const r=p("el-table-column"),s=p("el-checkbox"),c=p("el-divider"),i=p("el-button"),b=p("el-tag"),k=p("el-table"),V=p("el-scrollbar");return u(),d("div",y,[h(V,null,{default:v((()=>[h(k,{size:"mini",data:l.privTableData,"span-method":n.objectSpanMethod,border:"",style:{width:"100%"},"row-style":n.rowStyle,"cell-style":n.cellStyle},{default:v((()=>[h(r,{prop:"page_nice_name",label:"页面",width:"120px"}),h(r,{prop:"sub_page_nice_name",label:"子页面",width:"120px"},{default:v((({row:e})=>[h("div",x,[h("span",null,m(e.sub_page_nice_name+" "),1),h(s,{modelValue:l.checkBoxSelect[`${e.page_name}_${e.sub_page_name}`],"onUpdate:modelValue":a=>l.checkBoxSelect[`${e.page_name}_${e.sub_page_name}`]=a,onChange:a=>n.checkboxChange(a,e)},null,8,["modelValue","onUpdate:modelValue","onChange"])])])),_:1}),h(r,{label:"权限列表","scoped-slot":""},{header:v((()=>[w,h(c,{direction:"vertical"}),h(i,{size:"mini",onClick:a[1]||(a[1]=e=>n.selectAll(!0))},{default:v((()=>[S])),_:1}),h(i,{size:"mini",onClick:a[2]||(a[2]=e=>n.selectAll(!1))},{default:v((()=>[$])),_:1}),h(i,{size:"mini",type:"primary",onClick:a[3]||(a[3]=e=>n.save())},{default:v((()=>[C])),_:1}),h(i,{size:"mini",type:"warning",onClick:a[4]||(a[4]=e=>n.reload())},{default:v((()=>[I])),_:1}),h(i,{size:"mini",type:"info",onClick:a[5]||(a[5]=e=>n.goBack())},{default:v((()=>[j])),_:1}),h("span",B,[h(b,{size:"small"},{default:v((()=>[g("当前组名："+m(l.groupInfo.group_name),1)])),_:1})])])),default:v((({row:e})=>[(u(!0),d(_,null,f(e.func_list,(e=>(u(),d(s,{key:e.id,modelValue:e.checked,"onUpdate:modelValue":a=>e.checked=a},{default:v((()=>[g(m(e.func_nice_name),1)])),_:2},1032,["modelValue","onUpdate:modelValue"])))),128))])),_:1})])),_:1},8,["data","span-method","row-style","cell-style"])])),_:1})])}));k.render=V,k.__scopeId="data-v-00ec0c28";export default k;
