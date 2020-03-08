package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"testing"

	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/harm-matthias-harms/rpm/backend/model"
	"github.com/harm-matthias-harms/rpm/backend/storage"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestMedicalCase(t *testing.T) {
	//setup
	_ = storage.SetMongoDatabase()
	jwtCookie := loginUser(t)

	//models
	mc := model.MedicalCase{Title: "test"}
	mcInvalid := model.MedicalCase{}

	//tests
	t.Run("create", func(t *testing.T) {
		body, header := createFormData(mc, "./medical_case_test.go", false)
		bodyInvalid, headerInvalid := createFormData(mcInvalid, "", false)
		bodyBroken, headerBroken := createFormData(mc, "", true)

		// create medical case
		rec, err := testRequest(http.MethodPost, "/api/medical_cases", body, HandleMedicalCaseCreate, header, nil, jwtCookie)
		if assert.NoError(t, err) {
			response := &jsonStatus{}
			parseResponse(rec, response)
			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.Equal(t, mc.Title, response.Data.(map[string]interface{})["title"])
			assert.NotNil(t, response.Data.(map[string]interface{})["author"].(map[string]interface{})["id"])
			assert.NotNil(t, response.Data.(map[string]interface{})["author"].(map[string]interface{})["username"])
			assert.NotNil(t, response.Data.(map[string]interface{})["createdAt"])

			jsonString, _ := json.Marshal(response.Data)
			json.Unmarshal(jsonString, &mc)
		}

		// no payload error
		rec, err = testRequest(http.MethodPost, "/api/medical_cases", nil, HandleMedicalCaseCreate, header, nil, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
		}

		//invalid medical case
		rec, err = testRequest(http.MethodPost, "/api/medical_cases", bodyInvalid, HandleMedicalCaseCreate, headerInvalid, nil, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
			assert.Equal(t, "title not set", err.(*echo.HTTPError).Message)
		}

		// broken payload
		rec, err = testRequest(http.MethodPost, "/api/medical_cases", bodyBroken, HandleMedicalCaseCreate, headerBroken, nil, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
			assert.Equal(t, errorParseRequest, err.(*echo.HTTPError).Message)
		}
	})

	t.Run("update", func(t *testing.T) {
		mc.Medical.Allergies = "allergy"
		falseMC := mc
		id := primitive.NewObjectID()
		falseMC.ID = id

		body, header := createFormData(mc, "./medical_case_test.go", false)
		body2 := bytes.NewReader(body.Bytes())
		body3 := bytes.NewReader(body.Bytes())
		bodyFalse, headerFalse := createFormData(falseMC, "", false)
		bodyBroken, headerBroken := createFormData(mc, "", true)
		bodyInvalid, headerInvalid := createFormData(mcInvalid, "", false)

		// update medical case
		rec, err := testRequest(http.MethodPut, "/api/medical_cases/:id", body, HandleMedicalCaseEdit, header, map[string]string{"id": mc.ID.Hex()}, jwtCookie)
		if assert.NoError(t, err) {
			response := &model.MedicalCase{}
			parseResponse(rec, response)
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, mc.Title, response.Title)
			assert.NotNil(t, response.Editor.ID)
			assert.NotNil(t, response.Editor.Username)
			assert.NotNil(t, response.EditedAt)
		}

		// no payload error
		_, err = testRequest(http.MethodPost, "/api/medical_cases/:id", nil, HandleMedicalCaseEdit, header, map[string]string{"id": mc.ID.Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
		}

		// no id provided
		_, err = testRequest(http.MethodPut, "/api/medical_cases/:id", body2, HandleMedicalCaseEdit, header, nil, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, errorNoIDParam, err.(*echo.HTTPError).Message)
		}

		// id not match
		_, err = testRequest(http.MethodPut, "/api/medical_cases/:id", body3, HandleMedicalCaseEdit, header, map[string]string{"id": primitive.NewObjectID().Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, errorIDNotMatch, err.(*echo.HTTPError).Message)
		}

		// false id
		_, err = testRequest(http.MethodPut, "/api/medical_cases/:id", bodyFalse, HandleMedicalCaseEdit, headerFalse, map[string]string{"id": id.Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, errorUpdated, err.(*echo.HTTPError).Message)
		}

		// broken payload
		_, err = testRequest(http.MethodPut, "/api/medical_cases/:id", bodyBroken, HandleMedicalCaseEdit, headerBroken, map[string]string{"id": id.Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
			assert.Equal(t, errorParseRequest, err.(*echo.HTTPError).Message)
		}

		// invalid mc
		_, err = testRequest(http.MethodPost, "/api/medical_cases/:id", bodyInvalid, HandleMedicalCaseEdit, headerInvalid, map[string]string{"id": mcInvalid.ID.Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
			assert.Equal(t, "title not set", err.(*echo.HTTPError).Message)
		}
	})

	t.Run("find", func(t *testing.T) {
		// no id provided
		_, err := testRequest(http.MethodGet, "/api/medical_cases/:id", nil, HandleMedicalCaseFind, nil, map[string]string{"id": ""}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, errorNoIDParam, err.(*echo.HTTPError).Message)
		}
		// false id
		_, err = testRequest(http.MethodGet, "/api/medical_cases/:id", nil, HandleMedicalCaseFind, nil, map[string]string{"id": primitive.NewObjectID().Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, errorFind, err.(*echo.HTTPError).Message)
		}
		// good id
		rec, err := testRequest(http.MethodGet, "/api/medical_cases/:id", nil, HandleMedicalCaseFind, nil, map[string]string{"id": mc.ID.Hex()}, jwtCookie)
		if assert.NoError(t, err) {
			response := &model.MedicalCase{}
			parseResponse(rec, response)
			assert.Equal(t, mc.Title, response.Title)
			assert.Equal(t, "allergy", response.Medical.Allergies)
		}
	})

	t.Run("get", func(t *testing.T) {
		q := make(url.Values)
		rec, err := testRequest(http.MethodGet, "/api/medical_cases?"+q.Encode(), nil, HandleMedicalCaseGet, nil, nil, jwtCookie)
		if assert.NoError(t, err) {
			response := model.MedicalCaseList{}
			parseResponse(rec, &response)
			assert.Greater(t, response.Count, int64(0))
		}
		q.Add("title", mc.Title)
		rec, err = testRequest(http.MethodGet, "/api/medical_cases?"+q.Encode(), nil, HandleMedicalCaseGet, nil, nil, jwtCookie)
		if assert.NoError(t, err) {
			response := model.MedicalCaseList{}
			parseResponse(rec, &response)
			assert.Greater(t, response.Count, int64(0))
			assert.Equal(t, mc.Title, response.MedicalCases[0].Title)
		}
		q.Add("page", "1")
		q.Add("limit", "1")
		rec, err = testRequest(http.MethodGet, "/api/medical_cases?"+q.Encode(), nil, HandleMedicalCaseGet, nil, nil, jwtCookie)
		if assert.NoError(t, err) {
			response := model.MedicalCaseList{}
			parseResponse(rec, &response)
			assert.Equal(t, 1, len(response.MedicalCases))
			assert.Equal(t, mc.Title, response.MedicalCases[0].Title)
		}
		q.Del("page")
		q.Del("limit")
		q.Add("author", mc.Author.Username)
		rec, err = testRequest(http.MethodGet, "/api/medical_cases?"+q.Encode(), nil, HandleMedicalCaseGet, nil, nil, jwtCookie)
		if assert.NoError(t, err) {
			response := model.MedicalCaseList{}
			parseResponse(rec, &response)
			assert.Greater(t, response.Count, int64(0))
			assert.Equal(t, mc.Title, response.MedicalCases[0].Title)
			assert.Equal(t, mc.Author.Username, response.MedicalCases[0].Author.Username)
		}
	})

	t.Run("get file", func(t *testing.T) {
		rec, err := testRequest(http.MethodGet, "/api/medical_cases/:mc_id/documents/:id", nil, HandleMedicalCaseFileGet, nil, map[string]string{"mc_id": mc.ID.Hex(), "id": mc.Files[0].ID.Hex()}, jwtCookie)
		if assert.NoError(t, err) {
			assert.Greater(t, len(rec.Body.Bytes()), 0)
		}

		// no id
		_, err = testRequest(http.MethodGet, "/api/medical_cases/:mc_id/documents/:id", nil, HandleMedicalCaseFileGet, nil, map[string]string{"mc_id": "", "id": mc.Files[0].ID.Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
			assert.Equal(t, errorNoIDParam, err.(*echo.HTTPError).Message)
		}

		// no document id
		_, err = testRequest(http.MethodGet, "/api/medical_cases/:mc_id/documents/:id", nil, HandleMedicalCaseFileGet, nil, map[string]string{"mc_id": mc.ID.Hex(), "id": ""}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
			assert.Equal(t, errorNoIDParam, err.(*echo.HTTPError).Message)
		}

		// false mc id
		_, err = testRequest(http.MethodGet, "/api/medical_cases/:mc_id/documents/:id", nil, HandleMedicalCaseFileGet, nil, map[string]string{"mc_id": primitive.NewObjectID().Hex(), "id": mc.Files[0].ID.Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
			assert.Equal(t, errorFind, err.(*echo.HTTPError).Message)
		}

		// false document id
		_, err = testRequest(http.MethodGet, "/api/medical_cases/:mc_id/documents/:id", nil, HandleMedicalCaseFileGet, nil, map[string]string{"mc_id": mc.ID.Hex(), "id": primitive.NewObjectID().Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
			assert.Equal(t, errorFind, err.(*echo.HTTPError).Message)
		}
	})

	t.Run("delete file", func(t *testing.T) {
		//setup
		mc2 := mc
		mc2.ID = primitive.NilObjectID
		mc2.Author.ID = primitive.NewObjectID()
		_ = storage.CreateMedicalCase(nil, &mc2)

		// no id
		_, err := testRequest(http.MethodDelete, "/api/medical_cases/:mc_id/documents/:id", nil, HandleMedicalCaseFileDelete, nil, map[string]string{"mc_id": "", "id": mc.Files[0].ID.Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
			assert.Equal(t, errorNoIDParam, err.(*echo.HTTPError).Message)
		}

		// no document id
		_, err = testRequest(http.MethodDelete, "/api/medical_cases/:mc_id/documents/:id", nil, HandleMedicalCaseFileDelete, nil, map[string]string{"mc_id": mc.ID.Hex(), "id": ""}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
			assert.Equal(t, errorNoIDParam, err.(*echo.HTTPError).Message)
		}

		// false mc id
		_, err = testRequest(http.MethodDelete, "/api/medical_cases/:mc_id/documents/:id", nil, HandleMedicalCaseFileDelete, nil, map[string]string{"mc_id": primitive.NewObjectID().Hex(), "id": mc.Files[0].ID.Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
			assert.Equal(t, errorFind, err.(*echo.HTTPError).Message)
		}

		// false document id
		_, err = testRequest(http.MethodDelete, "/api/medical_cases/:mc_id/documents/:id", nil, HandleMedicalCaseFileDelete, nil, map[string]string{"mc_id": mc.ID.Hex(), "id": primitive.NewObjectID().Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
			assert.Equal(t, errorFind, err.(*echo.HTTPError).Message)
		}

		// not author
		_, err = testRequest(http.MethodDelete, "/api/medical_cases/:mc_id/documents/:id", nil, HandleMedicalCaseFileDelete, nil, map[string]string{"mc_id": mc2.ID.Hex(), "id": mc2.Files[0].ID.Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
			assert.Equal(t, errorNotAuthorized, err.(*echo.HTTPError).Message)
		}

		// good
		fileID := mc.Files[0].ID.Hex()
		rec, err := testRequest(http.MethodDelete, "/api/medical_cases/:mc_id/documents/:id", nil, HandleMedicalCaseFileDelete, nil, map[string]string{"mc_id": mc.ID.Hex(), "id": fileID}, jwtCookie)
		if assert.NoError(t, err) {
			parseResponse(rec, &mc)
		}
	})

	t.Run("delete", func(t *testing.T) {
		// no id provided
		_, err := testRequest(http.MethodDelete, "/api/medical_cases/:id", nil, HandleMedicalCaseDelete, nil, map[string]string{"id": ""}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, errorNoIDParam, err.(*echo.HTTPError).Message)
		}
		// false id
		_, err = testRequest(http.MethodDelete, "/api/medical_cases/:id", nil, HandleMedicalCaseDelete, nil, map[string]string{"id": primitive.NewObjectID().Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, errorDelete, err.(*echo.HTTPError).Message)
		}
		// good id
		_, err = testRequest(http.MethodDelete, "/api/medical_cases/:id", nil, HandleMedicalCaseDelete, nil, map[string]string{"id": mc.ID.Hex()}, jwtCookie)
		assert.NoError(t, err)
	})

	// cleanup
	resetMedicalCase(&mc)
}

func createFormData(mc model.MedicalCase, filePath string, broken bool) (body *bytes.Buffer, header http.Header) {
	body = new(bytes.Buffer)
	header = http.Header{}
	writer := multipart.NewWriter(body)
	if broken {
		writer.WriteField("medicalCase", structToJSONString(mc)[0:4])
	} else {
		writer.WriteField("medicalCase", structToJSONString(mc))
	}

	if filePath != "" {
		file, _ := os.Open(filePath)
		defer file.Close()
		part, _ := writer.CreateFormFile("files", filepath.Base(filePath))
		io.Copy(part, file)
	}

	header.Set(echo.HeaderContentType, writer.FormDataContentType())
	writer.Close()
	return
}

func fileBucket() (bucket *gridfs.Bucket, err error) {
	bucket, err = gridfs.NewBucket(storage.MongoSession, options.GridFSBucket().SetName("files"))
	return
}

func resetMedicalCaseFile(mc *model.MedicalCase) {
	bucket, _ := fileBucket()
	for _, f := range mc.Files {
		_ = bucket.Delete(f.ID)
	}
}

func resetMedicalCase(mc *model.MedicalCase) {
	_, _ = storage.MongoSession.Collection("medicalCases").DeleteMany(nil, bson.M{"title": mc.Title})
}
