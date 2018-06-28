package factory

type Item interface {
	MakeHTML() string
}

type Link interface {
	Item
}

type tray interface {
	Item
	Add(item Item)
}

type Tray struct {
	trays []tray
}

type Page interface {
	Item
	Add(item Item)
	Output()
}

type Factroy interface {
	CreateLink(caption string, url string) Link
	CreateTray(caption string) Tray
	CreatePage(title string, author string) Page
}

type ListFactory struct {
}

func (lf *ListFactory) CreateLink(caption string, url string) Link {
	return &ListLink{caption, url}
}

func (lf *ListFactory) CreateTray(caption string) Tray {
	return &ListTray{caption, url}
}

func (lf *ListFactory) CreatePage(title string, author string) Page {
	return &ListPage{caption, url}
}

type ListLink struct {
	caption string
	url     string
}

func (ll *ListLink) MakeHTML() string {
	return "<li><a href =\"" + ll.url + "\">" + ll.caption + "</a></li>\n"
}

type ListTray struct {
	caption string
}

func (lt *ListTray) MakeHTML() string {
	return "<li><a href =\"" + lt.url + "\">" + lt.caption + "</a></li>\n"
}
