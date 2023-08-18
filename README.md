## gRPC Go服务端demo

### 关于.proto文件的一些语法说明
- https://www.cnblogs.com/zhangcaiwang/p/15755264.html


### 测试
1. 编译Go的grpc代码，编写server端逻辑
```
cd pbs;
./generate.sh
```
2. 启动Server端，进入项目根目录，编译&执行
```
go build -o server .
./server
```