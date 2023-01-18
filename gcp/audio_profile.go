package gcp

type AudioProfile string

const (
	WearableClassDevice               AudioProfile = "wearable-class-device"
	HandsetClassDevice                AudioProfile = "handset-class-device"
	HeadphoneClassDevice              AudioProfile = "headphone-class-device"
	SmallBluetoothSpeakerClassDevice  AudioProfile = "small-bluetooth-speaker-class-device"
	MediumBluetoothSpeakerClassDevice AudioProfile = "medium-bluetooth-speaker-class-device"
	LargeHomeEntertainmentClassDevice AudioProfile = "large-home-entertainment-class-device"
	LargeAutomotiveClassDevice        AudioProfile = "large-automotive-class-device"
	TelephonyClassApplication         AudioProfile = "telephony-class-application"
)

const DefaultAudioProfile = HeadphoneClassDevice
