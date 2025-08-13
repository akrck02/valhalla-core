package apimodels

import (
	"fmt"
	"mime/multipart"
	"strconv"
	"strings"

	verrors "github.com/akrck02/valhalla-core/sdk/errors"
)

type Request struct {
	Authorization string                             `json:"authorization,omitempty"`
	UserID        *int64                             `json:"userID,omitempty"`
	IP            string                             `json:"ip,omitempty"`
	UserAgent     string                             `json:"userAgent,omitempty"`
	Params        map[string]string                  `json:"params,omitempty"`
	Body          any                                `json:"body,omitempty"`
	Headers       map[string]string                  `json:"headers,omitempty"`
	Files         map[string][]*multipart.FileHeader `json:"files,omitempty"`
}

func (r Request) GetParamInt64(name string) (*int64, error) {
	paramStr := r.Params[name]
	if strings.TrimSpace(paramStr) == "" {
		return nil, fmt.Errorf("parameter %s is empty", name)
	}

	param, err := strconv.ParseInt(paramStr, 10, 64)
	if nil != err {
		return nil, err
	}

	return &param, nil
}

func (r Request) GetFile(name string, maxFileSize int64) (*multipart.FileHeader, error) {
	if nil == r.Files[name] || len(r.Files[name]) == 0 {
		return nil, nil
	}

	file := r.Files[name][0]
	if file.Size > maxFileSize {
		return nil, fmt.Errorf(verrors.FileTooLargeMessage, name, maxFileSize/1024/1024)
	}

	return file, nil
}
