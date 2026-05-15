# Chapter 07 — Security（网络安全）

> 对应原书 *Network Programming with Go Language, 2nd Ed.* 第 7 章。  
> 安全不是「调几个加密 API」的补丁集合，而是**信任模型 + 分层机制 + 运维生命周期**的系统工程；脱离框架写密码学，就像在流沙上盖楼——小疏忽会拖垮整条信任链。

**性质**：与第 **1** 章（分层与谬误）、第 **5** 章（协议边界）、第 **6** 章（表示与编码）、后续 **HTTP/TLS/gRPC** 强相关。

---

## 7.1 ISO 安全架构（ISO 7498-2 心智）

标准把**安全服务（Security Services）**与**安全机制（Security Mechanisms）**分开：前者是「要达到什么安全属性」，后者是「用什么技术实现」。

**常见安全服务目标**

| 服务 | 含义 |
|------|------|
| **认证（Authentication）** | 确认通信实体身份，防冒充 |
| **访问控制（Access Control）** | 策略下谁能访问何种资源 |
| **保密性（Confidentiality）** | 非授权者不可读 |
| **完整性（Integrity）** | 可检测篡改与意外损坏 |
| **不可否认性（Non-repudiation）** | 参与方难以事后抵赖（常依赖签名与审计） |

**架构视角**：只做强加密而**不做身份认证**，仍可能掉进 **MITM（中间人）**——对方换成攻击者的密钥你并不知道。宏观上缺任一维度，信任链都会断。

---

## 7.2 安全功能与层级（深度防御）

**深度防御**：在栈的多个位置叠安全能力；但别误以为「某一层安全 = 端到端安全」。

| 层级（教学 OSI） | 典型能力 | Go 网络实践提醒 |
|------------------|-----------|------------------|
| L1/L2 | 链路加密、物理隔离 | 机房/VLAN；**不等于**跨公网端到端 |
| L3 | IPsec 等 | 主机到主机隧道；与业务进程 TLS **可叠加** |
| L4 | **TLS** | `crypto/tls`：进程间传输机密性 + 完整性 + 身份（证书） |
| L7 | 鉴权、审计、输入校验 | 与传输层安全**分工**：TLS 不替代业务授权 |

**现实复杂度**：如 **HTTPS** 会跨越「会话 / 表示 / 应用」的教科书边界（与第 1 章分层叙事一致）。微服务里若只信「入口 LB 到后端是内网」，要意识到流量可能在**中间件解密再转发**——策略上是否仍满足合规，要单独论证。

**工程偏好**：**端到端 TLS**（或 mTLS）+ 应用层最小权限；让 **TLS** 处理传输密码学细节，让 **业务代码** 处理授权模型（RBAC/ABAC、令牌范围等）。

---

## 7.3 安全机制（Mechanisms）

服务靠机制落地，常见映射：

- **加密（Encipherment）**：隐藏内容。  
- **数字签名（Digital Signature）**：身份绑定 + 完整性 + 不可否认（视设计与法域）。  
- **流量填充（Traffic padding）**：对抗流量分析（成本与延迟权衡）。  
- **公证 / 可信第三方**：如 **CA** 对公钥与身份的绑定。

**战略评估**：**单点依赖一种机制**往往不够。典型反例：**只加密不做完整性（无 MAC / 非 AEAD）**时，攻击者可能做**密文操控**或配合协议缺陷破坏业务语义。现代对称加密应优先 **AEAD**（如 **AES-GCM**）：**机密性 + 完整性**一体。

---

## 7.4 数据完整性（Data Integrity）

网络既不绝对可靠也不绝对安全（对照八大谬误 #1、#4）。**哈希**能发现内容被改，但**不**提供保密性，也**不**单独解决「是谁改的」——后者常交给 **MAC / 签名 / AEAD**。

**工程警示（解析一致性）**：例如对 **IP / URL** 等输入，不同组件、不同版本对「宽松字面量」的接受度可能不同；若安全策略依赖「字符串相等」而非**规范化后的结构化对象**，易出现 **SSRF / 绕过** 类问题。做法：**白名单**（允许的网络、协议、端口）、**显式规范化**、**禁止仅靠单一解析函数**充当安全边界。

**SHA-256 示例**（小内存一次哈希；大文件用 `hash.Hash` 流式 `Write`）：

```go
import "crypto/sha256"

func checksum256(data []byte) [32]byte {
	return sha256.Sum256(data)
}
```

- **禁用 MD5 / SHA-1** 做安全相关完整性（碰撞风险）。  
- **大对象**：分块读入并写入 `sha256.New()`，尽早发现坏块。

---

## 7.5 对称密钥加密（Symmetric）

**优点**：吞吐高、适合大数据量。  
**痛点**：**密钥分发**——没有预先安全信道时，无法凭空共享对称密钥。

**AES-GCM（AEAD）** 要点：`nonce` **唯一**（通常随机且长度满足 `NonceSize()`），**可明文前缀**随密文发送；**同一 key 下 nonce 复用是灾难性错误**。

```go
import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

func sealAESGCM(key, plaintext, aad []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	return gcm.Seal(nonce, nonce, plaintext, aad), nil
}
```

密钥泄露则**历史密文**在「未实现前向保密 / 未轮换密钥」前提下也可能被解密——需 **KMS、轮换、mTLS** 等运维体系配合。

---

## 7.6 公钥加密与混合加密（Asymmetric）

**PKI 心智**：解决「在不安全信道上如何信一个公钥真的是对方」的问题。

**信箱比喻**：公钥像邮筒口（谁都能投密文）；私钥像钥匙（只有持有者能开）。

**RSA 签名示例（PKCS#1 v1.5）**——演示「私钥签、公钥验」；**新协议**更常推荐 **RSA-PSS** 或 **ECDSA / Ed25519**（按团队标准与合规选型）。

```go
import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

func signVerifyRSA(priv *rsa.PrivateKey, pub *rsa.PublicKey, msg []byte) error {
	h := sha256.Sum256(msg)
	sig, err := rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA256, h[:])
	if err != nil {
		return err
	}
	return rsa.VerifyPKCS1v15(pub, crypto.SHA256, h[:], sig)
}
```

**新手大忌**：**不要**用 RSA 公钥直接加密大量业务数据（慢、包大小受限）。标准路径是 **混合加密**：握手阶段用非对称交换或协商出对称密钥，再用 **AES-GCM** 等保护 bulk data。

**生产提示**：密钥生成昂贵——**进程启动或部署时**加载/KMS 取出，**禁止**每请求 `GenerateKey`。

---

## 7.7 X.509 证书

证书把 **身份（Subject / SAN）** 与 **公钥** 绑在一起，并由 **Issuer（CA）** 背书；带 **Validity** 窗口。

**Go 校验心智**（抽象流程）：`pem.Decode` → `x509.ParseCertificate` → 构造 `x509.CertPool`（根/中间）→ `cert.Verify(x509.VerifyOptions{ DNSName, Roots, Intermediates, … })`。

```go
import (
	"crypto/x509"
)

// 伪代码骨架：真实代码需处理 PEM 链、中间证书、系统根与错误分支
func verifyPeerCert(leaf *x509.Certificate, roots *x509.CertPool, dnsName string) error {
	opts := x509.VerifyOptions{
		DNSName: dnsName,
		Roots:   roots,
	}
	_, err := leaf.Verify(opts)
	return err
}
```

---

## 7.8 TLS（`crypto/tls`）

TLS 握手阶段建立身份与密钥材料，记录层用协商出的对称算法保护流量；证书链校验失败则**必须**中止。

**HTTPS 客户端**（生产**禁止** `InsecureSkipVerify: true`，除非受控联调且带风险隔离）：

```go
// 需 import: crypto/tls, net/http
client := &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
			// RootCAs / 自定义校验仅在需要时设置
		},
	},
}
_, _ = client.Get("https://example.com")
```

**TLS 服务端（自签名仅联调）**：`tls.LoadX509KeyPair` + `tls.Listen`；生产使用 **CA 签发**证书并注意 **SNI、证书链、自动续期**。

```go
// 需 import: crypto/tls, net
cert, err := tls.LoadX509KeyPair("server.pem", "server.key")
if err != nil {
	panic(err)
}
cfg := &tls.Config{
	Certificates: []tls.Certificate{cert},
	MinVersion:   tls.VersionTLS12,
}
ln, err := tls.Listen("tcp", ":8443", cfg)
if err != nil {
	panic(err)
}
defer ln.Close()
// Accept 循环内同样建议 SetDeadline、限流等（见第 3 章）
```

---

## 7.9 本章小结

- **ISO 心智**：服务目标 vs 实现机制；缺认证则加密也可能白做。  
- **分层与 E2E**：TLS 解决传输面大量问题，**不**替代应用授权。  
- **对称**：AES-GCM 等 AEAD；**密钥与 nonce 生命周期**是事故高发区。  
- **非对称**：签名 / 密钥交换 / 混合加密；别用大消息直接 RSA encrypt。  
- **X.509 + TLS**：证书链 + 主机名/SAN 校验是默认信任根。

**背诵版**：**威胁模型 → 选服务目标 → AEAD / 签名 / TLS 组合 → 校验证书与业务授权分层。**  
后续 **HTTP/2、gRPC、WSS** 都把 TLS 当默认底座；本章是读那些章的「密码学与信任」前置。

**前后章节**：[`chapter06` 编码](../chapter06-charset-encoding/note.md) · [`chapter04` 序列化](../chapter04-data-serialization/note.md) · [`chapter08` HTTP](../../03-web-core-stage/chapter08-http/note.md)
