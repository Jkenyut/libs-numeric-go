package helper

import (
	"github.com/go-playground/validator/v10"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestConvertTimeWIBFromString(t *testing.T) {
	Convey("Positive Helper ConvertTimeWIBFromString", t, func() {
		convertTime, err := ConvertTimeWIBFromString("2006-01-31T15:04:05-07:00", "02/01/2006 15:04:05")
		So(convertTime, ShouldEqual, "31/01/2006 15:04:05")
		So(err, ShouldBeNil)
	})
	Convey("Positive Helper ConvertTimeWIBFromString Contains", t, func() {
		convertTime, err := ConvertTimeWIBFromString("0001-01-01T00:00:00", "02/01/2006 15:04:05")
		So(convertTime, ShouldEqual, "")
		So(err, ShouldBeNil)
	})
	Convey("Positive Helper ConvertTimeWIBFromString Empty", t, func() {
		convertTime, err := ConvertTimeWIBFromString("", "02/01/2006 15:04:05")
		So(convertTime, ShouldEqual, "")
		So(err, ShouldBeNil)
	})
	Convey("Negative Helper ConvertTimeWIBFromString", t, func() {
		convertTime, err := ConvertTimeWIBFromString("1902/1999", "02/01/2006 15:04:05")
		So(convertTime, ShouldBeEmpty)
		So(err, ShouldNotBeNil)
	})
}

func TestFunction(t *testing.T) {

	Convey("Convert Norek", t, func() {
		So(ConvertNorekFormatBRI("348601006415103"), ShouldEqual, "3486-01-006415-10-3")
	})
	Convey("Convert Norek Empty", t, func() {
		So(ConvertNorekFormatBRI(""), ShouldEqual, "")
	})
	Convey("Convert Amount", t, func() {
		So(FormatAmountIDR(7), ShouldEqual, "7.00")
		So(FormatAmountIDR(70), ShouldEqual, "70.00")
		So(FormatAmountIDR(700), ShouldEqual, "700.00")
		So(FormatAmountIDR(7000), ShouldEqual, "7,000.00")
		So(FormatAmountIDR(70000), ShouldEqual, "70,000.00")
		So(FormatAmountIDR(700000), ShouldEqual, "700,000.00")
		So(FormatAmountIDR(7000000), ShouldEqual, "7,000,000.00")
		So(FormatAmountIDR(70000000), ShouldEqual, "70,000,000.00")
		So(FormatAmountIDR(700000000), ShouldEqual, "700,000,000.00")
		So(FormatAmountIDR(7000000000), ShouldEqual, "7,000,000,000.00")
		So(FormatAmountIDR(70000000000), ShouldEqual, "70,000,000,000.00")
		So(FormatAmountIDR(700000000000), ShouldEqual, "700,000,000,000.00")
		So(FormatAmountIDR(7000000000000), ShouldEqual, "7,000,000,000,000.00")
		So(FormatAmountIDR(70000000000000), ShouldEqual, "70,000,000,000,000.00")
		So(FormatAmountIDR(700000000000000), ShouldEqual, "700,000,000,000,000.00")
		So(FormatAmountIDR(7000000000000000), ShouldEqual, "7,000,000,000,000,000.00")
		So(FormatAmountIDR(70000000000000000), ShouldEqual, "70,000,000,000,000,000.00")
	})

	Convey("Add Unique Value", t, func() {
		var s []string
		Array := []string{"Berhasil", "Gagal", "Kalo", "Ini", "ini", "Atau", "Jika", ""}
		for _, i := range Array {
			AddUnique(i, &s)
		}
		So(len(s), ShouldEqual, 6)
	})

}

func TestRequestValidate(t *testing.T) {
	Convey("Given a list of validation errors", t, func() {
		type Required struct {
			ID string `json:"ID,omitempty" validate:"required"`
		}
		type Min struct {
			ID string `json:"ID,omitempty" validate:"min=2"`
		}
		type Max struct {
			ID string `json:"ID,omitempty" validate:"max=2"`
		}
		type Numeric struct {
			ID string `json:"ID,omitempty" validate:"numeric=2"`
		}
		type ascii struct {
			ID string `json:"ID" validate:"ascii"`
		}
		type uuid struct {
			ID string `json:"ID" validate:"uuid"`
		}
		// Create a validation error object.
		valid := validator.New()

		Convey("When RequestValidate is called", func() {
			Convey("Then it should return the expected error message required", func() {
				e := valid.Struct(Required{ID: ""})
				result := RequestValidateID(e)
				So(result, ShouldNotBeNil)
			})

			Convey("Then it should return the expected error message min", func() {
				e := valid.Struct(Min{ID: "c"})
				result := RequestValidateID(e)
				So(result, ShouldNotBeNil)
			})
			Convey("Then it should return the expected error message max", func() {
				e := valid.Struct(Max{ID: "eneneiovneic"})
				result := RequestValidateID(e)
				So(result, ShouldNotBeNil)
			})
			Convey("Then it should return the expected error message Numeric", func() {
				e := valid.Struct(Numeric{ID: "eneneiovneic"})
				result := RequestValidateID(e)
				So(result, ShouldNotBeNil)
			})
			Convey("Then it should return the expected error message ascii", func() {
				e := valid.Struct(ascii{ID: "é"})
				result := RequestValidateID(e)
				So(result, ShouldNotBeNil)
			})
			Convey("Then it should return the expected error message uuid", func() {
				e := valid.Struct(uuid{ID: "28febdfd-2b58-4a03-ac1e-e541a07bfac"})
				result := RequestValidateID(e)
				So(result, ShouldNotBeNil)
			})
		})
	})
}
