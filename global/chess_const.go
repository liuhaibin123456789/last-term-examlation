// Package global
// @Author cold bin
// @DATE 2022/7/16
package global

const (
	RedWinner   = 0 //红方赢了
	BlackWinner = 1 //黑方赢了
)

//棋子枚举常量
const (
	Space = iota //没有棋子占领

	RedChe   //红车
	RedMa    //红马
	RedXiang //红象
	RedSHi   //红士
	RedShuai //红帅
	RedPao   //红炮
	RedBing  //红兵

	BlackChe   = 3 + iota //黑车
	BlackMa               //黑马
	BlackXiang            //黑象
	BlackSHi              //黑士
	BlackShuai            //黑帅
	BlackPao              //黑炮
	BlackBing             //黑兵
)
