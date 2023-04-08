https://coredns.io/plugins/forward/
https://coredns.io/plugins/hosts/#the-hosts-file
https://coredns.io/plugins/whoami/


https://help.aliyun.com/document_detail/380963.html
https://zhuanlan.zhihu.com/p/387806561


### 配置示例：
```
.:53 {
        hosts {
          192.168.1.105 go.cn
          192.168.1.100 qk.cn
          192.168.1.103 at.cn
          192.168.1.105 g.cn
          192.168.1.100 q.cn
          192.168.1.103 a.cn
          192.168.1.101 x.cn
          fallthrough
        }
        forward . 223.5.5.5 119.29.29.29 114.114.114.114
        log
        errors
        whoami
}
```


