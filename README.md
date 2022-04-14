# micro-logistics


## Generate
```sh
$ protoc -I communicate/ communicate/logistic.proto --go_out=plugins=grpc:communicate
$ protoc communicate/client.proto --go_out=plugins=grpc:.
```


## Keys google directions

protoc -I=communicate --go_out=. --go-grpc_out=. communicate/.proto protoc --go_out=. --proto_path=./communicate communicate/.proto --go-grpc_out=./communicate --go-grpc_opt=paths=source_relative communicate/slack.proto