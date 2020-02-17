# leoCoin
#### envs:
* golang 12.5

----
## 2020.02.17
>- 新建一个transaction数据结构用于记录转账信息

>- 将Block中的Data信息替换为transaction数据
>- Block中增加timestamp
>- Block中修改打包生成哈希的源内容

>- BlockChain新增difficulty
>- BlockChain新增transaction pool，用于存储所有的转账信息等待矿工打包
>- BlockChain新增矿工奖励
>- BlockChain修改挖矿和block的添加方式
>- BlockChain新增向链上添加转账记录的方法
#### TBD
>- block、chain的构造函数
>- 加密
>- 新增单元测试
>- BlockChain难度调整
>- BlockChain奖励减半
>- 数据持久化（transaction存入数据库）
>- 网络化（分布式网络）
----
## 2019.12.22

> - 修改生成当前hash值为工作量证明(POW))方式
> - 增加Block.Diff成员，使用difficult来判断hash前多少位为0，而不是在mine()中写死
> - 在验证函数中加入根据当前difficult验证hash前difficult位是否为0  
#### TBD
> - 后续加入转账、网络化功能

------

## 2019.11.13
> - 当前已经完成大概框架
#### TBD
> - 后续继续加入工作量证明、网络化等工作 

------

