# micro-logistics


## Generate
```sh
$ protoc -I communicate/ communicate/authenticate.proto --go_out=plugins=grpc:communicate
$ protoc -I communicate/ communicate/users.proto --go_out=plugins=grpc:communicate
```


## Keys google directions

KEY_GOOGLE="AIzaSyBPBrahfw9qmxMQAtTtDI54qBpjgF4I6wA"

protoc -I=communicate --go_out=. --go-grpc_out=. communicate/.proto protoc --go_out=. --proto_path=./communicate communicate/.proto --go-grpc_out=./communicate --go-grpc_opt=paths=source_relative communicate/slack.proto