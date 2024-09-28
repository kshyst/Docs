#include <stdlib.h>

int main() {
    
    system("cd S:/Docs");
    system("git add .");
    system("git commit -m \"Update\"");
    system("git push doc main");
    return 0;
}