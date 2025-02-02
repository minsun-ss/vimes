package main

import (
	"fmt"
	"log/slog"
	"vimes/internal/webserver"
)

func main() {
	slog.Info("Starting webserver...", "port", 24680)
	webserver.Webserver(":24680")
	fmt.Println("Whoa")
}
