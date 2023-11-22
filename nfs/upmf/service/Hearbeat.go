package service

func SendHB() {
	// node.upfcli= httpdp.NewClientWithAddr(nf.upfaddr.String()) ==> sbi/upf.go => register

	// node = upmf.prod
	// for (nodes)

	// if hbrsp, err := upf2upmf.Heartbeat(upfcli, msg); err != nil {
	// 	prod.Errorf("Send Heartbeat to %s failed: %+v", upfaddr, err)
	// } else {
	// 	prod.Infof("Nonce=%d, Time=%s", hbrsp.Nonce, hbrsp.Msg.RecoveryTimeStamp.RecoveryTimeStamp.String())
	// }
}

// upfaddr: ip + port ==> for creating upfcli
// func (addr *UpfAddress) String() string {
// 	return fmt.Sprintf("%s:%d", addr.Ip.String(), addr.Port)
// }


// lap lai request ==> heartbeatloop
// for {
// 	//do registering here
// 	if err := nf.heartbeat(); err == nil {
// 		return
// 	} else {
// 		log.Errorf("Registration error: %+v", err)
// 	}
// 	t := time.NewTimer(5 * time.Second)
// 	select {
// 	case <-t.C:
// 		continue
// 	case <-nf.quit:
// 		return
// 	}
// }