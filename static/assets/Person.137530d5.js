var e=Object.defineProperty,t=Object.getOwnPropertySymbols,a=Object.prototype.hasOwnProperty,s=Object.prototype.propertyIsEnumerable,l=(t,a,s)=>a in t?e(t,a,{enumerable:!0,configurable:!0,writable:!0,value:s}):t[a]=s;import{r as o,_ as r}from"./index.6f2c4da2.js";import{s as i}from"./request.c24bac87.js";import{p as n,j as d,r as u,o as m,c,a as p,l as h,t as g,F as f,n as _,q as v,k as w,s as y}from"./vendor.684f159e.js";const b={data(){return{src:"/imgs/xiaohei.jpg",uname:"Linux 4.15.0-150-generic Unknow SMP Sat Jul 3 13:37:31 UTC 2021 x86_64 x86_64 x86_64 GNU/Linux",programmer_said:["选择你能够承担的，承担你已经选择的"],sayWhat:[],timer:null,userInfo:{phone:"",group_name:"",register_time:"",login_time:""},dialogVisible:!1,ruleForm:{old_pwd:"",new_pwd1:"",new_pwd2:""},rules:{old_pwd:[{required:!0,message:"请输入旧密码",trigger:"blur"}],new_pwd1:[{required:!0,message:"请输入新密码",trigger:"blur"}],new_pwd2:[{required:!0,validator:(e,t,a)=>{""===t?a(new Error("请再次输入密码")):t!==this.ruleForm.new_pwd1?a(new Error("两次输入密码不一致!")):a()},trigger:"blur"}]},userID:0,menuShow:{show_menu_prometheus_cfg_manager:!1,show_menu_administrator_cfg_manager:!1}}},mounted(){this.initUserInfo(),this.goMenu(),this.switchSay(),null===this.timer&&this.setTimeer(),this.loadProgramerSay(this)},beforeUnmount(){clearInterval(this.timer),this.timer=null},methods:{initUserInfo(){if(this.$store.getters.userInfo){this.userID=this.$store.getters.userID;const e=this.$store.getters.userInfo;this.userInfo={id:this.$store.getters.userID,phone:e.phone,group_name:e.group_name,register_time:e.create_at,login_time:e.update_at}}},goMenu(){var e;i({url:"/v1/manager/user/menu/priv",method:"get",params:e}).then((e=>{this.menuShow=e.data})).catch((e=>{console.log(e)}))},goToConfig(){this.$router.push({name:"menu"})},goToAdmin(){this.$router.push({name:"admin"})},loadProgramerSay(){i({url:"/v1/txt/programer/say",method:"get"}).then((e=>{this.programmer_said=e.data.programer_say,this.uname=e.data.uname,this.switchSay()})).catch((e=>console.log(e)))},setTimeer(){this.timer=setInterval((()=>{this.switchSay()}),15e3)},switchSay(){const e=this.programmer_said[Math.floor(Math.random()*this.programmer_said.length)];this.sayWhat=[e]},fixedLengthFormatString(e,t){if(null==e||null==e)return null;if(!/^[0-9]*[1-9][0-9]*$/.test(t))return null;for(var a=new Array,s=e.length,l=0;l<s/t;l++)(l+1)*t>s?a.push(e.substring(l*t,s)):a.push(e.substring(l*t,(l+1)*t));return a},parseTimeSelf(e){var t=new Date(Date.parse(e));return t.toLocaleDateString()+" "+t.toTimeString().split(" ")[0]},changePassword(){this.dialogVisible=!0},quitLog(){this.$confirm("退出账号, 是否继续?","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning"}).then((()=>{i({url:"/v1/logout",method:"post"}).then((e=>{this.$notify({title:"成功",message:"退出成功！",type:"success"}),this.$store.dispatch("resetToken"),o(),this.$router.push({name:"login"})}))})).catch((()=>{this.$notify({title:"消息",message:"取消退出",type:"info"})}))},handleClose(){this.dialogVisible=!1},submitChangePassword(e){this.$refs[e].validate((e=>{if(e){const e=((e,o)=>{for(var r in o||(o={}))a.call(o,r)&&l(e,r,o[r]);if(t)for(var r of t(o))s.call(o,r)&&l(e,r,o[r]);return e})({},this.ruleForm);(o=e,i({url:"/v1/manager/user/password",method:"post",data:o})).then((e=>{this.$notify({title:"成功",message:"密码更新成功！",type:"success"}),this.dialogVisible=!1})).catch((e=>console.log(e)))}var o}))}}},x=w();n("data-v-58f98e3a");const k={class:"person-board"},S={id:"box"},I={class:"meBox"},V=p("div",{class:"headPhoto"},null,-1),C=p("div",{class:"meBox-title"},[p("p",null,"I'm DianDian"),p("div",{class:"fg"})],-1),T={class:"meBox-text"},F={class:"meBox-Button"},$={id:"cmdBox"},j={class:"cmd"},P=y('<div class="click" data-v-58f98e3a><div class="red" data-v-58f98e3a></div><div class="yellow" data-v-58f98e3a></div><div class="green" data-v-58f98e3a></div></div><div class="title" data-v-58f98e3a><span data-v-58f98e3a>root - bash</span></div>',2),q={class:"cmdText"},U=p("span",{style:{color:"rgb(0, 190, 0)"}},"root@monitor:~",-1),D=p("span",{style:{color:"blue"}},"# ",-1),B=p("span",{style:{color:"rgb(39, 39, 39)"}},"./tianqi.sh",-1),z=p("br",null,null,-1),O=p("iframe",{scrolling:"no",src:"https://tianqiapi.com/api.php?style=tc&skin=pitaya",frameborder:"0",width:"350",height:"24",allowtransparency:"true"},null,-1),L=p("br",null,null,-1),M=p("span",{style:{color:"rgb(0, 190, 0)"}},"root@monitor:~",-1),A=p("span",{style:{color:"blue"}},"# ",-1),E=p("span",{style:{color:"rgb(39, 39, 39)"}},"uname",-1),W=p("br",null,null,-1),G=p("span",{style:{color:"rgb(0, 190, 0)"}},"root@monitor:~",-1),J=p("span",{style:{color:"blue"}},"# ",-1),N=p("span",{style:{color:"rgb(39, 39, 39)"}},"cat ./quotes_by_famous_people.txt",-1),H=p("br",null,null,-1),K=p("span",{style:{color:"rgb(0, 190, 0)"}},"root@monitor:~",-1),Q=p("span",{style:{color:"blue"}},"# ",-1),R=p("span",{style:{color:"rgb(39, 39, 39)"}},"sudo rm -rf /过去的自己/*",-1),X={class:"cmd cmd2"},Y=y('<div class="click" data-v-58f98e3a><div class="red" data-v-58f98e3a></div><div class="yellow" data-v-58f98e3a></div><div class="green" data-v-58f98e3a></div></div><div class="title" data-v-58f98e3a><span data-v-58f98e3a>root - bash</span></div>',2),Z={class:"cmdText"},ee=p("span",{style:{color:"rgb(0, 190, 0)"}},"root@monitor:~",-1),te=p("span",{style:{color:"blue"}},"# ",-1),ae=p("span",{style:{color:"rgb(39, 39, 39)"}},"awk '/menu/' ./prometheus.yml",-1),se=p("p",{class:"prompt-text"},"点击下面菜单进入：",-1),le={key:0},oe=h("Prometheus配置"),re={key:1,class:"menu_zero_padding"},ie=h("用户及权限管理"),ne={key:2},de={class:"no_menu_show"},ue=h("没有可供显示的菜单"),me=p("span",{style:{color:"rgb(0, 190, 0)"}},"root@monitor:~",-1),ce=p("span",{style:{color:"blue"}},"# ",-1),pe=h("取 消"),he=h("确 定");d();const ge=x(((e,t,a,s,l,o)=>{const i=u("el-descriptions-item"),n=u("el-descriptions"),d=r,w=u("el-button"),y=u("el-tag"),b=u("el-input"),ge=u("el-form-item"),fe=u("el-form"),_e=u("el-dialog");return m(),c("div",k,[p("div",S,[p("div",I,[V,C,p("div",T,[p(n,{title:"",column:1,size:""},{default:x((()=>[p(i,{label:"手机号"},{default:x((()=>[h(g(l.userInfo.phone),1)])),_:1}),p(i,{label:"所属组"},{default:x((()=>[h(g(l.userInfo.group_name),1)])),_:1}),p(i,{label:"注册时间"},{default:x((()=>[h(g(o.parseTimeSelf(l.userInfo.register_time)),1)])),_:1}),p(i,{label:"登录时间"},{default:x((()=>[h(g(o.parseTimeSelf(l.userInfo.login_time)),1)])),_:1})])),_:1})]),p("div",F,[p("a",{href:"javascript:void(0)",onClick:t[1]||(t[1]=e=>o.changePassword())},"更改密码"),p("a",{href:"javascript:void(0)",onClick:t[2]||(t[2]=e=>o.quitLog())},"退出登录")])]),p("div",$,[p("div",j,[P,p("div",q,[U,D,B,z,O,L,M,A,E,p("p",null,g(l.uname),1),W,G,J,N,(m(!0),c(f,null,_(l.sayWhat,((e,t)=>(m(),c("p",{key:t},g(e),1)))),128)),H,K,Q,R])]),p("div",X,[Y,p("div",Z,[ee,te,ae,se,p("dl",null,[!0===l.menuShow.show_menu_prometheus_cfg_manager?(m(),c("dt",le,[p(w,{type:"text",onClick:o.goToConfig},{default:x((()=>[p(d,{"icon-class":"xianrenzhang"}),oe])),_:1},8,["onClick"])])):v("",!0),!0===l.menuShow.show_menu_administrator_cfg_manager?(m(),c("dt",re,[p(d,{"icon-class":"yezhi"}),p(w,{type:"text",onClick:o.goToAdmin},{default:x((()=>[ie])),_:1},8,["onClick"])])):v("",!0),0===Object.keys(l.menuShow).length?(m(),c("dt",ne,[p("div",de,[p(y,{type:"warning",size:"small"},{default:x((()=>[p(d,{"icon-class":"shuzhi"}),ue])),_:1})])])):v("",!0)]),me,ce])])])]),p(_e,{title:"更新我的密码",modelValue:l.dialogVisible,"onUpdate:modelValue":t[7]||(t[7]=e=>l.dialogVisible=e),width:"450px","before-close":o.handleClose},{default:x((()=>[p(fe,{model:l.ruleForm,rules:l.rules,ref:"ruleForm","label-width":"auto",class:"demo-ruleForm",size:"small"},{default:x((()=>[p(ge,{label:"请输出旧密码",prop:"old_pwd"},{default:x((()=>[p(b,{placeholder:"请输入旧密码",modelValue:l.ruleForm.old_pwd,"onUpdate:modelValue":t[3]||(t[3]=e=>l.ruleForm.old_pwd=e),type:"password"},null,8,["modelValue"])])),_:1}),p(ge,{label:"新密码",prop:"new_pwd1"},{default:x((()=>[p(b,{placeholder:"请输入新密码",modelValue:l.ruleForm.new_pwd1,"onUpdate:modelValue":t[4]||(t[4]=e=>l.ruleForm.new_pwd1=e),type:"password"},null,8,["modelValue"])])),_:1}),p(ge,{label:"再次输入新密码",prop:"new_pwd2"},{default:x((()=>[p(b,{placeholder:"请输入新密码",modelValue:l.ruleForm.new_pwd2,"onUpdate:modelValue":t[5]||(t[5]=e=>l.ruleForm.new_pwd2=e),type:"password"},null,8,["modelValue"])])),_:1}),p(ge,null,{default:x((()=>[p(w,{onClick:o.handleClose},{default:x((()=>[pe])),_:1},8,["onClick"]),p(w,{type:"primary",onClick:t[6]||(t[6]=e=>o.submitChangePassword("ruleForm"))},{default:x((()=>[he])),_:1})])),_:1})])),_:1},8,["model","rules"])])),_:1},8,["modelValue","before-close"])])}));b.render=ge,b.__scopeId="data-v-58f98e3a";export default b;