# qqchat
gin框架+websocket聊天  

## 初始化系统   
第一步：``  编写获取配置函数InitConfig，初始化到系统中``    
第二步：``  编写日志初始化函数InitLogger,和编写日志LoggerMiddleware记录请求日志；``     
第三步：``  编写连接数据库InitGorm``    
第四步： ``  整理路由``    
第五步： ``  编写swagger``   

## 具体的系统路径访问处理  
第一步： 处理请求校验-验证器 和 返回  
``  请求校验：有俩部分，一手动校验，二自定义校验规则``    
<!--       请求校验：自定义方法              -->
        ValidateRequest        校验表单请求参数      ShouldBind 方式  
        ValidateJSONRequest    校验JSON请求参数     ShouldBindJSON方式      
        ValidateQueryRequest   校验请求路径上的参数   ShouldBindQuery方式 http://localhost:8080/search?name=John&age=30这种参数     
        ValidateHeaderRequest  校验Header头请求参数  ShouldBindHeader 方式 
``  返回： 需要定义正确返回和错误返回 ``

第二步：``  编写系统异常捕获处理 ErrorHandlerMiddleware``

第三步: ``  编写jwt; 然后编写上下文context; 然后编写鉴权校验 AuthMiddleware中间件``   
<!--       流程：              -->
    登录把个人信息通过jwt写入token；  
    然后下次请求时，通过AuthMiddleware中间件，校验token的jwt解析获取个人信息，写入context；  
    然后再具体的请求方法中从context中拿到登录的个人信息。  





## gin框架文档
https://github.com/gin-gonic/gin  
https://gin-gonic.com/zh-cn/docs/introduction/

## gorm 文档
https://gorm.io/
https://www.kancloud.cn/sliver_horn/gorm/1861153

## gin-swagger
链接里面有使用操作手册： https://github.com/swaggo/gin-swagger  
注释的规范格式： https://github.com/swaggo/swag/blob/master/README.md#declarative-comments-format    
go get -u github.com/swaggo/swag/cmd/swag    
go get -u github.com/swaggo/gin-swagger    
go get -u github.com/swaggo/files    
import "github.com/swaggo/gin-swagger" // gin-swagger middleware   
import "github.com/swaggo/files" // swagger embed files   
生成对应的文件： swag init

## go-playground/validator 验证器文档
https://github.com/go-playground/validator  这里面包含了很多验证参数规则  
【项目中】ValidateRequest() 手动验证加上了中文反馈 ； InitValidator()这个是自定义验证规则注册到了路由中