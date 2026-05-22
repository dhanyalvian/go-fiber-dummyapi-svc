//- apps/models/base_model.go

package models

import (
	"math"
	"reflect"
	"strconv"
	"strings"

	"github.com/dhanyalvian/go-fiber-packages/request"
	"github.com/dhanyalvian/go-fiber-packages/response"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetModelListData[T any](
	c *fiber.Ctx,
	dbCount *gorm.DB,
	dbRow *gorm.DB,
) response.ResponseData {
	var results []T
	var totalRecords int64

	// 1. Hitung Total Records (Tanpa Limit/Offset/Select)
	dbCount.Count(&totalRecords)

	// 2. Ambil Data (Gunakan Scan untuk mapping ke DTO)
	dbRow.Scopes(Paginate(c)).Scan(&results)

	page := request.GetPage(c)
	limit := request.GetLimit(c)

	// 3. Kalkulasi Pagination
	totalPages := int(math.Ceil(float64(totalRecords) / float64(limit)))
	next := page + 1
	if next > totalPages {
		next = 0
	}

	// Jika hasil kosong, pastikan mengembalikan slice kosong [] bukan nil
	if len(results) == 0 {
		results = []T{}
	}

	return response.ResponseData{
		Pagination: &response.ResponseDataPagination{
			Page:         page,
			Next:         next,
			Records:      len(results),
			TotalPages:   totalPages,
			TotalRecords: totalRecords,
		},
		Results: results,
	}
}

func GetListData[T any]( // T adalah tipe DTO (misal: InvoiceListDTO)
	c *fiber.Ctx,
	db *gorm.DB,
	entity interface{}, // Model asal (misal: Invoice{})
	where interface{},
	orders string,
) response.ResponseData {
	var results []T // Gunakan slice dari tipe T, bukan interface{}
	var totalRecords int64

	page := request.GetPage(c)
	limit := request.GetLimit(c)

	// Ambil kolom berdasarkan tipe T
	cols := GetColumns(new(T))

	// 1. Hitung Total Records (Tanpa Limit/Offset/Select)
	db.Model(entity).Where(where).Count(&totalRecords)

	// 2. Ambil Data (Gunakan Scan untuk mapping ke DTO)
	db.Model(entity).
		Select(cols).
		Where(where).
		Scopes(Paginate(c)). // Pastikan Paginate menggunakan limit dari request
		Order(orders).
		Scan(&results) // Gunakan Scan untuk DTO

	// 3. Kalkulasi Pagination
	totalPages := int(math.Ceil(float64(totalRecords) / float64(limit)))
	next := page + 1
	if next > totalPages {
		next = 0
	}

	// Jika hasil kosong, pastikan mengembalikan slice kosong [] bukan nil
	if len(results) == 0 {
		results = []T{}
	}

	return response.ResponseData{
		Pagination: &response.ResponseDataPagination{
			Page:         page,
			Next:         next,
			Records:      len(results),
			TotalPages:   totalPages,
			TotalRecords: totalRecords,
		},
		Results: results,
	}
}

func GetDetailById[T any](
	db *gorm.DB,
	entity interface{},
	id string,
	result T,
) response.ResponseData {
	idNumber, _ := strconv.Atoi(id)
	db.Model(entity).First(&result, uint(idNumber))

	return response.ResponseData{
		Result: result,
	}
}

func GetDetailData(
	db *gorm.DB,
	where interface{},
	result interface{},
) response.ResponseData {
	db.Where(where).First(result)

	return response.ResponseData{
		Result: result,
	}
}

func GetColumns(s interface{}) []string {
	var columns []string

	// 1. Dapatkan Value dari interface
	v := reflect.ValueOf(s)

	// 2. Jika Pointer, ambil elemen aslinya (Edisi rekursif untuk pointer ke pointer)
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// 3. VALIDASI: Pastikan sekarang kita memegang STRUCT
	// Jika bukan struct (misal nil atau tipe dasar), kembalikan slice kosong agar tidak panic
	if v.Kind() != reflect.Struct {
		return columns
	}

	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i) // Ini titik yang sebelumnya menyebabkan panic

		// 4. JIKA FIELD ADALAH EMBEDDED STRUCT (BaseID)
		if field.Anonymous {
			// Gunakan Interface() dari fieldValue untuk rekursi
			embeddedCols := GetColumns(fieldValue.Interface())
			columns = append(columns, embeddedCols...)
			continue
		}

		// 5. AMBIL DARI TAG GORM ATAU CONVERT NAME
		columnName := ""
		gormTag := field.Tag.Get("gorm")
		if gormTag != "" {
			parts := strings.Split(gormTag, ";")
			for _, part := range parts {
				if strings.HasPrefix(part, "column:") {
					columnName = strings.TrimPrefix(part, "column:")
					break
				}
			}
		}

		if columnName != "" {
			columns = append(columns, columnName)
		}
	}

	return columns
}

func Paginate(c *fiber.Ctx) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := request.GetPage(c)
		limit := request.GetLimit(c)
		offset := (page - 1) * limit

		return db.Offset(offset).Limit(limit)
	}
}
