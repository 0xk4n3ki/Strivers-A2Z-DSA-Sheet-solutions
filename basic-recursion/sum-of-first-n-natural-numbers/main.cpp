#include <bits/stdc++.h>

int naturalSum(int n) {
    if(n==1) {
        return 1;
    }
    return n + naturalSum(n-1);
}

int main() {
    int n;
    std::cout << "Enter n: ";
    std::cin >> n;

    int sum = naturalSum(n);
    std::cout << "Sum of first " << n << " natural numbers is: " << sum << std::endl;
    return 0;
}