import{_ as e,b as a}from"./index.6f2c4da2.js";import{p as l,j as s,r as n,o,c as t,a as i,T as d,u as r,k as u,l as c}from"./vendor.684f159e.js";const p={name:"Admin",data:()=>({}),mounted(){this.$router.push({name:"user"})},methods:{handleOpen(e,a){console.log(e,a)},handleClose(e,a){console.log(e,a)},handleSelect(e,a){"8"===e?this.$router.push({name:"user"}):"9"===e?this.$router.push({name:"group"}):"10"===e?this.$router.push({name:"adminSet"}):"11"===e&&this.$router.push({name:"privileges"})},goBackPerson(){this.$router.push({name:"person"})}}},m=u();l("data-v-90388900");const h={class:"admin-board"},f={class:"person-center-btn"},b=c("个人中心"),g=c(" 权限管理"),v={class:"base-config-board"},_={class:"base-config-menu"},C=i("i",{class:"el-icon-potato-strips"},null,-1),k=c("用户管理"),$=i("i",{class:"el-icon-milk-tea"},null,-1),x=c("组管理"),S=i("i",{class:"el-icon-lollipop"},null,-1),O=c("设置"),j={class:"base-config-router"};s();const w=m(((l,s,u,c,p,w)=>{const y=e,B=n("el-button"),P=a,z=n("el-menu-item"),A=n("el-menu"),I=n("router-view"),T=n("el-tab-pane"),q=n("el-tabs");return o(),t("div",h,[i("div",f,[i("div",null,[i(B,{size:"mini",type:"warning",onClick:w.goBackPerson},{default:m((()=>[i(y,{"icon-class":"pig"}),b])),_:1},8,["onClick"])])]),i(q,{type:"border-card"},{default:m((()=>[i(T,{label:"权限管理"},{label:m((()=>[i("span",null,[i(P,{"icon-class":"promgtshuiguo9"}),g])])),default:m((()=>[i("div",v,[i("div",_,[i(A,{"default-active":"8",class:"el-menu-vertical-demo",onOpen:w.handleOpen,onClose:w.handleClose,onSelect:w.handleSelect},{default:m((()=>[i(z,{index:"8"},{title:m((()=>[k])),default:m((()=>[C])),_:1}),i(z,{index:"9"},{title:m((()=>[x])),default:m((()=>[$])),_:1}),i(z,{index:"10"},{title:m((()=>[O])),default:m((()=>[S])),_:1})])),_:1},8,["onOpen","onClose","onSelect"])]),i("div",j,[i(I,null,{default:m((({Component:e})=>[i(d,{name:"slide-fade"},{default:m((()=>[(o(),t(r(e)))])),_:2},1024)])),_:1})])])])),_:1})])),_:1})])}));p.render=w,p.__scopeId="data-v-90388900";export default p;
