所以不要以为所有的异常都能够被 recover 到，实际上像 fatal error 和 runtime.throw 都是无法被 recover 到的，
甚至是 oom 也是直接中止程序的，也有反手就给你来个 exit(2) 教做人。
因此在写代码时你应该要相对注意些，“恐慌” 是存在无法恢复的场景的



