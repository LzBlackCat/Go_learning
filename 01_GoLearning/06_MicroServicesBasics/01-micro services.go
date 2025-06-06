package main

// 一、微服务
// 1.1.微服务的概念
// 1.微服务的定义：微服务是一种将应用程序拆分为小型、独立、轻量级服务的架构风格。
// 2.微服务的特点：每个服务都是小型、独立的，能够独立部署、扩展和升级。
// 3.微服务的优势：提高系统的可伸缩性、可维护性和可扩展性。

// 1.2、微服务的组成和架构
// 1.微服务的组成：微服务由多个小型服务器组成，每个服务器提供特定的功能。
// 2.微服务的架构：微服务采用分布式解决方案，类似于操作系统的分布式搭建。
// 3.微服务的例子：例如王者荣耀游戏，每个玩家都是一个微服务，通过服务间的交互完成游戏功能。

// 1.3、微服务的历史和应用
// 1.微服务的提出时间：微服务概念提出较早，但实际应用始于21世纪初期。
// 2.微服务的来源：最初借鉴自金融行业，后来逐渐应用于其他领域。
// 3.微服务的开发团队：开发团队较小，通常称为“二披萨团队”，成员三到四人。
// 		\微服务架构定义的精髓，可以用一句话来描述，那就是“**分而治之，合而用之**”。将复杂的系统进行拆分的方法，就是“分而治之”。
// 		分而治之，可以让复杂的事情变的简单，这很符合我们平时处理问题的方法。
// 		使用轻量级通讯等方式进行整合的设计，就是“合而用之”的方法，合而用之可以让微小的力量变动强大。

//
//
// 二、单体式服务
// 2.1.定义：单体式服务是一个完整的应用程序，包含多个模块。
// 2.2.缺点：复杂性高、技术债务高、维护成本高、持续交付时间长、技术选型成本高、扩展性差。
// 2.3.例子：如华为的BOSS计费系统，随着开发周期增长，复杂性和技术债务不断增加。

// 2.2.单体式服务的缺点（详细）：
// 2.2.1.单体式服务的复杂性
// 		1.复杂性：单体式服务的复杂性随着开发周期增长而增加，解决问题困难。
// 		2.例子：如华为的BOSS计费系统，代码量巨大，修改风险高。
// 		3.后果：导致开发人员不敢轻易修改代码，增加技术债务。

// 2.2.2.单体式服务的技术债务
// 		1.技术债务定义：指在开发过程中为了赶进度或满足需求而积累的技术问题。
// 		2.单体式服务的技术债务：随着开发周期增长，技术债务逐渐增加。
// 		3.例子：如某个开发人员离职前留下的64个bug，后续开发者需要花费更多时间和精力来修复。

// 2.2.3.单体式服务的偶合度
// 		1.偶合度定义：指模块之间的依赖和关联程度。
// 		2.单体式服务的偶合度：模块之间偶合度高，导致维护成本高。
// 		3.问题：出现bug时不容易排查，解决旧bug可能引入新bug。

// 2.2.4.单体式服务的持续交付时间
// 		1.持续交付时间定义：指从代码提交到正式环境部署的时间。
// 		2.单体式服务的持续交付时间：随着团队人员增多和沟通成本增加，交付时间变长。
// 		3.问题：人员利用率不高，开发团队之间沟通复杂。

// 2.2.5.单体式服务的技术选型成本
// 		1.技术选型成本定义：指在选择开发平台和架构时所花费的成本和风险。
// 		2.单体式服务的技术选型成本：选型不当可能导致后期扩展困难。
// 		3.例子：如选择C++作为开发语言，虽然可以应对多种场景，但开发周期长。

// 2.2.6.单体式服务的扩展性
// 		1.扩展性定义：指系统能够应对流量增长和业务扩展的能力。
// 		2.单体式服务的扩展性：垂直扩展（增加单个系统成员的复合）和水平扩展（增加更多系统成员）。
// 		3.问题：单体式服务在扩展时面临模块耦合度高、修改困难的挑战。

//
//
// 三、优缺点及对比
// 3.1.微服务的优点
// 		1.微服务的职责单一：每个服务负责单一的职责，便于独立测试和部署。
// 		2.轻量级通信：进程间的通信量级轻，支持跨语言通信。
// 		3.独立性：每个服务独立开发、测试、编译和部署，互不干扰。
// 		4.迭代开发：引入新功能时，旧服务不需要改动，简化迭代过程。

// 3.2.微服务的缺点
// 		1.运维成本高：每个服务需要单独的开发、测试、编译和部署。
// 		2.分布式复杂度高：微服务借鉴了分布式的思想，增加了系统复杂度。
// 		3.接口成本高：多个服务间的通信需要更多的接口。
// 		4.重复性劳动：每个服务都需要读配置文件，增加了开发工作量。
// 		5.业务分离困难：微服务越细，业务分离越困难，管理调用关系也越复杂。

// 3.3.单体式服务和微服务的对比
// 		1.单体式服务：传统架构，所有功能在一个应用中，容易部署，隔离性差，初期技术选型难度大。
// 		2.微服务：分布式架构，功能拆分成多个服务，经常部署，隔离性好，架构设计难度大，逻辑设计难度大。
// 		3.性能：单体式服务响应快但吞吐量小，微服务并发高但响应慢。
// 		4.运维难度：单体式服务简单，微服务复杂，学习曲线相似，但侧重点不同。
// 		5.技术：单体式服务技术单一，微服务技术多样且易于开发。
// 		6.测试和错误处理：单体式服务测试简单，微服务测试复杂，需要处理集群和负载均衡。
// 		7.扩展性：单体式服务扩展性差，微服务扩展性好。
// 		8.系统管理：单体式服务管理简单，微服务管理复杂，重点在于服务治理和调度。
