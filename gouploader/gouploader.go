package gouploader

import (
  "errors"
  "fmt"
  "io"
  "log"
  "os"
  "path/filepath"
  "mime/multipart"
  "github.com/sampila/go-randomstr/randomstr"
)

func UploadFile(f *multipart.FileHeader, t string) (string, error){
    var handler io.ReadCloser
    var err error
    if f == nil {
      return "", errors.New("file_error:empty file")
    }
      if handler, err = f.Open();err == nil {
        defer handler.Close()
        dir, err := os.Getwd()
        if err != nil {
          return "", err
        }
        filename := f.Filename
        randName := randomstr.RandStringRunes(14)
        if randName != "" {
          filename = fmt.Sprintf("%s%s", randName, filepath.Ext(f.Filename))
        }
        fileLocation := filepath.Join(dir, t, filename)
        targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
        defer targetFile.Close()
        if err != nil {
          return "", err
        }
        if _, err := io.Copy(targetFile, handler); err != nil {
          return "", err
        }
        return filename, nil
      }
    log.Println(err)
    return "", err
}
