package app

type ServiceContainer struct {
	VideoChat *VideoChatService
}

func NewApp() *ServiceContainer {
	wsService := NewVideoChatService()

	return &ServiceContainer{
		VideoChat: wsService,
	}
}
