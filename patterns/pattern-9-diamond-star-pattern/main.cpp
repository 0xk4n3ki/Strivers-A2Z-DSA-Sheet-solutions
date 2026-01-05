#include <iostream>

int main() {
    int n;
    std::cout << "Enter n: ";
    std::cin >> n;

    for(int i=1; i<=n; i++) {
        for(int j=1; j <= n-i; j++) {
            std::cout << " ";
        }
        for(int j=1; j<=2*i-1; j++) {
            std::cout << "*";
        }
        std::cout << std::endl;
    }

    for(int i=1; i<=n; i++) {
        for(int j=1; j < i; j++) {
            std::cout << " ";
        }
        for(int j=2*(n-i)+1; j > 0; j--) {
            std::cout << "*";
        }
        std::cout << std::endl;
    }
    return 0;
}