#include <iostream>

int main() {
    int n;
    std::cout << "Enter n: ";
    std::cin >> n;

    for(int i=0; i<n; i++) {
        for(char j='A'; j <='A'+i; j++) {
            std::cout << j;
        }
        std::cout << std::endl;
    }
    return 0;
}