package main

import (
    "fmt"
    "net"
    "os"
    "time"
    "math/rand"
)

type Pos_t struct {
    x int
    y int
    flag int
    score float64
}

//1-49 white | 2-50 black
var color int
var otherColor int

func main() {
    fmt.Println("start...")
    fmt.Println("username", os.Args[1])
    
    username := os.Args[1]
    port := os.Args[2]

    server := "localhost:" + port
    tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
    
    if err != nil {
        fmt.Println("fatal error:", err.Error())
    }

    conn, err := net.DialTCP("tcp", nil, tcpAddr)

    if err != nil {
        fmt.Println("fatal error", err.Error())
    }

    name := "N" + username
    conn.Write([]byte(name))
    nameReponse := make([]byte, 10)
    n, err := conn.Read(nameReponse)
    if err != nil {
        fmt.Println("register failed")
    }

    tmpNameResponse := nameReponse[1:2]
    if tmpNameResponse[0] == 49 {
        fmt.Println("black name:\n", string(nameReponse[:n]))
        color = 2
        otherColor = 1        
    } else {
        fmt.Println("white name:\n", string(nameReponse[:n]))
        color = 1
        otherColor = 2
    }

    fmt.Println("color: ", color)

    state := make([]byte, 65)

    for {
        n, err := conn.Read(state) 
                
        if err != nil {
            fmt.Println("get state failed")
        }

        if n == 1 {
            continue
        }

        if n != 65 {
            fmt.Println(state[:n])
            tmpState := state[:n]
            if (tmpState[0] == 71) {
                break
            } else {
                continue
            }
        }
       
        /* 解析当前棋盘 */
        var posList [8][8]Pos_t
        posList = analysis(state)

        /* 获取可落子位位置*/
        rightPosList := getRightPos(posList)

        /* 计算可落子位置得分，获取最佳位置 */
        top:= topPos(rightPosList)
        fmt.Println(top)

        /* 发送落子位置至服务器 */
        x := byte(top.x + 48) 
        y := byte(top.y + 48)
        tmpSend := []byte{77, x, y}
        conn.Write(tmpSend)
    }
}


func analysis(state []byte) [8][8]Pos_t {
    var posList  [8][8]Pos_t;
    for i := 1; i < 64; i += 8{
        start := i
        for j := 0; j < 8 ; j++ {
            var tmpPos Pos_t
            tmpPos.x = (i-1) / 8
            tmpPos.y = j
            tmpPos.score = 0
            tmpSlice := state[start+j:start+j+1]
            if tmpSlice[0] == 49 {
                tmpPos.flag = 1
            } else if tmpSlice[0] == 50 {
                tmpPos.flag = 2
            } else {
                tmpPos.flag = 0
            }
            posList[start/8][j] = tmpPos
        }
    }
    fmt.Println("=== 得到棋盘  ===", posList)
    return posList
}

func getRightPos(posList [8][8]Pos_t) []Pos_t {
    var rightPosList []Pos_t 
    var tmpMatchPosList []Pos_t

    /* 找出当前所有自身（w|b）颜色点 */
    for i := 0; i < 8; i++ {
        for j := 0; j < 8; j++ {
            tmpPos := posList[i][j]
            if tmpPos.flag == color {
                tmpMatchPosList = append(tmpMatchPosList, tmpPos)
            }
        }
    }

    fmt.Println("=== 取得我方所有棋子 ===", tmpMatchPosList)
  
    //对该点进行八个方向搜索，每个点判断能够夹住对方子
    //如果能则记录，否则放弃
    for _, tmpPos := range(tmpMatchPosList) {
        x := tmpPos.x
        y := tmpPos.y

        //y x+
        fmt.Println("y x+")
        var meetOtherColor bool = false
        if x < 7 {
            for i := x + 1; i < 8; i++ {
                curPos := posList[i][y]
                right, stop, tmpMeetOtherColor := isRightPos(curPos, meetOtherColor)
                meetOtherColor = tmpMeetOtherColor
                fmt.Println(right, stop, meetOtherColor)
                fmt.Println("y x+ ->x y -> new x y -> meet", x, y, i, y)
                if right {
                    rightPosList = append(rightPosList, curPos)
                    break
                }

                if (stop) {
                    break
                }
            }
        }

        //y x-
        fmt.Println("y x-")
        meetOtherColor = false
        if x > 0 {
            for i := x - 1; i >= 0; i-- {
                curPos := posList[i][y]
                right, stop, tmpMeetOtherColor:= isRightPos(curPos, meetOtherColor)
                meetOtherColor = tmpMeetOtherColor
                fmt.Println(right, stop, meetOtherColor)
                fmt.Println("y x+ ->x y -> new x y -> meet", x, y, i, y)
                if right {
                    rightPosList = append(rightPosList, curPos)
                    break
                }

                if (stop) {
                    break
                }
            }
        }

   
        //x y+
        fmt.Println("x y+")
        meetOtherColor = false
        if y < 7 {
            for i := y + 1; i < 8; i++ {
                curPos := posList[x][i]
                right, stop, tmpMeetOtherColor := isRightPos(curPos, meetOtherColor)
                meetOtherColor = tmpMeetOtherColor
                fmt.Println(right, stop, meetOtherColor)
                fmt.Println("y x+ ->x y -> new x y -> meet", x, y, x, i)
                if right {
                    rightPosList = append(rightPosList, curPos)
                    break
                }

                if (stop) {
                    break
                }
            }
        }

        //x y- 
        fmt.Println("x y-")
        meetOtherColor = false
        if y > 0 {
            for i := y - 1; i >= 0; i-- {
                curPos := posList[x][i]
                right, stop, tmpMeetOtherColor := isRightPos(curPos, meetOtherColor)
                meetOtherColor = tmpMeetOtherColor
                fmt.Println(right, stop, meetOtherColor)
                fmt.Println("y x+ ->x y -> new x y -> meet", x, y, x, i)
                if right {
                    rightPosList = append(rightPosList, curPos)
                    break
                }

                if (stop) {
                    break
                }
            }
        }

        //x- y+ 
        fmt.Println("x- y+")
        meetOtherColor = false
        if x > 0 {
            var j int = 1
            for i := x - 1; i >= 0; i-- {
                if (y + j) > 7 {
                    break
                }
                curPos := posList[i][y+j]
                right, stop, tmpMeetOtherColor := isRightPos(curPos, meetOtherColor)
                meetOtherColor = tmpMeetOtherColor
                fmt.Println(right, stop, meetOtherColor)
                fmt.Println("y x+ ->x y -> new x y -> meet", x, y, i, y+j)

                j += 1

                if right {
                    rightPosList = append(rightPosList, curPos)
                    break
                }

                if (stop) {
                    break
                }
            }
        }


        //x- y- 
        fmt.Println("x- y-")
        meetOtherColor = false
        if x > 0 {
            var j int = 1
            for i := x - 1; i >= 0; i-- {
                if (y - j) < 0 {
                    break
                }
                curPos := posList[i][y-j]
                right, stop, tmpMeetOtherColor := isRightPos(curPos, meetOtherColor)
                meetOtherColor = tmpMeetOtherColor
                fmt.Println(right, stop, meetOtherColor)
                fmt.Println("y x+ ->x y -> new x y -> meet", x, y, i, y-j)

                j += 1

                if right {
                    rightPosList = append(rightPosList, curPos)
                    break
                }

                if (stop) {
                    break
                }
            }
        }

        //x+ y+ 
        fmt.Println("x+ y+")
        meetOtherColor = false
        if x > 0 {
            for i := 1; i < 7; i++ {
                if (y + i) > 7  || (x + i) > 7{
                    break
                }
                curPos := posList[x+i][y+i]
                right, stop, tmpMeetOtherColor := isRightPos(curPos, meetOtherColor)
                meetOtherColor = tmpMeetOtherColor
                fmt.Println(right, stop, meetOtherColor)
                fmt.Println("y x+ ->x y -> new x y -> meet", x, y, x+i, y+i)

                if right {
                    rightPosList = append(rightPosList, curPos)
                    break
                }

                if (stop) {
                    break
                }
            }
        }

        //x+ y- 
        fmt.Println("x+ y-")
        meetOtherColor = false
        if x > 0 {
            var j int = 1
            for i := y - 1; i >= 0; i-- {
                if (x + j) > 7 {
                    break
                }
                curPos := posList[x+j][i]
                right, stop, tmpMeetOtherColor := isRightPos(curPos, meetOtherColor)
                meetOtherColor = tmpMeetOtherColor
                fmt.Println(right, stop, meetOtherColor)
                fmt.Println("y x+ ->x y -> new x y -> meet", x, y, x+j, i)
                j = j + 1
                if right {
                    rightPosList = append(rightPosList, curPos)
                    break
                }

                if (stop) {
                    break
                }
            }
        }

    }

    fmt.Println("=== 取得可落子位置 ===", rightPosList)
    return rightPosList
}

func topPos(posList []Pos_t) Pos_t{
    var topPos  Pos_t

    var rankList []Pos_t
    for _, tmpPos := range(posList) {
        /* 四角 */
        if (tmpPos.x == 0 && tmpPos.y == 0 ) || (tmpPos.x == 0 && tmpPos.y == 7 ) || (tmpPos.x == 7 && tmpPos.y == 0 ) || (tmpPos.x == 7 && tmpPos.y == 7) {
            return tmpPos
        }

        tmpPos := calPos(tmpPos)
        rankList = append(rankList, tmpPos)
    }
    
    var meet int = 0
    for _, tmpPos := range(rankList) {
        if tmpPos.score == 1 {
            topPos = tmpPos
            meet = 1
        }
    } 

    if meet == 0 {
        len := len(posList)
        r := rand.New(rand.NewSource(time.Now().UnixNano()))
        key := r.Intn(len)
        topPos = posList[key]
    }
    return topPos 
}

func calPos(pos Pos_t) Pos_t{
    posResult :=pos
    if (posResult.x ==1 && posResult.y == 1) || (posResult.x ==1 && posResult.y == 0) || (posResult.x ==0 && posResult.y == 1) || (posResult.x == 0&& posResult.y == 6) || (posResult.x == 1 && posResult.y == 6) || (posResult.x ==1 && posResult.y == 7) || (posResult.x == 6 && posResult.y == 0) || (posResult.x == 6 && posResult.y == 1)  || (posResult.x ==7 && posResult.y == 1) || (posResult.x ==6 && posResult.y == 6) || (posResult.x ==7 && posResult.y == 6) || (posResult.x ==6 && posResult.y == 7) {
        posResult.score = 0
    } else {
        posResult.score = 1
    }

    return posResult
}

func isRightPos(curPos Pos_t, meetOtherColor bool) (right bool, stop bool, returnMeetOtherColor bool)  {
    fmt.Println("in meet", meetOtherColor) 
    right = false
    stop = false
    returnMeetOtherColor = false
    if curPos.flag == 0 { //遇见空位直接停止
        stop = true
        if meetOtherColor  { //遇见空位前遇见了其他颜色，改点可以记录
           right = true
        }
    } else if curPos.flag == color { //遇见自身颜色直接停止
        right = false
        stop = true
    } else if curPos.flag == otherColor {
        returnMeetOtherColor  = true
    }
    return 
}
