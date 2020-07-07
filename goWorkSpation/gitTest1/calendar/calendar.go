package main

import (
	"bufio"
	"fmt"
	mapset "github.com/deckarep/golang-set"
	"github.com/gomodule/redigo/redis"
	"io"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

type MovieList struct {
	hotMovieList []int `json:"data"`
}
type AllowedShow struct {
	allowedList []int
}

var m atomic.Value

func getHotDataList() []int {

	for {
		time.Sleep(time.Second * 1)
		m1 := m.Load().(*MovieList)
		fmt.Println("用户读取数据中..", m1.hotMovieList)
		return m1.hotMovieList
	}

}
func getData() []int {

	m1 := m.Load().(*MovieList)

	fmt.Println("初始化之后取的数据", m1.hotMovieList)

	return m1.hotMovieList
}
func generateHotData(filename string, conn redis.Conn) {
	timeUnxi := time.Now().Unix()

	//获取到元热点数据

	srcHotData := getSrcHotData(filename)

	fmt.Println("元数据： ", srcHotData)
	//进行白名单操作,筛选数据
	filterWhiteList(srcHotData)
	fmt.Println("白名单过滤之后元数据 ", srcHotData, len(srcHotData))
	//判断元热点数据是否符合规范,不符合我们就不覆盖热点数据
	//如果第一次跑数据就不符合规范，这个需要初始化的时候就添加元热带上数据
	//如果第一次跑数据就不符合规范，这个需要初始的时候就添加元热带上数据

	if len(srcHotData) > 3 {
		//符合规范，覆盖前一天的热点数据
		fmt.Println("覆盖数据")
		conn.Do("del", "srcHotData")
		for i, k := range srcHotData {

			conn.Do("hset", "srcHotData", i, k)
		}

	}
	//如果元热点数据一个没有，那么我们需要查询前一天的数据作为支撑
	if len(srcHotData) == 0 {

		src, err := redis.StringMap(conn.Do("hgetAll", "srcHotData"))
		if err != nil {

		}
		//把map中的所有string转换为int
		for i, k := range src {
			index, err := strconv.Atoi(i)
			if err != nil {

			}
			//value是把string类型的数组进行
			val := stringToArr(k)
			srcHotData[index] = val

		}
		fmt.Println("元数据为空，查询reids数据库前一天数据", srcHotData)
	}

	//查询出redis中的所有数据，然后进行去重操作
	hotmap, err := redis.Strings(conn.Do("zrevrange", "calendar", "0", "-1"))

	if err != nil {

	}
	//迭代热点数据map，存入set集合中
	//获取近5天的热点数据作为多样性比较
	lateyData := make(map[int]int, 5)
	otherHotData := mapset.NewSet()
	currentList := make([]int, 0)
	for i, k := range hotmap {
		str := strings.Split(k, ":")
		key, _ := strconv.Atoi(str[0])
		val, _ := strconv.Atoi(str[1])
		if i < 5 {
			//保存近五天的热点数据，作为多样性参考
			lateyData[key] = val
		}
		//然后把查询的媒体id数据存入set中
		otherHotData.Add(val)
		//存入列表中
		currentList = append(currentList, val)
	}
	//去重和多样性操作,得到当天生成的hotdata
	hotdata, mediaId := filterRepeatData(otherHotData, lateyData, srcHotData)
	fmt.Println("今日热点数据： ", hotdata)
	//存入redis
	fmt.Println("今日存储   ====", timeUnxi, hotdata)
	conn.Do("zadd", "calendar", timeUnxi, hotdata)

	//刷新数据
	//取出前23的数据，然后把今天生成的数据加到最前面
	fmt.Println("redis中所有近期热点数据  ", currentList, "长度为：", len(currentList))
	fmt.Println("截取23为热点数据", currentList[0:23], "长度为：", len(currentList[0:23]))
	reverArr(currentList[0:23])
	current := append(currentList[0:23], mediaId)
	reverArr(current)
	fmt.Println("修改之后的热点数据", current, "长度为：", len(current))
	updataHotData(current)

}

func updataHotData(currentList []int) {

	m1 := m.Load().(*MovieList)
	fmt.Println("更改前： ", m1.hotMovieList)
	m2 := &MovieList{}
	m2.hotMovieList = currentList
	fmt.Println("更改后： ", m2.hotMovieList)

	m.Store(m2)

}

//去重操作
func filterRepeatData(set mapset.Set, latelyData map[int]int, data map[int][]int) (string, int) {
	var res string
	var remedy string
	var remedyId int
	var mediaId int
outerr:
	for i, k := range data {
		//多样性筛选
		if dataDiversify(latelyData, i) {

			//选出当天热门
			for _, m := range k {

				//判断当天热门是否重复,无重复返回res
				if set.Add(m) {

					res = strconv.Itoa(i) + ":" + strconv.Itoa(m)
					mediaId = m
					break outerr
				}
			}

		}
		//当多样性和去重失败时的补救措施，填充数据
		remedy = strconv.Itoa(i) + ":" + strconv.Itoa(k[0])
		remedyId = k[0]
	}

	//操作问题，如果去重操作和多样性操作把所有数据全部排除掉，必须实现补救方法
	if res == "" {

		return remedy, remedyId
	}
	return res, mediaId
}

//数据多样化操作，
func dataDiversify(latelyData map[int]int, data int) bool {
	//判断进5天的电影类型是否重复，重复返回false，否则为true
	for i, _ := range latelyData {

		if i == data {
			return false
		}
	}
	return true
}

func filterWhiteList(data map[int][]int) {
	allowedShow := AllowedShow{}
	//白名单列表
	allowedShow.allowedList = []int{36, 37, 38, 44, 46, 50}

	//白名单存入set集合好做筛选
	as := mapset.NewSet()
	for _, k := range allowedShow.allowedList {
		as.Add(k)
	}

	for i, _ := range data {
		//如果不在白名单中，set就是true，那么我们需要删除该分类
		if !as.Add(i) {
			delete(data, i)
		}
	}

}

//读取文件内容生成热点数据
func getSrcHotData(filename string) map[int][]int {

	file, err := os.Open(filename)

	if err != nil {

	}
	defer file.Close()

	buf := bufio.NewReader(file)

	hotDataSrc := make(map[int][]int, 0)
	for {
		line, err := buf.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		//把热点元数据存入map集合中
		list := make([]int, 0)
		//读取文件中的字符串会有\n换行作为最后一个结尾 ，所以必须把这个排除掉
		s := strings.Split(line, "\\t")

		index, err := strconv.Atoi(s[0])
		if err != nil {

		}
		val := s[1]
		//使用len(val) - 1来去除掉/n换行符
		for _, k := range strings.Split(val[0:len(val)-1], ":") {
			m, err := strconv.Atoi(k)
			if err != nil {

			}
			list = append(list, m)
		}
		fmt.Println("文件内容", index, list)
		hotDataSrc[index] = list
		fmt.Println("读取文件操作的内容", hotDataSrc)
	}
	return hotDataSrc
}
