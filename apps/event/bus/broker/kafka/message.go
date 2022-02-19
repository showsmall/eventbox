package kafka

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Shopify/sarama"

	"github.com/infraboard/eventbox/apps/event"
)

// 把我们事件 --> kafka的消息
func newProducerMessage(event *event.Event) (*sarama.ProducerMessage, error) {
	bytes, err := json.Marshal(event)
	if err != nil {
		return nil, fmt.Errorf("mashal event to json error, %s", err)
	}

	message := &sarama.ProducerMessage{
		Value: sarama.ByteEncoder(bytes),
	}

	if p, ok := event.GetMetaKey(MetaKafkaPartition); ok {
		intp, err := getInt32(p)
		if err != nil {
			return nil, err
		}
		message.Partition = intp
	}

	if key, ok := event.GetMetaKey(MetaKafkaKey); ok {
		strKey, err := getString(key)
		if err != nil {
			return nil, err
		}
		message.Key = sarama.StringEncoder(strKey)
	}

	if headers, ok := event.GetMetaKey(MetaKafkaHeaders); ok {
		strHeader, err := getString(headers)
		if err != nil {
			return nil, err
		}

		hdrs := []sarama.RecordHeader{}
		arrHdrs := strings.Split(strHeader, ",")
		for _, h := range arrHdrs {
			header := strings.Split(h, ":")
			if len(header) != 2 {
				return nil, fmt.Errorf("-header should be key:value. Example: -headers=foo:bar,bar:foo")
			}

			hdrs = append(hdrs, sarama.RecordHeader{
				Key:   []byte(header[0]),
				Value: []byte(header[1]),
			})
		}

		if len(hdrs) != 0 {
			message.Headers = hdrs
		}
	}

	return message, nil
}

func getInt32(data interface{}) (int32, error) {
	switch v := data.(type) {
	case int:
		return int32(v), nil
	case int32:
		return v, nil
	case int64:
		return int32(v), nil
	default:
		return 0, fmt.Errorf("not an number(int, int32, int64)")
	}
}

func getString(data interface{}) (string, error) {
	switch v := data.(type) {
	case string:
		return v, nil
	default:
		return "", fmt.Errorf("not an string")
	}
}
