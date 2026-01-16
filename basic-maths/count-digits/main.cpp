#include <bits/stdc++.h>

int main() {
    int n;
    std::cout << "Enter n: ";
    std::cin >> n;

    int count = 0;
    while(n>0) {
        count++;
        n = n/10;
        // std::cout << n << std::endl;
    }
    std::cout << "number of digits is: " << count << std::endl;
    
    std::cout << "Enter n: ";
    std::cin >> n;

    count = (int)(log10(n)+1);
    std::cout << "number of digits is: " << count << std::endl;
    
    return 0;
}