package config

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

var origins = []string{
	"https://MSyahidAlFajri.github.io",
	"https://msyahidalfajri.github.io",
	"https://penggajian.herokuapp.com/",
	"https://gocroot.github.io/",
	"http://127.0.0.1:5500",
	"http://127.0.0.1:5501",
	"https://auth.ulbi.ac.id",
	"https://sip.ulbi.ac.id",
	"https://euis.ulbi.ac.id",
	"https://home.ulbi.ac.id",
	"https://alpha.ulbi.ac.id",
	"https://dias.ulbi.ac.id",
	"https://iteung.ulbi.ac.id",
	"https://whatsauth.github.io/",
}

var Internalhost string = os.Getenv("INTERNALHOST") + ":" + os.Getenv("PORT")

var Cors = cors.Config{
	AllowOrigins:     strings.Join(origins[:], ","),
	AllowHeaders:     "Origin",
	ExposeHeaders:    "Content-Length",
	AllowCredentials: true,
}
