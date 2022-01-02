package main
import "fmt"
//import "crypto/sha256"
//求逆元
func getInverse(input int, mod int) int{
	var output int = -1
	for i:=1; i<mod; i++{
		if (input * i)%mod == 1{
			output = i
			break
		}
	}
	return output
}
//求最大公约数
func getGcd(a int, b int) int{
	if b == 0 {
		a, b = b, a
	}
	var ret int
	if a % b == 0{
		ret = b
		return ret
	}
	ret = getGcd(b, a % b)
	return ret
}
//求椭圆曲线加法
func getAdd(x1, y1, x2, y2, a, mod int)(x3 int,y3 int){
	var k int
	var dy, dx int
	var flag bool
	if x1 == x2 && y1 == y2{
		dx = 2*y1
		dy = (3*x1*x1 + a)
	}else{
		dy = y2-y1
		dx = x2-x1
	}
	if dy * dx < 0{
		flag = true
	}
	if dy < 0 {
		dy = -dy
	}
	if dx < 0 {
		dx = -dx
	}
	ret := getGcd(dy, dx)
	dy = dy/ret
	dx = dx/ret
	dx = getInverse(dx, mod)
	k = (dy * dx)
	if flag {
		k = -k
	}
	k = k % mod
	x3 = (k * k - x1 - x2) % mod
	y3 = (k * (x1 - x3) - y1) % mod
	if x3 < 0{
		x3 = mod + x3
	}
	if y3 < 0{
		y3 = mod + y3
	}
	return
}
//获得最大阶数
func getOrder(x1, y1, a, mod int) int{
	x0 := x1
	y0 := mod + (-1 * y1 ) % mod
	x2 := x1
	y2 := y1
	n := 1
	var graph [][]int = make([][]int, mod)
	for index, _ := range graph{
		graph[index] = make([]int, mod)
	}
	graph[x1][y1] = 1
	for {
		x2, y2 = getAdd(x1, y1, x2, y2, a, mod)
		if x2 == x0 && y2 == y0{
			break
		}
		n++
		graph[x2][y2] = n
	}
	for _, g := range graph{
		fmt.Println(g)
	}
	return n
}
//寻找加密椭圆曲线所有点
func getDot(x0, a, b, mod int) (y0, x1, y1 int, ok bool){
	for i:=0; i<mod; i++{
		left := (i*i)%mod
		right := (x0*x0*x0 + a*x0 + b)%mod
		if right < 0{
			right = mod + right 
		}
		if left == right{
			y0 = i
			ok = true
			break
		}
	}
	x1 = x0
	if -y0 < 0{
		y1 = mod + (-y0) % mod
	}else{
		y1 = y0 % mod
	}
	return
}
//画图
func getGraph(a, b, mod int){
	var graph [][]int = make([][]int, mod)
	for index, _ := range graph{
		graph[index] = make([]int, mod)
	}
	for x0:=0; x0<mod; x0++{
		y0, x1, y1, ok := getDot(x0, a, b, mod)
		if !ok {
			continue
		}
		graph[x0][y0] = 1
		graph[x1][y1] = 1
	}
	for _, g := range graph{
		fmt.Println(g)
	}
}
//生成非对称公钥，私钥
func getKey()(int, int, int, int, int, int, int, int){
here:
	fmt.Println("请输入椭圆曲线参数a=?,b=?,mod=?:y^2=x^3+ax+b(%mod)")
	var a, b, mod int
	fmt.Scanf("a=%d,b=%d,mod=%d\n", &a, &b, &mod)
	if (4*(a*a*a)+27*(b*b))%mod == 0{
		fmt.Println("椭圆曲线不合适")
		goto here
	}
	getGraph(a,b,mod)
	fmt.Println("请从上面点选一个起点x=?,y=?")
	var x1, y1 int
	fmt.Scanf("x=%d,y=%d\n", &x1, &y1)
	n := getOrder(x1, y1, a, mod)
	fmt.Println("n=",n)
	fmt.Println("请输入私钥key=?小于n")
	var key int
	fmt.Scanf("key=%d\n", &key)
	x2, y2 := getNG(x1, y1, x1, y1, a, mod, key)
	return a, b, mod, key, x1, y1, x2, y2
}
//计算rG, rQ
func getNG(x1, y1, x2, y2, a, mod, r int) (int, int){
	for i:=0; i<r; i++{
		x2, y2 = getAdd(x1, y1, x2, y2, a, mod)
		fmt.Println(x1, y1, x2, y2, a, mod)
	}
	return x2, y2
}
//公钥加密
func encrypt(x1, y1, x2, y2, a, mod int) []int {
here:
	fmt.Println("请输入r=?用于计算rG,rQ")
	var r int
	fmt.Scanf("r=%d\n", &r)
	//公钥加密：选择随机数r，将消息M生成密文C，该密文是一个点对，C = {rG, M+rQ}，其中Q为公钥。c={rG, M*rQx} 
	//私钥解密：M + rQ - d(rG) = M + r(dG) - d(rG) = M，其中d、Q分别为私钥、公钥。M*rQx/d(rGx) = M*rdGx/rdGx = M
	rGx, rGy := getNG(x1, y1, x1, y1, a, mod, r)
	rQx, _ := getNG(x1, y1, x2, y2, a, mod, r)
	if rQx == 0{
		fmt.Println("r 不合适，rQx==0")
		goto here
	}
	fmt.Println("rQx=", rQx, rGx, rGy)
	fmt.Println("请输入要加密的字符串:")
	var str string
	fmt.Scanf("%s\n", &str)
	var target []int
	target = append(target, rGx)
	target = append(target, rGy)
	var tmp int
	for i:=0; i<len(str); i++{
		tmp = int(str[i])
		tmp = tmp * rQx
		target = append(target, tmp)
	}
	fmt.Println(target)
	return target
}
//私钥解密
func decrypt(x1, y1 int, c []int, key, a, mod int){
	var m []byte
	var tmp int
	rGx := c[0]
	rGy := c[1]
	rQx, _ := getNG(x1, y1, rGx, rGy, a, mod, key)
	fmt.Println("rQx=", rQx, rGx, rGy)
	for i:=2; i<len(c); i++{
		tmp = c[i]/rQx
		m = append(m, byte(tmp))
	}
	fmt.Println(string(m))
}
//求hash
func hash256(str string) int {
	/*h := sha256.New()
	h.Write([]byte(str))
	return h.Sum(nil)*///256位,32字节
	return len(str)//压缩为长度信息
}
//签名消息结构
type Sign struct{
	M string
	RGx int
	RGy int
	S int
}
//私钥签名
//根据随机数r、消息M的哈希h、私钥d，计算s = (h + key*rGx)/r。将消息M、和签名{rG, s}发给接收方
func signature(x1, y1, key, a, b, mod int) Sign {
	var sign Sign
	var r int
	var str string
	fmt.Println("请输入r=?用于签名")
	fmt.Scanf("r=%d\n", &r)
	rGx, rGy := getNG(x1, y1, x1, y1, a, mod, r)
	sign.RGx = rGx
	sign.RGy = rGy
	
	fmt.Println("请输入要签名的字符串:")
	fmt.Scanf("%s\n", &str)	
	sign.M = str
	h := hash256(str)
	sign.S = ((h + key*rGx)*getInverse(r, mod)) % mod
	fmt.Println("rGx=", rGx, "rGy=", rGy, "h=", h, "s=", sign.S)
	return sign
}
//公钥验证
//使用发送方公钥Q计算：hG/s + (rGx)Q/s，并与rG比较，如相等即验签成功。
//原理：hG/s + rGxQ/s = hG/s + rGx*keyG/s = (h+key*rGx)G/s = r(h+key*rGx)G / (h+key*rGx) = rG
func verifySign(sign Sign, x1, y1, Qx, Qy, a, mod int) bool {
	h := hash256(sign.M)
	s2 := getInverse(sign.S, mod)
	u1 := (h*s2)%mod
	u2 := (sign.RGx*s2)%mod
	u1Gx, u1Gy := getNG(x1, y1, x1, y1, a, mod, u1)
	u2Qx, u2Qy := getNG(Qx, Qy, Qx, Qy, a, mod, u2)
	sx, sy := getAdd(u1Gx, u1Gy, u2Qx, u2Qy, a, mod)
	fmt.Println("sx=", sx, "sy=", sy, "s2=", s2)
	if sx == sign.RGx && sy == sign.RGy{
		return true
	}
	return false
}

func main(){
	//x3, y3 := getAdd(3, 10, 3, 10, 1, 23)
	//fmt.Println("x3=",x3,"y3=",y3)
	//pemx, pemy := getPem(3, 10, 10, 1, 23)
	//fmt.Println("pemx=",pemx,"pemy=",pemy)
	a, b, mod, key, x1, y1, x2, y2 := getKey()
	fmt.Printf("a=%d, b=%d, mod=%d, key=%d, x1=%d, y1=%d x2=%d, y2=%d\n",a, b, mod, key, x1, y1, x2, y2)
	c := encrypt(x1, y1, x2, y2, a, mod)
	decrypt(x1, y1, c, key, a, mod)
	sign := signature(x1, y1, key, a, b, mod)
	ret := verifySign(sign, x1, y1, x2, y2, a, mod)
	fmt.Println("verify=", ret)
}
