package v1

import (
	"net/http"

	httpUtil "github.com/tanmaij/friend-management/pkg/utils/http"
)

// Example handles an example request
func (handler *Handler) Example(_ *http.Request, w http.ResponseWriter) error {
	httpUtil.WriteString("Ok", http.StatusOK, w)
	return nil
}
