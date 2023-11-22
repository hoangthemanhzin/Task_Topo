package service

import "etrib5gc/logctx"

var log logctx.LogWriter

func _initLog() {
	if log == nil {
		log = logctx.WithFields(logctx.Fields{"mod": "service"})
	}
}
