package main

import (
	"github.com/sillyhatxu/microservice-template/common/logs"
	"github.com/sillyhatxu/microservice-template/demo-backend/api"
)

func init() {
	logs.InitialLogConfig(logs.Env("dev"), logs.Project("microservice"), logs.Module("demo-backend"), logs.Version("v1.0.0-beta.1"))
}

func main() {
	api.InitialAPI(8080)
}
