package im

import "github.com/wuyan94zl/chat"

func JoinChannelIds(uid uint64, channelIds ...string) {
	chart.JoinChannelIds(uid, channelIds...)
}

func SendMessageToUid(uid, toUId uint64, msg string, tp uint8) {
	chart.SendMessageToUid(uid, toUId, msg, tp)
}

func SendMessageToUids(uid uint64, msg string, tp uint8, toUIds ...uint64) {
	chart.SendMessageToUids(uid, msg, tp, toUIds...)
}

func SendMessageToChannelIds(uid uint64, msg string, tp uint8, channelIds ...string) {
	chart.SendMessageToChannelIds(uid, msg, tp, channelIds...)
}
