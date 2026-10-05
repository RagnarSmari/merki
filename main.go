package merki

import (
	"github.com/RagnarSmari/merki/internal/convert"
	"github.com/RagnarSmari/merki/internal/model"
	"github.com/RagnarSmari/merki/internal/preview"
)

func ConvertToZPL(data model.Label, vars map[string]string) (string, error) {
	return convert.ToZpl(data, vars)
}

func GetPngPreview(data model.Label, vars map[string]string) ([]byte, error){
	return preview.PreviewInBytes(data, vars)
}
