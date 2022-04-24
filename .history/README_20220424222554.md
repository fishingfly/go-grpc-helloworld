<!--
 * @Author: zhounanjun
 * @Date: 2022-04-24 22:09:34
 * @LastEditors: zhounanjun
 * @LastEditTime: 2022-04-24 22:25:54
 * @Description: 请填写简介
-->

步骤:
- 编写proto/helloworld.proto文件
- 运行
  ```
  protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/helloworld.proto
  ```
  生成了helloworld.pb.go文件和helloword_grpc.pb.go文件
- 编写server/main.go文件,启动grpc服务
- 运行grpcui -plaintext localhost:9001进行界面调试;

作为服务方,上面一个简单的helloworld的服务已经可以启动了,运行go run server/main.go就行

但是作为调用方还需要实现client,见client/main.go文件

