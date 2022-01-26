## repo
https://github.com/casbin/casbin

## 权限管理
支持常用的多种访问控制模型，如ACL/RBAC/ABAC等

## 权限管理link
https://zhuanlan.zhihu.com/p/70548562

## PERM模型
PERM（Policy，Effect，Request，Matchers） 元模型的配置文件。

- policy是策略或者说是规则的定义。它定义了具体的规则。
- request是对访问请求的抽象，它与e.Enforce()函数的参数是一一对应的
- matcher匹配器会将请求与定义的每个policy一一匹配，生成多个匹配结果。
- effect根据对请求运用匹配器得出的所有结果进行汇总，来决定该请求是允许还是拒绝。

PERM中又分为访问实体(Subject)，访问资源(Object) 和访问方法(Action)

## ACL
不同角色不同操作采用不同策略
缺点：ACL模型在用户和资源都比较少的情况下没什么问题，但是用户和资源量一大，ACL就会变得异常繁琐，每次增加一个用户需要重新设置他的各类权限

## RBAC
（role-based-access-control）引入role中间层来解决问题。每个用户都属于一个角色，例如开发者、管理员、运维等，每个角色都有其特定的权限，权限的增加和删除都通过角色来进行。这样新增一个用户时，我们只需要给他指派一个角色，他就能拥有该角色的所有权限。修改角色的权限时，属于这个角色的用户权限就会相应的修改。

在casbin中使用RBAC模型需要在模型文件中添加role_definition模块：
``` conf
[role_definition]
g = _,_

[matchers]
m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
```
定义了用户——角色，角色——角色的映射关系，前者是后者的成员，拥有后者的权限。然后在匹配器中，我们不需要判断r.sub与p.sub完全相等，只需要使用g(r.sub, p.sub)来判断请求主体r.sub是否属于p.sub这个角色即可。

### 多个RBAC
增加多个role_definition
``` conf
[role_definition]
g = _,_
g2 = _,_
```

### 多个角色

## ABAC
RBAC模型对于实现比较规则的、相对静态的权限管理非常有用。但是对于特殊的、动态的需求，RBAC就显得有点力不从心了。
不同的时间段对数据data实现不同的权限控制。正常工作时间9:00-18:00所有人都可以读写data，其他时间只有数据所有者能读写。这种需求我们可以很方便地使用ABAC（attribute base access list）模型完成：

## 模型存储
上面存储策略都存在model.conf文件中，也可以在代码中动态初始化

## 策略存储
上面存储策略都存在policy.csv文件中，实际情况应该在数据库中

## 使用函数
我们可以在匹配器中使用函数。casbin内置了一些函数keyMatch/keyMatch2/keyMatch3/keyMatch4都是匹配 URL 路径的，regexMatch使用正则匹配，ipMatch匹配 IP 地址。   
参见https://casbin.org/docs/en/function。使用内置函数我们能很容易对路由进行权限划分：