package opq

import "github.com/mcoo/OPQBot"

type GroupMessageFunc func(botQQ int64, packet *OPQBot.GroupMsgPack)
