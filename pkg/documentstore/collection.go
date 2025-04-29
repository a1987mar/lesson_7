package documentstore

import (
	"lesson4/pkg/err"
	"log/slog"
)

type Collection struct {
	documents map[string]Document
	config    CollectionConfig
}
type DTOCollection struct {
	Documents map[string]Document `json:"documents,omitempty"`
	Config    CollectionConfig    `json:"config"`
}

func (s *Collection) ToDto() DTOCollection {
	return DTOCollection{
		Documents: s.documents,
		Config:    s.config,
	}
}

type CollectionConfig struct {
	PrimaryKey string `json:"cgg"`
}

func (s *Collection) Put(doc Document) error {
	// Потрібно перевірити що документ містить поле `{cfg.PrimaryKey}` типу `string`

	keyFilds, ok := doc.Fields[s.config.PrimaryKey]
	if !ok {
		slog.Error("error: Document must contain a key field")
		return err.ErrUnsupportedDocumentField
	}

	if keyFilds.Type != DocumentFieldTypeString {
		slog.Error("error: Key field must be of type string")
		return err.ErrUnsupportedDocumentField
	}
	keyValue, ok := keyFilds.Value.(string)
	if !ok {
		slog.Error("Error: Key field value is not a string")
		return err.ErrUnsupportedDocumentField
	}
	if s.documents == nil {
		s.documents = map[string]Document{}
	}
	s.documents[keyValue] = doc
	slog.Info("document added")
	return nil
}

func (s *Collection) Get(key string) (*Document, error) {
	if doc, exists := s.documents[key]; exists {
		return &doc, nil
	}
	slog.Info("document not found")
	return nil, err.ErrDocumentNotFound
}

func (s *Collection) Delete(key string) bool {
	if _, exists := s.documents[key]; exists {
		delete(s.documents, key)
		slog.Info("document delete")
		return true
	}
	return false
}

func (s *Collection) List() []Document {
	sList := make([]Document, 0, len(s.documents))
	for _, v := range s.documents {
		sList = append(sList, v)
	}
	return sList
}
