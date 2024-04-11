package graph

import (
	"github.com/freightcms/common/models"
	"github.com/graphql-go/graphql"
)

// CurrencyCodeType is the GraphQL enum type for CurrencyCode
var CurrencyCodeGraphQLType = graphql.NewEnum(graphql.EnumConfig{
	Name:        "CurrencyCodeGraphQLtype",
	Description: "Stuff",
	Values: graphql.EnumValueConfigMap{
		"AFN": &graphql.EnumValueConfig{Value: models.AFN},
		"ALL": &graphql.EnumValueConfig{Value: models.ALL},
		"DZD": &graphql.EnumValueConfig{Value: models.DZD},
		"USD": &graphql.EnumValueConfig{Value: models.USD},
		"EUR": &graphql.EnumValueConfig{Value: models.EUR},
		"AOA": &graphql.EnumValueConfig{Value: models.AOA},
		"XCD": &graphql.EnumValueConfig{Value: models.XCD},
		"ARS": &graphql.EnumValueConfig{Value: models.ARS},
		"AMD": &graphql.EnumValueConfig{Value: models.AMD},
		"AWG": &graphql.EnumValueConfig{Value: models.AWG},
		"AUD": &graphql.EnumValueConfig{Value: models.AUD},
		"AZN": &graphql.EnumValueConfig{Value: models.AZN},
		"BSD": &graphql.EnumValueConfig{Value: models.BSD},
		"BHD": &graphql.EnumValueConfig{Value: models.BHD},
		"BDT": &graphql.EnumValueConfig{Value: models.BDT},
		"BBD": &graphql.EnumValueConfig{Value: models.BBD},
		"BYN": &graphql.EnumValueConfig{Value: models.BYN},
		"XOF": &graphql.EnumValueConfig{Value: models.XOF},
		"BMD": &graphql.EnumValueConfig{Value: models.BMD},
		"BTN": &graphql.EnumValueConfig{Value: models.BTN},
		"INR": &graphql.EnumValueConfig{Value: models.INR},
		"BOB": &graphql.EnumValueConfig{Value: models.BOB},
		"BOV": &graphql.EnumValueConfig{Value: models.BOV},
		"BAM": &graphql.EnumValueConfig{Value: models.BAM},
		"BWP": &graphql.EnumValueConfig{Value: models.BWP},
		"NOK": &graphql.EnumValueConfig{Value: models.NOK},
		"BRL": &graphql.EnumValueConfig{Value: models.BRL},
		"BND": &graphql.EnumValueConfig{Value: models.BND},
		"BGN": &graphql.EnumValueConfig{Value: models.BGN},
		"BIF": &graphql.EnumValueConfig{Value: models.BIF},
		"CVE": &graphql.EnumValueConfig{Value: models.CVE},
		"KHR": &graphql.EnumValueConfig{Value: models.KHR},
		"XAF": &graphql.EnumValueConfig{Value: models.XAF},
		"CAD": &graphql.EnumValueConfig{Value: models.CAD},
		"KYD": &graphql.EnumValueConfig{Value: models.KYD},
		"CLF": &graphql.EnumValueConfig{Value: models.CLF},
		"CLP": &graphql.EnumValueConfig{Value: models.CLP},
		"CNY": &graphql.EnumValueConfig{Value: models.CNY},
		"COP": &graphql.EnumValueConfig{Value: models.COP},
		"COU": &graphql.EnumValueConfig{Value: models.COU},
		"KMF": &graphql.EnumValueConfig{Value: models.KMF},
		"CDF": &graphql.EnumValueConfig{Value: models.CDF},
		"NZD": &graphql.EnumValueConfig{Value: models.NZD},
		"CRC": &graphql.EnumValueConfig{Value: models.CRC},
		"HRK": &graphql.EnumValueConfig{Value: models.HRK},
		"CUP": &graphql.EnumValueConfig{Value: models.CUP},
		"CUC": &graphql.EnumValueConfig{Value: models.CUC},
		"ANG": &graphql.EnumValueConfig{Value: models.ANG},
		"CZK": &graphql.EnumValueConfig{Value: models.CZK},
		"DKK": &graphql.EnumValueConfig{Value: models.DKK},
		"DJF": &graphql.EnumValueConfig{Value: models.DJF},
		"DOP": &graphql.EnumValueConfig{Value: models.DOP},
		"EGP": &graphql.EnumValueConfig{Value: models.EGP},
		"SVC": &graphql.EnumValueConfig{Value: models.SVC},
		"ERN": &graphql.EnumValueConfig{Value: models.ERN},
		"ETB": &graphql.EnumValueConfig{Value: models.ETB},
		"FKP": &graphql.EnumValueConfig{Value: models.FKP},
		"FJD": &graphql.EnumValueConfig{Value: models.FJD},
		"XPF": &graphql.EnumValueConfig{Value: models.XPF},
		"GMD": &graphql.EnumValueConfig{Value: models.GMD},
		"GEL": &graphql.EnumValueConfig{Value: models.GEL},
		"GHS": &graphql.EnumValueConfig{Value: models.GHS},
		"GIP": &graphql.EnumValueConfig{Value: models.GIP},
		"GTQ": &graphql.EnumValueConfig{Value: models.GTQ},
		"GBP": &graphql.EnumValueConfig{Value: models.GBP},
		"GNF": &graphql.EnumValueConfig{Value: models.GNF},
		"GYD": &graphql.EnumValueConfig{Value: models.GYD},
		"HTG": &graphql.EnumValueConfig{Value: models.HTG},
		"HNL": &graphql.EnumValueConfig{Value: models.HNL},
		"HKD": &graphql.EnumValueConfig{Value: models.HKD},
		"HUF": &graphql.EnumValueConfig{Value: models.HUF},
		"ISK": &graphql.EnumValueConfig{Value: models.ISK},
		"IDR": &graphql.EnumValueConfig{Value: models.IDR},
		"XDR": &graphql.EnumValueConfig{Value: models.XDR},
		"IRR": &graphql.EnumValueConfig{Value: models.IRR},
		"IQD": &graphql.EnumValueConfig{Value: models.IQD},
		"ILS": &graphql.EnumValueConfig{Value: models.ILS},
		"JMD": &graphql.EnumValueConfig{Value: models.JMD},
		"JPY": &graphql.EnumValueConfig{Value: models.JPY},
		"JOD": &graphql.EnumValueConfig{Value: models.JOD},
		"KZT": &graphql.EnumValueConfig{Value: models.KZT},
		"KES": &graphql.EnumValueConfig{Value: models.KES},
		"KPW": &graphql.EnumValueConfig{Value: models.KPW},
		"KRW": &graphql.EnumValueConfig{Value: models.KRW},
		"KWD": &graphql.EnumValueConfig{Value: models.KWD},
		"KGS": &graphql.EnumValueConfig{Value: models.KGS},
		"LAK": &graphql.EnumValueConfig{Value: models.LAK},
		"LBP": &graphql.EnumValueConfig{Value: models.LBP},
		"LSL": &graphql.EnumValueConfig{Value: models.LSL},
		"ZAR": &graphql.EnumValueConfig{Value: models.ZAR},
		"LRD": &graphql.EnumValueConfig{Value: models.LRD},
		"LYD": &graphql.EnumValueConfig{Value: models.LYD},
		"CHF": &graphql.EnumValueConfig{Value: models.CHF},
		"MOP": &graphql.EnumValueConfig{Value: models.MOP},
		"MKD": &graphql.EnumValueConfig{Value: models.MKD},
		"MGA": &graphql.EnumValueConfig{Value: models.MGA},
		"MWK": &graphql.EnumValueConfig{Value: models.MWK},
		"MYR": &graphql.EnumValueConfig{Value: models.MYR},
		"MVR": &graphql.EnumValueConfig{Value: models.MVR},
		"MRO": &graphql.EnumValueConfig{Value: models.MRO},
		"MUR": &graphql.EnumValueConfig{Value: models.MUR},
		"MXN": &graphql.EnumValueConfig{Value: models.MXN},
		"MXV": &graphql.EnumValueConfig{Value: models.MXV},
		"MDL": &graphql.EnumValueConfig{Value: models.MDL},
		"MNT": &graphql.EnumValueConfig{Value: models.MNT},
		"MAD": &graphql.EnumValueConfig{Value: models.MAD},
		"MZN": &graphql.EnumValueConfig{Value: models.MZN},
		"MMK": &graphql.EnumValueConfig{Value: models.MMK},
		"NAD": &graphql.EnumValueConfig{Value: models.NAD},
		"NPR": &graphql.EnumValueConfig{Value: models.NPR},
		"NIO": &graphql.EnumValueConfig{Value: models.NIO},
		"NGN": &graphql.EnumValueConfig{Value: models.NGN},
		"OMR": &graphql.EnumValueConfig{Value: models.OMR},
		"PKR": &graphql.EnumValueConfig{Value: models.PKR},
		"PAB": &graphql.EnumValueConfig{Value: models.PAB},
		"PGK": &graphql.EnumValueConfig{Value: models.PGK},
		"PYG": &graphql.EnumValueConfig{Value: models.PYG},
		"PEN": &graphql.EnumValueConfig{Value: models.PEN},
		"PHP": &graphql.EnumValueConfig{Value: models.PHP},
		"PLN": &graphql.EnumValueConfig{Value: models.PLN},
		"QAR": &graphql.EnumValueConfig{Value: models.QAR},
		"RON": &graphql.EnumValueConfig{Value: models.RON},
		"RUB": &graphql.EnumValueConfig{Value: models.RUB},
		"RWF": &graphql.EnumValueConfig{Value: models.RWF},
		"SHP": &graphql.EnumValueConfig{Value: models.SHP},
		"WST": &graphql.EnumValueConfig{Value: models.WST},
		"STD": &graphql.EnumValueConfig{Value: models.STD},
		"SAR": &graphql.EnumValueConfig{Value: models.SAR},
		"RSD": &graphql.EnumValueConfig{Value: models.RSD},
		"SCR": &graphql.EnumValueConfig{Value: models.SCR},
		"SLL": &graphql.EnumValueConfig{Value: models.SLL},
		"SGD": &graphql.EnumValueConfig{Value: models.SGD},
		"SBD": &graphql.EnumValueConfig{Value: models.SBD},
		"SOS": &graphql.EnumValueConfig{Value: models.SOS},
		"SSP": &graphql.EnumValueConfig{Value: models.SSP},
		"LKR": &graphql.EnumValueConfig{Value: models.LKR},
		"SDG": &graphql.EnumValueConfig{Value: models.SDG},
		"SRD": &graphql.EnumValueConfig{Value: models.SRD},
		"SZL": &graphql.EnumValueConfig{Value: models.SZL},
		"SEK": &graphql.EnumValueConfig{Value: models.SEK},
		"SYP": &graphql.EnumValueConfig{Value: models.SYP},
		"TWD": &graphql.EnumValueConfig{Value: models.TWD},
		"TJS": &graphql.EnumValueConfig{Value: models.TJS},
		"TZS": &graphql.EnumValueConfig{Value: models.TZS},
		"THB": &graphql.EnumValueConfig{Value: models.THB},
		"TOP": &graphql.EnumValueConfig{Value: models.TOP},
		"TTD": &graphql.EnumValueConfig{Value: models.TTD},
		"TND": &graphql.EnumValueConfig{Value: models.TND},
		"TRY": &graphql.EnumValueConfig{Value: models.TRY},
		"TMT": &graphql.EnumValueConfig{Value: models.TMT},
		"UGX": &graphql.EnumValueConfig{Value: models.UGX},
		"UAH": &graphql.EnumValueConfig{Value: models.UAH},
		"UYU": &graphql.EnumValueConfig{Value: models.UYU},
		"UYI": &graphql.EnumValueConfig{Value: models.UYI},
		"UZS": &graphql.EnumValueConfig{Value: models.UZS},
		"VUV": &graphql.EnumValueConfig{Value: models.VUV},
		"VEF": &graphql.EnumValueConfig{Value: models.VEF},
		"VND": &graphql.EnumValueConfig{Value: models.VND},
		"YER": &graphql.EnumValueConfig{Value: models.YER},
		"ZMW": &graphql.EnumValueConfig{Value: models.ZMW},
		"ZWL": &graphql.EnumValueConfig{Value: models.ZWL},
	},
})

var CurrencyType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Currency",
	Fields: graphql.Fields{
		"code": &graphql.Field{
			Name: "code",
			Type: CurrencyCodeGraphQLType,
		},
		"name": &graphql.Field{
			Name: "name",
			Type: graphql.String,
		},
		"symbol": &graphql.Field{
			Name: "symbol",
			Type: graphql.String,
		},
	},
})
