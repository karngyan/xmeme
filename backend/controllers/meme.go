package controllers

import (
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
	"net/http"
	"strings"
	"xmeme/models"
)

type MemeController struct {
	beego.Controller
}

// @Title GetAllMemes
// @Description Get Latest 100 Memes (later can be paginated)
// @Success 200 {array} []models.Meme "Retrieved all memes"
// @Failure 500 Internal Server Error
// @router /memes [get]
func (mc *MemeController) GetAllMemes() {

	memes := make([]*models.Meme, 0)
	limit := 100
	if _, err := models.AllMemes().OrderBy("-id").Limit(limit).RelatedSel().All(&memes); err == nil {
		mc.Ctx.Output.SetStatus(http.StatusOK)
		_ = mc.Ctx.Output.JSON(memes, true, false)
		return
	} else {
		mc.Ctx.Output.SetStatus(http.StatusInternalServerError)
		return
	}
}

// @Title GetMeme
// @Description Get Meme by Id
// @Param	mid	path int true	"the mid you want to get"
// @Success 200 {object} models.Meme "Retrieved Meme from mid successfully"
// @Failure 404 Meme not found
// @Failure 406 :mid is empty
// @router /memes/:mid [get]
func (mc *MemeController) GetMeme() {

	var m models.Meme

	mid := mc.Ctx.Input.Param(":mid")
	if mid != "" {
		if err := models.AllMemes().Filter("id", mid).RelatedSel().One(&m); err == nil {
			mc.Ctx.Output.SetStatus(http.StatusOK)
			_ = mc.Ctx.Output.JSON(m, true, false)
		} else {
			mc.Ctx.Output.SetStatus(http.StatusNotFound)
			return
		}
	} else {
		mc.Ctx.Output.SetStatus(http.StatusNotAcceptable)
		return
	}
}

// @Title CreateMeme
// @Description Create New XMeme
// @Param body body models.Meme true "Meme Request"
// @Success 200 {object} models.IResponse "Meme created successfully"
// @Failure 406 Not Acceptable
// @Failure 422 Unprocessable entity
// @Failure 409 Conflict
// @Failure 500 Internal Server Error
// @router /memes [post]
func (mc *MemeController) CreateMeme() {

	var m models.Meme

	if err := json.Unmarshal(mc.Ctx.Input.RequestBody, &m); err != nil {
		mc.Ctx.Output.SetStatus(http.StatusUnprocessableEntity)
		return
	}

	m.Name = strings.TrimSpace(m.Name)
	m.Url = strings.TrimSpace(m.Url)
	m.Caption = strings.TrimSpace(m.Caption)

	if m.Name == "" || m.Url == "" || m.Caption == "" {
		mc.Ctx.Output.SetStatus(http.StatusNotAcceptable)
		return
	}

	// already exists
	if prs := models.AllMemes().Filter("name", m.Name).Filter("url", m.Url).Filter("caption", m.Caption).Exist(); prs {
		mc.Ctx.Output.SetStatus(http.StatusConflict)
		return
	}


	if err := m.Insert(); err == nil {
		// successful insert
		ir := models.IResponse{
			Id: m.Id,
		}
		mc.Ctx.Output.SetStatus(http.StatusCreated)
		_ = mc.Ctx.Output.JSON(ir, true, false)
		return
	} else {
		mc.Ctx.Output.SetStatus(http.StatusInternalServerError)
		return
	}

}
