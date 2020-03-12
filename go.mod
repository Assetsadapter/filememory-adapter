module github.com/Assetsadapter/filememory-adapter

go 1.12

require (
	github.com/asdine/storm v2.1.2+incompatible
	github.com/astaxie/beego v1.11.1
	github.com/blocktree/go-owcdrivers v1.0.12
	github.com/blocktree/go-owcrypt v1.0.1
	github.com/blocktree/openwallet v1.5.5
	github.com/elastic/gosigar v0.10.5 // indirect
	github.com/ethereum/go-ethereum v1.9.1
	github.com/gin-gonic/gin v1.5.0
	github.com/imroc/req v0.2.3
	github.com/shopspring/decimal v0.0.0-20180709203117-cd690d0c9e24
	github.com/steakknife/bloomfilter v0.0.0-20180922174646-6819c0d2a570 // indirect
	github.com/steakknife/hamming v0.0.0-20180906055917-c99c65617cd3 // indirect
	github.com/tidwall/gjson v1.2.1
)

//replace github.com/blocktree/openwallet => ../../openwallet
