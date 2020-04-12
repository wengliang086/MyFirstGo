# 说明

[官方文档](https://github.com/golang/protobuf)

## go 中使用 protocol buffer

```
protoc --go_out=. *.proto
```



## gPPC支持

```
protoc --go_out=plugins=grpc:. *.proto
```

* 指定文件

  `protoc --go_out=plugins=grpc:./ ./test.proto`

* 生成内容对比

![image-20200316115455732](readme.assets/image-20200316115455732.png)



## [枚举类型生成多Type前缀](https://github.com/golang/protobuf/issues/513)

## 解决方式

[gogoExtension](https://github.com/gogo/protobuf/blob/master/extensions.md#goprotobuf-compatibility)

![image-20200316123854179](/Users/phoenix/go/src/goModTest/tests/1grpc/pb/readme.assets/image-20200316123854179.png)



## gogo 使用

```
protoc --gofast_out=. myproto.proto
```

## gogo GRPC

```
protoc --gofast_out=plugins=grpc:. my.proto
```

## 完整示例

![image-20200316143414702](readme.assets/image-20200316143414702.png)

