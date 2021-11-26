# 数据定义

模式，表，视图，索引

## 模式的定义与删除

create scheme <模式名> authorization <用户名>

//如果语句里面没有模式名，则默认模式名为用户名

drop scheme <模式名> <cascade|restrict>

//cascade：表示删除模式时同时把该模式下的所有数据库对象全部删除

//restrict:表示如果该模式下已经定义了数据库对象，则拒绝删除该模式