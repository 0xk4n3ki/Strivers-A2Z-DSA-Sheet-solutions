#include <bits/stdc++.h>

int main() {
    int n;
    std::cout << "Enter n: ";
    std::cin >> n;

    int reverse = 0;
    int tmp = n;
    while(tmp > 0) {
        int lastDigit = tmp % 10;
        reverse = reverse * 10 + lastDigit;
        tmp = tmp / 10;
    }

    if(reverse == n)
        std::cout << n << " is a palindrome" << std::endl;
    else
        std::cout << n << " is not a palindrome" << std::endl;

    return 0;
}