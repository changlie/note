最近因为项目需要，简单的试用了两款高可用开源方案：Keepalived和Heartbeat。两者都很流行，但差异还是很大的，现将试用过程中的感受以及相关知识点简单总结一下，供大家选择方案的时候参考。

1. Keepalived使用更简单：从安装、配置、使用、维护等角度上对比，Keepalived都比Heartbeat要简单得多，尤其是Heartbeat 2.1.4后拆分成3个子项目，安装、配置、使用都比较复杂，尤其是出问题的时候，都不知道具体是哪个子系统出问题了；而Keepalived只有1个安装文件、1个配置文件，配置文件也简单很多；

2. Heartbeat功能更强大：Heartbeat虽然复杂，但功能更强大，配套工具更全，适合做大型集群管理，而Keepalived主要用于集群倒换，基本没有管理功能；

3. 协议不同：Keepalived使用VRRP协议进行通信和选举，Heartbeat使用心跳进行通信和选举；Heartbeat除了走网络外，还可以通过串口通信，貌似更可靠；

4. 使用方式基本类似：如果要基于两者设计高可用方案，最终都要根据业务需要写自定义的脚本，Keepalived的脚本没有任何约束，随便怎么写都可以；Heartbeat的脚本有约束，即要支持service start/stop/restart这种方式，而且Heartbeart提供了很多默认脚本，简单的绑定ip，启动apache等操作都已经有了；



> 使用建议：优先使用Keepalived，当Keepalived不够用的时候才选择Heartbeat