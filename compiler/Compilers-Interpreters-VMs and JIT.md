>  https://www.i-programmer.info/babbages-bag/352-interpreters.html


The distinction between a compiler and an interpreter is one that can cause controversy. One programmer's compiler is another's interpreter and the whole subject gets very murky when you throw in the idea of the Virtual Machine and Just In Time compilation. So what is it all about?

If you want to start an argument between programmers or software engineers then just say “language X is implemented by an interpreter”.

If that doesn’t cause a heated exchange, then change “interpreter” to “compiler” in the same sentence. For while there are languages that are generally accepted as being implemented usually as one or the other, the situation is rarely clear cut.

Back in the early days of computing the distinction was much more important and the debate about compilers versus interpreters would rage in the pages of any magazine or journal and verbal wars were common. Today, with more powerful machines, the issue isn’t quite as important but it can still raise a good argument! 

It can also be the deeper cause of a language war. For example, C++ v C# isn't just about which language is better it is also about the way that the languages are generally implemented. C++ is usually compiled to machine code where C# is usually compiled to an intermediate code and then interpreted or JITed.

But this is getting ahead of our selves.

As well as being controversial it is also still at the leading edge of development although now we tend to talk more about “virtual machines” and “Just In Time Compilers” than straightforward interpreters and compilers and this makes everything dissolve into shades of grey.

In other articles we have looked at the progression from machine code to assembler and on to high-level languages. An inherent assumption in this discussion is that the language would always be translated, or compiled, to machine code before it was expected to do anything much - but this isn’t the only possible approach. 

The run time
The idea of an interpreter as opposed to a compiler evolved slowly and it isn’t very easy to say exactly where the idea came from.

In the early days assemblers and compilers would translate every last instruction of a high-level language program into machine code and the resulting machine code was then taken and run on a "real" machine. When you think about it what other way could you get a program written in say, Fortran, to run on hardware that knew nothing of Fortran and only understood machine code? 

In other words the process of compiling from a language to machine code seemed unavoidable and the only way to do things. This was a major problem however as no one had any idea how to do it and it wasn't at all obvious that the resulting automatically generated code would be fast enough to be useful. Eventually the first compilers were created and they proved valuable.

Then some clever programmer had an idea. Why not make use of a library of machine code subroutines to make life easier?

For example, if the high level language program contained the line

A=B*C

then a standard compiler approach would take the “B times C” part of the and translate it to the machine code equivalent. Which, if you were lucky would be something like:

MUL B,C

However, many early machines didn’t have a multiply command and so the multiplication had to be built up using add and shift instructions. A true compiler would and should take the high level multiplication B*C and convert it into a complete sequence of machine code instructions that performed the multiplication. 

You can see that implemented this way a simple operation of multiplication would generate a lot of machine code each time you used it.

The clever idea was to create a subroutine that would multiply two numbers together and then a multiplication operation would be compiled to a call to the new subroutine.

That is B*C compiles to:

CALL Multiply

Of course the values in B and C would have to be loaded into the locations in which the subroutine expected to find them before the call, but in principle this is still a great simplification.

The advantage of this approach is that the compiled program is smaller because the multiplication code isn’t repeated every time it is used.

However there are some disadvantages to the method. The first is that there is usually an overhead in calling a subroutine that makes this approach slower than simply compiling in the instructions needed to multiply two numbers together. The second is that now any program that runs on the machine compiled in this way needs access to a chunk of standard code – usually called a “run time library”.

Compilers that use run time libraries are common and they can hide the fact that they are doing so by including a copy of the library with each and every program. This, of course wastes space but it does produce programs that can run without the obvious help of a run time library.

Smarter compilers only compile in the bits of the run time library that a program actually uses.This is what a linker does. It takes the separate sections of a program including any routines needed from the run time library and stitches it all together to create a complete program.

You can even package much of the run time library into separate modules which are used at run time and not compiled into the application - this is what Windows DLLs and Linux SO files are all about. The are shared libraries that are precompiled and available for use by any program.

As time went on run time libraries tended to get bigger and bigger and more sophisticated. The consequence is that more and more of the compiler's time is spent not in compiling to machine code but simply putting in calls to routines in the run time. 

The Interpreter
The idea of using a run time library to implement multiplication and similar operations doesn’t seem like a deep philosophical principle – but it is!

What we have done is to notice that the high level language makes use of a facility that the machine doesn’t actually have.

The run time subroutine can be thought of as providing the missing facility by simulating it in software. If at a later date the machine suddenly gets a piece of hardware that does multiplication the compiler can be changed to make use of it by issuing a single multiply instruction instead of a call to the multiply subroutine. Alternatively the run time library can be changed to make use of the new hardware – this is very flexible.

Now think a little harder about his idea and you quickly see that it can be extended to greater things.

The multiplication routine is there to make up for the lack of a multiply command in the real hardware.

Why not simulate all of the hardware you need and make the run time library a software machine that runs the high-level language in question.

After all A*B is an instruction in say Fortran or Basic and the run time can simply do the multiplication as soon as it encounters the instruction. The same is true for all of the instructions written in the high-level language. Each instruction results in a subroutine being called to do what ever it is the instruction is about. 

Now we don't have an obvious translation step from the high level language to machine code. In fact no machine code is generated. All that happens is that the "run time" package now reads the high level language and obeys its instructions by calling appropriate subroutines. 

That is, the run time library has become a computer, implemented in software, which has the high-level language as its machine code.

This is an amazing idea and another name for the simulation is an interpreter.

In the early days languages such as Basic were implemented using interpreters and no compile step was required. The Basic Interpreter read each of the Basic instruction and did what it was told to do, usually by calling predefined routines. 

One of the earliest to have an impact was Palo Alto Tiny Basic. A miracle of economy by today's standards. A complete integer only Basic interpreter that used just 2K of an 8 bit processors memory. It worked by using a big selection statement to call the appropriate subroutine to handle whatever keyword was next LET, GOTO etc. The actual Basic was never converted to machine code it simply activated the subroutines in the runtime library now more correctly called an interpreter.

The Virtual Machine
However notice that there is another way to look at the interpreter code. It is a software implementation of a machine that runs the high-level language as its "machine code". This is what we generally refer to as a Virtual Machine or VM. 

In the case of Tiny Basic the interpreter can be viewed as a simulation of a machine that has Tiny Basic as its machine code or assembly language.

This idea can be generalized and you can design VMs that have a lower level language than Tiny Basic as their machine code. Such a VM can be used as a target for a range of languages with the help of a compiler that translates the high level language into the VMs machine code. And at this point you might think that world has gone mad - a compiler that compiles to a made up machine code and then run on a VM in the style of an interpreter! Surely this is crazy?!


Battle lines
Now we have two ways of implementing a high level language – we can compile it to machine code or we can run it on an interpreter for which it IS the machine code.

Traditionally the argument for and against the two approaches goes something like this:

A compiler produces “tight efficient code”.
This is supposed to mean that because it generates machine code everything happens as fast as possible. Of course this is nonsense because the machine code could make a lot of use of run time subroutines and so start to slide towards the interpreter approach.
A compiler produces “small stand-alone code”.
Clearly if it uses a run time library then it isn’t stand-alone unless the library is included in the compiled code when it isn’t going to be small!
Conversely an interpreter is said to be slow and wasteful of memory space. 
In fact an interpreter doesn’t have to be slow and a high-level language version of a program can be a lot smaller than a fully compiled machine code version.
So what is the truth?

The fact is that in the past implementations that have described themselves as compilers have been faster than ones that were called interpreters but there has always been a considerable overlap between the two.

Over time the two approaches have tended to become even more blurred and they have borrowed ideas from one another.

For example, the first generation of interpreters usually had excellent debugging facilities. Because the machines they ran on were implemented in software it was an easy task to provide additional facilities that would tell the programmer what was going on.

Interpreters invented idea such as tracing, i.e. following the execution of the program line by line, and dynamic inspection of variable contents etc. Yes this is where Basic's Tron and Troff originated from and from here all breakpoint and trace technology.

As time went on it became clear that these facilities could be built into a compiler as well by augmenting the run time environment to include them. Many compilers will produce a debug version of the code while you are still testing things and a production version when you have finished.

It is true, however, that in the past interpreted languages were more sophisticated and complex then compiled languages. The reason was simply that writing an interpreter seemed to be an easier thing to do than writing a compiler and so the implementation method chosen tended to limit, or expand, the language according to what was perceived as difficult or easy.

It is strange to think that the reason that compilers have seemed to be harder to write might well be the way that they are taught a part of formal grammar. Translating one machine language into another involves the use of a grammar and parsing techniques to work out the structure and this can be made very mathematical and so off putting to many. An interpreter on the other hand is a machine, a soft machine, that has a particular high level language as its machine code and this seems so much more like engineering.

If you make a language static and strongly typed then it seems to be easier to implement using a compiler approach. On the other hand if you use an interpreted approach then its natural to allow the language to be dynamic and allow self modification.

Today we are in a period where static languages such as Java and C# are giving ground to dynamic languages such as Ruby and even JavaScript.

These differences are not new and in many ways they represent the re-emergence of the compiler v interpreter approach to language design and implementation.

Virtual Machines And Intermediate Languages
There is one last development of the interpreter idea that is worth going into in more because it is important today.

An alternative to implementing a machine that runs the high-level language as its machine code is to compile the high-level language to a lower-level language and then run this using an interpreter or VM.  This is madness referred to at the end of the first page.

That is instead of writing an interpreter to run Java we first compile it to a simpler language called byte code. Notice we do not compile it to machine code and byte code is still fairly high level compared to machine code. To actually run the Java we use an interpreter or virtual machine for byte code. 

This might seem like a very strange idea in that you now have the worst of all possible worlds.

You have to use a compiler to translate the program from one language to another and then you have to use an interpreter to run it.

What could possibly be good about this idea?

The answer is a great deal.

The first advantage is that a compiler from a high-level language to an intermediate-level language is easier to write and can be very efficient.

The second is that an interpreter for an intermediate-level language is easier to write and can also be very efficient.

Looking at things another way we get the best, not the worst, of both approaches!

In addition there is one huge advantage which you might not notice at first. If the interpreter for the intermediate-level language is simple enough then it can be easily implemented on any hardware and this makes programs compiled to the intermediate-level code easily portable between different types of hardware.

If you are really clever then you even write the compiler in the intermediate-level language making it portable as well!

In this mode the interpreter is generally called a Virtual Machine or VM. 

That is we generally call a VM that works directly with a high level language an Interpreter. Hence Basic was generally executed by an interpreter. However if the VM runs an intermediate code produced by a compiler we generally call it a VM. Thus Java is executed by a VM and not an interpreter.

This is all the difference amounts to. 

The intermediate language is also generally called Pseudo Code, or P-Code for short. P-Code compilers and VMs were very popular in the time before the IBM PC came on the scene (USCD Pascal and Fortran being the best known). Then they more or less vanished, only to return with in a big way with Java but renamed "byte code".


Java’s main claim to fame is that it is the ultimate portable language.

Java VMs exist for most hardware platforms and up to a point you really can compile a Java program and expect it to run on any machine that has a VM. Not only this but the Java compiler and all of the Java system is itself compiled to byte code and so once you have a VM running on new hardware you also have the entire Java system – clever!

.NET languages such as C# and Visual Basic also use an intermediate language and VM approach but due to Microsoft's proprietary approach to computing neither is quite as portable as Java although with the open sourcing of .NET this is changing very rapidly. You can now find good implementations of the CLR and the entire .NET system on Linux and other operating systems.

This idea is so good that you can expect most language development in the future to be centred on the VM idea. One thing is sure - the future is virtual.

JIT and Not so JIT
The story so far is easy enough to understand. At one end of the spectrum of language implementations we have the pure compiler, which generates nothing but machine code and uses no run time library or package.

At the other end we have the interpreter, which generates no machine code and is all run time package in the form of a complete VM for the language.

Of course, in the real world these really are two ends of the spectrum and real compilers use different amounts of run time library, so slowly sliding towards the interpreter end of the spectrum. But what about interpreters? Do they have another way of sliding towards the compiler end of the spectrum?

An interpreter can generate some machine code to get a job done quicker if this is important. A modern VM will use all sorts of techniques to make it faster and more efficient. For example, for each instruction in the intermediate language the VM could ask which is going to be quicker: to call a routine or to generate machine code to get the job done. This approach is often called Just-In-Time or JIT compilation. It is usually explained as the VM compiling the intermediate language just before it is run, but this isn't really a good way to think of it.

The VM does compile the intermediate language, but mostly what it produces is just lots of calls to routines that constitute a runtime package. So the JIT is a sort of mixture of interpreting the code and compiling the code according to what makes best use of the real machine. 

This is still a source of lots of arguments in the programming world - is it a compiler, an interpreter, a JIT or what?

In practice there is a lot of overlap, but it is still true that languages at the compiler end of the spectrum run faster than languages at the interpreter or VM end of the spectrum. However, the gap isn't as wide as you might think and a lot depends on how well the compiler and VM are implemented. When it comes to efficiency and performance of implementing a language the devil is in the detail rather than the bigger choices.