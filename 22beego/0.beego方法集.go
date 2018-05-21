beego.Controller

Abort func(code string)
code : HTTP状态信息
中止此次请求并抛出异常，之后的代码不会再执行，而且会默认显示给用户404页面


Redirect func(url string, code int)
url : 跳转路径
code : HTTP状态码
通过 Redirect 方法来进行跳转，重定向。




