package context

import (
	"etrib5gc/common"
	"etrib5gc/mesh"
	meshmodels "etrib5gc/mesh/models"
	"etrib5gc/nfs/ausf/config"
	"etrib5gc/sbi/models"
)

type AusfContext struct {
	group  string
	plmnid models.PlmnId

	uelist UeList //ue pool indexed with supi
	idlist IdList //suciorsupi to supi (for searching supi)
}

func New(cfg *config.AusfConfig) *AusfContext {
	_initLog()
	ret := &AusfContext{
		group:  cfg.Group,
		plmnid: models.PlmnId(cfg.PlmnId),

		idlist: newIdList(),
		uelist: newUeList(),
	}
	return ret
}

func (c *AusfContext) AddUeContext(ue *UeContext, info *models.AuthenticationInfoResult) (err error) {
	if err = ue.update(info); err == nil {
		c.uelist.add(ue)
		c.idlist.add(ue.SupiOrSuci(), info.Supi)
	}
	return
}

func (c *AusfContext) GetSupi(suciORsupi string) (supi string) {
	supi = c.idlist.get(suciORsupi)
	return
}

func (c *AusfContext) NewUeContext(ueid string, snname string) (ue *UeContext, err error) {
	ue = newUeContext(ueid, snname)
	sid := common.UdmServiceName(&c.plmnid)
	ue.udmcli, err = mesh.Consumer(meshmodels.ServiceName(sid), nil, false)
	return
}
func (c *AusfContext) GetUeContext(suciORsupi string) (ue *UeContext) {
	if supi := c.idlist.get(suciORsupi); len(supi) > 0 {
		ue = c.uelist.get(supi)
	}
	return
}

func (c *AusfContext) Url() string {
	return mesh.CallbackAddress()
}
func (c *AusfContext) IsNetworkAuthorized(netname string) bool {
	return true
}
func (c *AusfContext) PlmnId() *models.PlmnId {
	return &c.plmnid
}
