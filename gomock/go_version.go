package gomock

import (
	"github.com/marmotedu/gopractise-demo/gomock/spider"
)

func GetGoVersion(s spider.Spider) string {
	body := s.GetBody()
	return body
}
