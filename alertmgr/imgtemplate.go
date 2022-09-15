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

var TestTmpl = `<!DOCTYPE html>
<html>
<head> 
<meta charset="utf-8"> 
<title>测试邮件</title> 
</head>
<body>

<h2>测试邮件</h2>

<span>发送的测试邮件</span> - 
<span>{{.Date}}</span>

</body>
</html>`
