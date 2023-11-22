package utils

import (
	"etrib5gc/sbi/models"
	"fmt"
)

func Param2String(v interface{}) string {
	//Dummy implementation
	return fmt.Sprintf("%v", v)
}

func String2Int32(s string) (v *int32, err error) {
	//TODO
	return
}

func String2Bool(s string) (v *bool, err error) {
	//TODO
	return
}

func String2ArrayOfstring(v string) (list []string, err error) {
	//TODO
	return
}

func String2ArrayOfRegistrationDataSetName(v string) (list []models.RegistrationDataSetName, err error) {
	//TODO
	return
}
func String2PlmnId(s string) (v *models.PlmnId, err error) {
	//TODO
	return
}
func String2HssAuthTypeInUri(s string) (v *models.HssAuthTypeInUri, err error) {
	//TODO
	return
}

func String2UpdateUeIdParameter(s string) (v *models.UpdateUeIdParameter, err error) {
	//TODO
	return
}

func String2UeContextInfoClass(s string) (v *models.UeContextInfoClass, err error) {
	//TODO
	return
}

func String2Guami(s string) (v *models.Guami, err error) {
	//TODO
	return
}

func String2ArrayOfNFType(s string) (nfs []models.NFType, err error) {
	//TODO
	return
}

func String2DataFilter(s string) (f *models.DataFilter, err error) {
	//TODO
	return
}

func String2ArrayOfContextDataSetName(s string) (list []models.ContextDataSetName, err error) {
	//TODO
	return
}

func String2ArrayOfDataSetName(s string) (list []models.DataSetName, err error) {
	//TODO
	return
}

func String2ArrayOfSnssai(s string) (list []models.Snssai, err error) {
	//TODO
	return
}

func String2Snssai(s string) (v *models.Snssai, err error) {
	//TODO
	return
}

func String2AppPortId(s string) (v *models.AppPortId, err error) {
	//TODO
	return
}

func String2ArrayOfIpv6Addr(s string) (list []models.Ipv6Addr, err error) {
	//TODO
	return
}
