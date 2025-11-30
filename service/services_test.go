package service

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func Test_genArticleId(t *testing.T) {
	convey.Convey("Test_genArticleId", t, func() {
		id := randArticleId()
		t.Logf("id: %s", id)
	})
}
