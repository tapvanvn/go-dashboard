package hub

import (
	"github.com/tapvanvn/go-dashboard/entity"
)

type Hub struct {
	Item          *entity.Item
	LastWriteTime int64
}

func (h *Hub) Signal(params map[string][]entity.Param) {

}
