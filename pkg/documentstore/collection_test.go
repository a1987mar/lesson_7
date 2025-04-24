package documentstore

import (
	"reflect"
	"testing"
)

func TestCollection_Put(t *testing.T) {
	type fields struct {
		Documents map[string]Document
		Config    CollectionConfig
	}
	type args struct {
		doc Document
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "valid document with correct primary key",
			fields: fields{
				Documents: map[string]Document{},
				Config:    CollectionConfig{PrimaryKey: "id"},
			},
			args: args{
				doc: Document{
					Fields: map[string]DocumentField{
						"id": {
							Type:  DocumentFieldTypeString,
							Value: "123",
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "missing primary key field",
			fields: fields{
				Documents: map[string]Document{},
				Config:    CollectionConfig{PrimaryKey: "id"},
			},
			args: args{
				doc: Document{
					Fields: map[string]DocumentField{
						"name": {
							Type:  DocumentFieldTypeString,
							Value: "Test",
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "primary key not string type",
			fields: fields{
				Documents: map[string]Document{},
				Config:    CollectionConfig{PrimaryKey: "id"},
			},
			args: args{
				doc: Document{
					Fields: map[string]DocumentField{
						"id": {
							Type:  DocumentFieldTypeNumber,
							Value: 123,
						},
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Collection{
				Documents: tt.fields.Documents,
				Config:    tt.fields.Config,
			}
			if err := s.Put(tt.args.doc); (err != nil) != tt.wantErr {
				t.Errorf("Put() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCollection_Get(t *testing.T) {
	type fields struct {
		Documents map[string]Document
		Config    CollectionConfig
	}
	type args struct {
		doc Document
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "valid document with correct primary key",
			fields: fields{
				Documents: map[string]Document{},
				Config:    CollectionConfig{PrimaryKey: "id"},
			},
			args: args{
				doc: Document{
					Fields: map[string]DocumentField{
						"id": {
							Type:  DocumentFieldTypeString,
							Value: "123",
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "missing primary key field",
			fields: fields{
				Documents: map[string]Document{},
				Config:    CollectionConfig{PrimaryKey: "id"},
			},
			args: args{
				doc: Document{
					Fields: map[string]DocumentField{
						"name": {
							Type:  DocumentFieldTypeString,
							Value: "Test User",
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "primary key field is not string type",
			fields: fields{
				Documents: map[string]Document{},
				Config:    CollectionConfig{PrimaryKey: "id"},
			},
			args: args{
				doc: Document{
					Fields: map[string]DocumentField{
						"id": {
							Type:  DocumentFieldTypeNumber, // not string!
							Value: 123,
						},
					},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Collection{
				Documents: tt.fields.Documents,
				Config:    tt.fields.Config,
			}
			if err := s.Put(tt.args.doc); (err != nil) != tt.wantErr {
				t.Errorf("Put() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCollection_Delete(t *testing.T) {
	type fields struct {
		Documents map[string]Document
		Config    CollectionConfig
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "valid document with correct primary key",
			fields: fields{
				Documents: map[string]Document{},
				Config:    CollectionConfig{PrimaryKey: "id"},
			},
			args: args{
				key: "id",
			},
		},
		{
			name: "missing primary key field",
			fields: fields{
				Documents: map[string]Document{},
				Config:    CollectionConfig{PrimaryKey: "id"},
			},
			args: args{
				"id",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Collection{
				Documents: tt.fields.Documents,
				Config:    tt.fields.Config,
			}
			if got := s.Delete(tt.args.key); got != tt.want {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCollection_List(t *testing.T) {
	tests := []struct {
		name string
		s    *Collection
		want []Document
	}{
		{name: "valid List with correct documnet",
			s: &Collection{
				Documents: map[string]Document{},
				Config: CollectionConfig{
					PrimaryKey: "id-1",
				},
			},
			want: []Document{},
		},

		{
			name: "collection with one document",
			s: &Collection{
				Documents: map[string]Document{
					"doc1": {
						Fields: map[string]DocumentField{
							"id": {Type: DocumentFieldTypeString, Value: "123"},
						},
					},
				},
				Config: CollectionConfig{
					PrimaryKey: "id",
				},
			},
			want: []Document{
				{
					Fields: map[string]DocumentField{
						"id": {Type: DocumentFieldTypeString, Value: "123"},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.List(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Collection.List() = %v, want %v", got, tt.want)
			}
		})
	}
}
