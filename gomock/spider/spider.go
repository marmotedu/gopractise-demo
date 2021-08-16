package spider

//go:generate mockgen -destination mock/mock_spider.go -package spider github.com/marmotedu/gopractise-demo/gomock/spider Spider

type Spider interface {
	GetBody() string
}
