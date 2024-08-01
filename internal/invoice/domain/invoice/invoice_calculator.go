package invoice

// 手数料に対する消費税率
const FeeConsumptionTaxRate = 0.1

func CalculateBillingAmount(paymentAmount float64, feeRate float64) float64 {
	return paymentAmount + (paymentAmount * feeRate / 100 * (1 + FeeConsumptionTaxRate))
}
