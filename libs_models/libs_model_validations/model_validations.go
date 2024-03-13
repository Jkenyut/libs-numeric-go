package libs_model_validations

// ValidationMessages struct is connected to LanguageDetails and its details to ValidationDetails
type ValidationMessages struct {
	Lang []LanguageDetails `json:"lang"`
}

type LanguageDetails struct {
	Language string            `json:"language"`
	Details  ValidationDetails `json:"details"`
}

// ValidationDetails fill in all validator 10 fields of this type
type ValidationDetails struct {
	Required string `json:"required"`
	Min      string `json:"min"`
	Max      string `json:"max"`
	Numeric  string `json:"numeric"`
	Ascii    string `json:"ascii"`
	UUID     string `json:"uuid"`
}
