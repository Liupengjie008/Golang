MySQL数据库设计 老男孩IT教育，只培养技术精英 
数据类型之整数类型 
MySQL不仅支持标准SQL中的integer和smallint类型，还支持一些自己的扩展的整数类型 
下表中指明了具体的类型，存储消耗的字节数，最小最大取值范围，unsigned代表不允许负数，则正整数的取值范围扩大一倍 

 
数据类型之整数类型 
MySQL可以为整数类型指定宽度，比如INT(11)，这个限制对大多数应用没有意义，因为这不是限制值的合法范围，对于存储和计算来说，INT(1)和INT(20)是相同的，只是对一些MySQL的交互工具规定了显示字符的个数，比如MySQL命令行客户端。  
mysql> create table temp1(id int(1),id2 int(20)); 
Query OK, 0 rows affected (0.01 sec)  
mysql> insert into temp1 values(1000000,1000000); 
Query OK, 1 row affected (0.00 sec)  
mysql> select * from temp1; 
+---------+---------+ 
| id      | id2     | 
+---------+---------+ 
| 1000000 | 1000000 | 
+---------+---------+ 
1 row in set (0.00 sec)  

 
数据类型之整数类型  
mysql> alter table temp1 modify id int(1) zerofill; 
Query OK, 0 rows affected (0.00 sec) 
Records: 0  Duplicates: 0  Warnings: 0  
mysql> alter table temp1 modify id2 int(20) zerofill; 
Query OK, 1 row affected (0.03 sec) 
Records: 1  Duplicates: 0  Warnings: 0  
mysql> select * from temp1; 
+------+----------------------+ 
| id   | id2                  | 
+------+----------------------+ 
|   10 | 00000000000000000010 | 
+------+----------------------+  

 
数据类型之固定浮点类型 
Decimal和numeric数据类型用来存储高精度数据，一般只在对小数进行精确计算时才使用，比如涉及财务数据的时候 
DECIMAL[(M[,D])] [UNSIGNED] 
在MySQL中，numeric和decimal的含义相同 
Decimal的使用方法举例为decimal(5,2) 
其中的5代表为精度，表示了可以使用多少位数字 
其中的2代表小数点后面的小数位数 
此例子的取值范围为-999.99到999.99 
当不需要指定小数时，可以使用decimal(M),decimal(M,0)表示 
当直接使用decimal时，则默认的M为10 
M的最大取值为65，D的最大取值为30，当D为0时可以用来存储比BIGINT更大范围的整数值 
当指定unsigned，表示不允许负数 
MySQL对decimal字段采用每4个字节存储9个数字的方式，例如decimal(18,9)小数点两边各存储9个数字，一共使用9个字节：小数点前的数字用4个字节，小数点后的数字用4个字节，小数点本身占1个字节 

 
数据类型之浮点类型 
浮点类型中包含float和double两种，与decimal相比是不精确类型 
FLOAT[(M,D)] [UNSIGNED]中的M代表可以使用的数字位数，D则代表小数点后的小数位数 
Unsigned(无)代表不允许使用负数 
Float的取值范围为-3.402823466E+38 to -1.175494351E-38, 0, and 1.175494351E-38 to 3.402823466E+38 
DOUBLE[(M,D)] [UNSIGNED]中的M代表可以使用的数字位数，D则代表小数点后的小数位数 
Double的取值范围对比float要大，-1.7976931348623157E+308 to -2.2250738585072014E-308, 0, and 2.2250738585072014E-308 to 1.7976931348623157E+308 
在存储同样范围的值时，通常比decimal使用更少的空间，float使用4个字节存储，double使用8个字节。 

 
数据类型之浮点类型 
mysql> create table temp2(id float(10,2),id2 double(10,2),id3 decimal(10,2));  
mysql> insert into temp2 values(1234567.21, 1234567.21,1234567.21),(9876543.21, 9876543.12, 9876543.12); 
Query OK, 2 rows affected (0.00 sec) 
Records: 2  Duplicates: 0  Warnings: 0  
mysql> select * from temp2; 
+------------+------------+------------+ 
| id         | id2        | id3        | 
+------------+------------+------------+ 
| 1234567.25 | 1234567.21 | 1234567.21 | 
| 9876543.00 | 9876543.12 | 9876543.12 | 
+------------+------------+------------+ 
2 rows in set (0.00 sec) 

 
数据类型之bit类型 
Bit数据类型用来存储bit值 
BIT(M)代表可以存储M个bit，M的取值范围为1到64 
如果手工指定bit值，则可以使用b’value’格式，比如b’111’和b‘10000000’分别代表7和128 
除非特殊情况，否则尽量不要使用这个类型 

 
数据类型之日期时间类型 
日期时间类型包括date,time,datetime,timestamp和year，用来指定不同范围的日期或时间值 
Date类型用来表示仅日期，MySQL默认的日期格式为yyyy-mm-dd，取值范围为1000-01-01到9999-12-31 
Datetime类型用来表示日期和时间，MySQL默认的格式为yyyy-mm-dd hh:mi:ss，取值范围为1000-01-01 00:00:00到9999-12-31 23:59:59 
Timestamp类型也用来表示日期和时间，其取值范围为1970-01-01 00:00:01到2038-01-19 03:14:07 
Datetime和timestamp两个类型都可以保存到微妙级别，即6位毫秒微妙精度，即1000-01-01 00:00:00.000000到9999-12-31 23:59:59.999999和1970-01-01 00:00:01.000000到2038-01-19 03:14:07.999999 
非法的date,datetime,timestamp值将被转换成0值，0000-00-00或者0000-00-00 00:00:00 

 
数据类型之日期时间类型 
Time类型用来仅表示时间，MySQL默认格式为HH:MM:SS，其取值范围为-838:59:59到838:59:59，小时字段可以超过24是因为time类型不光代表小时，也可以代表持续时长中的小时 
Time类型也可以包含6位的毫秒微秒精度，其取值范围为-838:59:59.000000到838:59:59.000000 
Year类型用来仅表示年份，MySQL默认格式为YYYY，其取值范围为1901到2155，和0000 
针对非法的year数据，则直接转化为0000 

 
数据类型之日期时间类型 
Timestamp和datetime日期时间类型可以被自动初始化和更新为当前的日期时间数据，当你默认指定current timestamp为默认值，或者指定此数据列为自动更新时 
指定默认值是指当插入新的数据而该列没有显视指定数值时，则插入当前日期时间值 
指定自动更新是指当行中的其他列被更新时，则此列被自动更新为当前日期时间值 
CREATE TABLE t1 ( 
  ts TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, 
  dt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP 
); CREATE TABLE t1 (   ts TIMESTAMP DEFAULT CURRENT_TIMESTAMP,   dt DATETIME DEFAULT CURRENT_TIMESTAMP ); 

 
数据类型之日期时间类型 
mysql> create table temp3(id int,tstamp datetime,tstamp2 datetime,tstamp3 timestamp,tstamp4 timestamp); 
Query OK, 0 rows affected (0.01 sec)  
mysql> insert into temp3(id) values(1); 
Query OK, 1 row affected (0.77 sec)  
mysql> insert into temp3(id) values(2); 
Query OK, 1 row affected (0.00 sec)  
mysql> select * from temp3; 
+------+--------+---------+---------------------+---------------------+ 
| id   | tstamp | tstamp2 | tstamp3             | tstamp4             | 
+------+--------+---------+---------------------+---------------------+ 
|    1 | NULL   | NULL    | 2017-05-12 23:48:47 | 0000-00-00 00:00:00 | 
|    2 | NULL   | NULL    | 2017-05-12 23:48:51 | 0000-00-00 00:00:00 | 
+------+--------+---------+---------------------+---------------------+ 

 
数据类型之日期时间类型 
mysql> create table temp4(id int,tstamp datetime default current_timestamp,tstamp2 datetime default current_timestamp,tstamp3 timestamp,tstamp4 timestamp default current_timestamp); 
Query OK, 0 rows affected (0.01 sec)  
mysql> insert into temp4(id) values(1); 
Query OK, 1 row affected (0.00 sec)  
mysql> insert into temp4(id) values(2); 
Query OK, 1 row affected (0.00 sec)  
mysql> select * from temp4; 
+------+---------------------+---------------------+---------------------+---------------------+ 
| id   | tstamp              | tstamp2             | tstamp3             | tstamp4             | 
+------+---------------------+---------------------+---------------------+---------------------+ 
|    1 | 2017-05-12 23:51:42 | 2017-05-12 23:51:42 | 2017-05-12 23:51:42 | 2017-05-12 23:51:42 | 
|    2 | 2017-05-12 23:51:44 | 2017-05-12 23:51:44 | 2017-05-12 23:51:44 | 2017-05-12 23:51:44 | 
+------+---------------------+---------------------+---------------------+---------------------+ 

 
数据类型之日期时间类型 
mysql> create table temp4(id int,tstamp datetime default current_timestamp on update current_timestamp,tstamp2 datetime default current_timestamp ,tstamp3 timestamp,tstamp4 timestamp default current_timestamp on update current_timestamp); 
mysql> insert into temp4(id) values(1); 
mysql> select * from temp4; 
+------+---------------------+---------------------+---------------------+---------------------+ 
| id   | tstamp              | tstamp2             | tstamp3             | tstamp4             | 
+------+---------------------+---------------------+---------------------+---------------------+ 
|    1 | 2017-05-12 23:55:11 | 2017-05-12 23:55:11 | 2017-05-12 23:55:11 | 2017-05-12 23:55:11 | 
+------+---------------------+---------------------+---------------------+---------------------+ 
mysql> update temp4 set id=2; 
mysql> select * from temp4; 
+------+---------------------+---------------------+---------------------+---------------------+ 
| id   | tstamp              | tstamp2             | tstamp3             | tstamp4             | 
+------+---------------------+---------------------+---------------------+---------------------+ 
|    2 | 2017-05-12 23:55:27 | 2017-05-12 23:55:11 | 2017-05-12 23:55:27 | 2017-05-12 23:55:27 | 
+------+---------------------+---------------------+---------------------+---------------------+ 
mysql> update temp4 set id=3; 
mysql> select * from temp4; 
+------+---------------------+---------------------+---------------------+---------------------+ 
| id   | tstamp              | tstamp2             | tstamp3             | tstamp4             | 
+------+---------------------+---------------------+---------------------+---------------------+ 
|    3 | 2017-05-12 23:55:46 | 2017-05-12 23:55:11 | 2017-05-12 23:55:46 | 2017-05-12 23:55:46 | 
1 row in set (0.00 sec) 

 
数据类型之日期时间类型 
当在time,timestamp,datetime中指定含有毫秒微秒数值时，则用type_name(fsp)来表达，其中fsp可以取0到6之间的数值 
CREATE TABLE t1 (t TIME(3), dt DATETIME(6)); 
mysql> CREATE TABLE fractest( c1 TIME(2), c2 DATETIME(2), c3 TIMESTAMP(2) ); 
Query OK, 0 rows affected (0.33 sec)  
mysql> INSERT INTO fractest VALUES 
     > ('17:51:04.777', '2014-09-08 17:51:04.777', '2014-09-08 17:51:04.777'); 
     Query OK, 1 row affected (0.03 sec)  
mysql> SELECT * FROM fractest; 
+-------------+------------------------+------------------------+ 
| c1          | c2                     | c3                     | 
+-------------+------------------------+------------------------+ 
| 17:51:04.78 | 2014-09-08 17:51:04.78 | 2014-09-08 17:51:04.78 | 
+-------------+------------------------+------------------------+ 
1 row in set (0.00 sec) 

 
数据类型之字符类型 
字符类型包含char, varchar, binary, varbinary, blob, text, enum和set 
Char和varchar可以通过char(M)和varchar(M)指定可以存储的最大字符数，比如char(30)表示可以存储最长30个字符Char类型的长度一旦指定就固定了，其范围可以是0到255，当被存储时，未达到指定长度的则在值右边填充空格，而获取数据时则会把右侧的空格去掉 
Varchar类型是变长的类型，其范围可以是0到65535，当存储是未达到指定长度则不填充空格 

 
数据类型之字符类型 
Varchar类型用来存储可变长字符串，是最常见的字符串数据类型，它比定长类型更节省空间，因为它仅使用必要的空间。 
另外varchar需要使用1或2个额外字节记录字符串的长度，如果列的最大长度小于等于255字节时，需要1个字节，否则需要2个字节。比如采用Latin1字符集，varchar(10)的列需要11个字节的存储空间，而varchar(1000)列需要1002个字节的存储空间。 
varchar节省了存储空间，所以对性能也有帮助。但由于行是变长的，在update时可能使行变得比原来更长，这就导致需要做额外的工作。如果一行占用的空间增长，并且物理数据页内没有更多空间存储时，MyISAM会将行拆成不同的片段存储，InnoDB需要分列页来讲行放到数据页里。 

 
数据类型之字符类型 
char类型是定长，MySQL总是根据定义的字符串长度分配足够的空间。当查询char值时，MySQL会删除所有的末尾空格 
char适合存储很短的字符串，或者所有值都接近同一个长度。对于经常变更的数据，char也比varchar更好，因为定长的char类型不容易产生碎片。而且对非常短的字符串，char不需要一个额外的字节记录长度 

 
数据类型之字符类型 
Char类型值右边的空格会被自动剔除，而varchar类型则不会 
mysql> CREATE TABLE vc (v VARCHAR(4), c CHAR(4)); 
Query OK, 0 rows affected (0.01 sec)  
mysql> INSERT INTO vc VALUES ('ab  ', 'ab  '); 
Query OK, 1 row affected (0.00 sec)  
mysql> SELECT CONCAT('(', v, ')'), CONCAT('(', c, ')') FROM vc; 
+---------------------+---------------------+ 
| CONCAT('(', v, ')') | CONCAT('(', c, ')') | 
+---------------------+---------------------+ 
| (ab  )              | (ab)                | 
+---------------------+---------------------+ 
1 row in set (0.06 sec) 

 
数据类型之二进制类型 
Binary和varbinary类型和char/varchar类似，只不过是存储二进制字符 
mysql> CREATE TABLE t (c BINARY(3)); 
Query OK, 0 rows affected (0.01 sec)  
mysql> INSERT INTO t SET c = 'a'; 
Query OK, 1 row affected (0.01 sec)  
mysql> SELECT HEX(c), c = 'a', c = 'a\0\0' from t; 
+--------+---------+-------------+ 
| HEX(c) | c = 'a' | c = 'a\0\0' | 
+--------+---------+-------------+ 
| 610000 |       0 |           1 | 
+--------+---------+-------------+ 
1 row in set (0.09 sec) 

 
数据类型之大数据类型 
Blob和text类型被用来存储大量的数据 
Blob是用来存储二进制的大量数据，其有四种类型，tinyblob、blob、mediumblob、longblob四种的区别是能存储的数据长度有所不同 
Text是用来存储字符型的大量数据，其有四种类型， tinytext、text、mediumtext、longtext四种的区别是能存储的数据长度有所不同 
Blob和text的列字段不能含有默认值 

 TINYTEXT 256 bytes   TEXT 65,535 bytes ~64kb MEDIUMTEXT  16,777,215 bytes ~16MB LONGTEXT 4,294,967,295 bytes ~4GB 
数据类型之枚举类型 
Enum枚举类型是字符串类型，其值是从事先指定的一系列值中选出，适用在某列的取值范围已经固定 
主要好处为MySQL在存储此类数据时，直接转化成数字存储而不是字符串，可以节省空间，并且在表的.frm文件中存储“数字-字符串”之间的对应关系 
CREATE TABLE shirts ( 
    name VARCHAR(40), 
    size ENUM('x-small', 'small', 'medium', 'large', 'x-large') 
); 
INSERT INTO shirts (name, size) VALUES ('dress shirt','large'), ('t-shirt','medium'), 
  ('polo shirt','small'); 
  SELECT name, size FROM shirts WHERE size = 'medium'; 
+---------+--------+ 
| name    | size   | 
+---------+--------+ 
| t-shirt | medium | 
+---------+--------+ 
UPDATE shirts SET size = 'small' WHERE size = 'large'; 
COMMIT; 

 
数据类型之枚举类型 
select size+0 from shirts;  
另外枚举类型的排序规则是按照存储顺序进行而不是按照值本身排序的 
select size from shirts order by size; 
mysql> select * from shirts order by size; 
+-------------+--------+ 
| name        | size   | 
+-------------+--------+ 
| polo shirt  | small  | 
| t-shirt     | medium | 
| dress shirt | large  |  
如果想要按照一般的排序规则进行排序，需要使用field()函数显示指定排序规则 
select size from shirts order by field(size,’large’,’medium’,’small’)  
枚举类型字段的取值的增加必须通过alter table命令 
mysql> alter table shirts modify size ENUM('x-small', 'small', 'medium', 'large', 'x-large','xx-large');  01 
数据类型之枚举类型 
Enum枚举类型最多可以有65535个值 
当插入数字到枚举类型字段时，数字会被当做枚举值的第几个值而插入 
numbers ENUM('0','1','2') 
mysql> INSERT INTO t (numbers) VALUES(2),('2'),('3'); 
mysql> SELECT * FROM t; 
+---------+ 
| numbers | 
+---------+ 
| 1       | 
| 2       | 
| 2       | 
+---------+ 

 
数据类型之枚举类型 
枚举类型中的空串和NULL值 
当插入一个非法的值到枚举字段时，则会报错 
如果枚举字段允许NULL，则NULL值为此枚举类型的默认值  
mysql> insert into shirts values('abc','smal'); 
ERROR 1265 (01000): Data truncated for column 'size' at row 1 

 
数据类型之集合类型 
Set集合类型是字符类型，可以含有0个或多个值，其中的每个值都需要是在创建字段时指定的集合中 
比如一个字段被指定为SET(‘one’, ‘two’) not null可以含有以下四种值 
'' 
'one' 
'two' 
'one,two‘ 
Set集合最大可以有255个值 
MySQL在存储set集合时，同样也是存储为数字类型 

 
数据类型之集合类型 
当存储一个数字到set集合字段时，就会按照二进制计算值 
SET('a','b','c','d')     
当该列插入9这个数字时，则转化为二进制的1001，即转化为’a’,’d’值 
Set集合列中各值的顺序无关紧要，且一个值如果出现多次也会被忽略 
mysql> CREATE TABLE myset (col SET('a', 'b', 'c', 'd')); 
mysql> INSERT INTO myset (col) VALUES  
-> ('a,d'), ('d,a'), ('a,d,a'), ('a,d,d'), ('d,a,d'); 
Query OK, 5 rows affected (0.01 sec) 
Records: 5  Duplicates: 0  Warnings: 0 
mysql> SELECT col FROM myset; 
+------+ 
| col  | 
+------+ 
| a,d  | 
| a,d  | 
| a,d  | 
| a,d  | 
| a,d  | 
+------+ 
5 rows in set (0.04 sec) 

 
数据类型之如何选择 
MySQL支持的数据类型很多，选择正确的数据类型对获得高性能至关重要 
更小的通常更好 
尽量使用可以正确存储数据的最小数据类型。更小的数据类型通常更快，因为它们占用更小的磁盘、内存和CPU缓存，并且处理时需要的CPU时间也更少 
比如如果知道某个数字列的存储值在0~200之间，就应该选取tinyint类型 
简单就好 
简单的数据类型操作通常需要更少的CPU周期。例如整型比字符操作代价更低，因为字符集和排序规则使得字符比较比整型比较更复杂 
尽量避免NULL 
通常情况下最好指定列为NOT NULL。因为如果查询中包含可为NULL的列，对MySQL来说更难优化，因为可为NULL的列使得索引、索引统计和值比较都更为复杂。当可为NULL的列被索引时，每个索引记录需要一个额外的字节，所以会使用更多的存储空间 

 
数据类型之设置默认值 
Default默认值用来指定一个列的默认值，但不能指定函数或表达式作为默认值，比如now()和current_date，但唯一的例外是可以指定current_timestamp作为timestamp和datetime列的默认值 
Blob,text列不能指定默认值 
如果一个列没有显视指定default默认值，则依照以下规则 
如果该列允许null值，则默认值为null 

 
自增长类型字段 
整型和浮点型字段可以被指定为自增长类型字段，意味着当插入行数据时这列为NULL时，则按照此列最大值+1的方式插入数据 
获取插入后的自增长列的值，可以用LAST_INSERT_ID()函数获取 
一个表中只能有一个自增长字段，且不能含有默认值 
自增长字段的数值从1开始递增，且不能插入负值 
CREATE TABLE animals ( 
     id MEDIUMINT NOT NULL AUTO_INCREMENT, 
     name CHAR(30) NOT NULL, 
     PRIMARY KEY (id) 
 );  
INSERT INTO animals (name) VALUES 
    ('dog'),('cat'),('penguin'), 
    ('lax'),('whale'),('ostrich');  
    SELECT * FROM animals; 

     
自增长类型字段 
+----+---------+ 
| id | name    | 
+----+---------+ 
|  1 | dog     | 
|  2 | cat     | 
|  3 | penguin | 
|  4 | lax     | 
|  5 | whale   | 
|  6 | ostrich | 
+----+---------+ 

 老男孩IT教育，只培养技术精英 
自增长类型字段 
当你显视的插入一个数值到自增长字段时，则下一个是表中所有值的最大值+1 
设置字段的auto_increment属性，可以有两种方法 
Create table的时候指定 
Alter table的时候指定： mysql> ALTER TABLE tbl AUTO_INCREMENT = 100; 
针对Myisam存储引擎，auto_increment属性可以添加到多列键值的第二列上，则自增列的值计算是根据第一个列分组计算得出 
CREATE TABLE animals ( 
    grp ENUM('fish','mammal','bird') NOT NULL, 
    id MEDIUMINT NOT NULL AUTO_INCREMENT, 
    name CHAR(30) NOT NULL, 
    PRIMARY KEY (grp,id) 
) ENGINE=MyISAM;  
INSERT INTO animals (grp,name) VALUES 
    ('mammal','dog'),('mammal','cat'), 
    ('bird','penguin'),('fish','lax'),('mammal','whale'), 
    ('bird','ostrich');  
    SELECT * FROM animals ORDER BY grp,id; 

     
自增长类型字段 
+--------+----+---------+ 
| grp    | id | name    | 
+--------+----+---------+ 
| fish   |  1 | lax     | 
| mammal |  1 | dog     | 
| mammal |  2 | cat     | 
| mammal |  3 | whale   | 
| bird   |  1 | penguin | 
| bird   |  2 | ostrich | 
+--------+----+---------+ 

 
MySQL存储引擎 
通过执行show engines命令查看MySQL中支持哪些存储引擎 mysql> SHOW ENGINES\G *************************** 1. row ***************************       Engine: PERFORMANCE_SCHEMA      Support: YES      Comment: Performance Schema Transactions: NO           XA: NO   Savepoints: NO *************************** 2. row ***************************       Engine: InnoDB      Support: DEFAULT      Comment: Supports transactions, row-level locking, and foreign keys Transactions: YES           XA: YES   Savepoints: YES ...... *************************** 5. row ***************************       Engine: MyISAM      Support: YES      Comment: MyISAM storage engine Transactions: NO           XA: NO   Savepoints: NO ..... 

 
MySQL存储引擎 
MySQL存储引擎属性对比 

 
MySQL存储引擎 
设置表的存储引擎的方法 
在my.cnf配置文件中设置default-storage-engine参数表示设置默认存储引擎 
在MySQL的连接上设置当前连接的默认存储引擎 
SET default_storage_engine=NDBCLUSTER; 
在创建表的时候通过engine=语句指定该表的存储引擎 
CREATE TABLE t1 (i INT) ENGINE = INNODB; 
-- Simple table definitions can be switched from one to another. 
CREATE TABLE t2 (i INT) ENGINE = CSV; 
CREATE TABLE t3 (i INT) ENGINE = MEMORY; 
在表创建之后通过alter语句修改表的存储引擎 
ALTER TABLE t ENGINE = InnoDB; 

 
MySQL存储引擎之InnoDB 
存储引擎InnoDB是目前MySQL版本默认的存储引擎，也是MySQL推荐使用的存储引擎，是集高可靠性和高性能于一身的存储引擎。 
在MySQL5.7版本中，除非在配置文件中显视指定default storage engine或者创建表时显视使用engine=语句指定其它的存储引擎，否则默认都是InnoDB 
使用InnoDB存储引擎的优势在于 
DML语句支持事务功能，保证ACID特性 
行级锁的使用保证了高并发的属性 
InnoDB对有主键的表会依据主键优化查询性能，也称聚簇索引，将所有数据存储在聚簇索引上以减少对主键查询的IO消耗 
为保证数据的一致性，InnoDB还支持外键属性，确保有外键约束的表之间不会有不一致的数据 
当服务器硬件或者软件故障导致MySQL重启后，InnoDB会自动识别已经在故障之前提交的数据，并回退所有故障时未提交的数据，最大限度的保护数据不会丢失(crash recovery) 

 
MySQL存储引擎之InnoDB 
InnoDB存储引擎的属性 

 
MySQL存储引擎之MyISAM 
MyISAM存储引擎是MySQL老版本的默认存储引擎，由于其表级锁的特性，所以限制了其在读写操作时的性能，常用在只读表上或者读操作占绝大多数的表上，比如一些web应用和数据仓库相关表 
每个MyISAM表都会在磁盘上生成三个文件，表名和文件名相同但后缀不同，.frm文件存储表的结构信息，.MYD文件存储表的数据信息，.MYI文件存储表的索引信息 

 
MySQL存储引擎之Memory 
Memory存储引擎将所有数据存储在内存中以便加快对某些不重要数据的访问速度 
此存储引擎的使用范围已经变小，因为InnoDB已经提供了数据缓存区以便对将经常访问的数据缓存在内存中 
当MySQL重启时，Memory表中的数据会丢失，但表结构还在 
Memory只适用在只读表或者读操作占绝大多数的情况，因为对表的写操作也会导致表锁，大大限制了并发性 
Memory表创建之后，在磁盘文件会生成一个相同表名的文件，后缀为.frm，仅存储表结构而不存储表数据 
mysql> CREATE TABLE test ENGINE=MEMORY 
    ->     SELECT ip,SUM(downloads) AS down 
    ->     FROM log_table GROUP BY ip; 
    mysql> SELECT COUNT(ip),AVG(down) FROM test; 
    mysql> DROP TABLE test; 

     
MySQL存储引擎之CSV 
Csv存储引擎下的表对应了文本文件，其中的数据用逗号隔开，csv表可用来以csv格式导入和导出表 
当创建一个csv表时，磁盘会生成三个以表名为名字的文件，.frm存储表的结构信息，而.CSV文件用来存储以逗号隔开的数据信息，.CSM文件用来存储表的元数据，包括表的状态和有多少行数据信息 
mysql> CREATE TABLE test (i INT NOT NULL, c CHAR(10) NOT NULL) 
    -> ENGINE = CSV; 
    Query OK, 0 rows affected (0.12 sec)  
mysql> INSERT INTO test VALUES(1,'record one'),(2,'record two'); 
Query OK, 2 rows affected (0.00 sec) 
Records: 2  Duplicates: 0  Warnings: 0  
mysql> SELECT * FROM test; 
+------+------------+ 
| i    | c          | 
+------+------------+ 
|    1 | record one | 
|    2 | record two | 
+------+------------+ 
2 rows in set (0.00 sec) 

 
MySQL存储引擎之ARCHIVE 
Archive存储引擎表用来存储大量未加索引的历史归档数据 
archive表会在磁盘创建两个文件，.frm文件用来存储表结构信息，.ARZ文件用来存储历史归档数据 
Archive表支持insert, replace和select语句，但不支持delete和update语句 
Archive表支持行级锁 
Archive支持auto_incrment列，且其列上可以包含一个索引，但在其他字段上不能创建索引 
Archive不支持对auto_incrment列插入一个小于当前最大值的数据 
Archive存储引擎会用zlib来压缩数据 

 
MySQL存储引擎之ARCHIVE  

 
MySQL存储引擎之Blackhole 
Blackhole存储引擎用来接收表插入请求，但不存储数据，所以查询表数据总是返回空，通常用在主从复制的情况下当主库不想保留数据而从库通过复制语句执行而保留数据的情况 
Blackhole表在磁盘会创建一个文件，.frm文件用来存储表结构 
mysql> CREATE TABLE test(i INT, c CHAR(10)) ENGINE = BLACKHOLE; 
Query OK, 0 rows affected (0.03 sec)  
mysql> INSERT INTO test VALUES(1,'record one'),(2,'record two'); 
Query OK, 2 rows affected (0.00 sec) 
Records: 2  Duplicates: 0  Warnings: 0  
mysql> SELECT * FROM test; 
Empty set (0.00 sec) 

 
MySQL存储引擎之Merge 
Merge存储引擎可以将一批字段相同，索引相同且顺序相同的MyISAM表在逻辑上看做是同一个 
Merge表在磁盘上创建两个文件，.frm文件保存表的结构信息，.MRG文件包含所有被视作同一个表的MyISAM表 
Merge表支持select,delete,update,insert语句执行 
创建merge表时需要执行union子句，用来将指定的MyISAM结合起来，insert_method选项用来指定插入语句是将数据插入到第一个表FIRST还是最后一个表LAST中，或者不指定或NO代表不允许插入 
mysql> CREATE TABLE t1 ( 
    ->    a INT NOT NULL AUTO_INCREMENT PRIMARY KEY, 
    ->    message CHAR(20)) ENGINE=MyISAM; 
    mysql> CREATE TABLE t2 ( 
    ->    a INT NOT NULL AUTO_INCREMENT PRIMARY KEY, 
    ->    message CHAR(20)) ENGINE=MyISAM; 
    mysql> INSERT INTO t1 (message) VALUES ('Testing'),('table'),('t1'); 
mysql> INSERT INTO t2 (message) VALUES ('Testing'),('table'),('t2'); 
mysql> CREATE TABLE total ( 
    ->    a INT NOT NULL AUTO_INCREMENT, 
    ->    message CHAR(20), INDEX(a)) 
    ->    ENGINE=MERGE UNION=(t1,t2) INSERT_METHOD=LAST; 

     
MySQL存储引擎之Merge 
mysql> SELECT * FROM total; 
+---+---------+ 
| a | message | 
+---+---------+ 
| 1 | Testing | 
| 2 | table   | 
| 3 | t1      | 
| 1 | Testing | 
| 2 | table   | 
| 3 | t2      | 
+---+---------+ 

 
MySQL存储引擎之Federated 
Federated存储引擎提供了从一个MySQL实例连接其它实例上数据的能力 
Federated存储引擎默认是disable状态，如果要开启，则需要在启动MySQL时使用—federated选项 
CREATE TABLE federated_table ( 
    id     INT(20) NOT NULL AUTO_INCREMENT, 
    name   VARCHAR(32) NOT NULL DEFAULT '', 
    other  INT(20) NOT NULL DEFAULT '0', 
    PRIMARY KEY  (id), 
    INDEX name (name), 
    INDEX other_key (other) 
) 
ENGINE=FEDERATED 
DEFAULT CHARSET=latin1 
CONNECTION='mysql://fed_user@remote_host:9306/federated/test_table'; 

 
MySQL存储引擎之Federated 
CREATE SERVER fedlink 
FOREIGN DATA WRAPPER mysql 
OPTIONS (USER 'fed_user', HOST 'remote_host', PORT 9306, DATABASE 'federated');  
CREATE TABLE test_table ( 
    id     INT(20) NOT NULL AUTO_INCREMENT, 
    name   VARCHAR(32) NOT NULL DEFAULT '', 
    other  INT(20) NOT NULL DEFAULT '0', 
    PRIMARY KEY  (id), 
    INDEX name (name), 
    INDEX other_key (other) 
) 
ENGINE=FEDERATED 
DEFAULT CHARSET=latin1 
CONNECTION='fedlink/test_table'; 

 
MySQL存储引擎之Example 
Example存储引擎只存在于MySQL源码中，只针对开发者，对实际的数据库使用者没有太大的意义 
Example表只保留表结构，本身不保存数据 
mysql> CREATE TABLE test (i INT) ENGINE = EXAMPLE; 
Query OK, 0 rows affected (0.78 sec)  
mysql> INSERT INTO test VALUES(1),(2),(3); 
ERROR 1031 (HY000): Table storage engine for 'test' doesn't ? 
                    have this option  
                    mysql> SELECT * FROM test; 
                    Empty set (0.31 sec) 

                     
MySQL存储引擎之NDB 
NDB存储引擎专用在MySQL Cluster软件中，是MySQL自己推出的提高可用性和可靠性的集群软件 

 
数据库设计方法之E-R模型 
E-R模型在数据库概念设计阶段广泛采用 
E-R模型的构成成分是实体集，属性和联系集 
实体是把具有属性、性质和特征相同的实体，用所有实体名和他的属性名称的来用抽象的形式描述同种类的实体;一般实体用矩形符号表示，矩形框内标注实体的名称 
属性是实体特有的一些特性，一个实体有包含有许多个属性来进行描述。用椭圆形的符号来表示，用无向直线将属性的椭圆和其相对应的实体图形串联起来 
联系是ER模型中的联系是用来反映实体内部和实体之间的属性关系。用菱形符号表示，在菱形框内标注联系的名称，然后使用无向直线将有关系的实体属性串联起来，还需要在无向直线上标出实体和联系的类型（1 : 1，1 : n或m : n） 

 
数据库设计方法之E-R模型 
例如系、学生和课程的联系的E-R模型 
系、学生和课程作为实体集；一个系有多个学生，而一个学生仅属于一个系，所以系和学生之间是一对多的联系；一个学生可以选修多门课程，而一门课程有多个学生选修，所以学生和课程之间是多对多的联系 
E-R模型的设计也要遵循三范式的要求 

 
数据库设计方法之第三范式 
第三范式是数据库逻辑设计阶段的有效方法 
第一范式(1NF)：表示每个属性都不可分，是数据库设计的基本要求 
仅符合第一范式存在数据冗余过大，插入删除和修改异常的问题 

 
数据库设计方法之第三范式 
假如学校新建了一个系，但是暂时还没有招收任何学生（比如3月份就新建了，但要等到8月份才招生），那么是无法将系名与系主任的数据单独地添加到数据表中去的 （注１）——插入异常 
假如将某个系中所有学生相关的记录都删除，那么所有系与系主任的数据也就随之消失了（一个系所有学生都没有了，并不表示这个系就没有了）。——删除异常 
假如李小明转系到法律系，那么为了保证数据库中数据的一致性，需要修改三条记录中系与系主任的数据。——修改异常 
第二范式(2NF)：在第一范式的基础上限制一个表只能表达一个实体 
第三范式(3NF)：在第二范式的基础上，表中的每列都和主键有直接关系，而不存在传递性依赖 

 
数据库设计方法之第三范式  

 
数据库设计方法之第三范式 
符合3NF要求的数据库设计，基本上解决了数据冗余过大，插入异常，修改异常，删除异常的问题。当然，在实际中，往往为了性能上或者应对扩展的需要，经常 做到2NF或者1NF，但是作为数据库设计人员，至少应该知道，3NF的要求是怎样的 
范式的优势在于： 
当数据较好的范式化时，就只有很少或者没有重复数据，所以只需要修改更少的数据，也就意味着更新操作更快 
范式化的表通常更小，可以更好地放在内存里，所以执行操作会更快 
很少有多余的数据意味着检索列表数据时更少需要distinct或者group by语句 
范式化的缺点在于通常需要表关联，有的时候会显得代价昂贵 

 
数据库设计方法之第三范式 
事实上在是否对每个应用场景使用第三范式设计表结构不是固定的，要具体问题具体分析，实际情况中经常混合使用 
比如按照第三范式只会在teacher表中存储老师姓名的字段，但如果经常会碰到查询课程信息和对应老师姓名的语句时，就可以考虑将老师姓名也存储在course表中一份，以便在查询时不再需要两个表关联查询。 
当然还要评估老师姓名是否会经常变化，如果经常变化那更新两个表的代价也比较大，所以要综合评估 
另外也有排序的需要，当子表中的字段要依赖父表的信息排序时，关联排序往往代价很大，可以考虑冗余排序字段到子表并建立索引。 
当比如要经常计算老师的课程数时，可以考虑将课程数冗余到老师表里，并在课程表更新时同时更新老师表里的该字段 

 
数据库常用设计工具之powerdesigner 
Power Designer 是Sybase公司的CASE工具集，使用它可以方便地对管理信息系统进行分析设计，他几乎包括了数据库模型设计的全过程。利用Power Designer可以制作数据流程图、概念数据模型、物理数据模型，还可以为数据仓库制作结构模型，也能对团队设计模型进行控制。 
power designer是能进行数据库设计的强大的软件，是一款开发人员常用的数据库建模工具。使用它可以分别从概念数据模型(Conceptual Data Model)和物理数据模型(Physical Data Model)两个层次对数据库进行设计。在这里，概念数据模型描述的是独立于数据库管理系统(DBMS)的实体定义和实体关系定义；物理数据模型是在概念数据模型的基础上针对目标数据库管理系统的具体化。 

 
数据库常用设计工具之powerdesigner  

 
数据库常用设计工具之powerdesigner  

 
数据库常用设计工具之powerdesigner  

 
数据库常用设计工具之powerdesigner  

 
数据库常用设计工具之powerdesigner  

 
数据库常用设计工具之powerdesigner  

 
数据库常用设计工具之powerdesigner  

 
数据库常用设计工具之powerdesigner  

 
数据库常用设计工具之workbench 
MySQL Workbench是一款专为MySQL设计的ER/数据库建模工具。它是著名的数据库设计工具DBDesigner4的继任者。你可以用MySQL Workbench设计和创建新的数据库图示，建立数据库文档，以及进行复杂的MySQL 迁移。 
MySQL Workbench是下一代的可视化数据库设计、管理的工具，它同时有开源和商业化的两个版本。该软件支持Windows和Linux系统，  

 
数据库常用设计工具之workbench 
下载地址https://dev.mysql.com/downloads/workbench/  

 
数据库常用设计工具之workbench  

 
数据库常用设计工具之workbench  

 
数据库常用设计工具之workbench  

 
数据库常用设计工具之workbench  

 
数据库常用设计工具之workbench  

 
数据库常用设计工具之workbench  

 
数据库常用设计工具之workbench  

 
数据库常用设计工具之workbench  

 
数据库常用设计工具之workbench  

 
数据库常用设计工具之workbench  

 
数据库设计方法之字段属性 
表字段在设计时还需考虑字段属性 
NOT NULL代表表中此列的数据必须存在，默认是NULL容许为空 
主键(primary key)代表此表的所有数据都可以被主键里的字段区分，创建完主键则默认会在对应字段上创建唯一性索引，且每个主键字段都需要是NOT NULL，一个表上只允许有一个主键 
外键(foreign key)可以将两个表的数据建立映射关系，并定义不同的外键约束条件以保证数据的一致性，通常为一个父表一个子表，子表中的数据映射到父表对应的列 

 
数据库设计方法之字段属性 
Index_name代表外键ID，默认情况下MySQL会在子表上创建一个外键索引 
MySQL推荐在父表和子表的相关字段上都创建索引，以避免全表扫描 
当创建好外键后，任何对子表的插入和修改操作如果对应的值没有在父表中有对应，都会被MySQL拒绝 
当在父表上update和delete操作时，对子表中对应数据的操作依赖设置 
Cascade代表子表中的数据也自动update和delete 
Set null代表子表中的数据自动修改成null 
Restrict(默认)代表如果子表中有对应的数据，则拒绝父表上的update和delete操作 
No action在MySQL中的含义和restrict一样 
Set default代表将子表中的数据自动修改成default值 
备注comment可以用在表或者字段上，表达表和字段的含义，最常可以包含1024个字符 

 
数据库设计方法之字段属性 
CREATE TABLE product ( 
    category INT NOT NULL, id INT NOT NULL, 
    price DECIMAL, 
    PRIMARY KEY(category, id) 
)   ENGINE=INNODB;  
CREATE TABLE customer ( 
    id INT NOT NULL, 
    PRIMARY KEY (id) 
)   ENGINE=INNODB;  
CREATE TABLE product_order ( 
    no INT NOT NULL AUTO_INCREMENT, 
    product_category INT NOT NULL, 
    product_id INT NOT NULL, 
    customer_id INT NOT NULL,  
    PRIMARY KEY(no), 
    INDEX (product_category, product_id), 
    INDEX (customer_id),  
    FOREIGN KEY (product_category, product_id) 
      REFERENCES product(category, id) 
      ON UPDATE CASCADE ON DELETE RESTRICT,  
    FOREIGN KEY (customer_id) 
      REFERENCES customer(id) 
      )   ENGINE=INNODB; 

       
THANKS 