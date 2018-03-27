ElasticSearch介绍与使用

https://www.elastic.co/cn/      Elasticsearch官网（下载ElasticSearch）

http://www.jianshu.com/p/05cff717563c
ElasticSearch，一款非常优秀的分布式搜索程序。

ElasticSearch安装
    1) 下载ElasticSearch，下载地：github.com/elastic/elasticsearch
    2) 修改config/elasticsearch.yml配置：
         network.host: 本地ip
         node.name:node-1
          （如果电脑配置不够，修改config/ jvm.options     -Xms512m     -Xmx512m）
    3) 启动ElasticSearch ： ./bin/./elasticsearch (Windows：./bin/elasticsearch.bat)
    

浏览器测试：
    http://192.168.1.103:9200/       （http://本地ip:9200/ ）

    浏览器输出：
    {
        "name" : "node-1",
        "cluster_name" : "elasticsearch",
        "cluster_uuid" : "g_1HOUylRHaKmuaRWzJo3g",
        "version" : {
          "number" : "5.5.2",
          "build_hash" : "b2f0c09",
          "build_date" : "2017-08-14T12:33:14.154Z",
          "build_snapshot" : false,
          "lucene_version" : "6.6.0"
        },
        "tagline" : "You Know, for Search"
      }