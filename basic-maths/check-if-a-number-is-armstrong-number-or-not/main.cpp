#include <bits/stdc++.h>

int main() {
    int n;
    std::cout << "enter n: ";
    std::cin >> n;

    int tmp = n;
    int numDigits = 0;

    while(tmp>0) {
        numDigits++;
        tmp = tmp/10;
    }

    tmp = n;
    int sum = 0;

    while(tmp>0) {
        sum += pow((tmp%10), numDigits) ;
        tmp = tmp/10;
    }
    if(sum==n)
        std::cout << n << " is an armstrong number" << std::endl;
    else
        std::cout << n << " is not an armstrong number" << std::endl;

    return 0;
}