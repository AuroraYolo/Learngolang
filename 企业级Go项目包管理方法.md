## Go包管理困镜
 - 没有统一的包管理方式
 - 包之间的依赖关系很难维护
 - 如果同时需要一个包的不同版本,非常麻烦

## 尝试解决
 - 尝试使用godep,govendor.glide等解决
 - 未彻底解决GOPATH存在的问题
 - 使用起来麻烦

## Go Modules
 - 本质上，一个Go包就是一个项目的源码
 - gomod的作用:将Go包和Git项目关联起来
 - Go包的版本就是git的tag
 - gomod就是解决需要那个git项目的什么版本

## 使用Modules
 - github.com/Jeffail/tunny
    - go get github.com/Jeffail/tunny
    - github.com/Jeffail/tunny@0.1.3
   
## Github无法访问怎么办
 - 使用goproxy.cn代理

## 想用本地文件替代怎么办
 - go.mod文件追加
   - replace github.com/Jeffail/tunny => xxx/xx
 - go wendor 缓存到本地
   - go mod vendor //不是之前的go vendor
   - go    build -mod vendor
   
## 创建Go Module
 - 删除本地go.mod
   - go mod init github.com/AuroraYolo/moody 
   
 - 推送至代码仓库
 - 增加新版本时,在仓库打新Tag
 