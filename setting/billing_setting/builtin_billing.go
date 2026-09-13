package billing_setting

// Built-in token prices use actual USD per million tokens. Keep new model
// defaults here instead of splitting them across the legacy ratio tables.
var builtinBillingExpr = map[string]string{
	// https://developers.openai.com/api/docs/pricing (Standard, 2026-09-09).
	// The Images API reports image output in output_tokens, normalized to c.
	"gpt-image-2":            `tier("standard", p * 5 + cr * 1.25 + img * 8 + img_cr * 2 + c * 30)`,
	"gpt-image-2.5-sunburst": `tier("standard", p * 5 + cr * 1.25 + img * 8 + img_cr * 2 + c * 30)`,
	"gpt-image-2.5-flare":    `tier("standard", p * 5 + cr * 1.25 + img * 8 + img_cr * 2 + c * 30)`,
	// https://developers.openai.com/api/docs/models/gpt-6-astra
	// Standard pricing; the long-context rates apply to the whole request.
	// Do not infer service-tier discounts from incoming request parameters:
	// channels filter service_tier by default, so it may not reach the upstream.
	"gpt-6-astra": `len <= 272000 ? tier("standard", p * 10 + c * 50 + cr * 1 + cc * 12.5) : tier("long_context", p * 20 + c * 75 + cr * 2 + cc * 25)`,
	// https://developers.openai.com/api/docs/pricing (Short context <=272K input tokens).
	// Long-context rates apply to the whole request once the input exceeds 272K.
	"gpt-5.6-sol":   `len <= 272000 ? tier("standard", p * 4 + c * 20 + cr * 0.4 + cc * 5) : tier("long_context", p * 8 + c * 30 + cr * 0.8 + cc * 10)`,
	"gpt-5.6-terra": `len <= 272000 ? tier("standard", p * 2 + c * 12 + cr * 0.2 + cc * 2.5) : tier("long_context", p * 4 + c * 18 + cr * 0.4 + cc * 5)`,
	"gpt-5.6-luna":  `len <= 272000 ? tier("standard", p * 0.2 + c * 1.2 + cr * 0.02 + cc * 0.25) : tier("long_context", p * 0.4 + c * 1.8 + cr * 0.04 + cc * 0.5)`,
}
