package ngapConvert

import (
	"etrib5gc/sbi/models"

	"github.com/free5gc/aper"
	"github.com/free5gc/ngap/ngapType"
)

// TS 38.413 9.3.1.85
func RATRestrictionInformationToNgap(ratType models.RatType) (ratResInfo ngapType.RATRestrictionInformation) {
	// Only support EUTRA & NR in version15.2.0
	switch ratType {
	case models.RATTYPE_EUTRA:
		ratResInfo.Value = aper.BitString{
			Bytes:     []byte{0x80},
			BitLength: 8,
		}
	case models.RATTYPE_NR:
		ratResInfo.Value = aper.BitString{
			Bytes:     []byte{0x40},
			BitLength: 8,
		}
	}
	return
}
