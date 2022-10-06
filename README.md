# GoLazyFrame 懒人框架

> 偷懒，是程序员的基本守则
>
> - 仓库地址：
>   - github：[https://github.com/go-lazy-frame/go-lazy-frame.git](https://github.com/go-lazy-frame/go-lazy-frame.git)
>   - gitee：[https://gitee.com/go-lazy-frame/gol-lazy-frame.git](https://gitee.com/go-lazy-frame/gol-lazy-frame.git)

[TOC]

## 依赖包网络加速

```shell
# 永久生效
go env -w GOPROXY=https://goproxy.cn,direct
# 环境变量生效
export GOPROXY=https://goproxy.cn
```

## 环境配置

- 配置项

  | 配置项         | 说明                                                         |
  | -------------- | ------------------------------------------------------------ |
  | `projectHome` | 项目根目录的绝对目录路径（或项目启动时的工作目录为项目的根目录），如果不配置，启动时无法自动更新接口文档，例如 `/home/user/go-project/my-app` |
  | `docsPath`    | 需要自动更新接口文档定义的 `go` 文件源码位置，例如 `{ProjectHome}/internal/router/docs.go`，可使用 `{ProjectHome}` 代表上面配置的路径 |
  | `RUN_ENV`  | 有三个值：prod、dev、local，分别代表线上环境、开发环境、本地环境。本地环境代表本地开发时的环境，dev 和 prod 代表不同的运行环境（二进制运行） |
  | `ENV_Xxx` | 启动时，动态根据环境变量 `ENV_Xxx`（`Xxx` 为对应的具体配置名），更新配置，如 `ENV_EnableMon=true`，则代表系统配置`EnableMon`在运行时应用为 `true` |

- 配置方式

  ![image-20211009204418186](http://image.qianlong168.cn/uPic/image-20211009204418186.png)

## 独有特性

- 代码一体化自动生成
- 超有好的接口文档，免除和前端的接口对接工作，例如：
  - ![image-20211125133210425](http://image.qianlong168.cn/20211125/V7oISi.png)
- 接口文档自动生成并更新，无需执行 `swag init`
  - 注意：只有在配置了`project.home`和`docs.path`环境变量情况下，才会更新接口文档。所以，若改变了接口，或者执行了代码生成更新了代码，请本地启动一次项目，使自动更新文档，再编译发布，否则接口文档可能不新鲜（编译操作不会更新接口文档）。

## 性能比较

在相同硬件条件下的路由访问测试：

![image-20211002180553425](http://image.qianlong168.cn/uPic/image-20211002180553425.png)

- x1:本框架的路由注册
- x2：gin 原本的路由注册

可以看出，没有性能损耗

## 项目结构介绍

> 介于以下几点：
>
> - Golang 的包管理不完善，特别是企业的开发，依赖管理很不友好。不过 Golang 的包管理，对于开源项目，却是比较友好
> - 接口文档自动生成。目前业界都是用 swagger 进行文档接口文档的自动化管理。但是 Golang 的 swagger 支持，很不友好，虽然本框架进行了最大程序的自动化集成和改造，但是依然需要手动指定所有源码文件位置（当然，在项目规范下，该操作已自动化进行，无需开发人员介入）
>
> 所以不同的项目，以及框架，都统一在一个项目中。通过目录结构的区分，以及自动代码的生成系统，进行高效率的开发。这样，框架的维护和公共代码的及时同步更新，也能得到及时的应用。

项目结构，总体遵循 [project-layout ](http://github.com/golang-standards/project-layout)目录规范，该规范是目前 Golang 开发者，最公认的目录规范，对于 Golang 语言开发，也是比较合理的目录规范，在此之上，根据项目和框架的实际情况，做一些相应的调整。

## 代码生成（重要）

介于 Golang 的语言特点，Golang 的框架，提高自动化的开发效率，都是通过代码生成的途径，毕竟，Golang 语言的反射机制…..（此处就不吐槽了）。代码自动生成使用：

```shell
# 1. cd 进入到应用项目的同级目录，比如应用目录为 $GOPATH/src/my-app，则 cd 到 $GOPATH/src
# 2. 拉取代码生成系统，若 github 速度慢，也可从 gitee 仓库拉取：
git clone https://github.com/go-lazy-frame/go-lazy-frame-generate.git
# git clone https://gitee.com/go-lazy-frame/go-lazy-frame-generate.git
cd go-lazy-frame-generate
# 3. 查看 README.md 说明
```



## 开发规范（重要）

### 通用规范

1. DTO 和 VO 的定义：
  
   1. DTO：接口请求的参数封装；
   2. VO：接口响应的参数封装
   
2. 公共服务方法，例如 util 等，使用结构体方法进行定义，并且结构体私有，同时使用公开的变量供外部访问，例如：

   ```go
   var (
     // 供外部调用
     MyUtil = new(myUtil)
   )
   type myUtil struct {}
   func (me *myUtil) MyFun() {}
   ```

3. 每个表设计时，必须有以下几个字段和定义：

   ```sql
   create table xxx_xxx
   (
       id         bigint unsigned auto_increment primary key,
       created_at datetime(3) null,
       updated_at datetime(3) null,
       deleted_at datetime(3) null,
       index `index_deleted_at` (`deleted_at`)
   ) ENGINE = InnoDB
     DEFAULT CHARSET = utf8mb4
     COLLATE = utf8mb4_0900_ai_ci COMMENT ='';
   ```

4. 表字段规范：

   1. 如果是关联其他表 `ID` 的字段，列名必须用 `_id`结尾，若不遵守该规则，在代码自动生成时，会影响字段的类型定义。
   2. 主键列一律使用 `bigint` 类型

5. 不建议修改编辑 `gen_` 开头的代码文件（若有特殊的业务情况要修改，修改操作也是生效的，不过要防止被再次生成覆盖，以及表结构变更时注意手动维护代码）：

   1. 该代码文件内有 `Code generated by LazyFrame Gen tool. DO NOT EDIT.`的提示注释，并且如果编辑，开发工具也会有提示：

![image-20211009193046118](http://image.qianlong168.cn/uPic/image-20211009193046118.png)

2. 数据库表（实体）结构变更：
   1. 修改对应的表结构，然后通过代码生成工具进行自动代码生成，变更代码实体以及相关的基础代码，请勿直接修改实体
3. 创建表时，必须要有表注释，简短就行，不宜过长，否则自动生成的接口文档不友好

### Controller 规范

> 代码生成工具，会自动给每一个实体（表）自动生成一个 controller，并同时注册到路由中。具体代码，可参考自动生成的代码。

1. 一个 `controller` 即为定义一个`struct`，每个接口都是一个 `struct` 方

2. `controller struct` 要注册在 `router.go` 的路由中

3. 每一个`struct`接口方法，都要对应一个 `WebXxx` 的字段定义（其中的 `Xxx` 为对应的 `struct` 方法，既访问的接口方法），用于描述该接口的定义，如：

   1. ```go
      type SandCar struct {
      	...
      	WebCreate interface{} `url:"/sand_car/create" method:"post"`
        ...
      }
      
      func (me SandCar) Create(c *gin.Context) {
      }
      ```

4. 接口方法，需要用注释进行接口文档定义，项目启动时，会自动更新接口文档，如（注意`@id`、`@Router`这两个的对应调整）：

   1. ```go
      // Create
      // 注意：以下的 id 必须设置，且必须全局唯一，否则接口文档页面无法正常显示
      // @id SandCarCreateUsingPOST
      // @Tags 车辆
      // @Summary 车辆创建
      // @Description 车辆创建
      // @Accept json
      // @Produce  json
      // @Param token header string true "登陆成功后的授权 Token，后续的所有接口header，都要带上 token"
      // @Param request body dto.SandCarDto{} true "创建"
      // 		参数名 参数类型 参数对象类型 是否必传 描述
      // @Success 200 {object} web.ResponseResult{}
      // @Failure 500 {object} web.ResponseResult{}
      // @Router /sand_car/create [post]
      func (me SandCar) Create(c *gin.Context) {
      	d := dto.SandCarDto{}
      	if err := c.ShouldBind(&d); err != nil {
      		me.Failed(c, err.Error())
      		return
      	}
      	err := service.SandCarService.CreateSandCar(d)
      	if err != nil {
      		me.Failed(c, err.Error())
      		return
      	}
      	me.Success(c, "OK")
      }
      ```

5. 除 GET 方法外，统一使用POST body 体传 JSON 的方式，进行参数传递，并且结构体的对象数据获取，统一使用结构体方法：`me.BindBodyJson(c, &d)`，如：

   ```golang
   ...	
   d := RbacLogUpdateDto{}
   if err := me.BindBodyJson(c, &d); err != nil {
     me.Failed(c, err.Error())
     return
   }
   ...
   ```

   
