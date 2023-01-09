package gcp

type AudioProfile string

const (
	WearableClassDevice               AudioProfile = "wearable-class-device"
	HandsetClassDevice                             = "handset-class-device"
	HeadphoneClassDevice                           = "headphone-class-device"
	SmallBluetoothSpeakerClassDevice               = "small-bluetooth-speaker-class-device"
	MediumBluetoothSpeakerClassDevice              = "medium-bluetooth-speaker-class-device"
	LargeHomeEntertainmentClassDevice              = "large-home-entertainment-class-device"
	LargeAutomotiveClassDevice                     = "large-automotive-class-device"
	TelephonyClassApplication                      = "telephony-class-application"
)
