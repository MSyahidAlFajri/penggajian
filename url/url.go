package url

import (
	"github.com/gocroot/penggajian/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Web(page *fiber.App) {
	page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth
	page.Get("/karyawan", controller.GetKaryawan)
	page.Get("/honor", controller.GetHonor)
	page.Get("/team", controller.GetTeam)
	page.Get("/Job", controller.GetJob)
	// page.Get("/tambah", controller.InsertKaryawan)
}
