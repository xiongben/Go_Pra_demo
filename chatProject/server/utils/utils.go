package utils

import (
	"awesomeProject1/chatProject/common/message"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	var pkglen uint32
	pkglen = uint32(len(data))
	//var buf [4]byte
	binary.BigEndian.PutUint32(this.Buf[0:4], pkglen)
	_, err = this.Conn.Write(this.Buf[0:4])
	if err != nil {
		fmt.Println("conn.write byte err = ", err)
		return
	}

	_, err = this.Conn.Write(data)
	if err != nil {
		fmt.Println("conn.write byte(data) err = ", err)
		return
	}
	return
}

func (this *Transfer) Readpkg() (mes message.Message, err error) {
	//buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据")
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		//fmt.Println("conn read err = ", err)
		//err = errors.New("read pkg header error")
		return
	}
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		//fmt.Println("conn read fail,err = ",err)
		err = errors.New("read pkg body error")
		return
	}
	//反序列化
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.unmarshal err = ", err)
		return
	}
	return
}
