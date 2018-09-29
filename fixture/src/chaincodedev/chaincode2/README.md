
#### chaincode (dev mode)
function statement in receipt  

* init:初始化 参数[json] 

```
./peer chaincode instantiate -n mycc -v 1.0 -c '{"Args":["init","{\"receiptId\":\"r00234387ed1849e4\",\"message\":\"name|sinodata\",\"image\":\"ox921389213\",\"waterImage\":\"0x677137289137\",\"owner\":{\"ownerId\":\"u000\",\"ownerName\":\"user1\",\"ownerPw\":\"1\"},\"shshareList\":[]}"]}' -o 127.0.0.1:7050 -C originchan

```

* readByReciptId:使用receiptId读取receipt 参数[receiptId]

``` 
./peer chaincode query -n mycc -v 1.0 -c '{"Args":["readByReciptId","r00234387ed1849e4"]}' -o 127.0.0.1:7050 -C originchan
```

* readByReciptIdAndOwner:使用receiptId和ownerId读取receipt 参数[receiptId,ownerId]

``` 
./peer chaincode query -n mycc -v 1.0 -c '{"Args":["readByReciptIdAndOwner","r00234387ed1849e4","u000"]}' -o 127.0.0.1:7050 -C originchan
```

* readByMessage:使用全文索引读取receipt 参数[message]

``` 
./peer chaincode query -n mycc -v 1.0 -c '{"Args":["readByMessage","sinodata"]}' -o 127.0.0.1:7050 -C originchan
```

* readByMessageAndOwner:使用全文索引和ownerId读取receipt 参数[message,ownerId]

``` 
./peer chaincode query -n mycc -v 1.0 -c '{"Args":["readByMessageAndOwner","sinodata","u000"]}' -o 127.0.0.1:7050 -C originchan
```

* write:新建receipt 参数[json]

```
./peer chaincode invoke -n mycc -v 1.0 -c '{"Args":["write","{\"receiptId\":\"r8287373de1839465de1839465\",\"message\":\"纳税人识别号|91110102560430172A\",\"image\":\"0x3216983762178\",\"waterImage\":\"0x13798127\",\"owner\":{\"ownerId\":\"u002\",\"ownerName\":\"user2\",\"ownerPw\":\"1\"},\"shshareList\":[]}"]}' -o 127.0.0.1:7050 -C originchan
 
```

* update:更新receipt 参数[json]

```
./peer chaincode invoke -n mycc -v 1.0 -c '{"Args":["update","{\"receiptId\":\"r8287373de1839465\",\"message\":\"纳税人识别号|91110172A\",\"image\":\"0x32162178\",\"waterImage\":\"0x138127\",\"owner\":{\"ownerId\":\"u002\",\"ownerName\":\"user2\",\"ownerPw\":\"1\"},\"shshareList\":[]}"]}' -o 127.0.0.1:7050 -C originchan

```

* readOwnerAll:读取所有的owner

```
./peer chaincode query -n mycc -v 1.0 -c '{"Args":["readOwnerAll"]}' -o 127.0.0.1:7050 -C originchan
```

* readReceiptAll:读取所有的receipt

```
./peer chaincode query -n mycc -v 1.0 -c '{"Args":["readReceiptAll"]}' -o 127.0.0.1:7050 -C originchan
```

* initOwner:新建owner 参数[json]

```
./peer chaincode invoke -n mycc -v 1.0 -c '{"Args":["initOwner","{\"ownerId\":\"u002\",\"ownerName\":\"user2\",\"ownerPw\":\"1\"}"]}' -o 127.0.0.1:7050 -C originchan
```

* setOwner:设置receipt的owner 参数[receiptId,ownerId]

```
./peer chaincode invoke -n mycc -v 1.0 -c '{"Args":["setOwner","r00234387ed1849e4","u002"]}' -o 127.0.0.1:7050 -C originchan
```

* setShareList:owner加入至共享者列表 参数[receiptId,ownerId]

```
./peer chaincode invoke -n mycc -v 1.0 -c '{"Args":["setShareList","r00234387ed1849e4","u002","user02"]}' -o 127.0.0.1:7050 -C originchan
```

* readShareList:通过receiptId读取sharelist 参数[receiptId]

```
./peer chaincode query -n mycc -v 1.0 -c '{"Args":["readShareList","r00234387ed1849e4"]}' -o 127.0.0.1:7050 -C originchan
```

* readByOwner:通过ownerId读取receipt 参数[ownerId]

```
./peer chaincode query -n mycc -v 1.0 -c '{"Args":["readByOwner","u002"]}' -o 127.0.0.1:7050 -C originchan
```

* readShareListByOwner:读取sharelist中包含ownerId所在的receipt 参数[ownerId]

```
./peer chaincode query -n mycc -v 1.0 -c '{"Args":["readShareListByOwner","u002"]}' -o 127.0.0.1:7050 -C originchan
```

* readByOidAll:sharelist中含有ownerId的receipt与owner中含有ownerId的receipt的并集 [ownerId]

```
./peer chaincode query -n mycc -v 1.0 -c '{"Args":["readByOidAll","u002"]}' -o 127.0.0.1:7050 -C originchan
```

* verityOwner:校验密码 参数[ownerId]

```
./peer chaincode invoke -n mycc -v 1.0 -c '{"Args":["verityOwner","u002","1"]}' -o 127.0.0.1:7050 -C originchan
```

* delete:删除receipt 参数[receiptId]

```
./peer chaincode invoke -n mycc -v 1.0 -c '{"Args":["delete","r00234387ed1849e4"]}' -o 127.0.0.1:7050 -C originchan

```