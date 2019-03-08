package welcome

import (
	"../../entity"
	"../common"
	"net/http"
)

// GetWelcomeHandler get policies for partner
func GetWelcomeHandler(w http.ResponseWriter, r *http.Request) {
	common.RenderJSON(w, r, entity.Welcome{Message: "Hello World"})
}
