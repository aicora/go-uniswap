package libraries

import (
	"math/big"

	"github.com/aicora/go-uniswap/utils"
)

type Slot0 struct {
    SqrtPriceX96 *big.Int
    Tick         int32
    ProtocolFee  utils.ProtocolFee
    LPFee        utils.LPFee
}