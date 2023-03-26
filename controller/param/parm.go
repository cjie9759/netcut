package param

type Page struct {
	PageNum  int
	PageSize int
	Statu    string
	Desc     bool
}

func (p *Page) Default() {
	if p.PageNum == 0 {
		p.PageNum = 1
	}
	if p.PageSize == 0 {
		p.PageSize = 50
	}
}
