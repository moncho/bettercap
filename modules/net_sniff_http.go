package modules

import (
	"bufio"
	"bytes"
	"net/http"

	"github.com/bettercap/bettercap/core"

	"github.com/bettercap/gopacket"
	"github.com/bettercap/gopacket/layers"
)

func httpParser(ip *layers.IPv4, pkt gopacket.Packet, tcp *layers.TCP) bool {
	data := tcp.Payload
	reader := bufio.NewReader(bytes.NewReader(data))
	req, err := http.ReadRequest(reader)

	if err == nil {
		NewSnifferEvent(
			pkt.Metadata().Timestamp,
			"http",
			ip.SrcIP.String(),
			req.Host,
			req,
			"%s %s %s %s %s",
			core.W(core.BG_RED+core.FG_BLACK, "http"),
			vIP(ip.SrcIP),
			core.W(core.BG_LBLUE+core.FG_BLACK, req.Method),
			vURL(req.URL.String()),
			core.Dim(req.UserAgent()),
		).Push()

		return true
	}

	return false
}
