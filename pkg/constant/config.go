package constant

const (
	XAPIKEY    = "X-Api-Key"
	XCTO       = "X-Content-Type-Options"
	XCTO_VALUE = "nosniff"
	HSTS       = "Strict-Transport-Security"
	HSTS_VALUE = "max-age=31536000"
	ACAO       = "Access-Control-Allow-Origin"
	ACAO_VALUE = "*"
	ACAM       = "Access-Control-Allow-Methods"
	ACAM_VALUE = "GET, POST, OPTIONS"
	ACAH       = "Access-Control-Allow-Headers"
	ACAH_VALUE = "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Origin, Cookie, X-Appname, X-Api-Key, X-Personal-Number, X-Token, Signature, Grpc-Metadata-Signature, Timestamp, Grpc-Metadata-Timestamp, Grpc-Metadata-Client, Grpc-Metadata-Secret, Grpc-Metadata-Device, Client-Signature"
	ACAC       = "Access-Control-Allow-Credentials"
	CC         = "Cache-Control"
	CC_VALUE   = "no-store"
	ACAC_VALUE = "false"
)
