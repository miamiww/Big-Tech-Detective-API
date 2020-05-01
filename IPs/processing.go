package IPs

import (
	"net/http"
)

// FormToIP -- fills a IP struct with submitted form data
// params:
// r - request reader to fetch form data or url params (unused here)
// returns:
// IP struct if successful
// array of strings of errors if any occur during processing
func FormToIP(r *http.Request) (IP, []string) {
	var ip IP
	var errStr string
	var errs []string
	
	ip.IPV4, errStr = processFormField(r, "ipv4")
	errs = appendError(errs, errStr)
	ip.Company, errStr = processFormField(r, "company")
	errs = appendError(errs, errStr)

	return ip, errs
}

func appendError(errs []string, errStr string) ([]string) {
	if len(errStr) > 0 {
		errs = append(errs, errStr)
	}
	return errs
}

func processFormField(r *http.Request, field string) (string, string) {
	fieldData := r.PostFormValue(field)
	if len(fieldData) == 0 {
		return "", "Missing '" + field + "' parameter, cannot continue"
	}
	return fieldData, ""
}
