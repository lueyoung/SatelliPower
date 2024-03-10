#include <stdio.h>
#include<math.h>

//母线电压：Ubus  
//光照期标志位：R（光照区为1，地影期为0）
//蓄电池电量：q（0~120Ah）
//母线电流：Iload；
//太阳帆板输出电流：Ipv；
//蓄电池电流：Ibat;(充电为正，放电为负)
//蓄电池组电压：Uoc；
//入射角：Angle （表示入射光线与帆板法线的夹角）
//Ipv1为+Y太阳阵电流；Ipv2为-Y太阳阵电流

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
double p[19440];//存储数组

double Iload;//母线电流
double Pload;//母线功率
double Uoc;
double Ipv;
double Ipv1;
double Ipv2;
double SOC;//蓄电池荷电状态
double Ubus;//母线电压
double R = 1;//标志位，光照期为1，地影期为0
double Ibat;//蓄电池电流

double M1, M2, M3, M4;
double C1, C2;
double Iout;
double Ipv0;
double Um = 2.43;
double Uk = 2.73;
double Im = 0.0167 * 24;
double Ik = 0.0172 * 24;//太阳电池板参数

double T = 80.0;//温度
double G = 1000.0;//光照强度
double Angle = 0.0;//入射角

//设置太阳电池板电流
void PVfunction() {
    M3 = ((T - T_ref) * a1 + 1) * G / G_ref;
    M4 = log((G - G_ref) * b1 + e) * (1 - c1 * (T - T_ref));
    C1 = (1 - Im / Ik) * exp(-Um / C2 / Uk);
    C2 = (1 - Um / Uk) / log(1 - Im / Ik);
    Iout = (1 - C1 * exp(Um / C2 / Uk - 1)) * M3 * Ik;
    Ipv0 = cos(Angle) * Iout * Np*R;
}

// 定义微分方程组
double f(double y, double q) {
    return Uoc-uref+(Ipv-Iload) * r;  // 假设 f(y, q)
}

double g(double y, double q) {
    return (Ipv-Iload) * k_1;  // 假设 g(y, q) = y + q
}

// 定义龙格库塔算法
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

//写入文件 
void writeExcel()
{		
		int i;
        FILE* fp = NULL;
        fp = fopen("test.csv","w");
        for (i = 0; i < 19440; i++)
            fprintf(fp, "%-10.5f,%-10.5f,%-10.5f,%-10.5f\n", m [i], n [i], hh [i], p [i]);
        fclose(fp);
}
    
int main() {
    double y = 0;   // 初始值 y0
    double q = 96;   // 蓄电池电量初始值 q0
    double h = 1; // 步长
    double t = 0;   // 初始时间

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
             q = q_min;//蓄电池Q值给定上下限

         Uoc = SOC*k_2+b;
         
         u1 = (Ubus - 45) * kp + ki * y;
        
         if (u1 >= u1_max)
             u1 = u1_max;
         if (u1 <= u1_min)
             u1 = u1_min;//控制电压给定上下限
         
         Ipv = (20-u1)*k_3*Ipv0;
        
         Ipv1= Ipv / 2;
         Ipv2 = Ipv / 2;

		runge_kutta(&y, &q, h);
        PVfunction(&Ipv0);

        Ubus = Uoc + Ibat * r;
        Ibat = Ipv - Iload;
        
        //设置负载功率变化
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
    writeExcel();
    return 0;

}
