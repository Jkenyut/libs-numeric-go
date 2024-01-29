package libs_model_validations

type ValidationMessages struct {
	Lang []LanguageDetails `json:"lang"`
}

type LanguageDetails struct {
	Language string            `json:"language"`
	Details  ValidationDetails `json:"details"`
}

type ValidationDetails struct {
	Required string `json:"required"`
	Min      string `json:"min"`
	Max      string `json:"max"`
	Numeric  string `json:"numeric"`
	Ascii    string `json:"ascii"`
	UUID     string `json:"uuid"`
}
