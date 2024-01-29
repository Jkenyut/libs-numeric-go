package helper

import (
	"encoding/json"
	"fmt"
	"github.com/Jkenyut/libs-numeric-go/libs_models/libs_model_validations"
	"github.com/go-playground/validator/v10"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"io"
	"os"
	"path"
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

func FormatAmountIDR(req float64) string {
	var res string
	step := strings.Split(fmt.Sprintf("%.2f", req), ".")
	for i, j := range step[0] {
		if i > 0 && (len(step[0])-i)%3 == 0 {
			res += ","
		}
		res += string(j)
	}
	res += "." + step[1]
	return res
}

// AddUnique duplicate array
func AddUnique(value string, slice *[]string) {
	//lopping data
	for _, eq := range *slice {
		if strings.ToLower(eq) == strings.ToLower(value) || value == "" {
			//jika sudah ada data selesai
			return
		}
	}
	// jika tidak ada data masukkan
	*slice = append(*slice, cases.Title(language.Und, cases.NoLower).String(value))
}

func RequestValidateID(err error, lang string) string {
	paths := path.Base("")
	fmt.Println(paths)
	err = LoadFileLang(path.Join(paths, "lang.json"))
	if err != nil {
		return fmt.Sprint(err)
	}
	fmt.Println(ValidationMessages, len(ValidationMessages.Lang))
	var index = 1
	for i := 0; i < len(ValidationMessages.Lang); i++ {
		if ValidationMessages.Lang[i].Language == lang {
			index = i
		}
	}

	messages := ValidationMessages.Lang[index].Details

	var message string
	for _, messageError := range err.(validator.ValidationErrors) {
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

var ValidationMessages *libs_model_validations.ValidationMessages

// Function that uses the global variable
func LoadFileLang(filename string) error {
	filePath := filename

	// Read the JSON content from the file
	jsonFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return err
	}

	// Unmarshal the JSON content into the struct

	err = json.Unmarshal(byteValue, &ValidationMessages)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return err
	}
	return nil
}
