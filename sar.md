sar -P 0 1 1
Linux 5.13.0-51-generic (suehyun-950XDC-951XDC-950XDX) 2025년 01월 11일 _x86_64_ (8 CPU)

12시 33분 34초 CPU %user %nice %system %iowait %steal %idle
12시 33분 35초 0 2.00 0.00 1.00 0.00 0.00 97.00
평균값: 0 2.00 0.00 1.00 0.00 0.00 97.00

user + nice 을 합한게 사용자 모드에서 프로세스를 실행하는 시간 비율

idle이 100에 가까울수록 CPU는 아무 일도 안했다는 뜻이다.

taskset -c 0 프로그램 실행명령어 &:

프로그램을 CPU 0에서 실행시키겠다는 의미이다.
