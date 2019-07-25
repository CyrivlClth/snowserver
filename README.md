# snowserver

## 说明

分布式id生成器节点，可通过http、grpc方式获取id。

## 安装

编译安装后，将`$GOPATH/bin`加入`PATH`。

```shell
go get -u github.com/CyrivlClth/snowserver
```

## 运行

```shell
snowserver -gp=50051 -hp=8080 server

# 获取更多运行信息
snowserver -h
```

## 接口

### HTTP

#### 获取运行信息

URL:`{host}/stats`

METHOD:`GET`

RESPONSE:

```json
{
	"data_center_id": 1,				# 数据中心ID
	"worker_id": 1,						# 工作节点ID
	"start_timestamp": 1558930914000,	# 工作节点计算开始时间戳
	"last_timestamp": 1558950914000,	# 上次时间戳
	"timestamp": 15589509981234,		# 工作节点当前时间戳
	"sequence": 400,					# 工作节点当前序列号
	"sequence_overload": 3,				# 工作节点毫秒内序列重置次数
	"errors_count": 0					# 工作节点时间回调错误发生次数
}
```

#### 获取ID

URL：`{host}/`

METHOD:`GET`

RESPONSE:

```json
{
	"id": "14622164842135477"
}
```

#### 获取多个ID

URL:`{host}/count/{:count}`

METHOD:`GET`

RESPONSE:

```json
{
	"count": 5,
	"ids": [
	"14622164842135477",
	"14622164842135478",
	"14622164842135479",
	"14622164842135480",
	"14622164842135481"
	]
}
```

### GRPC

#### 使用

`grpc/pb/snow.proto`可生成各语言代码，使用客户端请求服务器grpc端口即可

#### 获取运行信息

服务名：`Stats`

说明：返回节点运行信息

#### 获取ID

服务名：`NextID`

说明：实时获取一个ID

#### 获取多个ID

服务名：`GetIDs`

说明：一次性获取多个ID
