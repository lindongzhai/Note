

# ElasticSearch

## 一、ElaseticSearch 简介

> 官网：https://www.elastic.co/cn/elasticsearch

> 官方文档：https://www.elastic.co/guide/en/elasticsearch/reference/current/index.html

> Elasticsearch是Elastic Stack核心的分布式搜索和分析引擎。Logstash和Beats用于收集、聚合和丰富数据，并将其存储在Elasticsearch中。Kibana使您能够交互式地探索、可视化和共享数据洞察，并管理和监控整个堆栈。Elasticsearch是索引、搜索和分析数据的核心。
>
> Elasticsearch提供了几乎实时的搜索和分析功能，适用于各种类型的数据。无论您有结构化或非结构化文本、数值数据还是地理空间数据，Elasticsearch都可以高效地存储和索引，并支持快速搜索。您可以超越简单的数据检索，聚合信息以发现数据中的趋势和模式。随着数据和查询量的增长，Elasticsearch的分布式特性使得部署可以无缝地扩展。尽管不是每个问题都是搜索问题，但Elasticsearch提供了处理各种用例中的数据的速度和灵活性：
>
> - 在应用程序或网站中添加搜索框
>- 存储和分析日志、指标和安全事件数据
> - 使用机器学习实时自动对数据进行建模
> - 将Elasticsearch用作矢量数据库，创建、存储和搜索矢量嵌入
> - 使用Elasticsearch作为存储引擎自动化业务工作流程
> - 使用Elasticsearch作为地理信息系统（GIS）管理、集成和分析空间信息
> - 使用Elasticsearch作为生物信息学研究工具存储和处理基因数据
> 
> 我们对人们以搜索的方式处理数据的新颖方法感到不断惊讶。但无论您的用例与这些示例相似，还是使用Elasticsearch解决新问题，您在Elasticsearch中处理数据、文档和索引的方式都是相同的。

### 1.1、特点

1. 分布式架构：Elasticsearch 8采用分布式架构，可以将数据分布在多个节点上，实现高可用性和高性能的搜索和分析。
2. 实时搜索和分析：Elasticsearch 8能够实时地处理和索引大规模数据，提供快速的搜索和分析功能。
3. 灵活的数据模型：Elasticsearch 8采用文档型数据模型，可以存储和索引各种类型的结构化和非结构化数据。
4. 强大的查询语言：Elasticsearch 8支持丰富的查询语言，包括全文搜索、过滤、聚合等功能，可以灵活地进行数据分析和挖掘。
5. 可扩展性和可靠性：Elasticsearch 8能够轻松地水平扩展，支持集群部署，提供高可用性和容错性。
6. 开源和活跃的社区支持：Elasticsearch 8是开源软件，拥有庞大的开发者社区，提供了丰富的插件和扩展功能。



### 1.2、优点

1. 高性能：Elasticsearch 8通过优化的搜索算法和索引结构，实现了快速的搜索和分析性能。
2. 可扩展性：Elasticsearch 8能够轻松地水平扩展，支持集群部署，能够处理大规模数据。
3. 灵活的数据模型：Elasticsearch 8支持各种类型的数据存储和索引，适用于多种场景和应用。
4. 易用性：Elasticsearch 8提供了简单易用的RESTful API，可以通过HTTP请求进行交互，降低了使用和集成的复杂性。
5. 强大的查询和分析功能：Elasticsearch 8提供了丰富的查询语言和聚合功能，可以进行复杂的数据分析和挖掘。



### 1.3、缺点

1. 学习曲线较陡峭：对于新手来说，Elasticsearch 8的学习曲线可能会比较陡峭，需要花费一定的时间来熟悉其概念和技术。
2. 需要大量资源：Elasticsearch 8在处理大规模数据时可能需要大量的计算和存储资源，这可能会增加企业的运营成本。
3. 安全性问题：虽然Elasticsearch 8提供了强大的安全功能，但如果配置不当或者受到攻击，仍有可能导致数据泄露或者其他安全性问题。



### 1.4、主要用途

1. 搜索：Elasticsearch可以用于全文搜索、结构化搜索、地理位置搜索等多种类型的搜索，并且支持实时搜索和高并发访问。
2. 数据分析：Elasticsearch提供了丰富的分析工具，可以帮助企业提取有价值的信息和洞察。例如，它可以通过聚合功能进行数据统计和可视化，帮助企业更好地理解和优化其业务流程。
3. 日志管理：Elasticsearch可以用于收集、存储、索引和分析日志数据，帮助企业和开发人员监控系统运行状态和排查问题。
4. 机器学习：Elasticsearch提供了自动机器学习功能，可以自动识别异常情况并生成预测模型，帮助企业更好地理解和优化其业务流程。
5. 应用程序开发：Elasticsearch提供了一个易于使用的API，可以帮助开发者快速集成到自己的应用程序中，以实现数据存储、检索和分析等功能。



### 1.5、使用 ElasticSearch 的公司

1. Uber：Uber使用Elasticsearch来收集、存储和分析其全球业务的数据，并使用它的实时搜索功能为用户提供快速的乘车服务。
2. Netflix：Netflix使用Elasticsearch来索引和搜索其大量的电影和电视节目元数据，并使用它的可视化工具来分析用户行为和偏好。
3. Airbnb：Airbnb使用Elasticsearch来处理其全球民宿预订和评价数据，并使用它的机器学习功能来优化推荐算法。
4. LinkedIn：LinkedIn使用Elasticsearch来支持其职业社交网络的搜索和分析功能，并使用它的聚合功能来提供用户的个性化建议。
5. Twitter：Twitter使用Elasticsearch来索引和搜索其海量的社交媒体数据，并使用它的实时搜索功能来提供用户的即时消息推送。



## 二、ElastciSearch 基础概念

> Elasticsearch是面向文档(document oriented)的，这意味着它可以存储整个对象或文档(document)。然而它不仅仅是存储，还会索引(index)每个文档的内容使之可以被搜索。在Elasticsearch中，你可以对文档（而非成行成列的数据）进行索引、搜索、排序、过滤。

### 2.1、index（索引）

> 简易理解：MySQL 中的数据库

> 一个索引就是一个拥有几分相似特征的文档的集合。索引由一个名字来标识（必须全部是小写字母的），并且当我们要对对应于这个索引中的文档进行索引、搜索、更新和删除的时候，都要使用到这个名字。

### 2.2、type（类型）

> <font color="Orange">注意：在 ElasticSearch8 中已经废弃，默认为 "_doc"</font>

> 简易理解：MySQL 中的表

> 在一个索引中，你可以定义一种或多种类型。一个类型是你的索引的一个逻辑上的分类/分区，其语义完全由你来定。通常，会为具有一组共同字段的文档定义一个类型。

### 2.3、document（文档）

> <font color="Orange">注意：尽管一个文档，物理上存在于一个索引之中，文档必须被索引/赋予一个索引的type。</font>

> 简易理解：MySQL 中的行记录

> 一个文档是一个可被索引的基础信息单元。文档以JSON（Javascript Object Notation）格式来表示。在一个index/type里面，你可以存储任意多的文档。

### 2.4、Field（字段）

> 简易理解：MySQL 中的字段

> 对文档数据根据不同属性进行的分类标识 。

### 2.5、mapping（映射）

> 简易理解：MySQL 中的字段定义

> mapping是处理数据的方式和规则方面做一些限制，如某个字段的数据类型、默认值、分析器、是否被索引等等，这些都是映射里面可以设置的，其它就是处理es里面数据的一些使用规则设置也叫做映射，按着最优规则处理数据对性能提高很大，因此才需要建立映射，并且需要思考如何建立映射才能对性能更好。

### 2.6、cluster（集群）

> <font color="Orange">注意：一个节点只能通过指定某个集群的名字来加入这个集群</font>

> 简易理解：Kubernetes 中的集群

> 一个集群就是由一个或多个节点组织在一起，它们共同持有整个的数据，并一起提供索引和搜索功能。一个集群由 一个唯一的名字标识，这个名字默认就是“elasticsearch”。

### 2.7、node（节点）

> 简易理解：Kubernetes 中的节点

> 一个节点是集群中的一个服务器，作为集群的一部分，它存储数据，参与集群的索引和搜索功能。和集群类似，一 个节点也是由一个名字来标识的，默认情况下，这个名字是一个随机的漫威漫画角色的名字，这个名字会在启动的 时候赋予节点。这个名字对于管理工作来说挺重要的，因为在这个管理过程中，你会去确定网络中的哪些服务器对 应于Elasticsearch集群中的哪些节点。
>
> 一个节点可以通过配置集群名称的方式来加入一个指定的集群。默认情况下，每个节点都会被安排加入到一个叫 做“elasticsearch”的集群中，这意味着，如果你在你的网络中启动了若干个节点，并假定它们能够相互发现彼此， 它们将会自动地形成并加入到一个叫做“elasticsearch”的集群中。
> 在一个集群里，只要你想，可以拥有任意多个节点。而且，如果当前你的网络中没有运行任何Elasticsearch节点， 这时启动一个节点，会默认创建并加入一个叫做“elasticsearch”的集群。

### 2.8、shards&replicas（分片和复制）

> 简易理解：会将一个索引的文件进行分布式存储，且创建副本避免数据的丢失，类似 raid

> Elasticsearch允许将索引划分为多个分片，每个分片都是功能完整的独立“索引”，可以放置在集群中的任何节点上。分片的目的主要是为了水平分割内容容量和实现分布式并行操作以提高性能/吞吐量。同时，Elasticsearch还支持创建复制分片作为主分片的备份，用于故障转移和提供高可用性。复制分片也可以帮助扩展搜索量和吞吐量，因为搜索可以在所有复制分片上并行运行。默认情况下，Elasticsearch中的每个索引包含5个主分片和1个复制，如果集群中至少有两个节点，则每个索引会有5个主分片和5个复制分片（总共10个分片）。

### 2.9、正排索引&倒排索引

* 正排索引

  > 正排索引是一种记录文档中每个单词出现在哪些文档中的数据结构。它将每个单词作为主键，对应的值是一个列表，列表中的元素是包含该单词的文档编号。这样就可以通过查询一个单词来快速获取包含这个单词的所有文档编号。
  >
  > | ID   | 开发语言     |
  > | ---- | ------------ |
  > | 1    | python       |
  > | 2    | java         |
  > | 3    | python，java |
  > | 4    | golang       |
  >
  > 

* 倒排索引

  > 倒排索引是一种记录每个文档中哪些单词出现及次数的数据结构。与正排索引相反，倒排索引将每个文档作为主键，对应的值是一个列表，列表中的元素是文档中出现的单词及其频率。这样就可以通过查询一个文档来快速获取该文档中所有单词的信息。
  >
  > | 开发语言 | ID   |
  > | -------- | ---- |
  > | python   | 1，3 |
  > | java     | 2，3 |
  > | golang   | 4    |

### 2.10、文档得分

$$
score=(q,d) = coord(q,d)·queryNorm(q)·∑(tf(t in d)·idf(t)²·t.getBoot()·norm(t,d))
$$

> 公式中将查询作为输入，使用不同的手段来确定每一篇文档的得分，将每一个因素最后通过公式综合起来，返回该文档的最终得分。  

> 查看得分情况：`GET /索引名称/_search?explain=true`

![image-20240102165421156](C:\Users\dz\AppData\Roaming\Typora\typora-user-images\image-20240102165421156.png)

#### 2.10.1、名词讲解

* TF（Term Frequency）【词频】：搜索文本中的各个词条（term）在查询文本中出现了多少次，出现次数越多，就越相关，得分会比较高

* IDF（Inverse Document Frequency）【逆文档频率】：搜索文本中的各个词条（term）在整个索引的所有文档中出现了多少次，出现的次数越多，说明越不重要，也就越不相关，得分就比较低。



#### 2.10.2、TF 计算

```json
{
    "value": 0.45454544,
    "description": "tf, computed as freq / (freq + k1 * (1 - b + b * dl / avgdl)) from:",
    "details": [
        {
            "value": 1,
            "description": "freq, occurrences of term within document",
            "details": []
        },
        {
            "value": 1.2,
            "description": "k1, term saturation parameter",
            "details": []
        },
        {
            "value": 0.75,
            "description": "b, length normalization parameter",
            "details": []
        },
        {
            "value": 1,
            "description": "dl, length of field",
            "details": []
        },
        {
            "value": 1,
            "description": "avgdl, average length of field",
            "details": []
        }
    ]
}
```



$$
freq / (freq + k1 * (1 - b + b * dl / avgdl))
$$

| 参数  | 含义                                             | 默认值 |
| ----- | ------------------------------------------------ | ------ |
| freq  | 文档中出现词条的次数                             | 无     |
| k1    | 术语饱和参数                                     | 1.2    |
| b     | 长度规格化参数（单词长度对于整个文档的影响程度） | 0.75   |
| dl    | 当前文中分解的字段长度                           | 无     |
| avgdl | 查询文档中分解字段数量/查询文档数量              | 无     |

#### 2.10.3、IDF 计算

```json
{
    "value": 1.3862944,
    "description": "idf, computed as log(1 + (N - n + 0.5) / (n + 0.5)) from:",
    "details": [
        {
            "value": 1,
            "description": "n, number of documents containing term",
            "details": []
        },
        {
            "value": 5,
            "description": "N, total number of documents with field",
            "details": []
        }
    ]
}
```

$$
log(1 + (N - n + 0.5) / (n + 0.5))
$$

| 参数 | 含义                                         | 默认值 |
| ---- | -------------------------------------------- | ------ |
| N    | 包含查询字段的文档总数（不一定包含查询词条） | 无     |
| n    | 包含查询词条的文档数                         | 无     |

#### 2.10.4、文档得分 计算

$$
boost * idf * tf
$$

| 参数  | 含义                 | 默认值                      |
| ----- | -------------------- | --------------------------- |
| boost | 词条权重             | 2.2（基础值） * 查询权重(1) |
| idf   | 包含查询词条的文档数 | 无                          |
| tf    | 词频                 | 无                          |



## 三、ElasticSearch 配置文件说明

### 3.1、Elasticsearch配置文件（elasticsearch.yml）

> 官方文档：https://www.elastic.co/guide/en/elasticsearch/reference/current/settings.html

> 路径：$ES_HOME/config/elasticsearch.yml

> 更改方法：`export ES_PATH_CONF=/path/to/my/config`

```yaml
# ======================== Elasticsearch Configuration =========================
#
# NOTE: Elasticsearch comes with reasonable defaults for most settings.
#       Before you set out to tweak and tune the configuration, make sure you
#       understand what are you trying to accomplish and the consequences.
#
# The primary way of configuring a node is via this file. This template lists
# the most important settings you may want to configure for a production cluster.
#
# Please consult the documentation for further information on configuration options:
# https://www.elastic.co/guide/en/elasticsearch/reference/index.html
#
# ---------------------------------- Cluster -----------------------------------
# ---------------------------------- 集群配置 -----------------------------------
# 集群名称，使用（默认注释）
cluster.name: my-application

# ------------------------------------ Node ------------------------------------
# ------------------------------------ 节点 ------------------------------------
# 节点名称（默认注释）
node.name: node-1
# 节点的自定义属性（默认注释）
node.attr.rack: r1
# 是否有资格选举 node
node.master: true
# 

# ----------------------------------- Paths ------------------------------------
# ----------------------------------- 路径集 ------------------------------------
# 存储数据的目录路径（用逗号分隔多个位置）（默认注释）
path.data: /path/to/data
# 日志位置（默认注释）
path.logs: /path/to/logs
# 临时文件位置（默认注释）
path.work: /path/to/work
# 配置文件位置（默认注释）
path.conf: /path/to/conf
# 插件位置
path.plugins: /path/to/plugins

# ----------------------------------- Memory -----------------------------------
# ----------------------------------- 内存 -----------------------------------
# 启动时锁定内存（默认注释）
bootstrap.memory_lock: true

# ---------------------------------- Network -----------------------------------
# ---------------------------------- 网络 -----------------------------------
# 默认情况下，Elasticsearch 只能在本地主机上访问。 此处设置不同的地址以在网络上公开该节点：（默认注释）
# 当填写时，ES假定从开发模式转向生产模式，并将许多系统启动检查从警告升级为异常
network.host: 192.168.0.1
# 默认情况下，Elasticsearch 在第一个空闲端口上侦听 HTTP 流量（默认注释）
# 查找从 9200 开始。在此处设置特定的 HTTP 端口（默认注释）
http.port: 9200

# --------------------------------- Discovery ----------------------------------
# --------------------------------- 发现 ----------------------------------
# 传递初始主机列表以在该节点启动时执行发现：
discovery.seed_hosts: ["host1", "host2"]
# 使用一组初始的符合主节点资格的节点引导集群：
cluster.initial_master_nodes: ["node-1", "node-2"]

# ---------------------------------- Various -----------------------------------
# ---------------------------------- 各种各样 -----------------------------------
# Allow wildcard deletion of indices:
action.destructive_requires_name: false

#----------------------- BEGIN SECURITY AUTO CONFIGURATION -----------------------
#----------------------- 开始安全自动配置 -----------------------
#
# The following settings, TLS certificates, and keys have been automatically      
# generated to configure Elasticsearch security features on 29-12-2023 07:38:00
#
# --------------------------------------------------------------------------------

# 是否启用安全特性
xpack.security.enabled: true
# 启用X-Pack安全特性中的Enrollment特性。该特性允许节点通过证书自动加入集群。
xpack.security.enrollment.enabled: true

# 为HTTP API客户端连接(如Kibana、Logstash和Agents)启用加密
xpack.security.http.ssl:
  # 表示启用SSL/TLS加密。
  enabled: true
  # 指定SSL证书文件的路径。
  keystore.path: certs/http.p12

# 开启集群节点间的加密和相互认证功能
xpack.security.transport.ssl:
  # 表示启用SSL/TLS加密。
  enabled: true
  # 表示使用证书进行节点身份验证。
  verification_mode: certificate
  # 指定SSL证书文件的路径。
  keystore.path: certs/transport.p12
  # 指定SSL信任证书库文件的路径。
  truststore.path: certs/transport.p12

# 仅使用当前节点创建一个新集群
# 以后，其他节点仍然可以加入集群
cluster.initial_master_nodes: ["DZ-PC"]

# 监听地址，0.0.0.0 表示监听来自任何地方的 HTTP API 连接
http.host: 0.0.0.0

# 允许其他节点从任何地方加入集群
# 连接经过加密并相互验证
transport.host: 0.0.0.0

#----------------------- END SECURITY AUTO CONFIGURATION -------------------------
#----------------------- 结束安全自动配置 -------------------------
```

### 3.2、JVM配置文件（jvm.options）

> 官方文档：
>
> https://www.oracle.com/java/technologies/javase/vmoptions-jsp.html
>
> https://www.elastic.co/guide/en/elasticsearch/reference/8.11/advanced-configuration.html#set-jvm-options
>
> https://www.elastic.co/guide/en/elasticsearch/reference/8.11/important-settings.html#heap-size-settings

> 配置文件路径：$ES_HOME/config/jvm.options

```properties
################################################################
##
## JVM configuration
##
################################################################
##
## WARNING: DO NOT EDIT THIS FILE. If you want to override the
## JVM options in this file, or set any additional options, you
## should create one or more files in the jvm.options.d
## directory containing your adjustments.
##
## See https://www.elastic.co/guide/en/elasticsearch/reference/8.11/jvm-options.html
## for more information.
##
################################################################



################################################################
## IMPORTANT: JVM heap size
################################################################
##
## The heap size is automatically configured by Elasticsearch
## based on the available memory in your system and the roles
## each node is configured to fulfill. If specifying heap is
## required, it should be done through a file in jvm.options.d,
## which should be named with .options suffix, and the min and
## max should be set to the same value. For example, to set the
## heap to 4 GB, create a new file in the jvm.options.d
## directory containing these lines:
##
## -Xms4g
## -Xmx4g
##
## See https://www.elastic.co/guide/en/elasticsearch/reference/8.11/heap-size.html
## for more information
##
################################################################


################################################################
## Expert settings
################################################################
##
## All settings below here are considered expert settings. Do
## not adjust them unless you understand what you are doing. Do
## not edit them in this file; instead, create a new file in the
## jvm.options.d directory containing your adjustments.
##
################################################################

# 启用 G1 垃圾回收器。
-XX:+UseG1GC

## JVM临时目录
-Djava.io.tmpdir=${ES_TMPDIR}

# 利用加速矢量硬件指令;移除这个可能
# 导致矢量性能不太理想
20-:--add-modules=jdk.incubator.vector

# 当从Java堆分配失败时生成堆转储;堆转储将在JVM的工作目录中创建，除非指定了替代路径
-XX:+HeapDumpOnOutOfMemoryError

# 在内存不足错误发生堆转储后立即退出
-XX:+ExitOnOutOfMemoryError

# 指定堆转储的替代路径;确保目录存在并且有足够的空间
-XX:HeapDumpPath=data

# 为JVM致命错误日志指定一个替代路径
-XX:ErrorFile=logs/hs_err_pid%p.log

## GC日志记录
-Xlog:gc*,gc+age=trace,safepoint:file=logs/gc.log:utctime,level,pid,tags:filecount=32,filesize=64m
```

### 3.3、log4j2.properties

> 官方文档：https://logging.apache.org/log4j/2.x/manual/configuration.html

> 配置文件路径：$ES_HOME/config/log4j2.properties"

```properties
status = error

# 指定了一个名为“console”的日志输出器类型
appender.console.type = Console
# 日志输出器的名称为“console”
appender.console.name = console
# 日志输出器使用PatternLayout来格式化日志内容。PatternLayout允许用户通过指定模式字符串来自定义日志输出的格式。
appender.console.layout.type = PatternLayout
# PatternLayout的具体格式化模式
appender.console.layout.pattern = [%d{ISO8601}][%-5p][%-25c{1.}] [%node_name]%marker %m%consoleException%n

######## Server JSON ############################
# 输出日志的方式是滚动文件，即当文件达到一定条件（如大小或时间间隔）时会自动创建新的日志文件并保存旧日志。
appender.rolling.type = RollingFile
# 日志输出器的名字为"rolling"
appender.rolling.name = rolling
# 指定了日志输出的文件名。
# ${sys:es.logs.base_path} 表示从系统环境变量 es.logs.base_path 中获取基础日志目录路径。
# ${sys:file.separator} 是一个系统依赖的路径分隔符（例如Windows上的"\"或Unix上的"/"）。
# ${sys:es.logs.cluster_name}_server.json 表示加上集群名称前缀以及"_server.json"后缀，最终形成完整日志文件名，格式为 <集群名称>_server.json，并且日志内容将以JSON格式编码。
appender.rolling.fileName = ${sys:es.logs.base_path}${sys:file.separator}${sys:es.logs.cluster_name}_server.json

# 日志输出的布局格式为ECSJsonLayout。ECS代表Elastic Common Schema，是一个标准化的日志数据结构，用于方便地将日志数据导入Elasticsearch进行统一管理和分析。
appender.rolling.layout.type = ECSJsonLayout
# 设置了日志布局的dataset属性为"elasticsearch.server"，这个值会在生成的JSON日志中作为顶级键值对"dataset": "elasticsearch.server"出现，有助于区分不同的日志类型，并便于在Elasticsearch中进行索引和过滤。
appender.rolling.layout.dataset = elasticsearch.server

# 定义滚动产生的新日志文件的命名规则。
# ${sys:es.logs.base_path}${sys:file.separator}：使用系统环境变量es.logs.base_path确定的基础日志路径，然后用系统路径分隔符连接；
# ${sys:es.logs.cluster_name}：日志所属的Elasticsearch集群名称；
# %d{yyyy-MM-dd}：日期，格式为年-月-日；
# %i：表示文件编号，当达到某个滚动条件时自动增加；
# .json.gz：表示日志文件扩展名为.json.gz，即压缩后的JSON格式文件。
appender.rolling.filePattern = ${sys:es.logs.base_path}${sys:file.separator}${sys:es.logs.cluster_name}-%d{yyyy-MM-dd}-%i.json.gz

# 声明该appender使用多个策略来决定何时滚动日志。
appender.rolling.policies.type = Policies
# 定义了一个基于时间的触发策略，表示当达到指定的时间间隔时触发滚动。
appender.rolling.policies.time.type = TimeBasedTriggeringPolicy
# 设置时间间隔为1，通常会和后面的modulate配合，实现每天或者每小时滚动一次。
appender.rolling.policies.time.interval = 1
# 表示根据时间间隔进行调制，确保文件名中的时间戳按小时或天平滑递增（例如每天或每小时生成一个新的文件）。
appender.rolling.policies.time.modulate = true
# 定义了一个基于文件大小的触发策略，表示当单个日志文件大小达到某个阈值时触发滚动。
appender.rolling.policies.size.type = SizeBasedTriggeringPolicy
# 设置单个日志文件的最大大小为128MB。
appender.rolling.policies.size.size = 128MB

# 使用默认的滚动策略，来管理文件编号和滚动过程。
appender.rolling.strategy.type = DefaultRolloverStrategy
# 表示没有文件编号的最大限制，即可以无限滚动。
appender.rolling.strategy.fileIndex = nomax
# 定义了一个删除动作，当不再需要某些旧日志文件时执行删除操作。
appender.rolling.strategy.action.type = Delete
# 指定要执行删除操作的日志文件的基础路径。
appender.rolling.strategy.action.basepath = ${sys:es.logs.base_path}
# 定义了一个基于文件名的条件，用于确定哪些文件应该被删除。
appender.rolling.strategy.action.condition.type = IfFileName
# 匹配所有以当前集群名称开头的日志文件。
appender.rolling.strategy.action.condition.glob = ${sys:es.logs.cluster_name}-*
# 定义了一个嵌套条件，基于累计文件大小判断是否需要删除。
appender.rolling.strategy.action.condition.nested_condition.type = IfAccumulatedFileSize
# 当所有匹配文件的总大小超过2GB时，执行删除动作，清除超出空间限制的旧日志文件。
appender.rolling.strategy.action.condition.nested_condition.exceeds = 2GB
################################################
######## Server -  old style pattern ###########

# 定义了这个日志输出器的类型为RollingFile，即滚动文件日志Appender。
appender.rolling_old.type = RollingFile
# 给这个日志输出器起名为"rolling_old"，可用于在其他配置段中引用。
appender.rolling_old.name = rolling_old
# 定义了日志文件的位置和名称。其中 ${sys:es.logs.base_path} 指定的是系统环境变量中定义的日志基本路径；${sys:file.separator} 根据操作系统使用相应的路径分隔符；${sys:es.logs.cluster_name}.log 表示日志文件名是以集群名称后跟".log"的形式。
appender.rolling_old.fileName = ${sys:es.logs.base_path}${sys:file.separator}${sys:es.logs.cluster_name}.log
# 说明该输出器的日志格式采用PatternLayout，可以自定义日志输出格式。
appender.rolling_old.layout.type = PatternLayout
# 定义了具体的日志输出格式
# %d{ISO8601}：表示日志记录的时间，采用ISO8601格式
appender.rolling_old.layout.pattern = [%d{ISO8601}][%-5p][%-25c{1.}] [%node_name]%marker %m%n
# 
appender.rolling_old.filePattern = ${sys:es.logs.base_path}${sys:file.separator}${sys:es.logs.cluster_name}-%d{yyyy-MM-dd}-%i.log.gz
appender.rolling_old.policies.type = Policies
appender.rolling_old.policies.time.type = TimeBasedTriggeringPolicy
appender.rolling_old.policies.time.interval = 1
appender.rolling_old.policies.time.modulate = true
appender.rolling_old.policies.size.type = SizeBasedTriggeringPolicy
appender.rolling_old.policies.size.size = 128MB
appender.rolling_old.strategy.type = DefaultRolloverStrategy
appender.rolling_old.strategy.fileIndex = nomax
appender.rolling_old.strategy.action.type = Delete
appender.rolling_old.strategy.action.basepath = ${sys:es.logs.base_path}
appender.rolling_old.strategy.action.condition.type = IfFileName
appender.rolling_old.strategy.action.condition.glob = ${sys:es.logs.cluster_name}-*
appender.rolling_old.strategy.action.condition.nested_condition.type = IfAccumulatedFileSize
appender.rolling_old.strategy.action.condition.nested_condition.exceeds = 2GB
################################################

rootLogger.level = info
rootLogger.appenderRef.console.ref = console
rootLogger.appenderRef.rolling.ref = rolling
rootLogger.appenderRef.rolling_old.ref = rolling_old

######## Deprecation JSON #######################
appender.deprecation_rolling.type = RollingFile
appender.deprecation_rolling.name = deprecation_rolling
appender.deprecation_rolling.fileName = ${sys:es.logs.base_path}${sys:file.separator}${sys:es.logs.cluster_name}_deprecation.json
appender.deprecation_rolling.layout.type = ECSJsonLayout
# Intentionally follows a different pattern to above
appender.deprecation_rolling.layout.dataset = deprecation.elasticsearch
appender.deprecation_rolling.filter.rate_limit.type = RateLimitingFilter

appender.deprecation_rolling.filePattern = ${sys:es.logs.base_path}${sys:file.separator}${sys:es.logs.cluster_name}_deprecation-%i.json.gz
appender.deprecation_rolling.policies.type = Policies
appender.deprecation_rolling.policies.size.type = SizeBasedTriggeringPolicy
appender.deprecation_rolling.policies.size.size = 1GB
appender.deprecation_rolling.strategy.type = DefaultRolloverStrategy
appender.deprecation_rolling.strategy.max = 4

appender.header_warning.type = HeaderWarningAppender
appender.header_warning.name = header_warning
#################################################

logger.deprecation.name = org.elasticsearch.deprecation
logger.deprecation.level = WARN
logger.deprecation.appenderRef.deprecation_rolling.ref = deprecation_rolling
logger.deprecation.appenderRef.header_warning.ref = header_warning
logger.deprecation.additivity = false

######## Search slowlog JSON ####################
appender.index_search_slowlog_rolling.type = RollingFile
appender.index_search_slowlog_rolling.name = index_search_slowlog_rolling
appender.index_search_slowlog_rolling.fileName = ${sys:es.logs.base_path}${sys:file.separator}${sys:es.logs\
  .cluster_name}_index_search_slowlog.json
appender.index_search_slowlog_rolling.layout.type = ECSJsonLayout
appender.index_search_slowlog_rolling.layout.dataset = elasticsearch.index_search_slowlog

appender.index_search_slowlog_rolling.filePattern = ${sys:es.logs.base_path}${sys:file.separator}${sys:es.logs\
  .cluster_name}_index_search_slowlog-%i.json.gz
appender.index_search_slowlog_rolling.policies.type = Policies
appender.index_search_slowlog_rolling.policies.size.type = SizeBasedTriggeringPolicy
appender.index_search_slowlog_rolling.policies.size.size = 1GB
appender.index_search_slowlog_rolling.strategy.type = DefaultRolloverStrategy
appender.index_search_slowlog_rolling.strategy.max = 4
#################################################

#################################################
logger.index_search_slowlog_rolling.name = index.search.slowlog
logger.index_search_slowlog_rolling.level = trace
logger.index_search_slowlog_rolling.appenderRef.index_search_slowlog_rolling.ref = index_search_slowlog_rolling
logger.index_search_slowlog_rolling.additivity = false

######## Indexing slowlog JSON ##################
appender.index_indexing_slowlog_rolling.type = RollingFile
appender.index_indexing_slowlog_rolling.name = index_indexing_slowlog_rolling
appender.index_indexing_slowlog_rolling.fileName = ${sys:es.logs.base_path}${sys:file.separator}${sys:es.logs.cluster_name}\
  _index_indexing_slowlog.json
appender.index_indexing_slowlog_rolling.layout.type = ECSJsonLayout
appender.index_indexing_slowlog_rolling.layout.dataset = elasticsearch.index_indexing_slowlog


appender.index_indexing_slowlog_rolling.filePattern = ${sys:es.logs.base_path}${sys:file.separator}${sys:es.logs.cluster_name}\
  _index_indexing_slowlog-%i.json.gz
appender.index_indexing_slowlog_rolling.policies.type = Policies
appender.index_indexing_slowlog_rolling.policies.size.type = SizeBasedTriggeringPolicy
appender.index_indexing_slowlog_rolling.policies.size.size = 1GB
appender.index_indexing_slowlog_rolling.strategy.type = DefaultRolloverStrategy
appender.index_indexing_slowlog_rolling.strategy.max = 4
#################################################


logger.index_indexing_slowlog.name = index.indexing.slowlog.index
logger.index_indexing_slowlog.level = trace
logger.index_indexing_slowlog.appenderRef.index_indexing_slowlog_rolling.ref = index_indexing_slowlog_rolling
logger.index_indexing_slowlog.additivity = false


logger.org_apache_pdfbox.name = org.apache.pdfbox
logger.org_apache_pdfbox.level = off

logger.org_apache_poi.name = org.apache.poi
logger.org_apache_poi.level = off

logger.org_apache_fontbox.name = org.apache.fontbox
logger.org_apache_fontbox.level = off

logger.org_apache_xmlbeans.name = org.apache.xmlbeans
logger.org_apache_xmlbeans.level = off


logger.com_amazonaws.name = com.amazonaws
logger.com_amazonaws.level = warn

logger.com_amazonaws_jmx_SdkMBeanRegistrySupport.name = com.amazonaws.jmx.SdkMBeanRegistrySupport
logger.com_amazonaws_jmx_SdkMBeanRegistrySupport.level = error

logger.com_amazonaws_metrics_AwsSdkMetrics.name = com.amazonaws.metrics.AwsSdkMetrics
logger.com_amazonaws_metrics_AwsSdkMetrics.level = error

logger.com_amazonaws_auth_profile_internal_BasicProfileConfigFileLoader.name = com.amazonaws.auth.profile.internal.BasicProfileConfigFileLoader
logger.com_amazonaws_auth_profile_internal_BasicProfileConfigFileLoader.level = error

logger.com_amazonaws_services_s3_internal_UseArnRegionResolver.name = com.amazonaws.services.s3.internal.UseArnRegionResolver
logger.com_amazonaws_services_s3_internal_UseArnRegionResolver.level = error


appender.audit_rolling.type = RollingFile
appender.audit_rolling.name = audit_rolling
appender.audit_rolling.fileName = ${sys:es.logs.base_path}${sys:file.separator}${sys:es.logs.cluster_name}_audit.json
appender.audit_rolling.layout.type = PatternLayout
appender.audit_rolling.layout.pattern = {\
                "type":"audit", \
                "timestamp":"%d{yyyy-MM-dd'T'HH:mm:ss,SSSZ}"\
                %varsNotEmpty{, "cluster.name":"%enc{%map{cluster.name}}{JSON}"}\
                %varsNotEmpty{, "cluster.uuid":"%enc{%map{cluster.uuid}}{JSON}"}\
                %varsNotEmpty{, "node.name":"%enc{%map{node.name}}{JSON}"}\
                %varsNotEmpty{, "node.id":"%enc{%map{node.id}}{JSON}"}\
                %varsNotEmpty{, "host.name":"%enc{%map{host.name}}{JSON}"}\
                %varsNotEmpty{, "host.ip":"%enc{%map{host.ip}}{JSON}"}\
                %varsNotEmpty{, "event.type":"%enc{%map{event.type}}{JSON}"}\
                %varsNotEmpty{, "event.action":"%enc{%map{event.action}}{JSON}"}\
                %varsNotEmpty{, "authentication.type":"%enc{%map{authentication.type}}{JSON}"}\
                %varsNotEmpty{, "user.name":"%enc{%map{user.name}}{JSON}"}\
                %varsNotEmpty{, "user.run_by.name":"%enc{%map{user.run_by.name}}{JSON}"}\
                %varsNotEmpty{, "user.run_as.name":"%enc{%map{user.run_as.name}}{JSON}"}\
                %varsNotEmpty{, "user.realm":"%enc{%map{user.realm}}{JSON}"}\
                %varsNotEmpty{, "user.realm_domain":"%enc{%map{user.realm_domain}}{JSON}"}\
                %varsNotEmpty{, "user.run_by.realm":"%enc{%map{user.run_by.realm}}{JSON}"}\
                %varsNotEmpty{, "user.run_by.realm_domain":"%enc{%map{user.run_by.realm_domain}}{JSON}"}\
                %varsNotEmpty{, "user.run_as.realm":"%enc{%map{user.run_as.realm}}{JSON}"}\
                %varsNotEmpty{, "user.run_as.realm_domain":"%enc{%map{user.run_as.realm_domain}}{JSON}"}\
                %varsNotEmpty{, "user.roles":%map{user.roles}}\
                %varsNotEmpty{, "apikey.id":"%enc{%map{apikey.id}}{JSON}"}\
                %varsNotEmpty{, "apikey.name":"%enc{%map{apikey.name}}{JSON}"}\
                %varsNotEmpty{, "authentication.token.name":"%enc{%map{authentication.token.name}}{JSON}"}\
                %varsNotEmpty{, "authentication.token.type":"%enc{%map{authentication.token.type}}{JSON}"}\
                %varsNotEmpty{, "cross_cluster_access":%map{cross_cluster_access}}\
                %varsNotEmpty{, "origin.type":"%enc{%map{origin.type}}{JSON}"}\
                %varsNotEmpty{, "origin.address":"%enc{%map{origin.address}}{JSON}"}\
                %varsNotEmpty{, "realm":"%enc{%map{realm}}{JSON}"}\
                %varsNotEmpty{, "realm_domain":"%enc{%map{realm_domain}}{JSON}"}\
                %varsNotEmpty{, "url.path":"%enc{%map{url.path}}{JSON}"}\
                %varsNotEmpty{, "url.query":"%enc{%map{url.query}}{JSON}"}\
                %varsNotEmpty{, "request.method":"%enc{%map{request.method}}{JSON}"}\
                %varsNotEmpty{, "request.body":"%enc{%map{request.body}}{JSON}"}\
                %varsNotEmpty{, "request.id":"%enc{%map{request.id}}{JSON}"}\
                %varsNotEmpty{, "action":"%enc{%map{action}}{JSON}"}\
                %varsNotEmpty{, "request.name":"%enc{%map{request.name}}{JSON}"}\
                %varsNotEmpty{, "indices":%map{indices}}\
                %varsNotEmpty{, "opaque_id":"%enc{%map{opaque_id}}{JSON}"}\
                %varsNotEmpty{, "trace.id":"%enc{%map{trace.id}}{JSON}"}\
                %varsNotEmpty{, "x_forwarded_for":"%enc{%map{x_forwarded_for}}{JSON}"}\
                %varsNotEmpty{, "transport.profile":"%enc{%map{transport.profile}}{JSON}"}\
                %varsNotEmpty{, "rule":"%enc{%map{rule}}{JSON}"}\
                %varsNotEmpty{, "put":%map{put}}\
                %varsNotEmpty{, "delete":%map{delete}}\
                %varsNotEmpty{, "change":%map{change}}\
                %varsNotEmpty{, "create":%map{create}}\
                %varsNotEmpty{, "invalidate":%map{invalidate}}\
                }%n
# "node.name" node name from the `elasticsearch.yml` settings
# "node.id" node id which should not change between cluster restarts
# "host.name" unresolved hostname of the local node
# "host.ip" the local bound ip (i.e. the ip listening for connections)
# "origin.type" a received REST request is translated into one or more transport requests. This indicates which processing layer generated the event "rest" or "transport" (internal)
# "event.action" the name of the audited event, eg. "authentication_failed", "access_granted", "run_as_granted", etc.
# "authentication.type" one of "realm", "api_key", "token", "anonymous" or "internal"
# "user.name" the subject name as authenticated by a realm
# "user.run_by.name" the original authenticated subject name that is impersonating another one.
# "user.run_as.name" if this "event.action" is of a run_as type, this is the subject name to be impersonated as.
# "user.realm" the name of the realm that authenticated "user.name"
# "user.realm_domain" if "user.realm" is under a domain, this is the name of the domain
# "user.run_by.realm" the realm name of the impersonating subject ("user.run_by.name")
# "user.run_by.realm_domain" if "user.run_by.realm" is under a domain, this is the name of the domain
# "user.run_as.realm" if this "event.action" is of a run_as type, this is the realm name the impersonated user is looked up from
# "user.run_as.realm_domain" if "user.run_as.realm" is under a domain, this is the name of the domain
# "user.roles" the roles array of the user; these are the roles that are granting privileges
# "apikey.id" this field is present if and only if the "authentication.type" is "api_key"
# "apikey.name" this field is present if and only if the "authentication.type" is "api_key"
# "authentication.token.name" this field is present if and only if the authenticating credential is a service account token
# "authentication.token.type" this field is present if and only if the authenticating credential is a service account token
# "cross_cluster_access" this field is present if and only if the associated authentication occurred cross cluster
# "event.type" informs about what internal system generated the event; possible values are "rest", "transport", "ip_filter" and "security_config_change"
# "origin.address" the remote address and port of the first network hop, i.e. a REST proxy or another cluster node
# "realm" name of a realm that has generated an "authentication_failed" or an "authentication_successful"; the subject is not yet authenticated
# "realm_domain" if "realm" is under a domain, this is the name of the domain
# "url.path" the URI component between the port and the query string; it is percent (URL) encoded
# "url.query" the URI component after the path and before the fragment; it is percent (URL) encoded
# "request.method" the method of the HTTP request, i.e. one of GET, POST, PUT, DELETE, OPTIONS, HEAD, PATCH, TRACE, CONNECT
# "request.body" the content of the request body entity, JSON escaped
# "request.id" a synthetic identifier for the incoming request, this is unique per incoming request, and consistent across all audit events generated by that request
# "action" an action is the most granular operation that is authorized and this identifies it in a namespaced way (internal)
# "request.name" if the event is in connection to a transport message this is the name of the request class, similar to how rest requests are identified by the url path (internal)
# "indices" the array of indices that the "action" is acting upon
# "opaque_id" opaque value conveyed by the "X-Opaque-Id" request header
# "trace_id" an identifier conveyed by the part of "traceparent" request header
# "x_forwarded_for" the addresses from the "X-Forwarded-For" request header, as a verbatim string value (not an array)
# "transport.profile" name of the transport profile in case this is a "connection_granted" or "connection_denied" event
# "rule" name of the applied rule if the "origin.type" is "ip_filter"
# the "put", "delete", "change", "create", "invalidate" fields are only present
# when the "event.type" is "security_config_change" and contain the security config change (as an object) taking effect

appender.audit_rolling.filePattern = ${sys:es.logs.base_path}${sys:file.separator}${sys:es.logs.cluster_name}_audit-%d{yyyy-MM-dd}-%i.json.gz
appender.audit_rolling.policies.type = Policies
appender.audit_rolling.policies.time.type = TimeBasedTriggeringPolicy
appender.audit_rolling.policies.time.interval = 1
appender.audit_rolling.policies.time.modulate = true
appender.audit_rolling.policies.size.type = SizeBasedTriggeringPolicy
appender.audit_rolling.policies.size.size = 1GB
appender.audit_rolling.strategy.type = DefaultRolloverStrategy
appender.audit_rolling.strategy.fileIndex = nomax

logger.xpack_security_audit_logfile.name = org.elasticsearch.xpack.security.audit.logfile.LoggingAuditTrail
logger.xpack_security_audit_logfile.level = info
logger.xpack_security_audit_logfile.appenderRef.audit_rolling.ref = audit_rolling
logger.xpack_security_audit_logfile.additivity = false

logger.xmlsig.name = org.apache.xml.security.signature.XMLSignature
logger.xmlsig.level = error
logger.samlxml_decrypt.name = org.opensaml.xmlsec.encryption.support.Decrypter
logger.samlxml_decrypt.level = fatal
logger.saml2_decrypt.name = org.opensaml.saml.saml2.encryption.Decrypter
logger.saml2_decrypt.level = fatal
```

## 四、Kibana 使用







## 五、索引操作

### 5.1、新建索引

```http
PUT /indexName
```

```sh
curl -XPUT "https://172.19.0.2:9200/lindongzhai" -H "kbn-xsrf: reporting"
```



* 创建成功响应

    ```json
    {
      "acknowledged": true,
      "shards_acknowledged": true,
      "index": "lindongzhai"
    }
    ```

    ![image-20231227230355349](EleasticSearch.assets/image-20231227230355349.png)

* 索引已存在错误响应

    ```json
    {
      "error": {
        "root_cause": [
          {
            "type": "resource_already_exists_exception",
            "reason": "index [lindongzhai/A_iSA0wsQnS-gx3u_MsLMg] already exists",
            "index_uuid": "A_iSA0wsQnS-gx3u_MsLMg",
            "index": "lindongzhai"
          }
        ],
        "type": "resource_already_exists_exception",
        "reason": "index [lindongzhai/A_iSA0wsQnS-gx3u_MsLMg] already exists",
        "index_uuid": "A_iSA0wsQnS-gx3u_MsLMg",
        "index": "lindongzhai"
      },
      "status": 400
    }
    ```

    ![image-20231227230458356](EleasticSearch.assets/image-20231227230458356.png)

### 5.2、查询索引

```http
GET /indexName
```

* 索引存在

    ```JSON
    {
      "lindongzhai": {
        "aliases": {},
        "mappings": {},
        "settings": {
          "index": {
            "routing": {
              "allocation": {
                "include": {
                  "_tier_preference": "data_content"
                }
              }
            },
            "number_of_shards": "1",
            "provided_name": "lindongzhai",
            "creation_date": "1703688956328",
            "number_of_replicas": "1",
            "uuid": "A_iSA0wsQnS-gx3u_MsLMg",
            "version": {
              "created": "8500003"
            }
          }
        }
      }
    }
    ```

    ![image-20231227230811476](EleasticSearch.assets/image-20231227230811476.png)

* 索引不存在

    ```json
    {
      "error": {
        "root_cause": [
          {
            "type": "index_not_found_exception",
            "reason": "no such index [lindongzhai1]",
            "resource.type": "index_or_alias",
            "resource.id": "lindongzhai1",
            "index_uuid": "_na_",
            "index": "lindongzhai1"
          }
        ],
        "type": "index_not_found_exception",
        "reason": "no such index [lindongzhai1]",
        "resource.type": "index_or_alias",
        "resource.id": "lindongzhai1",
        "index_uuid": "_na_",
        "index": "lindongzhai1"
      },
      "status": 404
    }
    ```

    ![image-20231227230924382](EleasticSearch.assets/image-20231227230924382.png)

### 5.3、所有索引

```http
GET _cat/indices
```

```json
green  open .internal.alerts-observability.logs.alerts-default-000001      mP05Naz3TpyXv1fVV1wVdw 1 0 0 0 249b 249b 249b
green  open .internal.alerts-observability.uptime.alerts-default-000001    OMsgGuc3T2ujx8c7XdcB_g 1 0 0 0 249b 249b 249b
green  open .internal.alerts-ml.anomaly-detection.alerts-default-000001    N-Hf870fTEmlL9kD_VSNlQ 1 0 0 0 249b 249b 249b
green  open .internal.alerts-observability.slo.alerts-default-000001       Cg3Gv3dsQ6eWzbgPFc72Ww 1 0 0 0 249b 249b 249b
green  open .internal.alerts-observability.apm.alerts-default-000001       RQoc1NeSSSmXdpQiItxqLQ 1 0 0 0 249b 249b 249b
yellow open lindongzhai                                                    A_iSA0wsQnS-gx3u_MsLMg 1 1 0 0 249b 249b 249b
green  open .internal.alerts-observability.metrics.alerts-default-000001   BNLIQDMKRAGa8Vdsj597bQ 1 0 0 0 249b 249b 249b
green  open .kibana-observability-ai-assistant-conversations-000001        zEiFJFoqSwGJhare-nzmOA 1 0 0 0 249b 249b 249b
green  open .internal.alerts-observability.threshold.alerts-default-000001 1lMAyrvWQsu17Npw67nJkg 1 0 0 0 249b 249b 249b
green  open .kibana-observability-ai-assistant-kb-000001                   C95vr9FRR2eKfdzAcdE2bQ 1 0 0 0 249b 249b 249b
green  open .internal.alerts-security.alerts-default-000001                Sr0HV5lyQfy34q0_nk6tlw 1 0 0 0 249b 249b 249b
green  open .internal.alerts-stack.alerts-default-000001                   r4OFtkliQe2Tbh50gChqKw 1 0 0 0 249b 249b 249b
```

![image-20231227231846789](EleasticSearch.assets/image-20231227231846789.png)

> 查询结果说明
>
> | 内容                   | 含义           | 具体描述                                                     |
> | ---------------------- | -------------- | ------------------------------------------------------------ |
> | green                  | health         | 当前服务器健康状态，green（集群完整），yellow（单点正常、集群不完整），red（单点不正常） |
> | open                   | status         | 索引打开、关闭状态                                           |
> | lindongzhai            | index          | 索引名称                                                     |
> | A_iSA0wsQnS-gx3u_MsLMg | uuid           | 索引唯一编号                                                 |
> | 1                      | pri            | 主分片数量                                                   |
> | 1                      | rep            | 副本数量                                                     |
> | 0                      | docs.count     | 可用文档数（已有文档数）                                     |
> | 0                      | docs.deleted   | 已删除文档数（逻辑删除/软删除）                              |
> | 249b                   | store.size     | 主分片和副分片整体占空间大小                                 |
> | 249b                   | pri.store.size | 主分片占空间大小                                             |
>
> 

### 5.4、删除索引

```HTTP
DELETE /indexName
```

* 删除成功

    ```JSON
    {
      "acknowledged": true
    }
    ```

    ![image-20231227232550677](EleasticSearch.assets/image-20231227232550677.png)

* 删除不存在索引

    ```JSON
    {
      "error": {
        "root_cause": [
          {
            "type": "index_not_found_exception",
            "reason": "no such index [lindongzhai1]",
            "resource.type": "index_or_alias",
            "resource.id": "lindongzhai1",
            "index_uuid": "_na_",
            "index": "lindongzhai1"
          }
        ],
        "type": "index_not_found_exception",
        "reason": "no such index [lindongzhai1]",
        "resource.type": "index_or_alias",
        "resource.id": "lindongzhai1",
        "index_uuid": "_na_",
        "index": "lindongzhai1"
      },
      "status": 404
    }
    ```

    ![image-20231227232651458](EleasticSearch.assets/image-20231227232651458.png)

## 六、文档操作

### 6.1、新增文档

#### 6.1.1、方式一：自动生成ID

```http
POST /索引名/_doc
{
	"key": "value"
}
```

![image-20231229160045214](C:\Users\dz\AppData\Roaming\Typora\typora-user-images\image-20231229160045214.png)

```JSON
{
  "_index": "lindongzhai",       // 当前索引名称
  "_id": "1taUtIwBwSyoWLhSLjGX", // 文档ID
  "_version": 1,                 // 文档版本，增删改都会使版本产生变化
  "result": "created",           // 操作结果，create（创建），update（更新），delete（删除）
  "_shards": {                   // 分片信息
    "total": 2,                  // 总分片数
    "successful": 1,             // 成功分片数 
    "failed": 0                  // 失败分片数
  },
  "_seq_no": 0,                  // 文档序列号，用于分布式环境种保证顺序性
  "_primary_term": 1             // 主分片任期，用于确保在同一时间只有一个主分片处于活动状态，并保证变更操作的一致性。
}
```

#### 6.1.2、方式二：指定ID

```http
POST /索引名/_doc/文档ID
{
	"key": "value"
}
```

![image-20231229161533300](C:\Users\dz\AppData\Roaming\Typora\typora-user-images\image-20231229161533300.png)

```json
{
  "_index": "lindongzhai",       // 当前索引名称
  "_id": "001", 				 // 文档ID
  "_version": 1,                 // 文档版本，增删改都会使版本产生变化
  "result": "created",           // 操作结果，create（创建），update（更新），delete（删除）
  "_shards": {                   // 分片信息
    "total": 2,                  // 总分片数
    "successful": 1,             // 成功分片数 
    "failed": 0                  // 失败分片数
  },
  "_seq_no": 0,                  // 文档序列号，用于分布式环境种保证顺序性
  "_primary_term": 1             // 主分片任期，用于确保在同一时间只有一个主分片处于活动状态，并保证变更操作的一致性。
}
```

#### 6.1.3、方式三：使用 PUT 新增文档

```http
PUT /索引名称/_doc/文档ID
{
	"key": "value"
}
```

![image-20231229163823898](C:\Users\dz\AppData\Roaming\Typora\typora-user-images\image-20231229163823898.png)

```json
{
  "_index": "lindongzhai",       // 当前索引名称
  "_id": "002", 				 // 文档ID
  "_version": 1,                 // 文档版本，增删改都会使版本产生变化
  "result": "created",           // 操作结果，create（创建），update（更新），delete（删除）
  "_shards": {                   // 分片信息
    "total": 2,                  // 总分片数
    "successful": 1,             // 成功分片数 
    "failed": 0                  // 失败分片数
  },
  "_seq_no": 0,                  // 文档序列号，用于分布式环境种保证顺序性
  "_primary_term": 1             // 主分片任期，用于确保在同一时间只有一个主分片处于活动状态，并保证变更操作的一致性。
}
```



### 6.2、修改文档

#### 6.2.1、方式一：修改所有字段

>  <font color="Orange">注意：这种方式修改，如果不传所有字段，将会把未传的字段清空</font>

```http
POST /索引名称/_doc/文档ID
{
	"key": "value"
}
```

![image-20231229164246167](C:\Users\dz\AppData\Roaming\Typora\typora-user-images\image-20231229164246167.png)

```json
{
  "_index": "lindongzhai",       // 当前索引名称
  "_id": "002", 				 // 文档ID
  "_version": 1,                 // 文档版本，增删改都会使版本产生变化
  "result": "update",            // 操作结果，create（创建），update（更新），delete（删除）
  "_shards": {                   // 分片信息
    "total": 2,                  // 总分片数
    "successful": 1,             // 成功分片数 
    "failed": 0                  // 失败分片数
  },
  "_seq_no": 0,                  // 文档序列号，用于分布式环境种保证顺序性
  "_primary_term": 1             // 主分片任期，用于确保在同一时间只有一个主分片处于活动状态，并保证变更操作的一致性。
}
```

#### 6.2.2、方式二：修改指定字段



### 6.3、删除文档

> <font color="Orange">注意：删除一个文档不会立即从磁盘上移除，它只是被标记成已删除（逻辑删除）</font>

```http
DELETE /索引名称/_doc/文档ID
```

![image-20231229165612343](C:\Users\dz\AppData\Roaming\Typora\typora-user-images\image-20231229165612343.png)

```json
{
  "_index": "lindongzhai",  // 当前索引名称
  "_id": "002",             // 文档ID
  "_version": 2,            // 文档版本，增删改都会使版本产生变化
  "result": "deleted",      // 操作结果，create（创建），update（更新），delete（删除）
  "_shards": {              // 分片信息
    "total": 2,             // 总分片数
    "successful": 1,        // 成功分片数
    "failed": 0             // 失败分片数
  },
  "_seq_no": 4,             // 文档序列号，用于分布式环境种保证顺序性
  "_primary_term": 1        // 主分片任期，用于确保在同一时间只有一个主分片处于活动状态，并保证变更操作的一致性。
}
```

### 6.4、查询文档

#### 6.4.1、方式一：查询单个文档

```http
GET /索引名称/_doc/文档ID[?pretty=true]
```

![image-20231229165029235](C:\Users\dz\AppData\Roaming\Typora\typora-user-images\image-20231229165029235.png)

```json
{
  "_index": "lindongzhai",  // 索引名称
  "_id": "001",             // 文档ID
  "_version": 2,            // 文档版本，增删改都会使版本产生变化
  "_seq_no": 3,             // 文档序列号，用于分布式环境种保证顺序性
  "_primary_term": 1,       // 主分片任期，用于确保在同一时间只有一个主分片处于活动状态，并保证变更操作的一致性。
  "found": true,            // 布尔类型字段，表示请求所查找的文档是否存在。
  "_source": {              // 文档内容
    "name": "lindongzhai1",
    "age": 30
  }
}
```

#### 6.4.2、方式二：查询所有文档

```http
GET /索引名称/_search
```

![image-20231229172554390](C:\Users\dz\AppData\Roaming\Typora\typora-user-images\image-20231229172554390.png)

```json
{
  "took": 1,                            // 请求执行时间
  "timed_out": false,                   // 是否超时
  "_shards": {                          //  表示分片的相关信息
    "total": 1,                         // 总分片数
    "successful": 1,                    // 成功分片数
    "skipped": 0,                       // 跳过分片数
    "failed": 0                         // 失败分片数
  },
  "hits": {                             // 搜索结果信息
    "total": {                          // 搜索结果总数的统计信息。
      "value": 2,                       // 匹配的文档数
      "relation": "eq"                  // 与总数量的匹配关系
    },
    "max_score": 1,                     // 最高得分
    "hits": [                           // 命中的文档数组
      {
        "_index": "lindongzhai",        // 索引名称
        "_id": "1taUtIwBwSyoWLhSLjGX",  // 文档ID
        "_score": 1,                    // 文档得分
        "_source": {                    // 文档内容
          "name": "lindongzhai",
          "age": 24
        }
      },
      {
        "_index": "lindongzhai",         // 索引名称
        "_id": "001",                    // 文档ID
        "_score": 1,                     // 文档得分
        "_source": {                     // 文档内容
          "name": "lindongzhai1",
          "age": 30
        }
      }
    ]
  }
}
```



## 七、数据搜索

### 7.1、基础搜索

#### 7.1.1、完全匹配查询

```http
GET /apidoc/_search
{
  "query": {
    "match": {
      "name": "zhangsan"
    }
  }
}
```

![image-20240102143022381](C:\Users\dz\AppData\Roaming\Typora\typora-user-images\image-20240102143022381.png)

#### 7.1.2、只查询指定字段

```http
GET /apidoc/_search
{
  "_source": ["name", "age"],
  "query": {
    "match": {
      "name": "zhangsan"
    }
  }
}
```

![image-20240102172006063](C:\Users\dz\AppData\Roaming\Typora\typora-user-images\image-20240102172006063.png)

### 7.2、聚合搜索



## 八、索引模板

## 九、中文分词

### 9.1、原生分词

### 9.2、IK分词器



## 六、IK分词器

### 6.1、安装

<font color="Orange">注意：安装的版本需与ElasticSearch版本一致</font>

* 手动安装

	> 下载地址：https://github.com/medcl/elasticsearch-analysis-ik/releases

	```sh
	mkdir -p ES安装目录/plugins/ik
	unzip elasticsearch-analysis-ik-8.11.3.zip -d ES安装目录/plugins/ik
	docker restart ES-contains
	```

* ES命令安装

  ```sh
  cd ES安装目录/
  ```

  





## 十、SDK 使用

### 7.1、Golang



### 7.2、Java



### 7.3、Python



## 附录

### 1、ElasticSearch 安装

> 官方教程：https://www.elastic.co/guide/en/elasticsearch/reference/current/install-elasticsearch.html

#### 1.1、docker 安装

```SH
# 创建虚拟网络
docker network create elastic
# 拉取镜像
docker pull docker.elastic.co/elasticsearch/elasticsearch:8.11.3
# 安装证书
wget https://artifacts.elastic.co/cosign.pub
cosign verify --key cosign.pub docker.elastic.co/elasticsearch/elasticsearch:8.11.3
# 执行 ElasticSearch
docker run --name es01 --net elastic -p 9200:9200 -it -m 1GB docker.elastic.co/elasticsearch/elasticsearch:8.11.3
# 后台执行 ElasticSearch
docker run --name es01 --net elastic -p 9200:9200 -it -m 1GB -d docker.elastic.co/elasticsearch/elasticsearch:8.11.3
```

#### 1.2、Windows 安装



#### 1.3、Linux 安装



#### 1.4、Mac 安装



### 2、Kibana 安装

> 官方教程：https://www.elastic.co/guide/en/elasticsearch/reference/current/install-elasticsearch.html

#### 2.1、docker 安装

```sh
# 拉取镜像
docker pull docker.elastic.co/kibana/kibana:8.11.3
# 安装证书
wget https://artifacts.elastic.co/cosign.pub
cosign verify --key cosign.pub docker.elastic.co/kibana/kibana:8.11.3
# 执行 Kibana
docker run --name kib01 --net elastic -p 5601:5601 docker.elastic.co/kibana/kibana:8.11.3
# 后台执行 kibana
docker run --name kib01 --net elastic -p 5601:5601 -d docker.elastic.co/kibana/kibana:8.11.3
```



#### 2.2、Windows 安装

<font color="Orange">注意：需要安装java17+</font>

> 官方教程：https://www.elastic.co/guide/en/kibana/8.11/windows.html

> 下载地址：https://artifacts.elastic.co/downloads/kibana/kibana-8.11.3-windows-x86_64.zip



#### 2.3、Linux 安装

> 官方教程：https://www.elastic.co/guide/en/kibana/8.11/rpm.html#rpm-repo

* yum 安装

  ```shell
  sudo yum install kibana
  ```

* dnf 安装

  ```shell
  sudo dnf install kibana
  ```

* zypper 安装

  ```shell
  sudo zypper install kibana
  ```

  





#### 2.4、Mac 安装

