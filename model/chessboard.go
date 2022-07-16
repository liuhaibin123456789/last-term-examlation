package model

import (
	"fmt"
	"last-homework/global"
	"last-homework/tool"
	"math"
)

var MapResult = map[int]string{
	global.RedWinner:   "红棋赢了",
	global.BlackWinner: "黑棋赢了",
}

func GetResult(who int) string {
	return MapResult[who]
}

//ChessboardAbscissa 棋盘横纵坐标
type ChessboardAbscissa struct {
	Array  [10][9]int //棋盘
	Regret int        //悔棋暂存
	Winner chan int   //0表示红方赢，1表示黑方赢了
}

//NewChessboardAbscissa 初始化一个棋子的横纵坐标
func NewChessboardAbscissa() *ChessboardAbscissa {
	ca := new(ChessboardAbscissa)
	//初始化棋局
	arr := [10][9]int{
		{global.BlackChe, global.BlackMa, global.BlackXiang, global.BlackSHi, global.BlackShuai, global.BlackSHi, global.BlackXiang, global.BlackMa, global.BlackChe},
		{global.Space, global.Space, global.Space, global.Space, global.Space, global.Space, global.Space, global.Space, global.Space},
		{global.Space, global.BlackPao, global.Space, global.Space, global.Space, global.Space, global.Space, global.BlackPao, global.Space},
		{global.BlackBing, global.Space, global.BlackBing, global.Space, global.BlackBing, global.Space, global.BlackBing, global.Space, global.BlackBing},
		{global.Space, global.Space, global.Space, global.Space, global.Space, global.Space, global.Space, global.Space, global.Space},
		{},
		{},
		{},
		{},
		{},
	}
	//棋盘对称，相同棋子不同颜色差值为10
	for i := 0; i < 5; i++ {
		for j := 0; j < 9; j++ {
			if arr[i][j] != global.Space {
				arr[9-i][j] = arr[i][j] - 10
			} else {
				arr[9-i][j] = global.Space
			}
		}
	}
	ca.Array = arr
	ca.Winner = make(chan int)
	fmt.Println("初始化棋盘：")
	ca.PrintChessboard()
	return ca
}

// IsRed 是否是红棋
func (ca *ChessboardAbscissa) IsRed(x, y int) bool {
	return ca.Array[x][y] <= global.RedBing && ca.Array[x][y] >= global.RedChe
}

//IsBlack 是否是黑棋
func (ca *ChessboardAbscissa) IsBlack(x, y int) bool {
	return ca.Array[x][y] <= global.BlackBing && ca.Array[x][y] >= global.BlackChe
}

//IsSpace 是否是空位无子
func (ca *ChessboardAbscissa) IsSpace(x, y int) bool {
	return ca.Array[x][y] == global.Space
}

//IsRedShuai 是否是红方帅
func (ca *ChessboardAbscissa) IsRedShuai(x, y int) bool {
	return ca.Array[x][y] == global.RedShuai
}

//IsBlackShuai 是不是黑方帅
func (ca *ChessboardAbscissa) IsBlackShuai(x, y int) bool {
	return ca.Array[x][y] == global.BlackShuai
}

//ContainBlackShuai 某一列是否含有黑帅
func (ca *ChessboardAbscissa) ContainBlackShuai(y int) (toX int, toY int, res bool) {
	if y < 3 || y > 5 {
		return -1, -1, false
	}
	for i := 0; i < 3; i++ {
		if ca.Array[i][y] == global.BlackShuai {
			return i, y, true
		}
	}
	return -1, -1, false
}

//ContainRedShuai 某一列是否含有红帅
func (ca *ChessboardAbscissa) ContainRedShuai(y int) (toX int, toY int, res bool) {
	if y < 3 || y > 5 {
		return -1, -1, false
	}
	for i := 7; i < 10; i++ {
		if ca.Array[i][y] == global.RedShuai {
			return i, y, true
		}
	}
	return -1, -1, false
}

//PrintChessboard 打印棋盘地图:打印至控制台
func (ca *ChessboardAbscissa) PrintChessboard() {
	for i := 0; i < len(ca.Array); i++ {
		for j := 0; j < len(ca.Array[0]); j++ {
			fmt.Print(ca.ToString(ca.Array[i][j]) + "	")
		}
		fmt.Println()
	}
}

//ToString 将棋子编号转化为棋子名字
func (ca *ChessboardAbscissa) ToString(whichChess int) string {
	switch whichChess {
	case global.RedShuai:
		return "红帅"
	case global.RedSHi:
		return "红士"
	case global.RedXiang:
		return "红相"
	case global.RedMa:
		return "红马"
	case global.RedChe:
		return "红车"
	case global.RedPao:
		return "红炮"
	case global.RedBing:
		return "红兵"

	case global.BlackShuai:
		return "黑帅"
	case global.BlackSHi:
		return "黑士"
	case global.BlackXiang:
		return "黑相"
	case global.BlackMa:
		return "黑马"
	case global.BlackChe:
		return "黑车"
	case global.BlackPao:
		return "黑炮"
	case global.BlackBing:
		return "黑兵"
	case global.Space:
		return " "
	default:
		return "错误的棋子"
	}
}

//DropChess 吃子覆盖，不吃子填空位
func (ca *ChessboardAbscissa) DropChess(fromX, fromY, toX, toY int) {
	ca.Regret = ca.Array[toX][toY] //暂存棋子

	ca.Array[toX][toY] = ca.Array[fromX][fromY]
	ca.Array[fromX][fromY] = global.Space

	if ca.IsBlackShuai(toX, toY) { //如果吃掉了黑帅，红方赢了
		ca.Winner <- global.RedWinner
	} else if ca.IsRedShuai(toX, toY) { //如果吃掉了红帅，黑方赢了
		ca.Winner <- global.BlackWinner
	}
}

//RegretChess 悔棋
func (ca *ChessboardAbscissa) RegretChess(fromX, fromY, toX, toY int) {
	ca.Array[fromX][fromY] = ca.Array[toX][toY]
	ca.Array[toX][toY] = ca.Regret
}

//ChessRule 检查传入棋子的走法是否合理(包括吃子与不吃子的情况)——注意吃了子一定返回true
func (ca *ChessboardAbscissa) ChessRule(whichChess, fromY, fromX, toY, toX int) (res bool) {
	//横坐标绝对值之差
	xDifference := int(math.Abs(float64(toX - fromX)))
	//纵坐标绝对值之差
	yDifference := int(math.Abs(float64(toY - fromY)))
	switch whichChess {

	case global.BlackShuai: //todo 将帅不能在同一列
		if ca.IsBlack(toX, toY) { //目标位置为自己的棋子不能走
			return false
		} else { //吃子或占空位
			//将帅同列,且中间无子，就会被吃掉
			if toX1, toY1, res := ca.ContainRedShuai(toY); res {
				if toY == toY1 { //扫描一列
					start := 0
					if toX > toX1 {
						start = toX1
					} else {
						start = toX
					}
					for i := start + 1; i <= xDifference; i++ { //扫描中间是否有棋子
						if !ca.IsSpace(i, toY) {
							return false
						}
					}
					//如果将帅同位，就可以吃子:注意吃子顺序
					ca.DropChess(toX1, toY1, toX, toY)
					return true
				}
			}

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

	case global.RedShuai:
		if ca.IsRed(toX, toY) { //目标位置是红棋不可以走
			return false
		} else { //目标位置为黑棋或者空位，可以走，坐标都可以更新

			//将帅同列,且中间无子，就会被吃掉
			if toX1, toY1, res := ca.ContainBlackShuai(toY); res {
				if toY == toY1 { //扫描一列
					start := 0
					if toX > toX1 {
						start = toX1
					} else {
						start = toX
					}
					for i := start + 1; i <= xDifference; i++ { //扫描中间是否有棋子
						if !ca.IsSpace(i, toY) {
							return false
						}
					}
					//如果将帅同位，就可以吃子:注意吃子顺序
					ca.DropChess(toX1, toY1, toX, toY)
					return true
				}
			}
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
	case global.BlackSHi:
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

	case global.RedSHi:
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
	case global.BlackXiang:
		if ca.IsBlack(toX, toY) {
			return false
		} else {
			if toY > 4 { //过河了
				return false
			} else if yDifference == 2 && xDifference == 2 {
				//走"田"字
				centerX := (toX + fromX) / 2
				centerY := (toY + fromY) / 2
				if !ca.IsSpace(centerX, centerY) { // 象眼处有棋子
					return false
				}
				//走棋
				ca.DropChess(fromX, fromY, toX, toY)
				return true
			} else {
				return false
			}
		}
	case global.RedXiang:
		if ca.IsRed(toX, toY) {
			return false
		} else {
			if toY < 5 { //过河
				return false
			} else if yDifference == 2 && xDifference == 2 {
				//走"田"字
				centerX := (toX + fromX) / 2
				centerY := (toY + fromY) / 2
				if !ca.IsSpace(centerX, centerY) { // 象眼处有棋子
					return false
				}
				ca.DropChess(fromX, fromY, toX, toY)
				return true
			} else {
				return false
			}
		}

	case global.BlackMa:
		if ca.IsBlack(toX, toY) {
			return false
		} else {
			if yDifference == 2 && xDifference == 1 {
				if toY-fromY == 2 {
					if ca.IsSpace(fromX, toY-1) { //马蹄处没有棋子
						ca.DropChess(fromX, fromY, toX, toY)
						return true
					}
				} else if toY-fromY == -2 {
					if ca.IsSpace(fromX, toY-1) { //马蹄处没有棋子
						ca.DropChess(fromX, fromY, toX, toY)
						return true
					}
				}
				return false
			} else if yDifference == 1 && xDifference == 2 {
				if toX-fromX == 2 {
					if ca.IsSpace(toX-1, fromY) { //马蹄处没有棋子
						ca.DropChess(fromX, fromY, toX, toY)
						return true
					}
				} else if toX-fromX == -2 {
					if ca.IsSpace(toX-1, fromY) { //马蹄处没有棋子
						ca.DropChess(fromX, fromY, toX, toY)
						return true
					}
				}
				return false
			} else {
				return false
			}
		}

	case global.RedMa:
		if ca.IsRed(toX, toY) {
			return false
		} else {
			if yDifference == 2 && xDifference == 1 {
				if toY-fromY == 2 {
					if ca.IsSpace(fromX, toY-1) { //马蹄处没有棋子
						ca.DropChess(fromX, fromY, toX, toY)
						return true
					}
				} else if toY-fromY == -2 {
					if ca.IsSpace(fromX, fromY-1) { //马蹄处没有棋子
						ca.DropChess(fromX, fromY, toX, toY)
						return true
					}
				}
				return false
			} else if yDifference == 1 && xDifference == 2 {
				if toX-fromX == 2 {
					if ca.IsSpace(toX-1, fromY) { //马蹄处没有棋子
						ca.DropChess(fromX, fromY, toX, toY)
						return true
					}
				} else if toX-fromX == -2 {
					if ca.IsSpace(fromX-1, fromY) { //马蹄处没有棋子
						ca.DropChess(fromX, fromY, toX, toY)
						return true
					}
				}
				return false
			}
			return false
		}
	case global.BlackChe:
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
					for i := start + 1; i <= xDifference; i++ { //扫描中间是否有棋子
						if !ca.IsSpace(i, toY) {
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
					for i := start + 1; i <= yDifference; i++ { //扫描中间是否有棋子
						if !ca.IsSpace(i, toY) {
							return false
						}
					}
					ca.DropChess(fromX, fromY, toX, toY)
					return true
				}
			}
		}
	case global.RedChe:
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
					for i := start + 1; i <= xDifference; i++ { //扫描中间是否有棋子
						if !ca.IsSpace(i, toY) {
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
					for i := start + 1; i <= yDifference; i++ { //扫描中间是否有棋子
						if !ca.IsSpace(i, toY) {
							return false
						}
					}
					ca.DropChess(fromX, fromY, toX, toY)
					return true
				}
			}
		}

	case global.BlackPao:
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
					for i := start + 1; i <= xDifference; i++ { //扫描中间是否有棋子
						if !ca.IsSpace(i, toY) {
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
					for i := start + 1; i <= yDifference; i++ { //扫描中间是否有棋子
						if !ca.IsSpace(i, toY) {
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
					for i := start + 1; i <= xDifference; i++ { //扫描中间是否有棋子
						if !ca.IsSpace(i, toY) {
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
					for i := start + 1; i <= yDifference; i++ { //扫描中间是否有棋子
						if !ca.IsSpace(i, toY) {
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
	case global.RedPao:
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
					for i := start + 1; i <= xDifference; i++ { //扫描中间是否有棋子
						if !ca.IsSpace(i, toY) {
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
					for i := start + 1; i <= yDifference; i++ { //扫描中间是否有棋子
						if !ca.IsSpace(i, toY) {
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
					for i := start + 1; i <= xDifference; i++ { //扫描中间是否有棋子
						if !ca.IsSpace(i, toY) {
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
					for i := start + 1; i <= yDifference; i++ { //扫描中间是否有棋子
						if !ca.IsSpace(i, toY) {
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
	case global.BlackBing:
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
	case global.RedBing:
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
		tool.SugaredError("无效的棋子")
		return false
	}
	return false
}
