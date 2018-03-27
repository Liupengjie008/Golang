kibana介绍：
	Kibana是一个开源的分析与可视化平台，设计出来用于和Elasticsearch一起使用的。你可以用kibana搜索、查看、交互存放在Elasticsearch索引里的数据，使用各种不同的图表、表格、地图等kibana能够很轻易地展示高级数据分析与可视化。

kibana安装：
	1、下载kibana，下载地：https://www.elastic.co/downloads 
	2、修改config/kibana.yml配置 
		elasticsearch.url: “http://192.168.1.103:9200” 		（http://本地ip:9200）
	3、启动kibana： ./bin/kibana     （Windows启动： ./bin/kibana.bat ）


浏览器测试：
	http://localhost:5601