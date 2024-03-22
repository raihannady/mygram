package main

import (
	"mygram/router"

	"github.com/gin-gonic/gin"
)

func main() {
    // Inisialisasi router utama
    mainRouter := gin.Default()

    // Gabungkan rute dari setiap entitas
    router.UserRouter(mainRouter)
    router.PhotosRouter(mainRouter)
    router.CommentRouter(mainRouter)
    router.SocmedRouter(mainRouter)
    // Sisanya entitas lainnya...

    // Menjalankan server
    mainRouter.Run(":8080")
}
