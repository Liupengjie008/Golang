ElasticSearch介绍与使用

https://www.elastic.co/cn/      Elasticsearch官网（下载ElasticSearch）

http://www.jianshu.com/p/05cff717563c
ElasticSearch，一款非常优秀的分布式搜索程序。

ElasticSearch安装
    1) 下载ElasticSearch，下载地：github.com/elastic/elasticsearch
    2) 修改config/elasticsearch.ymal配置：
         network.host: 本地ip
         node.name:node-1
          （如果电脑配置不够，修改config/ jvm.options     -Xms512m     -Xmx512m）
    3) 启动ElasticSearch， ./bin/elasticsearch.bat