package merchantfinder

import "testing"

func TestMatchProcessor(t *testing.T) {
	tests := []struct {
		description string
		expected    string
	}{
		{"SQ*Joe's Coffee", "Square"},
		{"STRIPE*PAYMENT", "Stripe"},
		{"TOASTTAB*BurgerBar", "Toast"},
		{"PayPal * Donation", "PayPal"},
		{"SHOPIFY*Order123", "Shopify"},
		{"CLOVER POS", "Clover"},
		{"ADYEN N.V.", "Adyen"},
		{"CHASE PAYMENTS SETTLEMENT", "Chase"},
		{"FIRST DATA CORP", "First Data"},
		{"WORLDPAY PAYMENT", "Worldpay"},
		{"ELAVON PAYMENT", "Elavon"},
		{"GLOBAL PAYMENTS INC", "Global/TSYS"},
		{"BRAINTREE PMT", "Braintree"},
		{"WEPAY*INVOICE", "WePay"},
		{"CASH APP*SQC123", "Cash App Biz"},
		{"MINDBODY INC", "Mindbody"},
		{"LIGHTSPEED RETAIL", "Lightspeed"},
		{"FATTMERCHANT PMT", "Stax"},
		{"HELCIM DEPOSIT", "Helcim"},
		{"AUTHNET DEPOSIT", "Authorize.Net"},
		{"NMI GATEWAY", "NMI"},
		{"Unrelated Description", ""},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			result := MatchProcessor(tt.description)
			if result != tt.expected {
				t.Errorf("Expected %q but got %q", tt.expected, result)
			}
		})
	}
}
