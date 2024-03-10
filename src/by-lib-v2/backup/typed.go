package main

const (
	kp       = 1
	k_2      = 9
	b        = 36
	uref     = 45
	ki       = 0.1
	k_3      = 0.2
	q_0      = 120
	r        = 0.04
	k_1      = 1 / 3600
	T_ref    = 25
	G_ref    = 1000
	a1       = 0.003
	b1       = 0.0005
	c1       = 0.0029
	e        = 2.71828
	Np       = 216
	Ns       = 27
	MAX_SIZE = 19941
)

var (
	m  [MAX_SIZE]float64
	n  [MAX_SIZE]float64
	hh [MAX_SIZE]float64
	p  [MAX_SIZE]float64

	Iload float64
	Pload float64
	Uoc   float64
	Ipv   float64
	Ipv1  float64
	Ipv2  float64
	SOC   float64
	Ubus  float64
	R     float64 = 1
	Ibat  float64

	M1, M2, M3, M4 float64
	C1, C2         float64
	Iout           float64
	Ipv0           float64
	Um             float64 = 2.43
	Uk             float64 = 2.73
	Im             float64 = 0.4008
	Ik             float64 = 0.0172 * 24

	T     float64 = 80.0
	G     float64 = 1000
	Angle float64 = 0.0
)
