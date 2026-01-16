#include <bits/stdc++.h>

int fibonacci(std::vector<int>& vec, int n) {
    if(n<=1)
        return vec[n] = n;

    if(vec[n] != -1)
        return vec[n];

    return vec[n] = fibonacci(vec, n-1) + fibonacci(vec, n-2);
}

void fib2(int a, int b, int n, std::vector<int>& vec) {
    if (n==0)
        return
    
    vec.push_back(a);
    fib2(b, a+b, n-1, vec);
}

int main() {
    int n;
    std::cout << "Enter n: ";
    std::cin >> n;

    std::vector<int> vec(n+1, -1);
    fibonacci(vec, n);
    for(auto it = vec.begin(); it != vec.end(); it++) {
        std::cout << *it << " ";
    }
    std::cout << std::endl;

    std::vector<int> vec2;
    fib2(0, 1, n, vec2);
    for(auto it = vec.begin(); it != vec.end(); it++) {
        std::cout << *it << " ";
    }
    std::cout << std::endl;
    return 0;
}

