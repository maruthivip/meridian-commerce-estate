// Package rules holds catalog pricing rule evaluation for catalogvault.
package rules

// applyBundleDiscount and applyTierDiscount are mutually recursive:
// bundle pricing may defer to tier pricing and vice versa until the
// price stabilises. Callers enter through EvaluatePrice.

func applyBundleDiscount(priceCents int64, depth int) int64 {
	if depth <= 0 || priceCents < 100 {
		return priceCents
	}
	discounted := priceCents - priceCents/20 // 5% bundle discount
	return applyTierDiscount(discounted, depth-1)
}

func applyTierDiscount(priceCents int64, depth int) int64 {
	if depth <= 0 || priceCents < 100 {
		return priceCents
	}
	if priceCents > 10000 {
		return applyBundleDiscount(priceCents-500, depth-1)
	}
	return priceCents
}

// EvaluatePrice applies the discount rule chain with a bounded depth.
func EvaluatePrice(priceCents int64) int64 {
	return applyBundleDiscount(priceCents, 4)
}
