package topo

import "etrib5gc/logctx"

var log logctx.LogWriter
var _logfields logctx.Fields = logctx.Fields{
	"mod": "topo",
}

func _initLog() {
	if log == nil {
		log = logctx.WithFields(_logfields)
	}
}
