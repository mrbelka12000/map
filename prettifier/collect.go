package prettifier

import (
	"github.com/graincomg/graincom_logistik/models"
)

func Сonvenience(mp models.Oms) {
	for _, node := range mp.Node {
		if node.MapTag == nil {
			node.MapTag = make(map[string]string)
		}
		for _, tag := range node.Tag {
			node.MapTag[tag.K] = tag.V
		}
		models.ConvertedNodes[node.ID] = node
	}

}

func GetNodeByID(id string) (models.Node, bool) {
	val, ok := models.ConvertedNodes[id]
	return val, ok
}
