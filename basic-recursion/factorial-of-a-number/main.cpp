#include <bits/stdc++.h>

int factorial(int n) {
    if(n==1) {
        return 1;
    }
    return n*factorial(n-1);
}

int main() {
    int n;
    std::cout << "Enter n: ";
    std::cin >> n;

    int ans = factorial(n);
    std::cout << "factorial of " << n << " numbers is: " << ans << std::endl;
    return 0;
}