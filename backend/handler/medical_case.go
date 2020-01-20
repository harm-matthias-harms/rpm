package handler

import (
	"bufio"
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/harm-matthias-harms/rpm/backend/model"
	"github.com/harm-matthias-harms/rpm/backend/storage"
	"github.com/harm-matthias-harms/rpm/backend/utils"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// HandleMedicalCaseGet returns medical cases in a short version
func HandleMedicalCaseGet(c echo.Context) (err error) {
	params := new(model.MedicalCaseQuery)
	if err = c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorParseParams)
	}
	filter := map[string]interface{}{}
	if params.Title != "" {
		filter["title"] = primitive.Regex{Pattern: params.Title}
	}
	if params.Author != "" {
		filter["author.username"] = primitive.Regex{Pattern: params.Author}
	}
	mcs, err := storage.GetMedicalCases(c.Request().Context(), filter, params.Page, params.PageSize)
	count, err := storage.CountMedicalCases(c.Request().Context(), filter)
	response := model.MedicalCaseList{Count: count, MedicalCases: mcs}
	return c.JSON(http.StatusOK, response)
}

// HandleMedicalCaseFind returns a medical case
func HandleMedicalCaseFind(c echo.Context) (err error) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorNoIDParam)
	}
	mc, err := storage.FindMedicalCase(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorFind)
	}
	return c.JSON(http.StatusOK, mc)
}

// HandleMedicalCaseCreate updates a medical case from a formdata object
func HandleMedicalCaseCreate(c echo.Context) (err error) {
	form, err := c.MultipartForm()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error()) //"no multipart form provided"
	}

	jsonString := form.Value["medicalCase"]
	mc := new(model.MedicalCase)
	err = json.Unmarshal([]byte(jsonString[0]), mc)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorParseRequest)
	}
	cookie, _ := c.Cookie(echo.HeaderAuthorization)
	token, _ := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(utils.GetEnv("JWT_SECRET", "secret")), nil
	})
	claims := token.Claims.(jwt.MapClaims)
	objectID, err := primitive.ObjectIDFromHex(claims["id"].(string))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errorAuthParse)
	}
	mc.Author.ID = objectID
	mc.Author.Username = claims["username"].(string)
	mc.CreatedAt = time.Now()
	if err = mc.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// load files
	for _, file := range form.File["files"] {
		err := storage.StoreMedicalCaseFile(mc, file)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, errorCreated)
		}
	}
	// create medical case
	if err = storage.CreateMedicalCase(c.Request().Context(), mc); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errorCreated)
	}

	return c.JSON(http.StatusCreated, jsonStatus{Success: true, Data: mc})
}

// HandleMedicalCaseEdit creates a medical case from a formdata object
func HandleMedicalCaseEdit(c echo.Context) (err error) {
	form, err := c.MultipartForm()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error()) //"no multipart form provided"
	}

	jsonString := form.Value["medicalCase"]
	mc := new(model.MedicalCase)
	err = json.Unmarshal([]byte(jsonString[0]), mc)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorParseRequest)
	}

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorNoIDParam)
	}
	if mc.ID != id {
		return echo.NewHTTPError(http.StatusBadRequest, errorIDNotMatch)
	}

	cookie, _ := c.Cookie(echo.HeaderAuthorization)
	token, _ := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(utils.GetEnv("JWT_SECRET", "secret")), nil
	})
	claims := token.Claims.(jwt.MapClaims)
	objectID, err := primitive.ObjectIDFromHex(claims["id"].(string))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errorAuthParse)
	}
	mc.Editor.ID = objectID
	mc.Editor.Username = claims["username"].(string)
	mc.EditedAt = time.Now()
	if err = mc.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// load files
	for _, file := range form.File["files"] {
		err := storage.StoreMedicalCaseFile(mc, file)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, errorCreated)
		}
	}

	// update medical case
	if err = storage.UpdateMedicalCase(c.Request().Context(), mc); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errorUpdated)
	}

	return c.JSON(http.StatusOK, mc)
}

// HandleMedicalCaseDelete updates an preset
func HandleMedicalCaseDelete(c echo.Context) (err error) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorNoIDParam)
	}

	cookie, _ := c.Cookie(echo.HeaderAuthorization)
	token, _ := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(utils.GetEnv("JWT_SECRET", "secret")), nil
	})
	claims := token.Claims.(jwt.MapClaims)
	userID, err := primitive.ObjectIDFromHex(claims["id"].(string))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errorAuthParse)
	}

	count, err := storage.DeleteMedicalCase(c.Request().Context(), id, userID)
	if err != nil || count == int64(0) {
		return echo.NewHTTPError(http.StatusBadRequest, errorDelete)
	}

	return c.JSON(http.StatusOK, jsonStatus{Success: true})
}

// HandleMedicalCaseFileGet serves the files for a medical case
func HandleMedicalCaseFileGet(c echo.Context) error {
	mcID, err := primitive.ObjectIDFromHex(c.Param("mc_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorNoIDParam)
	}
	fileID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorNoIDParam)
	}
	mc, err := storage.FindMedicalCase(c.Request().Context(), mcID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorFind)
	}
	for _, file := range mc.Files {
		if file.ID == fileID {
			var b bytes.Buffer
			w := bufio.NewWriter(&b)
			if err = storage.GetMedicalCaseFile(fileID, w); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, errorFind)
			}
			c.Response().Header().Set("Content-Disposition", "attachment; filename="+file.Name)
			return c.Stream(http.StatusOK, http.DetectContentType(b.Bytes()), &b)
		}
	}
	return echo.NewHTTPError(http.StatusBadRequest, errorFind)
}
