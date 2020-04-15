# 这是一个爬虫项目，爬取B站

1. 这是一个爬取b站数据的爬虫，整体使用golang，目前爬虫开发已经接近尾声，感兴趣的同学可以运行dist中的exe文件
2. 关于exe文件参数说明，mongodbUrl默认localhost:27017，db默认bspider
3. 关于一些其他的参数可以运行如下命令 bspider.exe -help, 将会有提示
4. ![tut](https://github.com/jerryzhuochuxin/picture/blob/master/20200415232213.png)
5. 如果想设置product模式: bspider.exe -debug=false
6. 默认debug模式， 会将debug的日志打印出来，这个模型一运行会出现很多日志，如果mongodb数据库是不可用的将会阻塞小段时间并且报错，推荐用cmd运行
7. ![tut2](https://github.com/jerryzhuochuxin/picture/blob/master/20200415234421.png)
8. 注：window环境下什么依赖都不用准备，只用运行bspider.exe文件即可，另外确保mongodb的url输入正确，服务可用，有点尴尬如果mongodb不可用，这个程序有个超时，可能不是很友好，目前正在改进
9. 当mongodb不可用时，会爆如下错误
10. ![tut3](https://github.com/jerryzhuochuxin/picture/blob/master/20200415234803.png)

