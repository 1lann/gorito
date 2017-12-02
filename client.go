package main

import "strings"

type Client struct {
	BR   *Region
	EUNE *Region
	EUW  *Region
	JP   *Region
	KR   *Region
	LAN  *Region
	LAS  *Region
	NA   *Region
	OCE  *Region
	TR   *Region
	RU   *Region
	PBE  *Region
}

func main() {

}

func (c *Client) Region(region string) *Region {
	switch strings.ToUpper(region) {
	case "BR":
		return c.BR
	case "EUNE":
		return c.EUNE
	case "EUW":
		return c.EUW
	case "JP":
		return c.JP
	case "KR":
		return c.KR
	case "LAN":
		return c.LAN
	case "LAS":
		return c.LAS
	case "NA":
		return c.NA
	case "OCE":
		return c.OCE
	case "TR":
		return c.TR
	case "RU":
		return c.RU
	case "PBE":
		return c.PBE
	default:
		return nil
	}
}

func (c *Client) Platform(platform string) *Region {

}
