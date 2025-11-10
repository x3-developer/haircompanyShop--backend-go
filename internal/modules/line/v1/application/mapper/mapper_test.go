package mapper

import (
	"serv_shop_haircompany/internal/modules/line/v1/application/dto"
	"serv_shop_haircompany/internal/modules/line/v1/domain"
	"serv_shop_haircompany/internal/modules/line/v1/domain/valueobject"
	shareddomain "serv_shop_haircompany/internal/shared/domain"
	"testing"
	"time"
)

func TestToModelFromCreateDTO_ValidColor(t *testing.T) {
	in := dto.CreateDTO{Name: "Name", Color: "#00A1B2"}
	model, fields := ToModelFromCreateDTO(in)
	if fields != nil {
		t.Fatalf("unexpected fields: %v", fields)
	}
	if model == nil || model.Color != "#00A1B2" || model.Name != "Name" {
		t.Fatalf("unexpected model: %#v", model)
	}
}

func TestToModelFromCreateDTO_InvalidColor(t *testing.T) {
	cases := []string{
		"00A1B2",   // нет '#'
		"#XYZ123",  // не hex
		"#1234",    // неверная длина
		"#1234567", // неверная длина
		"",         // пусто
	}
	for _, c := range cases {
		in := dto.CreateDTO{Name: "Name", Color: c}
		model, fields := ToModelFromCreateDTO(in)
		if model != nil {
			t.Fatalf("want nil model for color=%q, got %#v", c, model)
		}
		if fields == nil || len(fields) == 0 {
			t.Fatalf("want validation fields for color=%q, got nil", c)
		}
		found := false
		for _, f := range fields {
			if f.Field == "color" {
				found = true
			}
		}
		if !found {
			t.Fatalf("want error on 'color' for %q, got %v", c, fields)
		}
	}
}

func TestToRespDTOFromModel(t *testing.T) {
	color, err := valueobject.NewColorVO("#AABBCC")
	if err != nil {
		t.Fatalf("unexpected error while constructing color vo: %v", err)
	}

	now := time.Now().UTC().Truncate(time.Second)

	line := domain.Line{
		ID:    10,
		Name:  "Test Line",
		Color: color,
		Timestamps: shareddomain.Timestamps{ // вот так
			CreatedAt: now,
			UpdatedAt: now,
		},
	}

	result := ToRespDTOFromModel(line)

	if result.Id != line.ID {
		t.Fatalf("expected Id %d, got %d", line.ID, result.Id)
	}
	if result.Name != line.Name {
		t.Fatalf("expected Name %q, got %q", line.Name, result.Name)
	}
	if result.Color != "#AABBCC" {
		t.Fatalf("expected Color %q, got %q", "#AABBCC", result.Color)
	}
	if !result.CreatedAt.Equal(now) {
		t.Fatalf("expected CreatedAt %v, got %v", now, result.CreatedAt)
	}
	if !result.UpdatedAt.Equal(now) {
		t.Fatalf("expected UpdatedAt %v, got %v", now, result.UpdatedAt)
	}

	if _, ok := interface{}(*result).(dto.ResponseDTO); !ok {
		t.Fatalf("expected type ResponseDTO, got %T", result)
	}
}
