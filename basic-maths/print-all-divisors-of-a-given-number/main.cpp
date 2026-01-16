#include <bits/stdc++.h>

int main() {
    int n;
    std::cout << "Enter n: ";
    std::cin >> n;

    std::vector<int> vec;
    for(int i=1; i*i <= n; i++) {
        if(n%i==0){
            vec.push_back(i);
            if(n/i == i){
                continue;
            }
            vec.push_back(n/i);
        }
    }
    sort(vec.begin(), vec.end());

    std::cout << "divisors of " << n << " are: ";
    for(auto it = vec.begin(); it != vec.end(); it++) {
        std::cout << *it << " ";
    }
    std::cout << std::endl;
    return 0;
}