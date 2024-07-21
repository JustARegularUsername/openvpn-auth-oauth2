package utils

import (
	"github.com/JustARegularUsername/openvpn-auth-oauth2/internal/config"
)

func TransformCommonName(mode config.OpenVPNCommonNameMode, commonName string) string {
	switch mode {
	case config.CommonNameModePlain:
		return commonName
	case config.CommonNameModeOmit:
		fallthrough
	default:
		return config.CommonNameModeOmitValue
	}
}
