package main

import (
	"net/http"

	now "github.com/devgar/now-go/bridge"
)

func main() {
	now.Start(http.HandlerFunc(__NOW_HANDLER_FUNC_NAME))
}
