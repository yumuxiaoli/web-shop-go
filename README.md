# web-shop-go
golang微服务

```go

syntax = "proto3";

package user1;

option go_package = "./proto/user;user";

service User{
    rpc Register(UserRegisterRequest) returns (UserRegisterResponse){}
    rpc Login(UserLoginRequest) returns (UserLoginResponse){}
    rpc UserInfo(UserInfoRequest) returns (UserInfoResponse) {}
}

message UserInfoRequest{
    int64 user_id = 1;
}

message UserInfoResponse{
    string user_name = 1;
    string pwd = 2;
    string frist_name = 3; 
}

message UserRegisterRequest{
    string user_name = 1;
    string pwd =2;
    string frist_name = 3;
}

message UserRegisterResponse {
    string message = 1;
}

message UserLoginRequest {
    string user_name = 1;
    string pwd = 2;
}

message UserLoginResponse {
    bool is_success = 1;
}
```