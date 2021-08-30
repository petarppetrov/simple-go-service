package internal_test

import (
	"net/http"
	"net/http/httptest"

	internal "github.com/nenov92/simple-go-service/cmd/simple-go-service/internal"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Presenter", func() {

	var presenter *internal.Presenter
	var recorder *httptest.ResponseRecorder
	var mockContext *gin.Context

	BeforeEach(func() {
		presenter = internal.NewPresenter()
		recorder = httptest.NewRecorder()
		mockContext, _ = gin.CreateTestContext(recorder)
		mockContext.Request = httptest.NewRequest("GET", "/v1/data", nil)
	})

	Describe("Get Data", func() {
		When("Get Data is called with proper If-Modified-Since Header", func() {
			It("should return a data list with JSON Content-Type", func() {
				mockContext.Request.Header.Set("If-Modified-Since", "Mon, 02 Jan 2006 15:04:05 GMT")
				presenter.GetData(mockContext)
				Expect(mockContext.Writer.Status()).To(Equal(http.StatusOK))
			})
		})

		When("Get Data is called without proper If-Modified-Since Header", func() {
			It("should return a data list with JSON Content-Type", func() {
				presenter.GetData(mockContext)
				Expect(mockContext.Writer.Status()).To(Equal(http.StatusBadRequest))
			})
		})
	})

})
