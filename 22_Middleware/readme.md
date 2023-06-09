## Sri Widaningsih


#### Resume Middleware

1. Middleware
- Middleware adalah lapisan perangkat lunak yang berfungsi sebagai penghubung antara aplikasi dan server web. Middleware digunakan untuk memodifikasi permintaan dan tanggapan pada saat mereka melewati aplikasi web. Di dalam framework Echo, middleware adalah fungsi yang mengambil parameter Context dan HandlerFunc untuk memodifikasi atau mengambil tindakan tertentu pada permintaan dan tanggapan.

2. Echo Middleware
- Echo middleware adalah middleware yang digunakan pada framework web Echo, yang dikembangkan dengan bahasa pemrograman Go.


3. Logger middleware, digunakan untuk mencatat log pada setiap permintaan dan tanggapan yang dilakukan pada aplikasi web.

- contohnya yaitu
```js
import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

func main() {
	// Create a new instance of Echo
	e := echo.New()

	// Create a new file for logging
	f, err := os.Create("echo.log")
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Set up logger middleware with custom logger
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: f,
	}))

	// Start the server
	e.Start(":8000")
}

```
4. Auth middleware, digunakan untuk memverifikasi dan mengecek otorisasi pengguna.
```js
import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	// Create a new instance of Echo
	e := echo.New()

	// Set up auth middleware with username and password
	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "admin" && password == "secret" {
			return true, nil
		}
		return false, nil
	}))

	// Protected route
	e.GET("/admin", func(c echo.Context) error {
		return c.String(http.StatusOK, "You are authorized to see this page!")
	})

	// Start the server
	e.Start(":8000")
}

```
5. JWT middleware, digunakan untuk memvalidasi token JSON Web Token (JWT) yang diterima dari klien.
```js
import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	// Create a new instance of Echo
	e := echo.New()

	// Set up JWT middleware with secret key
	e.Use(middleware.JWT([]byte("secret")))

	// Protected route
	e.GET("/admin", func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		name := claims["name"].(string)
		return c.String(http.StatusOK, "Welcome "+name+"!")
	})

	// Start the server
	e.Start(":8000")
}

```

#### Untuk Tugas sudah terlampir pada folder screenshot dan script yang digunakan ada pada folder praktikum