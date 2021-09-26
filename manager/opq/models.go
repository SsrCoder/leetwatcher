package opq

import "github.com/mcoo/OPQBot"

type GroupMessageFunc func(botQQ int64, packet *OPQBot.GroupMsgPack)

type Group = struct {
	GroupID          int64  `json:"GroupId"`
	GroupMemberCount int64  `json:"GroupMemberCount"`
	GroupName        string `json:"GroupName"`
	GroupNotice      string `json:"GroupNotice"`
	GroupOwner       int64  `json:"GroupOwner"`
	GroupTotalCount  int    `json:"GroupTotalCount"`
}
