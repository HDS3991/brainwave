package special

import (
	"regexp"
	"strings"
)

const DictionaryPidThreatTech = "未知威胁\n端口扫描\n漏洞扫描\n信息刺探\n弱密码尝试\n暴力破解\n路径遍历\n文件下载\n文件执行\n系统命令注入\nSQL注入\nXSS攻击\nXXE攻击\nLDAP注入\nOGNL注入\n反序列化\n释放重用\nCSRF攻击\nSSRF攻击\n缓冲区溢出\n未授权访问\nUDP泛洪\nSYN泛洪\nACK泛洪\nCC攻击\n鱼叉攻击\n网络钓鱼\n网页挂马\n中间人攻击\nDNS劫持\nTCP劫持\n木马\n病毒\n间谍软件\n后门程序\nWEBSHELL\n恶意工具\n命令控制\n系统控制\n心跳连接\nDGA域名\n信息回传\n传播扩散\n其他威胁\n未知网络扫描\n其他网络扫描\n未知口令破解\n其他口令破解\n未知漏洞攻击\n文件上传\n其他漏洞攻击\n未知网络风暴\n其他网络风暴\n未知钓鱼攻击\n其他网络钓鱼\n未知水坑攻击\n其他水坑攻击\n未知网络劫持\n其他网络劫持\n未知恶意文件\n其他恶意文件\n未知远程控制\n其他远程控制\n未知任务执行\n武器下载\n环境刺探\n其他任务执行\n下载器\n受感染文件\n未知威胁"
const DictionaryPidDiscoverMechan = "未知\n黑IP\n黑域名\n黑URL\n黑邮箱\n文件哈希\t\n数据哈希\n负载特征\n文件特征\n行为基线\n文件行为\n行为模型\n其他\n"
const VictimAssetsType = "未知\n服务器/云\nPC\n手机\n工控设备\nIOT设备\n网络设备\n其他"
const killchain = "不适用\n侦查跟踪\n武器化\n部署\n漏洞利用\n安装植入\n命令控制\n目标达成\n未知"
const strage = "未知\nTA0001-初始访问\nTA0002-执行\nTA0003-权限维持\nTA0004-权限提升\nTA0005-防御绕过\nTA0006-凭据访问\nTA0007-发现\nTA0008-横向移动\nTA0009-收集\nTA0010-数据渗出\nTA0011-命令与控制\nTA00040-影响\nTA0042-资源开发\nTA0043-侦察\n"
const techIds = "851\n852\n853\n854\n855\n856\n857\n858\n859\n860\n861\n862\n863\n864\n865"
const source = "Snort 2.X\nSnort 3X\nSuricata ET\nSuricata ETPro \nClamAV\nYara\n其他\n"

func ConvertLineBreaksToCommas(str string) string {
	return strings.ReplaceAll(str, "\n", ",")
}

func ReplaceSuricataSId(rule string, sid string) string {
	pattern := "[ |;]sid:(.*?);"
	re, err := regexp.Compile(pattern)
	if err != nil {
		return rule
	}
	var sidBuilder strings.Builder
	sidBuilder.WriteString(sid)
	sidBuilder.WriteString(";")

	matches := re.FindAllString(rule, -1)
	if len(matches) > 0 {
		newStr := strings.Replace(rule, matches[0], sidBuilder.String(), -1)
		return newStr
	}

	return rule
}
