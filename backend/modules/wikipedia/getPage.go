package wikipedia

import (
	"fmt"

	"cgt.name/pkg/go-mwclient"
	"github.com/gin-gonic/gin"
)

func GetPage(ctx *gin.Context) {

	res := ctx.Param("name")
	w, err := mwclient.New("https://en.wikipedia.org/w/api.php", "ilmli")
	if err != nil {
		fmt.Println(err)
	}

	resp, _, err := w.GetPageByName(res)
	if err != nil {
		fmt.Println(err)
	}

	ctx.JSON(200, resp)
}
