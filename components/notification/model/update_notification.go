package notimodel

type UpdateNotification struct {

	// Subject is a NotificationObject that represent the object do the Action
	Subject *NotificationObject `bson:"subject,omitempty" json:"subject,omitempty"`

	// Direct is a NotificationObject that represent the object was directly affected by the Action
	Direct *NotificationObject `bson:"direct,omitempty" json:"direct,omitempty"`

	// Indirect is a NotificationObject that represent the object was indirectly affected by the Action
	Indirect *NotificationObject `bson:"indirect,omitempty" json:"indirect,omitempty"`

	// Indirect is a NotificationObject that represent the object was appear in the action with a prep (in, for, of)
	Prep *NotificationObject `bson:"prep,omitempty" json:"prep,omitempty"`
}

func (UpdateNotification) CollectionName() string {
	return "notifications"
}
