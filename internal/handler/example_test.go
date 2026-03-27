package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/FedorSidorow/shortener/internal/auth"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

type MockTest struct{}

func (m *MockTest) Errorf(format string, args ...any) {}
func (m *MockTest) Fatalf(format string, args ...any) {}
func ExampleAPIHandler_GenerateShortKeyHandler() {

	var (
		ctrl        = gomock.NewController(&MockTest{})
		mockService = NewMockShortenerServicer(ctrl)

		h, _ = NewHandler(mockService)

		bodyRequest = strings.NewReader("http://e.com")
		response    = httptest.NewRecorder()
		mockReq, _  = http.NewRequest(http.MethodPost, "http://localhost/", bodyRequest)
		userID      = uuid.New()
		ctx         = auth.WithUserID(mockReq.Context(), userID)
	)
	mockReq = mockReq.WithContext(ctx)

	mockService.EXPECT().
		GenerateShortURL(mockReq.Context(), "http://e.com", mockReq.Host, userID).
		Return("http://localhost/testKey", nil)

	h.GenerateShortKeyHandler(response, mockReq)
	fmt.Println(response.Code)
	// Output:
	// 201

}
