package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/harm-matthias-harms/rpm/backend/model"
	"github.com/harm-matthias-harms/rpm/backend/storage"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestMedicalCaseGet(t *testing.T) {
	mc := model.MedicalCase{Title: "test get", Author: model.LimitedUser{Username: "test"}}
	mc2 := model.MedicalCase{Title: "test get", Author: model.LimitedUser{Username: "second"}}
	resetMedicalCase(&mc)
	resetMedicalCase(&mc2)
	_ = storage.CreateMedicalCase(nil, &mc)
	_ = storage.CreateMedicalCase(nil, &mc2)
	jwtCookie := loginUser(t)
	q := make(url.Values)
	rec, err := testRequest(http.MethodGet, "/api/medical_cases?"+q.Encode(), nil, HandleMedicalCaseGet, nil, nil, jwtCookie)
	if assert.NoError(t, err) {
		response := model.MedicalCaseList{}
		defer rec.Result().Body.Close()
		body, _ := ioutil.ReadAll(rec.Result().Body)
		_ = json.Unmarshal(body, &response)
		assert.Greater(t, response.Count, int64(0))
	}
	q.Add("title", mc.Title)
	rec, err = testRequest(http.MethodGet, "/api/medical_cases?"+q.Encode(), nil, HandleMedicalCaseGet, nil, nil, jwtCookie)
	if assert.NoError(t, err) {
		response := model.MedicalCaseList{}
		defer rec.Result().Body.Close()
		body, _ := ioutil.ReadAll(rec.Result().Body)
		_ = json.Unmarshal(body, &response)
		assert.Greater(t, response.Count, int64(0))
		assert.Equal(t, mc.Title, response.MedicalCases[0].Title)
	}
	q.Add("page", "1")
	q.Add("limit", "1")
	rec, err = testRequest(http.MethodGet, "/api/medical_cases?"+q.Encode(), nil, HandleMedicalCaseGet, nil, nil, jwtCookie)
	if assert.NoError(t, err) {
		response := model.MedicalCaseList{}
		defer rec.Result().Body.Close()
		body, _ := ioutil.ReadAll(rec.Result().Body)
		_ = json.Unmarshal(body, &response)
		assert.Equal(t, 1, len(response.MedicalCases))
		assert.Equal(t, mc.Title, response.MedicalCases[0].Title)
	}
	q.Del("page")
	q.Del("limit")
	q.Add("author", mc.Author.Username)
	rec, err = testRequest(http.MethodGet, "/api/medical_cases?"+q.Encode(), nil, HandleMedicalCaseGet, nil, nil, jwtCookie)
	if assert.NoError(t, err) {
		response := model.MedicalCaseList{}
		defer rec.Result().Body.Close()
		body, _ := ioutil.ReadAll(rec.Result().Body)
		_ = json.Unmarshal(body, &response)
		assert.Greater(t, response.Count, int64(0))
		assert.Equal(t, mc.Title, response.MedicalCases[0].Title)
		assert.Equal(t, mc.Author.Username, response.MedicalCases[0].Author.Username)
	}
}

func TestMedicalCaseFind(t *testing.T) {
	mc := model.MedicalCase{Title: "test"}
	mc.GeneralInformation.Surgical = true
	resetMedicalCase(&mc)
	_ = storage.CreateMedicalCase(nil, &mc)

	jwtCookie := loginUser(t)
	// no id provided
	_, err := testRequest(http.MethodGet, "/api/medical_cases/:id", nil, HandleMedicalCaseFind, nil, map[string]string{"id": ""}, jwtCookie)
	if assert.Error(t, err) {
		assert.Equal(t, "no or false id provided", err.(*echo.HTTPError).Message)
	}
	// false id
	_, err = testRequest(http.MethodGet, "/api/medical_cases/:id", nil, HandleMedicalCaseFind, nil, map[string]string{"id": primitive.NewObjectID().Hex()}, jwtCookie)
	if assert.Error(t, err) {
		assert.Equal(t, "couldn't find medical case", err.(*echo.HTTPError).Message)
	}
	// good id
	rec, err := testRequest(http.MethodGet, "/api/medical_cases/:id", nil, HandleMedicalCaseFind, nil, map[string]string{"id": mc.ID.Hex()}, jwtCookie)
	if assert.NoError(t, err) {
		response := &model.MedicalCase{}
		defer rec.Result().Body.Close()
		body, _ := ioutil.ReadAll(rec.Result().Body)
		_ = json.Unmarshal(body, response)
		assert.Equal(t, mc.Title, response.Title)
		assert.Equal(t, mc.GeneralInformation.Surgical, response.GeneralInformation.Surgical)
	}
}

func TestMedicalCaseCreate(t *testing.T) {
	mc := model.MedicalCase{Title: "test"}
	resetMedicalCase(&mc)
	mcString, _ := json.Marshal(mc)
	//header.Set(echo.HeaderContentType, echo.MIMEMultipartForm)
	jwtCookie := loginUser(t)

	//test file upload
	file, _ := os.Open("./medical_case_test.go")
	defer file.Close()
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("medicalCase", string(mcString))
	part, _ := writer.CreateFormFile("files", filepath.Base("./medical_case_test.go"))
	io.Copy(part, file)
	header := http.Header{}
	header.Set(echo.HeaderContentType, writer.FormDataContentType())
	writer.Close()

	// create medical case
	rec, err := testRequest(http.MethodPost, "/api/medical_cases", body, HandleMedicalCaseCreate, header, nil, jwtCookie)
	if assert.NoError(t, err) {
		response := &jsonStatus{}
		defer rec.Result().Body.Close()
		body, _ := ioutil.ReadAll(rec.Result().Body)
		_ = json.Unmarshal(body, response)
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, mc.Title, response.Data.(map[string]interface{})["title"])
		assert.NotNil(t, response.Data.(map[string]interface{})["author"].(map[string]interface{})["id"])
		assert.NotNil(t, response.Data.(map[string]interface{})["author"].(map[string]interface{})["username"])
		assert.NotNil(t, response.Data.(map[string]interface{})["createdAt"])
	}
	// no payload error
	rec, err = testRequest(http.MethodPost, "/api/medical_cases", nil, HandleMedicalCaseCreate, header, nil, jwtCookie)
	if assert.Error(t, err) {
		err, ok := err.(*echo.HTTPError)
		if ok {
			assert.Equal(t, http.StatusBadRequest, err.Code)
			assert.Error(t, err)
		}
	}
	// not create invalid
	mcInvalid := model.MedicalCase{}
	resetMedicalCase(&mcInvalid)
	mcInvalidString, _ := json.Marshal(mcInvalid)
	//test file upload
	bodyInvalid := new(bytes.Buffer)
	writerInvalid := multipart.NewWriter(bodyInvalid)
	writerInvalid.WriteField("medicalCase", string(mcInvalidString))
	header = http.Header{}
	header.Set(echo.HeaderContentType, writerInvalid.FormDataContentType())
	writerInvalid.Close()
	rec, err = testRequest(http.MethodPost, "/api/medical_cases", bodyInvalid, HandleMedicalCaseCreate, header, nil, jwtCookie)
	if assert.Error(t, err) {
		err, ok := err.(*echo.HTTPError)
		if ok {
			assert.Equal(t, http.StatusBadRequest, err.Code)
			assert.Equal(t, "title not set", err.Message)
		}
	}
}

func TestMedicalCaseEdit(t *testing.T) {
	mc := model.MedicalCase{Title: "update test", Author: model.LimitedUser{ID: primitive.NewObjectID(), Username: "username"}, CreatedAt: time.Now()}
	resetMedicalCase(&mc)
	_ = storage.CreateMedicalCase(nil, &mc)
	mcString, _ := json.Marshal(mc)
	//header.Set(echo.HeaderContentType, echo.MIMEMultipartForm)
	jwtCookie := loginUser(t)

	//test file upload
	file, _ := os.Open("./medical_case_test.go")
	defer file.Close()
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("medicalCase", string(mcString))
	part, _ := writer.CreateFormFile("files", filepath.Base("./medical_case_test.go"))
	io.Copy(part, file)
	header := http.Header{}
	header.Set(echo.HeaderContentType, writer.FormDataContentType())
	writer.Close()
	body2 := bytes.NewReader(body.Bytes())
	body3 := bytes.NewReader(body.Bytes())

	// update medical case
	rec, err := testRequest(http.MethodPut, "/api/medical_cases/:id", body, HandleMedicalCaseEdit, header, map[string]string{"id": mc.ID.Hex()}, jwtCookie)
	if assert.NoError(t, err) {
		response := &model.MedicalCase{}
		defer rec.Result().Body.Close()
		recBody, _ := ioutil.ReadAll(rec.Result().Body)
		_ = json.Unmarshal(recBody, response)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, mc.Title, response.Title)
		assert.NotNil(t, response.Editor.ID)
		assert.NotNil(t, response.Editor.Username)
		assert.NotNil(t, response.EditedAt)
	}

	// no payload error
	_, err = testRequest(http.MethodPost, "/api/medical_cases/:id", nil, HandleMedicalCaseEdit, header, map[string]string{"id": mc.ID.Hex()}, jwtCookie)
	if assert.Error(t, err) {
		err, ok := err.(*echo.HTTPError)
		if ok {
			assert.Equal(t, http.StatusBadRequest, err.Code)
			assert.Error(t, err)
		}
	}

	// no id provided
	_, err = testRequest(http.MethodPut, "/api/medical_cases/:id", body2, HandleMedicalCaseEdit, header, nil, jwtCookie)
	if assert.Error(t, err) {
		assert.Equal(t, "no or false id provided", err.(*echo.HTTPError).Message)
	}

	// id not match
	_, err = testRequest(http.MethodPut, "/api/medical_cases/:id", body3, HandleMedicalCaseEdit, header, map[string]string{"id": primitive.NewObjectID().Hex()}, jwtCookie)
	if assert.Error(t, err) {
		assert.Equal(t, "id's do not match", err.(*echo.HTTPError).Message)
	}

	// false id
	falseMC := mc
	id := primitive.NewObjectID()
	falseMC.ID = id
	falseMCString, _ := json.Marshal(falseMC)
	bodyFalse := new(bytes.Buffer)
	writerFalse := multipart.NewWriter(bodyFalse)
	writerFalse.WriteField("medicalCase", string(falseMCString))
	header = http.Header{}
	header.Set(echo.HeaderContentType, writerFalse.FormDataContentType())
	writerFalse.Close()
	_, err = testRequest(http.MethodPut, "/api/medical_cases/:id", bodyFalse, HandleMedicalCaseEdit, header, map[string]string{"id": id.Hex()}, jwtCookie)
	if assert.Error(t, err) {
		assert.Equal(t, "couldn't be updated", err.(*echo.HTTPError).Message)
	}

	// broken payload
	bodyBroken := new(bytes.Buffer)
	writerBroken := multipart.NewWriter(bodyBroken)
	writerBroken.WriteField("medicalCase", string(mcString[0:10]))
	header = http.Header{}
	header.Set(echo.HeaderContentType, writerBroken.FormDataContentType())
	writerBroken.Close()
	_, err = testRequest(http.MethodPut, "/api/medical_cases/:id", bodyBroken, HandleMedicalCaseEdit, header, map[string]string{"id": id.Hex()}, jwtCookie)
	if assert.Error(t, err) {
		assert.Equal(t, "couldn't parse request", err.(*echo.HTTPError).Message)
	}

	// not update invalid
	mcInvalid := model.MedicalCase{Title: "invalid update test"}
	resetMedicalCase(&mcInvalid)
	_ = storage.CreateMedicalCase(nil, &mcInvalid)
	mcInvalidString, _ := json.Marshal(mcInvalid)
	//test file upload
	bodyInvalid := new(bytes.Buffer)
	writerInvalid := multipart.NewWriter(bodyInvalid)
	writerInvalid.WriteField("medicalCase", string(mcInvalidString))
	header = http.Header{}
	header.Set(echo.HeaderContentType, writerInvalid.FormDataContentType())
	writerInvalid.Close()
	_, err = testRequest(http.MethodPost, "/api/medical_cases/:id", bodyInvalid, HandleMedicalCaseEdit, header, map[string]string{"id": mcInvalid.ID.Hex()}, jwtCookie)
	if assert.Error(t, err) {
		err, ok := err.(*echo.HTTPError)
		if ok {
			assert.Equal(t, http.StatusBadRequest, err.Code)
			assert.Equal(t, "author not set", err.Message)
		}
	}
}

func TestMedicalCaseFileGet(t *testing.T) {
	// setup
	_ = storage.SetMongoDatabase()
	bucket, _ := fileBucket()
	jwtCookie := loginUser(t)

	// upload a fil
	ustream, _ := bucket.OpenUploadStream("test.txt", options.GridFSUpload())
	_, _ = ustream.Write([]byte("test file"))
	fileID := ustream.FileID
	ustream.Close()

	// create medical case with file
	mc := model.MedicalCase{Title: "test"}
	resetMedicalCase(&mc)
	mc.Files = append(mc.Files, model.MedicalCaseFile{ID: fileID.(primitive.ObjectID), Name: "test.txt", Size: 9})
	_ = storage.CreateMedicalCase(nil, &mc)

	// test get
	rec, err := testRequest(http.MethodGet, "/api/medical_cases/:mc_id/documents/:id", nil, HandleMedicalCaseFileGet, nil, map[string]string{"mc_id": mc.ID.Hex(), "id": fileID.(primitive.ObjectID).Hex()}, jwtCookie)
	assert.NoError(t, err)
	assert.Equal(t, []byte("test file"), rec.Body.Bytes())
}

func fileBucket() (bucket *gridfs.Bucket, err error) {
	_ = storage.SetMongoDatabase()
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
	_ = storage.SetMongoDatabase()
	_, _ = storage.MongoSession.Collection("medical_cases").DeleteOne(nil, bson.M{"title": mc.Title})
}
