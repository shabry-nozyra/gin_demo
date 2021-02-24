package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shabry-nozyra/gin_demo/config"
	"github.com/shabry-nozyra/gin_demo/models"
	"log"
	"net/http"
)

func (c *Context) add(ctx *gin.Context) {

	db, err := config.MySQL()
	if err != nil {
		errorm := "gagal terhubung ke Server MySql"
		ctx.JSON(http.StatusBadRequest, errorm)
		return
	}
	var paslons []models.Paslon

	queryText := fmt.Sprintf("SELECT * FROM paslon Order By id DESC")
	//ORM

	rowQuery, err := db.QueryContext(ctx, queryText)
	if err != nil {
		errorm := "Gagal Menjalankan Query"
		ctx.JSON(http.StatusBadRequest, errorm)
		return
	}

	for rowQuery.Next() {
		var paslon models.Paslon
		if err = rowQuery.Scan(&paslon.Id,
			&paslon.Bupati,
			&paslon.Wakil,
			&paslon.No_urut,
			&paslon.Foto); err != nil {
			return
		}
		paslons = append(paslons, paslon)
	}
	ctx.JSON(http.StatusOK, paslons)
}


func (c *Context) add3(ctx *gin.Context) {
	req := models.Paslon{}
	err := ctx.ShouldBindJSON(&req)

	db, errr := config.MySQL()

	if errr != nil {
		log.Fatal("Can't connect to MySQL", err)
		return
	}

	queryText := fmt.Sprintf("INSERT INTO paslon (bupati, wakil, no_urut, foto) values('%v','%v','%v','%v')",
		req.Bupati,
		req.Wakil,
		req.No_urut,
		req.Foto)

	_, errr = db.ExecContext(ctx, queryText)

	if errr != nil {
		errorm := "Gagal Menjalankan Query"
		ctx.JSON(http.StatusBadRequest, errorm)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}
	ctx.JSON(http.StatusCreated, res)
}

func (c *Context) edit(ctx *gin.Context) {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
		return
	}

	req := models.Paslon{}
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		errorm := "Gagal Menjalankan Query"
		ctx.JSON(http.StatusBadRequest, errorm)
		return
	}

	queryText := fmt.Sprintf("UPDATE paslon set bupati = '%s', wakil ='%s', no_urut = '%d', foto = '%s' where id = '%d'",
		req.Bupati,
		req.Wakil,
		req.No_urut,
		req.Foto,
		req.Id,
	)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		errorm := "Gagal Menjalankan Query"
		ctx.JSON(http.StatusBadRequest, errorm)
		return
	}

	res := map[string]string{
		"status": "Succesfully update",
	}
	ctx.JSON(http.StatusCreated, res)
}

func (c *Context) delete(ctx *gin.Context) {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
		return
	}

	queryText := fmt.Sprintf("DELETE FROM paslon where id = 17")

	_, err = db.ExecContext(ctx, queryText)

	res := map[string]string{
		"status": "Succesfully delete",
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *Context) get(ctx *gin.Context) {

	tpss := map[string]string{
		"status": "Succesfully",
	}
	ctx.JSON(http.StatusOK, tpss)
}
