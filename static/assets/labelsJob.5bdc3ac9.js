import{s as t}from"./request.c24bac87.js";function r(r){return t({url:"/v1/job/group",method:"get",params:r})}function a(r){return t({url:"/v1/job/group",method:"put",data:r})}function u(r){return t({url:"/v1/job/group",method:"post",data:r})}function o(r){return t({url:"/v1/job/group",method:"delete",params:r})}function e(r){return t({url:"/v1/job/machines",method:"get",params:{id:r}})}function n(r){return t({url:"/v1/job/group/machines",method:"get",params:{id:r}})}function s(r,a){return t({url:"/v1/job/group/machines",method:"put",params:{id:r},data:a})}function p(r){return t({url:"/v1/job/group/labels",method:"get",params:r})}function m(r,a){return t({url:"/v1/job/group/labels",method:"post",params:{id:r},data:a})}function d(r,a){return t({url:"/v1/job/group/labels",method:"put",params:{id:r},data:a})}function l(r){return t({url:"/v1/job/group/labels",method:"delete",params:r})}function i(r){return t({url:"/v1/job/group/machines-and-labels",method:"get",params:r})}function b(r){return t({url:"/v1/job/group/status",method:"put",data:r})}function c(r){return t({url:"/v1/job/group/labels/status",method:"put",data:r})}export{n as a,p as b,u as c,m as d,d as e,r as f,e as g,a as h,o as i,l as j,i as k,b as l,c as m,s as p};
