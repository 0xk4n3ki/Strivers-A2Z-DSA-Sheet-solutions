#include <bits/stdc++.h>

void printNum(int n) {
    if(n==0) {
        std::cout << std::endl;
        return;
    }
    std::cout << n << " ";
    printNum(n-1);
}

int main() {
    int n;
    std::cout << "Enter n: ";
    std::cin >> n;

    printNum(n);
    return 0;
}