package utils

type Page struct {
	//当前页
	PageNum int
	//每页的数量
	PageSize int

	Total int

	//总页数
	Pages int

	//前一页
	PrePage int
	//下一页
	NextPage int
}

func (p *Page) Init(pageSize, pageNum, total int) {
	p.PageSize = pageSize
	p.PageNum = pageNum
	p.Total = total
	p.Pages = total/pageSize + 1
	p.PrePage = pageNum - 1
	p.NextPage = pageNum + 1
}
