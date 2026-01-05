#include <iostream>

int main() {
    int n;
    std::cout << "Enter n: ";
    std::cin >> n;

    for(int i=0; i < 2*n-1; i++) {
        int limit;
        if(i<n) {
            limit = i+1;
        }else {
            limit = 2*n-i-1;
        }

        for(int j=0; j<limit; j++) {
            std::cout << "*";
        }
        std::cout << std::endl;
    }
    return 0;
}