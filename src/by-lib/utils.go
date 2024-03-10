package main

import (
	"fmt"
	"math"
)

func PVfunction() {
	M3 = ((T-T_ref)*a1 + 1) * G / G_ref
	M4 = math.Log((G-G_ref)*b1+e) * (1 - c1*(T-T_ref))
	C1 = (1 - Im/Ik) * math.Exp(-Um/C2/Uk)
	C2 = (1 - Um/Uk) / math.Log(1-Im/Ik)
	Iout = (1 - C1*math.Exp(Um/C2/Uk-1)) * M3 * Ik
	Ipv0 = math.Cos(Angle) * Iout * Np * R
	//fmt.Println(M3, M4, C1, C2, Iout, Ipv0)
}

func f(y float64, q float64) float64 {
	return Uoc - uref + (Ipv-Iload)*r
}

func g(y float64, q float64) float64 {
	return (Ipv - Iload) * k_1
}

func runge_kutta(y *float64, q *float64, h float64) {
	var k1, k2, k3, k4, l1, l2, l3, l4 float64
	var temp_y, temp_q float64

	k1 = h * g(*y, *q)
	l1 = h * f(*y, *q)

	k2 = h * g(*y+l1/2, *q+k1/2)
	l2 = h * f(*y+l1/2, *q+k1/2)

	k3 = h * g(*y+l2/2, *q+k2/2)
	l3 = h * f(*y+l2/2, *q+k2/2)

	k4 = h * g(*y+l3, *q+k3)
	l4 = h * f(*y+l3, *q+k3)

	temp_q = *q + (k1+2*k2+2*k3+k4)/6
	temp_y = *y + (l1+2*l2+2*l3+l4)/6

	//*q = temp_q
	//*y = temp_y
        fmt.Printf("tmp q : %v, tmp y: %v\n",temp_q, temp_y)
	*q = temp_q
	*y = temp_y
        fmt.Printf("q : %v, y: %v\n",*q, *y)
}

func satellite_func() {
	var y float64 = 0
	var q float64 = 96
	var h float64 = 1
	var t float64 = 0

	fmt.Printf("t = %-6.2f, y1 = %-10.5f, Q = %-10.5f\n", t, y, q)

	var i int
	var u1 float64
	var u1_max float64 = 20.0
	var u1_min float64 = 0.0
	var q_max float64 = 120.0
	var q_min float64 = 0.0
	for i = 0; i < 19441; i++ {

		fmt.Println(q)
		SOC = q / float64(q_0)

		if q >= q_max {
			q = q_max
		}
		if q <= q_min {
			q = q_min
		}

		Uoc = SOC*k_2 + b

		u1 = (Ubus-45)*kp + ki*y

		if u1 >= u1_max {
			u1 = u1_max
		}
		if u1 <= u1_min {
			u1 = u1_min
		}

		Ipv = (20 - u1) * k_3 * Ipv0

		Ipv1 = Ipv / 2
		Ipv2 = Ipv / 2

		runge_kutta(&y, &q, h)
		PVfunction()
		//PVfunction(&Ipv0)
		//fmt.Printf("%v\n",Ipv0)

		Ubus = Uoc + Ibat*r
		Ibat = Ipv - Iload

		var period float64 = 6480.0
		var t_period int
		t_period = (int)(t / period)
		var t_mod float64 = t - period*float64(t_period)
		if t_mod < 0.0 {
			t_mod += 6480.0
		}
		if 0.0 < t_mod && t_mod <= 420.0 {
			Pload = 4500.0
		} else if 420.0 < t_mod && t_mod <= 780.0 {
			Pload = 3500.0
		} else if 780.0 < t_mod && t_mod <= 2100.0 {
			Pload = 2500.0
		} else if 2100.0 < t_mod && t_mod <= 3780.0 {
			Pload = 2000.0
		} else if 3780.0 < t_mod && t_mod <= 6480.0 {
			Pload = 1500.0
		}
		Iload = Pload / Ubus

		t += h

		//fmt.Printf("t = %-6.2f, y1 = %-10.5f, Q = %-10.5f ", t, y, q)

		//fmt.Printf("Ubus = %-10.5f,SOC = %-10.5f,Ipv = %-10.5f\n", Ubus, SOC, Ipv)
		m[i] = t
		n[i] = SOC
		hh[i] = Ipv
		p[i] = Ubus
	}
}
