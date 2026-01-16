#include <bits/stdc++.h>

int main() {
    int n;
    std::cout << "Enter n: ";
    std::cin >> n;

    int ans = 0;
    int tmp = n;
    while(tmp > 0) {
        int digit = tmp % 10;
        ans = ans*10 + digit;
        tmp = tmp/10;
    }
    std::cout << "reverse of " << n << " is: " << ans << std::endl;

    return 0;
}