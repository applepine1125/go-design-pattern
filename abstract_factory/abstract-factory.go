package abstract_factory

type Item interface {
	MakeHTML() string
}

type Link interface {
	Item
}

type Tray interface {
	Item
	Add(item Item)
}

type Page interface {
	Item
	Add(item Item)
	Output() string
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
	return &ListTray{caption: caption}
}

func (lf *ListFactory) CreatePage(title string, author string) Page {
	return &ListPage{title: title, author: author}
}

type ListLink struct {
	caption string
	url     string
}

func (ll *ListLink) MakeHTML() string {
	return "<li>\n<a href=\"" + ll.url + "\">\n" + ll.caption + "\n</a>\n</li>"
}

type ListTray struct {
	Tray
	trays   []Item
	caption string
}

func (lt *ListTray) Add(item Item) {
	lt.trays = append(lt.trays, item)
}

func (lt *ListTray) MakeHTML() string {
	tray := "<li>\n" + lt.caption + "\n<ul>\n"
	for _, item := range lt.trays {
		tray += item.MakeHTML() + "\n"
	}
	tray += "</ul>\n</li>\n"
	return tray
}

type ListPage struct {
	content []Item
	title   string
	author  string
}

func (lp *ListPage) MakeHTML() string {
	page := "<html>\n<head>\n<title>\n" + lp.title + "\n</title>\n</head>\n"
	page += "<body>\n"
	page += "<h1>\n" + lp.title + "\n</h1>\n"
	page += "<ul>\n"
	for _, item := range lp.content {
		page += item.MakeHTML()
	}
	page += "</ul>\n"
	page += "<hr>\n<address>\n" + lp.author + "\n</address>"
	page += "\n</body>\n</html>"
	return page
}

func (lp *ListPage) Add(item Item) {
	lp.content = append(lp.content, item)

}
func (lp *ListPage) Output() string {
	return lp.MakeHTML()
}
