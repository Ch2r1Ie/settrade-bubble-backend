package market_data

var stock_symbols_input = map[string]string{
	"DELTA":  "DELTA.BK",
	"ADVANC": "ADVANC.BK",
	"CPALL":  "CPALL.BK",
	"BDMS":   "BDMS.BK",
	"PTTEP":  "PTTEP.BK",
	"GULF":   "GULF.BK",
	"INTUCH": "INTUCH.BK",
	"KTB":    "KTB.BK",
	"PTT":    "PTT.BK",
	"BBL":    "BBL.BK",
	"BCP":    "BCP.BK",
	"AOT":    "AOT.BK",
	"IVL":    "IVL.BK",
	"CPN":    "CPN.BK",
	"TRUE":   "TRUE.BK",
	"BH":     "BH.BK",
	"SCC":    "SCC.BK",
	"BANPU":  "BANPU.BK",
	"TOP":    "TOP.BK",
	"BTS":    "BTS.BK",
	"PTTGC":  "PTTGC.BK",
	"SCB":    "SCB.BK",
	"VGI":    "VGI.BK",
	"GPSC":   "GPSC.BK",
	"MINT":   "MINT.BK",
	"EA":     "EA.BK",
	"TU":     "TU.BK",
	"CPF":    "CPF.BK",
	"WHA":    "WHA.BK",
	"KTC":    "KTC.BK",
	"HMPRO":  "HMPRO.BK",
	"CENTEL": "CENTEL.BK",
	"CRC":    "CRC.BK",
	"SKY":    "SKY.BK",
	"TISCO":  "TISCO.BK",
	"GLOBAL": "GLOBAL.BK",
	"KCE":    "KCE.BK",
	"BEM":    "BEM.BK",
	"HANA":   "HANA.BK",
	"OR":     "OR.BK",
	"TTB":    "TTB.BK",
	"SCGP":   "SCGP.BK",
	"SAWAD":  "SAWAD.BK",
	"LH":     "LH.BK",
	"RATCH":  "RATCH.BK",
	"ERW":    "ERW.BK",
	"BGRIM":  "BGRIM.BK",
	"COM7":   "COM7.BK",
	"JMT":    "JMT.BK",
	"AWC":    "AWC.BK",
	"BAM":    "BAM.BK",
	"BCH":    "BCH.BK",
	"TIDLOR": "TIDLOR.BK",
	"MTC":    "MTC.BK",
	"TLI":    "TLI.BK",
	"ICHI":   "ICHI.BK",
	"BJC":    "BJC.BK",
	"AP":     "AP.BK",
	"OSP":    "OSP.BK",
	"CBG":    "CBG.BK",
	"ITC":    "ITC.BK",
	"DOHOME": "DOHOME.BK",
	"SIRI":   "SIRI.BK",
	"BA":     "BA.BK",
	"SPALI":  "SPALI.BK",
	"TCAP":   "TCAP.BK",
	"AMATA":  "AMATA.BK",
	"AAV":    "AAV.BK",
	"KKP":    "KKP.BK",
	"TOA":    "TOA.BK",
	"EGCO":   "EGCO.BK",
	"TASCO":  "TASCO.BK",
	"MEGA":   "MEGA.BK",
	"STA":    "STA.BK",
	"PLANB":  "PLANB.BK",
	"CHG":    "CHG.BK",
	"PRM":    "PRM.BK",
	"CK":     "CK.BK",
	"MBK":    "MBK.BK",
	"SPRG":   "SPRC.BK",
	"STGT":   "STGT.BK",
	"SJWD":   "SJWD.BK",
	"BCPG":   "BCPG.BK",
	"RBF":    "RBF.BK",
	"JAS":    "JAS.BK",
	"SISB":   "SISB.BK",
	"IRPG":   "IRPC.BK",
	"M":      "M.BK",
	"RCL":    "RCL.BK",
	"QH":     "QH.BK",
	"BLA":    "BLA.BK",
	"JMART":  "JMART.BK",
	"SAPPE":  "SAPPE.BK",
	"GUNKUL": "GUNKUL.BK",
	"AEONTS": "AEONTS.BK",
	"CKP":    "CKP.BK",
	"TIPH":   "TIPH.BK",
	"BSRC":   "BSRC.BK",
	"BTG":    "BTG.BK",
}

var stock_symbols_output = map[string]string{
	"DELTA.BK":  "DELTA",
	"ADVANC.BK": "ADVANC",
	"CPALL.BK":  "CPALL",
	"BDMS.BK":   "BDMS",
	"PTTEP.BK":  "PTTEP",
	"GULF.BK":   "GULF",
	"INTUCH.BK": "INTUCH",
	"KTB.BK":    "KTB",
	"PTT.BK":    "PTT",
	"BBL.BK":    "BBL",
	"BCP.BK":    "BCP",
	"AOT.BK":    "AOT",
	"IVL.BK":    "IVL",
	"CPN.BK":    "CPN",
	"TRUE.BK":   "TRUE",
	"BH.BK":     "BH",
	"SCC.BK":    "SCC",
	"BANPU.BK":  "BANPU",
	"TOP.BK":    "TOP",
	"BTS.BK":    "BTS",
	"PTTGC.BK":  "PTTGC",
	"SCB.BK":    "SCB",
	"VGI.BK":    "VGI",
	"GPSC.BK":   "GPSC",
	"MINT.BK":   "MINT",
	"EA.BK":     "EA",
	"TU.BK":     "TU",
	"CPF.BK":    "CPF",
	"WHA.BK":    "WHA",
	"KTC.BK":    "KTC",
	"HMPRO.BK":  "HMPRO",
	"CENTEL.BK": "CENTEL",
	"CRC.BK":    "CRC",
	"SKY.BK":    "SKY",
	"TISCO.BK":  "TISCO",
	"GLOBAL.BK": "GLOBAL",
	"KCE.BK":    "KCE",
	"BEM.BK":    "BEM",
	"HANA.BK":   "HANA",
	"OR.BK":     "OR",
	"TTB.BK":    "TTB",
	"SCGP.BK":   "SCGP",
	"SAWAD.BK":  "SAWAD",
	"LH.BK":     "LH",
	"RATCH.BK":  "RATCH",
	"ERW.BK":    "ERW",
	"BGRIM.BK":  "BGRIM",
	"COM7.BK":   "COM7",
	"JMT.BK":    "JMT",
	"AWC.BK":    "AWC",
	"BAM.BK":    "BAM",
	"BCH.BK":    "BCH",
	"TIDLOR.BK": "TIDLOR",
	"MTC.BK":    "MTC",
	"TLI.BK":    "TLI",
	"ICHI.BK":   "ICHI",
	"BJC.BK":    "BJC",
	"AP.BK":     "AP",
	"OSP.BK":    "OSP",
	"CBG.BK":    "CBG",
	"ITC.BK":    "ITC",
	"DOHOME.BK": "DOHOME",
	"SIRI.BK":   "SIRI",
	"BA.BK":     "BA",
	"SPALI.BK":  "SPALI",
	"TCAP.BK":   "TCAP",
	"AMATA.BK":  "AMATA",
	"AAV.BK":    "AAV",
	"KKP.BK":    "KKP",
	"TOA.BK":    "TOA",
	"EGCO.BK":   "EGCO",
	"TASCO.BK":  "TASCO",
	"MEGA.BK":   "MEGA",
	"STA.BK":    "STA",
	"PLANB.BK":  "PLANB",
	"CHG.BK":    "CHG",
	"PRM.BK":    "PRM",
	"CK.BK":     "CK",
	"MBK.BK":    "MBK",
	"SPRG.BK":   "SPRC",
	"STGT.BK":   "STGT",
	"SJWD.BK":   "SJWD",
	"BCPG.BK":   "BCPG",
	"RBF.BK":    "RBF",
	"JAS.BK":    "JAS",
	"SISB.BK":   "SISB",
	"IRPG.BK":   "IRPC",
	"M.BK":      "M",
	"RCL.BK":    "RCL",
	"QH.BK":     "QH",
	"BLA.BK":    "BLA",
	"JMART.BK":  "JMART",
	"SAPPE.BK":  "SAPPE",
	"GUNKUL.BK": "GUNKUL",
	"AEONTS.BK": "AEONTS",
	"CKP.BK":    "CKP",
	"TIPH.BK":   "TIPH",
	"BSRC.BK":   "BSRC",
	"BTG.BK":    "BTG",
}
