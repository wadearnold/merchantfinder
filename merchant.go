// merchantfinder/merchant.go
package merchantfinder

import (
	"regexp"
	"strings"
)

// MatchProcessor identifies a merchant acquiring processor from a description string
func MatchProcessor(description string) string {
	merchantAcquirers := map[string]*regexp.Regexp{
		"Square":                 regexp.MustCompile(`(?i)\bSQ\*|SQUARE|SQC\b`),
		"Stripe":                 regexp.MustCompile(`(?i)\bSTRIPE\b`),
		"Toast":                  regexp.MustCompile(`(?i)\bTOAST|TOASTTAB\b`),
		"PayPal":                 regexp.MustCompile(`(?i)\bPAYPAL\b`),
		"Shopify":                regexp.MustCompile(`(?i)\bSHOPIFY\b`),
		"Clover":                 regexp.MustCompile(`(?i)\bCLOVER\b`),
		"Adyen":                  regexp.MustCompile(`(?i)\bADYEN\b`),
		"Chase":                  regexp.MustCompile(`(?i)\bCHASE PAYMENTS|\bPMT\*`),
		"First Data":             regexp.MustCompile(`(?i)\bFIRST DATA\b`),
		"Worldpay":               regexp.MustCompile(`(?i)\bWORLDPAY|VANTIV\b`),
		"Elavon":                 regexp.MustCompile(`(?i)\bELAVON|US BANK MERCHANT\b`),
		"Global/TSYS":            regexp.MustCompile(`(?i)\bTSYS|GLOBAL PAYMENTS\b`),
		"Braintree":              regexp.MustCompile(`(?i)\bBRAINTREE|BT\*`),
		"WePay":                  regexp.MustCompile(`(?i)\bWEPAY\b`),
		"Cash App Biz":           regexp.MustCompile(`(?i)\bCASH APP\*|SQC\*`),
		"Mindbody":               regexp.MustCompile(`(?i)\bMINDBODY|MB\*`),
		"Lightspeed":             regexp.MustCompile(`(?i)\bLIGHTSPEED\b`),
		"Stax":                   regexp.MustCompile(`(?i)\bSTAX|FATTMERCHANT\b`),
		"Helcim":                 regexp.MustCompile(`(?i)\bHELCIM\b`),
		"Authorize.Net":          regexp.MustCompile(`(?i)\bAUTHNET|AUTHORIZENET\b`),
		"NMI":                    regexp.MustCompile(`(?i)\bNMI\b`),
		"Payrix":                 regexp.MustCompile(`(?i)\bPAYRIX\b`),
		"BlueSnap":               regexp.MustCompile(`(?i)\bBLUESNAP\b`),
		"Durango":                regexp.MustCompile(`(?i)\bDURANGO\b`),
		"ProPay":                 regexp.MustCompile(`(?i)\bPROPAY\b`),
		"FISERV":                 regexp.MustCompile(`(?i)\bFISERV\b`),
		"North American Bancard": regexp.MustCompile(`(?i)\bNORTH AMERICAN BANCARD\b`),
		"CDGcommerce":            regexp.MustCompile(`(?i)\bCDGCOMMERCE\b`),
		"Payment Depot":          regexp.MustCompile(`(?i)\bPAYMENT DEPOT\b`),
		"Payline Data":           regexp.MustCompile(`(?i)\bPAYLINE\b`),
		"QuickBooks Payments":    regexp.MustCompile(`(?i)\bINTUIT|QUICKBOOKS PAYMENTS\b`),
		"SquareSpace Commerce":   regexp.MustCompile(`(?i)\bSQUARESPACE\b`),
		"Wave Payments":          regexp.MustCompile(`(?i)\bWAVE\b`),
		"Gravity Payments":       regexp.MustCompile(`(?i)\bGRAVITY PAYMENTS\b`),
		"Paysafe":                regexp.MustCompile(`(?i)\bPAYSAFE\b`),
		"iPayment":               regexp.MustCompile(`(?i)\bIPAYMENT\b`),
		"Payoneer":               regexp.MustCompile(`(?i)\bPAYONEER\b`),
		"Dwolla":                 regexp.MustCompile(`(?i)\bDWOLLA\b`),
		"Finix":                  regexp.MustCompile(`(?i)\bFINIX\b`),
		"Infinicept":             regexp.MustCompile(`(?i)\bINFINICEPT\b`),
		"CheckOut.com":           regexp.MustCompile(`(?i)\bCHECKOUT\.COM\b`),
		"AltruPay":               regexp.MustCompile(`(?i)\bALTRUPAY\b`),
		"Nexio":                  regexp.MustCompile(`(?i)\bNEXIO\b`),
		"Corpay":                 regexp.MustCompile(`(?i)\bCORPAY\b`),
		"GoCardless":             regexp.MustCompile(`(?i)\bGOCARDLESS\b`),
		"Melio":                  regexp.MustCompile(`(?i)\bMELIO\b`),
		"Plastiq":                regexp.MustCompile(`(?i)\bPLASTIQ\b`),
		"Veem":                   regexp.MustCompile(`(?i)\bVEEM\b`),
		"Zelle Biz":              regexp.MustCompile(`(?i)\bZELLE\b`),
		"Bill.com":               regexp.MustCompile(`(?i)\bBILL\.COM\b`),
		"AffiniPay":              regexp.MustCompile(`(?i)\bAFFINIPAY\b`),
		"LawPay":                 regexp.MustCompile(`(?i)\bLAWPAY\b`),
		"CPACharge":              regexp.MustCompile(`(?i)\bCPACHARGE\b`),
		"PaySimple":              regexp.MustCompile(`(?i)\bPAYSIMPLE\b`),
		"RevverPay":              regexp.MustCompile(`(?i)\bREVVERPAY\b`),
		"PPRO":                   regexp.MustCompile(`(?i)\bPPRO\b`),
		"Fortis":                 regexp.MustCompile(`(?i)\bFORTIS\b`),
		"iTransact":              regexp.MustCompile(`(?i)\bITRANSACT\b`),
		"Fawry":                  regexp.MustCompile(`(?i)\bFAWRY\b`),
		"PayFast":                regexp.MustCompile(`(?i)\bPAYFAST\b`),
		"Amaryllis":              regexp.MustCompile(`(?i)\bAMARYLLIS\b`),
		"CORE":                   regexp.MustCompile(`(?i)\bCORE\b`),
		"Tilled":                 regexp.MustCompile(`(?i)\bTILLED\b`),
		"Paymob":                 regexp.MustCompile(`(?i)\bPAYMOB\b`),
		"BANKSapi":               regexp.MustCompile(`(?i)\bBANKSAPI\b`),
		"HIPS":                   regexp.MustCompile(`(?i)\bHIPS\b`),
		"Repay":                  regexp.MustCompile(`(?i)\bREPAY\b`),
		"CARDZ3N":                regexp.MustCompile(`(?i)\bCARDZ3N\b`),
		"Razorpay":               regexp.MustCompile(`(?i)\bRAZORPAY\b`),
		"Moov":                   regexp.MustCompile(`(?i)\bMOOV\b`),
		"BluePay":                regexp.MustCompile(`(?i)\bBLUEPAY\b`),
		"Exact Payments":         regexp.MustCompile(`(?i)\bEXACT PAYMENTS\b`),
		"Forward":                regexp.MustCompile(`(?i)\bFORWARD\b`),
		"Cardknox":               regexp.MustCompile(`(?i)\bCARDKNOX\b`),
		"Basis Theory":           regexp.MustCompile(`(?i)\bBASIS THEORY\b`),
		"IntelliPay":             regexp.MustCompile(`(?i)\bINTELLIPAY\b`),
		"PayJunction":            regexp.MustCompile(`(?i)\bPAYJUNCTION\b`),
		"Chargebee":              regexp.MustCompile(`(?i)\bCHARGEBEE\b`),
		"ChargeOver":             regexp.MustCompile(`(?i)\bCHARGEOVER\b`),
		"Recurly":                regexp.MustCompile(`(?i)\bRECURLY\b`),
		"Spreedly":               regexp.MustCompile(`(?i)\bSPREEDLY\b`),
	}

	desc := strings.TrimSpace(description)
	for name, pattern := range merchantAcquirers {
		if pattern.MatchString(desc) {
			return name
		}
	}
	return ""
}
