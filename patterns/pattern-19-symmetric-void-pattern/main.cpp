#include <bits/stdc++.h>

class solution {
public:
    void pattern(int n) {
        for(int i=0; i<2*n; i++) {
            int stars = 0;
            if(i<n) {
                stars = n-i;
            }else{
                stars = i-n+1;
            }

            for(int j=0; j<stars; j++) {
                std::cout << "*";
            }
            for(int j=0; j<2*(n-stars); j++) {
                std::cout << " ";
            }
            for(int j=0; j <stars; j++) {
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