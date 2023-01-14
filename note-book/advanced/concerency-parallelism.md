Concurency:
    Processes drop the CPU core that it is using when it is in a "waiting state".
        - this is in contrast to serial processes that lock the CPU while it is waiting for some type of response from human or application itself
    GoLang is concurent by default but also has specific things built into the language to set concurency items


Stronger that typical "multi threaded" languages
    GO Routine = is a set of methods that run in paralell on a single thread

Concurency vs Parallelism
Watch this! : https://www.youtube.com/watch?v=oV9rvDllKEg&ab_channel=gnbitcom

Concurency is about "dealing" with a lot of things at once
Parallelism is about "doing" with a lot of things at once
    - related but not the same things
    - concurency is about structure and paralelism is about execution

Concurency is the structure that allows for parallelism (but doesn't enforce it)






Paper by tony hore "communicating sequential processes" (go read it)


