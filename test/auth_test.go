package test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/p-jirayusakul/golang-echo-homework-2/configs"
	"github.com/p-jirayusakul/golang-echo-homework-2/database"
	"github.com/p-jirayusakul/golang-echo-homework-2/handlers"
	"github.com/p-jirayusakul/golang-echo-homework-2/handlers/request"
	"github.com/p-jirayusakul/golang-echo-homework-2/test/mockup"
	"github.com/p-jirayusakul/golang-echo-homework-2/utils"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestRegister(t *testing.T) {
	var uuid uuid.UUID
	uuid.Scan("6825a472-35cb-48a9-a1ac-73a4f263464a")

	testCases := []struct {
		name          string
		body          string
		buildStubs    func(store *mockup.MockStore, body request.RegisterRequest)
		checkResponse func(t *testing.T, status int, err error)
	}{
		{
			name: "OK",
			body: `{"email":"test@email.com","password":"123456"}`,
			buildStubs: func(store *mockup.MockStore, body request.RegisterRequest) {
				store.EXPECT().IsEmailAlreadyExists(body.Email).Times(1).Return(false, nil)
				store.EXPECT().CreateAccounts(gomock.Any()).Times(1).Return(uuid, nil)
				store.EXPECT().CreateProfiles(database.Profiles{
					UserID: uuid,
					Email:  body.Email,
				}).Times(1).Return(nil)
			},
			checkResponse: func(t *testing.T, status int, err error) {
				require.NoError(t, err)
				require.Equal(t, http.StatusCreated, status)
			},
		},
		{
			name: "Bad Request - this email is already used",
			body: `{"email":"test@email.com","password":"123456"}`,
			buildStubs: func(store *mockup.MockStore, body request.RegisterRequest) {
				store.EXPECT().IsEmailAlreadyExists(body.Email).Times(1).Return(true, nil)
			},
			checkResponse: func(t *testing.T, status int, err error) {
				require.Error(t, err)
				require.Equal(t, replaceStringError(http.StatusBadRequest, err.Error()), utils.ErrEmailIsAlreadyUsed.Error())
			},
		},
		{
			name: "Bad Request - email invalid format",
			body: `{"email":"testemail.com","password":"123456"}`,
			buildStubs: func(store *mockup.MockStore, body request.RegisterRequest) {
			},
			checkResponse: func(t *testing.T, status int, err error) {
				require.Error(t, err)
				require.Equal(t, replaceStringError(http.StatusBadRequest, err.Error()), "Key: 'RegisterRequest.Email' Error:Field validation for 'Email' failed on the 'email' tag")
			},
		},
		{
			name: "Internal Server Error",
			body: `{"email":"test@email.com","password":"123456"}`,
			buildStubs: func(store *mockup.MockStore, body request.RegisterRequest) {
				store.EXPECT().IsEmailAlreadyExists(body.Email).Times(1).Return(false, nil)
				store.EXPECT().CreateAccounts(gomock.Any()).Times(1).Return(uuid, nil)
				store.EXPECT().CreateProfiles(database.Profiles{
					UserID: uuid,
					Email:  body.Email,
				}).Times(1).Return(errors.New("some thing went wrong"))
			},
			checkResponse: func(t *testing.T, status int, err error) {
				require.Error(t, err)
				require.Equal(t, replaceStringError(http.StatusInternalServerError, err.Error()), "some thing went wrong")
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			cfg := configs.InitConfigs(".env")
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockup.NewMockStore(ctrl)
			var dto request.RegisterRequest
			err := json.Unmarshal([]byte(tc.body), &dto)
			require.NoError(t, err)
			tc.buildStubs(store, dto)

			app := echo.New()
			app.Validator = utils.NewCustomValidator()
			app.Use(utils.ErrorHandler)

			req := httptest.NewRequest(http.MethodPost, "/auth/register", strings.NewReader(tc.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := app.NewContext(req, rec)
			handler := handlers.NewHandler(app, cfg, store)

			err = handler.Register(c)
			tc.checkResponse(t, c.Response().Status, err)
		})
	}
}

func TestUpdateProfiles(t *testing.T) {
	var uuid uuid.UUID
	uid := "6825a472-35cb-48a9-a1ac-73a4f263464a"
	uuid.Scan(uid)

	testCases := []struct {
		name          string
		body          string
		buildStubs    func(store *mockup.MockStore, body request.UpdateProfilesRequest)
		checkResponse func(t *testing.T, status int, err error)
	}{
		{
			name: "OK - Update Profiles",
			body: fmt.Sprintf(`{"userId":"%s","firstName":"testUpdate111","lastName":"testUpdate111"}`, uid),
			buildStubs: func(store *mockup.MockStore, body request.UpdateProfilesRequest) {
				store.EXPECT().UpdateProfiles(database.UpdateProfilesParams{
					UserID:    uuid,
					FirstName: body.FirstName,
					LastName:  body.LastName,
				}).Times(1).Return(nil)
			},
			checkResponse: func(t *testing.T, status int, err error) {
				require.NoError(t, err)
				require.Equal(t, http.StatusOK, status)
			},
		},
		{
			name: "Internal Server Error - Data not found",
			body: fmt.Sprintf(`{"userId":"%s","firstName":"testUpdate111","lastName":"testUpdate111"}`, uid),
			buildStubs: func(store *mockup.MockStore, body request.UpdateProfilesRequest) {
				store.EXPECT().UpdateProfiles(database.UpdateProfilesParams{
					UserID:    uuid,
					FirstName: body.FirstName,
					LastName:  body.LastName,
				}).Times(1).Return(utils.ErrDataNotFound)
			},
			checkResponse: func(t *testing.T, status int, err error) {
				require.Error(t, err)
				require.Equal(t, replaceStringError(http.StatusInternalServerError, err.Error()), utils.ErrDataNotFound.Error())
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			cfg := configs.InitConfigs(".env")
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockup.NewMockStore(ctrl)
			var dto request.UpdateProfilesRequest
			err := json.Unmarshal([]byte(tc.body), &dto)
			require.NoError(t, err)
			tc.buildStubs(store, dto)

			app := echo.New()
			app.Validator = utils.NewCustomValidator()
			app.Use(utils.ErrorHandler)

			req := httptest.NewRequest(http.MethodPatch, "/users/profiles", strings.NewReader(tc.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := app.NewContext(req, rec)
			handler := handlers.NewHandler(app, cfg, store)

			err = handler.UpdateProfiles(c)
			tc.checkResponse(t, c.Response().Status, err)
		})
	}
}

func replaceStringError(code int, msg string) string {
	replace := fmt.Sprintf("code=%d, message=", code)
	return strings.Replace(msg, replace, "", -1)
}
