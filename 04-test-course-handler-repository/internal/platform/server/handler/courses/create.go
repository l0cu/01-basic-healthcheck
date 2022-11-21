package courses

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	mooc "/internal/platform/server"

	"github.com/gin-gonic/gin"
)

type createRequest struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

// CreateHandler returns an HTTP handler for courses creation
func CreateHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		course := mooc.NewCourse(req.ID, req.Name, req.Duration)

		if err := Save(ctx, course); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.Status(http.StatusCreated)
	}
}

const (
	dbUser = "codely"
	dbPass = "codely"
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "codely"
)

// Save persist the course on the database
func Save(ctx context.Context, course mooc.Course) error {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	type sqlCourse struct {
		ID       string `db:"id"`
		Name     string `db:"Name"`
		Duration string `db:"Duration"`
	}

	courseSQLStruct := sqlbuilder.NewStruct(new(sqlCourse))
	query, args := courseSQLStruct.InsertInto("courses", sqlCourse{
		ID:       course.ID(),
		Name:     course.Name(),
		Duration: course.Duration(),
	}).Build()

	_, err = db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist course in database: %v", err.Error())
	}
	return nil
}
