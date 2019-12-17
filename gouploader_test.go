package gouploader

import(
  "testing"
  "mime/multipart"
  "github.com/stretchr/testify/assert"
)

func TestUploadFileEmpty(t *testing.T){
  var mp *multipart.FileHeader
  actual, err := UploadFile(mp,"./public/uploads")
  assert.NotNil(t, err)
  assert.EqualValues(t, err.Error(), "file_error:empty file")
  assert.EqualValues(t, actual, "")
}
