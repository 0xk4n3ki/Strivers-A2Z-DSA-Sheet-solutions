#include <iostream>

int main() {
    int n;
    std::cout << "Enter n: ";
    std::cin >> n;

    for(int i=0; i< n; i++) {
        for(int j=0; j <n-i-1; j++) {
            std::cout << " ";
        }
        char ch = 'A';
        for(int j=0; j<2*i+1; j++) {
            std::cout << ch;
            if(j>i+1){
                ch = ch-1;
            }else {
                ch++;
            }
        }
        std::cout << std::endl;
    }
    return 0;
}