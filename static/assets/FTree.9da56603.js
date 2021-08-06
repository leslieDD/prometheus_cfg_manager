import{c as t,C as e}from"./lint.624d88ac.js";/* empty css                */import{l as r,a as n}from"./publish.e4ef820f.js";import{y as i,p as s,j as o,r as a,o as l,c,a as h,k as u}from"./vendor.684f159e.js";import"./request.c24bac87.js";import"./index.6f2c4da2.js";var f;(f=t.exports).registerHelper("lint","json",(function(t){var e=[];if(!window.jsonlint)return window.console&&window.console.error("Error: window.jsonlint not defined, CodeMirror JSON linting cannot run."),e;var r=window.jsonlint.parser||window.jsonlint;r.parseError=function(t,r){var n=r.loc;e.push({from:f.Pos(n.first_line-1,n.first_column),to:f.Pos(n.last_line-1,n.last_column),message:t})};try{r.parse(t)}catch(n){}return e}));var p,d,y,g,m,x,b,_,v,E,w,S,k,j,N,O,$,A,I={exports:{}};p=I,k=!0,j={},x={'"':'"',"\\":"\\","/":"/",b:"\b",f:"\f",n:"\n",r:"\r",t:"\t"},b=function(t){throw{name:"SyntaxError",message:t,at:d,text:g}},_=function(t){return t&&t!==y&&b("Expected '"+t+"' instead of '"+y+"'"),y=g.charAt(d),d+=1,y},v=function(){var t,e="";for("-"===y&&(e="-",_("-"));y>="0"&&y<="9";)e+=y,_();if("."===y)for(e+=".";_()&&y>="0"&&y<="9";)e+=y;if("e"===y||"E"===y)for(e+=y,_(),"-"!==y&&"+"!==y||(e+=y,_());y>="0"&&y<="9";)e+=y,_();if(t=+e,isFinite(t))return t;b("Bad number")},E=function(){var t,e,r,n="";if('"'===y)for(;_();){if('"'===y)return _(),n;if("\\"===y)if(_(),"u"===y){for(r=0,e=0;e<4&&(t=parseInt(_(),16),isFinite(t));e+=1)r=16*r+t;n+=String.fromCharCode(r)}else{if("string"!=typeof x[y])break;n+=x[y]}else n+=y}b("Bad string")},w=function(){for(;y&&y<=" ";)_()},m=function(){switch(w(),y){case"{":return function(){var t,e={};if("{"===y){if(_("{"),w(),"}"===y)return _("}"),e;for(;y;){if(t=E(),w(),_(":"),Object.hasOwnProperty.call(e,t)&&b("Duplicate key '"+t+"'"),e[t]=m(),w(),"}"===y)return _("}"),e;_(","),w()}}b("Bad object")}();case"[":return function(){var t=[];if("["===y){if(_("["),w(),"]"===y)return _("]"),t;for(;y;){if(t.push(m()),w(),"]"===y)return _("]"),t;_(","),w()}}b("Bad array")}();case'"':return E();case"-":return v();default:return y>="0"&&y<="9"?v():function(){switch(y){case"t":return _("t"),_("r"),_("u"),_("e"),!0;case"f":return _("f"),_("a"),_("l"),_("s"),_("e"),!1;case"n":return _("n"),_("u"),_("l"),_("l"),null}b("Unexpected '"+y+"'")}()}},N=function(t,e){var r;return g=t,d=0,y=" ",r=m(),w(),y&&b("Syntax error"),"function"==typeof e?function t(r,n){var i,s,o=r[n];if(o&&"object"==typeof o)for(i in o)Object.prototype.hasOwnProperty.call(o,i)&&(void 0!==(s=t(o,i))?o[i]=s:delete o[i]);return e.call(r,n,o)}({"":r},""):r},(S={trace:function(){},yy:{},symbols_:{error:2,JSONString:3,STRING:4,JSONNumber:5,NUMBER:6,JSONNullLiteral:7,NULL:8,JSONBooleanLiteral:9,TRUE:10,FALSE:11,JSONText:12,JSONValue:13,EOF:14,JSONObject:15,JSONArray:16,"{":17,"}":18,JSONMemberList:19,JSONMember:20,":":21,",":22,"[":23,"]":24,JSONElementList:25,$accept:0,$end:1},terminals_:{2:"error",4:"STRING",6:"NUMBER",8:"NULL",10:"TRUE",11:"FALSE",14:"EOF",17:"{",18:"}",21:":",22:",",23:"[",24:"]"},productions_:[0,[3,1],[5,1],[7,1],[9,1],[9,1],[12,2],[13,1],[13,1],[13,1],[13,1],[13,1],[13,1],[15,2],[15,3],[20,3],[19,1],[19,3],[16,2],[16,3],[25,1],[25,3]],performAction:function(t,e,r,n,i,s,o){var a=s.length-1;switch(i){case 1:this.$=t.replace(/\\(\\|")/g,"$1").replace(/\\n/g,"\n").replace(/\\r/g,"\r").replace(/\\t/g,"\t").replace(/\\v/g,"\v").replace(/\\f/g,"\f").replace(/\\b/g,"\b");break;case 2:this.$=Number(t);break;case 3:this.$=null;break;case 4:this.$=!0;break;case 5:this.$=!1;break;case 6:return this.$=s[a-1];case 13:this.$={};break;case 14:this.$=s[a-1];break;case 15:this.$=[s[a-2],s[a]];break;case 16:this.$={},this.$[s[a][0]]=s[a][1];break;case 17:this.$=s[a-2],s[a-2][s[a][0]]=s[a][1];break;case 18:this.$=[];break;case 19:this.$=s[a-1];break;case 20:this.$=[s[a]];break;case 21:this.$=s[a-2],s[a-2].push(s[a])}},table:[{3:5,4:[1,12],5:6,6:[1,13],7:3,8:[1,9],9:4,10:[1,10],11:[1,11],12:1,13:2,15:7,16:8,17:[1,14],23:[1,15]},{1:[3]},{14:[1,16]},{14:[2,7],18:[2,7],22:[2,7],24:[2,7]},{14:[2,8],18:[2,8],22:[2,8],24:[2,8]},{14:[2,9],18:[2,9],22:[2,9],24:[2,9]},{14:[2,10],18:[2,10],22:[2,10],24:[2,10]},{14:[2,11],18:[2,11],22:[2,11],24:[2,11]},{14:[2,12],18:[2,12],22:[2,12],24:[2,12]},{14:[2,3],18:[2,3],22:[2,3],24:[2,3]},{14:[2,4],18:[2,4],22:[2,4],24:[2,4]},{14:[2,5],18:[2,5],22:[2,5],24:[2,5]},{14:[2,1],18:[2,1],21:[2,1],22:[2,1],24:[2,1]},{14:[2,2],18:[2,2],22:[2,2],24:[2,2]},{3:20,4:[1,12],18:[1,17],19:18,20:19},{3:5,4:[1,12],5:6,6:[1,13],7:3,8:[1,9],9:4,10:[1,10],11:[1,11],13:23,15:7,16:8,17:[1,14],23:[1,15],24:[1,21],25:22},{1:[2,6]},{14:[2,13],18:[2,13],22:[2,13],24:[2,13]},{18:[1,24],22:[1,25]},{18:[2,16],22:[2,16]},{21:[1,26]},{14:[2,18],18:[2,18],22:[2,18],24:[2,18]},{22:[1,28],24:[1,27]},{22:[2,20],24:[2,20]},{14:[2,14],18:[2,14],22:[2,14],24:[2,14]},{3:20,4:[1,12],20:29},{3:5,4:[1,12],5:6,6:[1,13],7:3,8:[1,9],9:4,10:[1,10],11:[1,11],13:30,15:7,16:8,17:[1,14],23:[1,15]},{14:[2,19],18:[2,19],22:[2,19],24:[2,19]},{3:5,4:[1,12],5:6,6:[1,13],7:3,8:[1,9],9:4,10:[1,10],11:[1,11],13:31,15:7,16:8,17:[1,14],23:[1,15]},{18:[2,17],22:[2,17]},{18:[2,15],22:[2,15]},{22:[2,21],24:[2,21]}],defaultActions:{16:[2,6]},parseError:function(t,e){throw new Error(t)},parse:function(t){var e=this,r=[0],n=[null],i=[],s=this.table,o="",a=0,l=0,c=0;this.lexer.setInput(t),this.lexer.yy=this.yy,this.yy.lexer=this.lexer,void 0===this.lexer.yylloc&&(this.lexer.yylloc={});var h=this.lexer.yylloc;function u(){var t;return"number"!=typeof(t=e.lexer.lex()||1)&&(t=e.symbols_[t]||t),t}i.push(h),"function"==typeof this.yy.parseError&&(this.parseError=this.yy.parseError);for(var f,p,d,y,g,m,x,b,_,v,E={};;){if(d=r[r.length-1],this.defaultActions[d]?y=this.defaultActions[d]:(null==f&&(f=u()),y=s[d]&&s[d][f]),void 0===y||!y.length||!y[0]){if(!c){for(m in _=[],s[d])this.terminals_[m]&&m>2&&_.push("'"+this.terminals_[m]+"'");var w="";w=this.lexer.showPosition?"Parse error on line "+(a+1)+":\n"+this.lexer.showPosition()+"\nExpecting "+_.join(", ")+", got '"+this.terminals_[f]+"'":"Parse error on line "+(a+1)+": Unexpected "+(1==f?"end of input":"'"+(this.terminals_[f]||f)+"'"),this.parseError(w,{text:this.lexer.match,token:this.terminals_[f]||f,line:this.lexer.yylineno,loc:h,expected:_})}if(3==c){if(1==f)throw new Error(w||"Parsing halted.");l=this.lexer.yyleng,o=this.lexer.yytext,a=this.lexer.yylineno,h=this.lexer.yylloc,f=u()}for(;!(2..toString()in s[d]);){if(0==d)throw new Error(w||"Parsing halted.");v=1,r.length=r.length-2*v,n.length=n.length-v,i.length=i.length-v,d=r[r.length-1]}p=f,f=2,y=s[d=r[r.length-1]]&&s[d][2],c=3}if(y[0]instanceof Array&&y.length>1)throw new Error("Parse Error: multiple actions possible at state: "+d+", token: "+f);switch(y[0]){case 1:r.push(f),n.push(this.lexer.yytext),i.push(this.lexer.yylloc),r.push(y[1]),f=null,p?(f=p,p=null):(l=this.lexer.yyleng,o=this.lexer.yytext,a=this.lexer.yylineno,h=this.lexer.yylloc,c>0&&c--);break;case 2:if(x=this.productions_[y[1]][1],E.$=n[n.length-x],E._$={first_line:i[i.length-(x||1)].first_line,last_line:i[i.length-1].last_line,first_column:i[i.length-(x||1)].first_column,last_column:i[i.length-1].last_column},void 0!==(g=this.performAction.call(E,o,l,a,this.yy,y[1],n,i)))return g;x&&(r=r.slice(0,-1*x*2),n=n.slice(0,-1*x),i=i.slice(0,-1*x)),r.push(this.productions_[y[1]][0]),n.push(E.$),i.push(E._$),b=s[r[r.length-2]][r[r.length-1]],r.push(b);break;case 3:return!0}}return!0}}).lexer={EOF:1,parseError:function(t,e){if(!this.yy.parseError)throw new Error(t);this.yy.parseError(t,e)},setInput:function(t){return this._input=t,this._more=this._less=this.done=!1,this.yylineno=this.yyleng=0,this.yytext=this.matched=this.match="",this.conditionStack=["INITIAL"],this.yylloc={first_line:1,first_column:0,last_line:1,last_column:0},this},input:function(){var t=this._input[0];return this.yytext+=t,this.yyleng++,this.match+=t,this.matched+=t,t.match(/\n/)&&this.yylineno++,this._input=this._input.slice(1),t},unput:function(t){return this._input=t+this._input,this},more:function(){return this._more=!0,this},less:function(t){this._input=this.match.slice(t)+this._input},pastInput:function(){var t=this.matched.substr(0,this.matched.length-this.match.length);return(t.length>20?"...":"")+t.substr(-20).replace(/\n/g,"")},upcomingInput:function(){var t=this.match;return t.length<20&&(t+=this._input.substr(0,20-t.length)),(t.substr(0,20)+(t.length>20?"...":"")).replace(/\n/g,"")},showPosition:function(){var t=this.pastInput(),e=new Array(t.length+1).join("-");return t+this.upcomingInput()+"\n"+e+"^"},next:function(){if(this.done)return this.EOF;var t,e,r,n,i;this._input||(this.done=!0),this._more||(this.yytext="",this.match="");for(var s=this._currentRules(),o=0;o<s.length&&(!(r=this._input.match(this.rules[s[o]]))||e&&!(r[0].length>e[0].length)||(e=r,n=o,this.options.flex));o++);return e?((i=e[0].match(/\n.*/g))&&(this.yylineno+=i.length),this.yylloc={first_line:this.yylloc.last_line,last_line:this.yylineno+1,first_column:this.yylloc.last_column,last_column:i?i[i.length-1].length-1:this.yylloc.last_column+e[0].length},this.yytext+=e[0],this.match+=e[0],this.yyleng=this.yytext.length,this._more=!1,this._input=this._input.slice(e[0].length),this.matched+=e[0],t=this.performAction.call(this,this.yy,this,s[n],this.conditionStack[this.conditionStack.length-1]),this.done&&this._input&&(this.done=!1),t||void 0):""===this._input?this.EOF:void this.parseError("Lexical error on line "+(this.yylineno+1)+". Unrecognized text.\n"+this.showPosition(),{text:"",token:null,line:this.yylineno})},lex:function(){var t=this.next();return void 0!==t?t:this.lex()},begin:function(t){this.conditionStack.push(t)},popState:function(){return this.conditionStack.pop()},_currentRules:function(){return this.conditions[this.conditionStack[this.conditionStack.length-1]].rules},topState:function(){return this.conditionStack[this.conditionStack.length-2]},pushState:function(t){this.begin(t)},options:{},performAction:function(t,e,r,n){switch(r){case 0:break;case 1:return 6;case 2:return e.yytext=e.yytext.substr(1,e.yyleng-2),4;case 3:return 17;case 4:return 18;case 5:return 23;case 6:return 24;case 7:return 22;case 8:return 21;case 9:return 10;case 10:return 11;case 11:return 8;case 12:return 14;case 13:return"INVALID"}},rules:[/^(?:\s+)/,/^(?:(-?([0-9]|[1-9][0-9]+))(\.[0-9]+)?([eE][-+]?[0-9]+)?\b)/,/^(?:"(?:\\[\\"bfnrt/]|\\u[a-fA-F0-9]{4}|[^\\\0-\x09\x0a-\x1f"])*")/,/^(?:\{)/,/^(?:\})/,/^(?:\[)/,/^(?:\])/,/^(?:,)/,/^(?::)/,/^(?:true\b)/,/^(?:false\b)/,/^(?:null\b)/,/^(?:$)/,/^(?:.)/],conditions:{INITIAL:{rules:[0,1,2,3,4,5,6,7,8,9,10,11,12,13],inclusive:!0}}},$=(O=S).parse,O.parse=function(t){var e=$.call(O,t),r=void 0===N?k("./doug-json-parse"):N;try{r(t)}catch(o){if(/Duplicate key|Bad string|Unexpected/.test(o.message)){var n=t.substring(0,o.at).split("\n"),i=n.length,s=n[i-1].length-1;throw this.parseError(o.message,{line:i,col:s,message:o.message.replace(/./,(function(t){return t.toLowerCase()}))}),SyntaxError(o.message+" on line "+i)}}return e},void 0!==j&&(j.parser=O,j.parse=function(){return O.parse.apply(O,arguments)},j.main=function(t){if(!t[1])throw new Error("Usage: "+t[0]+" FILE");if("undefined"!=typeof process)var e=k("fs").readFileSync(k("path").join(process.cwd(),t[1]),"utf8");else e=k("file").path(k("file").cwd()).join(t[1]).read({charset:"utf-8"});return j.parser.parse(e)},!1===k.main&&j.main("undefined"!=typeof process?process.argv.slice(1):k("system").args)),A=j,p.exports&&(p.exports=A);var L=I.exports;window.jsonlint=L;const F={name:"JsonEditor",data:()=>({data:[],defaultProps:{children:"children",label:"label"},jsonEditor:!1,value:JSON.stringify(JSON.parse('[{"怎么操作": "请选择左侧节点，以展示内容"}]'),null,4)}),methods:{doLoadAllFiles(){r().then((t=>{this.data=t.data})).catch((t=>{console.log(t)}))},doloadFileContent(t){n(t).then((t=>{this.jsonEditor.setValue(t.data)})).catch((t=>{console.log(t)}))},handleNodeClick(t,e,r){if(!t.path)return!1;const n={label:t.label,path:t.path};this.doloadFileContent(n)},refresh(){this.jsonEditor.refresh()}},beforeUnmount(){},mounted(){this.jsonEditor=i(e.fromTextArea(this.$refs.textareaGroup,{mode:"application/json",indentUnit:2,smartIndent:!0,styleActiveLine:!0,lineNumbers:!0,theme:"darcula",keyMap:"sublime",lineWrapping:!0,foldGutter:!0,gutters:["CodeMirror-linenumbers","CodeMirror-foldgutter","CodeMirror-lint-markers"],lint:!0,fullScreen:!1,matchBrackets:!0,autoCloseBrackets:!0,extraKeys:{F11:t=>{t.setOption("fullScreen",!t.getOption("fullScreen"))},Esc:t=>{t.getOption("fullScreen")&&t.setOption("fullScreen",!1)}}})),this.jsonEditor.refresh(),this.jsonEditor.setValue(this.value),this.jsonEditor.on("change",(t=>{this.$emit("changed",t.getValue()),this.$emit("input",t.getValue())})),this.doLoadAllFiles()}},C=u();s("data-v-46c87bcf");const J={class:"ftree-box"},P={class:"json-editor"},U={ref:"textareaGroup"};o();const B=C(((t,e,r,n,i,s)=>{const o=a("el-tree"),u=a("el-scrollbar");return l(),c("div",J,[h("div",null,[h(u,{class:"card-scrollbar"},{default:C((()=>[h(o,{data:i.data,props:i.defaultProps,accordion:"","default-expand-all":"",indent:0,onNodeClick:s.handleNodeClick,class:"tree"},null,8,["data","props","onNodeClick"])])),_:1})]),h("div",P,[h("textarea",U,null,512)])])}));F.render=B,F.__scopeId="data-v-46c87bcf";export default F;
