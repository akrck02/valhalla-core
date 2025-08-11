package apimodels

type HTTPMethod int

const (
	GetMethod HTTPMethod = iota
	PostMethod
	PutMethod
	DeleteMethod
	PatchMethod
	OptionsMethod
	HeadMethod
	TraceMethod
)

type HTTPHeaders string

const (
	HeaderContentType             HTTPHeaders = "Content-Type"
	HeaderAccept                  HTTPHeaders = "Accept"
	HeaderAuthorization           HTTPHeaders = "Authorization"
	HeaderUserAgent               HTTPHeaders = "User-Agent"
	HeaderContentLength           HTTPHeaders = "Content-Length"
	HeaderContentEncoding         HTTPHeaders = "Content-Encoding"
	HeaderContentDisposition      HTTPHeaders = "Content-Disposition"
	HeaderContentTransferEncoding HTTPHeaders = "Content-Transfer-Encoding"
	HeaderContentLanguage         HTTPHeaders = "Content-Language"
)

type MimeType string

const (
	MimeApplicationJSON        MimeType = "application/json"
	MimeApplicationXML         MimeType = "application/xml"
	MimeApplicationYaml        MimeType = "application/yaml"
	MimeApplicationForm        MimeType = "application/x-www-form-urlencoded"
	MimeApplicationOctetStream MimeType = "application/octet-stream"
	MimeTextPlain              MimeType = "text/plain"
	MimeTextHTML               MimeType = "text/html"
)
