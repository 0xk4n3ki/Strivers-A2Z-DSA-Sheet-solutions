#include <bits/stdc++.h>

void printName(int n) {
    if(n==0)
        return;
    std::cout << "k4n3ki" << std::endl;
 
    printName(n-1);
}

int main() {
    int n;
    std::cout << "Enter n: ";
    std::cin >> n;

    printName(n);
    return 0;
}