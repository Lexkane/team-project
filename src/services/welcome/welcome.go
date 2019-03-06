package welcome

import (
	"net/http"

	"github.com/mshto/team-project/src/entity"
	"github.com/mshto/team-project/src/services/common"
)

// GetWelcomeHandler get policies for partner
func GetWelcomeHandler(w http.ResponseWriter, r *http.Request) {
	common.RenderStatusOK(w, r, entity.Welcome{Message: "Hello World"})
}
