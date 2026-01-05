#include <iostream>

int main() {
    int n;
    std::cout << "Enter n: ";
    std::cin >> n;

    for(int i=0; i <n; i++) {
        int d = (i&1)^1;
        for(int j=0; j<=i; j++) {
            std::cout << d;
            d = d^1;
        }
        std::cout << std::endl;
    }
    return 0;
}