#include <bits/stdc++.h>

int main() {
    int a, b;
    std::cout << "Enter two number: ";
    std::cin >> a >> b;
    
    int n1 = a;
    int n2 = b;

    
    while(n1 > 0 && n2 > 0) {
        if(n1 > n2) {
            n1 = n1 % n2;
        }else {
            n2 = n2 % n1;
        }
    }

    if(a==0)
        std::cout << "GCD of " << a << " and " << b << " is: " << n2 << std::endl;
    else
        std::cout << "GCD of " << a << " and " << b << " is: " << n1 << std::endl;
    return 0;
}

