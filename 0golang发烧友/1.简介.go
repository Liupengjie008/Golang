Beego框架默认有captcha验证码插件。在utils/captcha下面

使用方法

import(
    "github.com/astaxie/beego/cache"
    "github.com/astaxie/beego/utils/captcha"
）

var cpt *captcha.Captcha

func init() {
    store := cache.NewMemoryCache()
    cpt = captcha.NewWithFilter("/captcha/", store) //一定要写在构造函数里面，要不然第一次打开页面有可能是X
}

在模板里面写上   

{{create_captcha}}
就ok了，最贴心的是居然连onclick事件也已经做在了里面，方便。

默认的验证码是6位，200px宽，这个是可以自己设置的

cpt是一个结构体：

// Captcha struct
type Captcha struct {
    // beego cache store
    store cache.Cache

    // url prefix for captcha image
    URLPrefix string

    // specify captcha id input field name
    FieldIdName string
    // specify captcha result input field name
    FieldCaptchaName string

    // captcha image width and height
    StdWidth  int
    StdHeight int

    // captcha chars nums
    ChallengeNums int

    // captcha expiration seconds
    Expiration int64

    // cache key prefix
    CachePrefix string
}
 

func init() {
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
	cpt.ChallengeNums = 4  // 设置字数
	cpt.StdWidth = 100     // 设置宽度
	cpt.StdHeight = 40     // 设置高度
}
 

