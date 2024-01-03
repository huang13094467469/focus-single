package model

import (
	"focus-single/internal/model/entity"
)

// CnnvdGetListInput 获取内容列表
type CnnvdGetListInput struct {
	Type       string // 内容模型
	CategoryId uint   // 栏目ID
	Page       int    // 分页号码
	Size       int    // 分页数量，最大50
	Sort       int    // 排序类型(0:最新, 默认。1:活跃, 2:热度)
	UserId     uint   // 要查询的用户ID
	Id         string
	Name       string
	Severity   string
}

// CnnvdGetListOutput 查询列表结果
type CnnvdGetListOutput struct {
	List     []CnnvdGetListOutputItem `json:"list" description:"列表"`
	Page     int                      `json:"page" description:"分页码"`
	Size     int                      `json:"size" description:"分页数量"`
	Total    int                      `json:"total" description:"数据总数"`
	Id       string
	Name     string
	Severity string
}

// CnnvdSearchInput 搜索列表
type CnnvdSearchInput struct {
	Key        string // 关键字
	Type       string // 内容模型
	CategoryId uint   // 栏目ID
	Page       int    // 分页号码
	Size       int    // 分页数量，最大50
	Sort       int    // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// CnnvdSearchOutput 搜索列表结果
type CnnvdSearchOutput struct {
	List  []CnnvdSearchOutputItem `json:"list"`  // 列表
	Stats map[string]int          `json:"stats"` // 搜索统计
	Page  int                     `json:"page"`  // 分页码
	Size  int                     `json:"size"`  // 分页数量
	Total int                     `json:"total"` // 数据总数
}

type CnnvdGetListOutputItem struct {
	Cnnvd    *CnnvdListItem           `json:"Cnnvd"`
	Category *ContentListCategoryItem `json:"category"`
	User     *ContentListUserItem     `json:"user"`
}

type CnnvdSearchOutputItem struct {
	CnnvdGetListOutputItem
}

// CnnvdGetDetailOutput 查询详情结果
type CnnvdGetDetailOutput struct {
	Cnnvd *entity.Cnnvd `json:"Cnnvd"`
	User  *entity.User  `json:"user"`
}

// CnnvdCreateUpdateBase 创建/修改内容基类
type CnnvdCreateUpdateBase struct {
	Type       string   // 内容模型
	CategoryId uint     // 栏目ID
	Title      string   // 标题
	Cnnvd      string   // 内容
	Brief      string   // 摘要
	Thumb      string   // 缩略图
	Tags       []string // 标签名称列表，以JSON存储
	Referer    string   // 内容来源，例如github/gitee
}

// CnnvdCreateInput 创建内容
type CnnvdCreateInput struct {
	CnnvdCreateUpdateBase
	UserId uint
}

// CnnvdCreateOutput 创建内容返回结果
type CnnvdCreateOutput struct {
	CnnvdId uint `json:"Cnnvd_id"`
}

// CnnvdUpdateInput 修改内容
type CnnvdUpdateInput struct {
	CnnvdCreateUpdateBase
	Id uint
}

// CnnvdListItem 主要用于列表展示
type CnnvdListItem struct {
	Id          string `json:"id"          description:"id"`
	CveId       string `json:"cveId"       description:"CVE编号"`
	Name        string `json:"name"        description:"漏洞名称"`
	Published   string `json:"published"   description:"收录时间"`
	Modified    string `json:"modified"    description:"更新时间"`
	Source      string `json:"source"      description:""`
	Severity    string `json:"severity"    description:"危害等级"`
	VulnType    string `json:"vulnType"    description:"漏洞类型"`
	Description string `json:"description" description:"漏洞简介"`
	BugtraqId   string `json:"bugtraqId"   description:"漏洞描述"`
	Solution    string `json:"solution"    description:""`
}
