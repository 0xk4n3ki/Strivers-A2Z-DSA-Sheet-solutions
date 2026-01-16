#include <bits/stdc++.h>

void printNum(int n) {
    if(n==0) {
        return;
    }
    printNum(n-1);
    std::cout << n << " " ;
}

int main() {
    int n;
    std::cout << "Enter n: ";
    std::cin >> n;

    printNum(n);
    return 0;
}