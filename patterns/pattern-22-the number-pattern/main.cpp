#include <bits/stdc++.h>

int min(int a, int b) {
    if(a<b) {
        return a;
    }
    return b;
}

class solution{
public:
    void pattern(int n) {
        for(int i=0; i<2*n-1; i++) {
            for(int j=0; j<2*n-1; j++) {
                int top = i;
                int left = j;
                int bottom = 2*n-2-i;
                int right = 2*n-2-j;

                int minDist = min(min(top,bottom), min(left, right));
                std::cout << n-minDist;
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