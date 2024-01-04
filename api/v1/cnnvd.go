package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CnnvdIndexReq struct {
	g.Meta `path:"/cnnvd" method:"get" tags:"国家漏洞库" summary:"查询列表"`
	CommonPaginationReq
	Id       string `json:"id" in:"query"`
	CveId    string `json:"cveId" in:"query"`
	Name     string `json:"name" in:"query"`
	Severity string `json:"severity" in:"query"`
}

type CnnvdIndexRes struct {
	ContentGetListCommonRes
	Id       string
	CveId    string
	Name     string
	Severity string
}

type CnnvdDetailReq struct {
	g.Meta `path:"/cnnvd/detail" method:"get" tags:"漏洞" summary:"展示cnnvd详情页面" `
	Id     string `dc:"内容id"`
}
type CnnvdDetailRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type CnnvdReq struct {
	g.Meta `path:"/cnnvd/list" method:"get" summary:"国家漏洞库" tags:"查询列表"`
}
type CnnvdRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type PostReq struct {
	g.Meta `path:"/cnnvd" method:"post" summary:"国家漏洞库" tags:"添加"`
}
type PostRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type DeleteReq struct {
	g.Meta `path:"/cnnvd/{id}" method:"delete" summary:"国家漏洞库" tags:"删除1"`
}
type DeleteRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}
