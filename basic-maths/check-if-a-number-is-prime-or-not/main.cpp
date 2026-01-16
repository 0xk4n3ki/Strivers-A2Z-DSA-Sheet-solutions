#include <bits/stdc++.h>

int main() {
    int n;
    std::cout << "Enter n: ";
    std::cin >> n;
    
    int i;
    for(i = 2; i*i <= n; i++) {
        if(n%i==0) {
            std::cout << n << " is not a prime number" << std::endl;
            break;
        }
    }

    if(i*i>n){
        std::cout << n << " is a prime number" << std::endl;
    }
    return 0;
}