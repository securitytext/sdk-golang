/*
 * securitytextorgapi_lib
 *
 * This file was automatically generated for SecurityTextOrg by APIMATIC v2.0 ( https://apimatic.io )
 */

package models_pkg

import "time"

/*
 * Structure for the custom type AttributesUuidModel
 */
type AttributesUuidModel struct {
    Value           string          `json:"value" form:"value"` //Value of UUID
}

/*
 * Structure for the custom type RequestsQueryModel
 */
type RequestsQueryModel struct {
    Domain          string          `json:"domain" form:"domain"` //Name of Domain to query
}

/*
 * Structure for the custom type AttributesTimestampsModel
 */
type AttributesTimestampsModel struct {
    CreatedAt       *time.Time      `json:"created_at" form:"created_at"` //Creation date
    UpdatedAt       *time.Time      `json:"updated_at" form:"updated_at"` //Last updated
}

/*
 * Structure for the custom type ResponsesQueryModel
 */
type ResponsesQueryModel struct {
    Found           bool            `json:"found" form:"found"` //Query matched exactly one Domain object
    Status          int64           `json:"status" form:"status"` //Status of matched Domain object (if any)
    Output          string          `json:"output" form:"output"` //Text output of Query
}
