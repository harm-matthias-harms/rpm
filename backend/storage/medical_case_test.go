package storage

import (
	"bufio"
	"bytes"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/harm-matthias-harms/rpm/backend/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestCountMedicalCases(t *testing.T) {
	mc := &model.MedicalCase{Title: "count test"}
	resetMedicalCaseDatabase(mc)
	_ = CreateMedicalCase(nil, mc)
	count, err := CountMedicalCases(nil, nil)
	if assert.NoError(t, err) {
		assert.Greater(t, count, int64(0))
	}
	filter := map[string]interface{}{"title": "count test"}
	count, err = CountMedicalCases(nil, filter)
	if assert.NoError(t, err) {
		assert.Equal(t, int64(1), count)
	}
}

func TestGetMedicalCases(t *testing.T) {
	mc := &model.MedicalCase{Title: "get medical cases test"}
	mc.GeneralInformation.Surgical = true
	resetMedicalCaseDatabase(mc)
	_ = CreateMedicalCase(nil, mc)
	// test no filter
	result, err := GetMedicalCases(nil, nil, 1, 1)
	if assert.NoError(t, err) {
		assert.Equal(t, 1, len(result))
	}
	// test filter
	filter := map[string]interface{}{"title": "get medical cases test"}
	result, err = GetMedicalCases(nil, filter, 1, 1)
	if assert.NoError(t, err) {
		assert.Equal(t, 1, len(result))
	}
	// test multiple filters
	filter["generalInformation.surgical"] = true
	result, err = GetMedicalCases(nil, filter, 1, 1)
	if assert.NoError(t, err) {
		assert.Equal(t, 1, len(result))
		assert.Equal(t, mc.ID, result[0].ID)
	}
	// test load all
	mc2 := &model.MedicalCase{Title: "get test"}
	resetMedicalCaseDatabase(mc2)
	_ = CreateMedicalCase(nil, mc2)
	result, err = GetMedicalCases(nil, nil, 1, 0)
	if assert.NoError(t, err) {
		assert.Greater(t, len(result), 1)
	}
}

func TestFindMedicalCase(t *testing.T) {
	// First create it because no medical case is inserted
	mc := &model.MedicalCase{Title: "test find medical case"}
	resetMedicalCaseDatabase(mc)
	err := CreateMedicalCase(nil, mc)
	assert.NoError(t, err)
	mcFound, err := FindMedicalCase(nil, mc.ID)
	if assert.NoError(t, err) {
		assert.Equal(t, "test find medical case", mcFound.Title)
	}
	notExist := &model.MedicalCase{Title: "title"}
	_, err = FindMedicalCase(nil, notExist.ID)
	assert.Error(t, err)
}

func TestCreateMedicalCase(t *testing.T) {
	mc := &model.MedicalCase{Title: "test create medical case"}
	resetMedicalCaseDatabase(mc)
	err := CreateMedicalCase(nil, mc)
	assert.NoError(t, err)
}

func TestCreateMedicalCaseFile(t *testing.T) {
	// Setup
	_ = SetMongoDatabase()
	mc := &model.MedicalCase{Title: "test create file"}

	// create multipart form file to recieve the FileHeader
	file, _ := os.Open("./medical_case_test.go")
	defer file.Close()
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("files", filepath.Base("./medical_case_test.go"))
	io.Copy(part, file)
	writer.Close()

	// reading form
	r := multipart.NewReader(body, writer.Boundary())
	form, _ := r.ReadForm(10)

	// test file create
	for _, file := range form.File["files"] {
		err := StoreMedicalCaseFile(mc, file)
		assert.NoError(t, err)
	}
	resetMedicalCaseFile(mc)
}

func TestGetMedicalCaseFile(t *testing.T) {
	// setup
	_ = SetMongoDatabase()
	bucket, _ := fileBucket()

	// upload a fil
	ustream, _ := bucket.OpenUploadStream("test.txt", options.GridFSUpload())
	_, _ = ustream.Write([]byte("test file"))
	fileID := ustream.FileID
	ustream.Close()

	// test get
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	err := GetMedicalCaseFile(fileID.(primitive.ObjectID), w)
	assert.NoError(t, err)
	assert.Equal(t, "test file", b.String())

	// cleanup
	_ = bucket.Delete(fileID)
}

func resetMedicalCaseFile(mc *model.MedicalCase) {
	bucket, _ := fileBucket()
	for _, f := range mc.Files {
		_ = bucket.Delete(f.ID)
	}
}

func resetMedicalCaseDatabase(mc *model.MedicalCase) {
	_ = SetMongoDatabase()
	_, _ = mcCollection().DeleteMany(nil, bson.M{"title": mc.Title})
}
