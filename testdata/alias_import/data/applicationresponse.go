package data

import (
	typesapplication "github.com/5999ft/swag/v2/testdata/alias_import/types"
)

type ApplicationResponse struct {
	typesapplication.TypeToEmbed

	Application      typesapplication.Application   `json:"application"`
	ApplicationArray []typesapplication.Application `json:"application_array"`
	ApplicationTime  typesapplication.DateOnly      `json:"application_time"`
}
