package controller

import (
	"github.com/MSyahidAlFajri/backend"
	"github.com/gocroot/penggajian/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/whatsauth/whatsauth"
)

var DataKaryawan = "karyawan"
var DataHonor = "honor"
var DataTeam = "team"
var DataJob = "job"

func WsWhatsAuthQR(c *websocket.Conn) {
	whatsauth.RunSocket(c, config.PublicKey, config.Usertables[:], config.Ulbimariaconn)
}

func PostWhatsAuthRequest(c *fiber.Ctx) error {
	if string(c.Request().Host()) == config.Internalhost {
		var req whatsauth.WhatsauthRequest
		err := c.BodyParser(&req)
		if err != nil {
			return err
		}
		ntfbtn := whatsauth.RunModuleLegacy(req, config.PrivateKey, config.Usertables[:], config.Ulbimariaconn)
		return c.JSON(ntfbtn)
	} else {
		var ws whatsauth.WhatsauthStatus
		ws.Status = string(c.Request().Host())
		return c.JSON(ws)
	}

}

func GetKaryawan(c *fiber.Ctx) error {
	getstatus := backend.GetDataKaryawan("Aktif")
	return c.JSON(getstatus)
}

func GetHonor(c *fiber.Ctx) error {
	getstatus := backend.GetDataHonor("Aktif")
	return c.JSON(getstatus)
}
func GetTeam(c *fiber.Ctx) error {
	getstats := backend.GetDataTeam("Uzumaki Memet")
	return c.JSON(getstats)
}
func GetJob(c *fiber.Ctx) error {
	getnamajob := backend.GetDataJob("Staff Administrasi")
	return c.JSON(getnamajob)
}
