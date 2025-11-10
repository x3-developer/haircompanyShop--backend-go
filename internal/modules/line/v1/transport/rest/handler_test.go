package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"serv_shop_haircompany/internal/modules/line/v1/domain"
	shareddomain "serv_shop_haircompany/internal/shared/domain"
	"strings"
	"testing"
	"time"

	"serv_shop_haircompany/internal/modules/line/v1/application/dto"
	"serv_shop_haircompany/internal/shared/utils/response"
)

type mockCreateUC struct {
	out       *domain.Line
	valFields []response.ErrorField
	err       error
	lastCtx   context.Context
	lastModel *domain.Line
}

func (m *mockCreateUC) Execute(ctx context.Context, model *domain.Line) (*domain.Line, []response.ErrorField, error) {
	m.lastCtx = ctx
	m.lastModel = model
	return m.out, m.valFields, m.err
}

func newRequest(body any) *http.Request {
	var buf bytes.Buffer
	if body != nil {
		_ = json.NewEncoder(&buf).Encode(body)
	}
	req := httptest.NewRequest(http.MethodPost, "/lines", &buf)
	req.Header.Set("Content-Type", "application/json")
	return req
}

func TestHandler_Create_Success(t *testing.T) {
	now := time.Now().UTC().Truncate(time.Second)
	uc := &mockCreateUC{
		out: &domain.Line{ID: 1, Name: "Ok", Color: "#112233", Timestamps: shareddomain.Timestamps{
			CreatedAt: now, UpdatedAt: now,
		}},
	}
	h := NewHandler(uc)

	rec := httptest.NewRecorder()
	req := newRequest(dto.CreateDTO{Name: "Okay", Color: "#112233"})
	h.Create(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("want 201, got %d body=%s", rec.Code, rec.Body.String())
	}
	if !bytes.Contains(rec.Body.Bytes(), []byte(`"name":"Ok"`)) {
		t.Fatalf("response missing name, body=%s", rec.Body.String())
	}
}

func TestHandler_Create_InvalidJSON(t *testing.T) {
	uc := &mockCreateUC{}
	h := NewHandler(uc)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/lines", bytes.NewBufferString("{bad json"))
	req.Header.Set("Content-Type", "application/json")
	h.Create(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("want 400, got %d", rec.Code)
	}
	if !bytes.Contains(rec.Body.Bytes(), []byte(`"errorCode":"BAD_REQUEST"`)) {
		t.Fatalf("want BAD_REQUEST code, got body=%s", rec.Body.String())
	}
}

func TestHandler_Create_DTOValidationError(t *testing.T) {
	uc := &mockCreateUC{}
	h := NewHandler(uc)

	rec := httptest.NewRecorder()
	req := newRequest(dto.CreateDTO{Name: "x", Color: ""})
	h.Create(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("want 400, got %d", rec.Code)
	}
	if !bytes.Contains(rec.Body.Bytes(), []byte(`"fields"`)) {
		t.Fatalf("expected validation fields, got %+s", rec.Body.String())
	}
}

func TestHandler_Create_BusinessValidation_NotUnique(t *testing.T) {
	uc := &mockCreateUC{
		valFields: []response.ErrorField{response.NewErrorField("name", string(response.NotUnique))},
	}
	h := NewHandler(uc)

	rec := httptest.NewRecorder()
	req := newRequest(dto.CreateDTO{Name: "Dup", Color: "#000000"})
	h.Create(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("want 400, got %d", rec.Code)
	}
	if !bytes.Contains(rec.Body.Bytes(), []byte(`"name"`)) {
		t.Fatalf("validation field 'name' expected, body=%s", rec.Body.String())
	}
}

func TestHandler_Create_ServerError(t *testing.T) {
	uc := &mockCreateUC{err: context.DeadlineExceeded}
	h := NewHandler(uc)

	rec := httptest.NewRecorder()
	req := newRequest(dto.CreateDTO{Name: "Any", Color: "#000001"})
	h.Create(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("want 500, got %d", rec.Code)
	}
	if !bytes.Contains(rec.Body.Bytes(), []byte(`"errorCode":"SERVER_ERROR"`)) {
		t.Fatalf("want SERVER_ERROR code, body=%s", rec.Body.String())
	}
}

func TestHandler_Create_InvalidColor_FromMapper(t *testing.T) {
	uc := &mockCreateUC{}
	h := NewHandler(uc)

	rec := httptest.NewRecorder()
	req := newRequest(dto.CreateDTO{Name: "ValidName", Color: "123456"})

	h.Create(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("want 400, got %d", rec.Code)
	}

	body := rec.Body.String()
	if !strings.Contains(body, `"color"`) {
		t.Fatalf("expected validation error for color, got body=%s", body)
	}
}
