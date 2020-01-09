package api_test

//
//import (
//	"fmt"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//
//	"github.com/astaxie/beego"
//	"github.com/cloudsonic/sonic-server/api/responses"
//	"github.com/cloudsonic/sonic-server/domain"
//	"github.com/cloudsonic/sonic-server/persistence"
//	. "github.com/cloudsonic/sonic-server/tests"
//	"github.com/cloudsonic/sonic-server/utils"
//	. "github.com/smartystreets/goconvey/convey"
//)
//
//func getCoverArt(params ...string) (*http.Request, *httptest.ResponseRecorder) {
//	url := AddParams("/rest/getCoverArt.view", params...)
//	r, _ := http.NewRequest("GET", url, nil)
//	w := httptest.NewRecorder()
//	beego.BeeApp.Handlers.ServeHTTP(w, r)
//	log.Debug(r, "testing TestGetCoverArt", fmt.Sprintf("\nUrl: %s\nStatus Code: [%d]\n%#v", r.URL, w.Code, w.HeaderMap))
//	return r, w
//}
//
//func TestGetCoverArt(t *testing.T) {
//	Init(t, false)
//
//	mockMediaFileRepo := persistence.CreateMockMediaFileRepo()
//	utils.DefineSingleton(new(domain.MediaFileRepository), func() domain.MediaFileRepository {
//		return mockMediaFileRepo
//	})
//
//	Convey("Subject: GetCoverArt Endpoint", t, func() {
//		Convey("Should fail if missing Id parameter", func() {
//			_, w := getCoverArt()
//
//			So(w.Body, ShouldReceiveError, responses.ErrorMissingParameter)
//		})
//		Convey("When id is found", func() {
//			mockMediaFileRepo.SetData(`[{"Id":"2","HasCoverArt":true,"Path":"tests/fixtures/01 Invisible (RED) Edit Version.mp3"}]`, 1)
//			_, w := getCoverArt("id=2")
//
//			So(w.Body.Bytes(), ShouldMatchMD5, "e859a71cd1b1aaeb1ad437d85b306668")
//			So(w.Header().Get("Content-Type"), ShouldEqual, "image/jpeg")
//		})
//		Convey("When id is found but file is unavailable", func() {
//			mockMediaFileRepo.SetData(`[{"Id":"2","HasCoverArt":true,"Path":"tests/fixtures/NOT_FOUND.mp3"}]`, 1)
//			_, w := getCoverArt("id=2")
//
//			So(w.Body, ShouldReceiveError, responses.ErrorDataNotFound)
//		})
//		Convey("When the engine reports an error", func() {
//			mockMediaFileRepo.SetData(`[{"Id":"2","HasCoverArt":true,"Path":"tests/fixtures/NOT_FOUND.mp3"}]`, 1)
//			mockMediaFileRepo.SetError(true)
//			_, w := getCoverArt("id=2")
//
//			So(w.Body, ShouldReceiveError, responses.ErrorGeneric)
//		})
//		Convey("When specifying a size", func() {
//			mockMediaFileRepo.SetData(`[{"Id":"2","HasCoverArt":true,"Path":"tests/fixtures/01 Invisible (RED) Edit Version.mp3"}]`, 1)
//			_, w := getCoverArt("id=2", "size=100")
//
//			So(w.Body.Bytes(), ShouldMatchMD5, "04378f523ca3e8ead33bf7140d39799e")
//			So(w.Header().Get("Content-Type"), ShouldEqual, "image/jpeg")
//		})
//		Reset(func() {
//			mockMediaFileRepo.SetData("[]", 0)
//			mockMediaFileRepo.SetError(false)
//		})
//	})
//}
