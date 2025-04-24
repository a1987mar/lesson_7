package documentstore

import (
	"encoding/json"
	"fmt"
	"lesson4/pkg/err"
	"log/slog"
	"os"
	"strings"
)

type Store struct {
	Collections map[string]*Collection `json:"collections,omitempty"`
}

func NewStore() *Store {
	return &Store{
		Collections: make(map[string]*Collection),
	}
}

func (s *Store) CreateCollection(name string, cfg *CollectionConfig) (error, *Collection) {
	// Створюємо нову колекцію і повертаємо `true` якщо колекція була створена
	// Якщо ж колекція вже створеня то повертаємо `false` та nil
	if _, exists := s.Collections[name]; exists {
		slog.Error("collection already exists")
		return err.ErrCollectionAlreadyExists, nil
	}
	coll := &Collection{
		Documents: make(map[string]Document),
		Config:    *cfg}
	s.Collections[name] = coll
	slog.Info("collection added")
	return nil, coll
}

func (s *Store) GetCollection(name string) (*Collection, error) {
	if colect, ok := s.Collections[name]; ok {
		return colect, nil
	}
	slog.Error("collection not found")
	return nil, err.ErrCollectionNotFound
}

func (s *Store) DeleteCollection(name string) bool {
	if _, ok := s.Collections[name]; ok {
		delete(s.Collections, name)
		slog.Info("collection delete - %s")
		return true
	}
	return false
}

func NewStoreFromDump(dump []byte) (*Store, error) {
	// Функція повинна створити та проініціалізувати новий `Store`
	// зі всіма колекціями та даними з вхідного дампу.
	var s Store
	if err := json.Unmarshal(dump, &s); err != nil {
		return nil, err
	}
	if len(s.Collections) == 0 {
		slog.Info("collection not added")
		return nil, err.ErrNotFound
	}
	return &s, nil
}

func (s *Store) Dump() ([]byte, error) {
	// Методи повинен віддати дамп нашого стору в який включені дані про колекції та документ
	sToJson, err := json.MarshalIndent(s, " ", "")
	if err != nil {
		return nil, err
	}
	return sToJson, nil
}

//
//// Значення яке повертає метод `store.Dump()` має без помилок оброблятись функцією `NewStoreFromDump`
//

func NewStoreFromFile(filename string) (*Store, error) {
	// Робить те ж саме що і функція `NewStoreFromDump`, але сам дамп має діставатись з файлу
	file := strings.Builder{}
	file.WriteString(filename + ".json")

	//	f := fmt.Sprintf("%s.json", filename)
	dump, err := os.ReadFile(file.String())
	if err != nil {
		slog.Error("file not read")
		return nil, err
	}
	slog.Info("file read successfully")
	var s Store
	if err := json.Unmarshal(dump, &s); err != nil {

		return nil, err
	}
	if len(s.Collections) == 0 {
		slog.Error("no collections found in store from file")
		return nil, fmt.Errorf("no collections in store")
	}

	return &s, nil
}

func (s *Store) DumpToFile(filename string) error {
	// Робить те ж саме що і метод  `Dump`, але записує у файл замість того щоб повертати сам дамп
	sDump, err := s.Dump()
	if err != nil {
		fmt.Println(err)
	}
	file := strings.Builder{}
	file.WriteString(filename + ".json")
	slog.Info(file.String())
	return os.WriteFile(file.String(), sDump, 0644)
}
