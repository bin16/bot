package bot

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"
	"path"
)

type FormData struct {
	buf *bytes.Buffer
	mp  *multipart.Writer
}

func (f *FormData) ContentType() string {
	return f.mp.FormDataContentType()
}

func (f *FormData) Append(fieldName string, value string) error {
	return f.mp.WriteField(fieldName, value)
}

func (f *FormData) AppendFile(fieldName string, pathname string) error {
	fieldFile, err := os.Open(pathname)
	if err != nil {
		return err
	}
	defer fieldFile.Close()

	filename := path.Base(pathname)
	part, err := f.mp.CreateFormFile(fieldName, filename)
	if err != nil {
		return err
	}

	_, err = io.Copy(part, fieldFile)
	return err
}

func (f *FormData) Init() *FormData {
	f.buf = new(bytes.Buffer)
	f.mp = multipart.NewWriter(f.buf)
	return f
}

func (f *FormData) Done() io.Reader {
	f.mp.Close()
	return f.buf
}

func NewFormData() *FormData {
	d := &FormData{}
	return d.Init()
}
