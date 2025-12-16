一个基于 Go + Gin + GORM 的个人播客博客系统，支持用户注册/登录（JWT 认证）、文章管理、评论功能，并具备完善的错误处理与日志记录。

### 功能特性
✅ 用户系统：注册、登录（密码加密存储）

🔐 JWT 认证：安全的 Token 鉴权机制

📝 文章管理：创建、查询、更新、删除（含软删除）

💬 评论系统：按文章关联评论，支持软删除

🛡️ 统一错误处理：标准化错误码与消息

📊 结构化日志：请求日志 + 错误追踪（Zap + 日志轮转）

📦 统一响应格式：{ code, message, data }

### 技术栈

类别 | 技术
---|---
语言| Go 1.23+
Web 框架 | Gin
ORM | GORM
密码加密 | golang.org/x/crypto/bcrypt
JWT | golang-jwt/jwt/v4
日志| Uber Zap + lumberjack
依赖管理| Go Modules

### 运行环境要求
- Go: 1.23+ 或更高版本
安装：https://golang.org/dl/
- MySQL: 8.0+（需支持 utf8mb4 字符集）
推荐使用 Docker 快速启动：
```bash
docker run --name myblog-mysql -e MYSQL_ROOT_PASSWORD=admin123 -e MYSQL_DATABASE=myblog -p 3306:3306 -d mysql:latest
```
- 操作系统: Linux / macOS / Windows
  
###  快速开始

1. 克隆项目到本地：
```bash
git clone https://github.com/martintree/go-homework.git
cd task4/server
```
2. 安装依赖：
```bash
go mod tidy
```
> 会自动下载以下关键依赖：
> - github.com/gin-gonic/gin
> - gorm.io/gorm
> - gorm.io/driver/mysql
> - github.com/golang-jwt/jwt/v4
> - go.uber.org/zap
> - gopkg.in/natefinch/lumberjack.v2

3. 配置数据库：
修改 `config/config.go` 文件中的数据库配置项：
```go
func GetDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Host:               "127.0.0.1",
		Port:               "3306",
		User:               "root",
		Password:           "admin123",
		Name:               "myblog",
		Charset:            "utf8mb4",
		MaxConnections:     10,
		MaxIdleConnections: 5,
	}
}
```
4. 配置 JWT 密钥：
在 `config/config.go`  中设置密钥：
```go
var jwtSecretKey = "udia#y387dyJkanadk7&54"
```
5. 运行服务：
```bash
go run main.go
```   
默认监听 http://localhost:8080

###  项目结构
```text
server/
├── main.go                  # 入口文件
├── config/
│   └── config.go            # 数据库及密钥配置
├── db/
│   └── db.go                # 初始化数据库
├── dto/
│   ├── user_dto.go          # user DTO
│   ├── comment_dto.go       # comment DTO
│   └── post_dto.go          # post DTO
├── models/
│   ├── user.go              # 用户模型
│   ├── post.go              # 文章模型
│   └── comment.go           # 评论模型
├── handlers/
│   ├── user.go              # 注册/登录逻辑
│   ├── post.go              # 文章管理
│   └── comment.go           # 评论管理
├── middleware/
│   ├── jwt_auth.go          # JWT 认证中间件
│   ├── error_handler.go     # error处理中间件
│   └── logger.go            # 请求日志中间件
├── router/
│   └── routes.go            # 路由配置
├── utils/
│   ├── jwt.go               # JWT 工具
│   ├── logger.go            # 日志封装函数
│   ├── response.go          # 响应封装函数
│   └── error.go             # 自定义错误类型
├── logs/                    # 自动生成的日志目录
├── .gitignore               # gitignore文件
├── go.mod                   # go.mod文件
├── go.sum                   # go.sum文件
├── request.http             # 请求用例文件（基于rest client）
└── README.md                # readme文件
``` 

###  API 请求示例
1. 用户注册
```bash
curl -X POST http://localhost:8080/api/v1/users/register \
  -H "Content-Type: application/json" \
  -d '{"username":"alice","email":"alice@example.com","password":"123456"}'
```
2. 用户登录
 ```bash
curl -X POST http://localhost:8080/api/v1/users/login \
  -H "Content-Type: application/json" \
  -d '{"username":"alice","password":"123456"}'
```
3. 新增文章
```bash
curl -X POST http://localhost:8080/api/v1/auth/posts \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your_jwt_token>" \
  -d '{"title":"计算机的主要特点","content":"运算速度快 (High Speed): 计算机能以极高的速度（纳秒级）进行计算，比人工快无数倍。"}'
```  
4. 获取有个用户所有文章
```bash
curl -X GET http://localhost:8080/api/v1/auth/users/posts \
  -H "Authorization: Bearer <your_jwt_token>"
```  
5. 获取某个文章
```bash
curl -X GET http://localhost:8080/api/v1/auth/posts/2 \
  -H "Authorization: Bearer <your_jwt_token>"
```  
6. 修改某个文章
```bash
curl -X PUT http://localhost:8080/api/v1/auth/posts \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your_jwt_token>" \
  -d '{ "id":1,"title":"如何使用gorm从数据库返回新创建的记录","content":"我有一个创建新用户的函数，但是获取用户值的推荐方法不包括由数据库(id，created_at)创建的自动生成的值"}'
```     
7. 删除某个文章
```bash
curl -X DELETE http://localhost:8080/api/v1/auth/posts/1 \
  -H "Authorization: Bearer <your_jwt_token>" 
```  
8. 新增评论
```bash
curl -X POST http://localhost:8080/api/v1/auth/comments \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your_jwt_token>" \
  -d '{"postId":1,"content":"写的不错哦"}'
```
9. 删除评论
```bash
curl -X DELETE  http://localhost:8080/api/v1/auth/comments/2 \
  -H "Authorization: Bearer <your_jwt_token>" 
```  
###  日志说明

- 控制台日志：开发时输出彩色日志（INFO/WARN/ERROR）
- 文件日志：自动写入 logs/app.log，每日轮转，保留日志30天，保留最近 5 份，单个文件最大 10MB
- 错误日志：包含完整堆栈，便于排查问题
