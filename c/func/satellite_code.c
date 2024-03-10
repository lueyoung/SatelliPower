#include "satellite_code.h"

double R = 1;

double Um = 2.43;
double Uk = 2.73;
double Im = 0.0167 * 24;
double Ik = 0.0172 * 24;

double T = 80.0;
double G = 1000.0;
double Angle = 0.0;

void PVfunction() {
    M3 = ((T - T_ref) * a1 + 1) * G / G_ref;
    M4 = log((G - G_ref) * b1 + e) * (1 - c1 * (T - T_ref));
    C1 = (1 - Im / Ik) * exp(-Um / C2 / Uk);
    C2 = (1 - Um / Uk) / log(1 - Im / Ik);
    Iout = (1 - C1 * exp(Um / C2 / Uk - 1)) * M3 * Ik;
    Ipv0 = cos(Angle) * Iout * Np*R;
}

double f(double y, double q) {
    return Uoc-uref+(Ipv-Iload) * r;
}

double g(double y, double q) {
    return (Ipv-Iload) * k_1;
}

void runge_kutta(double* y, double* q, double h) {
    double k1, k2, k3, k4, l1, l2, l3, l4;
    double temp_y, temp_q;

    k1 = h * g(*y, *q);
    l1 = h * f(*y, *q);

    k2 = h * g(*y + l1/2, *q + k1/2);
    l2 = h * f(*y + l1/2, *q + k1/2);

    k3 = h * g(*y + l2/2, *q + k2/2);
    l3 = h * f(*y + l2/2, *q + k2/2);

    k4 = h * g(*y + l3, *q + k3);
    l4 = h * f(*y + l3, *q + k3);

    temp_q = *q + (k1 + 2*k2 + 2*k3 + k4)/6;
    temp_y = *y + (l1 + 2*l2 + 2*l3 + l4)/6;

    *q = temp_q;
    *y = temp_y;
}

void writeExcel()
{
		int i;
        FILE* fp = NULL;
        fp = fopen("test.csv","w");
        for (i = 0; i < 19440; i++)
            fprintf(fp, "%-10.5f,%-10.5f,%-10.5f,%-10.5f\n", m [i], n [i], hh [i], p [i]);
        fclose(fp);
}

void satellite() {
    double y = 0;
    double q = 96;
    double h = 1;
    double t = 0;

    printf("t = %-6.2f, y1 = %-10.5f, Q = %-10.5f\n", t, y, q);
    int i;
    double u1;
    double u1_max = 20.0;
    double u1_min = 0.0;
    double q_max = 120.0;
    double q_min = 0.0;

    for (i = 0; i <= 19440; i++) {

		 SOC = q/q_0;

         if (q >= q_max)
             q = q_max;
         if (q <= q_min)
             q = q_min;

         Uoc = SOC*k_2+b;

         u1 = (Ubus - 45) * kp + ki * y;

         if (u1 >= u1_max)
             u1 = u1_max;
         if (u1 <= u1_min)
             u1 = u1_min;

         Ipv = (20-u1)*k_3*Ipv0;

         Ipv1= Ipv / 2;
         Ipv2 = Ipv / 2;

		runge_kutta(&y, &q, h);
        PVfunction(&Ipv0);

        Ubus = Uoc + Ibat * r;
        Ibat = Ipv - Iload;

        double period = 6480.0;
        int t_period = (int)(t / period);
        double t_mod = t - period * t_period;
        if (t_mod < 0.0) {
            t_mod += 6480.0;
        }
        if (0.0 < t_mod && t_mod <= 420.0) {
            Pload = 4500.0;
        }
        else if (420.0 < t_mod && t_mod <= 780.0) {
            Pload = 3500.0;
        }
        else if (780.0 < t_mod && t_mod <= 2100.0) {
            Pload = 2500.0;
        }
        else if (2100.0 < t_mod && t_mod <= 3780.0) {
            Pload = 2000.0;
        }
        else if (3780.0 < t_mod && t_mod <= 6480.0) {
            Pload = 1500.0;
        }
        Iload = Pload / Ubus;

        t += h;

        printf("t = %-6.2f, y1 = %-10.5f, Q = %-10.5f ", t, y, q);

        printf("Ubus = %-10.5f,SOC = %-10.5f,Ipv = %-10.5f\n",Ubus, SOC, Ipv);
        m [i] = t;
        n [i] = SOC;
        hh [i] = Ipv;
        p [i] = Ubus;
    }
    //writeExcel();
}
