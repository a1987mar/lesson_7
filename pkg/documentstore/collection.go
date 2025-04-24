package documentstore

import (
	"lesson4/pkg/err"
	"log/slog"
)

type Collection struct {
	Documents map[string]Document `json:"documents,omitempty"`
	Config    CollectionConfig    `json:"config"`
}

type CollectionConfig struct {
	PrimaryKey string `json:"cgg"`
}

func (s *Collection) Put(doc Document) error {
	// Потрібно перевірити що документ містить поле `{cfg.PrimaryKey}` типу `string`

	keyFilds, ok := doc.Fields[s.Config.PrimaryKey]
	if !ok {
		slog.Error("error: Document must contain a key field")
		return err.ErrUnsupportedDocumentField
	}

	if keyFilds.Type != DocumentFieldTypeString {
		slog.Error("error: Key field must be of type string")
		return err.ErrUnsupportedDocumentField
	}
	if s.Documents == nil {
		s.Documents = map[string]Document{}
	}
	s.Documents[s.Config.PrimaryKey] = doc
	slog.Info("document added")
	return nil
}

func (s *Collection) Get(key string) (*Document, error) {
	if doc, exists := s.Documents[key]; exists {
		return &doc, nil
	}
	slog.Info("document not found")
	return nil, err.ErrDocumentNotFound
}

func (s *Collection) Delete(key string) bool {
	if _, exists := s.Documents[key]; exists {
		delete(s.Documents, key)
		slog.Info("document delete")
		return true
	}
	return false
}

func (s *Collection) List() []Document {
	sList := make([]Document, 0, len(s.Documents))
	for _, v := range s.Documents {
		sList = append(sList, v)
	}
	return sList
}
