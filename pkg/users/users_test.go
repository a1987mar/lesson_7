package users

import (
	"lesson4/pkg/documentstore"
	"reflect"
	"testing"
)

func BenchmarkCreateUser(b *testing.B) {

	for i := 0; i < b.N; i++ {
		s := NewService()

		d1 := documentstore.Document{
			Fields: map[string]documentstore.DocumentField{
				"id1": {
					Type:  documentstore.DocumentFieldTypeString,
					Value: "user1",
				},
			},
		}
		s.CreateUser("id1", "user1", documentstore.CollectionConfig{PrimaryKey: "id1"}, &d1)
	}
}

func TestService_CreateUser(t *testing.T) {
	type args struct {
		id   string
		name string
		cfg  documentstore.CollectionConfig
		doc  *documentstore.Document
	}
	tests := []struct {
		name    string
		s       *Service
		args    args
		want    *User
		wantErr bool
	}{
		{
			name: "Create collection with valid doc",
			s:    NewService(),
			args: args{
				id:   "id1",
				name: "user1",
				cfg:  documentstore.CollectionConfig{PrimaryKey: "id1"},
				doc: &documentstore.Document{
					Fields: map[string]documentstore.DocumentField{
						"id1": {
							Type:  documentstore.DocumentFieldTypeString,
							Value: "user1",
						},
					},
				},
			},
			want: &User{
				ID:   "id1",
				Name: "user1",
				Cfg:  documentstore.CollectionConfig{PrimaryKey: "id1"},
			},
			wantErr: false,
		},
		{
			name: "Create collection with not valid doc",
			s:    NewService(),
			args: args{
				id:   "id1",
				name: "user1",
				cfg:  documentstore.CollectionConfig{PrimaryKey: "id1"},
				doc: &documentstore.Document{
					Fields: map[string]documentstore.DocumentField{
						"id1": {
							Type:  documentstore.DocumentFieldTypeString,
							Value: "user1",
						},
					},
				},
			},
			want: &User{
				ID:   "id1",
				Name: "user1",
				Cfg:  documentstore.CollectionConfig{PrimaryKey: "id1"},
			},
			wantErr: false,
		},
		{
			name: "Create collection with user nil",
			s:    NewService(),
			args: args{
				doc: &documentstore.Document{
					Fields: map[string]documentstore.DocumentField{},
				},
			},
			want:    &User{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.CreateUser(tt.args.id, tt.args.name, tt.args.cfg, tt.args.doc)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetUser(t *testing.T) {
	type args struct {
		userID string
	}
	tests := []struct {
		name    string
		s       *Service
		args    args
		want    *User
		wantErr bool
	}{
		{
			name: "user exists",
			s: &Service{
				users: map[string]User{
					"id1": {
						ID:   "id1",
						Name: "John",
						Cfg:  documentstore.CollectionConfig{PrimaryKey: "id"},
					},
				}},
			args: args{
				userID: "id1",
			},
			want: &User{
				ID:   "id1",
				Name: "John",
				Cfg:  documentstore.CollectionConfig{PrimaryKey: "id"},
			},
			wantErr: false,
		},
		{
			name: "user does not exist",
			s: &Service{
				users: map[string]User{},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetUser(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_DeleteUser(t *testing.T) {
	type args struct {
		userID string
	}
	tests := []struct {
		name    string
		s       *Service
		args    args
		wantErr bool
	}{
		{
			name: "user exists",
			s: &Service{
				users: map[string]User{
					"id1": {
						ID:   "id1",
						Name: "John",
						Cfg:  documentstore.CollectionConfig{PrimaryKey: "id"},
					},
				}},
			args: args{
				userID: "id1",
			},
			wantErr: false,
		},
		{
			name: "user does not exist",
			s:    &Service{},
			args: args{
				userID: "id1",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.DeleteUser(tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("Service.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_ListUsers(t *testing.T) {
	tests := []struct {
		name    string
		s       *Service
		want    []User
		wantErr bool
	}{
		{
			name: "valid List with user",
			s: &Service{
				users: map[string]User{
					"id1": {
						ID:   "id1",
						Name: "John",
						Cfg:  documentstore.CollectionConfig{PrimaryKey: "id1"},
					},
					"id2": {
						ID:   "id2",
						Name: "Smit",
						Cfg:  documentstore.CollectionConfig{PrimaryKey: "id2"},
					},
					"id3": {
						ID:   "id3",
						Name: "Jek",
						Cfg:  documentstore.CollectionConfig{PrimaryKey: "id3"},
					},
				},
			},
			want: []User{
				{ID: "id1",
					Name: "John",
					Cfg:  documentstore.CollectionConfig{PrimaryKey: "id1"}},
				{ID: "id2",
					Name: "Smit",
					Cfg:  documentstore.CollectionConfig{PrimaryKey: "id2"}},
				{ID: "id3",
					Name: "Jek",
					Cfg:  documentstore.CollectionConfig{PrimaryKey: "id3"}},
			},
			wantErr: false,
		},
		{
			name:    "empty user list",
			s:       &Service{users: map[string]User{}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ListUsers()
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.ListUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.ListUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}
