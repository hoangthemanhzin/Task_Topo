package context

import "etrib5gc/logctx"

var log logctx.LogWriter
var _logfields logctx.Fields = logctx.Fields{
	"mod": "service",
}

func _initLog() {
	if log == nil {
		log = logctx.WithFields(_logfields)
	}
}
