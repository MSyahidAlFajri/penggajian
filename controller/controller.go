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

func InsertData(c *fiber.Ctx) error {
	model := new(backend.Karyawan)
	tambahdata := backend.InsertKaryawan(config.MongoConn,
		model.Nama,
		model.Status,
		model.Jabatan,
		model.Gaji,
	)
	return c.JSON(tambahdata)
}

// func InsertKaryawan(c *fiber.Ctx) error {
// 	database := config.MongoConn
// 	var model backend.Karyawan
// 	if err := c.BodyParser(&model); err != nil {
// 		return err
// 	}
// 	Inserted := backend.InsertKaryawan(database,
// 		DataKaryawan,
// 		model.Nama,
// 		model.Status,
// 		model.Jabatan,
// 		model.Gaji,
// 	)
// 	return c.JSON(map[string]interface{}{
// 		"status":      http.StatusOK,
// 		"message":     "Data berhasil disimpan.",
// 		"inserted_id": Inserted,
// 	})
// }
