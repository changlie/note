> https://softwareengineering.stackexchange.com/questions/231758/would-this-interpreter-be-considered-a-virtual-machine


It depends. "executing" isn't really the defining part of a virtual machine.

To qualify as a virtual machine, an interpreter would have to offer:

- a memory model (ways for a program to acquire, use and release memory - is there garbage collection?)
- APIs for IO (ways for programs to access files, networks, etc.)
- optionally, a concurrency model (ways for programs to execute in parallel and coordinate)

And all of these would have to be well-defined and independent of the underlying OS and hardware. An interpreter that does all of this would have a good claim of being a VM.



This interprenter is a very simple one, and is a Java program that runs on the JVM. As to memory, basically when the programmer declares a variable, the inner implementation of the interprenter to deal with this is to use Java Collections to store and access data that the programmer consideres as 'variables'. Basically, the interprenter uses high-level Java operations to implement everything. Just like the first JVM used high-level C operations to implement things (although ofcourse far, far more complex).




