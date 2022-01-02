参考这篇文章https://zhuanlan.zhihu.com/p/101907402
go语言从底层实现ECC加密算法


go run .\ecc.go

请输入椭圆曲线参数a=?,b=?,mod=?:y^2=x^3+ax+b(%mod)
a=4,b=2,mod=73
请从上面点选一个起点x=?,y=?
x=5,y=1
n= 85
n是否太小y/n
n
请输入私钥key=?小于n
key=1
a=4, b=2, mod=73, key=1, x1=5, y1=1 x2=5, y2=1
x3 68 y3 21 x2 72 y2 17 x34 47 y34 34 x26 47 y26 34
Qx= 13 Qy= 39 rQx2= 47 rQy2= 34
rGx2= 72 rGy2= 17 rQx3= 47 rQy3= 34
请输入r=?用于计算rG,rQ
r=3
rQx= 68 68 21
请输入要加密的字符串:
ninini
[68 21 7480 7140 7480 7140 7480 7140]
rQx= 68 68 21
ninini
请输入r=?用于签名
r=49
请输入要签名的字符串:
xz
rGx= 16 rGy= 72 h= 2 s= 54
sx= 16 sy= 72 s2= 23 h 2 u1 46 u2 3
u1Gx 25 u1Gy 55 u2Qx 68 u2Qy 21
