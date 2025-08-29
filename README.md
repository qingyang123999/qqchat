# qqchat

gin框架+websocket聊天

## gin框架文档

https://gin-gonic.com/zh-cn/docs/introduction/

## gorm 文档

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




