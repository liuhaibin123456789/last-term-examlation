package model

import (
	"fmt"
	"math"
)

//对应棋子
const (
	Space = iota

	RedChe
	RedMa
	RedXiang
	RedSHi
	RedShuai
	RedPao
	RedBing

	BlackChe = 3 + iota
	BlackMa
	BlackXiang
	BlackSHi
	BlackShuai
	BlackPao
	BlackBing
)

//ChessboardAbscissa 棋盘横纵坐标
type ChessboardAbscissa [10][9]int

//NewChessboardAbscissa 初始化一个棋子的横纵坐标
func NewChessboardAbscissa() *ChessboardAbscissa {
	//初始化棋局
	ca := [10][9]int{
		{BlackChe, BlackMa, BlackXiang, BlackSHi, BlackShuai, BlackSHi, BlackXiang, BlackMa, BlackChe},
		{Space, Space, Space, Space, Space, Space, Space, Space, Space},
		{Space, BlackPao, Space, Space, Space, Space, Space, BlackPao, Space},
		{BlackBing, Space, BlackBing, Space, BlackBing, Space, BlackBing, Space, BlackBing},
		{Space, Space, Space, Space, Space, Space, Space, Space, Space},
		{},
		{},
		{},
		{},
		{},
	}
	//棋盘对称，相同棋子不同颜色差值为10
	for i := 0; i < 5; i++ {
		for j := 0; j < 9; j++ {
			if ca[i][j] != Space {
				ca[9-i][j] = ca[i][j] - 10
			} else {
				ca[9-i][j] = Space
			}

		}
	}
	fmt.Println("初始化棋盘：", ca)
	return (*ChessboardAbscissa)(&ca)
}

// IsRed 是否是红棋
func (ca *ChessboardAbscissa) IsRed(x, y int) bool {
	return ca[x][y] <= RedBing && ca[x][y] >= RedChe
}

//IsBlack 是否是黑棋
func (ca *ChessboardAbscissa) IsBlack(x, y int) bool {
	return ca[x][y] <= BlackBing && ca[x][y] >= BlackChe
}

//IsSpace 是否是空位无子
func (ca *ChessboardAbscissa) IsSpace(x, y int) bool {
	return ca[x][y] == Space
}

//DropChess 吃子覆盖，不吃子填空位
func (ca *ChessboardAbscissa) DropChess(fromX, fromY, toX, toY int) {
	//todo 暂存吃掉的棋子,该需求需要拓展成结构体才行，或者数组扩容

	ca[toX][toY] = ca[fromX][fromY]
	ca[fromX][fromY] = Space
}

//RegretChess 悔棋
func (ca *ChessboardAbscissa) RegretChess(fromX, fromY, toX, toY, eatedChess int) {
	ca[fromX][fromY] = ca[toX][toY]
	ca[toX][toY] = eatedChess
}

//ChessRule 检查传入棋子的走法是否合理(包括吃子与不吃子的情况)——注意吃了子一定返回true
func (ca *ChessboardAbscissa) ChessRule(whichChess, fromY, fromX, toY, toX int) (res bool) {
	//横坐标绝对值之差
	xDifference := int(math.Abs(float64(toX - fromX)))
	//纵坐标绝对值之差
	yDifference := int(math.Abs(float64(toY - fromY)))
	switch whichChess {
	case BlackShuai:
		if ca.IsBlack(toX, toY) { //目标位置为自己的棋子不能走
			return false
		} else { //吃子或占空位
			if toY > 2 || toX < 3 || toX > 5 { //出了九宫格
				return false
			} else if xDifference == 1 && toY == fromY {
				//更新位置
				ca.DropChess(fromX, fromY, toX, toY)
				return true
			} else {
				return false
			}
		}
	case RedShuai:
		if ca.IsRed(toX, toY) { //目标位置是红棋不可以走
			return false
		} else { //目标位置为黑棋或者空位，可以走，坐标都可以更新
			if toY < 7 || toX < 3 || toX > 5 { //出了九宫格
				return false
			} else if yDifference == 1 && toX == fromX {
				//只能走一格
				ca.DropChess(fromX, fromY, toX, toY)
				return true
			} else {
				return false
			}
		}
	case BlackSHi:
		if ca.IsBlack(toX, toY) {
			return false
		} else {
			if toY > 2 || toX < 3 || toX > 5 { //出了九宫格
				return false
			} else if xDifference == 1 && yDifference == 1 {
				//走斜线，直走一格
				ca.DropChess(fromX, fromY, toX, toY)
				return true
			} else {
				return false
			}
		}

	case RedSHi:
		if ca.IsRed(toX, toY) {
			return false
		} else {
			if toY < 7 || toX < 3 || toX > 5 { //出了九宫格
				return false
			} else if yDifference == 1 && xDifference == 1 {
				//走斜线，直走一格
				ca.DropChess(fromX, fromY, toX, toY)
				return true
			} else {
				return false
			}
		}
	case BlackXiang:
		if ca.IsBlack(toX, toY) {
			return false
		} else {
			if toY > 4 { //过河了
				return false
			} else if yDifference == 2 && xDifference == 2 {
				//走"田"字
				centerX := (toX + fromX) / 2
				centerY := (toY + fromY) / 2
				if ca[centerX][centerY] != Space { // 象眼处有棋子
					return false
				}
				//走棋
				ca.DropChess(fromX, fromY, toX, toY)
				return true
			} else {
				return false
			}
		}
	case RedXiang:
		if ca.IsRed(toX, toY) {
			return false
		} else {
			if toY < 5 { //过河
				return false
			} else if yDifference == 2 && xDifference == 2 {
				//走"田"字
				centerX := (toX + fromX) / 2
				centerY := (toY + fromY) / 2
				if ca[centerX][centerY] != Space { // 象眼处有棋子
					return false
				}
				ca.DropChess(fromX, fromY, toX, toY)
				return true
			} else {
				return false
			}
		}

	case BlackMa:
		if ca.IsBlack(toX, toY) {
			return false
		} else {
			if yDifference == 2 && xDifference == 1 {
				if toY-fromY == 2 {
					if ca[fromX][toY-1] == Space { //马蹄处没有棋子
						ca.DropChess(fromX, fromY, toX, toY)
						return true
					}
				} else if toY-fromY == -2 {
					if ca[fromX][fromY-1] == Space { //马蹄处没有棋子
						ca.DropChess(fromX, fromY, toX, toY)
						return true
					}
				}
				return false
			} else if yDifference == 1 && xDifference == 2 {
				if toX-fromX == 2 {
					if ca[toX-1][fromY] == Space { //马蹄处没有棋子
						ca.DropChess(fromX, fromY, toX, toY)
						return true
					}
				} else if toX-fromX == -2 {
					if ca[fromX-1][fromY] == Space { //马蹄处没有棋子
						ca.DropChess(fromX, fromY, toX, toY)
						return true
					}
				}
				return false
			} else {
				return false
			}
		}

	case RedMa:
		if ca.IsRed(toX, toY) {
			return false
		} else {
			if yDifference == 2 && xDifference == 1 {
				if toY-fromY == 2 {
					if ca[fromX][toY-1] == Space { //马蹄处没有棋子
						ca.DropChess(fromX, fromY, toX, toY)
						return true
					}
				} else if toY-fromY == -2 {
					if ca[fromX][fromY-1] == Space { //马蹄处没有棋子
						ca.DropChess(fromX, fromY, toX, toY)
						return true
					}
				}
				return false
			} else if yDifference == 1 && xDifference == 2 {
				if toX-fromX == 2 {
					if ca[toX-1][fromY] == Space { //马蹄处没有棋子
						ca.DropChess(fromX, fromY, toX, toY)
						return true
					}
				} else if toX-fromX == -2 {
					if ca[fromX-1][fromY] == Space { //马蹄处没有棋子
						ca.DropChess(fromX, fromY, toX, toY)
						return true
					}
				}
				return false
			}
			return false
		}
	case BlackChe:
		if ca.IsBlack(toX, toY) {
			return false
		} else {
			if toX != fromX && toY != fromY { //没在竖线或横线上
				return false
			} else { //扫描是否有中间棋子
				if toY == fromY { //扫描一列
					start := 0
					if toX > fromX {
						start = fromX
					} else {
						start = toX
					}
					for i := start; i < xDifference; i++ { //扫描中间是否有棋子
						if ca[i][toY] != Space {
							return false
						}
					}
					ca.DropChess(fromX, fromY, toX, toY)
					return true
				} else if toX == fromX { //扫描一行
					start := 0
					if toY > fromY {
						start = fromY
					} else {
						start = toY
					}
					for i := start; i < yDifference; i++ { //扫描中间是否有棋子
						if ca[i][toY] != Space {
							return false
						}
					}
					ca.DropChess(fromX, fromY, toX, toY)
					return true
				}
			}
		}
	case RedChe:
		if ca.IsRed(toX, toY) {
			return false
		} else {
			if toX != fromX && toY != fromY { //没在竖线或横线上
				return false
			} else { //扫描是否有中间棋子
				if toY == fromY { //扫描一列
					start := 0
					if toX > fromX {
						start = fromX
					} else {
						start = toX
					}
					for i := start; i < xDifference; i++ { //扫描中间是否有棋子
						if ca[i][toY] != Space {
							return false
						}
					}
					ca.DropChess(fromX, fromY, toX, toY)
					return true
				} else if toX == fromX { //扫描一行
					start := 0
					if toY > fromY {
						start = fromY
					} else {
						start = toY
					}
					for i := start; i < yDifference; i++ { //扫描中间是否有棋子
						if ca[i][toY] != Space {
							return false
						}
					}
					ca.DropChess(fromX, fromY, toX, toY)
					return true
				}
			}
		}

	case BlackPao:
		if ca.IsBlack(toX, toY) {
			return false
		} else if ca.IsSpace(toX, toY) { //目标位置为空,同车
			if toX != fromX && toY != fromY { //没在竖线或横线上
				return false
			} else { //扫描是否有中间棋子
				if toY == fromY { //扫描一列
					start := 0
					if toX > fromX {
						start = fromX
					} else {
						start = toX
					}
					for i := start; i < xDifference; i++ { //扫描中间是否有棋子
						if ca[i][toY] != Space {
							return false
						}
					}
					ca.DropChess(fromX, fromY, toX, toY)
					return true
				} else if toX == fromX { //扫描一行
					start := 0
					if toY > fromY {
						start = fromY
					} else {
						start = toY
					}
					for i := start; i < yDifference; i++ { //扫描中间是否有棋子
						if ca[i][toY] != Space {
							return false
						}
					}
					ca.DropChess(fromX, fromY, toX, toY)
					return true
				}
			}
			//吃子，中间需要有一个子，并且后面的子必须是对方棋子
		} else if ca.IsRed(toX, toY) { //目标位置含有红棋才可以吃子
			centerNum := 0                    //中间棋子个数
			if toX != fromX && toY != fromY { //没在竖线或横线上
				return false
			} else { //扫描是否有中间棋子
				if toY == fromY { //扫描一列
					start := 0
					if toX > fromX {
						start = fromX
					} else {
						start = toX
					}
					for i := start; i < xDifference; i++ { //扫描中间是否有棋子
						if ca[i][toY] != Space {
							centerNum++
						}
					}
					if centerNum == 1 { //中间有一个棋子才能吃子
						//更新棋盘
						ca.DropChess(fromX, fromY, toX, toY)
						return true
					}
					return false
				} else if toX == fromX { //扫描一行
					start := 0
					if toY > fromY {
						start = fromY
					} else {
						start = toY
					}
					for i := start; i < yDifference; i++ { //扫描中间是否有棋子
						if ca[i][toY] != Space {
							centerNum++
						}
					}
					if centerNum == 1 { //中间有一个棋子才能吃子
						//更新棋盘
						ca.DropChess(fromX, fromY, toX, toY)
						return true
					}
					return false
				}
			}
		}
		return false
	case RedPao:
		if ca.IsRed(toX, toY) {
			return false
		} else if ca.IsSpace(toX, toY) { //目标位置为空
			if toX != fromX && toY != fromY { //没在竖线或横线上
				return false
			} else { //扫描是否有中间棋子
				if toY == fromY { //扫描一列
					start := 0
					if toX > fromX {
						start = fromX
					} else {
						start = toX
					}
					for i := start; i < xDifference; i++ { //扫描中间是否有棋子
						if ca[i][toY] != Space {
							return false
						}
					}
					return true
				} else if toX == fromX { //扫描一行
					start := 0
					if toY > fromY {
						start = fromY
					} else {
						start = toY
					}
					for i := start; i < yDifference; i++ { //扫描中间是否有棋子
						if ca[i][toY] != Space {
							return false
						}
					}
					ca.DropChess(fromX, fromY, toX, toY)
					return true
				}
			}
			//吃子，中间需要有一个子，并且后面的子必须是对方棋子
		} else if ca.IsBlack(toX, toY) { //目标位置含有黑棋才可以吃子
			centerNum := 0                    //中间棋子个数
			if toX != fromX && toY != fromY { //没在竖线或横线上
				return false
			} else { //扫描是否有中间棋子
				if toY == fromY { //扫描一列
					start := 0
					if toX > fromX {
						start = fromX
					} else {
						start = toX
					}
					for i := start; i < xDifference; i++ { //扫描中间是否有棋子
						if ca[i][toY] != Space {
							centerNum++
						}
					}
					if centerNum == 1 { //中间有一个棋子才能吃子
						ca.DropChess(fromX, fromY, toX, toY)
						return true
					}
					return false
				} else if toX == fromX { //扫描一行
					start := 0
					if toY > fromY {
						start = fromY
					} else {
						start = toY
					}
					for i := start; i < yDifference; i++ { //扫描中间是否有棋子
						if ca[i][toY] != Space {
							centerNum++
						}
					}
					if centerNum == 1 { //中间有一个棋子才能吃子
						//更新棋盘
						ca.DropChess(fromX, fromY, toX, toY)
						return true
					}
					return false
				}
			}
		}
		return false
	case BlackBing:
		if ca.IsBlack(toX, toY) {
			return false
		} else {
			if xDifference == 1 && toY == fromY { //纵向走一步
				if fromX > toX { //黑棋兵不能后退
					return false
				} else if fromX < toX { //不论过河，都可以往前走
					ca.DropChess(fromX, fromY, toX, toY)
					return true
				}
			} else if yDifference == 1 && toX == fromX { //横向走一步,需要过河
				if toX > 4 { //过河
					ca.DropChess(fromX, fromY, toX, toY)
					return true
				}
			}
			return false
		}
	case RedBing:
		if ca.IsRed(toX, toY) {
			return false
		} else {
			if xDifference == 1 && toY == fromY { //纵向走一步
				if fromX < toX { //红棋兵不能后退
					return false
				} else if fromX > toX { //不论过河，都可以往前走
					ca.DropChess(fromX, fromY, toX, toY)
					return true
				}
			} else if yDifference == 1 && toX == fromX { //横向走一步,需要过河
				if toX < 5 { //过河
					ca.DropChess(fromX, fromY, toX, toY)
					return true
				}
			}
			return false
		}
	default:
		fmt.Println("无效的棋子")
		return false
	}
	return false
}
