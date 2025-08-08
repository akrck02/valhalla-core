package apimodels

type HttpMethod int

const (
	GetMethod HttpMethod = iota
	PostMethod
	PutMethod
	DeleteMethod
	PatchMethod
	OptionsMethod
	HeadMethod
	TraceMethod
)

type HttpHeaders string

const (
	HeaderContentType             HttpHeaders = "Content-Type"
	HeaderAccept                  HttpHeaders = "Accept"
	HeaderAuthorization           HttpHeaders = "Authorization"
	HeaderUserAgent               HttpHeaders = "User-Agent"
	HeaderContentLength           HttpHeaders = "Content-Length"
	HeaderContentEncoding         HttpHeaders = "Content-Encoding"
	HeaderContentDisposition      HttpHeaders = "Content-Disposition"
	HeaderContentTransferEncoding HttpHeaders = "Content-Transfer-Encoding"
	HeaderContentLanguage         HttpHeaders = "Content-Language"
)

type MimeType string

const (
	MimeApplicationJson        MimeType = "application/json"
	MimeApplicationXml         MimeType = "application/xml"
	MimeApplicationYaml        MimeType = "application/yaml"
	MimeApplicationForm        MimeType = "application/x-www-form-urlencoded"
	MimeApplicationOctetStream MimeType = "application/octet-stream"
	MimeTextPlain              MimeType = "text/plain"
	MimeTextHtml               MimeType = "text/html"
)
