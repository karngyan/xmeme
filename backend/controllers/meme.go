package controllers

import (
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
// @Param name query string true "name of the owner"
// @Param url query string true "url of the meme"
// @Param caption query string true "caption of the meme"
// @Success 200 {object} models.IResponse "Meme created successfully"
// @Failure 406 Not Acceptable
// @Failure 409 Conflict
// @Failure 500 Internal Server Error
// @router /memes [post]
func (mc *MemeController) CreateMeme() {

	name := strings.TrimSpace(mc.GetString("name"))
	url := strings.TrimSpace(mc.GetString("url"))
	caption := strings.TrimSpace(mc.GetString("caption"))

	if name == "" || url == "" || caption == "" {
		mc.Ctx.Output.SetStatus(http.StatusNotAcceptable)
		return
	}

	// already exists
	if prs := models.AllMemes().Filter("name", name).Filter("url", url).Filter("caption", caption).Exist(); prs {
		mc.Ctx.Output.SetStatus(http.StatusConflict)
		return
	}

	m := models.Meme{
		Name:    name,
		Url:     url,
		Caption: caption,
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
