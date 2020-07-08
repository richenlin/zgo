# 说明

## 图像

resource ---------------->
      |                     user
      |-------role------->
正常用户的权限来自直接资源授权和通过角色继承的授权
但是推荐直接通过角色授权允许权限,而对于临时权限给出直接授权的方式.



## adapter
github.com/memwey/casbin-sqlx-adapter
github.com/casbin/redis-adapter

## multi-threading
https://github.com/casbin/casbin/blob/master/enforcer_synced.go
StartAutoLoadPolicy: 自动刷新策略

## watcher
github.com/billcobbler/casbin-redis-watcher


多线程处理中,可以通过watcher中的callback重新加载,也可以通过synced的StartAutoLoadPolicy方法自动加载
在实现过程中,如果使用redis,推荐使用第一种,但是如果使用sqlx的方式,则推荐使用StartAutoLoadPolicy的方式
注意,刷新的频率不要过于频繁.

## examples
https://github.com/casbin/casbin/tree/master/examples

## db
数据库的列名为V0~V5, 所以这里的列是复用的,本身没有什么固定的含义
由于PType为p或者g, 其代表的意义也不近相同.同时如果系统是基于rbac模式,
其V0, 即可以标识策略内容,也可以表示用户.

table_name: casbin_rule
colume    type       enum
ID        unit
PType     string     p, p2, g, g2
V0        string
V1        string
V2        string
V3        string
V4        string
V5        string

## rbac VS abac
rbac: 基于角色的访问控制

RBAC 系统中的用户名称和角色名称不应相同。因为Casbin将用户名和角色识别为字符串， 
所以当前语境下Casbin无法得出这个字面量到底指代用户 alice 还是角色 alice。 
这时，使用明确的 role_alice ，问题便可迎刃而解。

abac: 基于属性的访问控制
可以使用主体、客体或动作的属性，而不是字符串本身来控制访问
```ini
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub_rule, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = eval(p.sub_rule) && r.obj == p.obj && r.act == p.act
```
```csv
p, r.sub.Age > 18, /data1, read
p, r.sub.Age < 60, /data2, write
```
使用eval和[xxx]_rule, 可以对csv中的规则进行匹配
当然,可以使用如下官方的标准形式
```ini
[matchers]
m = r.sub == r.obj.Owner
```
