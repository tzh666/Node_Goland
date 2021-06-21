用户管理
    1、登录
    2、用户增删拆改

用户信息

部门管理


<!-- 登录成功以后显示页面 -->
url 跳转到用户列表显示页面
这是一个新的需求，所以：(新需求都类似这个步骤)
    1、一个controller  (model.QueryUser)
    2、model(获取用户数据)  (在model层定义查询用户的方法[QueryUser]，供controller调用)
    3、view(html上展示数据) (如果不需要展示数据的，就跳过这个步骤)
    4、router(绑定数据)

<!-- 编辑功能 -->
    1、打开编辑页面
        a、Get id  ---> 查找数据  ----> 渲染数据
    2、提交
        b、Post id/xxx/xxx  跟新数据 ----> 302 到用户列表页面

<!-- 修改自己的密码 -->
controller  ----> 加载页面
views/html
router

提交数据  ------> 
            Form ParseForm  就是把提交的数据解析到一个结构体对象，然后对这个结构体中的数据进行验证
         ------>      
            form := &forms.PasswordModifyForm{}
            c.ParseForm(form)
1、验证旧密码
2、确认新密码
3、新旧密码不能一致
4、大小写以及数字、特殊密码
