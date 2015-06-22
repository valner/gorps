package types
type EFigureType byte
const (
    // Figure type
    EPAPER EFigureType = iota
    ESCISSORS = iota
    EROCK = iota
    ENONE = iota
)

var efigureTypes = [...]string {
    "Paper",
    "Scissors",
    "Rock",
}

func (figureType EFigureType) String() string {
    return efigureTypes[figureType]
}




type EMoveResult byte

const (
    EMOVE EMoveResult = iota
    EDEAD = iota
    EWRONG_TURN = iota
    EGAME_ENDED = iota
    EDRAW_WAITED = iota
)
