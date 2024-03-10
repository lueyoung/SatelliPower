#ifndef _SATELLITE_CODE_H_
#define _SATELLITE_CODE_H_
#include <stdio.h>
#include <math.h>

#define kp 1
#define k_2 9
#define b 36
#define uref 45
#define ki 0.1
#define k_3 0.2
#define q_0 120
#define r 0.04
#define k_1 1/3600
#define T_ref 25
#define G_ref 1000
#define a1 0.003
#define b1 0.0005
#define c1 0.0029
#define e 2.71828
#define Np 216
#define Ns 27

double m[19440];
double n[19440];
double hh[19440];
double p[19440];

double Iload;
double Pload;
double Uoc;
double Ipv;
double Ipv1;
double Ipv2;
double SOC;
double Ubus;
double R;
//double R = 1;
double Ibat;

double M1, M2, M3, M4;
double C1, C2;
double Iout;
double Ipv0;
double Um;
double Uk;
double Im;
double Ik;
//double Um = 2.43;
//double Uk = 2.73;
//double Im = 0.0167 * 24;
//double Ik = 0.0172 * 24;

double T;
double G;
double Angle;
//double T = 80.0;
//double G = 1000.0;
//double Angle = 0.0;

void PVfunction();

double f(double, double);

double g(double, double);

void runge_kutta(double*, double*, double);

void writeExcel();
void satellite();

#endif
