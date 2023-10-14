package router

// This Constant Variable Status Code For HTTP Response
// Follow This Opensource https://github.com/waldemarnt/http-status-codes

const (
	// HTTP STATUS CODE
	StatusOK         = 200
	StatusCreated    = 201
	StatusAccepted   = 202
	StatusNoResponse = 203

	StatusMoved = 301
	StatusFound = 302

	StatusBadRequest       = 400
	StatusUnauthorized     = 401
	StatusPaymentRequired  = 402
	StatusForbidden        = 403
	StatusNotFound         = 404
	StatusMethodNotAllowed = 405

	StatusInternalServerError     = 500
	StatusServiceOverloaded       = 502
	StatusGatewayTimeout          = 503
	StatusHttpVersionNotSupported = 505
)
