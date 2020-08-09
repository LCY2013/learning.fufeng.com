package main

import (
	"fmt"
	"os"
)

//广度遍历图

/**
读取原始图数据
*/
func readMazes(path string) [][]int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	//读取行数
	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	//创建一个二维数组结构
	mazes := make([][]int, row)
	for indexRow := range mazes {
		mazes[indexRow] = make([]int, col)
		for indexCol := range mazes[indexRow] {
			fmt.Fscanf(file, "%d", &mazes[indexRow][indexCol])
		}
	}
	return mazes
}

//定义点
type pointer struct {
	i, j int
}

//定义一个点的上左下右
var dirs = [4]pointer{
	{-1,0},{0,-1},{1,0},{0,1},
}

//定义点的计算
func (p *pointer) add(pt pointer) pointer {
	return pointer{
		i: p.i + pt.i,
		j: p.j + pt.j,
	}
}

//定义点是否相等
func (p *pointer) equal(pt pointer) bool {
	return p.i == pt.i && p.j == pt.j
}

func (p *pointer) legal() bool {
	return p.i < 0 || p.j < 0
}

//计算当前点位在途中的位置
func (p *pointer) at(mazes [][] int) (*int,bool) {
	if p.i < 0 || p.i > len(mazes)-1{
		return nil,false
	}
	if p.j < 0 || p.j > len(mazes[p.i])-1{
		return nil,false
	}
	return &mazes[p.i][p.j],true
}

//寻找，广度搜索图
func walking(mazes [][]int, start, end pointer) [][]int{
	//先判断起始点是否一致,点位是否合法
	if start.equal(end) || start.legal() || end.legal() {
		return mazes
	}
	//定义和图一致的行走路线
	steps := make([][]int,len(mazes))
	for stepRow := range steps {
		steps[stepRow] = make([]int,len(mazes[stepRow]))
	}

	//定义图待遍历的点位的队列,slice
	Q := []pointer{start}

	for len(Q) > 0 {
		//获取队列中的点位
		cur := Q[0]
		Q = Q[1:]
		//1、回到原点就退出
		if cur.equal(end){
			break
		}
		for dir := range dirs {
			//计算探索节点按上左下右顺序的周围节点
			next := cur.add(dirs[dir])
			//计算探索点是否合法，和探索点的值
			v, ok := next.at(mazes)
			//2、如果遇到墙就继续
			if !ok || v == nil || *v == 1{
				continue
			}
			//计算探索点在步数中是否合法
			v,ok = next.at(steps)
			//3、如果点位不是0，说明已经走过了
			if !ok || v == nil || *v != 0{
				continue
			}

			//4、如果该点位和起始点一样就跳过
			if next == start{
				continue
			}

			//5、计算在最终路线图中的情况
			v,ok = cur.at(steps)
			steps[next.i][next.j] = *v + 1
			Q = append(Q,next)

			if next == end{
				steps[next.i][next.j] = steps[next.i][next.j] + 1
			}
		}
		steps[cur.i][cur.j] = steps[cur.i][cur.j] + 1
	}

	return steps
}

func main() {
	mazes := readMazes("/Users/magicLuoMacBook/software/go/projects/reptile/learning.fufeng.com/project/reptile/maze/maze.in")
	for mazeRow := range mazes{
		for mazeCol := range mazes[mazeRow]{
			fmt.Printf("%4d",mazes[mazeRow][mazeCol])
		}
		fmt.Println()
	}

	fmt.Println()

	steps := walking(mazes,
		pointer{
			i: 0,
			j: 0,
		}, pointer{
			i: len(mazes) - 1,
			j: len(mazes[0]) - 1,
		})
	for stepRow := range steps{
		for stepCol := range steps[stepRow] {
			fmt.Printf("%4d",steps[stepRow][stepCol])
		}
		fmt.Println()
	}
}
