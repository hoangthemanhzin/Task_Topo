package producer

import "etrib5gc/logctx"

var log logctx.LogWriter
var _logfields logctx.Fields = logctx.Fields{
	"mod": "context",
}

func _initLog() {
	if log == nil {
		log = logctx.WithFields(_logfields)
	}
}
