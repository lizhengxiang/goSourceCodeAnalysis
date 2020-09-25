// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

// Common HTTP methods.
//
// Unless otherwise noted, these are defined in RFC 7231 section 4.3.
const (
	MethodGet     = "GET"
	MethodHead    = "HEAD"
	MethodPost    = "POST"
	MethodPut     = "PUT"
	MethodPatch   = "PATCH" // RFC 5789
	MethodDelete  = "DELETE"
	MethodConnect = "CONNECT"
	MethodOptions = "OPTIONS"
	MethodTrace   = "TRACE"
)

/*
   POST    在Request-URI所标识的资源后附加新的数据
　　HEAD    请求获取由Request-URI所标识的资源的响应消息报头
　　PUT     请求服务器存储一个资源，并用Request-URI作为其标识
　　DELETE  请求服务器删除Request-URI所标识的资源
　　TRACE   请求服务器回送收到的请求信息，主要用于测试或诊断
　　CONNECT 保留将来使用
　　OPTIONS 请求查询服务器的性能，或者查询与资源相关的选项和需求

    Accept：用于高速服务器，客户机支持的数据类型
　　Accept-Charset：用于告诉服务器，客户机采用的编码格式
　　Accept-Encoding：用于告诉服务器，客户机支持的数据压缩格式
　　Accept-Language：客户机的语言环境
　　Host：客户机通过这个头高速服务器，想访问的主机名
　　If-Modified-Since：客户机通过这个头告诉服务器，资源的缓存时间
　　Referer：客户机通过这个头告诉服务器，它是从哪个资源来访问服务器的（防盗链）
　　User-Agent：客户机通过这个头告诉服务器，客户机的软件环境
　　Cookie：客户机通过这个头可以向服务器带数据
　　Connection：处理完这次请求后是否断开连接还是继续保持连接
　　Date：当前时间值

　　Location：这个头配合302状态码使用，用于告诉客户找谁。
　　Server：服务器通过这个头告诉浏览器服务器的类型。
　　Content-Encoding：服务器通过这个头告诉浏览器数据的压缩格式。
　　Content-Length：服务器通过这个头告诉浏览器回送数据的长度
　　Content-Type：服务器通过这个头告诉浏览器回送数据的类型
　　Last-Modified：告诉浏览器当前资源的最后缓存时间
　　Refresh：告诉浏览器隔多久刷新一次
　　Content-Disposition：告诉浏览器以下载方式打开数据
　　Transfer-Encoding：告诉浏览器数据的传送格式
　　ETag：缓存相关的头
 */