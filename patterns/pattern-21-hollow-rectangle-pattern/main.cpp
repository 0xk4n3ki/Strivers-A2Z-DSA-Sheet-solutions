#include <bits/stdc++.h>

class solution {
public:
    void pattern(int n) {
        for(int i=0; i < n; i++) {
            for(int j=0; j<n; j++) {
                if(i==0 || j==0 || i==n-1 || j==n-1) {
                    std::cout << "*";
                }else{
                    std::cout << " ";
                }
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