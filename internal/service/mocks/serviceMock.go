package mocks

type NotificationServiceMock struct {
	SendNotificationFunc func(notificationType, userID, message string) error
}

func (m *NotificationServiceMock) SendNotification(notificationType, userID, message string) error {
	if m.SendNotificationFunc != nil {
		return m.SendNotificationFunc(notificationType, userID, message)
	}
	return nil
}
