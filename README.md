# 决策引擎系统
[![License](https://img.shields.io/:license-apache%202-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![GoDoc](https://godoc.org/github.com/skyhackvip/risk_engine?status.png)](https://godoc.org/github.com/skyhackvip/risk_engine)
[![Go Report Card](https://goreportcard.com/badge/github.com/skyhackvip/risk_engine)](https://goreportcard.com/report/github.com/skyhackvip/risk_engine)

## 开源声明
本项目用于学习和参考，采用 Apache License, Version 2.0 开源，转载使用请说明出处。代码不定期迭代更新，欢迎 Star & Watch，如需交流请添加公众号“技术岁月”。

## 关于版本
公众号讲解版本为 tag v1，由于早期代码规划不合理，已做重构，最新更新稳定版本为 master 分支。

## 决策引擎系统介绍
决策引擎系统，是构建于规则引擎和流程引擎基础上，满足复杂业务决策的一套系统，可用于反欺诈、信用评估、风险决策、推荐系统、精准营销、内容审核等领域。

风控是决策引擎的一个重要应用场景，由于黑产、羊毛党行业的盛行，风控决策引擎在电商、支付、游戏、社交等领域也有了长足的发展，刷单、套现、作弊，凡是和钱相关的业务都离不开风控决策引擎系统的支持保障。目前流行的“智能风控”，即是以决策引擎为核心驱动，以机器学习/AI为大脑，在大数据基础上构建的通用风控能力。

我结合工作实践及个人思考，从业务抽象建模，产品逻辑规划以及最终技术架构和代码实现等方面给出全方位的解决方案。

### 功能列表
- 规则
- 规则集
- 决策树
- 决策表
- 决策矩阵
- 评分卡
- 决策流
- 冠军挑战者 
- 支持特征类型：int、string、bool

## 快速开始
- 环境准备
go version go1.13 +

- Make 编译执行（推荐）
```shell
#下载
git clone https://github.com/skyhackvip/risk_engine
cd risk_engine/

#编译
make build

#启动
make run

#停止
make stop

```

- Go 编译执行
```shell
#下载
git clone https://github.com/skyhackvip/risk_engine
cd risk_engine/

#编译
mkdir -p dist/conf dist/bin
cp cmd/risk_engine/config.yaml dist/conf
cp demo dist/demo -r
GO111MODULE=on CGO_ENABLED=0 go build -o dist/bin/risk_engine cmd/risk_engine/engine.go

#启动
cd dist/
nohup bin/risk_engine -c conf/config.yaml >nohup.out 2>nohup.out &

#停止
pkill -f bin/risk_engine

```

- Docker 容器编译执行
```shell
#下载
git clone https://github.com/skyhackvip/risk_engine
cd risk_engine/

#制作镜像
docker build -t risk_engine:v1 .

#启动镜像
docker run -d --name risk_engine -p 8889:8889 risk_engine:v1

#进入容器
docker exec -it risk_engine /bin/sh

#停止容器
docker stop risk_engine

```


## 支持接口 HTTP API

### 获取支持的所有决策流
- 请求接口：
```shell
curl http://localhost:8889/engine/list -XPOST

```
- 接口返回：
```json
{
	"code": 200,
	"error": "",
	"result": [{
		"key": "flow_long",
		"version": "1.0",
		"md5": "387a3719ab82b4a1b014a6d912a5ebb5"
	}, {
		"key": "flow_simple",
		"version": "1.0",
		"md5": "4b73cb9dfbe3d55e2b80c144fc04f643"
	}, {
		"key": "flow_test",
		"version": "1.0",
		"md5": "69296c88f96bb15d44ba565ed8364d86"
	}, {
		"key": "flow_abtest",
		"version": "1.0",
		"md5": "004a951c5d3b2678ec4c28e21bf0eaaf"
	}]
}
```
目前支持的决策流以文件形式存在于demo/中，启动时加载到内存中

### 执行决策流
- 请求接口：

```shell
curl -XPOST http://localhost:8889/engine/run -d '{"key":"flow_abtest", "version":"1.0", "req_id":"123456", "uid":1,"features":{"feature_1":5,"feature_2":3,"feature_3":55,"feature_4":32,"feature_5":33,"feature_6":231,"feature_7":2,"feature_8":4}}'
```
- key: 决策流标识，目前支持的决策流以文件形式存在于目录 demo/ 中
- version：决策流版本标识
- req_id：请求ID
- uid： 用户ID
- features：入参传入的特征值

- 接口返回：
```json
{
	"code": 200,
	"error": "",
	"result": {
		"key": "flow_abtest",
		"req_id": "123456",
		"uid": 1,
		"features": [{
			"isDefault": false,
			"name": "feature_8",
			"value": 4
		}, {
			"isDefault": false,
			"name": "feature_1",
			"value": 5
		}, {
			"isDefault": false,
			"name": "feature_2",
			"value": 3
		}, {
			"isDefault": false,
			"name": "feature_4",
			"value": 32
		}, {
			"isDefault": false,
			"name": "feat1",
			"value": "aa"
		}, {
			"isDefault": false,
			"name": "feature_5",
			"value": 33
		}, {
			"isDefault": false,
			"name": "feature_6",
			"value": 231
		}, {
			"isDefault": false,
			"name": "feature_7",
			"value": 2
		}, {
			"isDefault": false,
			"name": "feature_3",
			"value": 55
		}, {
			"isDefault": false,
			"name": "feat2",
			"value": "bb"
		}],
		"tracks": [{
			"index": 1,
			"label": "",
			"name": "start_1"
		}, {
			"index": 2,
			"label": "分流实验",
			"name": "abtest_1"
		}, {
			"index": 3,
			"label": "内部规则集2",
			"name": "ruleset_2"
		}],
		"hit_rules": [{
			"id": "55",
			"label": "规则4",
			"name": "rule_4"
		}],
		"node_results": [{
			"IsBlock": false,
			"Kind": "start",
			"Score": 0,
			"Value": null,
			"id": 0,
			"label": "",
			"name": "start_1",
			"tag": ""
		}, {
			"IsBlock": false,
			"Kind": "abtest",
			"Score": 0,
			"Value": 34.20045077555402,
			"id": 11,
			"label": "分流实验",
			"name": "abtest_1",
			"tag": "tag_ab"
		}, {
			"IsBlock": true,
			"Kind": "ruleset",
			"Score": 100,
			"Value": "reject",
			"id": 333,
			"label": "内部规则集2",
			"name": "ruleset_2",
			"tag": "internal"
		}],
		"start_time": "2022-06-12 11:48:59",
		"end_time": "2022-06-12 11:48:59",
		"run_time": 1
	}
}
```
- features 所有特征值，包括执行过程中产生的衍生特征和赋值特征
- hit_rules 命中的规则列表
- tracks 流执行轨迹 
- node_results 各节点执行情况和产生值

## 系统解读 
### 代码结构
```
├── api  http接口逻辑
├── configs  配置文件
├── docs 文档
├── core 决策引擎解析核心目录
├── service 执行逻辑
├── cmd 启动文件
├── global 全局配置
├── demo 测试yaml文件
├── internal
│  ├── dto 数据传输对象
│  ├── errcode 错误异常定义
│  ├── feature 特征
│  └── operator 操作算子
├── test 测试用例
```

### DSL 语法结构
[dsl详解](https://github.com/skyhackvip/risk_engine/tree/master/docs/dsl.md)


### 决策引擎架构图
![决策引擎架构图](https://i.loli.net/2021/01/21/bOR1tyVPnCZNGoi.png)

### 风控系列文章解读
[智能风控决策引擎系统架构设计总纲](https://mp.weixin.qq.com/s?__biz=MzIyMzMxNjYwNw==&mid=2247484064&idx=1&sn=fecd2c7379208e84e7e3cd4eb1abfb6c&chksm=e8215db0df56d4a623bd6be2a706c0220952f0e045b0d6d9646616ee3aae742c574335fa228a&token=221471496&lang=zh_CN#rd)


[智能风控决策引擎系统可落地实现方案（一）规则引擎实现](https://mp.weixin.qq.com/s?__biz=MzIyMzMxNjYwNw==&mid=2247483738&idx=1&sn=111609f176f11de8357c51a820b089b5&chksm=e8215e4adf56d75c2e6e8b81b89c1faabab667f493ce809cb749994cc9cd776342fd17d4172e&token=227666410&lang=zh_CN#rd)

[智能风控决策引擎系统可落地实现方案（二）决策流实现](https://mp.weixin.qq.com/s?__biz=MzIyMzMxNjYwNw==&mid=2247483770&idx=1&sn=3166a6617ddb6b628261b8b7ff84cfac&chksm=e8215e6adf56d77cb76de41b63e63759221932f030e315acebbc4025939b2e02b354a9072ecc&scene=178#rd)

[智能风控决策引擎系统可落地实现方案（三）模型引擎实现](https://mp.weixin.qq.com/s?__biz=MzIyMzMxNjYwNw==&mid=2247483789&idx=1&sn=ddb5f31edfd3174d4551fecc3f120f42&chksm=e8215e9ddf56d78b520f7ab5c8db7e978b3078a1e2511d424ff272ac6c509fd4c13d893dfc09&token=1795265687&lang=zh_CN#rd)

[智能风控决策引擎系统可落地实现方案（四）风控决策实现](https://mp.weixin.qq.com/s?__biz=MzIyMzMxNjYwNw==&mid=2247483825&idx=1&sn=3ebf7c8ad42f870e48db56ca6bb99ade&chksm=e8215ea1df56d7b7d9b1c653c61ef011d72d46d090845d91deba39f635d03ce1282eaa433485&token=1795265687&lang=zh_CN#rd)

[智能风控决策引擎系统可落地实现方案（五）评分卡实现](https://mp.weixin.qq.com/s?__biz=MzIyMzMxNjYwNw==&mid=2247483860&idx=1&sn=45bfbf4e436001dc060d5d4718688e9b&chksm=e8215ec4df56d7d2396c6024b49fc67eb25ee5754da9ddd40365f72abd5c1535a45218ea79b1&token=1239858205&lang=zh_CN#rd)

[智能风控决策引擎系统可落地实现方案（六）风控监控大盘实现](https://mp.weixin.qq.com/s?__biz=MzIyMzMxNjYwNw==&mid=2247483882&idx=1&sn=cb1142ea342b03f2f4ada44383e4bcbe&chksm=e8215efadf56d7ecae2159b7f742678d6036e6df046513ccce0efb052029d13b4c7b67ae1bc6&token=290046129&lang=zh_CN#rd)


[金融智能风控系统演进开发实践](https://mp.weixin.qq.com/s?__biz=MzIyMzMxNjYwNw==&mid=2247484207&idx=1&sn=9ef3c9a1b9f6ca0ad6fca1072925b15d&chksm=e8215c3fdf56d529b23975054a36b3186303400fedd90daa2298dd23c09779895204bc58655d&token=2012091003&lang=zh_CN#rd)

[金融风控领域的 DDD 与中台思考](https://mp.weixin.qq.com/s?__biz=MzIyMzMxNjYwNw==&mid=2247484233&idx=1&sn=59f68324e1b35c3ea2bc642edc21b004&chksm=e8215c59df56d54f9846cb218069451dc247dab2b815a0cdcc044886cb738e1372e2d25ba864&scene=178&cur_album_id=1519884739007053825#rd)

[金融风控系统的演进与升级:从第一代到第四代](https://mp.weixin.qq.com/s?__biz=MzIyMzMxNjYwNw==&mid=2247484409&idx=1&sn=5b646fc06bdf7256f4ff341610878bbd&chksm=e8215ce9df56d5ff9e45b00ca2cbbe8bdc7cc46e3e0c759f5de44118312301677dac7f4807ea&token=2012091003&lang=zh_CN#rd)


扫码关注微信公众号 ***技术岁月*** 支持：

![技术岁月](https://i.loli.net/2021/01/21/orQm9BUkEqKAR6x.jpg)
