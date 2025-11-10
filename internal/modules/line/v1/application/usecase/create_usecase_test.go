package usecase

import (
	"context"
	"errors"
	"reflect"
	"serv_shop_haircompany/internal/modules/line/v1/domain"
	"serv_shop_haircompany/internal/shared/utils/response"
	"testing"
)

type stubRepo struct {
	exists    bool
	existsErr error
	createOut *domain.Line
	createErr error
}

func (s *stubRepo) ExistsByUniqueFields(_ context.Context, _ string) (bool, error) {
	return s.exists, s.existsErr
}
func (s *stubRepo) Create(_ context.Context, _ *domain.Line) (*domain.Line, error) {
	return s.createOut, s.createErr
}

func TestCreateUseCase_NameNotUnique(t *testing.T) {
	repo := &stubRepo{exists: true}
	uc := NewCreateUseCase(repo)

	gotModel, gotFields, gotErr := uc.Execute(context.Background(), &domain.Line{Name: "Basic", Color: "#112233"})
	if gotErr != nil {
		t.Fatalf("want no error, got %v", gotErr)
	}
	if gotModel != nil {
		t.Fatalf("want nil model, got %#v", gotModel)
	}
	want := []response.ErrorField{response.NewErrorField("name", string(response.NotUnique))}
	if !reflect.DeepEqual(gotFields, want) {
		t.Fatalf("want fields=%v, got %v", want, gotFields)
	}
}

func TestCreateUseCase_ExistsCheckError(t *testing.T) {
	repo := &stubRepo{existsErr: errors.New("db down")}
	uc := NewCreateUseCase(repo)

	_, _, err := uc.Execute(context.Background(), &domain.Line{Name: "X", Color: "#000000"})
	if err == nil {
		t.Fatalf("want error, got nil")
	}
}

func TestCreateUseCase_CreateOK(t *testing.T) {
	out := &domain.Line{ID: 10, Name: "New", Color: "#ABCDEF"}
	repo := &stubRepo{exists: false, createOut: out}
	uc := NewCreateUseCase(repo)

	got, fields, err := uc.Execute(context.Background(), &domain.Line{Name: "New", Color: "#ABCDEF"})
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if fields != nil {
		t.Fatalf("want no validation fields, got %v", fields)
	}
	if !reflect.DeepEqual(got, out) {
		t.Fatalf("want %#v, got %#v", out, got)
	}
}

func TestCreateUseCase_CreateError(t *testing.T) {
	repo := &stubRepo{exists: false, createErr: errors.New("insert fail")}
	uc := NewCreateUseCase(repo)

	_, _, err := uc.Execute(context.Background(), &domain.Line{Name: "New", Color: "#FFFFFF"})
	if err == nil {
		t.Fatalf("want error, got nil")
	}
}
