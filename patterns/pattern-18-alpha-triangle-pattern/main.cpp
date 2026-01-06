#include <iostream>

int main() {
    int n;
    std::cout << "Enter n: ";
    std::cin >> n;

    for(int i=0; i< n; i++) {
        char c = 'A'+n-1-i;
        for(int j=0; j <= i; j++) {
            std::cout << c;
            c++;
        }
        std::cout << std::endl;
    }
    return 0;
}