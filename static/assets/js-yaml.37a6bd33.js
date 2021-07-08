/*! js-yaml 4.1.0 https://github.com/nodeca/js-yaml @license MIT */
function e(e){return null==e}var t={isNothing:e,isObject:function(e){return"object"==typeof e&&null!==e},toArray:function(t){return Array.isArray(t)?t:e(t)?[]:[t]},repeat:function(e,t){var n,i="";for(n=0;n<t;n+=1)i+=e;return i},isNegativeZero:function(e){return 0===e&&Number.NEGATIVE_INFINITY===1/e},extend:function(e,t){var n,i,r,o;if(t)for(n=0,i=(o=Object.keys(t)).length;n<i;n+=1)e[r=o[n]]=t[r];return e}};function n(e,t){var n="",i=e.reason||"(unknown reason)";return e.mark?(e.mark.name&&(n+='in "'+e.mark.name+'" '),n+="("+(e.mark.line+1)+":"+(e.mark.column+1)+")",!t&&e.mark.snippet&&(n+="\n\n"+e.mark.snippet),i+" "+n):i}function i(e,t){Error.call(this),this.name="YAMLException",this.reason=e,this.mark=t,this.message=n(this,!1),Error.captureStackTrace?Error.captureStackTrace(this,this.constructor):this.stack=(new Error).stack||""}i.prototype=Object.create(Error.prototype),i.prototype.constructor=i,i.prototype.toString=function(e){return this.name+": "+n(this,e)};var r=i;function o(e,t,n,i,r){var o="",a="",l=Math.floor(r/2)-1;return i-t>l&&(t=i-l+(o=" ... ").length),n-i>l&&(n=i+l-(a=" ...").length),{str:o+e.slice(t,n).replace(/\t/g,"→")+a,pos:i-t+o.length}}function a(e,n){return t.repeat(" ",n-e.length)+e}var l=function(e,n){if(n=Object.create(n||null),!e.buffer)return null;n.maxLength||(n.maxLength=79),"number"!=typeof n.indent&&(n.indent=1),"number"!=typeof n.linesBefore&&(n.linesBefore=3),"number"!=typeof n.linesAfter&&(n.linesAfter=2);for(var i,r=/\r?\n|\r|\0/g,l=[0],c=[],s=-1;i=r.exec(e.buffer);)c.push(i.index),l.push(i.index+i[0].length),e.position<=i.index&&s<0&&(s=l.length-2);s<0&&(s=l.length-1);var u,p,f="",d=Math.min(e.line+n.linesAfter,c.length).toString().length,h=n.maxLength-(n.indent+d+3);for(u=1;u<=n.linesBefore&&!(s-u<0);u++)p=o(e.buffer,l[s-u],c[s-u],e.position-(l[s]-l[s-u]),h),f=t.repeat(" ",n.indent)+a((e.line-u+1).toString(),d)+" | "+p.str+"\n"+f;for(p=o(e.buffer,l[s],c[s],e.position,h),f+=t.repeat(" ",n.indent)+a((e.line+1).toString(),d)+" | "+p.str+"\n",f+=t.repeat("-",n.indent+d+3+p.pos)+"^\n",u=1;u<=n.linesAfter&&!(s+u>=c.length);u++)p=o(e.buffer,l[s+u],c[s+u],e.position-(l[s]-l[s+u]),h),f+=t.repeat(" ",n.indent)+a((e.line+u+1).toString(),d)+" | "+p.str+"\n";return f.replace(/\n$/,"")},c=["kind","multi","resolve","construct","instanceOf","predicate","represent","representName","defaultStyle","styleAliases"],s=["scalar","sequence","mapping"];var u=function(e,t){var n,i;if(t=t||{},Object.keys(t).forEach((function(t){if(-1===c.indexOf(t))throw new r('Unknown option "'+t+'" is met in definition of "'+e+'" YAML type.')})),this.options=t,this.tag=e,this.kind=t.kind||null,this.resolve=t.resolve||function(){return!0},this.construct=t.construct||function(e){return e},this.instanceOf=t.instanceOf||null,this.predicate=t.predicate||null,this.represent=t.represent||null,this.representName=t.representName||null,this.defaultStyle=t.defaultStyle||null,this.multi=t.multi||!1,this.styleAliases=(n=t.styleAliases||null,i={},null!==n&&Object.keys(n).forEach((function(e){n[e].forEach((function(t){i[String(t)]=e}))})),i),-1===s.indexOf(this.kind))throw new r('Unknown kind "'+this.kind+'" is specified for "'+e+'" YAML type.')};function p(e,t){var n=[];return e[t].forEach((function(e){var t=n.length;n.forEach((function(n,i){n.tag===e.tag&&n.kind===e.kind&&n.multi===e.multi&&(t=i)})),n[t]=e})),n}function f(e){return this.extend(e)}f.prototype.extend=function(e){var t=[],n=[];if(e instanceof u)n.push(e);else if(Array.isArray(e))n=n.concat(e);else{if(!e||!Array.isArray(e.implicit)&&!Array.isArray(e.explicit))throw new r("Schema.extend argument should be a Type, [ Type ], or a schema definition ({ implicit: [...], explicit: [...] })");e.implicit&&(t=t.concat(e.implicit)),e.explicit&&(n=n.concat(e.explicit))}t.forEach((function(e){if(!(e instanceof u))throw new r("Specified list of YAML types (or a single Type object) contains a non-Type object.");if(e.loadKind&&"scalar"!==e.loadKind)throw new r("There is a non-scalar type in the implicit list of a schema. Implicit resolving of such types is not supported.");if(e.multi)throw new r("There is a multi type in the implicit list of a schema. Multi tags can only be listed as explicit.")})),n.forEach((function(e){if(!(e instanceof u))throw new r("Specified list of YAML types (or a single Type object) contains a non-Type object.")}));var i=Object.create(f.prototype);return i.implicit=(this.implicit||[]).concat(t),i.explicit=(this.explicit||[]).concat(n),i.compiledImplicit=p(i,"implicit"),i.compiledExplicit=p(i,"explicit"),i.compiledTypeMap=function(){var e,t,n={scalar:{},sequence:{},mapping:{},fallback:{},multi:{scalar:[],sequence:[],mapping:[],fallback:[]}};function i(e){e.multi?(n.multi[e.kind].push(e),n.multi.fallback.push(e)):n[e.kind][e.tag]=n.fallback[e.tag]=e}for(e=0,t=arguments.length;e<t;e+=1)arguments[e].forEach(i);return n}(i.compiledImplicit,i.compiledExplicit),i};var d=f,h=new u("tag:yaml.org,2002:str",{kind:"scalar",construct:function(e){return null!==e?e:""}}),g=new u("tag:yaml.org,2002:seq",{kind:"sequence",construct:function(e){return null!==e?e:[]}}),m=new u("tag:yaml.org,2002:map",{kind:"mapping",construct:function(e){return null!==e?e:{}}}),y=new d({explicit:[h,g,m]});var b=new u("tag:yaml.org,2002:null",{kind:"scalar",resolve:function(e){if(null===e)return!0;var t=e.length;return 1===t&&"~"===e||4===t&&("null"===e||"Null"===e||"NULL"===e)},construct:function(){return null},predicate:function(e){return null===e},represent:{canonical:function(){return"~"},lowercase:function(){return"null"},uppercase:function(){return"NULL"},camelcase:function(){return"Null"},empty:function(){return""}},defaultStyle:"lowercase"});var A=new u("tag:yaml.org,2002:bool",{kind:"scalar",resolve:function(e){if(null===e)return!1;var t=e.length;return 4===t&&("true"===e||"True"===e||"TRUE"===e)||5===t&&("false"===e||"False"===e||"FALSE"===e)},construct:function(e){return"true"===e||"True"===e||"TRUE"===e},predicate:function(e){return"[object Boolean]"===Object.prototype.toString.call(e)},represent:{lowercase:function(e){return e?"true":"false"},uppercase:function(e){return e?"TRUE":"FALSE"},camelcase:function(e){return e?"True":"False"}},defaultStyle:"lowercase"});function v(e){return 48<=e&&e<=55}function w(e){return 48<=e&&e<=57}var k=new u("tag:yaml.org,2002:int",{kind:"scalar",resolve:function(e){if(null===e)return!1;var t,n,i=e.length,r=0,o=!1;if(!i)return!1;if("-"!==(t=e[r])&&"+"!==t||(t=e[++r]),"0"===t){if(r+1===i)return!0;if("b"===(t=e[++r])){for(r++;r<i;r++)if("_"!==(t=e[r])){if("0"!==t&&"1"!==t)return!1;o=!0}return o&&"_"!==t}if("x"===t){for(r++;r<i;r++)if("_"!==(t=e[r])){if(!(48<=(n=e.charCodeAt(r))&&n<=57||65<=n&&n<=70||97<=n&&n<=102))return!1;o=!0}return o&&"_"!==t}if("o"===t){for(r++;r<i;r++)if("_"!==(t=e[r])){if(!v(e.charCodeAt(r)))return!1;o=!0}return o&&"_"!==t}}if("_"===t)return!1;for(;r<i;r++)if("_"!==(t=e[r])){if(!w(e.charCodeAt(r)))return!1;o=!0}return!(!o||"_"===t)},construct:function(e){var t,n=e,i=1;if(-1!==n.indexOf("_")&&(n=n.replace(/_/g,"")),"-"!==(t=n[0])&&"+"!==t||("-"===t&&(i=-1),t=(n=n.slice(1))[0]),"0"===n)return 0;if("0"===t){if("b"===n[1])return i*parseInt(n.slice(2),2);if("x"===n[1])return i*parseInt(n.slice(2),16);if("o"===n[1])return i*parseInt(n.slice(2),8)}return i*parseInt(n,10)},predicate:function(e){return"[object Number]"===Object.prototype.toString.call(e)&&e%1==0&&!t.isNegativeZero(e)},represent:{binary:function(e){return e>=0?"0b"+e.toString(2):"-0b"+e.toString(2).slice(1)},octal:function(e){return e>=0?"0o"+e.toString(8):"-0o"+e.toString(8).slice(1)},decimal:function(e){return e.toString(10)},hexadecimal:function(e){return e>=0?"0x"+e.toString(16).toUpperCase():"-0x"+e.toString(16).toUpperCase().slice(1)}},defaultStyle:"decimal",styleAliases:{binary:[2,"bin"],octal:[8,"oct"],decimal:[10,"dec"],hexadecimal:[16,"hex"]}}),C=new RegExp("^(?:[-+]?(?:[0-9][0-9_]*)(?:\\.[0-9_]*)?(?:[eE][-+]?[0-9]+)?|\\.[0-9_]+(?:[eE][-+]?[0-9]+)?|[-+]?\\.(?:inf|Inf|INF)|\\.(?:nan|NaN|NAN))$");var x=/^[-+]?[0-9]+e/;var I=new u("tag:yaml.org,2002:float",{kind:"scalar",resolve:function(e){return null!==e&&!(!C.test(e)||"_"===e[e.length-1])},construct:function(e){var t,n;return n="-"===(t=e.replace(/_/g,"").toLowerCase())[0]?-1:1,"+-".indexOf(t[0])>=0&&(t=t.slice(1)),".inf"===t?1===n?Number.POSITIVE_INFINITY:Number.NEGATIVE_INFINITY:".nan"===t?NaN:n*parseFloat(t,10)},predicate:function(e){return"[object Number]"===Object.prototype.toString.call(e)&&(e%1!=0||t.isNegativeZero(e))},represent:function(e,n){var i;if(isNaN(e))switch(n){case"lowercase":return".nan";case"uppercase":return".NAN";case"camelcase":return".NaN"}else if(Number.POSITIVE_INFINITY===e)switch(n){case"lowercase":return".inf";case"uppercase":return".INF";case"camelcase":return".Inf"}else if(Number.NEGATIVE_INFINITY===e)switch(n){case"lowercase":return"-.inf";case"uppercase":return"-.INF";case"camelcase":return"-.Inf"}else if(t.isNegativeZero(e))return"-0.0";return i=e.toString(10),x.test(i)?i.replace("e",".e"):i},defaultStyle:"lowercase"}),S=y.extend({implicit:[b,A,k,I]}),O=S,j=new RegExp("^([0-9][0-9][0-9][0-9])-([0-9][0-9])-([0-9][0-9])$"),T=new RegExp("^([0-9][0-9][0-9][0-9])-([0-9][0-9]?)-([0-9][0-9]?)(?:[Tt]|[ \\t]+)([0-9][0-9]?):([0-9][0-9]):([0-9][0-9])(?:\\.([0-9]*))?(?:[ \\t]*(Z|([-+])([0-9][0-9]?)(?::([0-9][0-9]))?))?$");var N=new u("tag:yaml.org,2002:timestamp",{kind:"scalar",resolve:function(e){return null!==e&&(null!==j.exec(e)||null!==T.exec(e))},construct:function(e){var t,n,i,r,o,a,l,c,s=0,u=null;if(null===(t=j.exec(e))&&(t=T.exec(e)),null===t)throw new Error("Date resolve error");if(n=+t[1],i=+t[2]-1,r=+t[3],!t[4])return new Date(Date.UTC(n,i,r));if(o=+t[4],a=+t[5],l=+t[6],t[7]){for(s=t[7].slice(0,3);s.length<3;)s+="0";s=+s}return t[9]&&(u=6e4*(60*+t[10]+ +(t[11]||0)),"-"===t[9]&&(u=-u)),c=new Date(Date.UTC(n,i,r,o,a,l,s)),u&&c.setTime(c.getTime()-u),c},instanceOf:Date,represent:function(e){return e.toISOString()}});var F=new u("tag:yaml.org,2002:merge",{kind:"scalar",resolve:function(e){return"<<"===e||null===e}}),E="ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=\n\r";var M=new u("tag:yaml.org,2002:binary",{kind:"scalar",resolve:function(e){if(null===e)return!1;var t,n,i=0,r=e.length,o=E;for(n=0;n<r;n++)if(!((t=o.indexOf(e.charAt(n)))>64)){if(t<0)return!1;i+=6}return i%8==0},construct:function(e){var t,n,i=e.replace(/[\r\n=]/g,""),r=i.length,o=E,a=0,l=[];for(t=0;t<r;t++)t%4==0&&t&&(l.push(a>>16&255),l.push(a>>8&255),l.push(255&a)),a=a<<6|o.indexOf(i.charAt(t));return 0===(n=r%4*6)?(l.push(a>>16&255),l.push(a>>8&255),l.push(255&a)):18===n?(l.push(a>>10&255),l.push(a>>2&255)):12===n&&l.push(a>>4&255),new Uint8Array(l)},predicate:function(e){return"[object Uint8Array]"===Object.prototype.toString.call(e)},represent:function(e){var t,n,i="",r=0,o=e.length,a=E;for(t=0;t<o;t++)t%3==0&&t&&(i+=a[r>>18&63],i+=a[r>>12&63],i+=a[r>>6&63],i+=a[63&r]),r=(r<<8)+e[t];return 0===(n=o%3)?(i+=a[r>>18&63],i+=a[r>>12&63],i+=a[r>>6&63],i+=a[63&r]):2===n?(i+=a[r>>10&63],i+=a[r>>4&63],i+=a[r<<2&63],i+=a[64]):1===n&&(i+=a[r>>2&63],i+=a[r<<4&63],i+=a[64],i+=a[64]),i}}),L=Object.prototype.hasOwnProperty,_=Object.prototype.toString;var D=new u("tag:yaml.org,2002:omap",{kind:"sequence",resolve:function(e){if(null===e)return!0;var t,n,i,r,o,a=[],l=e;for(t=0,n=l.length;t<n;t+=1){if(i=l[t],o=!1,"[object Object]"!==_.call(i))return!1;for(r in i)if(L.call(i,r)){if(o)return!1;o=!0}if(!o)return!1;if(-1!==a.indexOf(r))return!1;a.push(r)}return!0},construct:function(e){return null!==e?e:[]}}),U=Object.prototype.toString;var q=new u("tag:yaml.org,2002:pairs",{kind:"sequence",resolve:function(e){if(null===e)return!0;var t,n,i,r,o,a=e;for(o=new Array(a.length),t=0,n=a.length;t<n;t+=1){if(i=a[t],"[object Object]"!==U.call(i))return!1;if(1!==(r=Object.keys(i)).length)return!1;o[t]=[r[0],i[r[0]]]}return!0},construct:function(e){if(null===e)return[];var t,n,i,r,o,a=e;for(o=new Array(a.length),t=0,n=a.length;t<n;t+=1)i=a[t],r=Object.keys(i),o[t]=[r[0],i[r[0]]];return o}}),Y=Object.prototype.hasOwnProperty;var R=new u("tag:yaml.org,2002:set",{kind:"mapping",resolve:function(e){if(null===e)return!0;var t,n=e;for(t in n)if(Y.call(n,t)&&null!==n[t])return!1;return!0},construct:function(e){return null!==e?e:{}}}),B=O.extend({implicit:[N,F],explicit:[M,D,q,R]}),K=Object.prototype.hasOwnProperty,W=/[\x00-\x08\x0B\x0C\x0E-\x1F\x7F-\x84\x86-\x9F\uFFFE\uFFFF]|[\uD800-\uDBFF](?![\uDC00-\uDFFF])|(?:[^\uD800-\uDBFF]|^)[\uDC00-\uDFFF]/,H=/[\x85\u2028\u2029]/,P=/[,\[\]\{\}]/,$=/^(?:!|!!|![a-z\-]+!)$/i,G=/^(?:!|[^,\[\]\{\}])(?:%[0-9a-f]{2}|[0-9a-z\-#;\/\?:@&=\+\$,_\.!~\*'\(\)\[\]])*$/i;function V(e){return Object.prototype.toString.call(e)}function Z(e){return 10===e||13===e}function z(e){return 9===e||32===e}function J(e){return 9===e||32===e||10===e||13===e}function Q(e){return 44===e||91===e||93===e||123===e||125===e}function X(e){var t;return 48<=e&&e<=57?e-48:97<=(t=32|e)&&t<=102?t-97+10:-1}function ee(e){return 48===e?"\0":97===e?"":98===e?"\b":116===e||9===e?"\t":110===e?"\n":118===e?"\v":102===e?"\f":114===e?"\r":101===e?"":32===e?" ":34===e?'"':47===e?"/":92===e?"\\":78===e?"":95===e?" ":76===e?"\u2028":80===e?"\u2029":""}function te(e){return e<=65535?String.fromCharCode(e):String.fromCharCode(55296+(e-65536>>10),56320+(e-65536&1023))}for(var ne=new Array(256),ie=new Array(256),re=0;re<256;re++)ne[re]=ee(re)?1:0,ie[re]=ee(re);function oe(e,t){this.input=e,this.filename=t.filename||null,this.schema=t.schema||B,this.onWarning=t.onWarning||null,this.legacy=t.legacy||!1,this.json=t.json||!1,this.listener=t.listener||null,this.implicitTypes=this.schema.compiledImplicit,this.typeMap=this.schema.compiledTypeMap,this.length=e.length,this.position=0,this.line=0,this.lineStart=0,this.lineIndent=0,this.firstTabInLine=-1,this.documents=[]}function ae(e,t){var n={name:e.filename,buffer:e.input.slice(0,-1),position:e.position,line:e.line,column:e.position-e.lineStart};return n.snippet=l(n),new r(t,n)}function le(e,t){throw ae(e,t)}function ce(e,t){e.onWarning&&e.onWarning.call(null,ae(e,t))}var se={YAML:function(e,t,n){var i,r,o;null!==e.version&&le(e,"duplication of %YAML directive"),1!==n.length&&le(e,"YAML directive accepts exactly one argument"),null===(i=/^([0-9]+)\.([0-9]+)$/.exec(n[0]))&&le(e,"ill-formed argument of the YAML directive"),r=parseInt(i[1],10),o=parseInt(i[2],10),1!==r&&le(e,"unacceptable YAML version of the document"),e.version=n[0],e.checkLineBreaks=o<2,1!==o&&2!==o&&ce(e,"unsupported YAML version of the document")},TAG:function(e,t,n){var i,r;2!==n.length&&le(e,"TAG directive accepts exactly two arguments"),i=n[0],r=n[1],$.test(i)||le(e,"ill-formed tag handle (first argument) of the TAG directive"),K.call(e.tagMap,i)&&le(e,'there is a previously declared suffix for "'+i+'" tag handle'),G.test(r)||le(e,"ill-formed tag prefix (second argument) of the TAG directive");try{r=decodeURIComponent(r)}catch(o){le(e,"tag prefix is malformed: "+r)}e.tagMap[i]=r}};function ue(e,t,n,i){var r,o,a,l;if(t<n){if(l=e.input.slice(t,n),i)for(r=0,o=l.length;r<o;r+=1)9===(a=l.charCodeAt(r))||32<=a&&a<=1114111||le(e,"expected valid JSON character");else W.test(l)&&le(e,"the stream contains non-printable characters");e.result+=l}}function pe(e,n,i,r){var o,a,l,c;for(t.isObject(i)||le(e,"cannot merge mappings; the provided source object is unacceptable"),l=0,c=(o=Object.keys(i)).length;l<c;l+=1)a=o[l],K.call(n,a)||(n[a]=i[a],r[a]=!0)}function fe(e,t,n,i,r,o,a,l,c){var s,u;if(Array.isArray(r))for(s=0,u=(r=Array.prototype.slice.call(r)).length;s<u;s+=1)Array.isArray(r[s])&&le(e,"nested arrays are not supported inside keys"),"object"==typeof r&&"[object Object]"===V(r[s])&&(r[s]="[object Object]");if("object"==typeof r&&"[object Object]"===V(r)&&(r="[object Object]"),r=String(r),null===t&&(t={}),"tag:yaml.org,2002:merge"===i)if(Array.isArray(o))for(s=0,u=o.length;s<u;s+=1)pe(e,t,o[s],n);else pe(e,t,o,n);else e.json||K.call(n,r)||!K.call(t,r)||(e.line=a||e.line,e.lineStart=l||e.lineStart,e.position=c||e.position,le(e,"duplicated mapping key")),"__proto__"===r?Object.defineProperty(t,r,{configurable:!0,enumerable:!0,writable:!0,value:o}):t[r]=o,delete n[r];return t}function de(e){var t;10===(t=e.input.charCodeAt(e.position))?e.position++:13===t?(e.position++,10===e.input.charCodeAt(e.position)&&e.position++):le(e,"a line break is expected"),e.line+=1,e.lineStart=e.position,e.firstTabInLine=-1}function he(e,t,n){for(var i=0,r=e.input.charCodeAt(e.position);0!==r;){for(;z(r);)9===r&&-1===e.firstTabInLine&&(e.firstTabInLine=e.position),r=e.input.charCodeAt(++e.position);if(t&&35===r)do{r=e.input.charCodeAt(++e.position)}while(10!==r&&13!==r&&0!==r);if(!Z(r))break;for(de(e),r=e.input.charCodeAt(e.position),i++,e.lineIndent=0;32===r;)e.lineIndent++,r=e.input.charCodeAt(++e.position)}return-1!==n&&0!==i&&e.lineIndent<n&&ce(e,"deficient indentation"),i}function ge(e){var t,n=e.position;return!(45!==(t=e.input.charCodeAt(n))&&46!==t||t!==e.input.charCodeAt(n+1)||t!==e.input.charCodeAt(n+2)||(n+=3,0!==(t=e.input.charCodeAt(n))&&!J(t)))}function me(e,n){1===n?e.result+=" ":n>1&&(e.result+=t.repeat("\n",n-1))}function ye(e,t){var n,i,r=e.tag,o=e.anchor,a=[],l=!1;if(-1!==e.firstTabInLine)return!1;for(null!==e.anchor&&(e.anchorMap[e.anchor]=a),i=e.input.charCodeAt(e.position);0!==i&&(-1!==e.firstTabInLine&&(e.position=e.firstTabInLine,le(e,"tab characters must not be used in indentation")),45===i)&&J(e.input.charCodeAt(e.position+1));)if(l=!0,e.position++,he(e,!0,-1)&&e.lineIndent<=t)a.push(null),i=e.input.charCodeAt(e.position);else if(n=e.line,ve(e,t,3,!1,!0),a.push(e.result),he(e,!0,-1),i=e.input.charCodeAt(e.position),(e.line===n||e.lineIndent>t)&&0!==i)le(e,"bad indentation of a sequence entry");else if(e.lineIndent<t)break;return!!l&&(e.tag=r,e.anchor=o,e.kind="sequence",e.result=a,!0)}function be(e){var t,n,i,r,o=!1,a=!1;if(33!==(r=e.input.charCodeAt(e.position)))return!1;if(null!==e.tag&&le(e,"duplication of a tag property"),60===(r=e.input.charCodeAt(++e.position))?(o=!0,r=e.input.charCodeAt(++e.position)):33===r?(a=!0,n="!!",r=e.input.charCodeAt(++e.position)):n="!",t=e.position,o){do{r=e.input.charCodeAt(++e.position)}while(0!==r&&62!==r);e.position<e.length?(i=e.input.slice(t,e.position),r=e.input.charCodeAt(++e.position)):le(e,"unexpected end of the stream within a verbatim tag")}else{for(;0!==r&&!J(r);)33===r&&(a?le(e,"tag suffix cannot contain exclamation marks"):(n=e.input.slice(t-1,e.position+1),$.test(n)||le(e,"named tag handle cannot contain such characters"),a=!0,t=e.position+1)),r=e.input.charCodeAt(++e.position);i=e.input.slice(t,e.position),P.test(i)&&le(e,"tag suffix cannot contain flow indicator characters")}i&&!G.test(i)&&le(e,"tag name cannot contain such characters: "+i);try{i=decodeURIComponent(i)}catch(l){le(e,"tag name is malformed: "+i)}return o?e.tag=i:K.call(e.tagMap,n)?e.tag=e.tagMap[n]+i:"!"===n?e.tag="!"+i:"!!"===n?e.tag="tag:yaml.org,2002:"+i:le(e,'undeclared tag handle "'+n+'"'),!0}function Ae(e){var t,n;if(38!==(n=e.input.charCodeAt(e.position)))return!1;for(null!==e.anchor&&le(e,"duplication of an anchor property"),n=e.input.charCodeAt(++e.position),t=e.position;0!==n&&!J(n)&&!Q(n);)n=e.input.charCodeAt(++e.position);return e.position===t&&le(e,"name of an anchor node must contain at least one character"),e.anchor=e.input.slice(t,e.position),!0}function ve(e,n,i,r,o){var a,l,c,s,u,p,f,d,h,g=1,m=!1,y=!1;if(null!==e.listener&&e.listener("open",e),e.tag=null,e.anchor=null,e.kind=null,e.result=null,a=l=c=4===i||3===i,r&&he(e,!0,-1)&&(m=!0,e.lineIndent>n?g=1:e.lineIndent===n?g=0:e.lineIndent<n&&(g=-1)),1===g)for(;be(e)||Ae(e);)he(e,!0,-1)?(m=!0,c=a,e.lineIndent>n?g=1:e.lineIndent===n?g=0:e.lineIndent<n&&(g=-1)):c=!1;if(c&&(c=m||o),1!==g&&4!==i||(d=1===i||2===i?n:n+1,h=e.position-e.lineStart,1===g?c&&(ye(e,h)||function(e,t,n){var i,r,o,a,l,c,s,u=e.tag,p=e.anchor,f={},d=Object.create(null),h=null,g=null,m=null,y=!1,b=!1;if(-1!==e.firstTabInLine)return!1;for(null!==e.anchor&&(e.anchorMap[e.anchor]=f),s=e.input.charCodeAt(e.position);0!==s;){if(y||-1===e.firstTabInLine||(e.position=e.firstTabInLine,le(e,"tab characters must not be used in indentation")),i=e.input.charCodeAt(e.position+1),o=e.line,63!==s&&58!==s||!J(i)){if(a=e.line,l=e.lineStart,c=e.position,!ve(e,n,2,!1,!0))break;if(e.line===o){for(s=e.input.charCodeAt(e.position);z(s);)s=e.input.charCodeAt(++e.position);if(58===s)J(s=e.input.charCodeAt(++e.position))||le(e,"a whitespace character is expected after the key-value separator within a block mapping"),y&&(fe(e,f,d,h,g,null,a,l,c),h=g=m=null),b=!0,y=!1,r=!1,h=e.tag,g=e.result;else{if(!b)return e.tag=u,e.anchor=p,!0;le(e,"can not read an implicit mapping pair; a colon is missed")}}else{if(!b)return e.tag=u,e.anchor=p,!0;le(e,"can not read a block mapping entry; a multiline key may not be an implicit key")}}else 63===s?(y&&(fe(e,f,d,h,g,null,a,l,c),h=g=m=null),b=!0,y=!0,r=!0):y?(y=!1,r=!0):le(e,"incomplete explicit mapping pair; a key node is missed; or followed by a non-tabulated empty line"),e.position+=1,s=i;if((e.line===o||e.lineIndent>t)&&(y&&(a=e.line,l=e.lineStart,c=e.position),ve(e,t,4,!0,r)&&(y?g=e.result:m=e.result),y||(fe(e,f,d,h,g,m,a,l,c),h=g=m=null),he(e,!0,-1),s=e.input.charCodeAt(e.position)),(e.line===o||e.lineIndent>t)&&0!==s)le(e,"bad indentation of a mapping entry");else if(e.lineIndent<t)break}return y&&fe(e,f,d,h,g,null,a,l,c),b&&(e.tag=u,e.anchor=p,e.kind="mapping",e.result=f),b}(e,h,d))||function(e,t){var n,i,r,o,a,l,c,s,u,p,f,d,h=!0,g=e.tag,m=e.anchor,y=Object.create(null);if(91===(d=e.input.charCodeAt(e.position)))a=93,s=!1,o=[];else{if(123!==d)return!1;a=125,s=!0,o={}}for(null!==e.anchor&&(e.anchorMap[e.anchor]=o),d=e.input.charCodeAt(++e.position);0!==d;){if(he(e,!0,t),(d=e.input.charCodeAt(e.position))===a)return e.position++,e.tag=g,e.anchor=m,e.kind=s?"mapping":"sequence",e.result=o,!0;h?44===d&&le(e,"expected the node content, but found ','"):le(e,"missed comma between flow collection entries"),f=null,l=c=!1,63===d&&J(e.input.charCodeAt(e.position+1))&&(l=c=!0,e.position++,he(e,!0,t)),n=e.line,i=e.lineStart,r=e.position,ve(e,t,1,!1,!0),p=e.tag,u=e.result,he(e,!0,t),d=e.input.charCodeAt(e.position),!c&&e.line!==n||58!==d||(l=!0,d=e.input.charCodeAt(++e.position),he(e,!0,t),ve(e,t,1,!1,!0),f=e.result),s?fe(e,o,y,p,u,f,n,i,r):l?o.push(fe(e,null,y,p,u,f,n,i,r)):o.push(u),he(e,!0,t),44===(d=e.input.charCodeAt(e.position))?(h=!0,d=e.input.charCodeAt(++e.position)):h=!1}le(e,"unexpected end of the stream within a flow collection")}(e,d)?y=!0:(l&&function(e,n){var i,r,o,a,l,c=1,s=!1,u=!1,p=n,f=0,d=!1;if(124===(a=e.input.charCodeAt(e.position)))r=!1;else{if(62!==a)return!1;r=!0}for(e.kind="scalar",e.result="";0!==a;)if(43===(a=e.input.charCodeAt(++e.position))||45===a)1===c?c=43===a?3:2:le(e,"repeat of a chomping mode identifier");else{if(!((o=48<=(l=a)&&l<=57?l-48:-1)>=0))break;0===o?le(e,"bad explicit indentation width of a block scalar; it cannot be less than one"):u?le(e,"repeat of an indentation width identifier"):(p=n+o-1,u=!0)}if(z(a)){do{a=e.input.charCodeAt(++e.position)}while(z(a));if(35===a)do{a=e.input.charCodeAt(++e.position)}while(!Z(a)&&0!==a)}for(;0!==a;){for(de(e),e.lineIndent=0,a=e.input.charCodeAt(e.position);(!u||e.lineIndent<p)&&32===a;)e.lineIndent++,a=e.input.charCodeAt(++e.position);if(!u&&e.lineIndent>p&&(p=e.lineIndent),Z(a))f++;else{if(e.lineIndent<p){3===c?e.result+=t.repeat("\n",s?1+f:f):1===c&&s&&(e.result+="\n");break}for(r?z(a)?(d=!0,e.result+=t.repeat("\n",s?1+f:f)):d?(d=!1,e.result+=t.repeat("\n",f+1)):0===f?s&&(e.result+=" "):e.result+=t.repeat("\n",f):e.result+=t.repeat("\n",s?1+f:f),s=!0,u=!0,f=0,i=e.position;!Z(a)&&0!==a;)a=e.input.charCodeAt(++e.position);ue(e,i,e.position,!1)}}return!0}(e,d)||function(e,t){var n,i,r;if(39!==(n=e.input.charCodeAt(e.position)))return!1;for(e.kind="scalar",e.result="",e.position++,i=r=e.position;0!==(n=e.input.charCodeAt(e.position));)if(39===n){if(ue(e,i,e.position,!0),39!==(n=e.input.charCodeAt(++e.position)))return!0;i=e.position,e.position++,r=e.position}else Z(n)?(ue(e,i,r,!0),me(e,he(e,!1,t)),i=r=e.position):e.position===e.lineStart&&ge(e)?le(e,"unexpected end of the document within a single quoted scalar"):(e.position++,r=e.position);le(e,"unexpected end of the stream within a single quoted scalar")}(e,d)||function(e,t){var n,i,r,o,a,l,c;if(34!==(l=e.input.charCodeAt(e.position)))return!1;for(e.kind="scalar",e.result="",e.position++,n=i=e.position;0!==(l=e.input.charCodeAt(e.position));){if(34===l)return ue(e,n,e.position,!0),e.position++,!0;if(92===l){if(ue(e,n,e.position,!0),Z(l=e.input.charCodeAt(++e.position)))he(e,!1,t);else if(l<256&&ne[l])e.result+=ie[l],e.position++;else if((a=120===(c=l)?2:117===c?4:85===c?8:0)>0){for(r=a,o=0;r>0;r--)(a=X(l=e.input.charCodeAt(++e.position)))>=0?o=(o<<4)+a:le(e,"expected hexadecimal character");e.result+=te(o),e.position++}else le(e,"unknown escape sequence");n=i=e.position}else Z(l)?(ue(e,n,i,!0),me(e,he(e,!1,t)),n=i=e.position):e.position===e.lineStart&&ge(e)?le(e,"unexpected end of the document within a double quoted scalar"):(e.position++,i=e.position)}le(e,"unexpected end of the stream within a double quoted scalar")}(e,d)?y=!0:!function(e){var t,n,i;if(42!==(i=e.input.charCodeAt(e.position)))return!1;for(i=e.input.charCodeAt(++e.position),t=e.position;0!==i&&!J(i)&&!Q(i);)i=e.input.charCodeAt(++e.position);return e.position===t&&le(e,"name of an alias node must contain at least one character"),n=e.input.slice(t,e.position),K.call(e.anchorMap,n)||le(e,'unidentified alias "'+n+'"'),e.result=e.anchorMap[n],he(e,!0,-1),!0}(e)?function(e,t,n){var i,r,o,a,l,c,s,u,p=e.kind,f=e.result;if(J(u=e.input.charCodeAt(e.position))||Q(u)||35===u||38===u||42===u||33===u||124===u||62===u||39===u||34===u||37===u||64===u||96===u)return!1;if((63===u||45===u)&&(J(i=e.input.charCodeAt(e.position+1))||n&&Q(i)))return!1;for(e.kind="scalar",e.result="",r=o=e.position,a=!1;0!==u;){if(58===u){if(J(i=e.input.charCodeAt(e.position+1))||n&&Q(i))break}else if(35===u){if(J(e.input.charCodeAt(e.position-1)))break}else{if(e.position===e.lineStart&&ge(e)||n&&Q(u))break;if(Z(u)){if(l=e.line,c=e.lineStart,s=e.lineIndent,he(e,!1,-1),e.lineIndent>=t){a=!0,u=e.input.charCodeAt(e.position);continue}e.position=o,e.line=l,e.lineStart=c,e.lineIndent=s;break}}a&&(ue(e,r,o,!1),me(e,e.line-l),r=o=e.position,a=!1),z(u)||(o=e.position+1),u=e.input.charCodeAt(++e.position)}return ue(e,r,o,!1),!!e.result||(e.kind=p,e.result=f,!1)}(e,d,1===i)&&(y=!0,null===e.tag&&(e.tag="?")):(y=!0,null===e.tag&&null===e.anchor||le(e,"alias node should not have any properties")),null!==e.anchor&&(e.anchorMap[e.anchor]=e.result)):0===g&&(y=c&&ye(e,h))),null===e.tag)null!==e.anchor&&(e.anchorMap[e.anchor]=e.result);else if("?"===e.tag){for(null!==e.result&&"scalar"!==e.kind&&le(e,'unacceptable node kind for !<?> tag; it should be "scalar", not "'+e.kind+'"'),s=0,u=e.implicitTypes.length;s<u;s+=1)if((f=e.implicitTypes[s]).resolve(e.result)){e.result=f.construct(e.result),e.tag=f.tag,null!==e.anchor&&(e.anchorMap[e.anchor]=e.result);break}}else if("!"!==e.tag){if(K.call(e.typeMap[e.kind||"fallback"],e.tag))f=e.typeMap[e.kind||"fallback"][e.tag];else for(f=null,s=0,u=(p=e.typeMap.multi[e.kind||"fallback"]).length;s<u;s+=1)if(e.tag.slice(0,p[s].tag.length)===p[s].tag){f=p[s];break}f||le(e,"unknown tag !<"+e.tag+">"),null!==e.result&&f.kind!==e.kind&&le(e,"unacceptable node kind for !<"+e.tag+'> tag; it should be "'+f.kind+'", not "'+e.kind+'"'),f.resolve(e.result,e.tag)?(e.result=f.construct(e.result,e.tag),null!==e.anchor&&(e.anchorMap[e.anchor]=e.result)):le(e,"cannot resolve a node with !<"+e.tag+"> explicit tag")}return null!==e.listener&&e.listener("close",e),null!==e.tag||null!==e.anchor||y}function we(e){var t,n,i,r,o=e.position,a=!1;for(e.version=null,e.checkLineBreaks=e.legacy,e.tagMap=Object.create(null),e.anchorMap=Object.create(null);0!==(r=e.input.charCodeAt(e.position))&&(he(e,!0,-1),r=e.input.charCodeAt(e.position),!(e.lineIndent>0||37!==r));){for(a=!0,r=e.input.charCodeAt(++e.position),t=e.position;0!==r&&!J(r);)r=e.input.charCodeAt(++e.position);for(i=[],(n=e.input.slice(t,e.position)).length<1&&le(e,"directive name must not be less than one character in length");0!==r;){for(;z(r);)r=e.input.charCodeAt(++e.position);if(35===r){do{r=e.input.charCodeAt(++e.position)}while(0!==r&&!Z(r));break}if(Z(r))break;for(t=e.position;0!==r&&!J(r);)r=e.input.charCodeAt(++e.position);i.push(e.input.slice(t,e.position))}0!==r&&de(e),K.call(se,n)?se[n](e,n,i):ce(e,'unknown document directive "'+n+'"')}he(e,!0,-1),0===e.lineIndent&&45===e.input.charCodeAt(e.position)&&45===e.input.charCodeAt(e.position+1)&&45===e.input.charCodeAt(e.position+2)?(e.position+=3,he(e,!0,-1)):a&&le(e,"directives end mark is expected"),ve(e,e.lineIndent-1,4,!1,!0),he(e,!0,-1),e.checkLineBreaks&&H.test(e.input.slice(o,e.position))&&ce(e,"non-ASCII line breaks are interpreted as content"),e.documents.push(e.result),e.position===e.lineStart&&ge(e)?46===e.input.charCodeAt(e.position)&&(e.position+=3,he(e,!0,-1)):e.position<e.length-1&&le(e,"end of the stream or a document separator is expected")}function ke(e,t){t=t||{},0!==(e=String(e)).length&&(10!==e.charCodeAt(e.length-1)&&13!==e.charCodeAt(e.length-1)&&(e+="\n"),65279===e.charCodeAt(0)&&(e=e.slice(1)));var n=new oe(e,t),i=e.indexOf("\0");for(-1!==i&&(n.position=i,le(n,"null byte is not allowed in input")),n.input+="\0";32===n.input.charCodeAt(n.position);)n.lineIndent+=1,n.position+=1;for(;n.position<n.length-1;)we(n);return n.documents}var Ce={loadAll:function(e,t,n){null!==t&&"object"==typeof t&&void 0===n&&(n=t,t=null);var i=ke(e,n);if("function"!=typeof t)return i;for(var r=0,o=i.length;r<o;r+=1)t(i[r])},load:function(e,t){var n=ke(e,t);if(0!==n.length){if(1===n.length)return n[0];throw new r("expected a single document in the stream, but found more")}}},xe=Object.prototype.toString,Ie=Object.prototype.hasOwnProperty,Se={0:"\\0",7:"\\a",8:"\\b",9:"\\t",10:"\\n",11:"\\v",12:"\\f",13:"\\r",27:"\\e",34:'\\"',92:"\\\\",133:"\\N",160:"\\_",8232:"\\L",8233:"\\P"},Oe=["y","Y","yes","Yes","YES","on","On","ON","n","N","no","No","NO","off","Off","OFF"],je=/^[-+]?[0-9_]+(?::[0-9_]+)+(?:\.[0-9_]*)?$/;function Te(e){var n,i,o;if(n=e.toString(16).toUpperCase(),e<=255)i="x",o=2;else if(e<=65535)i="u",o=4;else{if(!(e<=4294967295))throw new r("code point within a string may not be greater than 0xFFFFFFFF");i="U",o=8}return"\\"+i+t.repeat("0",o-n.length)+n}function Ne(e){this.schema=e.schema||B,this.indent=Math.max(1,e.indent||2),this.noArrayIndent=e.noArrayIndent||!1,this.skipInvalid=e.skipInvalid||!1,this.flowLevel=t.isNothing(e.flowLevel)?-1:e.flowLevel,this.styleMap=function(e,t){var n,i,r,o,a,l,c;if(null===t)return{};for(n={},r=0,o=(i=Object.keys(t)).length;r<o;r+=1)a=i[r],l=String(t[a]),"!!"===a.slice(0,2)&&(a="tag:yaml.org,2002:"+a.slice(2)),(c=e.compiledTypeMap.fallback[a])&&Ie.call(c.styleAliases,l)&&(l=c.styleAliases[l]),n[a]=l;return n}(this.schema,e.styles||null),this.sortKeys=e.sortKeys||!1,this.lineWidth=e.lineWidth||80,this.noRefs=e.noRefs||!1,this.noCompatMode=e.noCompatMode||!1,this.condenseFlow=e.condenseFlow||!1,this.quotingType='"'===e.quotingType?2:1,this.forceQuotes=e.forceQuotes||!1,this.replacer="function"==typeof e.replacer?e.replacer:null,this.implicitTypes=this.schema.compiledImplicit,this.explicitTypes=this.schema.compiledExplicit,this.tag=null,this.result="",this.duplicates=[],this.usedDuplicates=null}function Fe(e,n){for(var i,r=t.repeat(" ",n),o=0,a=-1,l="",c=e.length;o<c;)-1===(a=e.indexOf("\n",o))?(i=e.slice(o),o=c):(i=e.slice(o,a+1),o=a+1),i.length&&"\n"!==i&&(l+=r),l+=i;return l}function Ee(e,n){return"\n"+t.repeat(" ",e.indent*n)}function Me(e){return 32===e||9===e}function Le(e){return 32<=e&&e<=126||161<=e&&e<=55295&&8232!==e&&8233!==e||57344<=e&&e<=65533&&65279!==e||65536<=e&&e<=1114111}function _e(e){return Le(e)&&65279!==e&&13!==e&&10!==e}function De(e,t,n){var i=_e(e),r=i&&!Me(e);return(n?i:i&&44!==e&&91!==e&&93!==e&&123!==e&&125!==e)&&35!==e&&!(58===t&&!r)||_e(t)&&!Me(t)&&35===e||58===t&&r}function Ue(e,t){var n,i=e.charCodeAt(t);return i>=55296&&i<=56319&&t+1<e.length&&(n=e.charCodeAt(t+1))>=56320&&n<=57343?1024*(i-55296)+n-56320+65536:i}function qe(e){return/^\n* /.test(e)}function Ye(e,t,n,i,r,o,a,l){var c,s,u=0,p=null,f=!1,d=!1,h=-1!==i,g=-1,m=Le(s=Ue(e,0))&&65279!==s&&!Me(s)&&45!==s&&63!==s&&58!==s&&44!==s&&91!==s&&93!==s&&123!==s&&125!==s&&35!==s&&38!==s&&42!==s&&33!==s&&124!==s&&61!==s&&62!==s&&39!==s&&34!==s&&37!==s&&64!==s&&96!==s&&function(e){return!Me(e)&&58!==e}(Ue(e,e.length-1));if(t||a)for(c=0;c<e.length;u>=65536?c+=2:c++){if(!Le(u=Ue(e,c)))return 5;m=m&&De(u,p,l),p=u}else{for(c=0;c<e.length;u>=65536?c+=2:c++){if(10===(u=Ue(e,c)))f=!0,h&&(d=d||c-g-1>i&&" "!==e[g+1],g=c);else if(!Le(u))return 5;m=m&&De(u,p,l),p=u}d=d||h&&c-g-1>i&&" "!==e[g+1]}return f||d?n>9&&qe(e)?5:a?2===o?5:2:d?4:3:!m||a||r(e)?2===o?5:2:1}function Re(e,t,n,i,o){e.dump=function(){if(0===t.length)return 2===e.quotingType?'""':"''";if(!e.noCompatMode&&(-1!==Oe.indexOf(t)||je.test(t)))return 2===e.quotingType?'"'+t+'"':"'"+t+"'";var a=e.indent*Math.max(1,n),l=-1===e.lineWidth?-1:Math.max(Math.min(e.lineWidth,40),e.lineWidth-a),c=i||e.flowLevel>-1&&n>=e.flowLevel;switch(Ye(t,c,e.indent,l,(function(t){return function(e,t){var n,i;for(n=0,i=e.implicitTypes.length;n<i;n+=1)if(e.implicitTypes[n].resolve(t))return!0;return!1}(e,t)}),e.quotingType,e.forceQuotes&&!i,o)){case 1:return t;case 2:return"'"+t.replace(/'/g,"''")+"'";case 3:return"|"+Be(t,e.indent)+Ke(Fe(t,a));case 4:return">"+Be(t,e.indent)+Ke(Fe(function(e,t){var n,i,r=/(\n+)([^\n]*)/g,o=(l=e.indexOf("\n"),l=-1!==l?l:e.length,r.lastIndex=l,We(e.slice(0,l),t)),a="\n"===e[0]||" "===e[0];var l;for(;i=r.exec(e);){var c=i[1],s=i[2];n=" "===s[0],o+=c+(a||n||""===s?"":"\n")+We(s,t),a=n}return o}(t,l),a));case 5:return'"'+function(e){for(var t,n="",i=0,r=0;r<e.length;i>=65536?r+=2:r++)i=Ue(e,r),!(t=Se[i])&&Le(i)?(n+=e[r],i>=65536&&(n+=e[r+1])):n+=t||Te(i);return n}(t)+'"';default:throw new r("impossible error: invalid scalar style")}}()}function Be(e,t){var n=qe(e)?String(t):"",i="\n"===e[e.length-1];return n+(i&&("\n"===e[e.length-2]||"\n"===e)?"+":i?"":"-")+"\n"}function Ke(e){return"\n"===e[e.length-1]?e.slice(0,-1):e}function We(e,t){if(""===e||" "===e[0])return e;for(var n,i,r=/ [^ ]/g,o=0,a=0,l=0,c="";n=r.exec(e);)(l=n.index)-o>t&&(i=a>o?a:l,c+="\n"+e.slice(o,i),o=i+1),a=l;return c+="\n",e.length-o>t&&a>o?c+=e.slice(o,a)+"\n"+e.slice(a+1):c+=e.slice(o),c.slice(1)}function He(e,t,n,i){var r,o,a,l="",c=e.tag;for(r=0,o=n.length;r<o;r+=1)a=n[r],e.replacer&&(a=e.replacer.call(n,String(r),a)),($e(e,t+1,a,!0,!0,!1,!0)||void 0===a&&$e(e,t+1,null,!0,!0,!1,!0))&&(i&&""===l||(l+=Ee(e,t)),e.dump&&10===e.dump.charCodeAt(0)?l+="-":l+="- ",l+=e.dump);e.tag=c,e.dump=l||"[]"}function Pe(e,t,n){var i,o,a,l,c,s;for(a=0,l=(o=n?e.explicitTypes:e.implicitTypes).length;a<l;a+=1)if(((c=o[a]).instanceOf||c.predicate)&&(!c.instanceOf||"object"==typeof t&&t instanceof c.instanceOf)&&(!c.predicate||c.predicate(t))){if(n?c.multi&&c.representName?e.tag=c.representName(t):e.tag=c.tag:e.tag="?",c.represent){if(s=e.styleMap[c.tag]||c.defaultStyle,"[object Function]"===xe.call(c.represent))i=c.represent(t,s);else{if(!Ie.call(c.represent,s))throw new r("!<"+c.tag+'> tag resolver accepts not "'+s+'" style');i=c.represent[s](t,s)}e.dump=i}return!0}return!1}function $e(e,t,n,i,o,a,l){e.tag=null,e.dump=n,Pe(e,n,!1)||Pe(e,n,!0);var c,s=xe.call(e.dump),u=i;i&&(i=e.flowLevel<0||e.flowLevel>t);var p,f,d="[object Object]"===s||"[object Array]"===s;if(d&&(f=-1!==(p=e.duplicates.indexOf(n))),(null!==e.tag&&"?"!==e.tag||f||2!==e.indent&&t>0)&&(o=!1),f&&e.usedDuplicates[p])e.dump="*ref_"+p;else{if(d&&f&&!e.usedDuplicates[p]&&(e.usedDuplicates[p]=!0),"[object Object]"===s)i&&0!==Object.keys(e.dump).length?(!function(e,t,n,i){var o,a,l,c,s,u,p="",f=e.tag,d=Object.keys(n);if(!0===e.sortKeys)d.sort();else if("function"==typeof e.sortKeys)d.sort(e.sortKeys);else if(e.sortKeys)throw new r("sortKeys must be a boolean or a function");for(o=0,a=d.length;o<a;o+=1)u="",i&&""===p||(u+=Ee(e,t)),c=n[l=d[o]],e.replacer&&(c=e.replacer.call(n,l,c)),$e(e,t+1,l,!0,!0,!0)&&((s=null!==e.tag&&"?"!==e.tag||e.dump&&e.dump.length>1024)&&(e.dump&&10===e.dump.charCodeAt(0)?u+="?":u+="? "),u+=e.dump,s&&(u+=Ee(e,t)),$e(e,t+1,c,!0,s)&&(e.dump&&10===e.dump.charCodeAt(0)?u+=":":u+=": ",p+=u+=e.dump));e.tag=f,e.dump=p||"{}"}(e,t,e.dump,o),f&&(e.dump="&ref_"+p+e.dump)):(!function(e,t,n){var i,r,o,a,l,c="",s=e.tag,u=Object.keys(n);for(i=0,r=u.length;i<r;i+=1)l="",""!==c&&(l+=", "),e.condenseFlow&&(l+='"'),a=n[o=u[i]],e.replacer&&(a=e.replacer.call(n,o,a)),$e(e,t,o,!1,!1)&&(e.dump.length>1024&&(l+="? "),l+=e.dump+(e.condenseFlow?'"':"")+":"+(e.condenseFlow?"":" "),$e(e,t,a,!1,!1)&&(c+=l+=e.dump));e.tag=s,e.dump="{"+c+"}"}(e,t,e.dump),f&&(e.dump="&ref_"+p+" "+e.dump));else if("[object Array]"===s)i&&0!==e.dump.length?(e.noArrayIndent&&!l&&t>0?He(e,t-1,e.dump,o):He(e,t,e.dump,o),f&&(e.dump="&ref_"+p+e.dump)):(!function(e,t,n){var i,r,o,a="",l=e.tag;for(i=0,r=n.length;i<r;i+=1)o=n[i],e.replacer&&(o=e.replacer.call(n,String(i),o)),($e(e,t,o,!1,!1)||void 0===o&&$e(e,t,null,!1,!1))&&(""!==a&&(a+=","+(e.condenseFlow?"":" ")),a+=e.dump);e.tag=l,e.dump="["+a+"]"}(e,t,e.dump),f&&(e.dump="&ref_"+p+" "+e.dump));else{if("[object String]"!==s){if("[object Undefined]"===s)return!1;if(e.skipInvalid)return!1;throw new r("unacceptable kind of an object to dump "+s)}"?"!==e.tag&&Re(e,e.dump,t,a,u)}null!==e.tag&&"?"!==e.tag&&(c=encodeURI("!"===e.tag[0]?e.tag.slice(1):e.tag).replace(/!/g,"%21"),c="!"===e.tag[0]?"!"+c:"tag:yaml.org,2002:"===c.slice(0,18)?"!!"+c.slice(18):"!<"+c+">",e.dump=c+" "+e.dump)}return!0}function Ge(e,t){var n,i,r=[],o=[];for(Ve(e,r,o),n=0,i=o.length;n<i;n+=1)t.duplicates.push(r[o[n]]);t.usedDuplicates=new Array(i)}function Ve(e,t,n){var i,r,o;if(null!==e&&"object"==typeof e)if(-1!==(r=t.indexOf(e)))-1===n.indexOf(r)&&n.push(r);else if(t.push(e),Array.isArray(e))for(r=0,o=e.length;r<o;r+=1)Ve(e[r],t,n);else for(r=0,o=(i=Object.keys(e)).length;r<o;r+=1)Ve(e[i[r]],t,n)}function Ze(e,t){return function(){throw new Error("Function yaml."+e+" is removed in js-yaml 4. Use yaml."+t+" instead, which is now safe by default.")}}var ze=u,Je=d,Qe=y,Xe=S,et=O,tt=B,nt=Ce.load,it=Ce.loadAll,rt={dump:function(e,t){var n=new Ne(t=t||{});n.noRefs||Ge(e,n);var i=e;return n.replacer&&(i=n.replacer.call({"":i},"",i)),$e(n,0,i,!0,!0)?n.dump+"\n":""}}.dump,ot=r,at={binary:M,float:I,map:m,null:b,pairs:q,set:R,timestamp:N,bool:A,int:k,merge:F,omap:D,seq:g,str:h},lt=Ze("safeLoad","load"),ct=Ze("safeLoadAll","loadAll"),st=Ze("safeDump","dump"),ut={Type:ze,Schema:Je,FAILSAFE_SCHEMA:Qe,JSON_SCHEMA:Xe,CORE_SCHEMA:et,DEFAULT_SCHEMA:tt,load:nt,loadAll:it,dump:rt,YAMLException:ot,types:at,safeLoad:lt,safeLoadAll:ct,safeDump:st},pt=Object.freeze({__proto__:null,[Symbol.toStringTag]:"Module",default:ut,CORE_SCHEMA:et,DEFAULT_SCHEMA:tt,FAILSAFE_SCHEMA:Qe,JSON_SCHEMA:Xe,Schema:Je,Type:ze,YAMLException:ot,dump:rt,load:nt,loadAll:it,safeDump:st,safeLoad:lt,safeLoadAll:ct,types:at});export{pt as j};
