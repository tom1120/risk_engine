# 决策引擎提供接口 HTTP API
- [字段定义](#字段定义)
- [错误码定义](#错误码定义)
- [获取支持的决策流列表](#决策流列表)
- [执行决策流](#决策流执行)


### 字段定义

| 字段     | 说明                                                                                                   |
| -------- | ------------------------------------------------------------------------------------------------------ |
| key      | 决策流唯一标识 key                                                                                     |
| version  | 决策流版本号                                                                                           |


### 错误码定义

| 错误码 | 说明           |
| ------ | -------------- |
| 200    | 成功           |
| -400   | 请求参数错误   |
| -404   | 决策流不存在   |
| -500   | 未知错误       |

### 决策流列表

*HTTP*

GET http://HOST/engine/list

*请求参数*


*返回结果*

```json
{
	"code": 200,
	"error": "",
	"result": [{
		"key": "flow_ruleset",
		"version": "1.0",
		"md5": "46ef45be94e2f3c1ca23cc4ce9fe8947"
	}, {
		"key": "flow_abtest",
		"version": "1.0",
		"md5": "886ac3943749e9cf27b5033aad84a9b9"
	}, {
		"key": "flow_conditional",
		"version": "1.0",
		"md5": "df7a8a9431e38f1729d8d6788987df23"
	}, {
		"key": "flow_matrix",
		"version": "1.0",
		"md5": "37a47768ccb2b48d8585a16693d2279f"
	}]
}
```

*CURL*
```shell
curl http://localhost:8889/engine/list
```

### 决策流执行

*HTTP*

POST http://HOST/engine/run

*请求参数*

| 参数名   | 必选  | 类型              | 说明                             |
| -------- | ----- | ----------------- | -------------------------------- |
| key      | true  | string            | 决策流唯一标识                   |
| version  | true  | string            | 决策流版本                       |
| req_id   | true  | string            | 本次请求id                       |
| uid      | true  | string            | 用户id                           |
| features | false | json              | 依赖特征值列表                   |

*返回结果*

```json
*****成功*****
{
    "code":0,
    "message":""
}
****失败****
{
    "code":-400,
    "message":"-400"
}
```

- features 所有特征值，包括执行过程中产生的衍生特征和赋值特征
- hit_rules 命中的规则列表
- tracks 流执行轨迹
- node_results 各节点执行情况和产生值

*CURL*
```shell
curl -XPOST http://localhost:8889/engine/run -d '{"key":"flow_abtest", "version":"1.0", "req_id":"123456", "uid":1,"features":{"feature_1":5,"feature_2":3,"feature_3":55,"feature_4":32,"feature_5":33,"feature_6":231,"feature_7":2,"feature_8":4}}'
```

### demo/ruleset
### demo/abtest
### demo/conditional
### demo/matrix
### demo/tree
### demo/scorecard

