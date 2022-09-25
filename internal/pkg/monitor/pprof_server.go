package monitor

import (
	"github.com/spf13/viper"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func PProfServer() {
	pprofAddr := ":" + viper.GetString("pporf.port")
	log.Fatal(http.ListenAndServe(pprofAddr, nil))
}
