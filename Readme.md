## 基于gin 实现限流算法

### 项目目录
```
├── go.mod
├── go.sum
├── main.go // gin server 路由
└── src // 项目主逻辑
    ├── Bucket.go   // 令牌桶算法
    ├── ip_limit.go // 实现ip限流器（基于令牌桶、LRU都有实现）
    ├── limit.go    // 普通限流器
    ├── listCache.go // LRU算法   
    ├── listCache_test.go
    ├── paramlimit.go   // 请求参数限流器
    └── test    // 简单测试文件
        └── testlist.go

```

### 令牌桶
#### 1. 建立桶结构
```
1.桶容量
2.当前桶有多少令牌
3.互斥锁
```
