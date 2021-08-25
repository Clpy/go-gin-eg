目录结构

go-gin-eg
- conf：用于存储配置文件
- middleware：应用中间件
- models：应用数据库模型
- pkg：第三方包
- routers 路由逻辑处理
- runtime：应用运行时数据

涉及的知识点

GORM是golang下的ORM框架，处理golang struct与数据库结构之间的映射，能根据struct结构自动生成SQL, 方便的CRUD操作，并且支持表关联。


gorm所支持的回调方法：

创建：BeforeSave、BeforeCreate、AfterCreate、AfterSave
更新：BeforeSave、BeforeUpdate、AfterUpdate、AfterSave
删除：BeforeDelete、AfterDelete
查询：AfterFind