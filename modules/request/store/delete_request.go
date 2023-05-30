package requeststore

import (
	"context"
	"cs_chat_app_server/common"
	requestmdl "cs_chat_app_server/modules/request/model"
)

func (s *mongoStore) DeleteRequest(ctx context.Context, filter map[string]interface{}) error {
	_, err := s.database.Collection(requestmdl.Request{}.CollectionName()).DeleteOne(ctx, filter)
	if err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
