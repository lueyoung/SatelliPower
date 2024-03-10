#ifndef __SATELLITE_CODE_H__
#define __SATELLITE_CODE_H__
#include <stdio.h>
#include <math.h>

#define kp 1
#define k_2 9
#define b 36
#define uref 45
#define ki 0.1
#define k_3 0.2
#define q_0 120
#define rxxx 0.04
#define k_1 1/3600
#define T_ref 25
#define G_ref 1000
#define a1 0.003
#define b1 0.0005
#define c1 0.0029
#define e 2.71828
#define Np 216
#define Ns 27

void PVfunction();

double f(double, double);

double g(double, double);

void runge_kutta(double*, double*, double);

void writeExcel();

void satellite();

double test();
double add(double, double);

#endif
