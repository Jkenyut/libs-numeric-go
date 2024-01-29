package helper

import (
	"encoding/json"
	"fmt"
	"github.com/Jkenyut/libs-numeric-go/libs_models/libs_model_validations"
	"github.com/go-playground/validator/v10"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
	"time"
)

// ConvertTime Convert Time
func ConvertTimeWIBFromString(date string, layout string) (res string, err error) {
	// if date empty or null
	if date == "" || strings.Contains(date, "0001-01-01T00:00:00") {
		return "", nil
	} else {
		var parsed time.Time
		parsed, err = time.Parse(time.RFC3339, date)
		if err != nil {
			return "", fmt.Errorf("failed to convert time format must format 0001-01-01T00:00:00")
		}
		currentDate := parsed.Format(layout)
		return currentDate, nil
	}
}

func ConvertNorekFormatBRI(in string) string {
	if len(in) < 15 {
		return ""
	}
	return fmt.Sprintf("%s-%s-%s-%s-%s", in[:4], in[4:6], in[6:12], in[12:14], in[14:])
}

func FormatAmountIDR(req float64) (res string) {
	//format idr like 100,000.00

	step := strings.Split(fmt.Sprintf("%.2f", req), ".")
	for i, j := range step[0] {
		if i > 0 && (len(step[0])-i)%3 == 0 { //logic
			res += ","
		}
		res += string(j)
	}
	res += "." + step[1] //join value + .00
	return res
}

// AddUnique duplicate array
func AddStringUnique(value string, slice *[]string) {
	//lopping data
	for _, eq := range *slice { //loop
		if strings.ToLower(eq) == strings.ToLower(value) || value == "" {
			//jika sudah ada data selesai
			return
		}
	}
	// jika tidak ada data masukkan
	*slice = append(*slice, cases.Title(language.Und, cases.NoLower).String(value))
}

func MessageValidate(err error, lang string) (message string) {
	ValidationMessages, errLoad := LoadLang() //load lang
	if errLoad != nil {
		return fmt.Sprint(errLoad)
	}

	var index = 1 //default indonesia
	for i := 0; i < len(ValidationMessages.Lang); i++ {
		if strings.ToLower(strings.TrimSpace(ValidationMessages.Lang[i].Language)) == strings.ToLower(strings.TrimSpace(lang)) { // lang conditional
			index = i
			break
		}
	}

	messages := ValidationMessages.Lang[index].Details //get value
	for _, messageError := range err.(validator.ValidationErrors) {
		//tagging validate
		switch strings.ToLower(messageError.Tag()) {
		case "required":
			message = fmt.Sprint(messageError.StructField(), messages.Required, messageError.Param())
		case "min":
			message = fmt.Sprint(messageError.StructField(), messages.Min, messageError.Param())
		case "max":
			message = fmt.Sprint(messageError.StructField(), messages.Max, messageError.Param())
		case "numeric":
			message = fmt.Sprint(messageError.StructField(), messages.Numeric, messageError.Param())
		case "ascii":
			message = fmt.Sprint(messageError.StructField(), messages.Ascii, messageError.Param())
		case "uuid":
			message = fmt.Sprint(messageError.StructField(), messages.UUID, messageError.Param())
		}
		break
	}
	message = cases.Title(language.Und, cases.NoLower).String(message)
	return message
}

// Function that uses the global variable
func LoadLang() (ValidationMessages *libs_model_validations.ValidationMessages, err error) {
	//validator10 **add your lang to here
	byteValue := libs_model_validations.ValidationMessages{
		Lang: []libs_model_validations.LanguageDetails{
			{
				Language: "eng",
				Details: libs_model_validations.ValidationDetails{
					Required: " ",
					Min:      " minimum character length ",
					Max:      " maximum character length ",
					Numeric:  " only allowed ",
					Ascii:    " ",
					UUID:     " ",
				},
			},
			{
				Language: "id",
				Details: libs_model_validations.ValidationDetails{
					Required: " ",
					Min:      " minimal panjang karakter ",
					Max:      " maksimal panjang karakter ",
					Numeric:  " hanya di bolehkan ",
					Ascii:    " ",
					UUID:     " ",
				},
			},
		},
	}
	//parse
	marshal, err := json.Marshal(byteValue)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON content into the struct
	err = json.Unmarshal(marshal, &ValidationMessages)
	if err != nil {
		return nil, err
	}
	return ValidationMessages, nil
}
