

所有linux系统调用都会对应一个编号，可以直接使用这个编号在汇编语言中进行系统调用。
```
/usr/include/x86_64-linux-gnu/asm
```
```c
This file should list the numbers of the system calls the system knows.
But instead of duplicating this we use the information available
from the kernel sources.
# ifdef __i386__
#  include <asm/unistd_32.h>
# elif defined(__ILP32__)
#  include <asm/unistd_x32.h>
# else
#  include <asm/unistd_64.h>
# endif
```