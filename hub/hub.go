package hub

import (
	"github.com/tapvanvn/go-dashboard/entity"
	"github.com/tapvanvn/godashboard"
)

type Hub struct {
	Item          *entity.Item
	LastWriteTime int64
}

func (h *Hub) Signal(params map[string]godashboard.Param) {

}
