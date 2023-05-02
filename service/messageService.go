package service

import (
	"TikTokLite/dao"
	. "TikTokLite/message"
	"TikTokLite/model"
)

//func MessageSendService(fromUserId int64, toUserId int64, msg string) (*DouyinSendMessageResponse, error) {
//	// 创建一个通道，大多数API都是用过该通道操作的。
//	conn := common.GetRabbitMq()
//	ch, err := conn.Channel()
//	if err != nil {
//		fmt.Printf("open a channel failed, err:%v\n", err)
//		return nil, err
//	}
//	defer ch.Close()
//	// 要发送，我们必须声明要发送到的队列。
//	q, err := ch.QueueDeclare(
//		"task_queue", // name
//		true,         // 持久的
//		false,        // delete when unused
//		false,        // 独有的
//		false,        // no-wait
//		nil,          // arguments
//	)
//	if err != nil {
//		fmt.Printf("declare a queue failed, err:%v\n", err)
//		return nil, err
//	}
//	// 然后我们可以将消息发布到声明的队列
//	err = ch.Publish(
//		"",     // exchange
//		q.Name, // routing key
//		false,  // 立即
//		false,  // 强制
//		amqp.Publishing{
//			DeliveryMode: amqp.Persistent, // 持久
//			ContentType:  "text/plain",
//			Body:         []byte(msg),
//		})
//	if err != nil {
//		fmt.Printf("publish a message failed, err:%v\n", err)
//		return nil, err
//	}
//	log.Printf("Send Message Success!")
//
//	_, err = ch.Consume(
//		q.Name, // queue
//		"",     // consumer
//		true,   // auto-ack
//		false,  // exclusive
//		false,  // no-local
//		false,  // no-wait
//		nil,    // args
//	)
//	if err != nil {
//		fmt.Printf("Failed to register a consumer,err:%v\n", err)
//		return nil, err
//	}
//
//	forever := make(chan bool)
//
//	go func() {
//		// 读写数据库的操作并发执行
//		err = dao.InsertMessage(fromUserId, toUserId, msg)
//		if err != nil {
//			forever <- false
//		}
//		forever <- true
//	}()
//
//	if <-forever {
//		return &DouyinSendMessageResponse{Response: Response{
//			StatusCode: 0,
//		}}, nil
//	} else {
//		return nil, errors.New("insert Message failed")
//	}
//}

func MessageSendService(fromUserId int64, toUserId int64, msg string) (*DouyinSendMessageResponse, error) {
	err := dao.InsertMessage(fromUserId, toUserId, msg)
	if err != nil {
		return nil, err
	}
	return &DouyinSendMessageResponse{Response: SuccessResponse}, nil
}

func GetMessageList(userId int64, toUserId int64, preMsgTime int64) (*DouyinMessageChatResponse, error) {
	messageList, err := dao.GetMessageList(userId, toUserId, preMsgTime)
	if err != nil {
		return nil, err
	}
	return &DouyinMessageChatResponse{
		Response:    SuccessResponse,
		MessageList: MessageMessageList(messageList),
	}, nil
}

func MessageMessageList(MList []model.Message) []Message {
	var MessageList []Message
	for _, message := range MList {
		MessageList = append(MessageList, Message{
			Id:         message.MessageId,
			ToUserId:   message.ToUserId,
			FromUserId: message.FromUserId,
			Content:    message.Content,
			CreateTime: message.CreatTime,
		})
	}
	return MessageList
}
