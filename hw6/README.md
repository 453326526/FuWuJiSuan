# golang-datamysql

在 java 的世界中， jdbc pk orm 是最常见的，目前 ORM 占上风。 golang 中存在同样问题！

	1. xorm 推荐，有完善中文文档：![xorm](https://github.com/go-xorm/xorm)
	2. gorm 推荐，：![gorm](https://github.com/jinzhu/gorm)
	3. beego orm 有较完善的中支持： ![beego orm](https://beego.me/docs/mvc/model/overview.md)


## 命令列表

有如下可用的命令：

	1. reverse 反转一个数据库结构，生成代码
	2. shell 通用的数据库操作客户端，可对数据库结构和数据操作
	3. dump Dump数据库中所有结构和数据到标准输出
	4. source 从标注输入中执行SQL文件
	5. driver 列出所有支持的数据库驱动

## 特性

	1. 支持Struct和数据库表之间的灵活映射，并支持自动同步表结构
	2. 事务支持
	3. 支持原始SQL语句和ORM操作的混合执行
	4. 使用连写来简化调用
	5. 支持使用Id, In, Where, Limit, Join, Having, Table, Sql, Cols等函数和结构体等方式作为条件
	6. 支持级联加载Struct
	7. 支持LRU缓存(支持memory, memcache, leveldb, redis缓存Store) 和 Redis缓存
	8. 支持反转，即根据数据库自动生成xorm的结构体
	9. 支持事件
	10. 支持created, updated, deleted和version记录版本（即乐观锁）

## 程序结构

	data/sql编程使用的是java经典的“entity - dao - service”层次结构模型。而在xorm中，实现了dao的自动化。
	即，在service层中，就可以完成与数据库的交互，省去了dao（持久层）结构。
	不仅程序结构上得到了很大的简化，编程效率也有很大的提升。


## 测试

利用ApacheBench进行压力测试，比较data/sql与xorm：

```shell
stonelm@ubuntu:~/Desktop/gocode/hw6$ ab -n 1000 -c 100 http://localhost:8080/service/userinfo?userid=

This is ApacheBench, Version 2.3 <$Revision: 1757674 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /service/userinfo?userid=
Document Length:        663 bytes

Concurrency Level:      100
Time taken for tests:   0.412 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      683000 bytes
HTML transferred:       559000 bytes
Requests per second:    2956.25 [#/sec] (mean)
Time per request:       37.256 [ms] (mean)
Time per request:       0.422 [ms] (mean, across all concurrent requests)
Transfer rate:          1888.89 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   0.9      0       5
Processing:     1   34  29.5     27     137
Waiting:        1   33  29.5     27     137
Total:          1   34  29.7     27     139
WARNING: The median and mean for the initial connection time are not within a normal deviation
        These results are probably not that reliable.

Percentage of the requests served within a certain time (ms)
  50%     22
  66%     34
  75%     36
  80%     42
  90%     90
  95%    117
  98%    123
  99%    134
 100%    143 (longest request)
```

- 吞吐率(Request per second)：2956.25 [requests/sec]
- 用户平均等待时间(Time per request)：37.256 [ms]
- 服务器平均等待时间(Time per request:across all concurrent requests)：0.422 [ms]



## XORM

```shell
stonelm@ubuntu:~/Desktop/gocode/hw6$ ab -n 1000 -c 100 http://localhost:8080/service/userinfo?userid=

This is ApacheBench, Version 2.3 <$Revision: 1757674 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /service/userinfo?userid=
Document Length:        462 bytes

Concurrency Level:      100
Time taken for tests:   0.272 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      568000 bytes
HTML transferred:       444000 bytes
Requests per second:    3671.22 [#/sec] (mean)
Time per request:       23.673 [ms] (mean)
Time per request:       0.312 [ms] (mean, across all concurrent requests)
Transfer rate:          2399.358 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   0.9      0       7
Processing:     1   25  11.5     27      68
Waiting:        1   25  11.5     26      66
Total:          1   29  14.1     28      68
WARNING: The median and mean for the initial connection time are not within a normal deviation
        These results are probably not that reliable.

Percentage of the requests served within a certain time (ms)
  50%     26
  66%     30
  75%     33
  80%     36
  90%     41
  95%     47
  98%     53
  99%     60
 100%     72 (longest request)
```

- 吞吐率(Request per second)：3671.22 [requests/sec]
- 用户平均等待时间(Time per request)：23.673 [ms]
- 服务器平均等待时间(Time per request:across all concurrent requests)：0.312 [ms]


可以看到，xorm性能更好。
