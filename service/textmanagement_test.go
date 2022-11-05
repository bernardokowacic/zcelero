package service_test

import (
	"errors"
	"reflect"
	"testing"
	"zcelero/entity"
	mockrespository "zcelero/mocks/repository"
	"zcelero/repository"
	"zcelero/service"
)

func TestTextManagementService_Get(t *testing.T) {
	type fields struct {
		TextManagementRepository repository.TextManagementInterface
	}
	type args struct {
		textId           string
		privateKeyString string
		password         string
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		mockBehavior   func(f fields, a args)
		assertBehavior func(t *testing.T, f fields)
		want           string
		wantErr        bool
	}{
		{
			name:   "Get encrypted content",
			fields: fields{&mockrespository.TextManagementInterface{}},
			args: args{
				textId:           "2f13ed58-afc9-477a-bf0d-c90eb1b7db90",
				privateKeyString: "-----BEGIN RSA PRIVATE KEY-----\nProc-Type: 4,ENCRYPTED\nDEK-Info: AES-256-CBC,53ae7dd8c61749faa0dcc2ace0dff6cc\n\nH6QV4F6XfP2EJ4mebb0t2YZKDwK6HBJXO2bdgVKlkuOepf+LBLJWsF88Dbk1PYn8\nXGLTKBGEhtK3Qy3+IeCeWUQOxdEKc2RJKMLl5Hs1uZm6okvJSdECNBHlO4PC2WDC\nuLT+CsFWnVNQH/WT16EpUQOu0IU7GugnOqIPm1z0U4XtprX8Xcw1uNl9uyDEEUxg\nvwGEc2Qi4RpcAbK1bnPUQ6BjQt0DYiOyEJZFeTTkXYCTsnifTIcyp19ItLYzXjDz\nbsRJX52jCj66dEHmAiXq8u5vpU55Y7Gt73CjBUf3mcNMMzzdhurJeKXCXkcWFXb7\nj2sbJDdEjyLGAlE4Fxeyn7hpmj8yo01oNFcpJYgoEgp+AdYgvjhEykYFqm4NzP/Y\nxmyDKU3rqO1nl9jOckzGLq8Ca16psLgXbcrpkfxv9dik5HiAUaaBZ3bINkrisoRg\nd7by7ig03sDlQWMU9P2b0UZx5jD/MUB06dDIAK7MtQ2mGTs99PB/42IO6Xv1/PYn\nObMNgmuZwD3sW6oU331KrYLyN97LSJ5hhop5s5y/kdapvwVLhToUa8c4VRTHde8x\nN8XZyKWJRxZHXRx9FToXMkHEGhT7ei8QUvuKQ7xFsmtbMDm8FQv/O3gs5GCI3MKw\nnGQZOxLvP31ZZ0un+7CS0HuyGl4KowHCQvvqKxMn4kLGa3qCtJCjKRNVoasr78+m\nCNGN2IlaCJZsBHYNEej6/vP7TJiYoNQQupFqkzz9w90nPhxkGU6dnPH78pOcWBAY\nU4cMRsJqezGkXHph9Z6j8fZyqdo+zNpyRQIhR84V1xDiKBsK6F5pInT1PsjJRyQ7\n-----END RSA PRIVATE KEY-----\n",
				password:         "aaa",
			},
			mockBehavior: func(f fields, a args) {
				fileContent := `{"Content":"G+DVq/1yfH4+hOhSnVYwiS0FK7SFADA65C6kPKKOR2OG6qXe0F/C1OKSRjm3CKQyY2cchCs0WyZopDZwTFsLMmS+GPM842FbBm/5KSbiNXeS0PsmoQW2RmVVV7iCJDeiNlzJHkdqDcZU/VwHj0FW0Z9nfDpXeBQyKt1Wx1WUXRI=","Encrypted":true}`
				f.TextManagementRepository.(*mockrespository.TextManagementInterface).On("Load", a.textId).Return([]byte(fileContent), nil)
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.TextManagementRepository.(*mockrespository.TextManagementInterface).AssertExpectations(t)
			},
			want:    "aaaaaaaa",
			wantErr: false,
		},
		{
			name:   "Get non encrypted content",
			fields: fields{&mockrespository.TextManagementInterface{}},
			args: args{
				textId:           "2f13ed58-afc9-477a-bf0d-c90eb1b7db90",
				privateKeyString: "",
				password:         "",
			},
			mockBehavior: func(f fields, a args) {
				fileContent := `{"Content":"aaaaaaaa","Encrypted":false}`
				f.TextManagementRepository.(*mockrespository.TextManagementInterface).On("Load", a.textId).Return([]byte(fileContent), nil)
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.TextManagementRepository.(*mockrespository.TextManagementInterface).AssertExpectations(t)
			},
			want:    "aaaaaaaa",
			wantErr: false,
		},
		{
			name:   "Error when reading file",
			fields: fields{&mockrespository.TextManagementInterface{}},
			args: args{
				textId:           "2f13ed58-afc9-477a-bf0d-c90eb1b7db90",
				privateKeyString: "",
				password:         "",
			},
			mockBehavior: func(f fields, a args) {
				f.TextManagementRepository.(*mockrespository.TextManagementInterface).On("Load", a.textId).Return(nil, errors.New("error"))
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.TextManagementRepository.(*mockrespository.TextManagementInterface).AssertExpectations(t)
			},
			want:    "",
			wantErr: true,
		},
		{
			name:   "Get encrypted content without private key",
			fields: fields{&mockrespository.TextManagementInterface{}},
			args: args{
				textId:           "2f13ed58-afc9-477a-bf0d-c90eb1b7db90",
				privateKeyString: "",
				password:         "aaa",
			},
			mockBehavior: func(f fields, a args) {
				fileContent := `{"Content":"G+DVq/1yfH4+hOhSnVYwiS0FK7SFADA65C6kPKKOR2OG6qXe0F/C1OKSRjm3CKQyY2cchCs0WyZopDZwTFsLMmS+GPM842FbBm/5KSbiNXeS0PsmoQW2RmVVV7iCJDeiNlzJHkdqDcZU/VwHj0FW0Z9nfDpXeBQyKt1Wx1WUXRI=","Encrypted":true}`
				f.TextManagementRepository.(*mockrespository.TextManagementInterface).On("Load", a.textId).Return([]byte(fileContent), nil)
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.TextManagementRepository.(*mockrespository.TextManagementInterface).AssertExpectations(t)
			},
			want:    "",
			wantErr: true,
		},
		{
			name:   "Get encrypted content without password",
			fields: fields{&mockrespository.TextManagementInterface{}},
			args: args{
				textId:           "2f13ed58-afc9-477a-bf0d-c90eb1b7db90",
				privateKeyString: "-----BEGIN RSA PRIVATE KEY-----\nProc-Type: 4,ENCRYPTED\nDEK-Info: AES-256-CBC,53ae7dd8c61749faa0dcc2ace0dff6cc\n\nH6QV4F6XfP2EJ4mebb0t2YZKDwK6HBJXO2bdgVKlkuOepf+LBLJWsF88Dbk1PYn8\nXGLTKBGEhtK3Qy3+IeCeWUQOxdEKc2RJKMLl5Hs1uZm6okvJSdECNBHlO4PC2WDC\nuLT+CsFWnVNQH/WT16EpUQOu0IU7GugnOqIPm1z0U4XtprX8Xcw1uNl9uyDEEUxg\nvwGEc2Qi4RpcAbK1bnPUQ6BjQt0DYiOyEJZFeTTkXYCTsnifTIcyp19ItLYzXjDz\nbsRJX52jCj66dEHmAiXq8u5vpU55Y7Gt73CjBUf3mcNMMzzdhurJeKXCXkcWFXb7\nj2sbJDdEjyLGAlE4Fxeyn7hpmj8yo01oNFcpJYgoEgp+AdYgvjhEykYFqm4NzP/Y\nxmyDKU3rqO1nl9jOckzGLq8Ca16psLgXbcrpkfxv9dik5HiAUaaBZ3bINkrisoRg\nd7by7ig03sDlQWMU9P2b0UZx5jD/MUB06dDIAK7MtQ2mGTs99PB/42IO6Xv1/PYn\nObMNgmuZwD3sW6oU331KrYLyN97LSJ5hhop5s5y/kdapvwVLhToUa8c4VRTHde8x\nN8XZyKWJRxZHXRx9FToXMkHEGhT7ei8QUvuKQ7xFsmtbMDm8FQv/O3gs5GCI3MKw\nnGQZOxLvP31ZZ0un+7CS0HuyGl4KowHCQvvqKxMn4kLGa3qCtJCjKRNVoasr78+m\nCNGN2IlaCJZsBHYNEej6/vP7TJiYoNQQupFqkzz9w90nPhxkGU6dnPH78pOcWBAY\nU4cMRsJqezGkXHph9Z6j8fZyqdo+zNpyRQIhR84V1xDiKBsK6F5pInT1PsjJRyQ7\n-----END RSA PRIVATE KEY-----\n",
				password:         "",
			},
			mockBehavior: func(f fields, a args) {
				fileContent := `{"Content":"G+DVq/1yfH4+hOhSnVYwiS0FK7SFADA65C6kPKKOR2OG6qXe0F/C1OKSRjm3CKQyY2cchCs0WyZopDZwTFsLMmS+GPM842FbBm/5KSbiNXeS0PsmoQW2RmVVV7iCJDeiNlzJHkdqDcZU/VwHj0FW0Z9nfDpXeBQyKt1Wx1WUXRI=","Encrypted":true}`
				f.TextManagementRepository.(*mockrespository.TextManagementInterface).On("Load", a.textId).Return([]byte(fileContent), nil)
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.TextManagementRepository.(*mockrespository.TextManagementInterface).AssertExpectations(t)
			},
			want:    "",
			wantErr: true,
		},
		{
			name:   "Get encrypted content with wrong private key",
			fields: fields{&mockrespository.TextManagementInterface{}},
			args: args{
				textId:           "2f13ed58-afc9-477a-bf0d-c90eb1b7db90",
				privateKeyString: "-----BEGIN RSA PRIVATE KEY-----\nProc-Type: 4,ENCRYPTED\nDEK-Info: AES-256-CBC,53ae7dd8c61749faa0dcc2ace0dff6cc\nXGLTKBGEhtK3Qy3+IeCeWUQOxdEKc2RJKMLl5Hs1uZm6okvJSdECNBHlO4PC2WDC\nuLT+CsFWnVNQH/WT16EpUQOu0IU7GugnOqIPm1z0U4XtprX8Xcw1uNl9uyDEEUxg\nvwGEc2Qi4RpcAbK1bnPUQ6BjQt0DYiOyEJZFeTTkXYCTsnifTIcyp19ItLYzXjDz\nbsRJX52jCj66dEHmAiXq8u5vpU55Y7Gt73CjBUf3mcNMMzzdhurJeKXCXkcWFXb7\nj2sbJDdEjyLGAlE4Fxeyn7hpmj8yo01oNFcpJYgoEgp+AdYgvjhEykYFqm4NzP/Y\nxmyDKU3rqO1nl9jOckzGLq8Ca16psLgXbcrpkfxv9dik5HiAUaaBZ3bINkrisoRg\nd7by7ig03sDlQWMU9P2b0UZx5jD/MUB06dDIAK7MtQ2mGTs99PB/42IO6Xv1/PYn\nObMNgmuZwD3sW6oU331KrYLyN97LSJ5hhop5s5y/kdapvwVLhToUa8c4VRTHde8x\nN8XZyKWJRxZHXRx9FToXMkHEGhT7ei8QUvuKQ7xFsmtbMDm8FQv/O3gs5GCI3MKw\nnGQZOxLvP31ZZ0un+7CS0HuyGl4KowHCQvvqKxMn4kLGa3qCtJCjKRNVoasr78+m\nCNGN2IlaCJZsBHYNEej6/vP7TJiYoNQQupFqkzz9w90nPhxkGU6dnPH78pOcWBAY\nU4cMRsJqezGkXHph9Z6j8fZyqdo+zNpyRQIhR84V1xDiKBsK6F5pInT1PsjJRyQ7\n-----END RSA PRIVATE KEY-----\n",
				password:         "aaa",
			},
			mockBehavior: func(f fields, a args) {
				fileContent := `{"Content":"G+DVq/1yfH4+hOhSnVYwiS0FK7SFADA65C6kPKKOR2OG6qXe0F/C1OKSRjm3CKQyY2cchCs0WyZopDZwTFsLMmS+GPM842FbBm/5KSbiNXeS0PsmoQW2RmVVV7iCJDeiNlzJHkdqDcZU/VwHj0FW0Z9nfDpXeBQyKt1Wx1WUXRI=","Encrypted":true}`
				f.TextManagementRepository.(*mockrespository.TextManagementInterface).On("Load", a.textId).Return([]byte(fileContent), nil)
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.TextManagementRepository.(*mockrespository.TextManagementInterface).AssertExpectations(t)
			},
			want:    "",
			wantErr: true,
		},
		{
			name:   "Get encrypted content with wrong password",
			fields: fields{&mockrespository.TextManagementInterface{}},
			args: args{
				textId:           "2f13ed58-afc9-477a-bf0d-c90eb1b7db90",
				privateKeyString: "-----BEGIN RSA PRIVATE KEY-----\nProc-Type: 4,ENCRYPTED\nDEK-Info: AES-256-CBC,53ae7dd8c61749faa0dcc2ace0dff6cc\n\nH6QV4F6XfP2EJ4mebb0t2YZKDwK6HBJXO2bdgVKlkuOepf+LBLJWsF88Dbk1PYn8\nXGLTKBGEhtK3Qy3+IeCeWUQOxdEKc2RJKMLl5Hs1uZm6okvJSdECNBHlO4PC2WDC\nuLT+CsFWnVNQH/WT16EpUQOu0IU7GugnOqIPm1z0U4XtprX8Xcw1uNl9uyDEEUxg\nvwGEc2Qi4RpcAbK1bnPUQ6BjQt0DYiOyEJZFeTTkXYCTsnifTIcyp19ItLYzXjDz\nbsRJX52jCj66dEHmAiXq8u5vpU55Y7Gt73CjBUf3mcNMMzzdhurJeKXCXkcWFXb7\nj2sbJDdEjyLGAlE4Fxeyn7hpmj8yo01oNFcpJYgoEgp+AdYgvjhEykYFqm4NzP/Y\nxmyDKU3rqO1nl9jOckzGLq8Ca16psLgXbcrpkfxv9dik5HiAUaaBZ3bINkrisoRg\nd7by7ig03sDlQWMU9P2b0UZx5jD/MUB06dDIAK7MtQ2mGTs99PB/42IO6Xv1/PYn\nObMNgmuZwD3sW6oU331KrYLyN97LSJ5hhop5s5y/kdapvwVLhToUa8c4VRTHde8x\nN8XZyKWJRxZHXRx9FToXMkHEGhT7ei8QUvuKQ7xFsmtbMDm8FQv/O3gs5GCI3MKw\nnGQZOxLvP31ZZ0un+7CS0HuyGl4KowHCQvvqKxMn4kLGa3qCtJCjKRNVoasr78+m\nCNGN2IlaCJZsBHYNEej6/vP7TJiYoNQQupFqkzz9w90nPhxkGU6dnPH78pOcWBAY\nU4cMRsJqezGkXHph9Z6j8fZyqdo+zNpyRQIhR84V1xDiKBsK6F5pInT1PsjJRyQ7\n-----END RSA PRIVATE KEY-----\n",
				password:         "bbb",
			},
			mockBehavior: func(f fields, a args) {
				fileContent := `{"Content":"G+DVq/1yfH4+hOhSnVYwiS0FK7SFADA65C6kPKKOR2OG6qXe0F/C1OKSRjm3CKQyY2cchCs0WyZopDZwTFsLMmS+GPM842FbBm/5KSbiNXeS0PsmoQW2RmVVV7iCJDeiNlzJHkdqDcZU/VwHj0FW0Z9nfDpXeBQyKt1Wx1WUXRI=","Encrypted":true}`
				f.TextManagementRepository.(*mockrespository.TextManagementInterface).On("Load", a.textId).Return([]byte(fileContent), nil)
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.TextManagementRepository.(*mockrespository.TextManagementInterface).AssertExpectations(t)
			},
			want:    "",
			wantErr: true,
		},
		{
			name:   "Get encrypted empty content",
			fields: fields{&mockrespository.TextManagementInterface{}},
			args: args{
				textId:           "2f13ed58-afc9-477a-bf0d-c90eb1b7db90",
				privateKeyString: "-----BEGIN RSA PRIVATE KEY-----\nProc-Type: 4,ENCRYPTED\nDEK-Info: AES-256-CBC,53ae7dd8c61749faa0dcc2ace0dff6cc\n\nH6QV4F6XfP2EJ4mebb0t2YZKDwK6HBJXO2bdgVKlkuOepf+LBLJWsF88Dbk1PYn8\nXGLTKBGEhtK3Qy3+IeCeWUQOxdEKc2RJKMLl5Hs1uZm6okvJSdECNBHlO4PC2WDC\nuLT+CsFWnVNQH/WT16EpUQOu0IU7GugnOqIPm1z0U4XtprX8Xcw1uNl9uyDEEUxg\nvwGEc2Qi4RpcAbK1bnPUQ6BjQt0DYiOyEJZFeTTkXYCTsnifTIcyp19ItLYzXjDz\nbsRJX52jCj66dEHmAiXq8u5vpU55Y7Gt73CjBUf3mcNMMzzdhurJeKXCXkcWFXb7\nj2sbJDdEjyLGAlE4Fxeyn7hpmj8yo01oNFcpJYgoEgp+AdYgvjhEykYFqm4NzP/Y\nxmyDKU3rqO1nl9jOckzGLq8Ca16psLgXbcrpkfxv9dik5HiAUaaBZ3bINkrisoRg\nd7by7ig03sDlQWMU9P2b0UZx5jD/MUB06dDIAK7MtQ2mGTs99PB/42IO6Xv1/PYn\nObMNgmuZwD3sW6oU331KrYLyN97LSJ5hhop5s5y/kdapvwVLhToUa8c4VRTHde8x\nN8XZyKWJRxZHXRx9FToXMkHEGhT7ei8QUvuKQ7xFsmtbMDm8FQv/O3gs5GCI3MKw\nnGQZOxLvP31ZZ0un+7CS0HuyGl4KowHCQvvqKxMn4kLGa3qCtJCjKRNVoasr78+m\nCNGN2IlaCJZsBHYNEej6/vP7TJiYoNQQupFqkzz9w90nPhxkGU6dnPH78pOcWBAY\nU4cMRsJqezGkXHph9Z6j8fZyqdo+zNpyRQIhR84V1xDiKBsK6F5pInT1PsjJRyQ7\n-----END RSA PRIVATE KEY-----\n",
				password:         "aaa",
			},
			mockBehavior: func(f fields, a args) {
				fileContent := `{"Content":"","Encrypted":true}`
				f.TextManagementRepository.(*mockrespository.TextManagementInterface).On("Load", a.textId).Return([]byte(fileContent), nil)
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.TextManagementRepository.(*mockrespository.TextManagementInterface).AssertExpectations(t)
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockBehavior != nil {
				tt.mockBehavior(tt.fields, tt.args)
			}

			service := service.NewService(tt.fields.TextManagementRepository)

			got, err := service.Get(tt.args.textId, tt.args.privateKeyString, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("TextManagementService.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TextManagementService.Get() = %v, want %v", got, tt.want)
			}

			if tt.assertBehavior != nil {
				tt.assertBehavior(t, tt.fields)
			}
		})
	}
}

func TestTextManagementService_Insert(t *testing.T) {
	type fields struct {
		TextManagementRepository repository.TextManagementInterface
	}
	type args struct {
		text entity.TextManagement
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		mockBehavior   func(f fields, a args)
		assertBehavior func(t *testing.T, f fields)
		want           entity.TextManagement
		wantErr        bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockBehavior != nil {
				tt.mockBehavior(tt.fields, tt.args)
			}

			service := service.NewService(tt.fields.TextManagementRepository)

			got, err := service.Insert(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("TextManagementService.Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TextManagementService.Insert() = %v, want %v", got, tt.want)
			}

			if tt.assertBehavior != nil {
				tt.assertBehavior(t, tt.fields)
			}
		})
	}
}
