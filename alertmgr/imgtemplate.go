package alertmgr

var ImgTmpl = `<!DOCTYPE html>
<html>
<head> 
<meta charset="utf-8"> 
<title>{{.Title}}</title> 
</head>
<body>

<h2>{{.Title}}</h2>
<img border="0" src="{{.Image}}" alt="展示拆线图" width="{{.Width}}" height="{{.Height}}">

</body>
</html>`
