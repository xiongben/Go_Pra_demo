package utils

import (
	"awesomeProject1/chatProject/common/message"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

func WritePkg(conn net.Conn, data []byte) (err error) {
	var pkglen uint32
	pkglen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkglen)
	_, err = conn.Write(buf[0:4])
	if err != nil {
		fmt.Println("conn.write byte err = ", err)
		return
	}

	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.write byte(data) err = ", err)
		return
	}
	return
}

func Readpkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据")
	_, err = conn.Read(buf[:4])
	if err != nil {
		//fmt.Println("conn read err = ", err)
		//err = errors.New("read pkg header error")
		return
	}
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])
	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		//fmt.Println("conn read fail,err = ",err)
		err = errors.New("read pkg body error")
		return
	}
	//反序列化
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.unmarshal err = ", err)
		return
	}
	return
}
