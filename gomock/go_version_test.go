package gomock

import (
	"testing"

	"github.com/golang/mock/gomock"

	spider "github.com/marmotedu/gopractise-demo/gomock/spider/mock"
)

func TestGetGoVersion(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSpider := spider.NewMockSpider(ctrl)
	mockSpider.EXPECT().GetBody().Return("go1.8.3")
	goVer := GetGoVersion(mockSpider)

	if goVer != "go1.8.3" {
		t.Errorf("Get wrong version %s", goVer)
	}
}
