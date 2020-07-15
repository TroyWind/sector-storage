package dsectorslog

import (
	"github.com/filecoin-project/sector-storage/common/util"
	"go.uber.org/zap"
)

var L *zap.Logger

func init() {
	L = util.GetXDebugLog("sector-store")
}
