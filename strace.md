strace -T -o hello.log ./main

=> T옵션: 각 시스템 콜에 대해 소요시간을 마이크로초 단위로 알 수 있다.

각 리눅스 명령어와 프로그램이 사용하는 표준 C 라이브러리 알아보기: ldd

```
ldd /bin/echo
        linux-vdso.so.1 (0x00007fff921c7000)
        libc.so.6 => /lib/x86_64-linux-gnu/libc.so.6 (0x00007fd43485d000)
        /lib64/ld-linux-x86-64.so.2 (0x00007fd434a6e000)

suehyun@suehyun-950XDC-951XDC-950XDX:~/바탕화면/linux_study$ ldd /bin/cat
        linux-vdso.so.1 (0x00007ffd1e78e000)
        libc.so.6 => /lib/x86_64-linux-gnu/libc.so.6 (0x00007f0db4abf000)
        /lib64/ld-linux-x86-64.so.2 (0x00007f0db4cd1000)

suehyun@suehyun-950XDC-951XDC-950XDX:~/바탕화면/linux_study$ ldd ./main
        동적 실행 파일이 아닙니다

 ldd /usr/bin/go
        동적 실행 파일이 아닙니다
```
