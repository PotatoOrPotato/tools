
#### chaincode (dev mode)
function statement in receipt  

* init  

```
./peer chaincode instantiate -n mycc -v 1.0 -c '{"Args":["init","{\"id\":\"r15647670\",\"purchaser\":\"liwei\",\"seller\":\"chinamobile\",\"sum\":200,\"imageByte\":\"4AAQSkZ\",\"ownerRelation\":{\"id\":\"o12345\",\"username\":\"ocr1\"},\"shshareList\":[]}"]}' -o 127.0.0.1:7050 -C originchan

```

* read (need 1 args)

``` 
./peer chaincode invoke -n mycc -v 1.0 -c '{"Args":["read","r15647670"]}' -o 127.0.0.1:7050 -C originchan
```

* write 

```
./peer chaincode invoke -n mycc -v 1.0 -c '{"Args":["write","{\"id\":\"r876497\",\"purchaser\":\"xiaoming\",\"seller\":\"chinamobileBJ\",\"sum\":300,\"imageByte\":\"ABAQAA\",\"ownerRelation\":{\"id\":\"o225\",\"username\":\"siwa\"},\"shshareList\":[]}"]}' -o 127.0.0.1:7050 -C originchan
 
```

* update

```
./peer chaincode invoke -n mycc -v 1.0 -c '{"Args":["update","{\"id\":\"r876497\",\"purchaser\":\"xiaoming\",\"seller\":\"chinamobileSH\",\"sum\":500,\"imageByte\":\"ABAQAA\",\"ownerRelation\":{\"id\":\"o225\",\"username\":\"siwa\"},\"shshareList\":[]}"]}' -o 127.0.0.1:7050 -C originchan

```

* history (need 1 args) 

```
./peer chaincode invoke -n mycc -v 1.0 -c '{"Args":["history","r15647670"]}' -o 127.0.0.1:7050 -C originchan
```

* readOwnerAll

```
./peer chaincode invoke -n mycc -v 1.0 -c '{"Args":["readOwnerAll"]}' -o 127.0.0.1:7050 -C originchan
```

* readReceiptAll

```
./peer chaincode invoke -n mycc -v 1.0 -c '{"Args":["readReceiptAll"]}' -o 127.0.0.1:7050 -C originchan
```

* initOwner

```
./peer chaincode invoke -n mycc -v 1.0 -c '{"Args":["initOwner","{\"id\":\"o223\",\"username\":\"erwa\"}"]}' -o 127.0.0.1:7050 -C originchan
```

* setOwner (need 2 args)

```
./peer chaincode invoke -n mycc -v 1.0 -c '{"Args":["setOwner","r15647670","o228"]}' -o 127.0.0.1:7050 -C originchan
```

* setShareList (need 2 args)

```
./peer chaincode invoke -n mycc -v 1.0 -c '{"Args":["setShareList","r15647670","o223","erwa"]}' -o 127.0.0.1:7050 -C originchan
```

* readShareList (need 1 args)

```
./peer chaincode invoke -n mycc -v 1.0 -c '{"Args":["readShareList","r15647670"]}' -o 127.0.0.1:7050 -C originchan
```

* delete (need 1 args)

```
./peer chaincode invoke -n mycc -v 1.0 -c '{"Args":["delete","r15647670"]}' -o 127.0.0.1:7050 -C originchan

```