### 从github上拉项目并且成功运行

刚拉去下来的时候没有导入依赖，发现go get github...怎么样都不成功

拉去并运行正确步骤：
1. 查看是否有mod
2. 进入该目录添加mod 命令： go mod init 
3. 在go.mod下 添加go.sum 复制粘贴
4. 