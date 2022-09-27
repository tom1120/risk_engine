# 决策引擎系统
[![License](https://img.shields.io/:license-apache%202-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![GoDoc](https://godoc.org/github.com/skyhackvip/risk_engine?status.png)](https://godoc.org/github.com/skyhackvip/risk_engine)
[![Go Report Card](https://goreportcard.com/badge/github.com/skyhackvip/risk_engine)](https://goreportcard.com/report/github.com/skyhackvip/risk_engine)

## 开源声明
本项目用于学习和参考，采用 Apache License, Version 2.0 开源，转载使用请说明出处。代码不定期迭代更新，欢迎 Star & Watch，如需交流请添加公众号“技术岁月”。

## 关于版本
最新更新稳定版本为 master 分支。

## 决策引擎系统介绍
决策引擎系统，是构建于规则引擎和流程引擎基础上，满足复杂业务决策的一套系统，可用于反欺诈、信用评估、风险决策、推荐系统、精准营销、内容审核、数据清洗等领域场景。

风控是决策引擎的一个重要应用场景，在金融、保险、电商、支付、游戏、社交等领域皆有应用，凡是和钱相关的业务都离不开风控决策引擎系统的支持保障。目前流行的“智能风控”，即是以决策引擎为核心驱动，以机器学习 / AI 为大脑，在大数据基础上构建的通用风控能力。而决策引擎作为风控的核心系统，承担着复杂规则和多样化决策需求，这款开源决策引擎即满足了这样的诉求。


### 功能列表
- 规则、决策
- 规则集
- 决策树
- 决策表
- 决策矩阵(交叉决策表）
- 评分卡
- 决策流
- 冠军挑战者 
- 条件分流
- 支持特征类型：int、string、bool、date、array、map
- 支持运算类型：>、>=、<、<=、==、!=、between、before、after、in、not in、contain、not contain、like、key exist、value exist
- 支持自定义函数(udf)：内建函数sum、avg、min、max
- 支持并发执行和串行执行
- 支持决策流执行短路（触发阻断规则或策略）

## 快速开始
- 环境准备
go version go1.13 +

- 代码下载
git clone [git@github.com:skyhackvip/risk_engine.git](git@github.com:skyhackvip/risk_engine.git)

访问下载：
[https://github.com/skyhackvip/risk_engine/](https://github.com/skyhackvip/risk_engine/)


- Make 编译执行（推荐）
```shell
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

- Docker 镜像制作启动
```shell
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


## 支持 HTTP 接口执行
- 获取决策流列表
- 执行决策流

[接口文档详情](docs/api.md)

## 实践案例（测试样例）
[测试样例](docs/demo.md)

## DSL 语法结构
[Dsl 语法详解](docs/dsl.md)

## 代码结构
```
├── api  http接口逻辑
├── configs  配置文件
├── docs 文档
├── core 决策引擎解析核心目录
├── service 执行逻辑
├── cmd 启动文件
├── global 全局配置
├── demo 测试样例 dsl 文件
├── internal
│  ├── dto 数据传输对象
│  ├── errcode 错误异常定义
│  ├── util 工具包
│  ├── udf 内建自定义函数
│  └── operator 表达式操作算子
├── test 测试用例
```

### 决策引擎架构图
![决策引擎架构图](https://i.loli.net/2021/01/21/bOR1tyVPnCZNGoi.png)

### 风控系列文章解读
- [一款开源又好用的决策引擎](https://mp.weixin.qq.com/s?__biz=MzIyMzMxNjYwNw==&mid=2247484522&idx=1&sn=9cbd6bb463a5d5dc49da8da72e8db77a&chksm=e8215b7adf56d26cd332475a3d7fa8dc7536e526565749c282737fc8bcf7d2e57bc8161ef9da&token=2083215119&lang=zh_CN#rd)

- [智能风控决策引擎系统架构设计总纲](https://mp.weixin.qq.com/s?__biz=MzIyMzMxNjYwNw==&mid=2247484064&idx=1&sn=fecd2c7379208e84e7e3cd4eb1abfb6c&chksm=e8215db0df56d4a623bd6be2a706c0220952f0e045b0d6d9646616ee3aae742c574335fa228a&token=221471496&lang=zh_CN#rd)


- [智能风控决策引擎系统可落地实现方案（一）规则引擎实现](https://mp.weixin.qq.com/s?__biz=MzIyMzMxNjYwNw==&mid=2247483738&idx=1&sn=111609f176f11de8357c51a820b089b5&chksm=e8215e4adf56d75c2e6e8b81b89c1faabab667f493ce809cb749994cc9cd776342fd17d4172e&token=227666410&lang=zh_CN#rd)

- [智能风控决策引擎系统可落地实现方案（二）决策流实现](https://mp.weixin.qq.com/s?__biz=MzIyMzMxNjYwNw==&mid=2247483770&idx=1&sn=3166a6617ddb6b628261b8b7ff84cfac&chksm=e8215e6adf56d77cb76de41b63e63759221932f030e315acebbc4025939b2e02b354a9072ecc&scene=178#rd)

- [智能风控决策引擎系统可落地实现方案（三）模型引擎实现](https://mp.weixin.qq.com/s?__biz=MzIyMzMxNjYwNw==&mid=2247483789&idx=1&sn=ddb5f31edfd3174d4551fecc3f120f42&chksm=e8215e9ddf56d78b520f7ab5c8db7e978b3078a1e2511d424ff272ac6c509fd4c13d893dfc09&token=1795265687&lang=zh_CN#rd)

- [智能风控决策引擎系统可落地实现方案（四）风控决策实现](https://mp.weixin.qq.com/s?__biz=MzIyMzMxNjYwNw==&mid=2247483825&idx=1&sn=3ebf7c8ad42f870e48db56ca6bb99ade&chksm=e8215ea1df56d7b7d9b1c653c61ef011d72d46d090845d91deba39f635d03ce1282eaa433485&token=1795265687&lang=zh_CN#rd)

- [智能风控决策引擎系统可落地实现方案（五）评分卡实现](https://mp.weixin.qq.com/s?__biz=MzIyMzMxNjYwNw==&mid=2247483860&idx=1&sn=45bfbf4e436001dc060d5d4718688e9b&chksm=e8215ec4df56d7d2396c6024b49fc67eb25ee5754da9ddd40365f72abd5c1535a45218ea79b1&token=1239858205&lang=zh_CN#rd)

- [智能风控决策引擎系统可落地实现方案（六）风控监控大盘实现](https://mp.weixin.qq.com/s?__biz=MzIyMzMxNjYwNw==&mid=2247483882&idx=1&sn=cb1142ea342b03f2f4ada44383e4bcbe&chksm=e8215efadf56d7ecae2159b7f742678d6036e6df046513ccce0efb052029d13b4c7b67ae1bc6&token=290046129&lang=zh_CN#rd)


- [金融智能风控系统演进开发实践](https://mp.weixin.qq.com/s?__biz=MzIyMzMxNjYwNw==&mid=2247484207&idx=1&sn=9ef3c9a1b9f6ca0ad6fca1072925b15d&chksm=e8215c3fdf56d529b23975054a36b3186303400fedd90daa2298dd23c09779895204bc58655d&token=2012091003&lang=zh_CN#rd)

- [金融风控领域的 DDD 与中台思考](https://mp.weixin.qq.com/s?__biz=MzIyMzMxNjYwNw==&mid=2247484233&idx=1&sn=59f68324e1b35c3ea2bc642edc21b004&chksm=e8215c59df56d54f9846cb218069451dc247dab2b815a0cdcc044886cb738e1372e2d25ba864&scene=178&cur_album_id=1519884739007053825#rd)

- [金融风控系统的演进与升级:从第一代到第四代](https://mp.weixin.qq.com/s?__biz=MzIyMzMxNjYwNw==&mid=2247484409&idx=1&sn=5b646fc06bdf7256f4ff341610878bbd&chksm=e8215ce9df56d5ff9e45b00ca2cbbe8bdc7cc46e3e0c759f5de44118312301677dac7f4807ea&token=2012091003&lang=zh_CN#rd)


扫码关注微信公众号 ***技术岁月*** 支持：

![技术岁月](https://i.loli.net/2021/01/21/orQm9BUkEqKAR6x.jpg)
