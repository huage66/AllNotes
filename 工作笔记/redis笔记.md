## Redis的基本命令





## Redis 的配置文件项

二、Redis的配置
下面列举了Redis中的一些常用配置项：
daemonize 如果需要将Redis服务以守护进程在后台运行，则把该项的值改为yes

pidfile 配置多个pid的地址，默认在/var/run/redis/pid

bind 绑定ip，设置后只接受来自该ip的请求

port 监听端口，默认是6379

timeout 客户端连接超时的设定，单位是秒

loglevel 分为4级，debug、verbose、notice、warning

logfile 配置log文件地址

databases 设置数据库的个数，默认使用的数据库为0

save 设置redis进行数据库镜像的频率

rdbcompression 在进行镜像备份时，是否进行压缩

Dbfilename 镜像备份文件的文件名

Dir 数据库镜像备份文件的存放路径

Slaveof 设置数据库为其他数据库的从数据库

Masterauth 主数据库连接需要的密码验证
Requirepass 设置登录时，需要使用的密码
Maxclients 设置同时连接的最大客户端数量
Maxmemory 设置redis能够使用的最大内存
Appendonly 开启append only模式
Appendfsync 设置对appendonly.aof文件同步的频率
vm-enabled 是否开启虚拟内存支持
vm-swap-file 设置虚拟内存的交换文件路径
vm-max-memory 设置redis能够使用的最大虚拟内存
vm-page-size 设置虚拟内存的页大小
vm-pages 设置交换文件的总的page数量
vm-max-threads 设置VMIO同时使用的线程数量
Glueoutputbuf 把小的输出缓存存放在一起
hash-max-zipmap-entries 设置hash的临界值
Activerehashing 重新hash

修改redis的配置参数：
vi /usr/local/redis/etc/redis.conf
将daemonize no改为daemonize yes，保存退出。
再来启动redis服务器
cd /usr/local/redis/bin
./redis-server /usr/local/redis/etc/redis.conf 启动redis并指定配置文件

ps aux | grep redis 查看redis是否启动成功

netstat -tlun 查看主机的6379端口是否在使用（监听）

./redis-cli 打开redis的客户端

quit 退出redis的客户端

pkill redis-server 关闭redis服务器

**./redis-cli shutdown 也可以通过这条命令关闭redis服务器**

下面列举了Redis中的一些常用配置项：
daemonize 如果需要将Redis服务以守护进程在后台运行，则把该项的值改为yes

pidfile 配置多个pid的地址，默认在/var/run/redis/pid

bind 绑定ip，设置后只接受来自该ip的请求

port 监听端口，默认是6379

timeout 客户端连接超时的设定，单位是秒

loglevel 分为4级，debug、verbose、notice、warning

logfile 配置log文件地址

databases 设置数据库的个数，默认使用的数据库为0

save 设置redis进行数据库镜像的频率

rdbcompression 在进行镜像备份时，是否进行压缩

Dbfilename 镜像备份文件的文件名

Dir 数据库镜像备份文件的存放路径

Slaveof 设置数据库为其他数据库的从数据库

Masterauth 主数据库连接需要的密码验证
Requirepass 设置登录时，需要使用的密码
Maxclients 设置同时连接的最大客户端数量
Maxmemory 设置redis能够使用的最大内存
Appendonly 开启append only模式
Appendfsync 设置对appendonly.aof文件同步的频率
vm-enabled 是否开启虚拟内存支持
vm-swap-file 设置虚拟内存的交换文件路径
vm-max-memory 设置redis能够使用的最大虚拟内存
vm-page-size 设置虚拟内存的页大小
vm-pages 设置交换文件的总的page数量
vm-max-threads 设置VMIO同时使用的线程数量
Glueoutputbuf 把小的输出缓存存放在一起
hash-max-zipmap-entries 设置hash的临界值
Activerehashing 重新hash

修改redis的配置参数：
vi /usr/local/redis/etc/redis.conf
将daemonize no改为daemonize yes，保存退出。
再来启动redis服务器
cd /usr/local/redis/bin
./redis-server /usr/local/redis/etc/redis.conf 启动redis并指定配置文件

ps aux | grep redis 查看redis是否启动成功

netstat -tlun 查看主机的6379端口是否在使用（监听）

./redis-cli 打开redis的客户端

quit 退出redis的客户端

pkill redis-server 关闭redis服务器

./redis-cli shutdown 也可以通过这条命令关闭redis服务器



