# filememory-adapter

本项目适配了openwallet.AssetsAdapter接口，给应用提供了底层的区块链协议支持。

## 项目依赖库

- [go-owcrypt](https://github.com/blocktree/go-owcrypt.git)
- [go-owcdrivers](https://github.com/blocktree/.git)

## 如何测试

openwtester包下的测试用例已经集成了openwallet钱包体系，创建conf文件，新建ETH.ini文件，编辑如下内容：

```ini
# is enable scanner
isScan = false

# wallet api url
ServerAPI = "https://chain.fmchain.cc/exchange/"

# block chain ID
ChainID = 39482

# gas limit
GasLimit = 500000

# gas price
GasPrice = 18

# Summery transaction get addresses balance concurrency channel control, default value is 5;
SumThreadControl = 1

# Cache data file directory, default = "", current directory: ./data
dataDir = ""
```
