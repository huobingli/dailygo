## repo
github.com/darjun/go-daily-lib/twirp

## 简介
基于protobuf的rpc框架   

## 使用
``` cmd
$ go get github.com/twitchtv/twirp/protoc-gen-twirp
$ go get github.com/golang/protobuf/protoc-gen-go
```

增加proto文件 echo.proto
``` proto
syntax = "proto3";
option go_package = "proto";

service Echo {
  rpc Say(Request) returns (Response);
}

message Request {
  string text = 1;
}

message Response {
  string text = 2;
}
```

运行生成pb文件
``` cmd
$ protoc --twirp_out=. --go_out=. ./echo.proto
```
直接运行会报如下错误
``` cmd
The import path must contain at least one period ('.') or forward slash ('/') character. 
```
需要在上面go_package中增加'/' 或者 '.'