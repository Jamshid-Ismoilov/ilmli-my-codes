package wikipedia

import (
	"fmt"

	"cgt.name/pkg/go-mwclient"
	"github.com/gin-gonic/gin"
)

func GetPages(ctx *gin.Context) {

	res := ctx.Param("names")
	w, err := mwclient.New("https://en.wikipedia.org/w/api.php", "ilmli")
	if err != nil {
		fmt.Println(err)
	}

	resp, err := w.GetPagesByName(res)
	if err != nil {
		fmt.Println(err)
	}

	ctx.JSON(200, resp)
}
