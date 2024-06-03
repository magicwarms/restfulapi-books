package utils

import (
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

// ValidateFields is for response validation error
func ValidateFields(model interface{}) *ErrorDetail {
	var validate = validator.New()
	err := validate.Struct(model)
	if err != nil {
		validationErr := err.(validator.ValidationErrors)[0]

		// 	// fmt.Println(err.Namespace())
		// 	// fmt.Println(err.Field())
		// 	// fmt.Println(err.StructNamespace())
		// 	// fmt.Println(err.StructField())
		// 	// fmt.Println(err.Tag())
		// 	// fmt.Println(err.ActualTag())
		// 	// fmt.Println(err.Kind())
		// 	// fmt.Println(err.Type())
		// 	// fmt.Println(err.Value().(string))
		// 	// fmt.Println(err.Param())
		// 	// fmt.Println()

		return &ErrorDetail{
			Code:    http.StatusUnprocessableEntity,
			Message: messageForTag(validationErr.Tag(), validationErr.Param(), validationErr.Field()),
		}
	}

	return nil
}

func messageForTag(tag, value, field string) string {

	switch tag {
	case "required":
		return field + " wajib dimasukkan"
	case "min":
		return field + " minimal " + value + " karakter"
	case "max":
		return field + " maksimal " + value + " karakter"
	case "lowercase":
		return field + " harus kecil semua"
	case "e164":
		return "Format " + field + " tidak valid"
	case "uuid4":
		return field + " UUID tidak valid"
	case "latitude":
		return field + value + " tidak valid"
	case "longitude":
		return field + value + " tidak valid"
	case "numeric":
		return field + value + " harus berupa angka"
	case "alpha":
		return field + value + " harus berupa huruf saja dan tidak ada spasi"
	case "alphanum":
		return field + " harus berupa angka dan huruf"
	case "datetime":
		return field + " format salah, mohon ikuti format: 2006-11-14"
	case "number":
		return field + " tipe data salah, harus bertipe angka"
	}

	// default error
	return strings.ToUpper("WARNING!! pesan validasi belum di convert. type: " + tag)
}
