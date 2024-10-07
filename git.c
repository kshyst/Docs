#include <stdlib.h>

int main() {
    system("git add .");
    system("git commit -m \"Update\"");
    system("git push 1 main");
    return 0;
} 