// Mapping of error codes.
package endpoints

// Provides functionality to convert HTTP error codes from the SPI to
// match the errors that will be returned by the server.
// TODO: Generate from /google/appengine/tools/devappserver2/endpoints/generated_error_info.py


type ErrorInfo struct {
	http_status, rpc_status int
	reason, domain string
}

var UNSUPPORTED_ERROR = &ErrorInfo{404, 404, "unsupportedProtocol", "global"}
var BACKEND_ERROR = &ErrorInfo{503, -32099, "backendError", "global"}

var ERROR_MAP = map[int]*ErrorInfo {
	400: &ErrorInfo{400, 400, "badRequest", "global"},
	401: &ErrorInfo{401, 401, "required", "global"},
	402: &ErrorInfo{404, 404, "unsupportedProtocol", "global"},
	403: &ErrorInfo{403, 403, "forbidden", "global"},
	404: &ErrorInfo{404, 404, "notFound", "global"},
	405: &ErrorInfo{501, 501, "unsupportedMethod", "global"},
	406: &ErrorInfo{404, 404, "unsupportedProtocol", "global"},
	407: &ErrorInfo{404, 404, "unsupportedProtocol", "global"},
	408: &ErrorInfo{503, -32099, "backendError", "global"},
	409: &ErrorInfo{409, 409, "duplicate", "global"},
	410: &ErrorInfo{410, 410, "deleted", "global"},
	411: &ErrorInfo{404, 404, "unsupportedProtocol", "global"},
	412: &ErrorInfo{412, 412, "conditionNotMet", "global"},
	413: &ErrorInfo{413, 413, "uploadTooLarge", "global"},
	414: &ErrorInfo{404, 404, "unsupportedProtocol", "global"},
	415: &ErrorInfo{404, 404, "unsupportedProtocol", "global"},
	416: &ErrorInfo{404, 404, "unsupportedProtocol", "global"},
	417: &ErrorInfo{404, 404, "unsupportedProtocol", "global"},
}

// Get info that would be returned by the server for this HTTP status.
//
// Args:
//   lily_status: An integer containing the HTTP status returned by the SPI.
//
// Returns:
// An _ErrorInfo object containing information that would be returned by the
// live server for the provided lily_status.
func get_error_info(lily_status int) *ErrorInfo {
	if lily_status >= 500 {
		return BACKEND_ERROR
	}

	error_info, ok := ERROR_MAP[lily_status]
	if !ok {
		return UNSUPPORTED_ERROR
	}
	return error_info
}