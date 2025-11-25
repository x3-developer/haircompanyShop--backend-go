package mapper

import (
	"serv_shop_haircompany/internal/shared/application/dto"
)

func ToPaginationResponseDTO[T any](items []T, total int64) *dto.PaginatedDTO[T] {
	return &dto.PaginatedDTO[T]{
		Items: items,
		Total: total,
	}
}
