package constant

const (
	SignalingNotificationBegin               = 1600
	SignalingNotification                    = 1601
	RoomParticipantsConnectedNotification    = 1602
	RoomParticipantsDisconnectedNotification = 1603
	StreamChangedNotification                = 1604
	CustomSignalNotification                 = 1605

	SignalingNotificationEnd = 1699
)

const (
	MeetingNotEnd = 0
	MeetingEnd    = 1

	SingleCreateVideoMeeting   = 0
	SingleReceiveVideoMeeting  = 1
	SingleRefuseVideoMeeting   = 2
	SingleHangupVideoMeeting   = 3
	SingleNoAnswerVideoMeeting = 4
)
