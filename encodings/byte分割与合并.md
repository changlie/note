java 代码
```java
byte raw = (byte) 0b11101010;
System.out.println("init raw: " + raw);

byte head = (byte) (raw >>> 4 & 0x0f);
byte tail = (byte) (((byte)(raw <<  4)) >>>  4 & 0b00001111);
System.out.println("\nsplit: ");
System.out.println("head: " + head + ", " + Integer.toBinaryString(head));
System.out.println("tail: " + tail + ", " + Integer.toBinaryString(tail));

System.out.println("\nmerge: ");
System.out.println((byte)((head<<4 | tail) & 0xff));
System.out.println((byte)((head<<4 | tail)));

System.out.println("\nbit show: ");
System.out.println(Integer.toBinaryString((head<<4 | tail)));
System.out.println(-1 >>> 8); // 无符号位移对int类型才有效。
System.out.println(Integer.toBinaryString(-1 & 0xff));
System.out.println(Integer.toBinaryString(-2 & 0xff));
System.out.println(Integer.toBinaryString(-3));
System.out.println(Integer.toBinaryString(-127));
System.out.println((char)0b11101010);
```

输出
```
init raw: -22

split: 
head: 14, 1110
tail: 10, 1010

merge: 
-22
-22

bit show: 
11101010
16777215
11111111
11111110
11111111111111111111111111111101
11111111111111111111111110000001
ê
```
