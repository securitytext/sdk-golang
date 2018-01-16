/*
 * securitytextorgapi_lib
 *
 * This file was automatically generated for SecurityTextOrg by APIMATIC v2.0 ( https://apimatic.io )
 */

package whois_pkg

import "securitytextorgapi_lib/models_pkg"

/*
 * Interface for the WHOIS_IMPL
 */
type WHOIS interface {
    CreateApiWhoisQuery (*models_pkg.RequestsQueryModel) (*models_pkg.ResponsesQueryModel, error)
}

/*
 * Factory for the WHOIS interaface returning WHOIS_IMPL
 */
func NewWHOIS() WHOIS {
    return &WHOIS_IMPL{}
}