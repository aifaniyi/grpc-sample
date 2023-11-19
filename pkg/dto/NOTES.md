# grpc service method kinds

- Unary:
```
rpc SayHello(HelloRequest) returns (HelloResponse);
```

- Server streaming
```
rpc LotsOfReplies(HelloRequest) returns (stream HelloResponse);
```

- Client streaming
```
rpc LotsOfGreetings(stream HelloRequest) returns (HelloResponse);
```

- Bi-directional
```
rpc BidiHello(stream HelloRequest) returns (stream HelloResponse);
```