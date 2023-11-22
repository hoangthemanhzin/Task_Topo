package ngap

import "etrib5gc/logctx"

var log logctx.LogWriter
var _logfields logctx.Fields = logctx.Fields{
	"mod": "ngap",
}

func _initLog() {
	if log == nil {
		log = logctx.WithFields(_logfields)
	}
}

func logSendingReport(msgname string, err error) {
	if err == nil {
		log.Infof("Message %s sent", msgname)
	} else {
		log.Errorf("Send %s failed: %s", err.Error())
	}
}
