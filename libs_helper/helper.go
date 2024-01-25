package helper

import (
	"fmt"
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
