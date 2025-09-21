package login

import (
	reflect "reflect"
	"testing"
	"time"

	"github.com/nick6969/go-clean-project/internal/domain"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestUseCase_Execute(t *testing.T) {
	// 準備 Mock Controller
	ctrl := gomock.NewController(t)

	// 測試結束後釋放資源
	defer ctrl.Finish()

	// 建立 Mock 物件
	mockRepository := NewMockrepository(ctrl)
	mockPassword := NewMockpassword(ctrl)
	mockToken := NewMocktoken(ctrl)

	type fields struct {
		repository repository
		password   password
		token      token
	}

	defaultFields := fields{
		repository: mockRepository,
		password:   mockPassword,
		token:      mockToken,
	}

	getDomainUserModel := func(id int, email, password string) *domain.DBUserModel {
		user, err := domain.NewDBUserModel(id, email, password, time.Now(), time.Now())
		if err != nil {
			t.Fatalf("failed to create DBUserModel: %v", err)
		}
		return user
	}

	type args struct {
		input Input
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Output
		want1  *domain.GPError
		setup  func()
	}{
		{
			name:   "normal case",
			fields: defaultFields,
			args: args{
				input: NewInput("<Email>", "<Password>"),
			},
			want: &Output{
				AccessToken: "<Token>",
			},
			want1: nil,
			setup: func() {
				// 設定 Mock 物件的行為
				mockRepository.EXPECT().FindUserByEmail(gomock.Any(), "<Email>").Return(getDomainUserModel(1, "<Email>", "<HashedPassword>"), nil)
				mockPassword.EXPECT().Compare("<HashedPassword>", "<Password>").Return(nil)
				mockToken.EXPECT().GenerateAccessToken(1).Return("<Token>", nil)
			},
		},
	}

	for _, tt := range tests {
		tt.setup()
		t.Run(tt.name, func(t *testing.T) {
			u := &UseCase{
				repository: tt.fields.repository,
				password:   tt.fields.password,
				token:      tt.fields.token,
			}
			got, got1 := u.Execute(t.Context(), tt.args.input)

			// 使用 assert 來檢查結果
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)

			// 傳統的比較方式
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UseCase.Execute() got = %v, want %v", got, tt.want)
			}
			if (got1 == nil) != (tt.want1 == nil) || (got1 != nil && tt.want1 != nil && got1.Error() != tt.want1.Error()) {
				t.Errorf("UseCase.Execute() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
