#include <bits/stdc++.h>

int main() {
    std::vector<int> v = {4, 2, 5, 9, 1};
    std::cout << "The elements in the vector are: ";
    for(int i=0; i<v.size(); i++) {
        std::cout << v[i] << " ";
    }
    std::cout << std::endl;

    std::cout << "The minimum element is: " << *std::min_element(v.begin(), v.end()) << std::endl;
}