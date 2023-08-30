package special

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"os"
	"testing"
)

func TestConvertLineBreaksToCommas(t *testing.T) {
	file, err := os.Create("./example.txt")
	if err != nil {
		fmt.Println("无法创建文件:", err)
		return
	}
	defer file.Close()

	str := ConvertLineBreaksToCommas(source)
	data := []byte(str)

	// 将数据写入文件
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("无法写入文件:", err)
		return
	}

}

func TestReplaceSuricataSId(t *testing.T) {
	suricataRule := "alert http $HOME_NET any -> $EXTERNAL_NET any (msg:\"TROJAN Win32/Wacapew.C!ml CnC Checkin\"; flow:established,to_server; content:\"GET\"; http_method; content:\"/?cname=\"; http_uri; depth:8; fast_pattern; content:\"Go-http-client\"; http_user_agent; depth:14; http_header_names; content:\"|0d 0a|Host|0d 0a|User-Agent|0d 0a|Accept-Encoding|0d 0a 0d 0a|\"; reference:md5,7f939b3ba32224827485168077881948; classtype:trojan-activity; sid:2037777; rev:1; metadata:affected_product Windows_XP_Vista_7_8_10_Server_32_64_Bit, attack_target Client_Endpoint, created_at 2022_07_15, deployment Perimeter, former_category MALWARE, performance_impact Low, signature_severity Major, updated_at 2022_07_15;)"
	wantSchema := "alert http $HOME_NET any -> $EXTERNAL_NET any (msg:\"TROJAN Win32/Wacapew.C!ml CnC Checkin\"; flow:established,to_server; content:\"GET\"; http_method; content:\"/?cname=\"; http_uri; depth:8; fast_pattern; content:\"Go-http-client\"; http_user_agent; depth:14; http_header_names; content:\"|0d 0a|Host|0d 0a|User-Agent|0d 0a|Accept-Encoding|0d 0a 0d 0a|\"; reference:md5,7f939b3ba32224827485168077881948; classtype:trojan-activity;%s; rev:1; metadata:affected_product Windows_XP_Vista_7_8_10_Server_32_64_Bit, attack_target Client_Endpoint, created_at 2022_07_15, deployment Perimeter, former_category MALWARE, performance_impact Low, signature_severity Major, updated_at 2022_07_15;)"
	//expected := "alert http $HOME_NET any -> $EXTERNAL_NET any (msg:\"TROJAN Win32/Wacapew.C!ml CnC Checkin\"; flow:established,to_server; content:\"GET\"; http_method; content:\"/?cname=\"; http_uri; depth:8; fast_pattern; content:\"Go-http-client\"; http_user_agent; depth:14; http_header_names; content:\"|0d 0a|Host|0d 0a|User-Agent|0d 0a|Accept-Encoding|0d 0a 0d 0a|\"; reference:md5,7f939b3ba32224827485168077881948; classtype:trojan-activity;sid:test rev:1; metadata:affected_product Windows_XP_Vista_7_8_10_Server_32_64_Bit, attack_target Client_Endpoint, created_at 2022_07_15, deployment Perimeter, former_category MALWARE, performance_impact Low, signature_severity Major, updated_at 2022_07_15;)"
	sid := "sid:test"
	want := fmt.Sprintf(wantSchema, sid)
	assert.Equal(t, ReplaceSuricataSId(suricataRule, sid), want)
}
