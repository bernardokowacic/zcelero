package repository

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"testing"
	"zcelero/helper"
	mockhelper "zcelero/mocks/helper"
)

func Test_textManagementRepositoryStruct_Save(t *testing.T) {
	fileLocation = t.TempDir()
	type fields struct {
		Helper helper.HelperInterface
	}
	type args struct {
		fileName string
		content  string
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		mockBehavior   func(f fields, a args)
		assertBehavior func(t *testing.T, f fields)
		wantErr        bool
	}{
		{
			name:   "Save new file",
			fields: fields{&mockhelper.HelperInterface{}},
			args: args{
				fileName: "47b416d1-c5f2-417e-929e-7b83667c6654",
				content:  `{content:"base64",encrypted:true}`,
			},
			mockBehavior: func(f fields, a args) {
				file := os.NewFile(uintptr(10), fmt.Sprintf("%s/%s.json", fileLocation, a.fileName))
				f.Helper.(*mockhelper.HelperInterface).On("CreateFile", fmt.Sprintf("%s/%s.json", fileLocation, a.fileName)).Return(file, nil)
				f.Helper.(*mockhelper.HelperInterface).On("WriteFile", file, a.content).Return(len([]byte(a.content)), nil)
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.Helper.(*mockhelper.HelperInterface).AssertExpectations(t)
			},
			wantErr: false,
		},
		{
			name:   "Save new file with creation error",
			fields: fields{&mockhelper.HelperInterface{}},
			args: args{
				fileName: "nil",
				content:  `{content:"base64",encrypted:true}`,
			},
			mockBehavior: func(f fields, a args) {
				f.Helper.(*mockhelper.HelperInterface).On("CreateFile", fmt.Sprintf("%s/%s.json", fileLocation, a.fileName)).Return(nil, errors.New("error"))
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.Helper.(*mockhelper.HelperInterface).AssertExpectations(t)
			},
			wantErr: true,
		},
		{
			name:   "Save new file with write error",
			fields: fields{&mockhelper.HelperInterface{}},
			args: args{
				fileName: "47b416d1-c5f2-417e-929e-7b83667c6654",
				content:  `{content:"base64",encrypted:true}`,
			},
			mockBehavior: func(f fields, a args) {
				file := os.NewFile(uintptr(10), fmt.Sprintf("%s/%s.json", fileLocation, a.fileName))
				f.Helper.(*mockhelper.HelperInterface).On("CreateFile", fmt.Sprintf("%s/%s.json", fileLocation, a.fileName)).Return(file, nil)
				f.Helper.(*mockhelper.HelperInterface).On("WriteFile", file, a.content).Return(0, errors.New("error"))
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.Helper.(*mockhelper.HelperInterface).AssertExpectations(t)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockBehavior != nil {
				tt.mockBehavior(tt.fields, tt.args)
			}

			repository := NewRepository(tt.fields.Helper)
			if err := repository.Save(tt.args.fileName, tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("textManagementRepositoryStruct.Save() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.assertBehavior != nil {
				tt.assertBehavior(t, tt.fields)
			}
		})
	}
}

func Test_textManagementRepositoryStruct_Load(t *testing.T) {
	content := `{content:"base64",encrypted:true}`
	type fields struct {
		Helper helper.HelperInterface
	}
	type args struct {
		fileName string
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		mockBehavior   func(f fields, a args)
		assertBehavior func(t *testing.T, f fields)
		want           []byte
		wantErr        bool
	}{
		{
			name:   "Load file",
			fields: fields{&mockhelper.HelperInterface{}},
			args: args{
				fileName: "47b416d1-c5f2-417e-929e-7b83667c6654",
			},
			mockBehavior: func(f fields, a args) {
				f.Helper.(*mockhelper.HelperInterface).On("ReadFile", fmt.Sprintf("%s/%s.json", fileLocation, a.fileName)).Return([]byte(content), nil)
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.Helper.(*mockhelper.HelperInterface).AssertExpectations(t)
			},
			want:    []byte(content),
			wantErr: false,
		},
		{
			name:   "Load file with error",
			fields: fields{&mockhelper.HelperInterface{}},
			args: args{
				fileName: "47b416d1-c5f2-417e-929e-7b83667c6654",
			},
			mockBehavior: func(f fields, a args) {
				f.Helper.(*mockhelper.HelperInterface).On("ReadFile", fmt.Sprintf("%s/%s.json", fileLocation, a.fileName)).Return(nil, errors.New("error"))
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.Helper.(*mockhelper.HelperInterface).AssertExpectations(t)
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockBehavior != nil {
				tt.mockBehavior(tt.fields, tt.args)
			}

			repository := NewRepository(tt.fields.Helper)
			got, err := repository.Load(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("textManagementRepositoryStruct.Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("textManagementRepositoryStruct.Load() = %v, want %v", got, tt.want)
			}

			if tt.assertBehavior != nil {
				tt.assertBehavior(t, tt.fields)
			}
		})
	}
}
