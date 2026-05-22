//- apps/models/user_model.go

package models

import (
	"go-fiber-dummy-svc/apps/entities"

	"github.com/dhanyalvian/go-fiber-packages/response"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ListUser(c *fiber.Ctx, db *gorm.DB, q string) response.ResponseData {
	entity := entities.User{}
	cols := GetColumns(new(entities.RespListUser))
	orders := "id ASC"

	dbCount := db.Model(&entity)
	dbRow := db.Model(&entity).Select(cols).Order(orders)

	if q != "" {
		whereSearch := "firstname ILIKE ? OR lastname ILIKE ? OR email ILIKE ?"
		whereBind := "%" + q + "%"

		dbCount.Where(whereSearch, whereBind, whereBind, whereBind)
		dbRow.Where(whereSearch, whereBind, whereBind, whereBind)
	}

	return GetModelListData[entities.RespListUser](c, dbCount, dbRow)
}

func DetailUser(db *gorm.DB, id string) response.ResponseData {
	var result entities.RespDetailUser
	entity := entities.User{}
	return GetDetailById(db, entity, id, &result)
}
