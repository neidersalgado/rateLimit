package request

// NotificationRequest define la estructura esperada de la solicitud de notificación
type NotificationRequest struct {
	Type    string `json:"type" binding:"required"`
	UserID  string `json:"userId" binding:"required"`
	Message string `json:"message" binding:"required"`
}
