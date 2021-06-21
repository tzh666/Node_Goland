CSRF  网络攻击
    扩展请求伪造
开启CSRF防护
   a、 配置Token key => 过期时间
       过期时间
           EnableXSRF=true
           XSRFKey=d12e9b0a277a2b34d3ccaf53075489fe
           XSRFExpire=36000

   b、 打开页面生成Token
       从Controller生成，提交到页面
                方式一：
                c.Data["xsrf_input"] = template.HTML(c.XSRFFormHTML())
                方式二：
                c.Data["xsrf_token"] = c.XSRFToken()

   c、 提交数据提交Token
       beego自动验证（POST DELETE PUT三种请求的时候）
                方式一：
                csrf_token => {{ .xsrf_input }}
                方式二：
                自己写input标签
                <input type="hidden" name="_xsrf" value="{{ .xsrf_token }}" />

   
   d、如果没加的话 会报错 ('_xsrf' argument missing from POST)

日志：
const (
　　　　LevelEmergency = iota        // 紧急级别
　　　　LevelAlert                   // 报警级别
　　　　LevelCritical                // 严重错误级别
　　　　LevelError                   // 错误级别
　　　　LevelWarning                 // 警告级别
　　　　LevelNotice                  // 注意级别
　　　　LevelInformational           // 报告级别
　　　　LevelDebug                   // 除错级别
)


Flash:
    存储消息,修改成功后的给用户一个提示:  消息传递过去会删除这个cookie
        // 这个消息存在cookie中,Set-Cookie: BEEGO_FLASH=%00xxxx
		flash := beego.NewFlash()
		// key value 形式
		flash.Set("notice", "修改用户信息成功")
		// c.Controller是beego的Controller
		flash.Store(&c.Controller)
    模板使用：
     {{ if .flash.notice }}
        <div class="row">
            <div class="col-3 offset-4">
                <div class="alert alert-danger" role="alert">
                    <p>{{ .flash.notice }}</p>
                </div>
            </div>
        </div>
        {{ end }}


缓存:
