package server

import (
	"encoding/binary"
	"net"

	"github.com/Dreamacro/clash/log"
	"github.com/kr328/cfa/tun"
	"golang.org/x/sys/unix"
)

const (
	tunCommandEnd = 0x243
)

func handleTunStart(client *net.UnixConn) {
	buffer := make([]byte, unix.CmsgLen(4*1))

	_, noob, _, _, err := client.ReadMsgUnix(nil, buffer)
	if err != nil {
		log.Errorln("Read tun socket failure, %s", err.Error())
		return
	}

	msg, err := unix.ParseSocketControlMessage(buffer[:noob])
	if err != nil || len(msg) != 1 {
		log.Errorln("Parse tun socket failure, %s", err.Error())
		return
	}

	fds, err := unix.ParseUnixRights(&msg[0])
	if err != nil {
		log.Errorln("Parse tun socket failure, %s", err.Error())
		return
	}

	var mtu uint32
	var dns uint32
	var end uint32

	binary.Read(client, binary.BigEndian, &mtu)
	binary.Read(client, binary.BigEndian, &dns)
	binary.Read(client, binary.BigEndian, &end)

	if end != tunCommandEnd {
		log.Errorln("Invalid tun command end")
		return
	}

	if dns != 0 {
		tun.SetDnsHijacking(true)
	} else {
		tun.SetDnsHijacking(false)
	}

	if err := tun.StartTunProxy(fds[0], int(mtu)); err != nil {
		log.Errorln("Open tun device failure" + err.Error())
	}
}

func handleTunStop(client *net.UnixConn) {
	tun.StopTunProxy()
}
