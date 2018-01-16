/*
 * securitytextorgapi_lib
 *
 * This file was automatically generated for SecurityTextOrg by APIMATIC v2.0 ( https://apimatic.io )
 */
package whois_pkg


import(
	"errors"
	"encoding/json"
	"securitytextorgapi_lib/models_pkg"
	"github.com/apimatic/unirest-go"
	"securitytextorgapi_lib"
	"securitytextorgapi_lib/apihelper_pkg"
)
/*
 * Client structure as interface implementation
 */
type WHOIS_IMPL struct { }

/**
 * Query the server for a Domain object
 * @param    *models_pkg.RequestsQueryModel        body     parameter: Required
 * @return	Returns the *models_pkg.ResponsesQueryModel response from the API call
 */
func (me *WHOIS_IMPL) CreateApiWhoisQuery (
            body *models_pkg.RequestsQueryModel) (*models_pkg.ResponsesQueryModel, error) {
    //validating required parameters
    if (body == nil){
        return nil,errors.New("The parameter 'body' is a required parameter and cannot be nil.")
}     //the base uri for api requests
    _queryBuilder := securitytextorgapi_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/query"

    //variable to hold errors
    var err error = nil
    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "securitytextorg-sdk/v1",
        "accept" : "application/json",
        "content-type" : "application/json; charset=utf-8",
    }

    //prepare API request
    _request := unirest.Post(_queryBuilder, headers, body)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 400) {
        err = apihelper_pkg.NewAPIError("Bad Request", _response.Code, _response.RawBody)
    } else if (_response.Code == 406) {
        err = apihelper_pkg.NewAPIError("Not Acceptable", _response.Code, _response.RawBody)
    } else if (_response.Code == 500) {
        err = apihelper_pkg.NewAPIError("Internal Server Error", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper_pkg.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
        }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal *models_pkg.ResponsesQueryModel = &models_pkg.ResponsesQueryModel{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil
}

