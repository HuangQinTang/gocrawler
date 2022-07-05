package persist

import (
	"crawler/model"
	"crawler/persist"
	"github.com/olivere/elastic/v7"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaverService) Save(item model.SimpleInfo, result *string) error {
	docId, err := persist.Save(s.Client, s.Index, item)
	if err == nil {
		*result = docId
	}
	return err
}
