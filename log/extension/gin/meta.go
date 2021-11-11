package gin

import "github.com/illidaris/apocalypse/pkg/consts"

type HttpMetaData consts.MetaData

const (
	HTTPStatusCode  HttpMetaData = "statusCode"
	HTTPContentType HttpMetaData = "contentType"
	HTTPMethod      HttpMetaData = "httpMethod"
	HTTPPath        HttpMetaData = "httpPath"
	HTTPQuery       HttpMetaData = "httpQuery"
	HTTPClientIp    HttpMetaData = "httpClientIp"
	HTTPUserAgent   HttpMetaData = "httpUserAgent"
)
