#include <bits/stdc++.h>

class solution {
public:
    void pattern(int n) {
        for(int i=0; i<2*n-1; i++) {
            int stars = 0;
            if(i>=n){
                stars = 2*n-1-i;
            }else {
                stars = i+1;
            }

            for(int j=0; j<stars; j++) {
                std::cout << "*";
            }
            for(int j=0; j < 2*(n-stars); j++) {
                std::cout << " ";
            }
            for(int j=0; j<stars; j++) {
                std::cout << "*";
            }
            std::cout << std::endl;
        }
    }
};

int main() {
    int n;
    std::cout << "Enter n: ";
    std::cin >> n;

    solution sol;
    sol.pattern(n);

    return 0;
}