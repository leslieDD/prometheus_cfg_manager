import{c as r}from"./lint.d7baeab5.js";var o;(o=r.exports).registerHelper("lint","yaml",(function(r){var n=[];if(!window.jsyaml)return window.console&&window.console.error("Error: window.jsyaml not defined, CodeMirror YAML linting cannot run."),n;try{jsyaml.loadAll(r)}catch(i){var e=i.mark,a=e?o.Pos(e.line,e.column):o.Pos(0,0),s=a;n.push({from:a,to:s,message:i.message})}return n}));
