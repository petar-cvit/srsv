gen:
	gcc generator.c -o ./generator.exe; ./generator.exe $(tasks) $(duration)

posluzitelj:
	gcc -ggdb3 -O3 -std=c99 -Wall -Wextra -pedantic posluzitelj.c -o ./posluzitelj.exe -pthread; ./posluzitelj.exe $(timeToFinish) $(numberOfThreads)
