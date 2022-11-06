package service_test

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"reflect"
	"testing"
	"zcelero/entity"
	"zcelero/helper"
	mockhelper "zcelero/mocks/helper"
	mockrepository "zcelero/mocks/repository"
	"zcelero/repository"
	"zcelero/service"
)

func TestTextManagementService_Get(t *testing.T) {
	type fields struct {
		TextManagementRepository repository.TextManagementInterface
		Helper                   helper.HelperInterface
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
			name: "Get encrypted content",
			fields: fields{
				&mockrepository.TextManagementInterface{},
				&mockhelper.HelperInterface{},
			},
			args: args{
				textId:           "2f13ed58-afc9-477a-bf0d-c90eb1b7db90",
				privateKeyString: "-----BEGIN RSA PRIVATE KEY-----\nProc-Type: 4,ENCRYPTED\nDEK-Info: AES-256-CBC,53ae7dd8c61749faa0dcc2ace0dff6cc\n\nH6QV4F6XfP2EJ4mebb0t2YZKDwK6HBJXO2bdgVKlkuOepf+LBLJWsF88Dbk1PYn8\nXGLTKBGEhtK3Qy3+IeCeWUQOxdEKc2RJKMLl5Hs1uZm6okvJSdECNBHlO4PC2WDC\nuLT+CsFWnVNQH/WT16EpUQOu0IU7GugnOqIPm1z0U4XtprX8Xcw1uNl9uyDEEUxg\nvwGEc2Qi4RpcAbK1bnPUQ6BjQt0DYiOyEJZFeTTkXYCTsnifTIcyp19ItLYzXjDz\nbsRJX52jCj66dEHmAiXq8u5vpU55Y7Gt73CjBUf3mcNMMzzdhurJeKXCXkcWFXb7\nj2sbJDdEjyLGAlE4Fxeyn7hpmj8yo01oNFcpJYgoEgp+AdYgvjhEykYFqm4NzP/Y\nxmyDKU3rqO1nl9jOckzGLq8Ca16psLgXbcrpkfxv9dik5HiAUaaBZ3bINkrisoRg\nd7by7ig03sDlQWMU9P2b0UZx5jD/MUB06dDIAK7MtQ2mGTs99PB/42IO6Xv1/PYn\nObMNgmuZwD3sW6oU331KrYLyN97LSJ5hhop5s5y/kdapvwVLhToUa8c4VRTHde8x\nN8XZyKWJRxZHXRx9FToXMkHEGhT7ei8QUvuKQ7xFsmtbMDm8FQv/O3gs5GCI3MKw\nnGQZOxLvP31ZZ0un+7CS0HuyGl4KowHCQvvqKxMn4kLGa3qCtJCjKRNVoasr78+m\nCNGN2IlaCJZsBHYNEej6/vP7TJiYoNQQupFqkzz9w90nPhxkGU6dnPH78pOcWBAY\nU4cMRsJqezGkXHph9Z6j8fZyqdo+zNpyRQIhR84V1xDiKBsK6F5pInT1PsjJRyQ7\n-----END RSA PRIVATE KEY-----\n",
				password:         "aaa",
			},
			mockBehavior: func(f fields, a args) {
				fileContent := `{"Content":"G+DVq/1yfH4+hOhSnVYwiS0FK7SFADA65C6kPKKOR2OG6qXe0F/C1OKSRjm3CKQyY2cchCs0WyZopDZwTFsLMmS+GPM842FbBm/5KSbiNXeS0PsmoQW2RmVVV7iCJDeiNlzJHkdqDcZU/VwHj0FW0Z9nfDpXeBQyKt1Wx1WUXRI=","Encrypted":true}`
				f.TextManagementRepository.(*mockrepository.TextManagementInterface).On("Load", a.textId).Return([]byte(fileContent), nil)
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.TextManagementRepository.(*mockrepository.TextManagementInterface).AssertExpectations(t)
			},
			want:    "aaaaaaaa",
			wantErr: false,
		},
		{
			name: "Get non encrypted content",
			fields: fields{
				&mockrepository.TextManagementInterface{},
				&mockhelper.HelperInterface{},
			},
			args: args{
				textId:           "2f13ed58-afc9-477a-bf0d-c90eb1b7db90",
				privateKeyString: "",
				password:         "",
			},
			mockBehavior: func(f fields, a args) {
				fileContent := `{"Content":"aaaaaaaa","Encrypted":false}`
				f.TextManagementRepository.(*mockrepository.TextManagementInterface).On("Load", a.textId).Return([]byte(fileContent), nil)
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.TextManagementRepository.(*mockrepository.TextManagementInterface).AssertExpectations(t)
			},
			want:    "aaaaaaaa",
			wantErr: false,
		},
		{
			name: "Error when reading file",
			fields: fields{
				&mockrepository.TextManagementInterface{},
				&mockhelper.HelperInterface{},
			},
			args: args{
				textId:           "2f13ed58-afc9-477a-bf0d-c90eb1b7db90",
				privateKeyString: "",
				password:         "",
			},
			mockBehavior: func(f fields, a args) {
				f.TextManagementRepository.(*mockrepository.TextManagementInterface).On("Load", a.textId).Return(nil, errors.New("error"))
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.TextManagementRepository.(*mockrepository.TextManagementInterface).AssertExpectations(t)
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Get encrypted content without private key",
			fields: fields{
				&mockrepository.TextManagementInterface{},
				&mockhelper.HelperInterface{},
			},
			args: args{
				textId:           "2f13ed58-afc9-477a-bf0d-c90eb1b7db90",
				privateKeyString: "",
				password:         "aaa",
			},
			mockBehavior: func(f fields, a args) {
				fileContent := `{"Content":"G+DVq/1yfH4+hOhSnVYwiS0FK7SFADA65C6kPKKOR2OG6qXe0F/C1OKSRjm3CKQyY2cchCs0WyZopDZwTFsLMmS+GPM842FbBm/5KSbiNXeS0PsmoQW2RmVVV7iCJDeiNlzJHkdqDcZU/VwHj0FW0Z9nfDpXeBQyKt1Wx1WUXRI=","Encrypted":true}`
				f.TextManagementRepository.(*mockrepository.TextManagementInterface).On("Load", a.textId).Return([]byte(fileContent), nil)
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.TextManagementRepository.(*mockrepository.TextManagementInterface).AssertExpectations(t)
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Get encrypted content without password",
			fields: fields{
				&mockrepository.TextManagementInterface{},
				&mockhelper.HelperInterface{},
			},
			args: args{
				textId:           "2f13ed58-afc9-477a-bf0d-c90eb1b7db90",
				privateKeyString: "-----BEGIN RSA PRIVATE KEY-----\nProc-Type: 4,ENCRYPTED\nDEK-Info: AES-256-CBC,53ae7dd8c61749faa0dcc2ace0dff6cc\n\nH6QV4F6XfP2EJ4mebb0t2YZKDwK6HBJXO2bdgVKlkuOepf+LBLJWsF88Dbk1PYn8\nXGLTKBGEhtK3Qy3+IeCeWUQOxdEKc2RJKMLl5Hs1uZm6okvJSdECNBHlO4PC2WDC\nuLT+CsFWnVNQH/WT16EpUQOu0IU7GugnOqIPm1z0U4XtprX8Xcw1uNl9uyDEEUxg\nvwGEc2Qi4RpcAbK1bnPUQ6BjQt0DYiOyEJZFeTTkXYCTsnifTIcyp19ItLYzXjDz\nbsRJX52jCj66dEHmAiXq8u5vpU55Y7Gt73CjBUf3mcNMMzzdhurJeKXCXkcWFXb7\nj2sbJDdEjyLGAlE4Fxeyn7hpmj8yo01oNFcpJYgoEgp+AdYgvjhEykYFqm4NzP/Y\nxmyDKU3rqO1nl9jOckzGLq8Ca16psLgXbcrpkfxv9dik5HiAUaaBZ3bINkrisoRg\nd7by7ig03sDlQWMU9P2b0UZx5jD/MUB06dDIAK7MtQ2mGTs99PB/42IO6Xv1/PYn\nObMNgmuZwD3sW6oU331KrYLyN97LSJ5hhop5s5y/kdapvwVLhToUa8c4VRTHde8x\nN8XZyKWJRxZHXRx9FToXMkHEGhT7ei8QUvuKQ7xFsmtbMDm8FQv/O3gs5GCI3MKw\nnGQZOxLvP31ZZ0un+7CS0HuyGl4KowHCQvvqKxMn4kLGa3qCtJCjKRNVoasr78+m\nCNGN2IlaCJZsBHYNEej6/vP7TJiYoNQQupFqkzz9w90nPhxkGU6dnPH78pOcWBAY\nU4cMRsJqezGkXHph9Z6j8fZyqdo+zNpyRQIhR84V1xDiKBsK6F5pInT1PsjJRyQ7\n-----END RSA PRIVATE KEY-----\n",
				password:         "",
			},
			mockBehavior: func(f fields, a args) {
				fileContent := `{"Content":"G+DVq/1yfH4+hOhSnVYwiS0FK7SFADA65C6kPKKOR2OG6qXe0F/C1OKSRjm3CKQyY2cchCs0WyZopDZwTFsLMmS+GPM842FbBm/5KSbiNXeS0PsmoQW2RmVVV7iCJDeiNlzJHkdqDcZU/VwHj0FW0Z9nfDpXeBQyKt1Wx1WUXRI=","Encrypted":true}`
				f.TextManagementRepository.(*mockrepository.TextManagementInterface).On("Load", a.textId).Return([]byte(fileContent), nil)
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.TextManagementRepository.(*mockrepository.TextManagementInterface).AssertExpectations(t)
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Get encrypted content with wrong private key",
			fields: fields{
				&mockrepository.TextManagementInterface{},
				&mockhelper.HelperInterface{},
			},
			args: args{
				textId:           "2f13ed58-afc9-477a-bf0d-c90eb1b7db90",
				privateKeyString: "-----BEGIN RSA PRIVATE KEY-----\nProc-Type: 4,ENCRYPTED\nDEK-Info: AES-256-CBC,53ae7dd8c61749faa0dcc2ace0dff6cc\nXGLTKBGEhtK3Qy3+IeCeWUQOxdEKc2RJKMLl5Hs1uZm6okvJSdECNBHlO4PC2WDC\nuLT+CsFWnVNQH/WT16EpUQOu0IU7GugnOqIPm1z0U4XtprX8Xcw1uNl9uyDEEUxg\nvwGEc2Qi4RpcAbK1bnPUQ6BjQt0DYiOyEJZFeTTkXYCTsnifTIcyp19ItLYzXjDz\nbsRJX52jCj66dEHmAiXq8u5vpU55Y7Gt73CjBUf3mcNMMzzdhurJeKXCXkcWFXb7\nj2sbJDdEjyLGAlE4Fxeyn7hpmj8yo01oNFcpJYgoEgp+AdYgvjhEykYFqm4NzP/Y\nxmyDKU3rqO1nl9jOckzGLq8Ca16psLgXbcrpkfxv9dik5HiAUaaBZ3bINkrisoRg\nd7by7ig03sDlQWMU9P2b0UZx5jD/MUB06dDIAK7MtQ2mGTs99PB/42IO6Xv1/PYn\nObMNgmuZwD3sW6oU331KrYLyN97LSJ5hhop5s5y/kdapvwVLhToUa8c4VRTHde8x\nN8XZyKWJRxZHXRx9FToXMkHEGhT7ei8QUvuKQ7xFsmtbMDm8FQv/O3gs5GCI3MKw\nnGQZOxLvP31ZZ0un+7CS0HuyGl4KowHCQvvqKxMn4kLGa3qCtJCjKRNVoasr78+m\nCNGN2IlaCJZsBHYNEej6/vP7TJiYoNQQupFqkzz9w90nPhxkGU6dnPH78pOcWBAY\nU4cMRsJqezGkXHph9Z6j8fZyqdo+zNpyRQIhR84V1xDiKBsK6F5pInT1PsjJRyQ7\n-----END RSA PRIVATE KEY-----\n",
				password:         "aaa",
			},
			mockBehavior: func(f fields, a args) {
				fileContent := `{"Content":"G+DVq/1yfH4+hOhSnVYwiS0FK7SFADA65C6kPKKOR2OG6qXe0F/C1OKSRjm3CKQyY2cchCs0WyZopDZwTFsLMmS+GPM842FbBm/5KSbiNXeS0PsmoQW2RmVVV7iCJDeiNlzJHkdqDcZU/VwHj0FW0Z9nfDpXeBQyKt1Wx1WUXRI=","Encrypted":true}`
				f.TextManagementRepository.(*mockrepository.TextManagementInterface).On("Load", a.textId).Return([]byte(fileContent), nil)
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.TextManagementRepository.(*mockrepository.TextManagementInterface).AssertExpectations(t)
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Get encrypted content with wrong password",
			fields: fields{
				&mockrepository.TextManagementInterface{},
				&mockhelper.HelperInterface{},
			},
			args: args{
				textId:           "2f13ed58-afc9-477a-bf0d-c90eb1b7db90",
				privateKeyString: "-----BEGIN RSA PRIVATE KEY-----\nProc-Type: 4,ENCRYPTED\nDEK-Info: AES-256-CBC,53ae7dd8c61749faa0dcc2ace0dff6cc\n\nH6QV4F6XfP2EJ4mebb0t2YZKDwK6HBJXO2bdgVKlkuOepf+LBLJWsF88Dbk1PYn8\nXGLTKBGEhtK3Qy3+IeCeWUQOxdEKc2RJKMLl5Hs1uZm6okvJSdECNBHlO4PC2WDC\nuLT+CsFWnVNQH/WT16EpUQOu0IU7GugnOqIPm1z0U4XtprX8Xcw1uNl9uyDEEUxg\nvwGEc2Qi4RpcAbK1bnPUQ6BjQt0DYiOyEJZFeTTkXYCTsnifTIcyp19ItLYzXjDz\nbsRJX52jCj66dEHmAiXq8u5vpU55Y7Gt73CjBUf3mcNMMzzdhurJeKXCXkcWFXb7\nj2sbJDdEjyLGAlE4Fxeyn7hpmj8yo01oNFcpJYgoEgp+AdYgvjhEykYFqm4NzP/Y\nxmyDKU3rqO1nl9jOckzGLq8Ca16psLgXbcrpkfxv9dik5HiAUaaBZ3bINkrisoRg\nd7by7ig03sDlQWMU9P2b0UZx5jD/MUB06dDIAK7MtQ2mGTs99PB/42IO6Xv1/PYn\nObMNgmuZwD3sW6oU331KrYLyN97LSJ5hhop5s5y/kdapvwVLhToUa8c4VRTHde8x\nN8XZyKWJRxZHXRx9FToXMkHEGhT7ei8QUvuKQ7xFsmtbMDm8FQv/O3gs5GCI3MKw\nnGQZOxLvP31ZZ0un+7CS0HuyGl4KowHCQvvqKxMn4kLGa3qCtJCjKRNVoasr78+m\nCNGN2IlaCJZsBHYNEej6/vP7TJiYoNQQupFqkzz9w90nPhxkGU6dnPH78pOcWBAY\nU4cMRsJqezGkXHph9Z6j8fZyqdo+zNpyRQIhR84V1xDiKBsK6F5pInT1PsjJRyQ7\n-----END RSA PRIVATE KEY-----\n",
				password:         "bbb",
			},
			mockBehavior: func(f fields, a args) {
				fileContent := `{"Content":"G+DVq/1yfH4+hOhSnVYwiS0FK7SFADA65C6kPKKOR2OG6qXe0F/C1OKSRjm3CKQyY2cchCs0WyZopDZwTFsLMmS+GPM842FbBm/5KSbiNXeS0PsmoQW2RmVVV7iCJDeiNlzJHkdqDcZU/VwHj0FW0Z9nfDpXeBQyKt1Wx1WUXRI=","Encrypted":true}`
				f.TextManagementRepository.(*mockrepository.TextManagementInterface).On("Load", a.textId).Return([]byte(fileContent), nil)
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.TextManagementRepository.(*mockrepository.TextManagementInterface).AssertExpectations(t)
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Get encrypted empty content",
			fields: fields{
				&mockrepository.TextManagementInterface{},
				&mockhelper.HelperInterface{},
			},
			args: args{
				textId:           "2f13ed58-afc9-477a-bf0d-c90eb1b7db90",
				privateKeyString: "-----BEGIN RSA PRIVATE KEY-----\nProc-Type: 4,ENCRYPTED\nDEK-Info: AES-256-CBC,53ae7dd8c61749faa0dcc2ace0dff6cc\n\nH6QV4F6XfP2EJ4mebb0t2YZKDwK6HBJXO2bdgVKlkuOepf+LBLJWsF88Dbk1PYn8\nXGLTKBGEhtK3Qy3+IeCeWUQOxdEKc2RJKMLl5Hs1uZm6okvJSdECNBHlO4PC2WDC\nuLT+CsFWnVNQH/WT16EpUQOu0IU7GugnOqIPm1z0U4XtprX8Xcw1uNl9uyDEEUxg\nvwGEc2Qi4RpcAbK1bnPUQ6BjQt0DYiOyEJZFeTTkXYCTsnifTIcyp19ItLYzXjDz\nbsRJX52jCj66dEHmAiXq8u5vpU55Y7Gt73CjBUf3mcNMMzzdhurJeKXCXkcWFXb7\nj2sbJDdEjyLGAlE4Fxeyn7hpmj8yo01oNFcpJYgoEgp+AdYgvjhEykYFqm4NzP/Y\nxmyDKU3rqO1nl9jOckzGLq8Ca16psLgXbcrpkfxv9dik5HiAUaaBZ3bINkrisoRg\nd7by7ig03sDlQWMU9P2b0UZx5jD/MUB06dDIAK7MtQ2mGTs99PB/42IO6Xv1/PYn\nObMNgmuZwD3sW6oU331KrYLyN97LSJ5hhop5s5y/kdapvwVLhToUa8c4VRTHde8x\nN8XZyKWJRxZHXRx9FToXMkHEGhT7ei8QUvuKQ7xFsmtbMDm8FQv/O3gs5GCI3MKw\nnGQZOxLvP31ZZ0un+7CS0HuyGl4KowHCQvvqKxMn4kLGa3qCtJCjKRNVoasr78+m\nCNGN2IlaCJZsBHYNEej6/vP7TJiYoNQQupFqkzz9w90nPhxkGU6dnPH78pOcWBAY\nU4cMRsJqezGkXHph9Z6j8fZyqdo+zNpyRQIhR84V1xDiKBsK6F5pInT1PsjJRyQ7\n-----END RSA PRIVATE KEY-----\n",
				password:         "aaa",
			},
			mockBehavior: func(f fields, a args) {
				fileContent := `{"Content":"","Encrypted":true}`
				f.TextManagementRepository.(*mockrepository.TextManagementInterface).On("Load", a.textId).Return([]byte(fileContent), nil)
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.TextManagementRepository.(*mockrepository.TextManagementInterface).AssertExpectations(t)
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

			service := service.NewService(tt.fields.TextManagementRepository, tt.fields.Helper)

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
	encryptedMessage := []byte("!@#$%¨&")
	encryptedRequestEncryption := true
	encryptedRequest := entity.TextManagement{
		TextData:           "aaaaaaaa",
		Encryption:         &encryptedRequestEncryption,
		KeySize:            1024,
		PrivateKeyPassword: "aaa",
	}
	encryptedResponse := entity.TextManagement{
		TextData:           base64.StdEncoding.EncodeToString(encryptedMessage),
		Encryption:         encryptedRequest.Encryption,
		KeySize:            encryptedRequest.KeySize,
		PrivateKeyPassword: encryptedRequest.PrivateKeyPassword,
		Uuid:               "47b416d1-c5f2-417e-929e-7b83667c6654",
		PrivateKey:         "-----BEGIN RSA PRIVATE KEY-----\nProc-Type: 4,ENCRYPTED\nDEK-Info: AES-256-CBC,acb37115f9b09418b21bb6fcf33be50c\n\nusbQ2x4BnMlx7iws1d0jGTGrFqKXOixk84oR7HYKjFsQXCz0ro6+wC6k4pQRxpSA\n+KnFcfrpSfgx6GpkEotuyDNBBu7rqkSZykWyxd9xC6FVcRpdpHFnzX44qehHYiOk\njylsWoaG/yS/zaTGKJGB1AQRSHfB3+G6tgWtZCW70S6u8Y3QyfrokgQ7wHrCVMaJ\nnazR/klacUPDd48zY8AU+mh/duZctaYW56lsUuNFuwdLXsI7ViR63iFDYV2gvJn0\no73Ju5NHEwWBPmsbYaBusvP8KdpliJpwMiONr8hP9D6PLqADtXe2Xne4t2EnSbQN\n1sHswCjs7aHHxDdL2ZuL8Yo+uh3GjeXL8a9XM5SHyecGnVB+UUOK7mkk/XxX8eEn\nITZedRDjrLzImFfKfa65fFoCAotB/k+LQI95TUVl1Q/mRK/fG+5F91JYBcuRNWdX\nBM/J5zZznQo68sSnopWdiFxM6Hi85FseKBLbz2qIqd01xC5zYulAnldGJRSUeGbR\n6jBt/Bf5xdmUaHyRpmDt24b6Qv9bOz6LT3+My+e61nkwqs6yyFOaCbON27odK6gb\niOH583AG1bLVfnLuknTXmgeSrSsugqhficVXRCOGfsAZjjCtaM+nu6TkheUA2zA9\nNuy0tBLoD5W0MF5bK14G4oLTkVvUtNGd0qWJjhN81zfym31EMpTCliWhHD1BXTc9\nFzInfSd8b/bWqPS3H8t2iWxPgsWKoJYqx5EB5cOr3ysQMd+XBh8+j6T1UWYPmHux\neNibYpZLa28Lfp2vrmolBivZctAC2K97Juchc4GtY1InYGr46lsIhqVN5E2GtNhd\n-----END RSA PRIVATE KEY-----\n",
	}
	unencryptedRequestEncryption := false
	unencryptedRequest := entity.TextManagement{
		TextData:   "aaaaaaaa",
		Encryption: &unencryptedRequestEncryption,
	}
	unencryptedResponse := entity.TextManagement{
		TextData:   unencryptedRequest.TextData,
		Encryption: unencryptedRequest.Encryption,
		Uuid:       "47b416d1-c5f2-417e-929e-7b83667c6654",
	}

	type fields struct {
		TextManagementRepository repository.TextManagementInterface
		Helper                   helper.HelperInterface
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
		{
			name: "Insert encrypted content",
			fields: fields{
				&mockrepository.TextManagementInterface{},
				&mockhelper.HelperInterface{},
			},
			args: args{
				text: entity.TextManagement{
					TextData:           encryptedRequest.TextData,
					Encryption:         encryptedRequest.Encryption,
					KeySize:            encryptedRequest.KeySize,
					PrivateKeyPassword: encryptedRequest.PrivateKeyPassword,
				},
			},
			mockBehavior: func(f fields, a args) {
				fileData := struct {
					Content   string
					Encrypted bool
				}{
					Content:   encryptedResponse.TextData,
					Encrypted: *encryptedRequest.Encryption,
				}
				b, _ := json.Marshal(fileData)

				f.Helper.(*mockhelper.HelperInterface).On("GenerateUuid").Return(encryptedResponse.Uuid)
				f.Helper.(*mockhelper.HelperInterface).On("EncryptMessage", a.text.KeySize, a.text.TextData, a.text.PrivateKeyPassword).Return(encryptedMessage, encryptedResponse.PrivateKey, nil)
				f.TextManagementRepository.(*mockrepository.TextManagementInterface).On("Save", encryptedResponse.Uuid, string(b)).Return(nil)
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.TextManagementRepository.(*mockrepository.TextManagementInterface).AssertExpectations(t)
			},
			want:    encryptedResponse,
			wantErr: false,
		},
		{
			name: "Insert unencrypted content",
			fields: fields{
				&mockrepository.TextManagementInterface{},
				&mockhelper.HelperInterface{},
			},
			args: args{
				text: entity.TextManagement{
					TextData:   unencryptedRequest.TextData,
					Encryption: unencryptedRequest.Encryption,
				},
			},
			mockBehavior: func(f fields, a args) {
				fileData := struct {
					Content   string
					Encrypted bool
				}{
					Content:   unencryptedResponse.TextData,
					Encrypted: *unencryptedRequest.Encryption,
				}
				b, _ := json.Marshal(fileData)

				f.Helper.(*mockhelper.HelperInterface).On("GenerateUuid").Return(unencryptedResponse.Uuid)
				f.TextManagementRepository.(*mockrepository.TextManagementInterface).On("Save", unencryptedResponse.Uuid, string(b)).Return(nil)
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.TextManagementRepository.(*mockrepository.TextManagementInterface).AssertExpectations(t)
			},
			want:    unencryptedResponse,
			wantErr: false,
		},
		{
			name: "Insert encrypted content encryptation error",
			fields: fields{
				&mockrepository.TextManagementInterface{},
				&mockhelper.HelperInterface{},
			},
			args: args{
				text: entity.TextManagement{
					TextData:           encryptedRequest.TextData,
					Encryption:         encryptedRequest.Encryption,
					KeySize:            encryptedRequest.KeySize,
					PrivateKeyPassword: encryptedRequest.PrivateKeyPassword,
				},
			},
			mockBehavior: func(f fields, a args) {
				f.Helper.(*mockhelper.HelperInterface).On("GenerateUuid").Return(encryptedResponse.Uuid)
				f.Helper.(*mockhelper.HelperInterface).On("EncryptMessage", a.text.KeySize, a.text.TextData, a.text.PrivateKeyPassword).Return(nil, "", errors.New("error"))
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.TextManagementRepository.(*mockrepository.TextManagementInterface).AssertExpectations(t)
			},
			want:    entity.TextManagement{},
			wantErr: true,
		},
		{
			name: "Insert content save file error",
			fields: fields{
				&mockrepository.TextManagementInterface{},
				&mockhelper.HelperInterface{},
			},
			args: args{
				text: entity.TextManagement{
					TextData:   unencryptedRequest.TextData,
					Encryption: unencryptedRequest.Encryption,
				},
			},
			mockBehavior: func(f fields, a args) {
				fileData := struct {
					Content   string
					Encrypted bool
				}{
					Content:   unencryptedResponse.TextData,
					Encrypted: *unencryptedRequest.Encryption,
				}
				b, _ := json.Marshal(fileData)

				f.Helper.(*mockhelper.HelperInterface).On("GenerateUuid").Return(unencryptedResponse.Uuid)
				f.TextManagementRepository.(*mockrepository.TextManagementInterface).On("Save", unencryptedResponse.Uuid, string(b)).Return(errors.New("error"))
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.TextManagementRepository.(*mockrepository.TextManagementInterface).AssertExpectations(t)
			},
			want:    entity.TextManagement{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockBehavior != nil {
				tt.mockBehavior(tt.fields, tt.args)
			}

			service := service.NewService(tt.fields.TextManagementRepository, tt.fields.Helper)

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
